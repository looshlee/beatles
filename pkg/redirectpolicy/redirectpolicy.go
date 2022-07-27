//  Copyright 2020 Authors of Cilium
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package redirectpolicy

import (
	"fmt"
	"net"

	"github.com/cilium/cilium/pkg/k8s"
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/labels"
	k8sUtils "github.com/cilium/cilium/pkg/k8s/utils"
	"github.com/cilium/cilium/pkg/loadbalancer"
	lb "github.com/cilium/cilium/pkg/loadbalancer"
	"github.com/cilium/cilium/pkg/policy/api"

	"k8s.io/apimachinery/pkg/types"
)

type lrpConfigType = int

const (
	lrpConfigTypeNone = iota
	// LRP config is specified using IP/port tuple.
	lrpConfigTypeAddr
	// LRP config is specified using Kubernetes service name and namespace.
	lrpConfigTypeSvc
)

type frontendConfigType = int

const (
	frontendTypeUnknown = iota
	// Get frontends for service with clusterIP and all service ports.
	svcFrontendAll
	// Get frontends for service with clusterIP and parsed config named ports.
	svcFrontendNamedPorts
	// Get a frontend for service with clusterIP and parsed config port.
	svcFrontendSinglePort
	// Get a frontend with parsed config frontend IP and port.
	addrFrontendSinglePort
	// Get frontends with parsed config frontend IPs and named ports.
	addrFrontendNamedPorts
)

// LRPConfig is the internal representation of Cilium Local Redirect Policy.
type LRPConfig struct {
	// id is the parsed config name and namespace
	id k8s.ServiceID
	// uid is the unique identifier assigned by Kubernetes
	uid types.UID
	// lrpType is the type of either address matcher or service matcher policy
	lrpType lrpConfigType
	// frontendType is the type for the parsed config frontend.
	frontendType frontendConfigType
	// frontendMappings is a slice of policy config frontend mappings that include
	// frontend address, frontend port name, and a slice of its associated backends
	frontendMappings []*feMapping
	// serviceID is the parsed service name and namespace
	serviceID *k8s.ServiceID
	// backendSelector is an endpoint selector generated from the parsed policy selector
	backendSelector api.EndpointSelector
	// backendPorts is a slice of backend port and protocol along with the port name
	backendPorts []bePortInfo
	// backendPortsByPortName is a map indexed by port name with the value as
	// a pointer to bePortInfo for easy lookup into backendPorts
	backendPortsByPortName map[portName]*bePortInfo
}

type frontend = loadbalancer.L3n4Addr

type backend = loadbalancer.L3n4Addr

type portName = string

// feMapping stores frontend address and a list of associated backend addresses.
type feMapping struct {
	feAddr   *frontend
	backends []backend
	fePort   portName
}

type bePortInfo struct {
	// l4Addr is the port and protocol
	l4Addr lb.L4Addr
	// name is the port name
	name string
}

// PolicyID includes policy name and namespace
type policyID = k8s.ServiceID

// Parse parses the specified cilium local redirect policy spec, and returns
// a sanitized LRPConfig.
func Parse(clrp *v2.CiliumLocalRedirectPolicy, sanitize bool) (*LRPConfig, error) {
	name := clrp.ObjectMeta.Name
	if name == "" {
		return &LRPConfig{}, fmt.Errorf("CiliumLocalRedirectPolicy must have a name")
	}

	namespace := k8sUtils.ExtractNamespace(&clrp.ObjectMeta)
	if namespace == "" {
		// TODO CCLRP: addressed in follow-up PR. Details - GH-12831
		return &LRPConfig{}, fmt.Errorf("ClusterwideCiliumLocalRedirectPolicy is currently not supported")
	}

	if sanitize {
		return getSanitizedLRPConfig(name, namespace, clrp.UID, clrp.Spec)
	} else {
		return &LRPConfig{
			id: k8s.ServiceID{
				Name:      name,
				Namespace: namespace,
			},
		}, nil
	}
}

func getSanitizedLRPConfig(name, namespace string, uid types.UID, spec v2.CiliumLocalRedirectPolicySpec) (*LRPConfig, error) {

	var (
		addrMatcher    = spec.RedirectFrontend.AddressMatcher
		svcMatcher     = spec.RedirectFrontend.ServiceMatcher
		redirectTo     = spec.RedirectBackend
		frontendType   = frontendTypeUnknown
		checkNamedPort = false
		lrpType        lrpConfigType
		k8sSvc         *k8s.ServiceID
		fe             *frontend
		feMappings     []*feMapping
		bePorts        []bePortInfo
		bePortsMap     = make(map[portName]*bePortInfo)
	)

	// Parse frontend config
	switch {
	case addrMatcher == nil && svcMatcher == nil:
		return nil, fmt.Errorf("both address and service" +
			" matchers can not be empty")
	case addrMatcher != nil && svcMatcher != nil:
		return nil, fmt.Errorf("both address and service" +
			" matchers can not be specified")
	case addrMatcher != nil:
		// LRP specifies IP/port tuple config for traffic that needs to be redirected.
		ip := net.ParseIP(addrMatcher.IP)
		if ip == nil {
			return nil, fmt.Errorf("invalid address matcher IP %v",
				addrMatcher.IP)
		}
		if len(addrMatcher.ToPorts) > 1 {
			// If there are multiple ports, then the ports must be named.
			checkNamedPort = true
			frontendType = addrFrontendNamedPorts
		} else if len(addrMatcher.ToPorts) == 1 {
			frontendType = addrFrontendSinglePort
		}
		feMappings = make([]*feMapping, len(addrMatcher.ToPorts))
		for i, portInfo := range addrMatcher.ToPorts {
			p, pName, proto, err := portInfo.SanitizePortInfo(checkNamedPort)
			if err != nil {
				return nil, fmt.Errorf("invalid address matcher port %v", err)
			}
			// Set the scope to ScopeExternal as the externalTrafficPolicy is set to Cluster.
			fe = loadbalancer.NewL3n4Addr(proto, ip, p, loadbalancer.ScopeExternal)
			feM := &feMapping{
				feAddr: fe,
				fePort: pName,
			}
			feMappings[i] = feM
		}
		lrpType = lrpConfigTypeAddr
	case svcMatcher != nil:
		// LRP specifies service config for traffic that needs to be redirected.
		k8sSvc = &k8s.ServiceID{
			Name:      svcMatcher.Name,
			Namespace: svcMatcher.Namespace,
		}
		if k8sSvc.Name == "" {
			return nil, fmt.Errorf("kubernetes service name can" +
				"not be empty")
		}
		if namespace != "" && k8sSvc.Namespace != namespace {
			return nil, fmt.Errorf("kubernetes service namespace" +
				"does not match with the CiliumLocalRedirectPolicy namespace")
		}
		switch len(svcMatcher.ToPorts) {
		case 0:
			frontendType = svcFrontendAll
		case 1:
			frontendType = svcFrontendSinglePort
		default:
			frontendType = svcFrontendNamedPorts
			checkNamedPort = true
		}
		feMappings = make([]*feMapping, len(svcMatcher.ToPorts))
		for i, portInfo := range svcMatcher.ToPorts {
			p, pName, proto, err := portInfo.SanitizePortInfo(checkNamedPort)
			if err != nil {
				return nil, fmt.Errorf("invalid service matcher port %v", err)
			}
			// Set the scope to ScopeExternal as the externalTrafficPolicy is set to Cluster.
			// frontend ip will later be populated with the clusterIP of the service.
			fe = loadbalancer.NewL3n4Addr(proto, net.IP{}, p, loadbalancer.ScopeExternal)
			feM := &feMapping{
				feAddr: fe,
				fePort: pName,
			}
			feMappings[i] = feM
		}
		lrpType = lrpConfigTypeSvc
	default:
		return nil, fmt.Errorf("invalid local redirect policy %v", spec)
	}

	if lrpType == lrpConfigTypeNone || frontendType == frontendTypeUnknown {
		return nil, fmt.Errorf("invalid local redirect policy %v", spec)
	}

	// Parse backend config
	bePorts = make([]bePortInfo, len(redirectTo.ToPorts))
	if len(redirectTo.ToPorts) > 1 {
		// We check for backend named ports if either frontend/backend have
		// multiple ports.
		checkNamedPort = true
	}
	for i, portInfo := range redirectTo.ToPorts {
		p, pName, proto, err := portInfo.SanitizePortInfo(checkNamedPort)
		if err != nil {
			return nil, fmt.Errorf("invalid backend port %v", err)
		}
		beP := bePortInfo{
			l4Addr: lb.L4Addr{
				Protocol: proto,
				Port:     p,
			},
			name: pName,
		}
		bePorts[i] = beP
		if len(pName) > 0 {
			// Port name is present.
			bePortsMap[pName] = &bePorts[i]

		}
	}
	// When a single port is specified in the LRP frontend, the protocol for frontend and
	// backend must match.
	if len(feMappings) == 1 {
		if bePorts[0].l4Addr.Protocol != feMappings[0].feAddr.Protocol {
			return nil, fmt.Errorf("backend protocol must match with " +
				"frontend protocol")
		}
	}

	// Get an EndpointSelector from the passed policy labelSelector for optimized matching.
	selector := api.NewESFromK8sLabelSelector("", &redirectTo.LocalEndpointSelector)

	return &LRPConfig{
		uid:                    uid,
		serviceID:              k8sSvc,
		frontendMappings:       feMappings,
		backendSelector:        selector,
		backendPorts:           bePorts,
		backendPortsByPortName: bePortsMap,
		lrpType:                lrpType,
		frontendType:           frontendType,
		id: k8s.ServiceID{
			Name:      name,
			Namespace: namespace,
		},
	}, nil
}

// policyConfigSelectsPod determines if the given pod is selected by the policy
// config based on matching labels of config and pod.
func (config *LRPConfig) policyConfigSelectsPod(podInfo *podMetadata) bool {
	return config.backendSelector.Matches(labels.Set(podInfo.labels))
}

// checkNamespace returns true if config namespace matches with the given namespace.
// The namespace check isn't applicable for clusterwide LRPs.
func (config *LRPConfig) checkNamespace(namespace string) bool {
	if config.id.Namespace != "" {
		return namespace == config.id.Namespace
	}
	return true
}

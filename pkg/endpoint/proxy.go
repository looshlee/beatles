// Copyright 2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package endpoint

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cilium/cilium/pkg/completion"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/policy"
	"github.com/cilium/cilium/pkg/proxy/logger"
	"github.com/cilium/cilium/pkg/revert"
	"github.com/sirupsen/logrus"
)

// EndpointProxy defines any L7 proxy with which an Endpoint must interact.
type EndpointProxy interface {
	CreateOrUpdateRedirect(l4 policy.ProxyPolicy, id string, localEndpoint logger.EndpointUpdater, wg *completion.WaitGroup) (proxyPort uint16, err error, finalizeFunc revert.FinalizeFunc, revertFunc revert.RevertFunc)
	RemoveRedirect(id string, wg *completion.WaitGroup) (error, revert.FinalizeFunc, revert.RevertFunc)
	UpdateNetworkPolicy(ep logger.EndpointUpdater, policy *policy.L4Policy, ingressPolicyEnforced, egressPolicyEnforced bool, wg *completion.WaitGroup) (error, func() error)
	UseCurrentNetworkPolicy(ep logger.EndpointUpdater, policy *policy.L4Policy, wg *completion.WaitGroup)
	RemoveNetworkPolicy(ep logger.EndpointInfoSource)
}

// SetProxy sets the proxy for this endpoint.
func (e *Endpoint) SetProxy(p EndpointProxy) {
	e.unconditionalLock()
	defer e.unlock()
	e.proxy = p
}

// updateProxyRedirect updates the redirect rules in the proxy for a particular
// endpoint using the provided L4 filter. Returns the allocated proxy port
func (e *Endpoint) updateProxyRedirect(l4 *policy.L4Filter, proxyWaitGroup *completion.WaitGroup) (proxyPort uint16, err error, finalizeFunc revert.FinalizeFunc, revertFunc revert.RevertFunc) {
	if e.isProxyDisabled() {
		return 0, fmt.Errorf("can't redirect, proxy disabled"), nil, nil
	}
	return e.proxy.CreateOrUpdateRedirect(l4, e.ProxyID(l4), e, proxyWaitGroup)
}

// removeProxyRedirect removes a previously installed proxy redirect for an
// endpoint.
func (e *Endpoint) removeProxyRedirect(id string, proxyWaitGroup *completion.WaitGroup) (error, revert.FinalizeFunc, revert.RevertFunc) {
	if e.isProxyDisabled() {
		return nil, nil, nil
	}
	log.WithFields(logrus.Fields{
		logfields.EndpointID: e.ID,
		logfields.L4PolicyID: id,
	}).Debug("Removing redirect to endpoint")
	return e.proxy.RemoveRedirect(id, proxyWaitGroup)
}

func (e *Endpoint) removeNetworkPolicy() {
	if e.isProxyDisabled() {
		return
	}
	e.proxy.RemoveNetworkPolicy(e)
}

func (e *Endpoint) isProxyDisabled() bool {
	return e.proxy == nil || reflect.ValueOf(e.proxy).IsNil()
}

// OnProxyPolicyUpdate is a callback used to update the Endpoint's
// proxyPolicyRevision when the specified revision has been applied in the
// proxy.
func (e *Endpoint) OnProxyPolicyUpdate(revision uint64) {
	// NOTE: unconditionalLock is used here because this callback has no way of reporting an error
	e.unconditionalLock()
	if revision > e.proxyPolicyRevision {
		e.proxyPolicyRevision = revision
	}
	e.unlock()
}

// waitForProxyCompletions blocks until all proxy changes have been completed.
// Called with buildMutex held.
func (e *Endpoint) waitForProxyCompletions(proxyWaitGroup *completion.WaitGroup) error {
	if proxyWaitGroup == nil {
		return nil
	}

	err := proxyWaitGroup.Context().Err()
	if err != nil {
		return fmt.Errorf("context cancelled before waiting for proxy updates: %s", err)
	}

	start := time.Now()

	e.getLogger().Debug("Waiting for proxy updates to complete...")
	err = proxyWaitGroup.Wait()
	if err != nil {
		return fmt.Errorf("proxy state changes failed: %s", err)
	}
	e.getLogger().Debug("Wait time for proxy updates: ", time.Since(start))

	return nil
}

// GetProxyInfoByFields returns the ID, IPv4 address, IPv6 address, labels,
// SHA of labels, and identity of the endpoint. Returns an error if the endpoint
// is in the process of being deleted / has been deleted.
func (e *Endpoint) GetProxyInfoByFields() (uint64, string, string, []string, string, uint64, error) {
	// We use unconditional locking here because we explicitly handle state
	// in which the endpoint is being deleted.
	e.unconditionalRLock()
	defer e.runlock()
	var err error
	if e.IsDisconnecting() {
		err = fmt.Errorf("endpoint is in the process of being deleted")
	}
	return e.GetID(), e.GetIPv4Address(), e.GetIPv6Address(), e.GetLabels(), e.GetLabelsSHA(), uint64(e.GetIdentity()), err
}

// FakeEndpointProxy is a stub proxy used for testing.
type FakeEndpointProxy struct{}

// CreateOrUpdateRedirect does nothing.
func (f *FakeEndpointProxy) CreateOrUpdateRedirect(l4 policy.ProxyPolicy, id string, localEndpoint logger.EndpointUpdater, wg *completion.WaitGroup) (proxyPort uint16, err error, finalizeFunc revert.FinalizeFunc, revertFunc revert.RevertFunc) {
	return
}

// RemoveRedirect does nothing.
func (f *FakeEndpointProxy) RemoveRedirect(id string, wg *completion.WaitGroup) (error, revert.FinalizeFunc, revert.RevertFunc) {
	return nil, nil, nil
}

// UpdateNetworkPolicy does nothing.
func (f *FakeEndpointProxy) UpdateNetworkPolicy(ep logger.EndpointUpdater, policy *policy.L4Policy, ingressPolicyEnforced, egressPolicyEnforced bool, wg *completion.WaitGroup) (error, func() error) {
	return nil, nil
}

// UseCurrentNetworkPolicy does nothing.
func (f *FakeEndpointProxy) UseCurrentNetworkPolicy(ep logger.EndpointUpdater, policy *policy.L4Policy, wg *completion.WaitGroup) {
}

// RemoveNetworkPolicy does nothing.
func (f *FakeEndpointProxy) RemoveNetworkPolicy(ep logger.EndpointInfoSource) {}

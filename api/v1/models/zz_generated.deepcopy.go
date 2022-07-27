// +build !ignore_autogenerated

// Copyright 2017-2019 Authors of Cilium
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

// Code generated by deepcopy-gen. DO NOT EDIT.

package models

import (
	strfmt "github.com/go-openapi/strfmt"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CiliumHealth != nil {
		in, out := &in.CiliumHealth, &out.CiliumHealth
		*out = new(Status)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]*NodeElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatus) DeepCopyInto(out *ControllerStatus) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ControllerStatusConfiguration)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ControllerStatusStatus)
		(*in).DeepCopyInto(*out)
	}
	in.UUID.DeepCopyInto(&out.UUID)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatus.
func (in *ControllerStatus) DeepCopy() *ControllerStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatusStatus) DeepCopyInto(out *ControllerStatusStatus) {
	*out = *in
	in.LastFailureTimestamp.DeepCopyInto(&out.LastFailureTimestamp)
	in.LastSuccessTimestamp.DeepCopyInto(&out.LastSuccessTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatusStatus.
func (in *ControllerStatusStatus) DeepCopy() *ControllerStatusStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMStatus) DeepCopyInto(out *IPAMStatus) {
	*out = *in
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make(AllocationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IPV4 != nil {
		in, out := &in.IPV4, &out.IPV4
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPV6 != nil {
		in, out := &in.IPV6, &out.IPV6
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMStatus.
func (in *IPAMStatus) DeepCopy() *IPAMStatus {
	if in == nil {
		return nil
	}
	out := new(IPAMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sStatus) DeepCopyInto(out *K8sStatus) {
	*out = *in
	if in.K8sAPIVersions != nil {
		in, out := &in.K8sAPIVersions, &out.K8sAPIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sStatus.
func (in *K8sStatus) DeepCopy() *K8sStatus {
	if in == nil {
		return nil
	}
	out := new(K8sStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAddressing) DeepCopyInto(out *NodeAddressing) {
	*out = *in
	if in.IPV4 != nil {
		in, out := &in.IPV4, &out.IPV4
		*out = new(NodeAddressingElement)
		**out = **in
	}
	if in.IPV6 != nil {
		in, out := &in.IPV6, &out.IPV6
		*out = new(NodeAddressingElement)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAddressing.
func (in *NodeAddressing) DeepCopy() *NodeAddressing {
	if in == nil {
		return nil
	}
	out := new(NodeAddressing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeElement) DeepCopyInto(out *NodeElement) {
	*out = *in
	if in.HealthEndpointAddress != nil {
		in, out := &in.HealthEndpointAddress, &out.HealthEndpointAddress
		*out = new(NodeAddressing)
		(*in).DeepCopyInto(*out)
	}
	if in.PrimaryAddress != nil {
		in, out := &in.PrimaryAddress, &out.PrimaryAddress
		*out = new(NodeAddressing)
		(*in).DeepCopyInto(*out)
	}
	if in.SecondaryAddresses != nil {
		in, out := &in.SecondaryAddresses, &out.SecondaryAddresses
		*out = make([]*NodeAddressingElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeAddressingElement)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeElement.
func (in *NodeElement) DeepCopy() *NodeElement {
	if in == nil {
		return nil
	}
	out := new(NodeElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyStatistics) DeepCopyInto(out *ProxyStatistics) {
	*out = *in
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(RequestResponseStatistics)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStatistics.
func (in *ProxyStatistics) DeepCopy() *ProxyStatistics {
	if in == nil {
		return nil
	}
	out := new(ProxyStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestResponseStatistics) DeepCopyInto(out *RequestResponseStatistics) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(MessageForwardingStatistics)
		**out = **in
	}
	if in.Responses != nil {
		in, out := &in.Responses, &out.Responses
		*out = new(MessageForwardingStatistics)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestResponseStatistics.
func (in *RequestResponseStatistics) DeepCopy() *RequestResponseStatistics {
	if in == nil {
		return nil
	}
	out := new(RequestResponseStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusResponse) DeepCopyInto(out *StatusResponse) {
	*out = *in
	if in.Cilium != nil {
		in, out := &in.Cilium, &out.Cilium
		*out = new(Status)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(ClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(Status)
		**out = **in
	}
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make(ControllerStatuses, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ControllerStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IPAM != nil {
		in, out := &in.IPAM, &out.IPAM
		*out = new(IPAMStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(K8sStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Kvstore != nil {
		in, out := &in.Kvstore, &out.Kvstore
		*out = new(Status)
		**out = **in
	}
	if in.NodeMonitor != nil {
		in, out := &in.NodeMonitor, &out.NodeMonitor
		*out = new(MonitorStatus)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(ProxyStatus)
		**out = **in
	}
	if in.Stale != nil {
		in, out := &in.Stale, &out.Stale
		*out = make(map[string]strfmt.DateTime, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusResponse.
func (in *StatusResponse) DeepCopy() *StatusResponse {
	if in == nil {
		return nil
	}
	out := new(StatusResponse)
	in.DeepCopyInto(out)
	return out
}

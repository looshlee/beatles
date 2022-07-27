// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/tcp_proxy/v3/tcp_proxy.proto

package envoy_extensions_filters_network_tcp_proxy_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/accesslog/v3"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v32 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#next-free-field: 13]
type TcpProxy struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_tcp_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria. Only endpoints in the upstream
	// cluster with metadata matching that set in metadata_match will be
	// considered. The filter name should be specified as *envoy.lb*.
	MetadataMatch *v3.Metadata `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// The idle timeout for connections managed by the TCP proxy filter. The idle timeout
	// is defined as the period in which there are no bytes sent or received on either
	// the upstream or downstream connection. If not set, the default idle timeout is 1 hour. If set
	// to 0s, the timeout will be disabled.
	//
	// .. warning::
	//   Disabling this timeout has a highly likelihood of yielding connection leaks due to lost TCP
	//   FIN packets, etc.
	IdleTimeout *duration.Duration `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// [#not-implemented-hide:] The idle timeout for connections managed by the TCP proxy
	// filter. The idle timeout is defined as the period in which there is no
	// active traffic. If not set, there is no idle timeout. When the idle timeout
	// is reached the connection will be closed. The distinction between
	// downstream_idle_timeout/upstream_idle_timeout provides a means to set
	// timeout based on the last byte sent on the downstream/upstream connection.
	DownstreamIdleTimeout *duration.Duration `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	// [#not-implemented-hide:]
	UpstreamIdleTimeout *duration.Duration `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	// Configuration for :ref:`access logs <arch_overview_access_logs>`
	// emitted by the this tcp_proxy.
	AccessLog []*v31.AccessLog `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// The maximum number of unsuccessful connection attempts that will be made before
	// giving up. If the parameter is not specified, 1 connection attempt will be made.
	MaxConnectAttempts *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	// Optional configuration for TCP proxy hash policy. If hash_policy is not set, the hash-based
	// load balancing algorithms will select a host randomly. Currently the number of hash policies is
	// limited to 1.
	HashPolicy []*v32.HashPolicy `protobuf:"bytes,11,rep,name=hash_policy,json=hashPolicy,proto3" json:"hash_policy,omitempty"`
	// [#not-implemented-hide:] feature in progress
	// If set, this configures tunneling, e.g. configuration options to tunnel multiple TCP
	// payloads over a shared HTTP/2 tunnel. If this message is absent, the payload
	// will be proxied upstream as per usual.
	TunnelingConfig      *TcpProxy_TunnelingConfig `protobuf:"bytes,12,opt,name=tunneling_config,json=tunnelingConfig,proto3" json:"tunneling_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *v3.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v31.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

func (m *TcpProxy) GetHashPolicy() []*v32.HashPolicy {
	if m != nil {
		return m.HashPolicy
	}
	return nil
}

func (m *TcpProxy) GetTunnelingConfig() *TcpProxy_TunnelingConfig {
	if m != nil {
		return m.TunnelingConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// Allows for specification of multiple upstream clusters along with weights
// that indicate the percentage of traffic to be forwarded to each cluster.
// The router selects an upstream cluster based on these weights.
type TcpProxy_WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 0}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is
	// determined by its weight. The sum of weights across all entries in the
	// clusters array determines the total weight.
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints
	// in the upstream cluster with metadata matching what is set in this field will be considered
	// for load balancing. Note that this will be merged with what's provided in
	// :ref:`TcpProxy.metadata_match
	// <envoy_api_field_extensions.filters.network.tcp_proxy.v3.TcpProxy.metadata_match>`, with values
	// here taking precedence. The filter name should be specified as *envoy.lb*.
	MetadataMatch        *v3.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 0, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetMetadataMatch() *v3.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

// Configuration for tunneling TCP over other transports or application layers.
// Currently, only HTTP/2 is supported. When other options exist, HTTP/2 will
// remain the default.
type TcpProxy_TunnelingConfig struct {
	// The hostname to send in the synthesized CONNECT headers to the upstream proxy.
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_TunnelingConfig) Reset()         { *m = TcpProxy_TunnelingConfig{} }
func (m *TcpProxy_TunnelingConfig) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_TunnelingConfig) ProtoMessage()    {}
func (*TcpProxy_TunnelingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_926036763459e38f, []int{0, 1}
}

func (m *TcpProxy_TunnelingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Unmarshal(m, b)
}
func (m *TcpProxy_TunnelingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Marshal(b, m, deterministic)
}
func (m *TcpProxy_TunnelingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_TunnelingConfig.Merge(m, src)
}
func (m *TcpProxy_TunnelingConfig) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_TunnelingConfig.Size(m)
}
func (m *TcpProxy_TunnelingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_TunnelingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_TunnelingConfig proto.InternalMessageInfo

func (m *TcpProxy_TunnelingConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.WeightedCluster.ClusterWeight")
	proto.RegisterType((*TcpProxy_TunnelingConfig)(nil), "envoy.extensions.filters.network.tcp_proxy.v3.TcpProxy.TunnelingConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/tcp_proxy/v3/tcp_proxy.proto", fileDescriptor_926036763459e38f)
}

var fileDescriptor_926036763459e38f = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0xc7, 0xaf, 0x93, 0xdc, 0x24, 0x9d, 0xdc, 0xd0, 0xdc, 0x81, 0xab, 0xeb, 0x06, 0x68, 0xc3,
	0xc7, 0x22, 0x20, 0x61, 0xab, 0xc9, 0x06, 0x15, 0x58, 0xd4, 0x29, 0x22, 0x45, 0x54, 0xa4, 0x56,
	0x3f, 0x96, 0xd6, 0xd4, 0x9e, 0x38, 0x16, 0xf6, 0x8c, 0xe5, 0x19, 0x3b, 0xa9, 0xd8, 0x74, 0xc9,
	0x9a, 0x25, 0x8f, 0xc2, 0x1e, 0xc1, 0x96, 0xf7, 0xe0, 0x01, 0x50, 0x25, 0x24, 0x34, 0x33, 0x76,
	0x52, 0xa7, 0x41, 0xa5, 0x94, 0x55, 0x3c, 0xe7, 0xe3, 0x37, 0x67, 0xce, 0x39, 0xff, 0x80, 0x2f,
	0x30, 0xc9, 0xe8, 0xb5, 0x89, 0x17, 0x1c, 0x13, 0x16, 0x50, 0xc2, 0xcc, 0x69, 0x10, 0x72, 0x9c,
	0x30, 0x93, 0x60, 0x3e, 0xa7, 0xc9, 0x77, 0x26, 0x77, 0x63, 0x27, 0x4e, 0xe8, 0xe2, 0xda, 0xcc,
	0x86, 0xab, 0x83, 0x11, 0x27, 0x94, 0x53, 0xf8, 0x89, 0x4c, 0x37, 0x56, 0xe9, 0x46, 0x9e, 0x6e,
	0xe4, 0xe9, 0xc6, 0x2a, 0x23, 0x1b, 0x76, 0x3f, 0x52, 0xb7, 0xb9, 0x94, 0x4c, 0x03, 0xdf, 0x44,
	0xae, 0x8b, 0x19, 0x0b, 0xa9, 0x2f, 0xc8, 0xcb, 0x83, 0x22, 0x77, 0xf7, 0x4a, 0xa1, 0x2e, 0x4d,
	0xb0, 0x88, 0xba, 0x42, 0x0c, 0x97, 0x03, 0xf8, 0x75, 0x2c, 0x3d, 0x33, 0xc4, 0x66, 0x4e, 0x4c,
	0xc3, 0xc0, 0xcd, 0x6b, 0xeb, 0xee, 0xfa, 0x94, 0xfa, 0x21, 0x36, 0xe5, 0xe9, 0x2a, 0x9d, 0x9a,
	0x5e, 0x9a, 0x20, 0x1e, 0x50, 0xf2, 0x4f, 0xfe, 0x79, 0x82, 0xe2, 0x58, 0xd4, 0xae, 0xfc, 0xef,
	0xa6, 0x5e, 0x8c, 0x4c, 0x44, 0x08, 0xe5, 0x32, 0x8d, 0x99, 0x8c, 0x23, 0x9e, 0x16, 0xee, 0xf7,
	0xee, 0xb9, 0x33, 0x9c, 0x88, 0x1e, 0x04, 0xa4, 0x78, 0xc3, 0xeb, 0x0c, 0x85, 0x81, 0x87, 0x38,
	0x36, 0x8b, 0x0f, 0xe5, 0x78, 0xff, 0xaf, 0x16, 0x68, 0x9e, 0xb9, 0xf1, 0x44, 0xf4, 0x05, 0xf6,
	0x41, 0x4b, 0x80, 0x9d, 0x38, 0xc1, 0xd3, 0x60, 0xa1, 0x6b, 0x3d, 0xad, 0xbf, 0x65, 0x35, 0x6e,
	0xad, 0x5a, 0x52, 0xe9, 0x69, 0x36, 0x10, 0xbe, 0x89, 0x74, 0xc1, 0x2e, 0x68, 0xb8, 0x61, 0xca,
	0x38, 0x4e, 0xf4, 0x8a, 0x88, 0x1a, 0x3f, 0xb3, 0x0b, 0x03, 0xcc, 0xc0, 0xcb, 0x39, 0x0e, 0xfc,
	0x19, 0xc7, 0x9e, 0x93, 0xdb, 0x98, 0x0e, 0x7a, 0x5a, 0xbf, 0x35, 0xf8, 0xca, 0x78, 0xd4, 0x94,
	0x8c, 0xa2, 0x32, 0xe3, 0x32, 0x07, 0x8e, 0x14, 0x6f, 0xfc, 0xcc, 0xee, 0xcc, 0xcb, 0x26, 0x06,
	0xbf, 0x04, 0x6f, 0x44, 0x98, 0x23, 0x0f, 0x71, 0xe4, 0x44, 0x88, 0xbb, 0x33, 0x7d, 0x4b, 0x5e,
	0xba, 0x9b, 0x5f, 0xaa, 0x06, 0x68, 0x88, 0x01, 0x0a, 0xf6, 0x49, 0x1e, 0x6b, 0xb7, 0x8b, 0xac,
	0x13, 0x91, 0x04, 0x3f, 0x07, 0x2f, 0x02, 0x2f, 0xc4, 0x0e, 0x0f, 0x22, 0x4c, 0x53, 0xae, 0x37,
	0x25, 0x64, 0xc7, 0x50, 0x33, 0x32, 0x8a, 0x19, 0x19, 0x47, 0xf9, 0x0c, 0xed, 0x96, 0x08, 0x3f,
	0x53, 0xd1, 0xf0, 0x14, 0xbc, 0xf6, 0xe8, 0x9c, 0x30, 0x9e, 0x60, 0x14, 0x39, 0x25, 0x50, 0xf5,
	0x21, 0xd0, 0xab, 0x55, 0xe6, 0xf1, 0x1d, 0xe4, 0x09, 0x78, 0x95, 0xc6, 0x9b, 0x80, 0xb5, 0x87,
	0x80, 0x6f, 0x16, 0x79, 0x77, 0x71, 0x23, 0x00, 0xd4, 0x86, 0x3b, 0x21, 0xf5, 0xf5, 0xe7, 0xbd,
	0x6a, 0xbf, 0x35, 0xf8, 0xb0, 0xdc, 0xa2, 0x95, 0x02, 0xb2, 0xa1, 0x71, 0x28, 0x0f, 0xdf, 0x50,
	0xdf, 0xde, 0x42, 0xc5, 0x27, 0xbc, 0x04, 0x6f, 0x45, 0x68, 0xe1, 0xb8, 0x94, 0x10, 0xec, 0x72,
	0x07, 0x71, 0x8e, 0xa3, 0x98, 0x33, 0xbd, 0x21, 0x4b, 0x7a, 0xe7, 0x5e, 0x49, 0xe7, 0xc7, 0x84,
	0x0f, 0x07, 0x17, 0x28, 0x4c, 0xb1, 0x5c, 0xa8, 0x8f, 0x2b, 0x7d, 0xcd, 0x86, 0x11, 0x5a, 0x8c,
	0x14, 0xe1, 0x30, 0x07, 0xc0, 0x23, 0xd0, 0xba, 0xa3, 0x1f, 0xbd, 0x25, 0xcb, 0xdb, 0xc9, 0xcb,
	0x13, 0x0a, 0x13, 0x25, 0x8d, 0x11, 0x9b, 0x4d, 0x64, 0x80, 0xd5, 0xbc, 0xb5, 0x9e, 0xff, 0xa8,
	0x55, 0x3a, 0x9a, 0x0d, 0x66, 0x4b, 0x2b, 0x4c, 0x40, 0x87, 0xa7, 0x84, 0xe0, 0x30, 0x20, 0xbe,
	0xa3, 0x1e, 0xa5, 0xbf, 0x78, 0xda, 0x06, 0x9e, 0x15, 0xbc, 0x91, 0xc4, 0xd9, 0xdb, 0xbc, 0x6c,
	0xe8, 0xfe, 0x5a, 0x05, 0xdb, 0x6b, 0x6b, 0x0a, 0xbf, 0x07, 0xcd, 0xa5, 0x02, 0x34, 0xf9, 0x94,
	0xf3, 0xff, 0x49, 0x01, 0x46, 0xfe, 0xab, 0xcc, 0x45, 0x1b, 0x9a, 0x9a, 0xbd, 0xbc, 0xb0, 0xfb,
	0x87, 0x06, 0xda, 0xa5, 0x28, 0xf8, 0x36, 0xa8, 0x11, 0x14, 0xe1, 0x75, 0x61, 0x4b, 0x23, 0xdc,
	0x03, 0x75, 0x25, 0x29, 0xa9, 0xe8, 0xf6, 0x6a, 0x4c, 0xb9, 0x79, 0x83, 0xbe, 0xaa, 0xff, 0x41,
	0x5f, 0x07, 0x17, 0x3f, 0xfd, 0xf2, 0xc3, 0xee, 0x29, 0xf8, 0xb6, 0x94, 0xa4, 0x7a, 0xb0, 0xa9,
	0x05, 0x83, 0x7f, 0xd9, 0x82, 0x83, 0xb1, 0xe0, 0x8e, 0xc0, 0xe1, 0x93, 0xb9, 0xdd, 0x1b, 0x0d,
	0x6c, 0xaf, 0x8d, 0x1b, 0x7e, 0x00, 0x9a, 0x33, 0xca, 0xf8, 0xa6, 0xf6, 0x2d, 0x1d, 0x4f, 0x28,
	0x61, 0xed, 0xba, 0x83, 0x4f, 0x05, 0x69, 0x08, 0xf6, 0x1f, 0x4d, 0xb2, 0x74, 0xf0, 0x32, 0xdf,
	0x00, 0x87, 0xc5, 0xd8, 0x0d, 0xa6, 0x01, 0x4e, 0x60, 0xf5, 0x4f, 0x4b, 0xfb, 0xba, 0xd6, 0xac,
	0x77, 0x1a, 0x76, 0xdb, 0xc3, 0x71, 0x82, 0x5d, 0x24, 0xfe, 0x9d, 0xb3, 0x7d, 0xeb, 0xe2, 0xe7,
	0x9b, 0xdf, 0x7e, 0xaf, 0x57, 0x3a, 0x15, 0xf0, 0x59, 0x40, 0xd5, 0x20, 0x15, 0xf3, 0x51, 0x6b,
	0x6a, 0xb5, 0x8b, 0xfb, 0x27, 0x42, 0xf1, 0x13, 0xed, 0xaa, 0x2e, 0xa5, 0x3f, 0xfc, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x67, 0xac, 0xa5, 0x1b, 0xd6, 0x07, 0x00, 0x00,
}

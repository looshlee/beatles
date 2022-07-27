// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/lds.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	listener "github.com/cilium/proxy/go/envoy/api/v2/listener"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Listener_DrainType int32

const (
	// Drain in response to calling /healthcheck/fail admin endpoint (along with the health check
	// filter), listener removal/modification, and hot restart.
	Listener_DEFAULT Listener_DrainType = 0
	// Drain in response to listener removal/modification and hot restart. This setting does not
	// include /healthcheck/fail. This setting may be desirable if Envoy is hosting both ingress
	// and egress listeners.
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

type Listener struct {
	// The unique name by which this listener is known. If no name is provided,
	// Envoy will allocate an internal UUID for the listener. If the listener is to be dynamically
	// updated or removed via :ref:`LDS <config_listeners_lds>` a unique name must be provided.
	// By default, the maximum length of a listener's name is limited to 60 characters. This limit can
	// be increased by setting the :option:`--max-obj-name-len` command line argument to the desired
	// value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address that the listener should listen on. In general, the address must be unique, though
	// that is governed by the bind rules of the OS. E.g., multiple listeners can listen on port 0 on
	// Linux as the actual port will be allocated by the OS.
	Address *core.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// A list of filter chains to consider for this listener. The
	// :ref:`FilterChain <envoy_api_msg_listener.FilterChain>` with the most specific
	// :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` criteria is used on a
	// connection.
	//
	// Example using SNI for filter chain selection can be found in the
	// :ref:`FAQ entry <faq_how_to_setup_sni>`.
	FilterChains []*listener.FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	// If a connection is redirected using *iptables*, the port on which the proxy
	// receives it might be different from the original destination address. When this flag is set to
	// true, the listener hands off redirected connections to the listener associated with the
	// original destination address. If there is no listener associated with the original destination
	// address, the connection is handled by the listener that receives it. Defaults to false.
	//
	// .. attention::
	//
	//   This field is deprecated. Use :ref:`an original_dst <config_listener_filters_original_dst>`
	//   :ref:`listener filter <envoy_api_field_Listener.listener_filters>` instead.
	//
	//   Note that hand off to another listener is *NOT* performed without this flag. Once
	//   :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` is implemented this flag
	//   will be removed, as filter chain matching can be used to select a filter chain based on the
	//   restored destination address.
	UseOriginalDst *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst,proto3" json:"use_original_dst,omitempty"` // Deprecated: Do not use.
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Listener metadata.
	Metadata *core.Metadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	// Listener filters have the opportunity to manipulate and augment the connection metadata that
	// is used in connection filter chain matching, for example. These filters are run before any in
	// :ref:`filter_chains <envoy_api_field_Listener.filter_chains>`. Order matters as the
	// filters are processed sequentially right after a socket has been accepted by the listener, and
	// before a connection is created.
	ListenerFilters []*listener.ListenerFilter `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	// Whether the listener should be set as a transparent socket.
	// When this flag is set to true, connections can be redirected to the listener using an
	// *iptables* *TPROXY* target, in which case the original source and destination addresses and
	// ports are preserved on accepted connections. This flag should be used in combination with
	// :ref:`an original_dst <config_listener_filters_original_dst>` :ref:`listener filter
	// <envoy_api_field_Listener.listener_filters>` to mark the connections' local addresses as
	// "restored." This can be used to hand off each redirected connection to another listener
	// associated with the connection's destination address. Direct connections to the socket without
	// using *TPROXY* cannot be distinguished from connections redirected using *TPROXY* and are
	// therefore treated as if they were redirected.
	// When this flag is set to false, the listener's socket is explicitly reset as non-transparent.
	// Setting this flag requires Envoy to run with the *CAP_NET_ADMIN* capability.
	// When this flag is not set (default), the socket is not modified, i.e. the transparent option
	// is neither set nor reset.
	Transparent *wrappers.BoolValue `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	// Whether the listener should set the *IP_FREEBIND* socket option. When this
	// flag is set to true, listeners can be bound to an IP address that is not
	// configured on the system running Envoy. When this flag is set to false, the
	// option *IP_FREEBIND* is disabled on the socket. When this flag is not set
	// (default), the socket is not modified, i.e. the option is neither enabled
	// nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions []*core.SocketOption `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	// Whether the listener should accept TCP Fast Open (TFO) connections.
	// When this flag is set to a value greater than 0, the option TCP_FASTOPEN is enabled on
	// the socket, with a queue length of the specified size
	// (see `details in RFC7413 <https://tools.ietf.org/html/rfc7413#section-5.1>`_).
	// When this flag is set to 0, the option TCP_FASTOPEN is disabled on the socket.
	// When this flag is not set (default), the socket is not modified,
	// i.e. the option is neither enabled nor disabled.
	//
	// On Linux, the net.ipv4.tcp_fastopen kernel parameter must include flag 0x2 to enable
	// TCP_FASTOPEN.
	// See `ip-sysctl.txt <https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt>`_.
	//
	// On macOS, only values of 0, 1, and unset are valid; other values may result in an error.
	// To set the queue length on macOS, set the net.inet.tcp.fastopen_backlog kernel parameter.
	TcpFastOpenQueueLength *wrappers.UInt32Value `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

// Deprecated: Do not use.
func (m *Listener) GetUseOriginalDst() *wrappers.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*core.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

// [#not-implemented-hide:]
type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn’t
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// [V2-API-DIFF] This is deprecated in v2, all Listeners will bind to their
	// port. An additional filter chain must be created for every original
	// destination port this listener may redirect to in v2, with the original
	// port specified in the FilterChainMatch destination_port field.
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
}

func init() { proto.RegisterFile("envoy/api/v2/lds.proto", fileDescriptor_34e2cd84a105bcd1) }

var fileDescriptor_34e2cd84a105bcd1 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x8f, 0xdb, 0x44,
	0x18, 0xc6, 0x77, 0xb2, 0xdb, 0x26, 0x99, 0xfc, 0xd9, 0x68, 0x84, 0x5a, 0x2b, 0x2c, 0x9b, 0x10,
	0x40, 0x0a, 0x1c, 0x1c, 0x9a, 0x4a, 0x20, 0x55, 0x95, 0x50, 0xd3, 0x10, 0xb5, 0x52, 0x4a, 0xc0,
	0xbb, 0x5d, 0xba, 0xa7, 0xd1, 0xc4, 0x7e, 0x9d, 0xb5, 0x70, 0x66, 0x66, 0x67, 0xc6, 0x41, 0xb9,
	0x72, 0x42, 0x1c, 0xe1, 0xca, 0x07, 0xe0, 0x33, 0x70, 0xe2, 0xc8, 0x9d, 0x3b, 0x07, 0xc4, 0x05,
	0xf1, 0x25, 0x90, 0x27, 0xb6, 0x49, 0xb4, 0xbb, 0x2c, 0x87, 0xde, 0xde, 0x99, 0xf7, 0xf7, 0x3e,
	0x1e, 0x3f, 0x8f, 0x3d, 0xf8, 0x1e, 0xf0, 0x95, 0x58, 0x0f, 0x98, 0x8c, 0x06, 0xab, 0xe1, 0x20,
	0x0e, 0xb4, 0x2b, 0x95, 0x30, 0x82, 0xd4, 0xed, 0xbe, 0xcb, 0x64, 0xe4, 0xae, 0x86, 0xed, 0xce,
	0x0e, 0xe5, 0x0b, 0x05, 0x03, 0x16, 0x04, 0x0a, 0x74, 0x86, 0xb7, 0x8f, 0xae, 0x02, 0x73, 0xa6,
	0xe1, 0xda, 0x6e, 0x10, 0x69, 0x5f, 0xac, 0x40, 0xad, 0xb3, 0xee, 0xbb, 0xbb, 0x47, 0x88, 0xb4,
	0x01, 0x0e, 0xaa, 0x28, 0x72, 0x8d, 0x85, 0x10, 0x8b, 0x18, 0x2c, 0xc6, 0x38, 0x17, 0x86, 0x99,
	0x48, 0xf0, 0xfc, 0xf9, 0xc7, 0x59, 0xd7, 0xae, 0xe6, 0x49, 0x38, 0xf8, 0x5a, 0x31, 0x29, 0x41,
	0xe5, 0xfd, 0xfb, 0x2b, 0x16, 0x47, 0x01, 0x33, 0x30, 0xc8, 0x8b, 0xac, 0xf1, 0xc6, 0x42, 0x2c,
	0x84, 0x2d, 0x07, 0x69, 0xb5, 0xd9, 0xed, 0xfd, 0x58, 0xc1, 0x95, 0x69, 0xf6, 0x7c, 0x42, 0xf0,
	0x01, 0x67, 0x4b, 0x70, 0x50, 0x17, 0xf5, 0xab, 0x9e, 0xad, 0xc9, 0x18, 0x97, 0x33, 0x03, 0x9c,
	0x52, 0x17, 0xf5, 0x6b, 0xc3, 0xb6, 0xbb, 0x6d, 0x98, 0x9b, 0x3a, 0xe0, 0x3e, 0xd9, 0x10, 0xa3,
	0xe6, 0xaf, 0xbf, 0x77, 0xf6, 0x7e, 0xfe, 0xeb, 0x97, 0xfd, 0x3b, 0xdf, 0xa1, 0x52, 0x0b, 0x79,
	0xf9, 0x28, 0xf9, 0x12, 0x37, 0xc2, 0x28, 0x36, 0xa0, 0xa8, 0x7f, 0xc1, 0x22, 0xae, 0x9d, 0xfd,
	0xee, 0x7e, 0xbf, 0x36, 0xec, 0xed, 0x6a, 0x15, 0x46, 0x4c, 0x2c, 0xfb, 0x34, 0x45, 0xb7, 0x34,
	0xbf, 0x47, 0xa5, 0x0a, 0xf2, 0xea, 0xe1, 0xbf, 0x4d, 0x4d, 0x9e, 0xe1, 0x56, 0xa2, 0x81, 0x0a,
	0x15, 0x2d, 0x22, 0xce, 0x62, 0x1a, 0x68, 0xe3, 0x1c, 0x64, 0xe7, 0xdc, 0x38, 0xe5, 0xe6, 0x4e,
	0xb9, 0x23, 0x21, 0xe2, 0x33, 0x16, 0x27, 0x30, 0x2a, 0x39, 0xc8, 0x6b, 0x26, 0x1a, 0x66, 0xd9,
	0xd8, 0x58, 0x1b, 0x12, 0xe2, 0xb7, 0x65, 0x7a, 0x3e, 0xc1, 0x39, 0xf8, 0xa9, 0xe3, 0x74, 0x9e,
	0x84, 0x21, 0x28, 0x1a, 0x47, 0xcb, 0xc8, 0xd0, 0xf9, 0xda, 0x80, 0x76, 0xee, 0x58, 0xe9, 0xa3,
	0x2b, 0xd2, 0x2f, 0x9f, 0x73, 0xf3, 0x70, 0x68, 0xc5, 0xbd, 0xb7, 0x24, 0xa8, 0xa7, 0x85, 0xca,
	0xc8, 0x8a, 0x4c, 0x53, 0x8d, 0x51, 0x2a, 0x41, 0x3e, 0xc6, 0x95, 0x25, 0x18, 0x16, 0x30, 0xc3,
	0x9c, 0xbb, 0x56, 0xee, 0xcd, 0x6b, 0x1c, 0x7d, 0x91, 0x21, 0x5e, 0x01, 0x93, 0x67, 0xb8, 0x11,
	0x80, 0x54, 0xe0, 0x33, 0x03, 0x01, 0x5d, 0x3d, 0x70, 0xca, 0x76, 0xfa, 0x9d, 0xdd, 0xe9, 0x3c,
	0x4c, 0x77, 0x5c, 0xb0, 0x67, 0x0f, 0xbc, 0x7a, 0xb0, 0xb5, 0x22, 0x9f, 0x60, 0x1c, 0x28, 0x16,
	0x71, 0x6a, 0xd6, 0x12, 0x9c, 0x4a, 0x17, 0xf5, 0x9b, 0xc3, 0xee, 0x4d, 0x32, 0x29, 0x78, 0xba,
	0x96, 0xe0, 0x55, 0x83, 0xbc, 0x24, 0x67, 0xb8, 0x95, 0x67, 0x45, 0x37, 0x71, 0x68, 0xa7, 0x6a,
	0x13, 0x7d, 0xef, 0x86, 0x44, 0x73, 0xbd, 0x4d, 0xb2, 0xa3, 0x83, 0x34, 0x54, 0xef, 0x30, 0xde,
	0xd9, 0xd5, 0xe4, 0x31, 0xae, 0x19, 0xc5, 0xb8, 0x96, 0x4c, 0x01, 0x37, 0x0e, 0xbe, 0x2d, 0x48,
	0x6f, 0x1b, 0x27, 0x1f, 0xe1, 0x4a, 0xa8, 0x00, 0xe6, 0x11, 0x0f, 0x9c, 0xda, 0xad, 0xa3, 0x05,
	0x4b, 0x26, 0xb8, 0xa9, 0x85, 0xff, 0x15, 0x18, 0x2a, 0xa4, 0xfd, 0xd5, 0x9c, 0x86, 0x7d, 0x97,
	0xce, 0x35, 0xb9, 0x9c, 0x58, 0x70, 0x66, 0x39, 0xaf, 0xa1, 0xb7, 0x56, 0x9a, 0xbc, 0xc2, 0x6d,
	0xe3, 0x4b, 0x1a, 0x32, 0x9d, 0x2a, 0x01, 0xa7, 0x97, 0x09, 0x24, 0x40, 0x63, 0xe0, 0x0b, 0x73,
	0xe1, 0xd4, 0xff, 0xc7, 0xa7, 0x73, 0xcf, 0xf8, 0x72, 0xc2, 0xb4, 0x99, 0x49, 0xe0, 0x5f, 0xa4,
	0xc3, 0x53, 0x3b, 0xdb, 0x9e, 0xe2, 0xfa, 0x76, 0x9c, 0xe4, 0x31, 0xae, 0xa7, 0x27, 0xa7, 0x46,
	0x50, 0x29, 0x94, 0xb1, 0x3f, 0xec, 0x7f, 0xbf, 0x2d, 0x4e, 0xf9, 0x53, 0xf1, 0xb9, 0x50, 0xa6,
	0xf7, 0x3e, 0xae, 0x16, 0xa9, 0x92, 0x1a, 0x2e, 0x8f, 0x3f, 0x9d, 0x3c, 0x79, 0x39, 0x3d, 0x6d,
	0xed, 0x91, 0x43, 0x5c, 0x7b, 0x31, 0x1b, 0x3f, 0x9f, 0x9c, 0xd3, 0xd9, 0x67, 0xd3, 0xf3, 0x16,
	0x1a, 0xfe, 0x8d, 0xb0, 0x93, 0x47, 0x37, 0xce, 0x6f, 0xb3, 0x13, 0x50, 0xab, 0xc8, 0x07, 0xf2,
	0x0a, 0x1f, 0x9e, 0x18, 0x05, 0x6c, 0x99, 0x13, 0x9a, 0x1c, 0xef, 0x5a, 0x56, 0x8c, 0x78, 0x70,
	0x99, 0x80, 0x36, 0xed, 0xce, 0x8d, 0x7d, 0x2d, 0x05, 0xd7, 0xd0, 0xdb, 0xeb, 0xa3, 0x0f, 0x11,
	0x49, 0x70, 0x73, 0x02, 0xc6, 0xbf, 0x78, 0x8d, 0xc2, 0xbd, 0x6f, 0x7e, 0xfb, 0xf3, 0x87, 0xd2,
	0x51, 0xef, 0xfe, 0xce, 0xc5, 0xfc, 0x28, 0xff, 0x02, 0xf5, 0x23, 0xf4, 0xc1, 0xa8, 0xfc, 0xd3,
	0x1f, 0xc7, 0xe8, 0x5b, 0x84, 0xe6, 0x77, 0xad, 0x83, 0x0f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x2b, 0x09, 0xdd, 0x34, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/lds.proto",
}

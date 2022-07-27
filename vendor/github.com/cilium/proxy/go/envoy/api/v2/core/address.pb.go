// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/address.proto

package core

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	// [#not-implemented-hide:]
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}

func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{1, 0}
}

type Pipe struct {
	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{0}
}

func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (m *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(m, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.api.v2.core.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
	// to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
	// It is possible to distinguish a Listener address via the prefix/suffix matching
	// in :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>`.] When used
	// within an upstream :ref:`BindConfig <envoy_api_msg_core.BindConfig>`, the address
	// controls the source address of outbound connections. For :ref:`clusters
	// <envoy_api_msg_Cluster>`, the cluster type determines whether the
	// address must be an IP (*STATIC* or *EDS* clusters) or a hostname resolved by DNS
	// (*STRICT_DNS* or *LOGICAL_DNS* clusters). Address resolution can be customized
	// via :ref:`resolver_name <envoy_api_field_core.SocketAddress.resolver_name>`.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the resolver. This must have been registered with Envoy. If this is
	// empty, a context dependent default applies. If address is a hostname this
	// should be set for resolution other than DNS. If the address is a concrete
	// IP address, no resolution will occur.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibity
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat           bool     `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{1}
}

func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (m *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(m, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SocketAddress_OneofMarshaler, _SocketAddress_OneofUnmarshaler, _SocketAddress_OneofSizer, []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

func _SocketAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NamedPort)
	case nil:
	default:
		return fmt.Errorf("SocketAddress.PortSpecifier has unexpected type %T", x)
	}
	return nil
}

func _SocketAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SocketAddress)
	switch tag {
	case 3: // port_specifier.port_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PortSpecifier = &SocketAddress_PortValue{uint32(x)}
		return true, err
	case 4: // port_specifier.named_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortSpecifier = &SocketAddress_NamedPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _SocketAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NamedPort)))
		n += len(x.NamedPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TcpKeepalive struct {
	// Maximum number of keepalive probes to send without response before deciding
	// the connection is dead. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 9.)
	KeepaliveProbes *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes
	// start being sent. Default is to use the OS level configuration (unless
	// overridden, Linux defaults to 7200s (ie 2 hours.)
	KeepaliveTime *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. Default is to use the OS
	// level configuration (unless overridden, Linux defaults to 75s.)
	KeepaliveInterval    *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{2}
}

func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (m *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(m, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	// The address to bind to when creating a socket.
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// Whether to set the *IP_FREEBIND* option when creating the socket. When this
	// flag is set to true, allows the :ref:`source_address
	// <envoy_api_field_UpstreamBindConfig.source_address>` to be an IP address
	// that is not configured on the system running Envoy. When this flag is set
	// to false, the option *IP_FREEBIND* is disabled on the socket. When this
	// flag is not set (default), the socket is not modified, i.e. the option is
	// neither enabled nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions        []*SocketOption `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{3}
}

func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (m *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(m, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{4}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SocketAddress); err != nil {
			return err
		}
	case *Address_Pipe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipe); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Address.Address has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address.socket_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketAddress)
		err := b.DecodeMessage(msg)
		m.Address = &Address_SocketAddress{msg}
		return true, err
	case 2: // address.pipe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipe)
		err := b.DecodeMessage(msg)
		m.Address = &Address_Pipe{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		s := proto.Size(x.SocketAddress)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Address_Pipe:
		s := proto.Size(x.Pipe)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// CidrRange specifies an IP Address and a prefix length to construct
// the subnet mask for a `CIDR <https://tools.ietf.org/html/rfc4632>`_ range.
type CidrRange struct {
	// IPv4 or IPv6 address, e.g. ``192.0.0.0`` or ``2001:db8::``.
	AddressPrefix string `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	// Length of prefix, e.g. 0, 32.
	PrefixLen            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6906417f87bcce55, []int{5}
}

func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (m *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(m, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.api.v2.core.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.core.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.core.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.api.v2.core.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.core.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.core.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.core.CidrRange")
}

func init() { proto.RegisterFile("envoy/api/v2/core/address.proto", fileDescriptor_6906417f87bcce55) }

var fileDescriptor_6906417f87bcce55 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0x13, 0x3b,
	0x18, 0x8d, 0x33, 0xd3, 0x36, 0xf9, 0xd2, 0xe4, 0xa6, 0xd6, 0x95, 0x3a, 0x8a, 0x7a, 0x9b, 0x28,
	0xd5, 0x95, 0x42, 0x05, 0x33, 0x28, 0x45, 0xec, 0x99, 0x20, 0xda, 0xaa, 0x08, 0xc2, 0xd0, 0xc2,
	0x72, 0xe4, 0x24, 0x4e, 0xb0, 0x3a, 0x19, 0x5b, 0x9e, 0xe9, 0x00, 0x3b, 0xc4, 0x02, 0x21, 0xf6,
	0x3c, 0x01, 0x1b, 0x1e, 0x01, 0xb1, 0xea, 0x92, 0x57, 0x60, 0xc3, 0x02, 0x89, 0x45, 0x9f, 0xa2,
	0xc8, 0x9e, 0x9f, 0x0a, 0x02, 0x2a, 0xec, 0xec, 0xf3, 0x9d, 0x73, 0x7c, 0x3e, 0xfb, 0x33, 0xb4,
	0x69, 0x98, 0xf0, 0xe7, 0x0e, 0x11, 0xcc, 0x49, 0xfa, 0xce, 0x98, 0x4b, 0xea, 0x90, 0xc9, 0x44,
	0xd2, 0x28, 0xb2, 0x85, 0xe4, 0x31, 0xc7, 0x6b, 0x9a, 0x60, 0x13, 0xc1, 0xec, 0xa4, 0x6f, 0x2b,
	0x42, 0x6b, 0x63, 0x51, 0x33, 0x22, 0x11, 0x4d, 0x05, 0xad, 0xcd, 0x19, 0xe7, 0xb3, 0x80, 0x3a,
	0x7a, 0x37, 0x3a, 0x99, 0x3a, 0x4f, 0x25, 0x11, 0x82, 0xca, 0xcc, 0xb0, 0xb5, 0x9e, 0x90, 0x80,
	0x4d, 0x48, 0x4c, 0x9d, 0x7c, 0x91, 0x15, 0xfe, 0x9d, 0xf1, 0x19, 0xd7, 0x4b, 0x47, 0xad, 0x52,
	0xb4, 0xfb, 0x3f, 0x98, 0x43, 0x26, 0x28, 0xfe, 0x0f, 0x4c, 0x41, 0xe2, 0x27, 0x16, 0xea, 0xa0,
	0x5e, 0xd5, 0xad, 0x7e, 0x3c, 0x3b, 0x35, 0x4c, 0x59, 0xee, 0x20, 0x4f, 0xc3, 0xdd, 0xcf, 0x65,
	0xa8, 0x3f, 0xe4, 0xe3, 0x63, 0x1a, 0xdf, 0x4a, 0xe3, 0xe3, 0x07, 0x50, 0xd1, 0x0e, 0x63, 0x1e,
	0x68, 0x51, 0xa3, 0x7f, 0xc5, 0x5e, 0xe8, 0xc5, 0xfe, 0x41, 0x63, 0x0f, 0x33, 0x81, 0x0b, 0xca,
	0x7f, 0xe9, 0x25, 0x2a, 0x37, 0x91, 0x57, 0xd8, 0xe0, 0x2d, 0x58, 0xc9, 0x2e, 0xc7, 0x2a, 0xff,
	0x1c, 0x23, 0xaf, 0xe0, 0xab, 0x00, 0x82, 0xcb, 0xd8, 0x4f, 0x48, 0x70, 0x42, 0x2d, 0xa3, 0x83,
	0x7a, 0x75, 0xb7, 0xa6, 0x78, 0xcb, 0xdb, 0xa6, 0x75, 0x7e, 0x6e, 0xec, 0x95, 0xbc, 0xaa, 0x22,
	0x3c, 0x52, 0x75, 0xdc, 0x06, 0x08, 0xc9, 0x9c, 0x4e, 0x7c, 0x05, 0x59, 0xa6, 0x72, 0x55, 0x04,
	0x8d, 0x0d, 0xb9, 0x8c, 0xf1, 0x16, 0xd4, 0x25, 0x8d, 0x78, 0x90, 0x50, 0xe9, 0x2b, 0xd4, 0x5a,
	0x52, 0x1c, 0x6f, 0x35, 0x07, 0xef, 0x91, 0xb9, 0x72, 0xa9, 0x31, 0x91, 0xdc, 0xf0, 0xc7, 0x7c,
	0x2e, 0x48, 0x6c, 0x2d, 0x77, 0x50, 0xaf, 0xe2, 0x81, 0x82, 0x06, 0x1a, 0xe9, 0x76, 0xa1, 0x92,
	0xf7, 0x86, 0x57, 0xc0, 0x38, 0x1c, 0x0c, 0x9b, 0x25, 0xb5, 0x38, 0xba, 0x3d, 0x6c, 0xa2, 0x96,
	0xf9, 0xfa, 0xdd, 0x66, 0xc9, 0x5d, 0x87, 0x86, 0x0e, 0x1e, 0x09, 0x3a, 0x66, 0x53, 0x46, 0x25,
	0x5e, 0xfa, 0x70, 0x76, 0x6a, 0xa0, 0xee, 0x19, 0x82, 0xd5, 0xc3, 0xb1, 0x38, 0xa0, 0x54, 0x90,
	0x80, 0x25, 0x14, 0xef, 0x42, 0xf3, 0x38, 0xdf, 0xf8, 0x42, 0xf2, 0x11, 0x8d, 0xf4, 0x15, 0xd7,
	0xfa, 0x1b, 0x76, 0xfa, 0xfa, 0x76, 0xfe, 0xfa, 0xf6, 0xd1, 0x7e, 0x18, 0xef, 0xf4, 0x75, 0xb3,
	0xde, 0x3f, 0x85, 0x6a, 0xa8, 0x45, 0x78, 0x00, 0x8d, 0x0b, 0xa3, 0x98, 0xcd, 0xa9, 0xbe, 0xd7,
	0xcb, 0x6c, 0xea, 0x85, 0xe6, 0x90, 0xcd, 0x29, 0x3e, 0x00, 0x7c, 0x61, 0xc2, 0xc2, 0x98, 0xca,
	0x84, 0x04, 0xfa, 0xe2, 0x2f, 0x33, 0x5a, 0x2b, 0x74, 0xfb, 0x99, 0xac, 0xfb, 0x0d, 0x01, 0xb8,
	0x2c, 0x9c, 0x0c, 0x78, 0x38, 0x65, 0x33, 0xfc, 0x18, 0x1a, 0x11, 0x3f, 0x91, 0x63, 0xea, 0xe7,
	0x0f, 0x9f, 0xf6, 0xd9, 0xb9, 0x6c, 0x94, 0xdc, 0xc6, 0xa7, 0x2f, 0xed, 0x92, 0x9e, 0xa2, 0x37,
	0x7a, 0x8a, 0xea, 0xa9, 0x4f, 0x3e, 0x9d, 0x37, 0xa1, 0x32, 0x95, 0x94, 0x8e, 0x58, 0x38, 0xc9,
	0x7a, 0x6e, 0x2d, 0x44, 0x75, 0x39, 0x0f, 0xd2, 0xa0, 0x05, 0x17, 0xdf, 0x51, 0x81, 0xd4, 0x39,
	0x3e, 0x17, 0x31, 0xe3, 0x61, 0x64, 0x19, 0x1d, 0xa3, 0x57, 0xeb, 0xb7, 0x7f, 0x1b, 0xe8, 0xbe,
	0xe6, 0xa9, 0xf3, 0x2f, 0x76, 0x51, 0xf7, 0x2d, 0x82, 0x95, 0x3c, 0xcb, 0x7e, 0xe1, 0xf9, 0x97,
	0x4d, 0xee, 0x95, 0x72, 0xdb, 0xdc, 0xea, 0x1a, 0x98, 0x82, 0x89, 0xfc, 0x19, 0xd7, 0x7f, 0x61,
	0xa0, 0x3e, 0xf3, 0x5e, 0xc9, 0xd3, 0x34, 0xb7, 0x59, 0x7c, 0xa8, 0x7c, 0xd6, 0x5e, 0x21, 0xa8,
	0x0e, 0xd8, 0x44, 0x7a, 0x24, 0x9c, 0x51, 0x7c, 0x1d, 0x1a, 0x59, 0xdd, 0x17, 0x92, 0x4e, 0xd9,
	0xb3, 0xc5, 0xef, 0x5f, 0xcf, 0x08, 0x43, 0x5d, 0xc7, 0xbb, 0x00, 0x29, 0xd3, 0x0f, 0x68, 0xf8,
	0x27, 0xd3, 0x94, 0x7d, 0xf5, 0x6d, 0xc3, 0x7a, 0x81, 0xbc, 0x6a, 0xaa, 0xbd, 0x4b, 0x43, 0xd7,
	0x7c, 0xff, 0x75, 0x13, 0x8d, 0x96, 0xb5, 0x64, 0xe7, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x8d, 0xca, 0xde, 0x27, 0x05, 0x00, 0x00,
}

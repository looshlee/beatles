// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/fault/v2/fault.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Fixed delay (step function).
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}
var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fault_2b52d9bee2de8d9c, []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
type FaultDelay struct {
	// Delay type to use (fixed|exponential|..). Currently, only fixed delay (step function) is
	// supported.
	Type FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.config.filter.fault.v2.FaultDelay_FaultDelayType" json:"type,omitempty"`
	// An integer between 0-100 indicating the percentage of operations/connection requests
	// on which the delay will be injected.
	Percent uint32 `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"`
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	FaultDelaySecifier   isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_fault_2b52d9bee2de8d9c, []int{0}
}
func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay.Unmarshal(m, b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
}
func (dst *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(dst, src)
}
func (m *FaultDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay.Size(m)
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

func (m *FaultDelay) GetType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.Type
	}
	return FaultDelay_FIXED
}

func (m *FaultDelay) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	FixedDelay *duration.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetFixedDelay() *duration.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultDelay_OneofMarshaler, _FaultDelay_OneofUnmarshaler, _FaultDelay_OneofSizer, []interface{}{
		(*FaultDelay_FixedDelay)(nil),
	}
}

func _FaultDelay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedDelay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaultDelay.FaultDelaySecifier has unexpected type %T", x)
	}
	return nil
}

func _FaultDelay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultDelay)
	switch tag {
	case 3: // fault_delay_secifier.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(duration.Duration)
		err := b.DecodeMessage(msg)
		m.FaultDelaySecifier = &FaultDelay_FixedDelay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FaultDelay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		s := proto.Size(x.FixedDelay)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FaultDelay)(nil), "envoy.config.filter.fault.v2.FaultDelay")
	proto.RegisterEnum("envoy.config.filter.fault.v2.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/fault/v2/fault.proto", fileDescriptor_fault_2b52d9bee2de8d9c)
}

var fileDescriptor_fault_2b52d9bee2de8d9c = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x4d, 0x4b, 0xfb, 0x30,
	0x1c, 0xc7, 0x97, 0x6c, 0xfb, 0xff, 0x59, 0x86, 0x63, 0x84, 0x81, 0x75, 0x3e, 0x8d, 0x79, 0x29,
	0x3b, 0x24, 0x50, 0x0f, 0xde, 0xcb, 0x1c, 0x0a, 0x9e, 0x86, 0xa0, 0x78, 0x19, 0x5d, 0xfb, 0x4b,
	0x09, 0x94, 0xa6, 0xd4, 0xb4, 0xd8, 0xab, 0xaf, 0xc1, 0x83, 0xaf, 0xc1, 0x57, 0x20, 0x9e, 0xf6,
	0x4e, 0x3c, 0xef, 0x5d, 0x48, 0x93, 0x15, 0xf5, 0xe2, 0xed, 0x4b, 0xf2, 0xf9, 0x7d, 0x1f, 0x88,
	0x0b, 0x69, 0xa9, 0x2a, 0x1e, 0xaa, 0x54, 0xc8, 0x98, 0x0b, 0x99, 0x68, 0xc8, 0xb9, 0x08, 0x8a,
	0x44, 0xf3, 0xd2, 0xb3, 0x82, 0x65, 0xb9, 0xd2, 0x8a, 0x1e, 0x19, 0x92, 0x59, 0x92, 0x59, 0x92,
	0x59, 0xa0, 0xf4, 0xc6, 0x27, 0xb1, 0x52, 0x71, 0x02, 0xdc, 0xb0, 0xeb, 0x42, 0xf0, 0xa8, 0xc8,
	0x03, 0x2d, 0x55, 0x6a, 0xaf, 0xc7, 0xfb, 0x65, 0x90, 0xc8, 0x28, 0xd0, 0xc0, 0x1b, 0xb1, 0xfb,
	0x18, 0xc5, 0x2a, 0x56, 0x46, 0xf2, 0x5a, 0xd9, 0xd7, 0xe9, 0x0b, 0x26, 0x64, 0x51, 0x7b, 0xcf,
	0x21, 0x09, 0x2a, 0x7a, 0x47, 0x3a, 0xba, 0xca, 0xc0, 0x41, 0x13, 0xe4, 0x0e, 0xbc, 0x0b, 0xf6,
	0x57, 0x15, 0xf6, 0x7d, 0xf7, 0x43, 0xde, 0x56, 0x19, 0xf8, 0xe4, 0x63, 0xbb, 0x69, 0x77, 0x9f,
	0x11, 0x1e, 0xa2, 0xa5, 0x31, 0xa4, 0x67, 0xe4, 0x7f, 0x06, 0x79, 0x08, 0xa9, 0x76, 0xf0, 0x04,
	0xb9, 0x7b, 0x7e, 0xaf, 0x46, 0x3a, 0x33, 0xec, 0x44, 0xcb, 0xe6, 0x87, 0xde, 0x90, 0xbe, 0x90,
	0x4f, 0x10, 0xad, 0xa2, 0xda, 0xc9, 0x69, 0x4f, 0x90, 0xdb, 0xf7, 0x0e, 0x98, 0x5d, 0xcc, 0x9a,
	0xc5, 0x6c, 0xbe, 0x5b, 0xec, 0x0f, 0x5e, 0x3f, 0x4f, 0x91, 0x89, 0x7a, 0x43, 0x78, 0xd6, 0xba,
	0x6a, 0x2d, 0x89, 0xb9, 0x37, 0x45, 0xa6, 0x87, 0x64, 0xf0, 0xbb, 0x16, 0xed, 0x91, 0xee, 0xe2,
	0xfa, 0xfe, 0x72, 0x3e, 0x6c, 0xf9, 0xc7, 0x64, 0x64, 0x76, 0xd8, 0xa8, 0xd5, 0x23, 0x84, 0x52,
	0x48, 0xc8, 0x69, 0xf7, 0x7d, 0xbb, 0x69, 0x23, 0xbf, 0xf3, 0x80, 0x4b, 0x6f, 0xfd, 0xcf, 0x44,
	0x9e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x98, 0xe8, 0xe0, 0xbc, 0x01, 0x00, 0x00,
}

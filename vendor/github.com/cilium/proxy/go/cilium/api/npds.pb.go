// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/api/npds.proto

package cilium

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/api/v2"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	proto "github.com/golang/protobuf/proto"
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

// A network policy that is enforced by a filter on the network flows to/from
// associated hosts.
type NetworkPolicy struct {
	// The unique identifier of the network policy.
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The policy identifier associated with the network policy. Corresponds to
	// remote_policies entries in PortNetworkPolicyRule.
	// Required.
	Policy uint64 `protobuf:"varint,2,opt,name=policy,proto3" json:"policy,omitempty"`
	// The part of the policy to be enforced at ingress by the filter, as a set
	// of per-port network policies, one per destination L4 port.
	// Every PortNetworkPolicy element in this set has a unique port / protocol
	// combination.
	// Optional. If empty, all flows in this direction are denied.
	IngressPerPortPolicies []*PortNetworkPolicy `protobuf:"bytes,3,rep,name=ingress_per_port_policies,json=ingressPerPortPolicies,proto3" json:"ingress_per_port_policies,omitempty"`
	// The part of the policy to be enforced at egress by the filter, as a set
	// of per-port network policies, one per destination L4 port.
	// Every PortNetworkPolicy element in this set has a unique port / protocol
	// combination.
	// Optional. If empty, all flows in this direction are denied.
	EgressPerPortPolicies []*PortNetworkPolicy `protobuf:"bytes,4,rep,name=egress_per_port_policies,json=egressPerPortPolicies,proto3" json:"egress_per_port_policies,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *NetworkPolicy) Reset()         { *m = NetworkPolicy{} }
func (m *NetworkPolicy) String() string { return proto.CompactTextString(m) }
func (*NetworkPolicy) ProtoMessage()    {}
func (*NetworkPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{0}
}

func (m *NetworkPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPolicy.Unmarshal(m, b)
}
func (m *NetworkPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPolicy.Marshal(b, m, deterministic)
}
func (m *NetworkPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPolicy.Merge(m, src)
}
func (m *NetworkPolicy) XXX_Size() int {
	return xxx_messageInfo_NetworkPolicy.Size(m)
}
func (m *NetworkPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPolicy proto.InternalMessageInfo

func (m *NetworkPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkPolicy) GetPolicy() uint64 {
	if m != nil {
		return m.Policy
	}
	return 0
}

func (m *NetworkPolicy) GetIngressPerPortPolicies() []*PortNetworkPolicy {
	if m != nil {
		return m.IngressPerPortPolicies
	}
	return nil
}

func (m *NetworkPolicy) GetEgressPerPortPolicies() []*PortNetworkPolicy {
	if m != nil {
		return m.EgressPerPortPolicies
	}
	return nil
}

// A network policy to whitelist flows to a specific destination L4 port,
// as a conjunction of predicates on L3/L4/L7 flows.
// If all the predicates of a policy match a flow, the flow is whitelisted.
type PortNetworkPolicy struct {
	// The flows' destination L4 port number, as an unsigned 16-bit integer.
	// If 0, all destination L4 port numbers are matched by this predicate.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// The flows' L4 transport protocol.
	// Required.
	Protocol core.SocketAddress_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.api.v2.core.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The network policy rules to be enforced on the flows to the port.
	// Optional. A flow is matched by this predicate if either the set of
	// rules is empty or any of the rules matches it.
	Rules                []*PortNetworkPolicyRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PortNetworkPolicy) Reset()         { *m = PortNetworkPolicy{} }
func (m *PortNetworkPolicy) String() string { return proto.CompactTextString(m) }
func (*PortNetworkPolicy) ProtoMessage()    {}
func (*PortNetworkPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{1}
}

func (m *PortNetworkPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortNetworkPolicy.Unmarshal(m, b)
}
func (m *PortNetworkPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortNetworkPolicy.Marshal(b, m, deterministic)
}
func (m *PortNetworkPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortNetworkPolicy.Merge(m, src)
}
func (m *PortNetworkPolicy) XXX_Size() int {
	return xxx_messageInfo_PortNetworkPolicy.Size(m)
}
func (m *PortNetworkPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_PortNetworkPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_PortNetworkPolicy proto.InternalMessageInfo

func (m *PortNetworkPolicy) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PortNetworkPolicy) GetProtocol() core.SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return core.SocketAddress_TCP
}

func (m *PortNetworkPolicy) GetRules() []*PortNetworkPolicyRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A network policy rule, as a conjunction of predicates on L3/L7 flows.
// If all the predicates of a rule match a flow, the flow is matched by the
// rule.
type PortNetworkPolicyRule struct {
	// The set of identifiers of policies of remote hosts.
	// A flow is matched by this predicate if the identifier of the policy
	// applied on the flow's remote host is contained in this set.
	// Optional. If not specified, any remote host is matched by this predicate.
	RemotePolicies []uint64 `protobuf:"varint,1,rep,packed,name=remote_policies,json=remotePolicies,proto3" json:"remote_policies,omitempty"`
	// Optional L7 protocol parser name. This is only used if the parser is not
	// one of the well knows ones. If specified, the l7 parser having this name
	// needs to be built in to libcilium.so.
	L7Proto string `protobuf:"bytes,2,opt,name=l7_proto,json=l7Proto,proto3" json:"l7_proto,omitempty"`
	// Optional. If not specified, any L7 request is matched by this predicate.
	// All rules on any given port must have the same type of L7 rules!
	//
	// Types that are valid to be assigned to L7:
	//	*PortNetworkPolicyRule_HttpRules
	//	*PortNetworkPolicyRule_KafkaRules
	//	*PortNetworkPolicyRule_L7Rules
	L7                   isPortNetworkPolicyRule_L7 `protobuf_oneof:"l7"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PortNetworkPolicyRule) Reset()         { *m = PortNetworkPolicyRule{} }
func (m *PortNetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*PortNetworkPolicyRule) ProtoMessage()    {}
func (*PortNetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{2}
}

func (m *PortNetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortNetworkPolicyRule.Unmarshal(m, b)
}
func (m *PortNetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortNetworkPolicyRule.Marshal(b, m, deterministic)
}
func (m *PortNetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortNetworkPolicyRule.Merge(m, src)
}
func (m *PortNetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_PortNetworkPolicyRule.Size(m)
}
func (m *PortNetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_PortNetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_PortNetworkPolicyRule proto.InternalMessageInfo

func (m *PortNetworkPolicyRule) GetRemotePolicies() []uint64 {
	if m != nil {
		return m.RemotePolicies
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetL7Proto() string {
	if m != nil {
		return m.L7Proto
	}
	return ""
}

type isPortNetworkPolicyRule_L7 interface {
	isPortNetworkPolicyRule_L7()
}

type PortNetworkPolicyRule_HttpRules struct {
	HttpRules *HttpNetworkPolicyRules `protobuf:"bytes,100,opt,name=http_rules,json=httpRules,proto3,oneof"`
}

type PortNetworkPolicyRule_KafkaRules struct {
	KafkaRules *KafkaNetworkPolicyRules `protobuf:"bytes,101,opt,name=kafka_rules,json=kafkaRules,proto3,oneof"`
}

type PortNetworkPolicyRule_L7Rules struct {
	L7Rules *L7NetworkPolicyRules `protobuf:"bytes,102,opt,name=l7_rules,json=l7Rules,proto3,oneof"`
}

func (*PortNetworkPolicyRule_HttpRules) isPortNetworkPolicyRule_L7() {}

func (*PortNetworkPolicyRule_KafkaRules) isPortNetworkPolicyRule_L7() {}

func (*PortNetworkPolicyRule_L7Rules) isPortNetworkPolicyRule_L7() {}

func (m *PortNetworkPolicyRule) GetL7() isPortNetworkPolicyRule_L7 {
	if m != nil {
		return m.L7
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetHttpRules() *HttpNetworkPolicyRules {
	if x, ok := m.GetL7().(*PortNetworkPolicyRule_HttpRules); ok {
		return x.HttpRules
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetKafkaRules() *KafkaNetworkPolicyRules {
	if x, ok := m.GetL7().(*PortNetworkPolicyRule_KafkaRules); ok {
		return x.KafkaRules
	}
	return nil
}

func (m *PortNetworkPolicyRule) GetL7Rules() *L7NetworkPolicyRules {
	if x, ok := m.GetL7().(*PortNetworkPolicyRule_L7Rules); ok {
		return x.L7Rules
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PortNetworkPolicyRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PortNetworkPolicyRule_OneofMarshaler, _PortNetworkPolicyRule_OneofUnmarshaler, _PortNetworkPolicyRule_OneofSizer, []interface{}{
		(*PortNetworkPolicyRule_HttpRules)(nil),
		(*PortNetworkPolicyRule_KafkaRules)(nil),
		(*PortNetworkPolicyRule_L7Rules)(nil),
	}
}

func _PortNetworkPolicyRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PortNetworkPolicyRule)
	// l7
	switch x := m.L7.(type) {
	case *PortNetworkPolicyRule_HttpRules:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpRules); err != nil {
			return err
		}
	case *PortNetworkPolicyRule_KafkaRules:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KafkaRules); err != nil {
			return err
		}
	case *PortNetworkPolicyRule_L7Rules:
		b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.L7Rules); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PortNetworkPolicyRule.L7 has unexpected type %T", x)
	}
	return nil
}

func _PortNetworkPolicyRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PortNetworkPolicyRule)
	switch tag {
	case 100: // l7.http_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpNetworkPolicyRules)
		err := b.DecodeMessage(msg)
		m.L7 = &PortNetworkPolicyRule_HttpRules{msg}
		return true, err
	case 101: // l7.kafka_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KafkaNetworkPolicyRules)
		err := b.DecodeMessage(msg)
		m.L7 = &PortNetworkPolicyRule_KafkaRules{msg}
		return true, err
	case 102: // l7.l7_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(L7NetworkPolicyRules)
		err := b.DecodeMessage(msg)
		m.L7 = &PortNetworkPolicyRule_L7Rules{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PortNetworkPolicyRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PortNetworkPolicyRule)
	// l7
	switch x := m.L7.(type) {
	case *PortNetworkPolicyRule_HttpRules:
		s := proto.Size(x.HttpRules)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PortNetworkPolicyRule_KafkaRules:
		s := proto.Size(x.KafkaRules)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PortNetworkPolicyRule_L7Rules:
		s := proto.Size(x.L7Rules)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A set of network policy rules that match HTTP requests.
type HttpNetworkPolicyRules struct {
	// The set of HTTP network policy rules.
	// An HTTP request is matched if any of its rules matches the request.
	// Required and may not be empty.
	HttpRules            []*HttpNetworkPolicyRule `protobuf:"bytes,1,rep,name=http_rules,json=httpRules,proto3" json:"http_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HttpNetworkPolicyRules) Reset()         { *m = HttpNetworkPolicyRules{} }
func (m *HttpNetworkPolicyRules) String() string { return proto.CompactTextString(m) }
func (*HttpNetworkPolicyRules) ProtoMessage()    {}
func (*HttpNetworkPolicyRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{3}
}

func (m *HttpNetworkPolicyRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNetworkPolicyRules.Unmarshal(m, b)
}
func (m *HttpNetworkPolicyRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNetworkPolicyRules.Marshal(b, m, deterministic)
}
func (m *HttpNetworkPolicyRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNetworkPolicyRules.Merge(m, src)
}
func (m *HttpNetworkPolicyRules) XXX_Size() int {
	return xxx_messageInfo_HttpNetworkPolicyRules.Size(m)
}
func (m *HttpNetworkPolicyRules) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNetworkPolicyRules.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNetworkPolicyRules proto.InternalMessageInfo

func (m *HttpNetworkPolicyRules) GetHttpRules() []*HttpNetworkPolicyRule {
	if m != nil {
		return m.HttpRules
	}
	return nil
}

// An HTTP network policy rule, as a conjunction of predicates on HTTP requests.
// If all the predicates of a rule match an HTTP request, the request is allowed. Otherwise, it is
// denied.
type HttpNetworkPolicyRule struct {
	// A set of matchers on the HTTP request's headers' names and values.
	// If all the matchers in this set match an HTTP request, the request is allowed by this rule.
	// Otherwise, it is denied.
	//
	// Some special header names are:
	//
	// * *:uri*: The HTTP request's URI.
	// * *:method*: The HTTP request's method.
	// * *:authority*: Also maps to the HTTP 1.1 *Host* header.
	//
	// Optional. If empty, matches any HTTP request.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HttpNetworkPolicyRule) Reset()         { *m = HttpNetworkPolicyRule{} }
func (m *HttpNetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*HttpNetworkPolicyRule) ProtoMessage()    {}
func (*HttpNetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{4}
}

func (m *HttpNetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpNetworkPolicyRule.Unmarshal(m, b)
}
func (m *HttpNetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpNetworkPolicyRule.Marshal(b, m, deterministic)
}
func (m *HttpNetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpNetworkPolicyRule.Merge(m, src)
}
func (m *HttpNetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_HttpNetworkPolicyRule.Size(m)
}
func (m *HttpNetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpNetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpNetworkPolicyRule proto.InternalMessageInfo

func (m *HttpNetworkPolicyRule) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// A set of network policy rules that match Kafka requests.
type KafkaNetworkPolicyRules struct {
	// The set of Kafka network policy rules.
	// A Kafka request is matched if any of its rules matches the request.
	// Required and may not be empty.
	KafkaRules           []*KafkaNetworkPolicyRule `protobuf:"bytes,1,rep,name=kafka_rules,json=kafkaRules,proto3" json:"kafka_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KafkaNetworkPolicyRules) Reset()         { *m = KafkaNetworkPolicyRules{} }
func (m *KafkaNetworkPolicyRules) String() string { return proto.CompactTextString(m) }
func (*KafkaNetworkPolicyRules) ProtoMessage()    {}
func (*KafkaNetworkPolicyRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{5}
}

func (m *KafkaNetworkPolicyRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaNetworkPolicyRules.Unmarshal(m, b)
}
func (m *KafkaNetworkPolicyRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaNetworkPolicyRules.Marshal(b, m, deterministic)
}
func (m *KafkaNetworkPolicyRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaNetworkPolicyRules.Merge(m, src)
}
func (m *KafkaNetworkPolicyRules) XXX_Size() int {
	return xxx_messageInfo_KafkaNetworkPolicyRules.Size(m)
}
func (m *KafkaNetworkPolicyRules) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaNetworkPolicyRules.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaNetworkPolicyRules proto.InternalMessageInfo

func (m *KafkaNetworkPolicyRules) GetKafkaRules() []*KafkaNetworkPolicyRule {
	if m != nil {
		return m.KafkaRules
	}
	return nil
}

// A Kafka network policy rule, as a conjunction of predicates on Kafka requests.
// If all the predicates of a rule match a Kafka request, the request is allowed. Otherwise, it is
// denied.
type KafkaNetworkPolicyRule struct {
	// The Kafka request's API key.
	// If <0, all Kafka requests are matched by this predicate.
	ApiKey int32 `protobuf:"varint,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// The Kafka request's API version.
	// If <0, all Kafka requests are matched by this predicate.
	ApiVersion int32 `protobuf:"varint,2,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The Kafka request's topic.
	// Optional. If not specified, all Kafka requests are matched by this predicate.
	// If specified, this predicates only matches requests that contain this topic, and never
	// matches requests that don't contain any topic.
	Topic string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	// The Kafka request's client ID.
	// Optional. If not specified, all Kafka requests are matched by this predicate.
	// If specified, this predicates only matches requests that contain this client ID, and never
	// matches requests that don't contain any client ID.
	ClientId             string   `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaNetworkPolicyRule) Reset()         { *m = KafkaNetworkPolicyRule{} }
func (m *KafkaNetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*KafkaNetworkPolicyRule) ProtoMessage()    {}
func (*KafkaNetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{6}
}

func (m *KafkaNetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaNetworkPolicyRule.Unmarshal(m, b)
}
func (m *KafkaNetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaNetworkPolicyRule.Marshal(b, m, deterministic)
}
func (m *KafkaNetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaNetworkPolicyRule.Merge(m, src)
}
func (m *KafkaNetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_KafkaNetworkPolicyRule.Size(m)
}
func (m *KafkaNetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaNetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaNetworkPolicyRule proto.InternalMessageInfo

func (m *KafkaNetworkPolicyRule) GetApiKey() int32 {
	if m != nil {
		return m.ApiKey
	}
	return 0
}

func (m *KafkaNetworkPolicyRule) GetApiVersion() int32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *KafkaNetworkPolicyRule) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *KafkaNetworkPolicyRule) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

// A set of network policy rules that match generic L7 requests.
type L7NetworkPolicyRules struct {
	// The set of generic key-value pair policy rules.
	// A request is matched if any of these rules matches the request.
	// Required and may not be empty.
	L7Rules              []*L7NetworkPolicyRule `protobuf:"bytes,1,rep,name=l7_rules,json=l7Rules,proto3" json:"l7_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *L7NetworkPolicyRules) Reset()         { *m = L7NetworkPolicyRules{} }
func (m *L7NetworkPolicyRules) String() string { return proto.CompactTextString(m) }
func (*L7NetworkPolicyRules) ProtoMessage()    {}
func (*L7NetworkPolicyRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{7}
}

func (m *L7NetworkPolicyRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L7NetworkPolicyRules.Unmarshal(m, b)
}
func (m *L7NetworkPolicyRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L7NetworkPolicyRules.Marshal(b, m, deterministic)
}
func (m *L7NetworkPolicyRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L7NetworkPolicyRules.Merge(m, src)
}
func (m *L7NetworkPolicyRules) XXX_Size() int {
	return xxx_messageInfo_L7NetworkPolicyRules.Size(m)
}
func (m *L7NetworkPolicyRules) XXX_DiscardUnknown() {
	xxx_messageInfo_L7NetworkPolicyRules.DiscardUnknown(m)
}

var xxx_messageInfo_L7NetworkPolicyRules proto.InternalMessageInfo

func (m *L7NetworkPolicyRules) GetL7Rules() []*L7NetworkPolicyRule {
	if m != nil {
		return m.L7Rules
	}
	return nil
}

// A generic L7 policy rule, as a conjunction of predicates on l7 requests.
// If all the predicates of a rule match a request, the request is allowed. Otherwise, it is
// denied.
type L7NetworkPolicyRule struct {
	// Optional. If empty, matches any request.
	Rule                 map[string]string `protobuf:"bytes,1,rep,name=rule,proto3" json:"rule,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L7NetworkPolicyRule) Reset()         { *m = L7NetworkPolicyRule{} }
func (m *L7NetworkPolicyRule) String() string { return proto.CompactTextString(m) }
func (*L7NetworkPolicyRule) ProtoMessage()    {}
func (*L7NetworkPolicyRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04d25916f7381d1, []int{8}
}

func (m *L7NetworkPolicyRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L7NetworkPolicyRule.Unmarshal(m, b)
}
func (m *L7NetworkPolicyRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L7NetworkPolicyRule.Marshal(b, m, deterministic)
}
func (m *L7NetworkPolicyRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L7NetworkPolicyRule.Merge(m, src)
}
func (m *L7NetworkPolicyRule) XXX_Size() int {
	return xxx_messageInfo_L7NetworkPolicyRule.Size(m)
}
func (m *L7NetworkPolicyRule) XXX_DiscardUnknown() {
	xxx_messageInfo_L7NetworkPolicyRule.DiscardUnknown(m)
}

var xxx_messageInfo_L7NetworkPolicyRule proto.InternalMessageInfo

func (m *L7NetworkPolicyRule) GetRule() map[string]string {
	if m != nil {
		return m.Rule
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkPolicy)(nil), "cilium.NetworkPolicy")
	proto.RegisterType((*PortNetworkPolicy)(nil), "cilium.PortNetworkPolicy")
	proto.RegisterType((*PortNetworkPolicyRule)(nil), "cilium.PortNetworkPolicyRule")
	proto.RegisterType((*HttpNetworkPolicyRules)(nil), "cilium.HttpNetworkPolicyRules")
	proto.RegisterType((*HttpNetworkPolicyRule)(nil), "cilium.HttpNetworkPolicyRule")
	proto.RegisterType((*KafkaNetworkPolicyRules)(nil), "cilium.KafkaNetworkPolicyRules")
	proto.RegisterType((*KafkaNetworkPolicyRule)(nil), "cilium.KafkaNetworkPolicyRule")
	proto.RegisterType((*L7NetworkPolicyRules)(nil), "cilium.L7NetworkPolicyRules")
	proto.RegisterType((*L7NetworkPolicyRule)(nil), "cilium.L7NetworkPolicyRule")
	proto.RegisterMapType((map[string]string)(nil), "cilium.L7NetworkPolicyRule.RuleEntry")
}

func init() { proto.RegisterFile("cilium/api/npds.proto", fileDescriptor_c04d25916f7381d1) }

var fileDescriptor_c04d25916f7381d1 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xdf, 0x49, 0x9c, 0xb4, 0x79, 0xd1, 0x2e, 0xec, 0x6c, 0x93, 0xba, 0x61, 0x9b, 0x06, 0x03,
	0x52, 0xb6, 0x52, 0x9d, 0x55, 0x7a, 0x08, 0x2d, 0x07, 0xb4, 0x11, 0x8b, 0x8a, 0x0a, 0x28, 0x72,
	0x57, 0x1c, 0x16, 0xb1, 0xd1, 0xac, 0xfd, 0xda, 0x8e, 0xe2, 0x7a, 0xcc, 0x78, 0x12, 0x14, 0x8e,
	0x2b, 0x2e, 0x5c, 0xe1, 0x73, 0x20, 0x71, 0xe6, 0xb4, 0xdf, 0x81, 0xaf, 0x00, 0x07, 0x3e, 0x45,
	0x91, 0x67, 0xec, 0x6c, 0xad, 0x75, 0xca, 0x85, 0x8b, 0x35, 0xf6, 0xef, 0x8f, 0xdf, 0xfb, 0xcd,
	0x1b, 0x1b, 0x5a, 0x3e, 0x0f, 0xf9, 0xfc, 0x6a, 0xc0, 0x62, 0x3e, 0x88, 0xe2, 0x20, 0x71, 0x63,
	0x29, 0x94, 0xa0, 0x75, 0xf3, 0xb8, 0xb3, 0x87, 0xd1, 0x42, 0x2c, 0x35, 0xba, 0x18, 0x0e, 0x7c,
	0x21, 0x71, 0xc0, 0x82, 0x40, 0x62, 0x92, 0x11, 0x3b, 0x0f, 0x0b, 0x84, 0x80, 0x27, 0xbe, 0x58,
	0xa0, 0x5c, 0x66, 0x68, 0xb7, 0x80, 0x4a, 0x31, 0x57, 0x68, 0xae, 0xb9, 0xfa, 0x42, 0x88, 0x8b,
	0x10, 0x35, 0x81, 0x45, 0x91, 0x50, 0x4c, 0x71, 0x11, 0xe5, 0xde, 0xdb, 0x0b, 0x16, 0xf2, 0x80,
	0x29, 0x1c, 0xe4, 0x0b, 0x03, 0x38, 0x7f, 0x13, 0xb8, 0xfb, 0x35, 0xaa, 0x1f, 0x84, 0x9c, 0x4d,
	0x44, 0xc8, 0xfd, 0x25, 0xa5, 0x60, 0x45, 0xec, 0x0a, 0x6d, 0xd2, 0x23, 0xfd, 0x86, 0xa7, 0xd7,
	0xb4, 0x0d, 0xf5, 0x58, 0xa3, 0x76, 0xa5, 0x47, 0xfa, 0x96, 0x97, 0xdd, 0xd1, 0x67, 0xb0, 0xc3,
	0xa3, 0x8b, 0xb4, 0x87, 0x69, 0x8c, 0x72, 0x1a, 0x0b, 0xa9, 0xa6, 0x1a, 0xe2, 0x98, 0xd8, 0xd5,
	0x5e, 0xb5, 0xdf, 0x1c, 0xee, 0xb8, 0xa6, 0x7f, 0x77, 0x22, 0xa4, 0x2a, 0xbc, 0xc9, 0x6b, 0x67,
	0xda, 0x09, 0xca, 0x14, 0x9c, 0x64, 0x42, 0xea, 0x81, 0x8d, 0xeb, 0x4c, 0xad, 0xff, 0x32, 0x6d,
	0x61, 0x99, 0xa7, 0xf3, 0x3b, 0x81, 0xfb, 0x6f, 0x91, 0xe9, 0x1e, 0x58, 0xa9, 0xbd, 0xee, 0xf5,
	0xee, 0xb8, 0xf9, 0xc7, 0x3f, 0xaf, 0xab, 0xf5, 0x7d, 0xcb, 0xbe, 0xbe, 0xae, 0x7a, 0x1a, 0xa0,
	0x4f, 0x61, 0x53, 0xe7, 0xe4, 0x8b, 0x50, 0xb7, 0x7e, 0x6f, 0xf8, 0xc8, 0xd5, 0x1b, 0xe1, 0xb2,
	0x98, 0xbb, 0x8b, 0xa1, 0x9b, 0xee, 0xa3, 0x7b, 0x26, 0xfc, 0x19, 0xaa, 0x27, 0xd9, 0x6e, 0x4e,
	0x32, 0x81, 0xb7, 0x92, 0xd2, 0x43, 0xa8, 0xc9, 0x79, 0xb8, 0xca, 0x64, 0x77, 0x7d, 0xf9, 0xf3,
	0x10, 0x3d, 0xc3, 0x75, 0x7e, 0xab, 0x40, 0xab, 0x94, 0x40, 0x0f, 0xe1, 0x1d, 0x89, 0x57, 0x42,
	0xe1, 0x9b, 0x5c, 0x48, 0xaf, 0xda, 0xb7, 0xc6, 0x90, 0x76, 0x50, 0xfb, 0x85, 0x54, 0x6c, 0xe2,
	0xdd, 0x33, 0x94, 0x55, 0xaa, 0x3b, 0xb0, 0x19, 0x8e, 0xa6, 0xba, 0x24, 0xdd, 0x4a, 0xc3, 0xdb,
	0x08, 0x47, 0xba, 0x56, 0xfa, 0x29, 0xc0, 0xa5, 0x52, 0xf1, 0xd4, 0xd4, 0x18, 0xf4, 0x48, 0xbf,
	0x39, 0xec, 0xe6, 0x35, 0x9e, 0x28, 0x15, 0xbf, 0x55, 0x42, 0x72, 0x72, 0xc7, 0x6b, 0xa4, 0x1a,
	0x7d, 0x43, 0xc7, 0xd0, 0x9c, 0xb1, 0xf3, 0x19, 0xcb, 0x1c, 0x50, 0x3b, 0xec, 0xe5, 0x0e, 0xa7,
	0x29, 0x54, 0x6a, 0x01, 0x5a, 0x65, 0x3c, 0x8e, 0x74, 0x7d, 0xc6, 0xe0, 0x5c, 0x1b, 0x3c, 0xcc,
	0x0d, 0xbe, 0x1c, 0x95, 0xaa, 0x37, 0xc2, 0x91, 0x5e, 0x8e, 0x2d, 0xa8, 0x84, 0x23, 0xe7, 0x25,
	0xb4, 0xcb, 0x6b, 0xa5, 0x27, 0x85, 0xfe, 0x48, 0x71, 0x0f, 0x4a, 0x35, 0x6f, 0x92, 0xdc, 0x24,
	0x37, 0x1a, 0x75, 0x9e, 0x41, 0xab, 0x94, 0x4f, 0x3f, 0x81, 0x8d, 0x4b, 0x64, 0x01, 0xca, 0xdc,
	0xff, 0xfd, 0xe2, 0x9c, 0x98, 0xa3, 0x7a, 0xa2, 0x29, 0x5f, 0x31, 0xe5, 0x5f, 0xa2, 0xf4, 0x72,
	0x85, 0x73, 0x0e, 0xdb, 0x6b, 0x32, 0xa2, 0xa7, 0xc5, 0x64, 0x8d, 0x77, 0xf7, 0xf6, 0x64, 0x0b,
	0xc5, 0xdf, 0x88, 0xd8, 0x79, 0x4d, 0xa0, 0x5d, 0x2e, 0xa1, 0xdb, 0xb0, 0xc1, 0x62, 0x3e, 0x9d,
	0xe1, 0x52, 0x1f, 0x86, 0x9a, 0x57, 0x67, 0x31, 0x3f, 0xc5, 0xf4, 0x88, 0x34, 0x53, 0x60, 0x81,
	0x32, 0xe1, 0x22, 0xd2, 0x93, 0x53, 0xf3, 0x80, 0xc5, 0xfc, 0x1b, 0xf3, 0x24, 0x9d, 0x6d, 0x25,
	0x62, 0xee, 0xdb, 0xd5, 0x74, 0xa8, 0xc6, 0xbb, 0xe9, 0xbb, 0x6d, 0xd9, 0xb6, 0xaf, 0xc9, 0xf0,
	0xfe, 0x8b, 0x6f, 0xd9, 0xc1, 0x8f, 0x4f, 0x0e, 0x9e, 0x3f, 0x3e, 0x38, 0x72, 0xa7, 0x07, 0xdf,
	0xed, 0x7f, 0xe8, 0x19, 0x2e, 0x1d, 0x41, 0xc3, 0x0f, 0x39, 0x46, 0x6a, 0xca, 0x03, 0xdb, 0xd2,
	0xc2, 0x4e, 0x2a, 0x6c, 0xc9, 0x07, 0x65, 0xaa, 0x4d, 0x43, 0xfe, 0x22, 0x70, 0x9e, 0xc3, 0x56,
	0xd9, 0x34, 0xd0, 0xf1, 0x8d, 0xe9, 0x31, 0x21, 0xbd, 0x77, 0xcb, 0xf4, 0x14, 0x12, 0xca, 0xc7,
	0xc8, 0xf9, 0x99, 0xc0, 0x83, 0x12, 0x32, 0x3d, 0x02, 0x2b, 0x35, 0xce, 0x7c, 0x3f, 0xba, 0xc5,
	0xd7, 0x4d, 0x2f, 0x4f, 0x23, 0x25, 0x97, 0x9e, 0x96, 0x74, 0x46, 0xd0, 0x58, 0x3d, 0xa2, 0xef,
	0x42, 0x35, 0xcf, 0xb7, 0xe1, 0xa5, 0x4b, 0xba, 0x05, 0xb5, 0x05, 0x0b, 0xe7, 0x98, 0x1d, 0x48,
	0x73, 0x73, 0x5c, 0xf9, 0x98, 0x0c, 0x7f, 0xaa, 0xc0, 0x6e, 0xc1, 0xfe, 0xb3, 0xfc, 0x7f, 0x70,
	0x86, 0x72, 0xc1, 0x7d, 0xa4, 0x2f, 0xa0, 0x75, 0xa6, 0x24, 0xb2, 0xab, 0x9b, 0xb4, 0xf4, 0xa0,
	0x77, 0x8b, 0x93, 0xb7, 0x12, 0x7a, 0xf8, 0xfd, 0x1c, 0x13, 0xd5, 0xd9, 0x5b, 0x8b, 0x27, 0xb1,
	0x88, 0x12, 0x74, 0xee, 0xf4, 0xc9, 0x63, 0x42, 0x5f, 0x11, 0xd8, 0xfa, 0x1c, 0x95, 0x7f, 0xf9,
	0xbf, 0xfb, 0x3f, 0x7a, 0xf5, 0xe7, 0x5f, 0xbf, 0x56, 0x3e, 0x70, 0xba, 0x85, 0xff, 0xdc, 0x71,
	0x64, 0xde, 0xb3, 0xfa, 0xa6, 0x1d, 0x93, 0xfd, 0x97, 0x75, 0xfd, 0xbd, 0x3a, 0xfc, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x87, 0x8e, 0x34, 0x5c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkPolicyDiscoveryServiceClient is the client API for NetworkPolicyDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkPolicyDiscoveryServiceClient interface {
	StreamNetworkPolicies(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient, error)
	FetchNetworkPolicies(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type networkPolicyDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkPolicyDiscoveryServiceClient(cc *grpc.ClientConn) NetworkPolicyDiscoveryServiceClient {
	return &networkPolicyDiscoveryServiceClient{cc}
}

func (c *networkPolicyDiscoveryServiceClient) StreamNetworkPolicies(ctx context.Context, opts ...grpc.CallOption) (NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkPolicyDiscoveryService_serviceDesc.Streams[0], "/cilium.NetworkPolicyDiscoveryService/StreamNetworkPolicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkPolicyDiscoveryServiceStreamNetworkPoliciesClient{stream}
	return x, nil
}

type NetworkPolicyDiscoveryService_StreamNetworkPoliciesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type networkPolicyDiscoveryServiceStreamNetworkPoliciesClient struct {
	grpc.ClientStream
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkPolicyDiscoveryServiceClient) FetchNetworkPolicies(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/cilium.NetworkPolicyDiscoveryService/FetchNetworkPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkPolicyDiscoveryServiceServer is the server API for NetworkPolicyDiscoveryService service.
type NetworkPolicyDiscoveryServiceServer interface {
	StreamNetworkPolicies(NetworkPolicyDiscoveryService_StreamNetworkPoliciesServer) error
	FetchNetworkPolicies(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterNetworkPolicyDiscoveryServiceServer(s *grpc.Server, srv NetworkPolicyDiscoveryServiceServer) {
	s.RegisterService(&_NetworkPolicyDiscoveryService_serviceDesc, srv)
}

func _NetworkPolicyDiscoveryService_StreamNetworkPolicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkPolicyDiscoveryServiceServer).StreamNetworkPolicies(&networkPolicyDiscoveryServiceStreamNetworkPoliciesServer{stream})
}

type NetworkPolicyDiscoveryService_StreamNetworkPoliciesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type networkPolicyDiscoveryServiceStreamNetworkPoliciesServer struct {
	grpc.ServerStream
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkPolicyDiscoveryServiceStreamNetworkPoliciesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NetworkPolicyDiscoveryService_FetchNetworkPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkPolicyDiscoveryServiceServer).FetchNetworkPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cilium.NetworkPolicyDiscoveryService/FetchNetworkPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkPolicyDiscoveryServiceServer).FetchNetworkPolicies(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkPolicyDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cilium.NetworkPolicyDiscoveryService",
	HandlerType: (*NetworkPolicyDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchNetworkPolicies",
			Handler:    _NetworkPolicyDiscoveryService_FetchNetworkPolicies_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkPolicies",
			Handler:       _NetworkPolicyDiscoveryService_StreamNetworkPolicies_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cilium/api/npds.proto",
}

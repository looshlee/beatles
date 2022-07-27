// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package core

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

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// REST-JSON legacy corresponds to the v1 API.
	ApiConfigSource_REST_LEGACY ApiConfigSource_ApiType = 0
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "REST_LEGACY",
	1: "REST",
	2: "GRPC",
}
var ApiConfigSource_ApiType_value = map[string]int32{
	"REST_LEGACY": 0,
	"REST":        1,
	"GRPC":        2,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}
func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_source_86f287b481865861, []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
type ApiConfigSource struct {
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// Cluster names should be used only with REST_LEGACY/REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay         *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_86f287b481865861, []int{0}
}
func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (dst *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(dst, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_REST_LEGACY
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_86f287b481865861, []int{1}
}
func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedConfigSource.Unmarshal(m, b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
}
func (dst *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(dst, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return xxx_messageInfo_AggregatedConfigSource.Size(m)
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager_cluster>`, :ref:`routes
// <config_http_conn_man_route_table>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_86f287b481865861, []int{2}
}
func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (dst *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(dst, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigSource_OneofMarshaler, _ConfigSource_OneofUnmarshaler, _ConfigSource_OneofSizer, []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
	}
}

func _ConfigSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Path)
	case *ConfigSource_ApiConfigSource:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiConfigSource); err != nil {
			return err
		}
	case *ConfigSource_Ads:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ads); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigSource.ConfigSourceSpecifier has unexpected type %T", x)
	}
	return nil
}

func _ConfigSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigSource)
	switch tag {
	case 1: // config_source_specifier.path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigSourceSpecifier = &ConfigSource_Path{x}
		return true, err
	case 2: // config_source_specifier.api_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApiConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_ApiConfigSource{msg}
		return true, err
	case 3: // config_source_specifier.ads
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AggregatedConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_Ads{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Path)))
		n += len(x.Path)
	case *ConfigSource_ApiConfigSource:
		s := proto.Size(x.ApiConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigSource_Ads:
		s := proto.Size(x.Ads)
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
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.api.v2.core.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.api.v2.core.AggregatedConfigSource")
	proto.RegisterType((*ConfigSource)(nil), "envoy.api.v2.core.ConfigSource")
	proto.RegisterEnum("envoy.api.v2.core.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/config_source.proto", fileDescriptor_config_source_86f287b481865861)
}

var fileDescriptor_config_source_86f287b481865861 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x6e, 0xdb, 0x3c,
	0x14, 0xc5, 0x4d, 0x49, 0xf9, 0x92, 0xd0, 0xf6, 0x67, 0x87, 0x08, 0x1a, 0x35, 0x83, 0x2b, 0xa8,
	0x2d, 0xe0, 0x76, 0xa0, 0x00, 0x75, 0xee, 0xe0, 0x3f, 0x81, 0x3d, 0x14, 0x45, 0x4a, 0x67, 0xe9,
	0x24, 0x30, 0x12, 0xcd, 0x10, 0x70, 0x4d, 0x82, 0x94, 0x05, 0x78, 0xed, 0x53, 0xf4, 0x11, 0xfa,
	0x08, 0x45, 0xa7, 0x0c, 0x7d, 0x87, 0x8e, 0x05, 0xba, 0xe5, 0x2d, 0x0a, 0xd2, 0x32, 0x1a, 0x27,
	0x06, 0xba, 0xdd, 0x7b, 0xee, 0x3d, 0xd2, 0xb9, 0x3f, 0x09, 0xbe, 0x64, 0xcb, 0x4a, 0xae, 0x13,
	0xaa, 0x44, 0x52, 0xa5, 0x49, 0x2e, 0x35, 0x4b, 0x72, 0xb9, 0x9c, 0x0b, 0x9e, 0x19, 0xb9, 0xd2,
	0x39, 0xc3, 0x4a, 0xcb, 0x52, 0xa2, 0x13, 0xb7, 0x86, 0xa9, 0x12, 0xb8, 0x4a, 0xb1, 0x5d, 0x3b,
	0x7f, 0xf1, 0xd8, 0xc9, 0xb5, 0xca, 0x33, 0xc3, 0x74, 0x25, 0xb6, 0xc6, 0xf3, 0x1e, 0x97, 0x92,
	0x2f, 0x58, 0xe2, 0xba, 0xeb, 0xd5, 0x3c, 0x29, 0x56, 0x9a, 0x96, 0x42, 0x2e, 0xeb, 0xf9, 0x59,
	0x45, 0x17, 0xa2, 0xa0, 0x25, 0x4b, 0xb6, 0x45, 0x3d, 0x38, 0xe5, 0x92, 0x4b, 0x57, 0x26, 0xb6,
	0xda, 0xa8, 0xf1, 0x0f, 0x0f, 0x76, 0x06, 0x4a, 0x8c, 0x5c, 0xc4, 0x99, 0x4b, 0x88, 0x3e, 0xc0,
	0x23, 0xaa, 0x44, 0x56, 0xae, 0x15, 0x0b, 0x41, 0x04, 0xfa, 0xff, 0xa7, 0xaf, 0xf1, 0xa3, 0xb8,
	0xf8, 0x81, 0xcb, 0xf6, 0x57, 0x6b, 0xc5, 0x86, 0xf0, 0xfb, 0xdd, 0xad, 0x7f, 0xf0, 0x19, 0x78,
	0x5d, 0x40, 0x0e, 0xe9, 0x46, 0x44, 0xcf, 0x61, 0x3b, 0x5f, 0xac, 0x4c, 0xc9, 0x74, 0xb6, 0xa4,
	0x9f, 0x98, 0x09, 0xbd, 0xc8, 0xef, 0x1f, 0x93, 0x56, 0x2d, 0xbe, 0xb7, 0x1a, 0x1a, 0xc1, 0xf6,
	0xfd, 0x83, 0x4d, 0x18, 0x44, 0x7e, 0xbf, 0x99, 0xf6, 0xf6, 0xbc, 0x7c, 0xa2, 0x55, 0x3e, 0xdb,
	0xac, 0x91, 0x16, 0xff, 0xdb, 0x18, 0x34, 0x86, 0x6d, 0xcd, 0xe6, 0x9a, 0x99, 0x9b, 0xac, 0x60,
	0x0b, 0xba, 0x0e, 0xfd, 0x08, 0xf4, 0x9b, 0xe9, 0x53, 0xbc, 0xe1, 0x86, 0xb7, 0xdc, 0xf0, 0xb8,
	0xe6, 0x36, 0x0c, 0xbe, 0xfc, 0x7a, 0x06, 0x48, 0xab, 0x76, 0x8d, 0xad, 0x29, 0xc6, 0xf0, 0xb0,
	0xbe, 0x07, 0x75, 0x60, 0x93, 0x5c, 0xcc, 0xae, 0xb2, 0x77, 0x17, 0x93, 0xc1, 0xe8, 0x63, 0xb7,
	0x81, 0x8e, 0x60, 0x60, 0x85, 0x2e, 0xb0, 0xd5, 0x84, 0x5c, 0x8e, 0xba, 0x5e, 0x1c, 0xc2, 0x27,
	0x03, 0xce, 0x35, 0xe3, 0xb4, 0x64, 0xc5, 0x7d, 0x2c, 0xf1, 0x4f, 0x00, 0x5b, 0x3b, 0x74, 0x4f,
	0x61, 0xa0, 0x68, 0x79, 0xe3, 0xc8, 0x1e, 0x4f, 0x1b, 0xc4, 0x75, 0xe8, 0x12, 0x9e, 0x58, 0xe6,
	0x3b, 0xbf, 0x4a, 0xe8, 0xb9, 0xe8, 0xf1, 0xbf, 0xe1, 0x4f, 0x1b, 0xa4, 0x43, 0x1f, 0x7c, 0xc5,
	0xb7, 0xd0, 0xa7, 0x85, 0xa9, 0xcf, 0x7f, 0xb5, 0xef, 0x19, 0x7b, 0x03, 0x4f, 0x1b, 0xc4, 0xfa,
	0x86, 0x11, 0x3c, 0xdb, 0x09, 0x93, 0x19, 0xc5, 0x72, 0x31, 0x17, 0x4c, 0xa3, 0x83, 0x6f, 0x77,
	0xb7, 0x3e, 0x18, 0x06, 0x5f, 0x7f, 0xf7, 0xc0, 0xf5, 0x7f, 0x0e, 0xe8, 0x9b, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x58, 0x46, 0xbd, 0x57, 0xf8, 0x02, 0x00, 0x00,
}

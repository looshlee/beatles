// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The tracing configuration specifies global
// settings for the HTTP tracer used by Envoy. The configuration is defined by
// the :ref:`Bootstrap <envoy_api_msg_config.bootstrap.v2.Bootstrap>` :ref:`tracing
// <envoy_api_field_config.bootstrap.v2.Bootstrap.tracing>` field. Envoy may support other tracers
// in the future, but right now the HTTP tracer is the only one supported.
type Tracing struct {
	// Provides configuration for the HTTP tracer.
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	// The name of the HTTP trace driver to instantiate. The name must match a
	// supported HTTP trace driver. Built-in trace drivers:
	//
	// - *envoy.lightstep*
	// - *envoy.zipkin*
	// - *envoy.dynamic.ot*
	// - *envoy.tracers.datadog*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being instantiated.
	// See the trace drivers for examples:
	//
	// - :ref:`LightstepConfig <envoy_api_msg_config.trace.v2.LightstepConfig>`
	// - :ref:`ZipkinConfig <envoy_api_msg_config.trace.v2.ZipkinConfig>`
	// - :ref:`DynamicOtConfig <envoy_api_msg_config.trace.v2.DynamicOtConfig>`
	// - :ref:`DatadogConfig <envoy_api_msg_config.trace.v2.DatadogConfig>`
	//
	// Types that are valid to be assigned to ConfigType:
	//	*Tracing_Http_Config
	//	*Tracing_Http_TypedConfig
	ConfigType           isTracing_Http_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTracing_Http_ConfigType interface {
	isTracing_Http_ConfigType()
}

type Tracing_Http_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Tracing_Http_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Tracing_Http_Config) isTracing_Http_ConfigType() {}

func (*Tracing_Http_TypedConfig) isTracing_Http_ConfigType() {}

func (m *Tracing_Http) GetConfigType() isTracing_Http_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Tracing_Http) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Tracing_Http_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Tracing_Http) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Tracing_Http_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tracing_Http_Config)(nil),
		(*Tracing_Http_TypedConfig)(nil),
	}
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <https://lightstep.com/>`_ API.
	AccessTokenFile      string   `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{1}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	// The cluster manager cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit bool `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	// Determines whether client and server spans will shared the same span id.
	// The default value is true.
	SharedSpanContext    *wrappers.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *wrappers.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
type DynamicOtConfig struct {
	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{3}
}

func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (m *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(m, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration for the Datadog tracer.
type DatadogConfig struct {
	// The cluster to use for submitting traces to the Datadog agent.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The name used for the service when traces are generated by envoy.
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatadogConfig) Reset()         { *m = DatadogConfig{} }
func (m *DatadogConfig) String() string { return proto.CompactTextString(m) }
func (*DatadogConfig) ProtoMessage()    {}
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{4}
}

func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (m *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(m, src)
}
func (m *DatadogConfig) XXX_Size() int {
	return xxx_messageInfo_DatadogConfig.Size(m)
}
func (m *DatadogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatadogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatadogConfig proto.InternalMessageInfo

func (m *DatadogConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *DatadogConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

// Configuration structure.
type TraceServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{5}
}

func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (m *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(m, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v2.DatadogConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_0785d24fc8ab55c7) }

var fileDescriptor_0785d24fc8ab55c7 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x4e, 0x14, 0x4d,
	0x14, 0xa5, 0x87, 0xf9, 0xe0, 0xe3, 0xce, 0x90, 0x81, 0x52, 0x03, 0x4e, 0x94, 0x60, 0x63, 0x8c,
	0x0b, 0xd3, 0x1d, 0xc6, 0x28, 0xb8, 0x74, 0x40, 0x45, 0xe3, 0x0f, 0x69, 0x88, 0x0b, 0x36, 0x9d,
	0x9a, 0xea, 0xa2, 0xa9, 0xd0, 0x53, 0x55, 0xa9, 0xae, 0x69, 0xe9, 0x9d, 0x6b, 0x1f, 0xc3, 0xc7,
	0x70, 0xe5, 0xdb, 0x18, 0xf7, 0x3e, 0x80, 0xa9, 0x9f, 0x11, 0xa4, 0x59, 0x98, 0xb0, 0xeb, 0xba,
	0xe7, 0x9c, 0xba, 0xf7, 0x9e, 0xa9, 0x33, 0x70, 0x8f, 0xf2, 0x4a, 0xd4, 0x31, 0x11, 0xfc, 0x98,
	0xe5, 0xb1, 0x56, 0x98, 0xd0, 0xb8, 0x1a, 0xb8, 0x8f, 0x48, 0x2a, 0xa1, 0x05, 0xba, 0x65, 0x29,
	0x91, 0xa3, 0x44, 0x0e, 0xa9, 0x06, 0xfd, 0xfb, 0x4e, 0x89, 0x25, 0x33, 0x02, 0x22, 0x14, 0x8d,
	0x73, 0x25, 0x49, 0x5a, 0x52, 0x55, 0xb1, 0xa9, 0xb8, 0x7f, 0x3b, 0x17, 0x22, 0x2f, 0x68, 0x6c,
	0x4f, 0xa3, 0xc9, 0x71, 0x8c, 0x79, 0xed, 0xa1, 0x3b, 0x97, 0xa1, 0x52, 0xab, 0x09, 0xd1, 0x1e,
	0x5d, 0xbb, 0x8c, 0x7e, 0x52, 0x58, 0x4a, 0xaa, 0x4a, 0x8f, 0xaf, 0x54, 0xb8, 0x60, 0x19, 0xd6,
	0x34, 0x9e, 0x7e, 0x38, 0x20, 0xfc, 0x11, 0xc0, 0xfc, 0xa1, 0xc2, 0x84, 0xf1, 0x1c, 0x6d, 0x41,
	0xfb, 0x44, 0x6b, 0xb9, 0x1a, 0xac, 0x07, 0x0f, 0x3b, 0x83, 0x8d, 0xe8, 0xca, 0x4d, 0x22, 0xcf,
	0x8e, 0xf6, 0xb4, 0x96, 0x89, 0x15, 0xf4, 0xbf, 0x06, 0xd0, 0x36, 0x47, 0x74, 0x17, 0xda, 0x1c,
	0x8f, 0xa9, 0xbd, 0x61, 0x61, 0xb8, 0xf0, 0xed, 0xe7, 0xf7, 0xd9, 0xb6, 0x6a, 0xad, 0x07, 0x89,
	0x2d, 0xa3, 0x4d, 0x98, 0x73, 0xb7, 0xad, 0xb6, 0x6c, 0x8b, 0x95, 0xc8, 0x8d, 0x1d, 0x4d, 0xc7,
	0x8e, 0x0e, 0xec, 0x52, 0x7b, 0x33, 0x89, 0x27, 0xa2, 0x67, 0xd0, 0xd5, 0xb5, 0xa4, 0x59, 0xea,
	0x85, 0xb3, 0x56, 0x78, 0xb3, 0x21, 0x7c, 0xce, 0xeb, 0xbd, 0x99, 0xa4, 0x63, 0xb9, 0x3b, 0x96,
	0x3a, 0x5c, 0x84, 0x8e, 0x13, 0xa5, 0xa6, 0x1a, 0x7e, 0x0e, 0xa0, 0xf7, 0x96, 0xe5, 0x27, 0xba,
	0xd4, 0x54, 0x3a, 0x0a, 0x7a, 0x0a, 0xcb, 0x44, 0x14, 0x05, 0x25, 0x5a, 0xa8, 0x94, 0x14, 0x93,
	0x52, 0x53, 0xd5, 0x1c, 0x7e, 0xe9, 0x0f, 0x67, 0xc7, 0x51, 0xd0, 0x13, 0x58, 0xc6, 0x84, 0xd0,
	0xb2, 0x4c, 0xb5, 0x38, 0xa5, 0x3c, 0x3d, 0x66, 0x05, 0xb5, 0x3b, 0xfd, 0xa5, 0xeb, 0x39, 0xce,
	0xa1, 0xa1, 0xbc, 0x64, 0x05, 0x0d, 0x7f, 0x05, 0xd0, 0x3d, 0x62, 0xf2, 0x94, 0xf1, 0x6b, 0xf6,
	0xdf, 0x06, 0x74, 0xae, 0xa3, 0x3c, 0x93, 0x82, 0x71, 0xdd, 0x1c, 0xe0, 0xfc, 0xf2, 0x17, 0x9e,
	0x83, 0x1e, 0x40, 0xcf, 0xfe, 0x92, 0x29, 0xcb, 0xd2, 0xcd, 0xc1, 0xf6, 0x88, 0x69, 0x6b, 0xe9,
	0xff, 0xc9, 0xa2, 0x2d, 0xbf, 0xce, 0x5c, 0x11, 0xbd, 0x81, 0x1b, 0xe5, 0x09, 0x56, 0x34, 0x4b,
	0x4b, 0x89, 0xb9, 0x71, 0x5f, 0xd3, 0x33, 0xbd, 0xda, 0xb6, 0xf6, 0xf7, 0x1b, 0xf6, 0x0f, 0x85,
	0x28, 0x3e, 0xe2, 0x62, 0x42, 0x93, 0x65, 0x27, 0x3b, 0x90, 0xd8, 0x2c, 0x69, 0x44, 0x61, 0x0e,
	0xbd, 0xdd, 0x9a, 0xe3, 0x31, 0x23, 0x1f, 0xb4, 0x5f, 0x7c, 0x03, 0xe6, 0x0b, 0x36, 0x52, 0x58,
	0xd5, 0xcd, 0x75, 0xa7, 0x08, 0x8a, 0xff, 0xf1, 0xb9, 0x4c, 0x1f, 0x4b, 0x38, 0x81, 0xc5, 0x5d,
	0xac, 0x71, 0x26, 0xf2, 0x6b, 0xfa, 0xfb, 0x08, 0xba, 0x3e, 0x98, 0xa9, 0x7d, 0xcf, 0x0d, 0x67,
	0x3b, 0x1e, 0x7e, 0x8f, 0xc7, 0x34, 0x24, 0x80, 0x4c, 0x28, 0xe8, 0x81, 0xab, 0xf9, 0xde, 0xef,
	0xa0, 0x7b, 0x31, 0xe1, 0x3e, 0x55, 0x6b, 0x3e, 0x55, 0x58, 0x32, 0x13, 0x26, 0xf3, 0x47, 0x10,
	0xbd, 0x52, 0x92, 0x78, 0xed, 0x10, 0x4c, 0x8f, 0xff, 0xbe, 0x04, 0xad, 0xa5, 0x20, 0xe9, 0xe4,
	0x17, 0x80, 0x2d, 0xd8, 0x60, 0xc2, 0x89, 0xa5, 0x12, 0x67, 0xf5, 0xd5, 0xe9, 0x1c, 0x82, 0x9d,
	0x64, 0xdf, 0x18, 0xb4, 0x1f, 0x1c, 0xb5, 0xaa, 0xc1, 0x68, 0xce, 0xba, 0xf5, 0xf8, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x3b, 0x5b, 0x9f, 0xbc, 0x04, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/service/accesslog/v2/als.proto

It has these top-level messages:
	StreamAccessLogsResponse
	StreamAccessLogsMessage
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import envoy_config_filter_accesslog_v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/filter/accesslog/v2"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Empty response for the StreamAccessLogs API. Will never be sent. See below.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type StreamAccessLogsResponse struct {
}

func (m *StreamAccessLogsResponse) Reset()                    { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()               {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// [#proto-status: experimental]
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever incude one type of log entry.
	//
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
}

func (m *StreamAccessLogsMessage) Reset()                    { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string            { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()               {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,oneof"`
}
type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}
func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries()  {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamAccessLogsMessage_OneofMarshaler, _StreamAccessLogsMessage_OneofUnmarshaler, _StreamAccessLogsMessage_OneofSizer, []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

func _StreamAccessLogsMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpLogs); err != nil {
			return err
		}
	case *StreamAccessLogsMessage_TcpLogs:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpLogs); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamAccessLogsMessage.LogEntries has unexpected type %T", x)
	}
	return nil
}

func _StreamAccessLogsMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamAccessLogsMessage)
	switch tag {
	case 2: // log_entries.http_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_HTTPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_HttpLogs{msg}
		return true, err
	case 3: // log_entries.tcp_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_TCPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_TcpLogs{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamAccessLogsMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		s := proto.Size(x.HttpLogs)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamAccessLogsMessage_TcpLogs:
		s := proto.Size(x.TcpLogs)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamAccessLogsMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node *envoy_api_v2_core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	// The friendly name of the log configured in AccessLogServiceConfig.
	LogName string `protobuf:"bytes,2,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) GetNode() *envoy_api_v2_core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry []*envoy_config_filter_accesslog_v2.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry" json:"log_entry,omitempty"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*envoy_config_filter_accesslog_v2.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry []*envoy_config_filter_accesslog_v2.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry" json:"log_entry,omitempty"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*envoy_config_filter_accesslog_v2.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v2.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.TCPAccessLogEntries")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccessLogService service

type AccessLogServiceClient interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AccessLogService_serviceDesc.Streams[0], c.cc, "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AccessLogService service

type AccessLogServiceServer interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v2.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v2/als.proto",
}

func init() { proto.RegisterFile("envoy/service/accesslog/v2/als.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe7, 0x74, 0xa5, 0xed, 0xeb, 0xa5, 0x32, 0x93, 0x5a, 0x45, 0x1c, 0xaa, 0x69, 0x87,
	0x9e, 0x1c, 0x94, 0x4e, 0xe2, 0x86, 0x44, 0x11, 0xa2, 0x48, 0x30, 0xa1, 0xac, 0x27, 0x04, 0x4c,
	0x5e, 0xf2, 0x1a, 0x2c, 0xd2, 0x38, 0xb2, 0x4d, 0xa4, 0x1e, 0xb9, 0x21, 0x8e, 0x1c, 0xf8, 0x2a,
	0x20, 0x4e, 0xfb, 0x3a, 0xfb, 0x16, 0xc8, 0x4e, 0x16, 0xa0, 0x6b, 0x85, 0xd6, 0x5b, 0x14, 0xbf,
	0xff, 0xef, 0xe7, 0xe7, 0x67, 0xc3, 0x09, 0xe6, 0xa5, 0x5c, 0x07, 0x1a, 0x55, 0x29, 0x62, 0x0c,
	0x78, 0x1c, 0xa3, 0xd6, 0x99, 0x4c, 0x83, 0x32, 0x0c, 0x78, 0xa6, 0x59, 0xa1, 0xa4, 0x91, 0xd4,
	0x77, 0x55, 0xac, 0xae, 0x62, 0x4d, 0x15, 0x2b, 0x43, 0xff, 0x41, 0x45, 0xe0, 0x85, 0xb0, 0x99,
	0x58, 0x2a, 0x0c, 0x2e, 0xb9, 0xc6, 0x2a, 0xe9, 0x3f, 0xac, 0x56, 0x63, 0x99, 0x2f, 0x45, 0x1a,
	0x2c, 0x45, 0x66, 0x50, 0x6d, 0x58, 0x1a, 0x58, 0x95, 0x18, 0x96, 0x3c, 0x13, 0x09, 0x37, 0x18,
	0xdc, 0x7c, 0x54, 0x0b, 0xc7, 0x3e, 0x8c, 0xce, 0x8d, 0x42, 0xbe, 0x7a, 0xe2, 0x12, 0x2f, 0x65,
	0xaa, 0x23, 0xd4, 0x85, 0xcc, 0x35, 0x1e, 0xff, 0x68, 0xc3, 0x70, 0x73, 0xf1, 0x15, 0x6a, 0xcd,
	0x53, 0xa4, 0xef, 0x01, 0x44, 0x82, 0xb9, 0x11, 0x4b, 0x81, 0x6a, 0x44, 0xc6, 0x64, 0xd2, 0x0f,
	0x1f, 0xb3, 0xdd, 0x1d, 0xb1, 0x1d, 0x20, 0xf6, 0xa2, 0xa1, 0x44, 0x7f, 0x11, 0x69, 0x0a, 0xbd,
	0x0f, 0xc6, 0x14, 0x17, 0x99, 0x4c, 0xf5, 0xc8, 0x73, 0xf8, 0xf9, 0x3e, 0xf8, 0xf9, 0x62, 0xf1,
	0xba, 0xf9, 0xfb, 0x2c, 0x37, 0x4a, 0xa0, 0x9e, 0x1f, 0x44, 0x5d, 0x0b, 0xb7, 0x75, 0x34, 0x81,
	0xae, 0x89, 0x6b, 0x4f, 0xcb, 0x79, 0x9e, 0xef, 0xe3, 0x59, 0x3c, 0xdd, 0xa6, 0xe9, 0x98, 0xd8,
	0x59, 0xfc, 0x8f, 0x00, 0x7f, 0x1a, 0xa5, 0x8f, 0xe0, 0x30, 0x97, 0x09, 0xd6, 0xc7, 0x36, 0xac,
	0x7d, 0xbc, 0x10, 0xd6, 0x60, 0x87, 0xcd, 0xce, 0x64, 0x82, 0x33, 0xf8, 0x75, 0x7d, 0xd5, 0x6a,
	0x7f, 0x25, 0xde, 0x80, 0x44, 0x2e, 0x40, 0x4f, 0xa0, 0x9b, 0xc9, 0xf4, 0x22, 0xe7, 0x2b, 0x74,
	0x87, 0xd2, 0x9b, 0xf5, 0x6c, 0xcd, 0xa1, 0xf2, 0xc6, 0x24, 0xea, 0x64, 0x32, 0x3d, 0xe3, 0x2b,
	0xf4, 0x3f, 0xc1, 0xd1, 0xb6, 0xb6, 0xe9, 0x3b, 0xe8, 0xd9, 0x34, 0xe6, 0x46, 0xad, 0x47, 0x64,
	0xdc, 0x9a, 0xf4, 0xc3, 0xd3, 0xda, 0x5d, 0x5d, 0x25, 0x56, 0x5d, 0xa5, 0x7f, 0x3b, 0xbe, 0x85,
	0x5a, 0xd7, 0x1b, 0xfb, 0x46, 0xbc, 0x2e, 0x89, 0xec, 0x86, 0xdc, 0x5f, 0x5f, 0xc3, 0xfd, 0x2d,
	0xa7, 0x40, 0xdf, 0xde, 0xb6, 0x4e, 0xff, 0x6f, 0xdd, 0x24, 0xed, 0x90, 0xce, 0x8e, 0xa0, 0x7f,
	0x43, 0xb7, 0xb2, 0xf6, 0xcf, 0xeb, 0xab, 0x16, 0x09, 0xbf, 0x13, 0x18, 0x34, 0xf1, 0xf3, 0x6a,
	0x8e, 0xf4, 0x33, 0x81, 0xc1, 0xe6, 0xf8, 0xe8, 0x74, 0x8f, 0x61, 0xfb, 0xa7, 0x77, 0x09, 0x35,
	0xcf, 0xe9, 0x60, 0x42, 0x66, 0x9d, 0x37, 0x5e, 0x19, 0x7e, 0x21, 0xe4, 0xf2, 0x9e, 0x7b, 0x7e,
	0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x6b, 0xdd, 0x4d, 0x2b, 0x04, 0x00, 0x00,
}
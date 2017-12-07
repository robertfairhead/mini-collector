// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PublishRequest
	PublishResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type PublishRequest struct {
	UnixTime      uint64 `protobuf:"varint,1,opt,name=unix_time,json=unixTime" json:"unix_time,omitempty"`
	Running       bool   `protobuf:"varint,2,opt,name=running" json:"running,omitempty"`
	MilliCpuUsage uint64 `protobuf:"varint,3,opt,name=milli_cpu_usage,json=milliCpuUsage" json:"milli_cpu_usage,omitempty"`
	MemoryTotalMb uint64 `protobuf:"varint,4,opt,name=memory_total_mb,json=memoryTotalMb" json:"memory_total_mb,omitempty"`
	MemoryRssMb   uint64 `protobuf:"varint,5,opt,name=memory_rss_mb,json=memoryRssMb" json:"memory_rss_mb,omitempty"`
	MemoryLimitMb uint64 `protobuf:"varint,6,opt,name=memory_limit_mb,json=memoryLimitMb" json:"memory_limit_mb,omitempty"`
	DiskUsageMb   uint64 `protobuf:"varint,7,opt,name=disk_usage_mb,json=diskUsageMb" json:"disk_usage_mb,omitempty"`
	DiskLimitMb   uint64 `protobuf:"varint,8,opt,name=disk_limit_mb,json=diskLimitMb" json:"disk_limit_mb,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishRequest) GetUnixTime() uint64 {
	if m != nil {
		return m.UnixTime
	}
	return 0
}

func (m *PublishRequest) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *PublishRequest) GetMilliCpuUsage() uint64 {
	if m != nil {
		return m.MilliCpuUsage
	}
	return 0
}

func (m *PublishRequest) GetMemoryTotalMb() uint64 {
	if m != nil {
		return m.MemoryTotalMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryRssMb() uint64 {
	if m != nil {
		return m.MemoryRssMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryLimitMb() uint64 {
	if m != nil {
		return m.MemoryLimitMb
	}
	return 0
}

func (m *PublishRequest) GetDiskUsageMb() uint64 {
	if m != nil {
		return m.DiskUsageMb
	}
	return 0
}

func (m *PublishRequest) GetDiskLimitMb() uint64 {
	if m != nil {
		return m.DiskLimitMb
	}
	return 0
}

type PublishResponse struct {
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "PublishResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Aggregator service

type AggregatorClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type aggregatorClient struct {
	cc *grpc.ClientConn
}

func NewAggregatorClient(cc *grpc.ClientConn) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/Aggregator/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Aggregator service

type AggregatorServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
}

func RegisterAggregatorServer(s *grpc.Server, srv AggregatorServer) {
	s.RegisterService(&_Aggregator_serviceDesc, srv)
}

func _Aggregator_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aggregator/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Aggregator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Aggregator_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x40, 0x69, 0x29, 0x4d, 0x7a, 0x08, 0x02, 0x9e, 0x22, 0x58, 0xaa, 0x0c, 0xa8, 0x53, 0x06,
	0x58, 0x59, 0x10, 0x2b, 0x95, 0x50, 0x54, 0x66, 0x2b, 0x01, 0x2b, 0x9c, 0xf0, 0x17, 0x3e, 0x5b,
	0x82, 0xdf, 0xc4, 0x9f, 0x44, 0x76, 0x42, 0x09, 0x8c, 0x7e, 0x7a, 0xf7, 0x74, 0x3a, 0xc3, 0xaa,
	0xb5, 0x58, 0x5b, 0x67, 0xbc, 0xa9, 0xbe, 0xe6, 0x70, 0xfa, 0x18, 0x3a, 0x89, 0xf4, 0xda, 0x88,
	0xf7, 0x20, 0xc8, 0xb3, 0x4b, 0x58, 0x05, 0x8d, 0x1f, 0xdc, 0xa3, 0x12, 0xe5, 0x6c, 0x3d, 0xdb,
	0x2c, 0x9a, 0x3c, 0x82, 0x1d, 0x2a, 0xc1, 0x4a, 0xc8, 0x5c, 0xd0, 0x1a, 0x75, 0x5f, 0xce, 0xd7,
	0xb3, 0x4d, 0xde, 0xfc, 0x3c, 0xd9, 0x15, 0x14, 0x0a, 0xa5, 0x44, 0xfe, 0x6c, 0x03, 0x0f, 0xd4,
	0xf6, 0xa2, 0x3c, 0x4c, 0xc3, 0x27, 0x09, 0xdf, 0xdb, 0xf0, 0x14, 0x61, 0xf2, 0x84, 0x32, 0xee,
	0x93, 0x7b, 0xe3, 0x5b, 0xc9, 0x55, 0x57, 0x2e, 0x46, 0x2f, 0xe1, 0x5d, 0xa4, 0xdb, 0x8e, 0x55,
	0x30, 0x02, 0xee, 0x88, 0xa2, 0x75, 0x94, 0xac, 0xe3, 0x01, 0x36, 0x44, 0xdb, 0x6e, 0xd2, 0x92,
	0xa8, 0xd0, 0x47, 0x6b, 0x39, 0x6d, 0x3d, 0x44, 0x3a, 0xb4, 0x5e, 0x90, 0xde, 0x86, 0xb5, 0xa2,
	0x95, 0x0d, 0xad, 0x08, 0xd3, 0x56, 0x13, 0x67, 0x5f, 0xca, 0x7f, 0x9d, 0xb1, 0x53, 0x9d, 0x43,
	0xb1, 0x3f, 0x16, 0x59, 0xa3, 0x49, 0x5c, 0xdf, 0x02, 0xdc, 0xf5, 0xbd, 0x13, 0x7d, 0xeb, 0x8d,
	0x63, 0x35, 0x64, 0xa3, 0xc0, 0x8a, 0xfa, 0xef, 0x5d, 0x2f, 0xce, 0xea, 0x7f, 0xb3, 0xd5, 0x41,
	0xb7, 0x4c, 0xbf, 0x70, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x59, 0x87, 0xe8, 0x5b, 0x92, 0x01,
	0x00, 0x00,
}

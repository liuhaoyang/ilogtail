// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TraceSegmentServiceClient is the client API for TraceSegmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceSegmentServiceClient interface {
	Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error)
}

type traceSegmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceSegmentServiceClient(cc grpc.ClientConnInterface) TraceSegmentServiceClient {
	return &traceSegmentServiceClient{cc}
}

func (c *traceSegmentServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (TraceSegmentService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &TraceSegmentService_ServiceDesc.Streams[0], "/TraceSegmentService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceSegmentServiceCollectClient{stream}
	return x, nil
}

type TraceSegmentService_CollectClient interface {
	Send(*UpstreamSegment) error
	CloseAndRecv() (*Downstream, error)
	grpc.ClientStream
}

type traceSegmentServiceCollectClient struct {
	grpc.ClientStream
}

func (x *traceSegmentServiceCollectClient) Send(m *UpstreamSegment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectClient) CloseAndRecv() (*Downstream, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Downstream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentServiceServer is the server API for TraceSegmentService service.
// All implementations should embed UnimplementedTraceSegmentServiceServer
// for forward compatibility
type TraceSegmentServiceServer interface {
	Collect(TraceSegmentService_CollectServer) error
}

// UnimplementedTraceSegmentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTraceSegmentServiceServer struct {
}

func (UnimplementedTraceSegmentServiceServer) Collect(TraceSegmentService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeTraceSegmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceSegmentServiceServer will
// result in compilation errors.
type UnsafeTraceSegmentServiceServer interface {
	mustEmbedUnimplementedTraceSegmentServiceServer()
}

func RegisterTraceSegmentServiceServer(s grpc.ServiceRegistrar, srv TraceSegmentServiceServer) {
	s.RegisterService(&TraceSegmentService_ServiceDesc, srv)
}

func _TraceSegmentService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceSegmentServiceServer).Collect(&traceSegmentServiceCollectServer{stream})
}

type TraceSegmentService_CollectServer interface {
	SendAndClose(*Downstream) error
	Recv() (*UpstreamSegment, error)
	grpc.ServerStream
}

type traceSegmentServiceCollectServer struct {
	grpc.ServerStream
}

func (x *traceSegmentServiceCollectServer) SendAndClose(m *Downstream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceSegmentServiceCollectServer) Recv() (*UpstreamSegment, error) {
	m := new(UpstreamSegment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceSegmentService_ServiceDesc is the grpc.ServiceDesc for TraceSegmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceSegmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TraceSegmentService",
	HandlerType: (*TraceSegmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _TraceSegmentService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "language-agent/TraceSegmentService.proto",
}

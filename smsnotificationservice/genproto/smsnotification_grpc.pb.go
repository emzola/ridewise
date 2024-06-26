// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/smsnotificationservice/smsnotification.proto

package genproto

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

const (
	SMSNotificationService_SendSMS_FullMethodName = "/sms.SMSNotificationService/SendSMS"
)

// SMSNotificationServiceClient is the client API for SMSNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSNotificationServiceClient interface {
	SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error)
}

type sMSNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSNotificationServiceClient(cc grpc.ClientConnInterface) SMSNotificationServiceClient {
	return &sMSNotificationServiceClient{cc}
}

func (c *sMSNotificationServiceClient) SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error) {
	out := new(SendSMSResponse)
	err := c.cc.Invoke(ctx, SMSNotificationService_SendSMS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSNotificationServiceServer is the server API for SMSNotificationService service.
// All implementations must embed UnimplementedSMSNotificationServiceServer
// for forward compatibility
type SMSNotificationServiceServer interface {
	SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
	mustEmbedUnimplementedSMSNotificationServiceServer()
}

// UnimplementedSMSNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSMSNotificationServiceServer struct {
}

func (UnimplementedSMSNotificationServiceServer) SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMS not implemented")
}
func (UnimplementedSMSNotificationServiceServer) mustEmbedUnimplementedSMSNotificationServiceServer() {
}

// UnsafeSMSNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSNotificationServiceServer will
// result in compilation errors.
type UnsafeSMSNotificationServiceServer interface {
	mustEmbedUnimplementedSMSNotificationServiceServer()
}

func RegisterSMSNotificationServiceServer(s grpc.ServiceRegistrar, srv SMSNotificationServiceServer) {
	s.RegisterService(&SMSNotificationService_ServiceDesc, srv)
}

func _SMSNotificationService_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSNotificationServiceServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SMSNotificationService_SendSMS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSNotificationServiceServer).SendSMS(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SMSNotificationService_ServiceDesc is the grpc.ServiceDesc for SMSNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMSNotificationService",
	HandlerType: (*SMSNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSMS",
			Handler:    _SMSNotificationService_SendSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/smsnotificationservice/smsnotification.proto",
}

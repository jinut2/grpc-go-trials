// Code generated by protoc-gen-go.
// source: protos/user.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/user.proto

It has these top-level messages:
	GetUserDetailsRequest
	GetUserDetailsResponse
	SetUserDetailsRequest
	SetUserDetailsResponse
*/
package protos

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

type GetUserDetailsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserDetailsRequest) Reset()                    { *m = GetUserDetailsRequest{} }
func (m *GetUserDetailsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserDetailsRequest) ProtoMessage()               {}
func (*GetUserDetailsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetUserDetailsResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *GetUserDetailsResponse) Reset()                    { *m = GetUserDetailsResponse{} }
func (m *GetUserDetailsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserDetailsResponse) ProtoMessage()               {}
func (*GetUserDetailsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SetUserDetailsRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *SetUserDetailsRequest) Reset()                    { *m = SetUserDetailsRequest{} }
func (m *SetUserDetailsRequest) String() string            { return proto.CompactTextString(m) }
func (*SetUserDetailsRequest) ProtoMessage()               {}
func (*SetUserDetailsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SetUserDetailsResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SetUserDetailsResponse) Reset()                    { *m = SetUserDetailsResponse{} }
func (m *SetUserDetailsResponse) String() string            { return proto.CompactTextString(m) }
func (*SetUserDetailsResponse) ProtoMessage()               {}
func (*SetUserDetailsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*GetUserDetailsRequest)(nil), "protos.GetUserDetailsRequest")
	proto.RegisterType((*GetUserDetailsResponse)(nil), "protos.GetUserDetailsResponse")
	proto.RegisterType((*SetUserDetailsRequest)(nil), "protos.SetUserDetailsRequest")
	proto.RegisterType((*SetUserDetailsResponse)(nil), "protos.SetUserDetailsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for UserService service

type UserServiceClient interface {
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	SetUserDetails(ctx context.Context, in *SetUserDetailsRequest, opts ...grpc.CallOption) (*SetUserDetailsResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := grpc.Invoke(ctx, "/protos.UserService/GetUserDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetUserDetails(ctx context.Context, in *SetUserDetailsRequest, opts ...grpc.CallOption) (*SetUserDetailsResponse, error) {
	out := new(SetUserDetailsResponse)
	err := grpc.Invoke(ctx, "/protos.UserService/SetUserDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
	SetUserDetails(context.Context, *SetUserDetailsRequest) (*SetUserDetailsResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserService/SetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetUserDetails(ctx, req.(*SetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
		{
			MethodName: "SetUserDetails",
			Handler:    _UserService_SetUserDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("protos/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0xb3, 0x85, 0xd8, 0x20, 0x42, 0x4a, 0xea,
	0x5c, 0xa2, 0xee, 0xa9, 0x25, 0xa1, 0x40, 0x09, 0x97, 0xd4, 0x92, 0xc4, 0xcc, 0x9c, 0xe2, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x20, 0x4b, 0xc9, 0x8d, 0x4b, 0x0c, 0x5d, 0x61, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0x2a, 0x54, 0x39, 0x8c, 0x2b,
	0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x06, 0xb3, 0x95, 0xac, 0xb9,
	0x44, 0x83, 0x89, 0xb1, 0x10, 0xab, 0x66, 0x23, 0x2e, 0xb1, 0x60, 0x12, 0x1d, 0x61, 0xb4, 0x9e,
	0x91, 0x8b, 0x1b, 0xa4, 0x23, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9f, 0x8b, 0x0f,
	0xd5, 0x23, 0x42, 0xb2, 0x90, 0x30, 0x29, 0xd6, 0xc3, 0x1a, 0x12, 0x52, 0x72, 0xb8, 0xa4, 0xa1,
	0x56, 0x03, 0x0d, 0x0c, 0xc6, 0x61, 0x60, 0x30, 0x7e, 0x03, 0xb1, 0xfb, 0x25, 0x09, 0x12, 0x37,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x05, 0x57, 0x46, 0xb7, 0x01, 0x00, 0x00,
}
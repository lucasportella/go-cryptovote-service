// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: voteservice/service.proto

package voteservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_voteservice_service_proto protoreflect.FileDescriptor

var file_voteservice_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x6f, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1a, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf4, 0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x12, 0x20, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x67, 0x65, 0x72, 0x69,
	0x6f, 0x34, 0x31, 0x30, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x76, 0x6f, 0x74, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_voteservice_service_proto_goTypes = []interface{}{
	(*GetAllCriptoRequest)(nil),  // 0: voteservice.GetAllCriptoRequest
	(*VoteRequest)(nil),          // 1: voteservice.VoteRequest
	(*RemoveVoteRequest)(nil),    // 2: voteservice.RemoveVoteRequest
	(*GetAllCriptoResponse)(nil), // 3: voteservice.GetAllCriptoResponse
	(*VoteResponse)(nil),         // 4: voteservice.VoteResponse
	(*RemoveVoteResponse)(nil),   // 5: voteservice.RemoveVoteResponse
}
var file_voteservice_service_proto_depIdxs = []int32{
	0, // 0: voteservice.VoteService.GetAllCripto:input_type -> voteservice.GetAllCriptoRequest
	1, // 1: voteservice.VoteService.Vote:input_type -> voteservice.VoteRequest
	2, // 2: voteservice.VoteService.RemoveVote:input_type -> voteservice.RemoveVoteRequest
	3, // 3: voteservice.VoteService.GetAllCripto:output_type -> voteservice.GetAllCriptoResponse
	4, // 4: voteservice.VoteService.Vote:output_type -> voteservice.VoteResponse
	5, // 5: voteservice.VoteService.RemoveVote:output_type -> voteservice.RemoveVoteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_voteservice_service_proto_init() }
func file_voteservice_service_proto_init() {
	if File_voteservice_service_proto != nil {
		return
	}
	file_voteservice_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voteservice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voteservice_service_proto_goTypes,
		DependencyIndexes: file_voteservice_service_proto_depIdxs,
	}.Build()
	File_voteservice_service_proto = out.File
	file_voteservice_service_proto_rawDesc = nil
	file_voteservice_service_proto_goTypes = nil
	file_voteservice_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VoteServiceClient interface {
	GetAllCripto(ctx context.Context, in *GetAllCriptoRequest, opts ...grpc.CallOption) (*GetAllCriptoResponse, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	RemoveVote(ctx context.Context, in *RemoveVoteRequest, opts ...grpc.CallOption) (*RemoveVoteResponse, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) GetAllCripto(ctx context.Context, in *GetAllCriptoRequest, opts ...grpc.CallOption) (*GetAllCriptoResponse, error) {
	out := new(GetAllCriptoResponse)
	err := c.cc.Invoke(ctx, "/voteservice.VoteService/GetAllCripto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, "/voteservice.VoteService/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) RemoveVote(ctx context.Context, in *RemoveVoteRequest, opts ...grpc.CallOption) (*RemoveVoteResponse, error) {
	out := new(RemoveVoteResponse)
	err := c.cc.Invoke(ctx, "/voteservice.VoteService/RemoveVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServiceServer is the server API for VoteService service.
type VoteServiceServer interface {
	GetAllCripto(context.Context, *GetAllCriptoRequest) (*GetAllCriptoResponse, error)
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	RemoveVote(context.Context, *RemoveVoteRequest) (*RemoveVoteResponse, error)
}

// UnimplementedVoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (*UnimplementedVoteServiceServer) GetAllCripto(context.Context, *GetAllCriptoRequest) (*GetAllCriptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCripto not implemented")
}
func (*UnimplementedVoteServiceServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (*UnimplementedVoteServiceServer) RemoveVote(context.Context, *RemoveVoteRequest) (*RemoveVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVote not implemented")
}

func RegisterVoteServiceServer(s *grpc.Server, srv VoteServiceServer) {
	s.RegisterService(&_VoteService_serviceDesc, srv)
}

func _VoteService_GetAllCripto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCriptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetAllCripto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voteservice.VoteService/GetAllCripto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetAllCripto(ctx, req.(*GetAllCriptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voteservice.VoteService/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_RemoveVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).RemoveVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voteservice.VoteService/RemoveVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).RemoveVote(ctx, req.(*RemoveVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "voteservice.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCripto",
			Handler:    _VoteService_GetAllCripto_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _VoteService_Vote_Handler,
		},
		{
			MethodName: "RemoveVote",
			Handler:    _VoteService_RemoveVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voteservice/service.proto",
}

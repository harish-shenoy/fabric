// Code generated by protoc-gen-go.
// source: peer/fabric_service.proto
// DO NOT EDIT!

package peer

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Endorser service

type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := grpc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Endorser service

type EndorserServer interface {
	ProcessProposal(context.Context, *Proposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor12,
}

func init() { proto.RegisterFile("peer/fabric_service.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0x4d, 0x2d,
	0xd2, 0x4f, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0x52, 0xc8, 0x4a, 0x0a, 0x8a,
	0xf2, 0x0b, 0xf2, 0x8b, 0x13, 0x73, 0x20, 0x6a, 0xa4, 0x94, 0xb1, 0xc9, 0xc5, 0x17, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0x43, 0x0d, 0x32, 0xf2, 0xe6, 0xe2, 0x70, 0xcd, 0x4b, 0xc9, 0x2f, 0x2a,
	0x4e, 0x2d, 0x12, 0xb2, 0xe7, 0xe2, 0x0f, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0x0e, 0x80, 0xaa,
	0x16, 0x12, 0x80, 0x28, 0x2b, 0xd6, 0x83, 0x89, 0x48, 0x49, 0xa0, 0x8b, 0x04, 0x41, 0x0d, 0x54,
	0x62, 0x70, 0xd2, 0x8e, 0xd2, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0xcf, 0xa8, 0x2c, 0x48, 0x2d, 0xca, 0x49, 0x4d, 0x49, 0x87, 0xbb, 0x42, 0x1f, 0xa2, 0x55, 0x1f,
	0xe4, 0xb0, 0x24, 0x88, 0x17, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x4a, 0x29, 0x52,
	0xe6, 0x00, 0x00, 0x00,
}
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gstake/lsgridiron/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryLiquidValidatorsRequest is the request type for the
// Query/LiquidValidators RPC method.
type QueryLiquidValidatorsRequest struct {
}

func (m *QueryLiquidValidatorsRequest) Reset()         { *m = QueryLiquidValidatorsRequest{} }
func (m *QueryLiquidValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidValidatorsRequest) ProtoMessage()    {}
func (*QueryLiquidValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{2}
}
func (m *QueryLiquidValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidValidatorsRequest.Merge(m, src)
}
func (m *QueryLiquidValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidValidatorsRequest proto.InternalMessageInfo

// QueryLiquidValidatorsResponse is the response type for the
// Query/LiquidValidators RPC method.
type QueryLiquidValidatorsResponse struct {
	LiquidValidators []LiquidValidatorState `protobuf:"bytes,1,rep,name=liquid_validators,json=liquidValidators,proto3" json:"liquid_validators"`
}

func (m *QueryLiquidValidatorsResponse) Reset()         { *m = QueryLiquidValidatorsResponse{} }
func (m *QueryLiquidValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLiquidValidatorsResponse) ProtoMessage()    {}
func (*QueryLiquidValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{3}
}
func (m *QueryLiquidValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLiquidValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLiquidValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLiquidValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLiquidValidatorsResponse.Merge(m, src)
}
func (m *QueryLiquidValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLiquidValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLiquidValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLiquidValidatorsResponse proto.InternalMessageInfo

func (m *QueryLiquidValidatorsResponse) GetLiquidValidators() []LiquidValidatorState {
	if m != nil {
		return m.LiquidValidators
	}
	return nil
}

// QueryStatesRequest is the request type for the Query/States RPC method.
type QueryStatesRequest struct {
}

func (m *QueryStatesRequest) Reset()         { *m = QueryStatesRequest{} }
func (m *QueryStatesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStatesRequest) ProtoMessage()    {}
func (*QueryStatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{4}
}
func (m *QueryStatesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStatesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatesRequest.Merge(m, src)
}
func (m *QueryStatesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatesRequest proto.InternalMessageInfo

// QueryStatesResponse is the response type for the Query/States RPC method.
type QueryStatesResponse struct {
	NetAmountState NetAmountState `protobuf:"bytes,1,opt,name=net_amount_state,json=netAmountState,proto3" json:"net_amount_state"`
}

func (m *QueryStatesResponse) Reset()         { *m = QueryStatesResponse{} }
func (m *QueryStatesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStatesResponse) ProtoMessage()    {}
func (*QueryStatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f125a17ab68cc0, []int{5}
}
func (m *QueryStatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatesResponse.Merge(m, src)
}
func (m *QueryStatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatesResponse proto.InternalMessageInfo

func (m *QueryStatesResponse) GetNetAmountState() NetAmountState {
	if m != nil {
		return m.NetAmountState
	}
	return NetAmountState{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "gstake.lsgridiron.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "gstake.lsgridiron.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryLiquidValidatorsRequest)(nil), "gstake.lsgridiron.v1beta1.QueryLiquidValidatorsRequest")
	proto.RegisterType((*QueryLiquidValidatorsResponse)(nil), "gstake.lsgridiron.v1beta1.QueryLiquidValidatorsResponse")
	proto.RegisterType((*QueryStatesRequest)(nil), "gstake.lsgridiron.v1beta1.QueryStatesRequest")
	proto.RegisterType((*QueryStatesResponse)(nil), "gstake.lsgridiron.v1beta1.QueryStatesResponse")
}

func init() {
	proto.RegisterFile("gstake/lsgridiron/v1beta1/query.proto", fileDescriptor_c3f125a17ab68cc0)
}

var fileDescriptor_c3f125a17ab68cc0 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x63, 0x06, 0x3d, 0x78, 0x12, 0x2a, 0x66, 0x87, 0x11, 0x8d, 0xb0, 0x05, 0x4d, 0x6c,
	0x87, 0xc4, 0x5a, 0x39, 0x8c, 0x1b, 0x62, 0x67, 0x04, 0x6c, 0x48, 0x93, 0xe0, 0x52, 0x39, 0xd4,
	0x32, 0x16, 0xa9, 0x9d, 0xc6, 0x4e, 0xc5, 0xae, 0xf0, 0x05, 0x10, 0x7c, 0x15, 0x3e, 0xc4, 0x8e,
	0x93, 0xb8, 0x70, 0x42, 0xa8, 0xe5, 0xc0, 0xc7, 0x40, 0xb1, 0xdd, 0xb4, 0x29, 0x4a, 0xda, 0xdd,
	0xa2, 0xf7, 0xfe, 0xef, 0xbd, 0x9f, 0xff, 0xef, 0x05, 0xee, 0x33, 0xa5, 0xc9, 0x07, 0x8a, 0x53,
	0xc5, 0x72, 0x3e, 0xe0, 0xb9, 0x14, 0x78, 0x7c, 0x94, 0x50, 0x4d, 0x8e, 0xf0, 0xa8, 0xa0, 0xf9,
	0x45, 0x9c, 0xe5, 0x52, 0x4b, 0x74, 0xcf, 0xca, 0xe2, 0xb9, 0x2c, 0x76, 0x32, 0x7f, 0x87, 0x49,
	0xc9, 0x52, 0x8a, 0x49, 0xc6, 0x31, 0x11, 0x42, 0x6a, 0xa2, 0xb9, 0x14, 0xca, 0x16, 0xfa, 0x51,
	0x73, 0xff, 0x94, 0x8f, 0x0a, 0x3e, 0x28, 0xd3, 0x5c, 0x30, 0x27, 0xdf, 0x62, 0x92, 0x49, 0xf3,
	0x89, 0xcb, 0x2f, 0x1b, 0x0d, 0xb7, 0x20, 0x3a, 0x2d, 0x61, 0x5e, 0x91, 0x9c, 0x0c, 0xd5, 0x19,
	0x1d, 0x15, 0x54, 0xe9, 0xf0, 0x1c, 0xde, 0xad, 0x45, 0x55, 0x26, 0x85, 0xa2, 0xe8, 0x29, 0xec,
	0x64, 0x26, 0xb2, 0x0d, 0x76, 0xc1, 0xc1, 0x66, 0x6f, 0x2f, 0x6e, 0x64, 0x8f, 0x6d, 0xe9, 0xc9,
	0xcd, 0xcb, 0x5f, 0x0f, 0xbc, 0x33, 0x57, 0x16, 0x06, 0x70, 0xc7, 0xf4, 0x7d, 0x6e, 0xf8, 0xce,
	0x49, 0xca, 0x07, 0x44, 0xcb, 0xbc, 0x9a, 0xfb, 0x19, 0xc0, 0xfb, 0x0d, 0x02, 0x87, 0x90, 0xc0,
	0x3b, 0xf6, 0x71, 0xfd, 0x71, 0x95, 0xdc, 0x06, 0xbb, 0x1b, 0x07, 0x9b, 0x3d, 0xdc, 0x42, 0xb3,
	0xd4, 0xef, 0xb5, 0x26, 0x9a, 0x3a, 0xb6, 0x6e, 0xba, 0x34, 0xab, 0xf2, 0xc4, 0xa8, 0x2a, 0xb6,
	0xcc, 0x79, 0x32, 0x8b, 0x3a, 0xa0, 0x37, 0xb0, 0x2b, 0xa8, 0xee, 0x93, 0xa1, 0x2c, 0x84, 0xee,
	0xab, 0x32, 0xe9, 0xdc, 0x39, 0x6c, 0xe1, 0x79, 0x41, 0xf5, 0x33, 0x53, 0xb1, 0x48, 0x72, 0x5b,
	0xd4, 0xa2, 0xbd, 0xbf, 0x1b, 0xf0, 0x96, 0x19, 0x89, 0xbe, 0x02, 0xd8, 0xb1, 0x86, 0xa2, 0xa8,
	0xa5, 0xeb, 0xff, 0x9b, 0xf4, 0xe3, 0x75, 0xe5, 0xf6, 0x39, 0xe1, 0xe1, 0xa7, 0x1f, 0x7f, 0xbe,
	0xdd, 0x78, 0x88, 0xf6, 0x70, 0xf3, 0x75, 0xd9, 0x65, 0xa2, 0xef, 0x00, 0x76, 0x97, 0xf7, 0x84,
	0x8e, 0x57, 0xcd, 0x6b, 0x58, 0xbd, 0xff, 0xe4, 0xfa, 0x85, 0x0e, 0x39, 0x32, 0xc8, 0x8f, 0xd0,
	0x7e, 0x0b, 0xf2, 0xfc, 0x58, 0x8c, 0x97, 0x76, 0x87, 0xab, 0xbd, 0xac, 0x5d, 0xc0, 0x6a, 0x2f,
	0xeb, 0xa7, 0xb1, 0x96, 0x97, 0xe6, 0x60, 0xd4, 0xc9, 0xe9, 0xe5, 0x24, 0x00, 0x57, 0x93, 0x00,
	0xfc, 0x9e, 0x04, 0xe0, 0xcb, 0x34, 0xf0, 0xae, 0xa6, 0x81, 0xf7, 0x73, 0x1a, 0x78, 0x6f, 0x8f,
	0x19, 0xd7, 0xef, 0x8b, 0x24, 0x7e, 0x27, 0x87, 0x78, 0x56, 0xfd, 0x52, 0x50, 0xd7, 0x32, 0x12,
	0x44, 0xf3, 0x31, 0xc5, 0x1f, 0x17, 0x7b, 0xeb, 0x8b, 0x8c, 0xaa, 0xa4, 0x63, 0x7e, 0xf0, 0xc7,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0xf1, 0x7e, 0x39, 0x87, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns parameters of the liquidstaking module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// LiquidValidators returns liquid validators with states of the liquidstaking
	// module.
	LiquidValidators(ctx context.Context, in *QueryLiquidValidatorsRequest, opts ...grpc.CallOption) (*QueryLiquidValidatorsResponse, error)
	// States returns states of the liquidstaking module.
	States(ctx context.Context, in *QueryStatesRequest, opts ...grpc.CallOption) (*QueryStatesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/gstake.lsgridiron.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidValidators(ctx context.Context, in *QueryLiquidValidatorsRequest, opts ...grpc.CallOption) (*QueryLiquidValidatorsResponse, error) {
	out := new(QueryLiquidValidatorsResponse)
	err := c.cc.Invoke(ctx, "/gstake.lsgridiron.v1beta1.Query/LiquidValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) States(ctx context.Context, in *QueryStatesRequest, opts ...grpc.CallOption) (*QueryStatesResponse, error) {
	out := new(QueryStatesResponse)
	err := c.cc.Invoke(ctx, "/gstake.lsgridiron.v1beta1.Query/States", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns parameters of the liquidstaking module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// LiquidValidators returns liquid validators with states of the liquidstaking
	// module.
	LiquidValidators(context.Context, *QueryLiquidValidatorsRequest) (*QueryLiquidValidatorsResponse, error)
	// States returns states of the liquidstaking module.
	States(context.Context, *QueryStatesRequest) (*QueryStatesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) LiquidValidators(ctx context.Context, req *QueryLiquidValidatorsRequest) (*QueryLiquidValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidValidators not implemented")
}
func (*UnimplementedQueryServer) States(ctx context.Context, req *QueryStatesRequest) (*QueryStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method States not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gstake.lsgridiron.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gstake.lsgridiron.v1beta1.Query/LiquidValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidValidators(ctx, req.(*QueryLiquidValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_States_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).States(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gstake.lsgridiron.v1beta1.Query/States",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).States(ctx, req.(*QueryStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gstake.lsgridiron.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "LiquidValidators",
			Handler:    _Query_LiquidValidators_Handler,
		},
		{
			MethodName: "States",
			Handler:    _Query_States_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gstake/lsgridiron/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryLiquidValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLiquidValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLiquidValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLiquidValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LiquidValidators) > 0 {
		for iNdEx := len(m.LiquidValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryStatesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStatesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStatesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryStatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.NetAmountState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryLiquidValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLiquidValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LiquidValidators) > 0 {
		for _, e := range m.LiquidValidators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryStatesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryStatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NetAmountState.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLiquidValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLiquidValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLiquidValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidValidators = append(m.LiquidValidators, LiquidValidatorState{})
			if err := m.LiquidValidators[len(m.LiquidValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStatesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStatesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStatesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAmountState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetAmountState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

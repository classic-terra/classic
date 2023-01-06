// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/abs/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{0}
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

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{1}
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

type QueryBreakFactorRequest struct {
}

func (m *QueryBreakFactorRequest) Reset()         { *m = QueryBreakFactorRequest{} }
func (m *QueryBreakFactorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBreakFactorRequest) ProtoMessage()    {}
func (*QueryBreakFactorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{2}
}
func (m *QueryBreakFactorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBreakFactorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBreakFactorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBreakFactorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBreakFactorRequest.Merge(m, src)
}
func (m *QueryBreakFactorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBreakFactorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBreakFactorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBreakFactorRequest proto.InternalMessageInfo

type QueryBreakFactorResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryBreakFactorResponse) Reset()         { *m = QueryBreakFactorResponse{} }
func (m *QueryBreakFactorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBreakFactorResponse) ProtoMessage()    {}
func (*QueryBreakFactorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{3}
}
func (m *QueryBreakFactorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBreakFactorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBreakFactorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBreakFactorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBreakFactorResponse.Merge(m, src)
}
func (m *QueryBreakFactorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBreakFactorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBreakFactorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBreakFactorResponse proto.InternalMessageInfo

func (m *QueryBreakFactorResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QueryWatchlistRequest struct {
}

func (m *QueryWatchlistRequest) Reset()         { *m = QueryWatchlistRequest{} }
func (m *QueryWatchlistRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWatchlistRequest) ProtoMessage()    {}
func (*QueryWatchlistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{4}
}
func (m *QueryWatchlistRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWatchlistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWatchlistRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWatchlistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWatchlistRequest.Merge(m, src)
}
func (m *QueryWatchlistRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWatchlistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWatchlistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWatchlistRequest proto.InternalMessageInfo

type QueryWatchlistResponse struct {
	Watchlist Watchlist `protobuf:"bytes,1,opt,name=watchlist,proto3" json:"watchlist"`
}

func (m *QueryWatchlistResponse) Reset()         { *m = QueryWatchlistResponse{} }
func (m *QueryWatchlistResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWatchlistResponse) ProtoMessage()    {}
func (*QueryWatchlistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea6c72040e49278, []int{5}
}
func (m *QueryWatchlistResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWatchlistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWatchlistResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWatchlistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWatchlistResponse.Merge(m, src)
}
func (m *QueryWatchlistResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWatchlistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWatchlistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWatchlistResponse proto.InternalMessageInfo

func (m *QueryWatchlistResponse) GetWatchlist() Watchlist {
	if m != nil {
		return m.Watchlist
	}
	return Watchlist{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "terra.abs.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "terra.abs.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryBreakFactorRequest)(nil), "terra.abs.v1beta1.QueryBreakFactorRequest")
	proto.RegisterType((*QueryBreakFactorResponse)(nil), "terra.abs.v1beta1.QueryBreakFactorResponse")
	proto.RegisterType((*QueryWatchlistRequest)(nil), "terra.abs.v1beta1.QueryWatchlistRequest")
	proto.RegisterType((*QueryWatchlistResponse)(nil), "terra.abs.v1beta1.QueryWatchlistResponse")
}

func init() { proto.RegisterFile("terra/abs/v1beta1/query.proto", fileDescriptor_eea6c72040e49278) }

var fileDescriptor_eea6c72040e49278 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0x86, 0x5b, 0xf5, 0x92, 0xdc, 0x31, 0x31, 0x3a, 0xa2, 0xf7, 0xde, 0x06, 0xc7, 0x6b, 0x15,
	0x45, 0x88, 0x1d, 0xc1, 0x85, 0x5b, 0x65, 0xe1, 0xd2, 0x28, 0x1b, 0x13, 0x36, 0x64, 0xda, 0x8c,
	0xa5, 0x91, 0x76, 0xca, 0xcc, 0x14, 0x65, 0xe1, 0xc6, 0x17, 0xd0, 0xc4, 0x8d, 0xef, 0xe1, 0x4b,
	0xb0, 0x24, 0x71, 0xe3, 0xca, 0x18, 0xf0, 0x41, 0x4c, 0x67, 0xa6, 0x35, 0x50, 0x88, 0xec, 0xa0,
	0xe7, 0x3f, 0xff, 0xff, 0x9d, 0x73, 0x5a, 0x70, 0x4b, 0x52, 0xce, 0x09, 0x26, 0xbe, 0xc0, 0xb3,
	0xae, 0x4f, 0x25, 0xe9, 0xe2, 0x69, 0x46, 0xf9, 0xdc, 0x4b, 0x39, 0x93, 0x0c, 0x5e, 0x53, 0x65,
	0x8f, 0xf8, 0xc2, 0x33, 0x65, 0xa7, 0x1e, 0xb2, 0x90, 0xa9, 0x2a, 0xce, 0x7f, 0x69, 0xa1, 0xd3,
	0x08, 0x19, 0x0b, 0x27, 0x14, 0x93, 0x34, 0xc2, 0x24, 0x49, 0x98, 0x24, 0x32, 0x62, 0x89, 0x30,
	0xd5, 0x76, 0xc0, 0x44, 0xcc, 0x04, 0xf6, 0x89, 0xa0, 0xda, 0xbf, 0x4c, 0x4b, 0x49, 0x18, 0x25,
	0x4a, 0x6c, 0xb4, 0xa8, 0x4a, 0x94, 0x12, 0x4e, 0xe2, 0xc2, 0xeb, 0x4e, 0xb5, 0xfe, 0x9e, 0xc8,
	0x60, 0x3c, 0x89, 0x84, 0xd4, 0x12, 0xb7, 0x0e, 0xe0, 0xeb, 0x3c, 0xe4, 0x95, 0xea, 0x1b, 0xd0,
	0x69, 0x46, 0x85, 0x74, 0x5f, 0x82, 0xeb, 0x1b, 0x4f, 0x45, 0xca, 0x12, 0x41, 0xe1, 0x53, 0x50,
	0xd3, 0xfe, 0xa7, 0xf6, 0xb9, 0xdd, 0xba, 0xdc, 0x3b, 0xf3, 0x2a, 0x33, 0x7b, 0xba, 0xa5, 0x7f,
	0x69, 0xf1, 0xeb, 0xb6, 0x35, 0x30, 0x72, 0xf7, 0x0c, 0x9c, 0x28, 0xbf, 0x3e, 0xa7, 0xe4, 0xdd,
	0x0b, 0x12, 0x48, 0xc6, 0x8b, 0xa8, 0xc7, 0xe0, 0xb4, 0x5a, 0x32, 0x79, 0x75, 0x70, 0x34, 0x23,
	0x93, 0x8c, 0xaa, 0xb8, 0xe3, 0x81, 0xfe, 0xe3, 0x9e, 0x80, 0x1b, 0xaa, 0xe3, 0x4d, 0x31, 0x4a,
	0x61, 0x35, 0x04, 0x37, 0xb7, 0x0b, 0xc6, 0xe8, 0x19, 0x38, 0x2e, 0x07, 0x37, 0xec, 0x8d, 0x1d,
	0xec, 0x65, 0xa3, 0xc1, 0xff, 0xd7, 0xd4, 0xfb, 0x7e, 0x11, 0x1c, 0x29, 0x73, 0xf8, 0x11, 0xd4,
	0xf4, 0x8c, 0xb0, 0xb9, 0xc3, 0xa2, 0xba, 0x4c, 0xe7, 0xfe, 0xff, 0x64, 0x1a, 0xd2, 0xbd, 0xf7,
	0xe9, 0xc7, 0x9f, 0xaf, 0x17, 0x10, 0x6c, 0x60, 0xa5, 0x7f, 0x14, 0xb3, 0x84, 0xce, 0x71, 0xc0,
	0x38, 0x55, 0x17, 0xd4, 0xab, 0x84, 0xdf, 0x6c, 0x70, 0x75, 0x7b, 0x61, 0xb0, 0xbd, 0x2f, 0xa2,
	0xba, 0x70, 0xa7, 0x73, 0x90, 0xd6, 0x30, 0x75, 0x14, 0x53, 0x13, 0xde, 0xdd, 0x60, 0xca, 0x71,
	0xd4, 0xab, 0x39, 0xf2, 0xf3, 0x9e, 0xd1, 0x5b, 0x4d, 0xf1, 0xd9, 0x06, 0x57, 0x36, 0x0f, 0x00,
	0x5b, 0xfb, 0xc2, 0xb6, 0x8f, 0xe7, 0x3c, 0x3c, 0x40, 0x69, 0xa0, 0x5a, 0x0a, 0xca, 0x85, 0xe7,
	0x7b, 0xa0, 0xca, 0xab, 0xf5, 0x9f, 0x2f, 0x56, 0xc8, 0x5e, 0xae, 0x90, 0xfd, 0x7b, 0x85, 0xec,
	0x2f, 0x6b, 0x64, 0x2d, 0xd7, 0xc8, 0xfa, 0xb9, 0x46, 0xd6, 0xf0, 0x41, 0x18, 0xc9, 0x71, 0xe6,
	0x7b, 0x01, 0x8b, 0xab, 0xeb, 0xfe, 0xa0, 0xcc, 0xe4, 0x3c, 0xa5, 0xc2, 0xaf, 0xa9, 0xef, 0xe4,
	0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x75, 0x9f, 0x3b, 0xfe, 0x03, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of QueryBreakFactor items.
	QueryBreakFactor(ctx context.Context, in *QueryBreakFactorRequest, opts ...grpc.CallOption) (*QueryBreakFactorResponse, error)
	// Queries watchlist.
	QueryWatchlist(ctx context.Context, in *QueryWatchlistRequest, opts ...grpc.CallOption) (*QueryWatchlistResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/terra.abs.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryBreakFactor(ctx context.Context, in *QueryBreakFactorRequest, opts ...grpc.CallOption) (*QueryBreakFactorResponse, error) {
	out := new(QueryBreakFactorResponse)
	err := c.cc.Invoke(ctx, "/terra.abs.v1beta1.Query/QueryBreakFactor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryWatchlist(ctx context.Context, in *QueryWatchlistRequest, opts ...grpc.CallOption) (*QueryWatchlistResponse, error) {
	out := new(QueryWatchlistResponse)
	err := c.cc.Invoke(ctx, "/terra.abs.v1beta1.Query/QueryWatchlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of QueryBreakFactor items.
	QueryBreakFactor(context.Context, *QueryBreakFactorRequest) (*QueryBreakFactorResponse, error)
	// Queries watchlist.
	QueryWatchlist(context.Context, *QueryWatchlistRequest) (*QueryWatchlistResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) QueryBreakFactor(ctx context.Context, req *QueryBreakFactorRequest) (*QueryBreakFactorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBreakFactor not implemented")
}
func (*UnimplementedQueryServer) QueryWatchlist(ctx context.Context, req *QueryWatchlistRequest) (*QueryWatchlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWatchlist not implemented")
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
		FullMethod: "/terra.abs.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryBreakFactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBreakFactorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryBreakFactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.abs.v1beta1.Query/QueryBreakFactor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryBreakFactor(ctx, req.(*QueryBreakFactorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.abs.v1beta1.Query/QueryWatchlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryWatchlist(ctx, req.(*QueryWatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.abs.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryBreakFactor",
			Handler:    _Query_QueryBreakFactor_Handler,
		},
		{
			MethodName: "QueryWatchlist",
			Handler:    _Query_QueryWatchlist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/abs/v1beta1/query.proto",
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

func (m *QueryBreakFactorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBreakFactorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBreakFactorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBreakFactorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBreakFactorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBreakFactorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWatchlistRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWatchlistRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWatchlistRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryWatchlistResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWatchlistResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWatchlistResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Watchlist.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryBreakFactorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBreakFactorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWatchlistRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryWatchlistResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Watchlist.Size()
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
func (m *QueryBreakFactorRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBreakFactorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBreakFactorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryBreakFactorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBreakFactorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBreakFactorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
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
func (m *QueryWatchlistRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWatchlistRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWatchlistRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryWatchlistResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWatchlistResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWatchlistResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watchlist", wireType)
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
			if err := m.Watchlist.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
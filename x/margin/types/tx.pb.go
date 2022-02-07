// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/margin/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgOpenLong struct {
	Signer           string                                  `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	CollateralAsset  string                                  `protobuf:"bytes,2,opt,name=collateral_asset,json=collateralAsset,proto3" json:"collateral_asset,omitempty"`
	CollateralAmount github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=collateral_amount,json=collateralAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"collateral_amount"`
	BorrowAsset      string                                  `protobuf:"bytes,4,opt,name=borrow_asset,json=borrowAsset,proto3" json:"borrow_asset,omitempty"`
}

func (m *MsgOpenLong) Reset()         { *m = MsgOpenLong{} }
func (m *MsgOpenLong) String() string { return proto.CompactTextString(m) }
func (*MsgOpenLong) ProtoMessage()    {}
func (*MsgOpenLong) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd3bc05d7e781ea, []int{0}
}
func (m *MsgOpenLong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenLong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenLong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenLong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenLong.Merge(m, src)
}
func (m *MsgOpenLong) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenLong) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenLong.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenLong proto.InternalMessageInfo

func (m *MsgOpenLong) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgOpenLong) GetCollateralAsset() string {
	if m != nil {
		return m.CollateralAsset
	}
	return ""
}

func (m *MsgOpenLong) GetBorrowAsset() string {
	if m != nil {
		return m.BorrowAsset
	}
	return ""
}

type MsgOpenLongResponse struct {
}

func (m *MsgOpenLongResponse) Reset()         { *m = MsgOpenLongResponse{} }
func (m *MsgOpenLongResponse) String() string { return proto.CompactTextString(m) }
func (*MsgOpenLongResponse) ProtoMessage()    {}
func (*MsgOpenLongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd3bc05d7e781ea, []int{1}
}
func (m *MsgOpenLongResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgOpenLongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgOpenLongResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgOpenLongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgOpenLongResponse.Merge(m, src)
}
func (m *MsgOpenLongResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgOpenLongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgOpenLongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgOpenLongResponse proto.InternalMessageInfo

type MsgCloseLong struct {
	Signer          string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	CollateralAsset string `protobuf:"bytes,2,opt,name=collateral_asset,json=collateralAsset,proto3" json:"collateral_asset,omitempty"`
	BorrowAsset     string `protobuf:"bytes,3,opt,name=borrow_asset,json=borrowAsset,proto3" json:"borrow_asset,omitempty"`
}

func (m *MsgCloseLong) Reset()         { *m = MsgCloseLong{} }
func (m *MsgCloseLong) String() string { return proto.CompactTextString(m) }
func (*MsgCloseLong) ProtoMessage()    {}
func (*MsgCloseLong) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd3bc05d7e781ea, []int{2}
}
func (m *MsgCloseLong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseLong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseLong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseLong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseLong.Merge(m, src)
}
func (m *MsgCloseLong) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseLong) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseLong.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseLong proto.InternalMessageInfo

func (m *MsgCloseLong) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgCloseLong) GetCollateralAsset() string {
	if m != nil {
		return m.CollateralAsset
	}
	return ""
}

func (m *MsgCloseLong) GetBorrowAsset() string {
	if m != nil {
		return m.BorrowAsset
	}
	return ""
}

type MsgCloseLongResponse struct {
}

func (m *MsgCloseLongResponse) Reset()         { *m = MsgCloseLongResponse{} }
func (m *MsgCloseLongResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCloseLongResponse) ProtoMessage()    {}
func (*MsgCloseLongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd3bc05d7e781ea, []int{3}
}
func (m *MsgCloseLongResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseLongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseLongResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseLongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseLongResponse.Merge(m, src)
}
func (m *MsgCloseLongResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseLongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseLongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseLongResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgOpenLong)(nil), "sifnode.margin.v1.MsgOpenLong")
	proto.RegisterType((*MsgOpenLongResponse)(nil), "sifnode.margin.v1.MsgOpenLongResponse")
	proto.RegisterType((*MsgCloseLong)(nil), "sifnode.margin.v1.MsgCloseLong")
	proto.RegisterType((*MsgCloseLongResponse)(nil), "sifnode.margin.v1.MsgCloseLongResponse")
}

func init() { proto.RegisterFile("sifnode/margin/v1/tx.proto", fileDescriptor_4dd3bc05d7e781ea) }

var fileDescriptor_4dd3bc05d7e781ea = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x51, 0x4b, 0xc2, 0x50,
	0x14, 0xc7, 0xb7, 0x0c, 0xc9, 0xab, 0x50, 0x2e, 0x13, 0x19, 0x34, 0xcb, 0x87, 0xac, 0xa0, 0x0d,
	0xeb, 0x13, 0x68, 0x3d, 0x26, 0x81, 0x15, 0x41, 0x04, 0x31, 0xe7, 0xf5, 0x3a, 0xda, 0xee, 0x19,
	0x3b, 0x57, 0xb3, 0x6f, 0xd1, 0x87, 0xe9, 0x43, 0xf8, 0x68, 0x6f, 0xd1, 0x83, 0x84, 0x7e, 0x91,
	0x70, 0x9b, 0x3a, 0x52, 0xea, 0xa5, 0xa7, 0x6d, 0xe7, 0xff, 0xdf, 0xb9, 0xbf, 0xff, 0x3d, 0x87,
	0xa8, 0x68, 0xb7, 0x39, 0xb4, 0xa8, 0xe1, 0x9a, 0x3e, 0xb3, 0xb9, 0xd1, 0xab, 0x18, 0xa2, 0xaf,
	0x7b, 0x3e, 0x08, 0x50, 0xb2, 0x91, 0xa6, 0x87, 0x9a, 0xde, 0xab, 0xa8, 0x39, 0x06, 0x0c, 0x02,
	0xd5, 0x98, 0xbe, 0x85, 0x46, 0x75, 0x77, 0x45, 0x93, 0x17, 0x8f, 0x62, 0x28, 0x97, 0xde, 0x65,
	0x92, 0xae, 0x23, 0xbb, 0xf2, 0x28, 0xbf, 0x04, 0xce, 0x94, 0x3c, 0x49, 0xa2, 0xcd, 0x38, 0xf5,
	0x0b, 0xf2, 0x9e, 0x7c, 0x98, 0x6a, 0x44, 0x5f, 0xca, 0x11, 0xd9, 0xb2, 0xc0, 0x71, 0x4c, 0x41,
	0x7d, 0xd3, 0x79, 0x34, 0x11, 0xa9, 0x28, 0xac, 0x05, 0x8e, 0xcd, 0x45, 0xbd, 0x3a, 0x2d, 0x2b,
	0x0f, 0x24, 0x1b, 0xb7, 0xba, 0xd0, 0xe5, 0xa2, 0x90, 0x98, 0x7a, 0x6b, 0xc6, 0x60, 0x54, 0x94,
	0x3e, 0x47, 0xc5, 0x32, 0xb3, 0x45, 0xa7, 0xdb, 0xd4, 0x2d, 0x70, 0x0d, 0x0b, 0xd0, 0x05, 0x8c,
	0x1e, 0x27, 0xd8, 0x7a, 0x8a, 0xf8, 0x6e, 0x6d, 0x2e, 0x1a, 0xb1, 0x43, 0xab, 0x41, 0x23, 0x65,
	0x9f, 0x64, 0x9a, 0xe0, 0xfb, 0xf0, 0x1c, 0x41, 0xac, 0x07, 0x10, 0xe9, 0xb0, 0x16, 0x00, 0x94,
	0x76, 0xc8, 0x76, 0x2c, 0x52, 0x83, 0xa2, 0x07, 0x1c, 0x69, 0x49, 0x90, 0x4c, 0x1d, 0xd9, 0xb9,
	0x03, 0x48, 0xff, 0x2b, 0xea, 0x4f, 0x98, 0xc4, 0x32, 0x4c, 0x9e, 0xe4, 0xe2, 0xa7, 0xce, 0x68,
	0x4e, 0xdf, 0x64, 0x92, 0xa8, 0x23, 0x53, 0x6e, 0xc8, 0xc6, 0xfc, 0xf2, 0x35, 0x7d, 0x69, 0xaa,
	0x7a, 0x2c, 0x89, 0x7a, 0xf0, 0xbb, 0x3e, 0x4f, 0x2a, 0x29, 0x77, 0x24, 0xb5, 0x08, 0x5a, 0x5c,
	0xfd, 0xdb, 0xdc, 0xa0, 0x96, 0xff, 0x30, 0x2c, 0x1a, 0xd7, 0x2e, 0x06, 0x63, 0x4d, 0x1e, 0x8e,
	0x35, 0xf9, 0x6b, 0xac, 0xc9, 0xaf, 0x13, 0x4d, 0x1a, 0x4e, 0x34, 0xe9, 0x63, 0xa2, 0x49, 0xf7,
	0xc7, 0xb1, 0x99, 0x5e, 0xdb, 0x6d, 0xab, 0x63, 0xda, 0xdc, 0x98, 0x2d, 0x5f, 0x7f, 0xb6, 0x7e,
	0xc1, 0x6c, 0x9b, 0xc9, 0x60, 0xf9, 0xce, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x70, 0x94, 0xfc,
	0x49, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	OpenLong(ctx context.Context, in *MsgOpenLong, opts ...grpc.CallOption) (*MsgOpenLongResponse, error)
	CloseLong(ctx context.Context, in *MsgCloseLong, opts ...grpc.CallOption) (*MsgCloseLongResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) OpenLong(ctx context.Context, in *MsgOpenLong, opts ...grpc.CallOption) (*MsgOpenLongResponse, error) {
	out := new(MsgOpenLongResponse)
	err := c.cc.Invoke(ctx, "/sifnode.margin.v1.Msg/OpenLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseLong(ctx context.Context, in *MsgCloseLong, opts ...grpc.CallOption) (*MsgCloseLongResponse, error) {
	out := new(MsgCloseLongResponse)
	err := c.cc.Invoke(ctx, "/sifnode.margin.v1.Msg/CloseLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	OpenLong(context.Context, *MsgOpenLong) (*MsgOpenLongResponse, error)
	CloseLong(context.Context, *MsgCloseLong) (*MsgCloseLongResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) OpenLong(ctx context.Context, req *MsgOpenLong) (*MsgOpenLongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenLong not implemented")
}
func (*UnimplementedMsgServer) CloseLong(ctx context.Context, req *MsgCloseLong) (*MsgCloseLongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLong not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_OpenLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgOpenLong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).OpenLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.margin.v1.Msg/OpenLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).OpenLong(ctx, req.(*MsgOpenLong))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseLong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sifnode.margin.v1.Msg/CloseLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseLong(ctx, req.(*MsgCloseLong))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sifnode.margin.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenLong",
			Handler:    _Msg_OpenLong_Handler,
		},
		{
			MethodName: "CloseLong",
			Handler:    _Msg_CloseLong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sifnode/margin/v1/tx.proto",
}

func (m *MsgOpenLong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenLong) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenLong) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BorrowAsset) > 0 {
		i -= len(m.BorrowAsset)
		copy(dAtA[i:], m.BorrowAsset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BorrowAsset)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.CollateralAmount.Size()
		i -= size
		if _, err := m.CollateralAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.CollateralAsset) > 0 {
		i -= len(m.CollateralAsset)
		copy(dAtA[i:], m.CollateralAsset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CollateralAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgOpenLongResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgOpenLongResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgOpenLongResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCloseLong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseLong) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseLong) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BorrowAsset) > 0 {
		i -= len(m.BorrowAsset)
		copy(dAtA[i:], m.BorrowAsset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BorrowAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CollateralAsset) > 0 {
		i -= len(m.CollateralAsset)
		copy(dAtA[i:], m.CollateralAsset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CollateralAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCloseLongResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseLongResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseLongResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgOpenLong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CollateralAsset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.CollateralAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.BorrowAsset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgOpenLongResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCloseLong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CollateralAsset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BorrowAsset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCloseLongResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgOpenLong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgOpenLong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenLong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorrowAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgOpenLongResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgOpenLongResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgOpenLongResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCloseLong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCloseLong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseLong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BorrowAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCloseLongResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCloseLongResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseLongResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
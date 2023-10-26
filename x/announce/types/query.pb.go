// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/announce/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GetAnnouncedStorageLocationsRequest is the request type for the
// GetAnnouncedStorageLocations RPC method.
type GetAnnouncedStorageLocationsRequest struct {
	// list of validators where each validator is in hex-encoded eth address
	// format (20 bytes)
	Validator [][]byte `protobuf:"bytes,1,rep,name=validator,proto3" json:"validator,omitempty"`
}

func (m *GetAnnouncedStorageLocationsRequest) Reset()         { *m = GetAnnouncedStorageLocationsRequest{} }
func (m *GetAnnouncedStorageLocationsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAnnouncedStorageLocationsRequest) ProtoMessage()    {}
func (*GetAnnouncedStorageLocationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06dd56980c12f1, []int{0}
}

func (m *GetAnnouncedStorageLocationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GetAnnouncedStorageLocationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAnnouncedStorageLocationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GetAnnouncedStorageLocationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnnouncedStorageLocationsRequest.Merge(m, src)
}

func (m *GetAnnouncedStorageLocationsRequest) XXX_Size() int {
	return m.Size()
}

func (m *GetAnnouncedStorageLocationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnnouncedStorageLocationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnnouncedStorageLocationsRequest proto.InternalMessageInfo

func (m *GetAnnouncedStorageLocationsRequest) GetValidator() [][]byte {
	if m != nil {
		return m.Validator
	}
	return nil
}

// GetAnnouncedStorageLocationsResponse is the response type for the
// GetAnnouncedStorageLocations RPC method.
type GetAnnouncedStorageLocationsResponse struct {
	Metadata []*StorageMetadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *GetAnnouncedStorageLocationsResponse) Reset()         { *m = GetAnnouncedStorageLocationsResponse{} }
func (m *GetAnnouncedStorageLocationsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAnnouncedStorageLocationsResponse) ProtoMessage()    {}
func (*GetAnnouncedStorageLocationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06dd56980c12f1, []int{1}
}

func (m *GetAnnouncedStorageLocationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GetAnnouncedStorageLocationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAnnouncedStorageLocationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GetAnnouncedStorageLocationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnnouncedStorageLocationsResponse.Merge(m, src)
}

func (m *GetAnnouncedStorageLocationsResponse) XXX_Size() int {
	return m.Size()
}

func (m *GetAnnouncedStorageLocationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnnouncedStorageLocationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnnouncedStorageLocationsResponse proto.InternalMessageInfo

func (m *GetAnnouncedStorageLocationsResponse) GetMetadata() []*StorageMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// GetAnnouncedValidatorsRequest is the request type for the
// GetAnnouncedValidators RPC method.
type GetAnnouncedValidatorsRequest struct{}

func (m *GetAnnouncedValidatorsRequest) Reset()         { *m = GetAnnouncedValidatorsRequest{} }
func (m *GetAnnouncedValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAnnouncedValidatorsRequest) ProtoMessage()    {}
func (*GetAnnouncedValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06dd56980c12f1, []int{2}
}

func (m *GetAnnouncedValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GetAnnouncedValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAnnouncedValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GetAnnouncedValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnnouncedValidatorsRequest.Merge(m, src)
}

func (m *GetAnnouncedValidatorsRequest) XXX_Size() int {
	return m.Size()
}

func (m *GetAnnouncedValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnnouncedValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnnouncedValidatorsRequest proto.InternalMessageInfo

// GetAnnouncedValidatorsResponse is the response type for the
// GetAnnouncedValidators RPC method.
type GetAnnouncedValidatorsResponse struct {
	// list of validators where each validator is in hex-encoded eth address
	// format (20 bytes)
	Validator []string `protobuf:"bytes,1,rep,name=validator,proto3" json:"validator,omitempty"`
}

func (m *GetAnnouncedValidatorsResponse) Reset()         { *m = GetAnnouncedValidatorsResponse{} }
func (m *GetAnnouncedValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAnnouncedValidatorsResponse) ProtoMessage()    {}
func (*GetAnnouncedValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06dd56980c12f1, []int{3}
}

func (m *GetAnnouncedValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GetAnnouncedValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAnnouncedValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GetAnnouncedValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnnouncedValidatorsResponse.Merge(m, src)
}

func (m *GetAnnouncedValidatorsResponse) XXX_Size() int {
	return m.Size()
}

func (m *GetAnnouncedValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnnouncedValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnnouncedValidatorsResponse proto.InternalMessageInfo

func (m *GetAnnouncedValidatorsResponse) GetValidator() []string {
	if m != nil {
		return m.Validator
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAnnouncedStorageLocationsRequest)(nil), "hyperlane.announce.v1.GetAnnouncedStorageLocationsRequest")
	proto.RegisterType((*GetAnnouncedStorageLocationsResponse)(nil), "hyperlane.announce.v1.GetAnnouncedStorageLocationsResponse")
	proto.RegisterType((*GetAnnouncedValidatorsRequest)(nil), "hyperlane.announce.v1.GetAnnouncedValidatorsRequest")
	proto.RegisterType((*GetAnnouncedValidatorsResponse)(nil), "hyperlane.announce.v1.GetAnnouncedValidatorsResponse")
}

func init() { proto.RegisterFile("hyperlane/announce/v1/query.proto", fileDescriptor_3e06dd56980c12f1) }

var fileDescriptor_3e06dd56980c12f1 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0xa8, 0x2c, 0x48,
	0x2d, 0xca, 0x49, 0xcc, 0x4b, 0xd5, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0xcd, 0x4b, 0x4e, 0xd5, 0x2f,
	0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85,
	0x2b, 0xd1, 0x83, 0x29, 0xd1, 0x2b, 0x33, 0x94, 0xc2, 0xa1, 0xb3, 0xa4, 0xb2, 0x20, 0xb5, 0x18,
	0xa2, 0x53, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x13, 0xac, 0xa6,
	0x24, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0x2a, 0xab, 0xe4, 0xcc, 0xa5, 0xec, 0x9e, 0x5a, 0xe2, 0x08,
	0xd5, 0x9b, 0x12, 0x5c, 0x92, 0x5f, 0x94, 0x98, 0x9e, 0xea, 0x93, 0x9f, 0x0c, 0x51, 0x15, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc3, 0xc5, 0x59, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58,
	0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x13, 0x84, 0x10, 0x50, 0xca, 0xe2, 0x52, 0xc1,
	0x6f, 0x48, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x13, 0x17, 0x47, 0x6e, 0x6a, 0x49, 0x62,
	0x4a, 0x62, 0x49, 0x22, 0xd8, 0x10, 0x6e, 0x23, 0x35, 0x3d, 0xac, 0xfe, 0xd2, 0x83, 0x1a, 0xe1,
	0x0b, 0x55, 0x1d, 0x04, 0xd7, 0xa7, 0x24, 0xcf, 0x25, 0x8b, 0x6c, 0x57, 0x18, 0xcc, 0x11, 0x30,
	0xa7, 0x2a, 0xd9, 0x71, 0xc9, 0xe1, 0x52, 0x00, 0x75, 0x06, 0x86, 0x67, 0x38, 0x91, 0x3c, 0x63,
	0xb4, 0x88, 0x99, 0x8b, 0x35, 0x10, 0x14, 0xf2, 0x42, 0xd7, 0x19, 0xb9, 0x64, 0xf0, 0xf9, 0x4b,
	0xc8, 0x0a, 0x87, 0xeb, 0x89, 0x08, 0x51, 0x29, 0x6b, 0xb2, 0xf4, 0x42, 0x7c, 0xa0, 0x64, 0xd7,
	0x74, 0xf9, 0xc9, 0x64, 0x26, 0x0b, 0x21, 0x33, 0x7d, 0xec, 0xf1, 0x9f, 0x9e, 0x5a, 0x12, 0x0f,
	0xe3, 0xa7, 0xc4, 0x17, 0x43, 0x8c, 0x89, 0xcf, 0x81, 0x3b, 0x7c, 0x1f, 0x23, 0x97, 0x18, 0xf6,
	0x40, 0x12, 0x32, 0x21, 0xc2, 0x5d, 0x18, 0x81, 0x2e, 0x65, 0x4a, 0xa2, 0x2e, 0xa8, 0x3f, 0xcc,
	0xc1, 0xfe, 0x30, 0x14, 0xd2, 0x27, 0xca, 0x1f, 0xf0, 0x38, 0x2a, 0x76, 0x8a, 0x3e, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xc7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x92, 0xa2, 0xc4, 0xbc, 0xf4, 0xd4, 0x9c, 0xfc, 0xb2, 0x54, 0xdd,
	0xb2, 0xd4, 0xbc, 0x92, 0xd2, 0xa2, 0xd4, 0x62, 0x84, 0x4d, 0xba, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9,
	0xc5, 0xfa, 0x15, 0x08, 0x2b, 0xc1, 0xf9, 0x26, 0x89, 0x0d, 0x9c, 0x35, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0xf4, 0x00, 0xbc, 0x97, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Gets the announced storage locations (where signatures are stored) for the
	// requested validators
	GetAnnouncedStorageLocations(ctx context.Context, in *GetAnnouncedStorageLocationsRequest, opts ...grpc.CallOption) (*GetAnnouncedStorageLocationsResponse, error)
	// Gets a list of validators that have made announcements
	GetAnnouncedValidators(ctx context.Context, in *GetAnnouncedValidatorsRequest, opts ...grpc.CallOption) (*GetAnnouncedValidatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetAnnouncedStorageLocations(ctx context.Context, in *GetAnnouncedStorageLocationsRequest, opts ...grpc.CallOption) (*GetAnnouncedStorageLocationsResponse, error) {
	out := new(GetAnnouncedStorageLocationsResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.announce.v1.Query/GetAnnouncedStorageLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAnnouncedValidators(ctx context.Context, in *GetAnnouncedValidatorsRequest, opts ...grpc.CallOption) (*GetAnnouncedValidatorsResponse, error) {
	out := new(GetAnnouncedValidatorsResponse)
	err := c.cc.Invoke(ctx, "/hyperlane.announce.v1.Query/GetAnnouncedValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Gets the announced storage locations (where signatures are stored) for the
	// requested validators
	GetAnnouncedStorageLocations(context.Context, *GetAnnouncedStorageLocationsRequest) (*GetAnnouncedStorageLocationsResponse, error)
	// Gets a list of validators that have made announcements
	GetAnnouncedValidators(context.Context, *GetAnnouncedValidatorsRequest) (*GetAnnouncedValidatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct{}

func (*UnimplementedQueryServer) GetAnnouncedStorageLocations(ctx context.Context, req *GetAnnouncedStorageLocationsRequest) (*GetAnnouncedStorageLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncedStorageLocations not implemented")
}

func (*UnimplementedQueryServer) GetAnnouncedValidators(ctx context.Context, req *GetAnnouncedValidatorsRequest) (*GetAnnouncedValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnnouncedValidators not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetAnnouncedStorageLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncedStorageLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAnnouncedStorageLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.announce.v1.Query/GetAnnouncedStorageLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAnnouncedStorageLocations(ctx, req.(*GetAnnouncedStorageLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAnnouncedValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnouncedValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAnnouncedValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hyperlane.announce.v1.Query/GetAnnouncedValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAnnouncedValidators(ctx, req.(*GetAnnouncedValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hyperlane.announce.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnnouncedStorageLocations",
			Handler:    _Query_GetAnnouncedStorageLocations_Handler,
		},
		{
			MethodName: "GetAnnouncedValidators",
			Handler:    _Query_GetAnnouncedValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyperlane/announce/v1/query.proto",
}

func (m *GetAnnouncedStorageLocationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAnnouncedStorageLocationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAnnouncedStorageLocationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		for iNdEx := len(m.Validator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validator[iNdEx])
			copy(dAtA[i:], m.Validator[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Validator[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetAnnouncedStorageLocationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAnnouncedStorageLocationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAnnouncedStorageLocationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for iNdEx := len(m.Metadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetAnnouncedValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAnnouncedValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAnnouncedValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetAnnouncedValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAnnouncedValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAnnouncedValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		for iNdEx := len(m.Validator) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validator[iNdEx])
			copy(dAtA[i:], m.Validator[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Validator[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
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

func (m *GetAnnouncedStorageLocationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validator) > 0 {
		for _, b := range m.Validator {
			l = len(b)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetAnnouncedStorageLocationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for _, e := range m.Metadata {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetAnnouncedValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetAnnouncedValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validator) > 0 {
		for _, s := range m.Validator {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *GetAnnouncedStorageLocationsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAnnouncedStorageLocationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAnnouncedStorageLocationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator, make([]byte, postIndex-iNdEx))
			copy(m.Validator[len(m.Validator)-1], dAtA[iNdEx:postIndex])
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

func (m *GetAnnouncedStorageLocationsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAnnouncedStorageLocationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAnnouncedStorageLocationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			m.Metadata = append(m.Metadata, &StorageMetadata{})
			if err := m.Metadata[len(m.Metadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func (m *GetAnnouncedValidatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAnnouncedValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAnnouncedValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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

func (m *GetAnnouncedValidatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAnnouncedValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAnnouncedValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
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
			m.Validator = append(m.Validator, string(dAtA[iNdEx:postIndex]))
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
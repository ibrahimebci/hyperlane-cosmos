// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hyperlane/igp/v1/types.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Hyperlane's gas oracle to configure exchange rates between origin and
// destination
type GasOracleConfig struct {
	// The IGP that this gas oracle config belongs to
	IgpId uint32 `protobuf:"varint,1,opt,name=igp_id,json=igpId,proto3" json:"igp_id,omitempty"`
	// The address that can update gas oracle configs for the remote domain
	GasOracle string `protobuf:"bytes,2,opt,name=gas_oracle,json=gasOracle,proto3" json:"gas_oracle,omitempty"`
	// The domain that the gas oracle can update gas related information for
	RemoteDomain uint32 `protobuf:"varint,3,opt,name=remote_domain,json=remoteDomain,proto3" json:"remote_domain,omitempty"`
}

func (m *GasOracleConfig) Reset()         { *m = GasOracleConfig{} }
func (m *GasOracleConfig) String() string { return proto.CompactTextString(m) }
func (*GasOracleConfig) ProtoMessage()    {}
func (*GasOracleConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d4b4953a4c266e9, []int{0}
}

func (m *GasOracleConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GasOracleConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasOracleConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GasOracleConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasOracleConfig.Merge(m, src)
}

func (m *GasOracleConfig) XXX_Size() int {
	return m.Size()
}

func (m *GasOracleConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GasOracleConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GasOracleConfig proto.InternalMessageInfo

// Hyperlane's gas oracle to configure exchange rates between origin and
// destination
type GasOracle struct {
	// Address of the oracle that can update this config
	GasOracle string `protobuf:"bytes,1,opt,name=gas_oracle,json=gasOracle,proto3" json:"gas_oracle,omitempty"`
	// The domain of the message's destination chain
	RemoteDomain      uint32                                 `protobuf:"varint,2,opt,name=remote_domain,json=remoteDomain,proto3" json:"remote_domain,omitempty"`
	TokenExchangeRate github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=token_exchange_rate,json=tokenExchangeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_exchange_rate"`
	GasPrice          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=gas_price,json=gasPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_price"`
	// gas overhead for the remote domain
	GasOverhead github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=gas_overhead,json=gasOverhead,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_overhead"`
}

func (m *GasOracle) Reset()         { *m = GasOracle{} }
func (m *GasOracle) String() string { return proto.CompactTextString(m) }
func (*GasOracle) ProtoMessage()    {}
func (*GasOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d4b4953a4c266e9, []int{1}
}

func (m *GasOracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GasOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasOracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GasOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasOracle.Merge(m, src)
}

func (m *GasOracle) XXX_Size() int {
	return m.Size()
}

func (m *GasOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_GasOracle.DiscardUnknown(m)
}

var xxx_messageInfo_GasOracle proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GasOracleConfig)(nil), "hyperlane.igp.v1.GasOracleConfig")
	proto.RegisterType((*GasOracle)(nil), "hyperlane.igp.v1.GasOracle")
}

func init() { proto.RegisterFile("hyperlane/igp/v1/types.proto", fileDescriptor_3d4b4953a4c266e9) }

var fileDescriptor_3d4b4953a4c266e9 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x33, 0xfd, 0x87, 0x19, 0x5b, 0xb4, 0xb1, 0x42, 0x2c, 0x92, 0x2d, 0x15, 0xa4, 0x08,
	0x49, 0x28, 0x1e, 0x04, 0xf1, 0xd2, 0x55, 0x91, 0x3d, 0x29, 0x11, 0x3c, 0x78, 0x30, 0x4c, 0x33,
	0xaf, 0x93, 0xa1, 0x9b, 0x99, 0x71, 0x66, 0x1a, 0xda, 0x6f, 0xe0, 0xc1, 0x83, 0x1f, 0xc1, 0x63,
	0x8f, 0x1e, 0xfa, 0x21, 0x7a, 0x2c, 0x3d, 0x89, 0x87, 0x45, 0x76, 0x0f, 0x7e, 0x0d, 0xc9, 0x64,
	0x5c, 0x44, 0x8f, 0xee, 0x25, 0x7f, 0x9e, 0x79, 0x79, 0x7e, 0xcf, 0xbc, 0xef, 0x8b, 0xef, 0xd6,
	0xa7, 0x0a, 0xf4, 0x98, 0x08, 0xc8, 0x39, 0x53, 0x79, 0xbb, 0x9f, 0xdb, 0x53, 0x05, 0x26, 0x53,
	0x5a, 0x5a, 0x19, 0xdd, 0x9c, 0x9f, 0x66, 0x9c, 0xa9, 0xac, 0xdd, 0xdf, 0xbe, 0x53, 0x49, 0xd3,
	0x48, 0x53, 0xba, 0xf3, 0xbc, 0xff, 0xe9, 0x8b, 0xb7, 0xb7, 0x98, 0x64, 0xb2, 0xd7, 0xbb, 0x2f,
	0xaf, 0x6e, 0x92, 0x86, 0x0b, 0x99, 0xbb, 0x67, 0x2f, 0xed, 0x7e, 0x42, 0xf8, 0xc6, 0x0b, 0x62,
	0x5e, 0x6a, 0x52, 0x8d, 0xe1, 0xa9, 0x14, 0xef, 0x39, 0x8b, 0x6e, 0xe3, 0x35, 0xce, 0x54, 0xc9,
	0x69, 0x8c, 0x76, 0xd0, 0xde, 0x46, 0xb1, 0xca, 0x99, 0x1a, 0xd1, 0xe8, 0x11, 0xc6, 0x8c, 0x98,
	0x52, 0xba, 0xd2, 0x78, 0x69, 0x07, 0xed, 0x85, 0xc3, 0xf8, 0xea, 0x3c, 0xdd, 0xf2, 0xe4, 0x03,
	0x4a, 0x35, 0x18, 0xf3, 0xda, 0x6a, 0x2e, 0x58, 0x11, 0xb2, 0xdf, 0xae, 0xd1, 0x3d, 0xbc, 0xa1,
	0xa1, 0x91, 0x16, 0x4a, 0x2a, 0x1b, 0xc2, 0x45, 0xbc, 0xec, 0x6c, 0xd7, 0x7b, 0xf1, 0x99, 0xd3,
	0x1e, 0xaf, 0x7c, 0xfc, 0x32, 0x08, 0x76, 0xcf, 0x96, 0x71, 0x38, 0x8f, 0xf3, 0x17, 0x11, 0xfd,
	0x07, 0x71, 0xe9, 0x5f, 0x62, 0xf4, 0x01, 0xdf, 0xb2, 0xf2, 0x08, 0x44, 0x09, 0x27, 0x55, 0x4d,
	0x04, 0x83, 0x52, 0x13, 0x0b, 0x2e, 0x5c, 0x38, 0x3c, 0xb8, 0x98, 0x0c, 0x82, 0xef, 0x93, 0xc1,
	0x7d, 0xc6, 0x6d, 0x7d, 0x7c, 0x98, 0x55, 0xb2, 0xf1, 0x1d, 0xf6, 0xaf, 0xd4, 0xd0, 0x23, 0x3f,
	0x9f, 0x91, 0xb0, 0x57, 0xe7, 0x29, 0xf6, 0xa1, 0x46, 0xc2, 0x9e, 0xfd, 0xfc, 0xfa, 0x00, 0x15,
	0x9b, 0xce, 0xfd, 0xb9, 0x37, 0x2f, 0x88, 0x85, 0xe8, 0x1d, 0xee, 0x42, 0x96, 0x4a, 0xf3, 0x0a,
	0xe2, 0x95, 0x45, 0x81, 0xae, 0x31, 0x62, 0x5e, 0x75, 0x96, 0x11, 0xc5, 0xeb, 0xae, 0x61, 0x2d,
	0xe8, 0x1a, 0x08, 0x8d, 0x57, 0x17, 0x85, 0xb8, 0xde, 0xf5, 0xd6, 0xbb, 0xf6, 0xa3, 0x1a, 0xbe,
	0xb9, 0x98, 0x26, 0xe8, 0x72, 0x9a, 0xa0, 0x1f, 0xd3, 0x04, 0x7d, 0x9e, 0x25, 0xc1, 0xe5, 0x2c,
	0x09, 0xbe, 0xcd, 0x92, 0xe0, 0xed, 0x93, 0x3f, 0x38, 0xc6, 0xea, 0xee, 0xf6, 0x63, 0xd9, 0x42,
	0xda, 0x82, 0xb0, 0xc7, 0x1a, 0x4c, 0x3e, 0xdf, 0xe4, 0xd4, 0xc7, 0x38, 0x71, 0x0b, 0xef, 0x12,
	0x1c, 0xae, 0xb9, 0xc5, 0x7c, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xc0, 0xd4, 0xce, 0x0e,
	0x03, 0x00, 0x00,
}

func (m *GasOracleConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasOracleConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasOracleConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RemoteDomain != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RemoteDomain))
		i--
		dAtA[i] = 0x18
	}
	if len(m.GasOracle) > 0 {
		i -= len(m.GasOracle)
		copy(dAtA[i:], m.GasOracle)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.GasOracle)))
		i--
		dAtA[i] = 0x12
	}
	if m.IgpId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.IgpId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GasOracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasOracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasOracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.GasOverhead.Size()
		i -= size
		if _, err := m.GasOverhead.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.GasPrice.Size()
		i -= size
		if _, err := m.GasPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TokenExchangeRate.Size()
		i -= size
		if _, err := m.TokenExchangeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.RemoteDomain != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RemoteDomain))
		i--
		dAtA[i] = 0x10
	}
	if len(m.GasOracle) > 0 {
		i -= len(m.GasOracle)
		copy(dAtA[i:], m.GasOracle)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.GasOracle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *GasOracleConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IgpId != 0 {
		n += 1 + sovTypes(uint64(m.IgpId))
	}
	l = len(m.GasOracle)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.RemoteDomain != 0 {
		n += 1 + sovTypes(uint64(m.RemoteDomain))
	}
	return n
}

func (m *GasOracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GasOracle)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.RemoteDomain != 0 {
		n += 1 + sovTypes(uint64(m.RemoteDomain))
	}
	l = m.TokenExchangeRate.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.GasPrice.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.GasOverhead.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *GasOracleConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GasOracleConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasOracleConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgpId", wireType)
			}
			m.IgpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IgpId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasOracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasOracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDomain", wireType)
			}
			m.RemoteDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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

func (m *GasOracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GasOracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasOracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasOracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasOracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDomain", wireType)
			}
			m.RemoteDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenExchangeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasOverhead", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GasOverhead.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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

func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)

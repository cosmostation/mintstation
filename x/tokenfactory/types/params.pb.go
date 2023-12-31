// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the tokenfactory module.
type Params struct {
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
	// if denom_creation_fee is an empty array, then this field is used to add more gas consumption
	// to the base cost.
	// https://github.com/CosmWasm/token-factory/issues/11
	DenomCreationGasConsume uint64 `protobuf:"varint,2,opt,name=denom_creation_gas_consume,json=denomCreationGasConsume,proto3" json:"denom_creation_gas_consume,omitempty" yaml:"denom_creation_gas_consume"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8299d306f3ff47, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func (m *Params) GetDenomCreationGasConsume() uint64 {
	if m != nil {
		return m.DenomCreationGasConsume
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.tokenfactory.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/params.proto", fileDescriptor_cc8299d306f3ff47)
}

var fileDescriptor_cc8299d306f3ff47 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xc2, 0x40,
	0x18, 0xc7, 0x5b, 0x34, 0x0c, 0x75, 0x31, 0x8d, 0x89, 0x40, 0x4c, 0x8b, 0x9d, 0x60, 0xb0, 0x17,
	0xd4, 0xc1, 0x38, 0xd2, 0x44, 0x27, 0x12, 0xc3, 0xe8, 0xd2, 0x7c, 0x2d, 0x47, 0x7b, 0xc1, 0xf6,
	0x23, 0xbd, 0xab, 0xb1, 0x8f, 0xe0, 0xe6, 0xe4, 0x43, 0xf8, 0x24, 0x8c, 0x8c, 0x4e, 0xd5, 0xc0,
	0x1b, 0xf0, 0x04, 0x86, 0xeb, 0x61, 0x40, 0x8d, 0x53, 0xfb, 0xe5, 0xff, 0xff, 0xff, 0xee, 0xff,
	0xdd, 0x19, 0x5d, 0xe4, 0x09, 0x72, 0xc6, 0x89, 0xc0, 0x09, 0x4d, 0xc7, 0x10, 0x0a, 0xcc, 0x0a,
	0xf2, 0xd8, 0x0b, 0xa8, 0x80, 0x1e, 0x99, 0x42, 0x06, 0x09, 0x77, 0xa7, 0x19, 0x0a, 0x34, 0x4f,
	0x94, 0xd5, 0xdd, 0xb6, 0xba, 0xca, 0xda, 0x3a, 0x8a, 0x30, 0x42, 0x69, 0x24, 0xeb, 0xbf, 0x2a,
	0xd3, 0xba, 0xfc, 0x17, 0x0f, 0xb9, 0x88, 0x31, 0x63, 0xa2, 0x18, 0x50, 0x01, 0x23, 0x10, 0xa0,
	0x52, 0xcd, 0x50, 0xc6, 0xfc, 0x0a, 0x57, 0x0d, 0x4a, 0xb2, 0xaa, 0x89, 0x04, 0xc0, 0xe9, 0x37,
	0x27, 0x44, 0x96, 0x56, 0xba, 0xf3, 0x5c, 0x33, 0xea, 0x77, 0xb2, 0xb5, 0xf9, 0xaa, 0x1b, 0xe6,
	0x88, 0xa6, 0x98, 0xf8, 0x61, 0x46, 0x41, 0x30, 0x4c, 0xfd, 0x31, 0xa5, 0x0d, 0xbd, 0xbd, 0xd7,
	0x39, 0x38, 0x6f, 0xba, 0x0a, 0xbb, 0x06, 0x6d, 0x96, 0x70, 0x3d, 0x64, 0x69, 0x7f, 0x30, 0x2b,
	0x6d, 0x6d, 0x55, 0xda, 0xcd, 0x02, 0x92, 0x87, 0x6b, 0xe7, 0x37, 0xc2, 0x79, 0xfb, 0xb0, 0x3b,
	0x11, 0x13, 0x71, 0x1e, 0xb8, 0x21, 0x26, 0xaa, 0xa0, 0xfa, 0x9c, 0xf1, 0xd1, 0x84, 0x88, 0x62,
	0x4a, 0xb9, 0xa4, 0xf1, 0xe1, 0xa1, 0x04, 0x78, 0x2a, 0x7f, 0x43, 0xa9, 0x39, 0x36, 0x5a, 0x3f,
	0xa0, 0x11, 0x70, 0x3f, 0xc4, 0x94, 0xe7, 0x09, 0x6d, 0xd4, 0xda, 0x7a, 0x67, 0xbf, 0xdf, 0x9d,
	0x95, 0xb6, 0xbe, 0x2a, 0xed, 0xd3, 0x3f, 0x4b, 0x6c, 0xf9, 0x9d, 0xe1, 0xf1, 0xce, 0x01, 0xb7,
	0xc0, 0xbd, 0x4a, 0xe9, 0x0f, 0x67, 0x0b, 0x4b, 0x9f, 0x2f, 0x2c, 0xfd, 0x73, 0x61, 0xe9, 0x2f,
	0x4b, 0x4b, 0x9b, 0x2f, 0x2d, 0xed, 0x7d, 0x69, 0x69, 0xf7, 0x57, 0x5b, 0xed, 0xbd, 0x98, 0xc5,
	0x39, 0xc4, 0x39, 0x78, 0x31, 0xb0, 0x94, 0x84, 0x9b, 0x91, 0x3c, 0xed, 0x3e, 0x9a, 0xdc, 0x29,
	0xa8, 0xcb, 0x6b, 0xbe, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x37, 0x1a, 0xf5, 0xac, 0x38, 0x02,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomCreationGasConsume != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DenomCreationGasConsume))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DenomCreationGasConsume != 0 {
		n += 1 + sovParams(uint64(m.DenomCreationGasConsume))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationGasConsume", wireType)
			}
			m.DenomCreationGasConsume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomCreationGasConsume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

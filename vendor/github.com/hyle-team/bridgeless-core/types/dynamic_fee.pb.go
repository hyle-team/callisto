// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/dynamic_fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ExtensionOptionDynamicFeeTx is an extension option that specifies the maxPrioPrice for cosmos tx
type ExtensionOptionDynamicFeeTx struct {
	// max_priority_price is the same as `max_priority_fee_per_gas` in eip-1559 spec
	MaxPriorityPrice github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_priority_price,json=maxPriorityPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_priority_price"`
}

func (m *ExtensionOptionDynamicFeeTx) Reset()         { *m = ExtensionOptionDynamicFeeTx{} }
func (m *ExtensionOptionDynamicFeeTx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionDynamicFeeTx) ProtoMessage()    {}
func (*ExtensionOptionDynamicFeeTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7cf05c9992c480, []int{0}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionDynamicFeeTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.Merge(m, src)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionDynamicFeeTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionDynamicFeeTx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionOptionDynamicFeeTx)(nil), "ethermint.types.v1.ExtensionOptionDynamicFeeTx")
}

func init() {
	proto.RegisterFile("ethermint/types/v1/dynamic_fee.proto", fileDescriptor_9d7cf05c9992c480)
}

var fileDescriptor_9d7cf05c9992c480 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0xf0, 0xec, 0x45, 0x30, 0x27, 0x09, 0x1e, 0x44, 0x61, 0x2b, 0x22, 0x22, 0x42, 0x76,
	0x29, 0xbe, 0x41, 0xa9, 0x82, 0x27, 0x8b, 0x78, 0x12, 0xa1, 0x24, 0x9b, 0x31, 0x59, 0xec, 0xee,
	0x84, 0xdd, 0xb1, 0x24, 0xf8, 0x12, 0x3e, 0x56, 0x8f, 0x3d, 0x8a, 0x87, 0x22, 0xc9, 0x8b, 0x48,
	0xfe, 0x20, 0x3d, 0xcd, 0x07, 0xf3, 0xf1, 0x63, 0x26, 0xbc, 0x04, 0x2a, 0xc0, 0x19, 0x6d, 0x49,
	0x52, 0x5d, 0x82, 0x97, 0xeb, 0xa9, 0xcc, 0x6a, 0x9b, 0x18, 0xad, 0x96, 0x6f, 0x00, 0xa2, 0x74,
	0x48, 0x18, 0x45, 0xff, 0x2d, 0xd1, 0xb7, 0xc4, 0x7a, 0x7a, 0x7a, 0x9c, 0x63, 0x8e, 0xfd, 0x5a,
	0x76, 0x69, 0x68, 0x5e, 0x7c, 0x86, 0x67, 0x77, 0x15, 0x81, 0xf5, 0x1a, 0xed, 0x63, 0x49, 0x1a,
	0xed, 0x7c, 0xd0, 0xee, 0x01, 0x9e, 0xab, 0xe8, 0x35, 0x8c, 0x4c, 0x52, 0x2d, 0x4b, 0xa7, 0xd1,
	0x69, 0xaa, 0xbb, 0xa0, 0xe0, 0x84, 0x9d, 0xb3, 0xeb, 0xc3, 0x99, 0xd8, 0xec, 0x26, 0xc1, 0xcf,
	0x6e, 0x72, 0x95, 0x6b, 0x2a, 0x3e, 0x52, 0xa1, 0xd0, 0x48, 0x85, 0xde, 0xa0, 0x1f, 0x47, 0xec,
	0xb3, 0xf7, 0xe1, 0x4a, 0xf1, 0x60, 0xe9, 0xe9, 0xc8, 0x24, 0xd5, 0x62, 0x84, 0x16, 0x9d, 0x33,
	0x9b, 0x6f, 0x1a, 0xce, 0xb6, 0x0d, 0x67, 0xbf, 0x0d, 0x67, 0x5f, 0x2d, 0x0f, 0xb6, 0x2d, 0x0f,
	0xbe, 0x5b, 0x1e, 0xbc, 0xdc, 0xec, 0x99, 0x45, 0xbd, 0x82, 0x98, 0x20, 0x31, 0x32, 0x75, 0x3a,
	0xcb, 0x61, 0x05, 0xde, 0xc7, 0x0a, 0x1d, 0x0c, 0x76, 0x7a, 0xd0, 0x7f, 0x72, 0xfb, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x89, 0x1d, 0xf0, 0x2b, 0x1b, 0x01, 0x00, 0x00,
}

func (m *ExtensionOptionDynamicFeeTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionDynamicFeeTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionDynamicFeeTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxPriorityPrice.Size()
		i -= size
		if _, err := m.MaxPriorityPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDynamicFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDynamicFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtensionOptionDynamicFeeTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxPriorityPrice.Size()
	n += 1 + l + sovDynamicFee(uint64(l))
	return n
}

func sovDynamicFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicFee(x uint64) (n int) {
	return sovDynamicFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtensionOptionDynamicFeeTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicFee
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
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriorityPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicFee
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
				return ErrInvalidLengthDynamicFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPriorityPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDynamicFee
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
func skipDynamicFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
				return 0, ErrInvalidLengthDynamicFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicFee = fmt.Errorf("proto: unexpected end of group")
)
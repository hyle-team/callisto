// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/nft/nft.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// NFT defines the NFT.
type NFT struct {
	Address             string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owner               string     `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Uri                 string     `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	VestingPeriod       int64      `protobuf:"varint,4,opt,name=vesting_period,json=vestingPeriod,proto3" json:"vesting_period,omitempty"`
	RewardPerPeriod     types.Coin `protobuf:"bytes,5,opt,name=reward_per_period,json=rewardPerPeriod,proto3" json:"reward_per_period"`
	VestingPeriodsCount int64      `protobuf:"varint,6,opt,name=vesting_periods_count,json=vestingPeriodsCount,proto3" json:"vesting_periods_count,omitempty"`
	AvailableToWithdraw types.Coin `protobuf:"bytes,7,opt,name=available_to_withdraw,json=availableToWithdraw,proto3" json:"available_to_withdraw"`
	LastVestingTime     int64      `protobuf:"varint,8,opt,name=last_vesting_time,json=lastVestingTime,proto3" json:"last_vesting_time,omitempty"`
	VestingCounter      int64      `protobuf:"varint,9,opt,name=vesting_counter,json=vestingCounter,proto3" json:"vesting_counter,omitempty"`
	Denom               string     `protobuf:"bytes,10,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *NFT) Reset()         { *m = NFT{} }
func (m *NFT) String() string { return proto.CompactTextString(m) }
func (*NFT) ProtoMessage()    {}
func (*NFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3395a297238036, []int{0}
}
func (m *NFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFT.Merge(m, src)
}
func (m *NFT) XXX_Size() int {
	return m.Size()
}
func (m *NFT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFT.DiscardUnknown(m)
}

var xxx_messageInfo_NFT proto.InternalMessageInfo

func (m *NFT) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NFT) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *NFT) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NFT) GetVestingPeriod() int64 {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func (m *NFT) GetRewardPerPeriod() types.Coin {
	if m != nil {
		return m.RewardPerPeriod
	}
	return types.Coin{}
}

func (m *NFT) GetVestingPeriodsCount() int64 {
	if m != nil {
		return m.VestingPeriodsCount
	}
	return 0
}

func (m *NFT) GetAvailableToWithdraw() types.Coin {
	if m != nil {
		return m.AvailableToWithdraw
	}
	return types.Coin{}
}

func (m *NFT) GetLastVestingTime() int64 {
	if m != nil {
		return m.LastVestingTime
	}
	return 0
}

func (m *NFT) GetVestingCounter() int64 {
	if m != nil {
		return m.VestingCounter
	}
	return 0
}

func (m *NFT) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type Owner struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	NftAddress string `protobuf:"bytes,2,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a3395a297238036, []int{1}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return m.Size()
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Owner) GetNftAddress() string {
	if m != nil {
		return m.NftAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*NFT)(nil), "cosmos.nft.NFT")
	proto.RegisterType((*Owner)(nil), "cosmos.nft.Owner")
}

func init() { proto.RegisterFile("cosmos/nft/nft.proto", fileDescriptor_2a3395a297238036) }

var fileDescriptor_2a3395a297238036 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0xdd, 0x98, 0x6e, 0x6b, 0xa7, 0xe8, 0xda, 0xe9, 0x16, 0xa6, 0x45, 0xd2, 0xa5, 0x20, 0x2e,
	0x05, 0x33, 0xb4, 0x7e, 0x81, 0xbb, 0xe0, 0x8b, 0xa0, 0x65, 0x5d, 0x14, 0x7c, 0x09, 0x93, 0xcd,
	0x24, 0x1d, 0xdc, 0xcc, 0x8d, 0x33, 0x93, 0x8d, 0xfb, 0x17, 0xfe, 0x89, 0xbf, 0xd1, 0xc7, 0x3e,
	0xfa, 0x24, 0xb2, 0xfb, 0x23, 0x92, 0x99, 0x89, 0xd8, 0x07, 0xc1, 0x87, 0x90, 0xb9, 0xe7, 0x9c,
	0x7b, 0xef, 0xe1, 0xde, 0x8b, 0x86, 0x0b, 0xd0, 0x25, 0x68, 0x2a, 0x73, 0xd3, 0x7e, 0x71, 0xa5,
	0xc0, 0x00, 0x46, 0x0e, 0x8d, 0x65, 0x6e, 0x4e, 0x4f, 0x0a, 0x80, 0x62, 0xc9, 0xa9, 0x65, 0xd2,
	0x3a, 0xa7, 0x4c, 0xae, 0x9d, 0xec, 0xf4, 0xc2, 0x27, 0xa7, 0x4c, 0x73, 0xfa, 0xa5, 0xe6, 0x6a,
	0x4d, 0x57, 0x97, 0x29, 0x37, 0xec, 0x92, 0x56, 0xac, 0x10, 0x92, 0x19, 0x01, 0xd2, 0x6b, 0x9f,
	0xfa, 0x32, 0xac, 0x12, 0x94, 0x49, 0x09, 0xc6, 0x92, 0xda, 0xb3, 0xc3, 0x02, 0x0a, 0xb0, 0x4f,
	0xda, 0xbe, 0x3c, 0x1a, 0xfd, 0x5d, 0xbf, 0xab, 0xbc, 0x00, 0xe1, 0x6b, 0x9e, 0x7f, 0x0f, 0x51,
	0xf8, 0xf6, 0xf5, 0x1c, 0x13, 0xb4, 0xc7, 0xb2, 0x4c, 0x71, 0xad, 0x49, 0x30, 0x0a, 0xc6, 0xfb,
	0xb3, 0x2e, 0xc4, 0x43, 0xd4, 0x87, 0x46, 0x72, 0x45, 0x1e, 0x58, 0xdc, 0x05, 0xf8, 0x09, 0x0a,
	0x6b, 0x25, 0x48, 0x68, 0xb1, 0xf6, 0x89, 0x9f, 0xa1, 0xc7, 0x2b, 0xae, 0x8d, 0x90, 0x45, 0x52,
	0x71, 0x25, 0x20, 0x23, 0x3b, 0xa3, 0x60, 0x1c, 0xce, 0x1e, 0x79, 0xf4, 0xda, 0x82, 0xf8, 0x0d,
	0x3a, 0x54, 0xbc, 0x61, 0x2a, 0x6b, 0x55, 0x9d, 0xb2, 0x3f, 0x0a, 0xc6, 0x07, 0x57, 0x27, 0xb1,
	0x9f, 0x59, 0x6b, 0x36, 0xf6, 0x66, 0xe3, 0x29, 0x08, 0x39, 0xd9, 0xb9, 0xfd, 0x79, 0xd6, 0x9b,
	0x0d, 0x5c, 0xe6, 0x35, 0x57, 0xbe, 0xd8, 0x15, 0x3a, 0xbe, 0xdf, 0x53, 0x27, 0x0b, 0xa8, 0xa5,
	0x21, 0xbb, 0xb6, 0xf5, 0xd1, 0xbd, 0xd6, 0x7a, 0xda, 0x52, 0xf8, 0x3d, 0x3a, 0x66, 0x2b, 0x26,
	0x96, 0x2c, 0x5d, 0xf2, 0xc4, 0x40, 0xd2, 0x08, 0x73, 0x93, 0x29, 0xd6, 0x90, 0xbd, 0xff, 0x33,
	0x71, 0xf4, 0x27, 0x7b, 0x0e, 0x1f, 0x7d, 0x2e, 0xbe, 0x40, 0x87, 0x4b, 0xa6, 0x4d, 0xd2, 0xb9,
	0x31, 0xa2, 0xe4, 0xe4, 0xa1, 0x35, 0x31, 0x68, 0x89, 0x0f, 0x0e, 0x9f, 0x8b, 0x92, 0xe3, 0xe7,
	0x68, 0xd0, 0xc9, 0xac, 0x59, 0xae, 0xc8, 0xbe, 0x55, 0x76, 0xf3, 0x9b, 0x3a, 0xb4, 0x9d, 0x7c,
	0xc6, 0x25, 0x94, 0x04, 0xb9, 0xc9, 0xdb, 0xe0, 0x7c, 0x82, 0xfa, 0xef, 0xec, 0x0a, 0xfe, 0xbd,
	0xb2, 0x33, 0x74, 0x20, 0x73, 0x93, 0x74, 0xac, 0x5b, 0x1c, 0x92, 0xb9, 0x79, 0xe5, 0x90, 0xc9,
	0xe4, 0x76, 0x13, 0x05, 0x77, 0x9b, 0x28, 0xf8, 0xb5, 0x89, 0x82, 0x6f, 0xdb, 0xa8, 0x77, 0xb7,
	0x8d, 0x7a, 0x3f, 0xb6, 0x51, 0xef, 0xd3, 0xb8, 0x10, 0xe6, 0xa6, 0x4e, 0xe3, 0x05, 0x94, 0xd4,
	0x9f, 0x8e, 0xfb, 0xbd, 0xd0, 0xd9, 0x67, 0xfa, 0xd5, 0x1e, 0xb9, 0x59, 0x57, 0x5c, 0xa7, 0xbb,
	0xf6, 0x80, 0x5e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x11, 0x02, 0x91, 0xff, 0x02, 0x00,
	0x00,
}

func (m *NFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x52
	}
	if m.VestingCounter != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.VestingCounter))
		i--
		dAtA[i] = 0x48
	}
	if m.LastVestingTime != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.LastVestingTime))
		i--
		dAtA[i] = 0x40
	}
	{
		size, err := m.AvailableToWithdraw.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.VestingPeriodsCount != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.VestingPeriodsCount))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.RewardPerPeriod.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.VestingPeriod != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.VestingPeriod))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Owner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Owner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Owner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftAddress) > 0 {
		i -= len(m.NftAddress)
		copy(dAtA[i:], m.NftAddress)
		i = encodeVarintNft(dAtA, i, uint64(len(m.NftAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if m.VestingPeriod != 0 {
		n += 1 + sovNft(uint64(m.VestingPeriod))
	}
	l = m.RewardPerPeriod.Size()
	n += 1 + l + sovNft(uint64(l))
	if m.VestingPeriodsCount != 0 {
		n += 1 + sovNft(uint64(m.VestingPeriodsCount))
	}
	l = m.AvailableToWithdraw.Size()
	n += 1 + l + sovNft(uint64(l))
	if m.LastVestingTime != 0 {
		n += 1 + sovNft(uint64(m.LastVestingTime))
	}
	if m.VestingCounter != 0 {
		n += 1 + sovNft(uint64(m.VestingCounter))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func (m *Owner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.NftAddress)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: NFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			m.VestingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPerPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriodsCount", wireType)
			}
			m.VestingPeriodsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriodsCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableToWithdraw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AvailableToWithdraw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastVestingTime", wireType)
			}
			m.LastVestingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastVestingTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingCounter", wireType)
			}
			m.VestingCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func (m *Owner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: Owner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Owner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/incentive.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// RewardDistCache is the cache for reward distribution of BTC validators at a height
type RewardDistCache struct {
	TotalVotingPower uint64 `protobuf:"varint,1,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	// btc_vals is a list of BTC validators' voting power information
	BtcVals []*BTCValDistInfo `protobuf:"bytes,2,rep,name=btc_vals,json=btcVals,proto3" json:"btc_vals,omitempty"`
}

func (m *RewardDistCache) Reset()         { *m = RewardDistCache{} }
func (m *RewardDistCache) String() string { return proto.CompactTextString(m) }
func (*RewardDistCache) ProtoMessage()    {}
func (*RewardDistCache) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{0}
}
func (m *RewardDistCache) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardDistCache) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardDistCache.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardDistCache) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardDistCache.Merge(m, src)
}
func (m *RewardDistCache) XXX_Size() int {
	return m.Size()
}
func (m *RewardDistCache) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardDistCache.DiscardUnknown(m)
}

var xxx_messageInfo_RewardDistCache proto.InternalMessageInfo

func (m *RewardDistCache) GetTotalVotingPower() uint64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *RewardDistCache) GetBtcVals() []*BTCValDistInfo {
	if m != nil {
		return m.BtcVals
	}
	return nil
}

// BTCValDistInfo is the reward distribution of a BTC validator and its BTC delegations
type BTCValDistInfo struct {
	// btc_pk is the Bitcoin secp256k1 PK of this BTC validator
	// the PK follows encoding in BIP-340 spec
	BtcPk *github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,opt,name=btc_pk,json=btcPk,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"btc_pk,omitempty"`
	// babylon_pk is the Babylon public key of the BTC validator
	BabylonPk *secp256k1.PubKey `protobuf:"bytes,2,opt,name=babylon_pk,json=babylonPk,proto3" json:"babylon_pk,omitempty"`
	// commission defines the commission rate of BTC validator
	Commission *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=commission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission,omitempty"`
	// total_voting_power is the total voting power of the BTC validator
	TotalVotingPower uint64 `protobuf:"varint,4,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
	// btc_dels is a list of BTC delegations' voting power information under this BTC validator
	BtcDels []*BTCDelDistInfo `protobuf:"bytes,5,rep,name=btc_dels,json=btcDels,proto3" json:"btc_dels,omitempty"`
}

func (m *BTCValDistInfo) Reset()         { *m = BTCValDistInfo{} }
func (m *BTCValDistInfo) String() string { return proto.CompactTextString(m) }
func (*BTCValDistInfo) ProtoMessage()    {}
func (*BTCValDistInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{1}
}
func (m *BTCValDistInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCValDistInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCValDistInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCValDistInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCValDistInfo.Merge(m, src)
}
func (m *BTCValDistInfo) XXX_Size() int {
	return m.Size()
}
func (m *BTCValDistInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCValDistInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BTCValDistInfo proto.InternalMessageInfo

func (m *BTCValDistInfo) GetBabylonPk() *secp256k1.PubKey {
	if m != nil {
		return m.BabylonPk
	}
	return nil
}

func (m *BTCValDistInfo) GetTotalVotingPower() uint64 {
	if m != nil {
		return m.TotalVotingPower
	}
	return 0
}

func (m *BTCValDistInfo) GetBtcDels() []*BTCDelDistInfo {
	if m != nil {
		return m.BtcDels
	}
	return nil
}

// BTCDelDistInfo contains the information related to reward distribution for a BTC delegations
type BTCDelDistInfo struct {
	// babylon_pk is the Babylon public key of the BTC delegations
	BabylonPk *secp256k1.PubKey `protobuf:"bytes,1,opt,name=babylon_pk,json=babylonPk,proto3" json:"babylon_pk,omitempty"`
	// voting_power is the voting power of the BTC delegation
	VotingPower uint64 `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *BTCDelDistInfo) Reset()         { *m = BTCDelDistInfo{} }
func (m *BTCDelDistInfo) String() string { return proto.CompactTextString(m) }
func (*BTCDelDistInfo) ProtoMessage()    {}
func (*BTCDelDistInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac354c3bd6d7a66b, []int{2}
}
func (m *BTCDelDistInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCDelDistInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCDelDistInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCDelDistInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCDelDistInfo.Merge(m, src)
}
func (m *BTCDelDistInfo) XXX_Size() int {
	return m.Size()
}
func (m *BTCDelDistInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCDelDistInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BTCDelDistInfo proto.InternalMessageInfo

func (m *BTCDelDistInfo) GetBabylonPk() *secp256k1.PubKey {
	if m != nil {
		return m.BabylonPk
	}
	return nil
}

func (m *BTCDelDistInfo) GetVotingPower() uint64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*RewardDistCache)(nil), "babylon.btcstaking.v1.RewardDistCache")
	proto.RegisterType((*BTCValDistInfo)(nil), "babylon.btcstaking.v1.BTCValDistInfo")
	proto.RegisterType((*BTCDelDistInfo)(nil), "babylon.btcstaking.v1.BTCDelDistInfo")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/incentive.proto", fileDescriptor_ac354c3bd6d7a66b)
}

var fileDescriptor_ac354c3bd6d7a66b = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0xee, 0x0f, 0xcc, 0x9d, 0x00, 0x45, 0x20, 0x95, 0x1d, 0xd2, 0x52, 0x69, 0xa8,
	0x07, 0x66, 0xd3, 0x0e, 0x26, 0x4e, 0x08, 0x65, 0xb9, 0x4c, 0x80, 0x14, 0x45, 0xa8, 0x42, 0x5c,
	0x2a, 0xdb, 0x35, 0xa9, 0x95, 0x34, 0x8e, 0x6a, 0x2f, 0x23, 0x47, 0xbe, 0x01, 0x1f, 0x86, 0x0f,
	0xc1, 0x71, 0xe2, 0x84, 0x76, 0x98, 0x50, 0x7b, 0xe2, 0x5b, 0xa0, 0x38, 0x1e, 0x4b, 0x11, 0x50,
	0x71, 0x4a, 0xde, 0x3c, 0xcf, 0xfb, 0xbe, 0x4f, 0x7e, 0x89, 0xe1, 0x3e, 0x25, 0xb4, 0x48, 0x64,
	0x8a, 0xa9, 0x66, 0x4a, 0x93, 0x58, 0xa4, 0x11, 0xce, 0x07, 0x58, 0xa4, 0x8c, 0xa7, 0x5a, 0xe4,
	0x1c, 0x65, 0x73, 0xa9, 0xa5, 0x73, 0xcf, 0xda, 0xd0, 0xb5, 0x0d, 0xe5, 0x83, 0xbd, 0xbb, 0x91,
	0x8c, 0xa4, 0x71, 0xe0, 0xf2, 0xae, 0x32, 0xef, 0xdd, 0x67, 0x52, 0xcd, 0xa4, 0x1a, 0x57, 0x42,
	0x55, 0x58, 0xa9, 0x57, 0x55, 0x98, 0xcd, 0x8b, 0x4c, 0x4b, 0xac, 0x38, 0xcb, 0x86, 0x4f, 0x8f,
	0xe2, 0x01, 0x8e, 0x79, 0x61, 0x3d, 0xbd, 0x8f, 0x00, 0xde, 0x0e, 0xf9, 0x19, 0x99, 0x4f, 0x7c,
	0xa1, 0xf4, 0x31, 0x61, 0x53, 0xee, 0x3c, 0x82, 0x8e, 0x96, 0x9a, 0x24, 0xe3, 0x5c, 0x6a, 0x91,
	0x46, 0xe3, 0x4c, 0x9e, 0xf1, 0x79, 0x1b, 0x74, 0x41, 0x7f, 0x33, 0xbc, 0x63, 0x94, 0x91, 0x11,
	0x82, 0xf2, 0xb9, 0xf3, 0x02, 0xde, 0xa4, 0x9a, 0x8d, 0x73, 0x92, 0xa8, 0x76, 0xb3, 0xbb, 0xd1,
	0x6f, 0x0d, 0xf7, 0xd1, 0x1f, 0x5f, 0x00, 0x79, 0x6f, 0x8e, 0x47, 0x24, 0x29, 0xf7, 0x9c, 0xa4,
	0xef, 0x65, 0x78, 0x83, 0x6a, 0x36, 0x22, 0x89, 0xea, 0xfd, 0x68, 0xc2, 0x5b, 0xab, 0x9a, 0xf3,
	0x1a, 0x6e, 0x97, 0x43, 0xb3, 0xd8, 0xac, 0xdd, 0xf5, 0x8e, 0x2e, 0x2e, 0x3b, 0xc3, 0x48, 0xe8,
	0xe9, 0x29, 0x45, 0x4c, 0xce, 0xb0, 0x5d, 0xc0, 0xa6, 0x44, 0xa4, 0x57, 0x05, 0xd6, 0x45, 0xc6,
	0x15, 0xf2, 0x4e, 0x82, 0xc3, 0x27, 0x8f, 0x83, 0x53, 0xfa, 0x92, 0x17, 0xe1, 0x16, 0xd5, 0x2c,
	0x88, 0x9d, 0xe7, 0x10, 0x5a, 0x53, 0x39, 0xb2, 0xd9, 0x05, 0xfd, 0xd6, 0xb0, 0x83, 0x2c, 0xac,
	0x0a, 0x0f, 0xfa, 0x85, 0x07, 0xd9, 0xde, 0x1d, 0xdb, 0x12, 0xc4, 0xce, 0x5b, 0x08, 0x99, 0x9c,
	0xcd, 0x84, 0x52, 0x42, 0xa6, 0xed, 0x8d, 0x2e, 0xe8, 0xef, 0x78, 0xcf, 0x2e, 0x2e, 0x3b, 0x0f,
	0x6b, 0x91, 0xae, 0x60, 0x9b, 0xcb, 0x81, 0x9a, 0xc4, 0x36, 0x8f, 0xcf, 0xd9, 0xd7, 0xcf, 0x07,
	0xd0, 0x2e, 0xf3, 0x39, 0x0b, 0x6b, 0xb3, 0xfe, 0xc2, 0x7a, 0xf3, 0xdf, 0xac, 0x27, 0x3c, 0x51,
	0xed, 0xad, 0x75, 0xac, 0x7d, 0xbe, 0xca, 0xda, 0xe7, 0x89, 0xea, 0x29, 0x83, 0xba, 0x26, 0xfd,
	0xc6, 0x06, 0xfc, 0x37, 0x9b, 0x07, 0x70, 0x77, 0x25, 0x7b, 0xd3, 0x64, 0x6f, 0xe5, 0xd7, 0xb1,
	0xbd, 0x57, 0x5f, 0x16, 0x2e, 0x38, 0x5f, 0xb8, 0xe0, 0xfb, 0xc2, 0x05, 0x9f, 0x96, 0x6e, 0xe3,
	0x7c, 0xe9, 0x36, 0xbe, 0x2d, 0xdd, 0xc6, 0xbb, 0xb5, 0xdf, 0xf4, 0x43, 0xfd, 0xac, 0x18, 0xa0,
	0x74, 0xdb, 0xfc, 0xb9, 0x87, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x48, 0xf9, 0xd9, 0x4e,
	0x03, 0x00, 0x00,
}

func (m *RewardDistCache) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardDistCache) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardDistCache) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BtcVals) > 0 {
		for iNdEx := len(m.BtcVals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BtcVals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BTCValDistInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCValDistInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCValDistInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BtcDels) > 0 {
		for iNdEx := len(m.BtcDels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BtcDels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.TotalVotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x20
	}
	if m.Commission != nil {
		{
			size := m.Commission.Size()
			i -= size
			if _, err := m.Commission.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BabylonPk != nil {
		{
			size, err := m.BabylonPk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BtcPk != nil {
		{
			size := m.BtcPk.Size()
			i -= size
			if _, err := m.BtcPk.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BTCDelDistInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCDelDistInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCDelDistInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if m.BabylonPk != nil {
		{
			size, err := m.BabylonPk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIncentive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardDistCache) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalVotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.TotalVotingPower))
	}
	if len(m.BtcVals) > 0 {
		for _, e := range m.BtcVals {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *BTCValDistInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcPk != nil {
		l = m.BtcPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.BabylonPk != nil {
		l = m.BabylonPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.Commission != nil {
		l = m.Commission.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.TotalVotingPower))
	}
	if len(m.BtcDels) > 0 {
		for _, e := range m.BtcDels {
			l = e.Size()
			n += 1 + l + sovIncentive(uint64(l))
		}
	}
	return n
}

func (m *BTCDelDistInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BabylonPk != nil {
		l = m.BabylonPk.Size()
		n += 1 + l + sovIncentive(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovIncentive(uint64(m.VotingPower))
	}
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardDistCache) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: RewardDistCache: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardDistCache: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcVals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcVals = append(m.BtcVals, &BTCValDistInfo{})
			if err := m.BtcVals[len(m.BtcVals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *BTCValDistInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: BTCValDistInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCValDistInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcPk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.BtcPk = &v
			if err := m.BtcPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonPk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BabylonPk == nil {
				m.BabylonPk = &secp256k1.PubKey{}
			}
			if err := m.BabylonPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Commission = &v
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcDels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BtcDels = append(m.BtcDels, &BTCDelDistInfo{})
			if err := m.BtcDels[len(m.BtcDels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *BTCDelDistInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: BTCDelDistInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCDelDistInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonPk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BabylonPk == nil {
				m.BabylonPk = &secp256k1.PubKey{}
			}
			if err := m.BabylonPk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
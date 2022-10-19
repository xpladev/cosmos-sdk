// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// GenesisState defines the slashing module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// signing_infos represents a map between validator addresses and their
	// signing infos.
	SigningInfos []SigningInfo `protobuf:"bytes,2,rep,name=signing_infos,json=signingInfos,proto3" json:"signing_infos"`
	// missed_blocks represents a map between validator addresses and their
	// missed blocks.
	MissedBlocks []ValidatorMissedBlocks `protobuf:"bytes,3,rep,name=missed_blocks,json=missedBlocks,proto3" json:"missed_blocks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1923b9188b635394, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSigningInfos() []SigningInfo {
	if m != nil {
		return m.SigningInfos
	}
	return nil
}

func (m *GenesisState) GetMissedBlocks() []ValidatorMissedBlocks {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// SigningInfo stores validator signing info of corresponding address.
type SigningInfo struct {
	// address is the validator address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// validator_signing_info represents the signing info of this validator.
	ValidatorSigningInfo ValidatorSigningInfo `protobuf:"bytes,2,opt,name=validator_signing_info,json=validatorSigningInfo,proto3" json:"validator_signing_info"`
}

func (m *SigningInfo) Reset()         { *m = SigningInfo{} }
func (m *SigningInfo) String() string { return proto.CompactTextString(m) }
func (*SigningInfo) ProtoMessage()    {}
func (*SigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1923b9188b635394, []int{1}
}
func (m *SigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningInfo.Merge(m, src)
}
func (m *SigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *SigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SigningInfo proto.InternalMessageInfo

func (m *SigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SigningInfo) GetValidatorSigningInfo() ValidatorSigningInfo {
	if m != nil {
		return m.ValidatorSigningInfo
	}
	return ValidatorSigningInfo{}
}

// ValidatorMissedBlocks contains array of missed blocks of corresponding
// address.
type ValidatorMissedBlocks struct {
	// address is the validator address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// missed_blocks is an array of missed blocks by the validator.
	MissedBlocks []MissedBlock `protobuf:"bytes,2,rep,name=missed_blocks,json=missedBlocks,proto3" json:"missed_blocks"`
}

func (m *ValidatorMissedBlocks) Reset()         { *m = ValidatorMissedBlocks{} }
func (m *ValidatorMissedBlocks) String() string { return proto.CompactTextString(m) }
func (*ValidatorMissedBlocks) ProtoMessage()    {}
func (*ValidatorMissedBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_1923b9188b635394, []int{2}
}
func (m *ValidatorMissedBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMissedBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMissedBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMissedBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMissedBlocks.Merge(m, src)
}
func (m *ValidatorMissedBlocks) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMissedBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMissedBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMissedBlocks proto.InternalMessageInfo

func (m *ValidatorMissedBlocks) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorMissedBlocks) GetMissedBlocks() []MissedBlock {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// MissedBlock contains height and missed status as boolean.
type MissedBlock struct {
	// index is the height at which the block was missed.
	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// missed is the missed status.
	Missed bool `protobuf:"varint,2,opt,name=missed,proto3" json:"missed,omitempty"`
}

func (m *MissedBlock) Reset()         { *m = MissedBlock{} }
func (m *MissedBlock) String() string { return proto.CompactTextString(m) }
func (*MissedBlock) ProtoMessage()    {}
func (*MissedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_1923b9188b635394, []int{3}
}
func (m *MissedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissedBlock.Merge(m, src)
}
func (m *MissedBlock) XXX_Size() int {
	return m.Size()
}
func (m *MissedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MissedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MissedBlock proto.InternalMessageInfo

func (m *MissedBlock) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MissedBlock) GetMissed() bool {
	if m != nil {
		return m.Missed
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.slashing.v1beta1.GenesisState")
	proto.RegisterType((*SigningInfo)(nil), "cosmos.slashing.v1beta1.SigningInfo")
	proto.RegisterType((*ValidatorMissedBlocks)(nil), "cosmos.slashing.v1beta1.ValidatorMissedBlocks")
	proto.RegisterType((*MissedBlock)(nil), "cosmos.slashing.v1beta1.MissedBlock")
}

func init() {
	proto.RegisterFile("cosmos/slashing/v1beta1/genesis.proto", fileDescriptor_1923b9188b635394)
}

var fileDescriptor_1923b9188b635394 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xb1, 0x6e, 0xd4, 0x30,
	0x1c, 0xc6, 0xe3, 0x3b, 0x71, 0x50, 0xa7, 0x5d, 0xac, 0x50, 0x42, 0x87, 0xb4, 0x3a, 0x01, 0xaa,
	0x90, 0x92, 0xa8, 0xc7, 0xc8, 0x44, 0x96, 0x8a, 0x01, 0x09, 0xe5, 0x2a, 0x06, 0x06, 0x22, 0x27,
	0x71, 0x5d, 0xab, 0x89, 0x7d, 0xca, 0xdf, 0x9c, 0xda, 0x37, 0x60, 0xe4, 0x11, 0xe0, 0x0d, 0x40,
	0xe2, 0x21, 0x3a, 0x56, 0x4c, 0x4c, 0x08, 0xdd, 0x0d, 0xbc, 0x06, 0xaa, 0x9d, 0xa3, 0x29, 0x6a,
	0x74, 0x12, 0xcb, 0x5d, 0x1c, 0xff, 0xbe, 0xef, 0xff, 0xe5, 0x8b, 0x83, 0x1f, 0x17, 0x0a, 0x6a,
	0x05, 0x31, 0x54, 0x14, 0x4e, 0x84, 0xe4, 0xf1, 0xfc, 0x20, 0x67, 0x9a, 0x1e, 0xc4, 0x9c, 0x49,
	0x06, 0x02, 0xa2, 0x59, 0xa3, 0xb4, 0x22, 0x0f, 0x2c, 0x16, 0xad, 0xb0, 0xa8, 0xc5, 0x76, 0x3c,
	0xae, 0xb8, 0x32, 0x4c, 0x7c, 0x75, 0x65, 0xf1, 0x9d, 0x27, 0x7d, 0xae, 0x7f, 0xf5, 0x96, 0x7b,
	0x68, 0xb9, 0xcc, 0x1a, 0xb4, 0x33, 0x6e, 0x5a, 0xd4, 0x70, 0xa5, 0x8e, 0x2b, 0xc6, 0x69, 0x71,
	0x9e, 0xd1, 0x5a, 0x48, 0x15, 0x9b, 0x5f, 0xcb, 0x8d, 0x3f, 0x0c, 0xf0, 0xe6, 0xa1, 0xcd, 0x3a,
	0xd5, 0x54, 0x33, 0x92, 0xe0, 0xd1, 0x8c, 0x36, 0xb4, 0x06, 0x1f, 0xed, 0xa1, 0x7d, 0x77, 0xb2,
	0x1b, 0xf5, 0x64, 0x8f, 0x5e, 0x1b, 0x2c, 0xd9, 0xb8, 0xf8, 0xb9, 0xeb, 0x7c, 0xfa, 0xfd, 0xe5,
	0x29, 0x4a, 0x5b, 0x25, 0x39, 0xc2, 0x5b, 0x20, 0xb8, 0x14, 0x92, 0x67, 0x42, 0x1e, 0x2b, 0xf0,
	0x07, 0x7b, 0xc3, 0x7d, 0x77, 0xf2, 0xa8, 0xd7, 0x6a, 0x6a, 0xe9, 0x97, 0xf2, 0x58, 0x75, 0xfd,
	0x36, 0xe1, 0xfa, 0x3e, 0x90, 0x77, 0x78, 0xab, 0x16, 0x00, 0xac, 0xcc, 0xf2, 0x4a, 0x15, 0xa7,
	0xe0, 0x0f, 0x8d, 0x6b, 0xd4, 0xeb, 0xfa, 0x86, 0x56, 0xa2, 0xa4, 0x5a, 0x35, 0xaf, 0x8c, 0x2c,
	0x31, 0xaa, 0x1b, 0xfe, 0x75, 0x67, 0x63, 0xfc, 0x15, 0x61, 0xb7, 0x13, 0x84, 0x4c, 0xf0, 0x5d,
	0x5a, 0x96, 0x0d, 0x03, 0x5b, 0xc5, 0x46, 0xe2, 0x7f, 0xff, 0x16, 0x7a, 0xed, 0xb0, 0x17, 0x76,
	0x67, 0xaa, 0x1b, 0x21, 0x79, 0xba, 0x02, 0x89, 0xc4, 0xdb, 0xf3, 0xd5, 0xd4, 0xac, 0xdb, 0x81,
	0x3f, 0x30, 0x6d, 0x86, 0xeb, 0xc3, 0xf6, 0x74, 0xe1, 0xcd, 0x6f, 0x01, 0xc6, 0x9f, 0x11, 0xbe,
	0x7f, 0xeb, 0x63, 0xfe, 0x57, 0xfa, 0xa3, 0x7f, 0x1b, 0x5e, 0xf7, 0xde, 0x3a, 0x13, 0xfb, 0x7b,
	0x7d, 0x8e, 0xdd, 0x0e, 0x47, 0x3c, 0x7c, 0x47, 0xc8, 0x92, 0x9d, 0x99, 0x58, 0xc3, 0xd4, 0x2e,
	0xc8, 0x36, 0x1e, 0x59, 0x91, 0x29, 0xea, 0x5e, 0xda, 0xae, 0x92, 0xc3, 0x8b, 0x45, 0x80, 0x2e,
	0x17, 0x01, 0xfa, 0xb5, 0x08, 0xd0, 0xc7, 0x65, 0xe0, 0x5c, 0x2e, 0x03, 0xe7, 0xc7, 0x32, 0x70,
	0xde, 0x86, 0x5c, 0xe8, 0x93, 0xf7, 0x79, 0x54, 0xa8, 0xba, 0x3d, 0xfa, 0xed, 0x5f, 0x08, 0xe5,
	0x69, 0x7c, 0x76, 0xfd, 0xf1, 0xe8, 0xf3, 0x19, 0x83, 0x7c, 0x64, 0xce, 0xfb, 0xb3, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x74, 0xd0, 0xac, 0x6f, 0xb2, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SigningInfos) > 0 {
		for iNdEx := len(m.SigningInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SigningInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValidatorSigningInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMissedBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMissedBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMissedBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Missed {
		i--
		if m.Missed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SigningInfos) > 0 {
		for _, e := range m.SigningInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ValidatorSigningInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorMissedBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *MissedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	if m.Missed {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningInfos = append(m.SigningInfos, SigningInfo{})
			if err := m.SigningInfos[len(m.SigningInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, ValidatorMissedBlocks{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *SigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSigningInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorSigningInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorMissedBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorMissedBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMissedBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, MissedBlock{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *MissedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MissedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Missed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Missed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

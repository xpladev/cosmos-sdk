// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/v1beta1/auth.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BaseAccount) Reset()      { *m = BaseAccount{} }
func (*BaseAccount) ProtoMessage() {}
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{0}
}
func (m *BaseAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAccount.Merge(m, src)
}
func (m *BaseAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAccount proto.InternalMessageInfo

// ModuleAccount defines an account for modules that holds coins on a pool.
type ModuleAccount struct {
	*BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions  []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *ModuleAccount) Reset()      { *m = ModuleAccount{} }
func (*ModuleAccount) ProtoMessage() {}
func (*ModuleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{1}
}
func (m *ModuleAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleAccount.Merge(m, src)
}
func (m *ModuleAccount) XXX_Size() int {
	return m.Size()
}
func (m *ModuleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleAccount proto.InternalMessageInfo

// Params defines the parameters for the auth module.
type Params struct {
	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty"`
	SigVerifyCostED25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty"`
	SigVerifyCostSecp256k1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{2}
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

func (m *Params) GetMaxMemoCharacters() uint64 {
	if m != nil {
		return m.MaxMemoCharacters
	}
	return 0
}

func (m *Params) GetTxSigLimit() uint64 {
	if m != nil {
		return m.TxSigLimit
	}
	return 0
}

func (m *Params) GetTxSizeCostPerByte() uint64 {
	if m != nil {
		return m.TxSizeCostPerByte
	}
	return 0
}

func (m *Params) GetSigVerifyCostED25519() uint64 {
	if m != nil {
		return m.SigVerifyCostED25519
	}
	return 0
}

func (m *Params) GetSigVerifyCostSecp256k1() uint64 {
	if m != nil {
		return m.SigVerifyCostSecp256k1
	}
	return 0
}

func init() {
	proto.RegisterType((*BaseAccount)(nil), "cosmos.auth.v1beta1.BaseAccount")
	proto.RegisterType((*ModuleAccount)(nil), "cosmos.auth.v1beta1.ModuleAccount")
	proto.RegisterType((*Params)(nil), "cosmos.auth.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/auth/v1beta1/auth.proto", fileDescriptor_7e1f7e915d020d2d) }

var fileDescriptor_7e1f7e915d020d2d = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0x8e, 0x21, 0x3f, 0xfe, 0x6c, 0x00, 0x09, 0x93, 0x1f, 0x35, 0x39, 0x24, 0x16, 0x12, 0x6d,
	0x8a, 0x1a, 0xbb, 0x09, 0xa2, 0x55, 0xb9, 0x61, 0xda, 0x03, 0x6a, 0x69, 0x91, 0xa3, 0x72, 0xe8,
	0xc5, 0x5a, 0x3b, 0x83, 0xb1, 0xc8, 0x7a, 0x5d, 0xef, 0x1a, 0xc5, 0x3c, 0x41, 0xd5, 0x53, 0x7b,
	0xab, 0x7a, 0xe2, 0x11, 0x38, 0xf0, 0x10, 0x2d, 0x27, 0xd4, 0x53, 0xa5, 0x4a, 0x51, 0x15, 0x0e,
	0x54, 0x7d, 0x8a, 0x2a, 0xbb, 0x0e, 0x4d, 0x28, 0xb7, 0x9d, 0xef, 0xfb, 0x66, 0x66, 0xe7, 0x9b,
	0x5d, 0x54, 0xf6, 0x28, 0x23, 0x94, 0x99, 0x38, 0xe1, 0x07, 0xe6, 0x51, 0xdd, 0x05, 0x8e, 0xeb,
	0x22, 0x30, 0xa2, 0x98, 0x72, 0xaa, 0x2e, 0x48, 0xde, 0x10, 0x50, 0xc6, 0x97, 0xee, 0x66, 0x49,
	0x84, 0xf9, 0x66, 0x1b, 0x7c, 0xec, 0xa5, 0x0e, 0x26, 0x41, 0x48, 0xcd, 0xa3, 0xba, 0x29, 0x0e,
	0x32, 0xb9, 0xb4, 0x24, 0x75, 0x8e, 0x88, 0xcc, 0xac, 0x92, 0xa4, 0x8a, 0x3e, 0xf5, 0xa9, 0xc4,
	0xfb, 0xa7, 0x41, 0x82, 0x4f, 0xa9, 0xdf, 0x06, 0x53, 0x44, 0x6e, 0xb2, 0x6f, 0xe2, 0x30, 0x95,
	0xd4, 0xf2, 0xc7, 0x31, 0x54, 0xb0, 0x30, 0x83, 0x4d, 0xcf, 0xa3, 0x49, 0xc8, 0xd5, 0x06, 0x9a,
	0xc4, 0xad, 0x56, 0x0c, 0x8c, 0x69, 0x8a, 0xae, 0x54, 0xa7, 0x2d, 0xed, 0xdb, 0x59, 0xad, 0x98,
	0xf5, 0xd8, 0x94, 0x4c, 0x93, 0xc7, 0x41, 0xe8, 0xdb, 0x03, 0xa1, 0xba, 0x87, 0x26, 0xa3, 0xc4,
	0x75, 0x0e, 0x21, 0xd5, 0xc6, 0x74, 0xa5, 0x5a, 0x68, 0x14, 0x0d, 0xd9, 0xd0, 0x18, 0x34, 0x34,
	0x36, 0xc3, 0xd4, 0xba, 0xf7, 0xbb, 0x5b, 0x29, 0x46, 0x89, 0xdb, 0x0e, 0xbc, 0xbe, 0xf6, 0x01,
	0x25, 0x01, 0x07, 0x12, 0xf1, 0xf4, 0xf3, 0xd5, 0xe9, 0x2a, 0xfa, 0x4b, 0xd8, 0x13, 0x51, 0xe2,
	0x3e, 0x87, 0x54, 0x5d, 0x41, 0x73, 0x58, 0x5e, 0xcb, 0x09, 0x13, 0xe2, 0x42, 0xac, 0x8d, 0xeb,
	0x4a, 0x35, 0x6f, 0xcf, 0x66, 0xe8, 0x4b, 0x01, 0xaa, 0x25, 0x34, 0xc5, 0xe0, 0x6d, 0x02, 0xa1,
	0x07, 0x5a, 0x5e, 0x08, 0xae, 0xe3, 0x8d, 0xb5, 0x77, 0x27, 0x95, 0xdc, 0xa7, 0x93, 0x4a, 0xee,
	0xd7, 0x49, 0x25, 0x77, 0x7e, 0x56, 0x9b, 0xca, 0xe6, 0xdc, 0x7e, 0x7f, 0x75, 0xba, 0xba, 0x28,
	0xc7, 0xaa, 0xb1, 0xd6, 0xa1, 0x39, 0xe4, 0xc1, 0xf2, 0x57, 0x05, 0xcd, 0xee, 0xd0, 0x56, 0xd2,
	0xbe, 0x76, 0x65, 0x1b, 0xcd, 0xb8, 0x98, 0x81, 0x93, 0x35, 0x16, 0xd6, 0x14, 0x1a, 0xba, 0x71,
	0xcb, 0x16, 0x8d, 0xa1, 0x4a, 0x56, 0xfe, 0xa2, 0x5b, 0x51, 0xec, 0x82, 0x3b, 0x64, 0xb0, 0x8a,
	0xf2, 0x21, 0x26, 0x20, 0x9c, 0x9a, 0xb6, 0xc5, 0x59, 0xd5, 0x51, 0x21, 0x82, 0x98, 0x04, 0x8c,
	0x05, 0x34, 0x64, 0xda, 0xb8, 0x3e, 0x5e, 0x9d, 0xb6, 0x87, 0xa1, 0x8d, 0xc7, 0x83, 0x39, 0xce,
	0xcf, 0x6a, 0x73, 0x23, 0x77, 0x13, 0x93, 0x68, 0x43, 0x93, 0x8c, 0xb0, 0xcb, 0x3f, 0xc6, 0xd0,
	0xc4, 0x2e, 0x8e, 0x31, 0x61, 0xaa, 0x81, 0x16, 0x08, 0xee, 0x38, 0x04, 0x08, 0x75, 0xbc, 0x03,
	0x1c, 0x63, 0x8f, 0x43, 0x2c, 0xd7, 0x9c, 0xb7, 0xe7, 0x09, 0xee, 0xec, 0x00, 0xa1, 0x5b, 0xd7,
	0x84, 0xaa, 0xa3, 0x19, 0xde, 0x71, 0x58, 0xe0, 0x3b, 0xed, 0x80, 0x04, 0x5c, 0xdc, 0x38, 0x6f,
	0x23, 0xde, 0x69, 0x06, 0xfe, 0x8b, 0x3e, 0xa2, 0x3e, 0x44, 0xff, 0x0b, 0xc5, 0x31, 0x38, 0x1e,
	0x65, 0xdc, 0x89, 0x20, 0x76, 0xdc, 0x94, 0x43, 0xb6, 0xa7, 0xf9, 0xbe, 0xf4, 0x18, 0xb6, 0x28,
	0xe3, 0xbb, 0x10, 0x5b, 0x29, 0x07, 0xf5, 0x15, 0xba, 0xd3, 0x2f, 0x78, 0x04, 0x71, 0xb0, 0x9f,
	0xca, 0x24, 0x68, 0x35, 0xd6, 0xd7, 0xeb, 0x4f, 0xe4, 0xea, 0x2c, 0xad, 0xd7, 0xad, 0x14, 0x9b,
	0x81, 0xbf, 0x27, 0x14, 0xfd, 0xd4, 0x67, 0x4f, 0x05, 0x6f, 0x17, 0xd9, 0x08, 0x2a, 0xb3, 0xd4,
	0xd7, 0x68, 0xe9, 0x66, 0x41, 0x06, 0x5e, 0xd4, 0x58, 0x7f, 0x74, 0x58, 0xd7, 0xfe, 0x13, 0x25,
	0x4b, 0xbd, 0x6e, 0x65, 0x71, 0xa4, 0x64, 0x73, 0xa0, 0xb0, 0x17, 0xd9, 0xad, 0xf8, 0xc6, 0x4a,
	0xf6, 0x66, 0x94, 0x9b, 0xee, 0x76, 0xe4, 0x97, 0x96, 0x96, 0x5a, 0x5b, 0x5f, 0x7a, 0x65, 0xe5,
	0xa2, 0x57, 0x56, 0x7e, 0xf6, 0xca, 0xca, 0x87, 0xcb, 0x72, 0xee, 0xe2, 0xb2, 0x9c, 0xfb, 0x7e,
	0x59, 0xce, 0xbd, 0xb9, 0xef, 0x07, 0xfc, 0x20, 0x71, 0x0d, 0x8f, 0x92, 0xec, 0x87, 0x9a, 0xff,
	0x56, 0xe1, 0x69, 0x04, 0xcc, 0x9d, 0x10, 0xbf, 0x64, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0x90, 0x75, 0x8d, 0x34, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MaxMemoCharacters != that1.MaxMemoCharacters {
		return false
	}
	if this.TxSigLimit != that1.TxSigLimit {
		return false
	}
	if this.TxSizeCostPerByte != that1.TxSizeCostPerByte {
		return false
	}
	if this.SigVerifyCostED25519 != that1.SigVerifyCostED25519 {
		return false
	}
	if this.SigVerifyCostSecp256k1 != that1.SigVerifyCostSecp256k1 {
		return false
	}
	return true
}
func (m *BaseAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if m.AccountNumber != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x18
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModuleAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Permissions[iNdEx])
			copy(dAtA[i:], m.Permissions[iNdEx])
			i = encodeVarintAuth(dAtA, i, uint64(len(m.Permissions[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.SigVerifyCostSecp256k1 != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.SigVerifyCostSecp256k1))
		i--
		dAtA[i] = 0x28
	}
	if m.SigVerifyCostED25519 != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.SigVerifyCostED25519))
		i--
		dAtA[i] = 0x20
	}
	if m.TxSizeCostPerByte != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.TxSizeCostPerByte))
		i--
		dAtA[i] = 0x18
	}
	if m.TxSigLimit != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.TxSigLimit))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxMemoCharacters != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.MaxMemoCharacters))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovAuth(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovAuth(uint64(m.Sequence))
	}
	return n
}

func (m *ModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			l = len(s)
			n += 1 + l + sovAuth(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxMemoCharacters != 0 {
		n += 1 + sovAuth(uint64(m.MaxMemoCharacters))
	}
	if m.TxSigLimit != 0 {
		n += 1 + sovAuth(uint64(m.TxSigLimit))
	}
	if m.TxSizeCostPerByte != 0 {
		n += 1 + sovAuth(uint64(m.TxSizeCostPerByte))
	}
	if m.SigVerifyCostED25519 != 0 {
		n += 1 + sovAuth(uint64(m.SigVerifyCostED25519))
	}
	if m.SigVerifyCostSecp256k1 != 0 {
		n += 1 + sovAuth(uint64(m.SigVerifyCostSecp256k1))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: BaseAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuth
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
func (m *ModuleAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: ModuleAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuth
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMemoCharacters", wireType)
			}
			m.MaxMemoCharacters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMemoCharacters |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSigLimit", wireType)
			}
			m.TxSigLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxSigLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSizeCostPerByte", wireType)
			}
			m.TxSizeCostPerByte = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxSizeCostPerByte |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostED25519", wireType)
			}
			m.SigVerifyCostED25519 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigVerifyCostED25519 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostSecp256k1", wireType)
			}
			m.SigVerifyCostSecp256k1 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigVerifyCostSecp256k1 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuth = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/new_tx/tx.proto

package tx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/iost-official/Go-IOS-Protocol/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionRaw struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	ActionName           string   `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRaw) Reset()         { *m = ActionRaw{} }
func (m *ActionRaw) String() string { return proto.CompactTextString(m) }
func (*ActionRaw) ProtoMessage()    {}
func (*ActionRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_3443414c29495769, []int{0}
}
func (m *ActionRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRaw.Unmarshal(m, b)
}
func (m *ActionRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRaw.Marshal(b, m, deterministic)
}
func (dst *ActionRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRaw.Merge(dst, src)
}
func (m *ActionRaw) XXX_Size() int {
	return xxx_messageInfo_ActionRaw.Size(m)
}
func (m *ActionRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRaw.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRaw proto.InternalMessageInfo

func (m *ActionRaw) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ActionRaw) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *ActionRaw) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type TxRaw struct {
	Time                 int64                  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Expiration           int64                  `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	GasLimit             uint64                 `protobuf:"varint,3,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasPrice             uint64                 `protobuf:"varint,4,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Actions              []*ActionRaw           `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	Signers              [][]byte               `protobuf:"bytes,6,rep,name=signers,proto3" json:"signers,omitempty"`
	Signs                []*common.SignatureRaw `protobuf:"bytes,7,rep,name=signs,proto3" json:"signs,omitempty"`
	Publisher            *common.SignatureRaw   `protobuf:"bytes,8,opt,name=publisher,proto3" json:"publisher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TxRaw) Reset()         { *m = TxRaw{} }
func (m *TxRaw) String() string { return proto.CompactTextString(m) }
func (*TxRaw) ProtoMessage()    {}
func (*TxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_3443414c29495769, []int{1}
}
func (m *TxRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxRaw.Unmarshal(m, b)
}
func (m *TxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxRaw.Marshal(b, m, deterministic)
}
func (dst *TxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRaw.Merge(dst, src)
}
func (m *TxRaw) XXX_Size() int {
	return xxx_messageInfo_TxRaw.Size(m)
}
func (m *TxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxRaw proto.InternalMessageInfo

func (m *TxRaw) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TxRaw) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *TxRaw) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TxRaw) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *TxRaw) GetActions() []*ActionRaw {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *TxRaw) GetSigners() [][]byte {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *TxRaw) GetSigns() []*common.SignatureRaw {
	if m != nil {
		return m.Signs
	}
	return nil
}

func (m *TxRaw) GetPublisher() *common.SignatureRaw {
	if m != nil {
		return m.Publisher
	}
	return nil
}

type ReceiptRaw struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptRaw) Reset()         { *m = ReceiptRaw{} }
func (m *ReceiptRaw) String() string { return proto.CompactTextString(m) }
func (*ReceiptRaw) ProtoMessage()    {}
func (*ReceiptRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_3443414c29495769, []int{2}
}
func (m *ReceiptRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptRaw.Unmarshal(m, b)
}
func (m *ReceiptRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptRaw.Marshal(b, m, deterministic)
}
func (dst *ReceiptRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptRaw.Merge(dst, src)
}
func (m *ReceiptRaw) XXX_Size() int {
	return xxx_messageInfo_ReceiptRaw.Size(m)
}
func (m *ReceiptRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptRaw.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptRaw proto.InternalMessageInfo

func (m *ReceiptRaw) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReceiptRaw) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type StatusRaw struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRaw) Reset()         { *m = StatusRaw{} }
func (m *StatusRaw) String() string { return proto.CompactTextString(m) }
func (*StatusRaw) ProtoMessage()    {}
func (*StatusRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_3443414c29495769, []int{3}
}
func (m *StatusRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRaw.Unmarshal(m, b)
}
func (m *StatusRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRaw.Marshal(b, m, deterministic)
}
func (dst *StatusRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRaw.Merge(dst, src)
}
func (m *StatusRaw) XXX_Size() int {
	return xxx_messageInfo_StatusRaw.Size(m)
}
func (m *StatusRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRaw.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRaw proto.InternalMessageInfo

func (m *StatusRaw) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusRaw) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TxReceiptRaw struct {
	TxHash               []byte        `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	GasUsage             uint64        `protobuf:"varint,2,opt,name=gasUsage,proto3" json:"gasUsage,omitempty"`
	Status               *StatusRaw    `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	SuccActionNum        int32         `protobuf:"varint,4,opt,name=succActionNum,proto3" json:"succActionNum,omitempty"`
	Receipts             []*ReceiptRaw `protobuf:"bytes,5,rep,name=receipts,proto3" json:"receipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TxReceiptRaw) Reset()         { *m = TxReceiptRaw{} }
func (m *TxReceiptRaw) String() string { return proto.CompactTextString(m) }
func (*TxReceiptRaw) ProtoMessage()    {}
func (*TxReceiptRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_3443414c29495769, []int{4}
}
func (m *TxReceiptRaw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxReceiptRaw.Unmarshal(m, b)
}
func (m *TxReceiptRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxReceiptRaw.Marshal(b, m, deterministic)
}
func (dst *TxReceiptRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReceiptRaw.Merge(dst, src)
}
func (m *TxReceiptRaw) XXX_Size() int {
	return xxx_messageInfo_TxReceiptRaw.Size(m)
}
func (m *TxReceiptRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReceiptRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxReceiptRaw proto.InternalMessageInfo

func (m *TxReceiptRaw) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TxReceiptRaw) GetGasUsage() uint64 {
	if m != nil {
		return m.GasUsage
	}
	return 0
}

func (m *TxReceiptRaw) GetStatus() *StatusRaw {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TxReceiptRaw) GetSuccActionNum() int32 {
	if m != nil {
		return m.SuccActionNum
	}
	return 0
}

func (m *TxReceiptRaw) GetReceipts() []*ReceiptRaw {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionRaw)(nil), "tx.ActionRaw")
	proto.RegisterType((*TxRaw)(nil), "tx.TxRaw")
	proto.RegisterType((*ReceiptRaw)(nil), "tx.ReceiptRaw")
	proto.RegisterType((*StatusRaw)(nil), "tx.StatusRaw")
	proto.RegisterType((*TxReceiptRaw)(nil), "tx.TxReceiptRaw")
}

func init() { proto.RegisterFile("core/new_tx/tx.proto", fileDescriptor_tx_3443414c29495769) }

var fileDescriptor_tx_3443414c29495769 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0xb2, 0xd9, 0x7c, 0x0c, 0x29, 0x07, 0x53, 0x21, 0xab, 0x07, 0x14, 0xad, 0x40, 0x44,
	0x3d, 0x24, 0x52, 0x39, 0xc1, 0x8d, 0x1b, 0x07, 0x54, 0x21, 0xb7, 0x9c, 0x38, 0x20, 0xd7, 0xb5,
	0xb6, 0x96, 0xd8, 0xf5, 0xca, 0x9e, 0x55, 0xcd, 0x1f, 0xe1, 0xcf, 0xf0, 0xe7, 0x90, 0xc7, 0x5e,
	0x37, 0x1c, 0x7a, 0xca, 0xbc, 0x79, 0x33, 0xcf, 0x2f, 0x6f, 0x16, 0xce, 0x95, 0x75, 0xfa, 0xd8,
	0xeb, 0xc7, 0x9f, 0x18, 0x8e, 0x18, 0x0e, 0x83, 0xb3, 0x68, 0xd9, 0x1c, 0xc3, 0xc5, 0x2b, 0x65,
	0xbb, 0xce, 0xf6, 0xc7, 0xf4, 0x93, 0x88, 0xe6, 0x07, 0x6c, 0x3e, 0x2b, 0x34, 0xb6, 0x17, 0xf2,
	0x91, 0x5d, 0xc0, 0x5a, 0xd9, 0x1e, 0x9d, 0x54, 0xc8, 0x67, 0xbb, 0xd9, 0x7e, 0x23, 0x0a, 0x66,
	0x6f, 0x00, 0x24, 0x0d, 0x5e, 0xcb, 0x4e, 0xf3, 0x39, 0xb1, 0x27, 0x1d, 0xc6, 0x60, 0x71, 0x2f,
	0x51, 0xf2, 0x8a, 0x18, 0xaa, 0x9b, 0x3f, 0x73, 0xa8, 0x6f, 0x43, 0x54, 0x66, 0xb0, 0x40, 0xd3,
	0x69, 0x52, 0xad, 0x04, 0xd5, 0x51, 0x51, 0x87, 0xc1, 0x38, 0x19, 0x35, 0x48, 0xb1, 0x12, 0x27,
	0x9d, 0xe8, 0xa6, 0x95, 0xfe, 0xab, 0xe9, 0x0c, 0x92, 0xea, 0x42, 0x14, 0x9c, 0xb9, 0x6f, 0xce,
	0x28, 0xcd, 0x17, 0x85, 0x23, 0xcc, 0xde, 0xc3, 0x2a, 0xf9, 0xf2, 0xbc, 0xde, 0x55, 0xfb, 0x17,
	0x57, 0x67, 0x07, 0x0c, 0x87, 0xf2, 0x2f, 0xc5, 0xc4, 0x32, 0x0e, 0x2b, 0x6f, 0xda, 0x5e, 0x3b,
	0xcf, 0x97, 0xbb, 0x6a, 0xbf, 0x15, 0x13, 0x64, 0x97, 0x50, 0xc7, 0xd2, 0xf3, 0x15, 0x09, 0x9c,
	0x1f, 0x72, 0x66, 0x37, 0xa6, 0xed, 0x25, 0x8e, 0x4e, 0x47, 0x9d, 0x34, 0xc2, 0xae, 0x60, 0x33,
	0x8c, 0x77, 0xbf, 0x8c, 0x7f, 0xd0, 0x8e, 0xaf, 0x77, 0xb3, 0x67, 0xe7, 0x9f, 0xc6, 0x9a, 0x4f,
	0x00, 0x42, 0x2b, 0x6d, 0x06, 0x9c, 0xc2, 0xf9, 0x3d, 0xa4, 0x70, 0x6a, 0x41, 0x75, 0xf4, 0x16,
	0xa3, 0xd7, 0x3d, 0xe6, 0xac, 0x27, 0xd8, 0x7c, 0x84, 0xcd, 0x0d, 0x4a, 0x1c, 0x7d, 0x5e, 0x55,
	0xf6, 0xbe, 0xac, 0xc6, 0x3a, 0xae, 0x76, 0xda, 0x7b, 0xd9, 0x4e, 0x67, 0x9a, 0x60, 0xf3, 0x77,
	0x06, 0xdb, 0xdb, 0x70, 0xf2, 0xf2, 0x6b, 0x58, 0x62, 0xf8, 0x22, 0xfd, 0x03, 0x09, 0x6c, 0x45,
	0x46, 0x39, 0xde, 0xef, 0x45, 0x23, 0xc5, 0x4b, 0x98, 0xbd, 0x83, 0xa5, 0xa7, 0xf7, 0xe9, 0x28,
	0x39, 0xdd, 0xe2, 0x48, 0x64, 0x92, 0xbd, 0x85, 0x33, 0x3f, 0x2a, 0x95, 0x62, 0xbf, 0x1e, 0x3b,
	0x3a, 0x53, 0x2d, 0xfe, 0x6f, 0xb2, 0x4b, 0x58, 0xbb, 0x64, 0x67, 0x3a, 0xd6, 0xcb, 0x28, 0xf7,
	0x64, 0x51, 0x14, 0xfe, 0x6e, 0x49, 0x5f, 0xec, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15,
	0xed, 0x83, 0x53, 0xe2, 0x02, 0x00, 0x00,
}
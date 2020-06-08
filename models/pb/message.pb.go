// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package MSG

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transaction struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	Value                float64  `protobuf:"fixed64,4,opt,name=Value,proto3" json:"Value,omitempty"`
	FileID               string   `protobuf:"bytes,5,opt,name=FileID,proto3" json:"FileID,omitempty"`
	Timestamp            string   `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Type                 int32    `protobuf:"varint,7,opt,name=Type,proto3" json:"Type,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Transaction) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Transaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Transaction) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Transaction) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *Transaction) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Transaction) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Transaction) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Block struct {
	Number               int64          `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	Hash                 string         `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Creator              string         `protobuf:"bytes,3,opt,name=Creator,proto3" json:"Creator,omitempty"`
	PrevHash             string         `protobuf:"bytes,4,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Timestamp            string         `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	MerkleRoot           string         `protobuf:"bytes,6,opt,name=MerkleRoot,proto3" json:"MerkleRoot,omitempty"`
	Txs                  []*Transaction `protobuf:"bytes,7,rep,name=Txs,proto3" json:"Txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Block) GetMerkleRoot() string {
	if m != nil {
		return m.MerkleRoot
	}
	return ""
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Account struct {
	Address              string   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	FileNumber           int32    `protobuf:"varint,2,opt,name=FileNumber,proto3" json:"FileNumber,omitempty"`
	Contribution         float64  `protobuf:"fixed64,3,opt,name=Contribution,proto3" json:"Contribution,omitempty"`
	ContributionFile     float64  `protobuf:"fixed64,4,opt,name=ContributionFile,proto3" json:"ContributionFile,omitempty"`
	ContributionTx       float64  `protobuf:"fixed64,5,opt,name=ContributionTx,proto3" json:"ContributionTx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`

	XXX_unrecognized []byte `json:"-"`
	XXX_sizecache    int32  `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetFileNumber() int32 {
	if m != nil {
		return m.FileNumber
	}
	return 0
}

func (m *Account) GetContribution() float64 {
	if m != nil {
		return m.Contribution
	}
	return 0
}

func (m *Account) GetContributionFile() float64 {
	if m != nil {
		return m.ContributionFile
	}
	return 0
}

func (m *Account) GetContributionTx() float64 {
	if m != nil {
		return m.ContributionTx
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "MSG.Transaction")
	proto.RegisterType((*Block)(nil), "MSG.Block")
	proto.RegisterType((*Account)(nil), "MSG.Account")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0xc6, 0xd3, 0x8d, 0x31, 0x38, 0x28, 0x21, 0x8d, 0x31, 0x8d, 0x31, 0x66, 0xd9, 0x85, 0x59,
	0xbc, 0xe0, 0x42, 0x9f, 0x00, 0x31, 0xa0, 0x17, 0x18, 0x53, 0x16, 0xef, 0xcb, 0x68, 0x74, 0x61,
	0x5b, 0x49, 0xdb, 0x19, 0x7c, 0x36, 0xe3, 0x13, 0xf8, 0x52, 0xa6, 0xa5, 0xe0, 0xc0, 0xbb, 0xf3,
	0xfd, 0x4e, 0xff, 0x7c, 0xdf, 0x69, 0xe1, 0xb4, 0xe4, 0x4a, 0xb1, 0x37, 0x3e, 0x5c, 0x4b, 0xa1,
	0x05, 0xf6, 0x67, 0xf3, 0x69, 0xfc, 0x8d, 0xa0, 0x97, 0x4a, 0x56, 0x29, 0x96, 0xe9, 0x5c, 0x54,
	0x18, 0x43, 0xeb, 0x91, 0xa9, 0x77, 0x82, 0x22, 0x94, 0x74, 0xa9, 0xad, 0x0d, 0x9b, 0x48, 0x51,
	0x12, 0x6f, 0xcb, 0x4c, 0x8d, 0xfb, 0xe0, 0xa5, 0x82, 0xf8, 0x96, 0x78, 0xa9, 0xc0, 0x67, 0x10,
	0xbc, 0xb2, 0xa2, 0xe6, 0xa4, 0x15, 0xa1, 0x04, 0xd1, 0xad, 0xc0, 0xe7, 0xd0, 0x9e, 0xe4, 0x05,
	0x7f, 0x7a, 0x20, 0x81, 0x5d, 0xe9, 0x14, 0xbe, 0x84, 0x6e, 0x9a, 0x97, 0x5c, 0x69, 0x56, 0xae,
	0x49, 0xdb, 0xb6, 0xfe, 0x80, 0xb9, 0x2f, 0xfd, 0x5c, 0x73, 0x12, 0x46, 0x28, 0x09, 0xa8, 0xad,
	0xcd, 0x49, 0x73, 0xcd, 0x74, 0xad, 0x48, 0xc7, 0x52, 0xa7, 0xe2, 0x1f, 0x04, 0xc1, 0x7d, 0x21,
	0xb2, 0x95, 0x59, 0xf1, 0x5c, 0x97, 0x0b, 0x2e, 0xad, 0x77, 0x9f, 0x3a, 0xb5, 0x4f, 0xe4, 0x35,
	0x12, 0x11, 0x08, 0xc7, 0x92, 0x33, 0x2d, 0xa4, 0x8b, 0xb0, 0x93, 0xf8, 0x02, 0x3a, 0x2f, 0x92,
	0x7f, 0xd8, 0x1d, 0x2d, 0xdb, 0xda, 0xeb, 0x43, 0xd7, 0xc1, 0xb1, 0xeb, 0x2b, 0x80, 0x19, 0x97,
	0xab, 0x82, 0x53, 0x21, 0xb4, 0x0b, 0xd5, 0x20, 0x38, 0x06, 0x3f, 0xdd, 0x28, 0x12, 0x46, 0x7e,
	0xd2, 0xbb, 0x1d, 0x0c, 0x67, 0xf3, 0xe9, 0xb0, 0x31, 0x78, 0x6a, 0x9a, 0xf1, 0x17, 0x82, 0x70,
	0x94, 0x65, 0xa2, 0xae, 0xb4, 0xf1, 0x38, 0x5a, 0x2e, 0x25, 0x57, 0xca, 0x3d, 0xc6, 0x4e, 0x9a,
	0x9b, 0xcc, 0x1c, 0x5d, 0x5a, 0xcf, 0xce, 0xa3, 0x41, 0x70, 0x0c, 0x27, 0x63, 0x51, 0x69, 0x99,
	0x2f, 0x6a, 0x73, 0xb4, 0x8d, 0x88, 0xe8, 0x01, 0xc3, 0x37, 0x30, 0x68, 0x6a, 0xb3, 0xdb, 0x3d,
	0xdd, 0x3f, 0x8e, 0xaf, 0xa1, 0xdf, 0x64, 0xe9, 0xc6, 0x86, 0x47, 0xf4, 0x88, 0x2e, 0xda, 0xf6,
	0x5f, 0xdd, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x43, 0x31, 0x23, 0x68, 0x02, 0x00, 0x00,
}

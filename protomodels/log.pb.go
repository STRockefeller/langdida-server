// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

package protomodels

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Log struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ReviewCards          int32                `protobuf:"varint,2,opt,name=review_cards,json=reviewCards,proto3" json:"review_cards,omitempty"`
	NewCards             int32                `protobuf:"varint,3,opt,name=new_cards,json=newCards,proto3" json:"new_cards,omitempty"`
	Streak               int32                `protobuf:"varint,4,opt,name=streak,proto3" json:"streak,omitempty"`
	StreakUpdated        bool                 `protobuf:"varint,5,opt,name=streak_updated,json=streakUpdated,proto3" json:"streak_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a153da538f858886, []int{0}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Log) GetReviewCards() int32 {
	if m != nil {
		return m.ReviewCards
	}
	return 0
}

func (m *Log) GetNewCards() int32 {
	if m != nil {
		return m.NewCards
	}
	return 0
}

func (m *Log) GetStreak() int32 {
	if m != nil {
		return m.Streak
	}
	return 0
}

func (m *Log) GetStreakUpdated() bool {
	if m != nil {
		return m.StreakUpdated
	}
	return false
}

func init() {
	proto.RegisterType((*Log)(nil), "protomodels.Log")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_a153da538f858886) }

var fileDescriptor_a153da538f858886 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xc9, 0x4f, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x53, 0xb9, 0xf9, 0x29, 0xa9, 0x39, 0xc5, 0x52,
	0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0xb1, 0xa4, 0xd2, 0x34, 0xfd, 0x92, 0xcc,
	0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x6a, 0xa5, 0x2d, 0x8c, 0x5c, 0xcc, 0x3e, 0xf9,
	0xe9, 0x42, 0x7a, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x52, 0x7a, 0x10, 0x7d, 0x7a, 0x30, 0x7d, 0x7a, 0x21, 0x30, 0x7d, 0x41, 0x60, 0x75, 0x42, 0x8a,
	0x5c, 0x3c, 0x45, 0xa9, 0x65, 0x99, 0xa9, 0xe5, 0xf1, 0xc9, 0x89, 0x45, 0x29, 0xc5, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xdc, 0x10, 0x31, 0x67, 0x90, 0x90, 0x90, 0x34, 0x17, 0x67, 0x1e,
	0x5c, 0x9e, 0x19, 0x2c, 0xcf, 0x91, 0x07, 0x93, 0x14, 0xe3, 0x62, 0x2b, 0x2e, 0x29, 0x4a, 0x4d,
	0xcc, 0x96, 0x60, 0x01, 0xcb, 0x40, 0x79, 0x42, 0xaa, 0x5c, 0x7c, 0x10, 0x56, 0x7c, 0x69, 0x01,
	0xc8, 0xa2, 0x14, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x5e, 0x88, 0x68, 0x28, 0x44, 0x30,
	0x89, 0x0d, 0xec, 0x30, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x13, 0x38, 0x3c, 0xf8,
	0x00, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: card.proto

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

type Language int32

const (
	Language_ENGLISH  Language = 0
	Language_JAPANESE Language = 1
	Language_FRENCH   Language = 2
)

var Language_name = map[int32]string{
	0: "ENGLISH",
	1: "JAPANESE",
	2: "FRENCH",
}

var Language_value = map[string]int32{
	"ENGLISH":  0,
	"JAPANESE": 1,
	"FRENCH":   2,
}

func (x Language) String() string {
	return proto.EnumName(Language_name, int32(x))
}

func (Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_95fd8cb6caa913ee, []int{0}
}

type Card struct {
	Index                *CardIndex           `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Labels               []string             `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	ExampleSentences     []string             `protobuf:"bytes,3,rep,name=example_sentences,json=exampleSentences,proto3" json:"example_sentences,omitempty"`
	Familiarity          int32                `protobuf:"varint,4,opt,name=familiarity,proto3" json:"familiarity,omitempty"`
	ReviewDate           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=review_date,json=reviewDate,proto3" json:"review_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_95fd8cb6caa913ee, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetIndex() *CardIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Card) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Card) GetExampleSentences() []string {
	if m != nil {
		return m.ExampleSentences
	}
	return nil
}

func (m *Card) GetFamiliarity() int32 {
	if m != nil {
		return m.Familiarity
	}
	return 0
}

func (m *Card) GetReviewDate() *timestamp.Timestamp {
	if m != nil {
		return m.ReviewDate
	}
	return nil
}

type RelatedCards struct {
	Synonyms             []*CardIndex `protobuf:"bytes,1,rep,name=synonyms,proto3" json:"synonyms,omitempty"`
	Antonyms             []*CardIndex `protobuf:"bytes,2,rep,name=antonyms,proto3" json:"antonyms,omitempty"`
	Origin               *CardIndex   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	Derivatives          []*CardIndex `protobuf:"bytes,4,rep,name=derivatives,proto3" json:"derivatives,omitempty"`
	InOtherLanguages     []*CardIndex `protobuf:"bytes,5,rep,name=in_other_languages,json=inOtherLanguages,proto3" json:"in_other_languages,omitempty"`
	Others               []*CardIndex `protobuf:"bytes,6,rep,name=others,proto3" json:"others,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RelatedCards) Reset()         { *m = RelatedCards{} }
func (m *RelatedCards) String() string { return proto.CompactTextString(m) }
func (*RelatedCards) ProtoMessage()    {}
func (*RelatedCards) Descriptor() ([]byte, []int) {
	return fileDescriptor_95fd8cb6caa913ee, []int{1}
}

func (m *RelatedCards) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedCards.Unmarshal(m, b)
}
func (m *RelatedCards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedCards.Marshal(b, m, deterministic)
}
func (m *RelatedCards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedCards.Merge(m, src)
}
func (m *RelatedCards) XXX_Size() int {
	return xxx_messageInfo_RelatedCards.Size(m)
}
func (m *RelatedCards) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedCards.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedCards proto.InternalMessageInfo

func (m *RelatedCards) GetSynonyms() []*CardIndex {
	if m != nil {
		return m.Synonyms
	}
	return nil
}

func (m *RelatedCards) GetAntonyms() []*CardIndex {
	if m != nil {
		return m.Antonyms
	}
	return nil
}

func (m *RelatedCards) GetOrigin() *CardIndex {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *RelatedCards) GetDerivatives() []*CardIndex {
	if m != nil {
		return m.Derivatives
	}
	return nil
}

func (m *RelatedCards) GetInOtherLanguages() []*CardIndex {
	if m != nil {
		return m.InOtherLanguages
	}
	return nil
}

func (m *RelatedCards) GetOthers() []*CardIndex {
	if m != nil {
		return m.Others
	}
	return nil
}

type CardIndex struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language             Language `protobuf:"varint,2,opt,name=language,proto3,enum=protomodels.Language" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardIndex) Reset()         { *m = CardIndex{} }
func (m *CardIndex) String() string { return proto.CompactTextString(m) }
func (*CardIndex) ProtoMessage()    {}
func (*CardIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_95fd8cb6caa913ee, []int{2}
}

func (m *CardIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardIndex.Unmarshal(m, b)
}
func (m *CardIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardIndex.Marshal(b, m, deterministic)
}
func (m *CardIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardIndex.Merge(m, src)
}
func (m *CardIndex) XXX_Size() int {
	return xxx_messageInfo_CardIndex.Size(m)
}
func (m *CardIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_CardIndex.DiscardUnknown(m)
}

var xxx_messageInfo_CardIndex proto.InternalMessageInfo

func (m *CardIndex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CardIndex) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_ENGLISH
}

func init() {
	proto.RegisterEnum("protomodels.Language", Language_name, Language_value)
	proto.RegisterType((*Card)(nil), "protomodels.Card")
	proto.RegisterType((*RelatedCards)(nil), "protomodels.RelatedCards")
	proto.RegisterType((*CardIndex)(nil), "protomodels.CardIndex")
}

func init() { proto.RegisterFile("card.proto", fileDescriptor_95fd8cb6caa913ee) }

var fileDescriptor_95fd8cb6caa913ee = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xda, 0x86, 0xf6, 0x65, 0x42, 0xe1, 0x49, 0x4c, 0xd6, 0x2e, 0x44, 0x3d, 0x55,
	0x80, 0x32, 0xad, 0x5c, 0x90, 0x38, 0x4d, 0x5b, 0x61, 0x43, 0x53, 0x41, 0x2e, 0xf7, 0xca, 0x5d,
	0xde, 0x82, 0xa5, 0xc4, 0xa9, 0x6c, 0xaf, 0xb4, 0xff, 0x28, 0xe2, 0xcf, 0x41, 0x71, 0xe2, 0xaa,
	0x5c, 0x72, 0x4a, 0xf2, 0xf9, 0xf7, 0xbd, 0xf7, 0x7d, 0x91, 0x01, 0x1e, 0x85, 0xce, 0xb3, 0xad,
	0xae, 0x6d, 0x8d, 0xb1, 0x7b, 0x54, 0x75, 0x4e, 0xa5, 0xb9, 0x78, 0x5b, 0xd4, 0x75, 0x51, 0xd2,
	0xa5, 0xd3, 0x36, 0xcf, 0x4f, 0x97, 0x56, 0x56, 0x64, 0xac, 0xa8, 0xb6, 0x2d, 0x3d, 0xfd, 0x1b,
	0xc0, 0xf0, 0x46, 0xe8, 0x1c, 0x3f, 0xc0, 0x48, 0xaa, 0x9c, 0xf6, 0x2c, 0x48, 0x83, 0x59, 0x3c,
	0x3f, 0xcf, 0x4e, 0xc6, 0x64, 0x0d, 0x71, 0xdf, 0x9c, 0xf2, 0x16, 0xc2, 0x73, 0x88, 0x4a, 0xb1,
	0xa1, 0xd2, 0xb0, 0x30, 0x1d, 0xcc, 0x26, 0xbc, 0xfb, 0xc2, 0xf7, 0xf0, 0x9a, 0xf6, 0xa2, 0xda,
	0x96, 0xb4, 0x36, 0xa4, 0x2c, 0xa9, 0x47, 0x32, 0x6c, 0xe0, 0x90, 0xa4, 0x3b, 0x58, 0x79, 0x1d,
	0x53, 0x88, 0x9f, 0x44, 0x25, 0x4b, 0x29, 0xb4, 0xb4, 0x07, 0x36, 0x4c, 0x83, 0xd9, 0x88, 0x9f,
	0x4a, 0xf8, 0x19, 0x62, 0x4d, 0x3b, 0x49, 0xbf, 0xd7, 0xb9, 0xb0, 0xc4, 0x46, 0x2e, 0xda, 0x45,
	0xd6, 0x96, 0xca, 0x7c, 0xa9, 0xec, 0xa7, 0x2f, 0xc5, 0xa1, 0xc5, 0x6f, 0x85, 0xa5, 0xe9, 0x9f,
	0x10, 0xce, 0x38, 0x95, 0xc2, 0x52, 0xde, 0xe4, 0x37, 0x38, 0x87, 0xb1, 0x39, 0xa8, 0x5a, 0x1d,
	0x2a, 0xc3, 0x82, 0x74, 0xd0, 0xd3, 0xf2, 0xc8, 0x35, 0x1e, 0xa1, 0x6c, 0xeb, 0x09, 0xfb, 0x3d,
	0x9e, 0xc3, 0x0c, 0xa2, 0x5a, 0xcb, 0x42, 0x2a, 0x36, 0xe8, 0xfd, 0x97, 0x1d, 0x85, 0x9f, 0x20,
	0xce, 0x49, 0xcb, 0x9d, 0xb0, 0x72, 0x47, 0x86, 0x0d, 0x7b, 0xd7, 0x9c, 0xa2, 0x78, 0x0b, 0x28,
	0xd5, 0xba, 0xb6, 0xbf, 0x48, 0xaf, 0x4b, 0xa1, 0x8a, 0x67, 0x51, 0x90, 0x61, 0xa3, 0xde, 0x01,
	0x89, 0x54, 0xdf, 0x1b, 0xc3, 0x83, 0xe7, 0x5d, 0xde, 0x46, 0x31, 0x2c, 0xea, 0x75, 0x76, 0xd4,
	0x94, 0xc3, 0xe4, 0x28, 0x22, 0xc2, 0x50, 0x89, 0x8a, 0xdc, 0xb5, 0x99, 0x70, 0xf7, 0x8e, 0x57,
	0x30, 0xf6, 0x69, 0x58, 0x98, 0x06, 0xb3, 0x57, 0xf3, 0x37, 0xff, 0x8d, 0xf4, 0xab, 0xf9, 0x11,
	0x7b, 0x77, 0x05, 0x63, 0xaf, 0x62, 0x0c, 0x2f, 0x17, 0xcb, 0xaf, 0x0f, 0xf7, 0xab, 0xbb, 0xe4,
	0x05, 0x9e, 0xc1, 0xf8, 0xdb, 0xf5, 0x8f, 0xeb, 0xe5, 0x62, 0xb5, 0x48, 0x02, 0x04, 0x88, 0xbe,
	0xf0, 0xc5, 0xf2, 0xe6, 0x2e, 0x09, 0x37, 0x91, 0x1b, 0xf9, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x04, 0x8b, 0x41, 0x55, 0xfd, 0x02, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spotigraph/squad/v1/squad.proto

package squadv1

import (
	fmt "fmt"

	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Squad repesents a collection of user of the same expertise.
type Squad struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Urn                  string   `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Squad) Reset()         { *m = Squad{} }
func (m *Squad) String() string { return proto.CompactTextString(m) }
func (*Squad) ProtoMessage()    {}
func (*Squad) Descriptor() ([]byte, []int) {
	return fileDescriptor_25c458d348b37ca7, []int{0}
}

func (m *Squad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squad.Unmarshal(m, b)
}

func (m *Squad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squad.Marshal(b, m, deterministic)
}

func (m *Squad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squad.Merge(m, src)
}

func (m *Squad) XXX_Size() int {
	return xxx_messageInfo_Squad.Size(m)
}

func (m *Squad) XXX_DiscardUnknown() {
	xxx_messageInfo_Squad.DiscardUnknown(m)
}

var xxx_messageInfo_Squad proto.InternalMessageInfo

func (m *Squad) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Squad) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Squad) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func init() {
	proto.RegisterType((*Squad)(nil), "spotigraph.squad.v1.Squad")
}

func init() { proto.RegisterFile("spotigraph/squad/v1/squad.proto", fileDescriptor_25c458d348b37ca7) }

var fileDescriptor_25c458d348b37ca7 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0xc8, 0x2f,
	0xc9, 0x4c, 0x2f, 0x4a, 0x2c, 0xc8, 0xd0, 0x2f, 0x2e, 0x2c, 0x4d, 0x4c, 0xd1, 0x2f, 0x33, 0x84,
	0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x11, 0x0a, 0xf4, 0x20, 0xe2, 0x65, 0x86,
	0x4a, 0xf6, 0x5c, 0xac, 0xc1, 0x20, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x08, 0x17, 0x6b, 0x4e, 0x62, 0x52, 0x6a, 0x8e,
	0x04, 0x13, 0x58, 0x08, 0xc2, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x93, 0x60, 0x06, 0x8b,
	0x81, 0x98, 0x4e, 0xd9, 0x5c, 0x0a, 0xf9, 0x45, 0xe9, 0x7a, 0x55, 0xa9, 0x79, 0x99, 0x25, 0x19,
	0x89, 0x45, 0x7a, 0x58, 0x2c, 0x71, 0xe2, 0x02, 0x5b, 0x11, 0x00, 0x72, 0x45, 0x00, 0x63, 0x14,
	0x3b, 0x58, 0xbc, 0xcc, 0x70, 0x11, 0x13, 0x73, 0x70, 0x70, 0xc4, 0x2a, 0x26, 0xe1, 0x60, 0x84,
	0x16, 0xb0, 0x42, 0xbd, 0x30, 0xc3, 0x53, 0xc8, 0xa2, 0x31, 0x60, 0xd1, 0x98, 0x30, 0xc3, 0x24,
	0x36, 0xb0, 0x4f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x0d, 0x1d, 0x34, 0xec, 0x00,
	0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spotigraph/events/v1/event.proto

package eventsv1

import (
	fmt "fmt"

	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// EventType enumerates all event type values.
type EventType int32

const (
	// Default value when no enumeration is specified.
	EventType_EVENT_TYPE_INVALID EventType = 0
	// Explicitly Unknown object value.
	EventType_EVENT_TYPE_UNKNOWN                EventType = 1
	EventType_EVENT_TYPE_CHAPTER_CREATED        EventType = 2
	EventType_EVENT_TYPE_CHAPTER_DELETED        EventType = 3
	EventType_EVENT_TYPE_CHAPTER_LABEL_UPDATED  EventType = 4
	EventType_EVENT_TYPE_CHAPTER_LEADER_UPDATED EventType = 5
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_INVALID",
	1: "EVENT_TYPE_UNKNOWN",
	2: "EVENT_TYPE_CHAPTER_CREATED",
	3: "EVENT_TYPE_CHAPTER_DELETED",
	4: "EVENT_TYPE_CHAPTER_LABEL_UPDATED",
	5: "EVENT_TYPE_CHAPTER_LEADER_UPDATED",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_INVALID":                0,
	"EVENT_TYPE_UNKNOWN":                1,
	"EVENT_TYPE_CHAPTER_CREATED":        2,
	"EVENT_TYPE_CHAPTER_DELETED":        3,
	"EVENT_TYPE_CHAPTER_LABEL_UPDATED":  4,
	"EVENT_TYPE_CHAPTER_LEADER_UPDATED": 5,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{0}
}

// Event describes event contract.
type Event struct {
	EventType     EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=spotigraph.events.v1.EventType" json:"event_type,omitempty"`
	EventId       string    `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateType string    `protobuf:"bytes,3,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	AggregateId   string    `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	Meta          *any.Any  `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Event_ChapterCreated
	//	*Event_ChapterDeleted
	//	*Event_ChapterLabelUpdated
	//	*Event_ChapterLeaderUpdated
	Payload              isEvent_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}

func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}

func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}

func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}

func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventType() EventType {
	if m != nil {
		return m.EventType
	}
	return EventType_EVENT_TYPE_INVALID
}

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetMeta() *any.Any {
	if m != nil {
		return m.Meta
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_ChapterCreated struct {
	ChapterCreated *ChapterCreated `protobuf:"bytes,10,opt,name=chapter_created,json=chapterCreated,proto3,oneof"`
}

type Event_ChapterDeleted struct {
	ChapterDeleted *ChapterDeleted `protobuf:"bytes,11,opt,name=chapter_deleted,json=chapterDeleted,proto3,oneof"`
}

type Event_ChapterLabelUpdated struct {
	ChapterLabelUpdated *ChapterLabelUpdated `protobuf:"bytes,12,opt,name=chapter_label_updated,json=chapterLabelUpdated,proto3,oneof"`
}

type Event_ChapterLeaderUpdated struct {
	ChapterLeaderUpdated *ChapterLeaderUpdated `protobuf:"bytes,13,opt,name=chapter_leader_updated,json=chapterLeaderUpdated,proto3,oneof"`
}

func (*Event_ChapterCreated) isEvent_Payload() {}

func (*Event_ChapterDeleted) isEvent_Payload() {}

func (*Event_ChapterLabelUpdated) isEvent_Payload() {}

func (*Event_ChapterLeaderUpdated) isEvent_Payload() {}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetChapterCreated() *ChapterCreated {
	if x, ok := m.GetPayload().(*Event_ChapterCreated); ok {
		return x.ChapterCreated
	}
	return nil
}

func (m *Event) GetChapterDeleted() *ChapterDeleted {
	if x, ok := m.GetPayload().(*Event_ChapterDeleted); ok {
		return x.ChapterDeleted
	}
	return nil
}

func (m *Event) GetChapterLabelUpdated() *ChapterLabelUpdated {
	if x, ok := m.GetPayload().(*Event_ChapterLabelUpdated); ok {
		return x.ChapterLabelUpdated
	}
	return nil
}

func (m *Event) GetChapterLeaderUpdated() *ChapterLeaderUpdated {
	if x, ok := m.GetPayload().(*Event_ChapterLeaderUpdated); ok {
		return x.ChapterLeaderUpdated
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_ChapterCreated)(nil),
		(*Event_ChapterDeleted)(nil),
		(*Event_ChapterLabelUpdated)(nil),
		(*Event_ChapterLeaderUpdated)(nil),
	}
}

// ChapterCreated is raised on chapter entity creation.
type ChapterCreated struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	LeaderId             string   `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterCreated) Reset()         { *m = ChapterCreated{} }
func (m *ChapterCreated) String() string { return proto.CompactTextString(m) }
func (*ChapterCreated) ProtoMessage()    {}
func (*ChapterCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{1}
}

func (m *ChapterCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterCreated.Unmarshal(m, b)
}

func (m *ChapterCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterCreated.Marshal(b, m, deterministic)
}

func (m *ChapterCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterCreated.Merge(m, src)
}

func (m *ChapterCreated) XXX_Size() int {
	return xxx_messageInfo_ChapterCreated.Size(m)
}

func (m *ChapterCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterCreated.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterCreated proto.InternalMessageInfo

func (m *ChapterCreated) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ChapterCreated) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ChapterCreated) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

// ChapterDeleted is raised on chapter entity deletion.
type ChapterDeleted struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterDeleted) Reset()         { *m = ChapterDeleted{} }
func (m *ChapterDeleted) String() string { return proto.CompactTextString(m) }
func (*ChapterDeleted) ProtoMessage()    {}
func (*ChapterDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{2}
}

func (m *ChapterDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterDeleted.Unmarshal(m, b)
}

func (m *ChapterDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterDeleted.Marshal(b, m, deterministic)
}

func (m *ChapterDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterDeleted.Merge(m, src)
}

func (m *ChapterDeleted) XXX_Size() int {
	return xxx_messageInfo_ChapterDeleted.Size(m)
}

func (m *ChapterDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterDeleted proto.InternalMessageInfo

func (m *ChapterDeleted) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

// ChapterLabelUpdated is raised on chapter entity label updates.
type ChapterLabelUpdated struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Old                  string   `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New                  string   `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterLabelUpdated) Reset()         { *m = ChapterLabelUpdated{} }
func (m *ChapterLabelUpdated) String() string { return proto.CompactTextString(m) }
func (*ChapterLabelUpdated) ProtoMessage()    {}
func (*ChapterLabelUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{3}
}

func (m *ChapterLabelUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterLabelUpdated.Unmarshal(m, b)
}

func (m *ChapterLabelUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterLabelUpdated.Marshal(b, m, deterministic)
}

func (m *ChapterLabelUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterLabelUpdated.Merge(m, src)
}

func (m *ChapterLabelUpdated) XXX_Size() int {
	return xxx_messageInfo_ChapterLabelUpdated.Size(m)
}

func (m *ChapterLabelUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterLabelUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterLabelUpdated proto.InternalMessageInfo

func (m *ChapterLabelUpdated) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ChapterLabelUpdated) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

func (m *ChapterLabelUpdated) GetNew() string {
	if m != nil {
		return m.New
	}
	return ""
}

// ChapterLeaderUpdated is raised on chapter entity leader updates.
type ChapterLeaderUpdated struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Old                  string   `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New                  string   `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChapterLeaderUpdated) Reset()         { *m = ChapterLeaderUpdated{} }
func (m *ChapterLeaderUpdated) String() string { return proto.CompactTextString(m) }
func (*ChapterLeaderUpdated) ProtoMessage()    {}
func (*ChapterLeaderUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae32168c2ebb17ec, []int{4}
}

func (m *ChapterLeaderUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterLeaderUpdated.Unmarshal(m, b)
}

func (m *ChapterLeaderUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterLeaderUpdated.Marshal(b, m, deterministic)
}

func (m *ChapterLeaderUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterLeaderUpdated.Merge(m, src)
}

func (m *ChapterLeaderUpdated) XXX_Size() int {
	return xxx_messageInfo_ChapterLeaderUpdated.Size(m)
}

func (m *ChapterLeaderUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterLeaderUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterLeaderUpdated proto.InternalMessageInfo

func (m *ChapterLeaderUpdated) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *ChapterLeaderUpdated) GetOld() string {
	if m != nil {
		return m.Old
	}
	return ""
}

func (m *ChapterLeaderUpdated) GetNew() string {
	if m != nil {
		return m.New
	}
	return ""
}

func init() {
	proto.RegisterEnum("spotigraph.events.v1.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "spotigraph.events.v1.Event")
	proto.RegisterType((*ChapterCreated)(nil), "spotigraph.events.v1.ChapterCreated")
	proto.RegisterType((*ChapterDeleted)(nil), "spotigraph.events.v1.ChapterDeleted")
	proto.RegisterType((*ChapterLabelUpdated)(nil), "spotigraph.events.v1.ChapterLabelUpdated")
	proto.RegisterType((*ChapterLeaderUpdated)(nil), "spotigraph.events.v1.ChapterLeaderUpdated")
}

func init() { proto.RegisterFile("spotigraph/events/v1/event.proto", fileDescriptor_ae32168c2ebb17ec) }

var fileDescriptor_ae32168c2ebb17ec = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xda, 0x4c,
	0x10, 0xc6, 0x1c, 0xfe, 0xc4, 0x43, 0xc2, 0x8f, 0x36, 0x6e, 0x44, 0x52, 0xa9, 0x05, 0x2b, 0x91,
	0x68, 0x2e, 0x8c, 0x48, 0xef, 0x2b, 0x19, 0xbc, 0x12, 0x28, 0x16, 0x41, 0x0e, 0xd0, 0x83, 0x22,
	0x59, 0x0b, 0xbb, 0x35, 0x48, 0xae, 0x6d, 0x19, 0x43, 0x45, 0x1f, 0xa7, 0x97, 0x7d, 0x92, 0xa8,
	0x4f, 0x55, 0x79, 0x17, 0x1b, 0xd2, 0xb8, 0x51, 0xd5, 0xbb, 0xf1, 0xb7, 0xdf, 0x61, 0x67, 0x34,
	0x6b, 0xa8, 0x2f, 0x03, 0x3f, 0x5a, 0x38, 0x21, 0x09, 0xe6, 0x2d, 0xb6, 0x66, 0x5e, 0xb4, 0x6c,
	0xad, 0xdb, 0xa2, 0xd2, 0x82, 0xd0, 0x8f, 0x7c, 0xa4, 0xec, 0x18, 0x9a, 0x60, 0x68, 0xeb, 0xf6,
	0xf9, 0x99, 0xe3, 0xfb, 0x8e, 0xcb, 0x5a, 0x9c, 0x33, 0x5d, 0x7d, 0x6e, 0x11, 0x6f, 0x23, 0x04,
	0xea, 0x43, 0x11, 0x4a, 0x38, 0x26, 0xa2, 0x77, 0x00, 0x5c, 0x61, 0x47, 0x9b, 0x80, 0xd5, 0xa4,
	0xba, 0xd4, 0xac, 0x5c, 0xbf, 0xd6, 0xb2, 0xfc, 0x34, 0x2e, 0x18, 0x6d, 0x02, 0x66, 0xc9, 0x2c,
	0x29, 0xd1, 0x19, 0x1c, 0x0a, 0xfd, 0x82, 0xd6, 0xf2, 0x75, 0xa9, 0x29, 0x5b, 0x07, 0xfc, 0xbb,
	0x4f, 0xd1, 0x25, 0x54, 0x88, 0xe3, 0x84, 0xcc, 0x21, 0x11, 0x13, 0xf6, 0x05, 0x4e, 0x38, 0x4e,
	0x51, 0xee, 0xd0, 0x80, 0xa3, 0x1d, 0x6d, 0x41, 0x6b, 0x45, 0x4e, 0x2a, 0xa7, 0x58, 0x9f, 0xa2,
	0x26, 0x14, 0xbf, 0xb0, 0x88, 0xd4, 0x4a, 0x75, 0xa9, 0x59, 0xbe, 0x56, 0x34, 0xd1, 0x98, 0x96,
	0x34, 0xa6, 0xe9, 0xde, 0xc6, 0xe2, 0x0c, 0x74, 0x0b, 0xff, 0xcf, 0xe6, 0x24, 0x88, 0x58, 0x68,
	0xcf, 0x42, 0x46, 0x22, 0x46, 0x6b, 0xc0, 0x45, 0x17, 0xd9, 0x3d, 0x75, 0x05, 0xb9, 0x2b, 0xb8,
	0xbd, 0x9c, 0x55, 0x99, 0x3d, 0x42, 0xf6, 0x0d, 0x29, 0x73, 0x59, 0x6c, 0x58, 0xfe, 0x0b, 0x43,
	0x43, 0x70, 0xf7, 0x0c, 0xb7, 0x08, 0xb2, 0xe1, 0x45, 0x62, 0xe8, 0x92, 0x29, 0x73, 0xed, 0x55,
	0x40, 0xf9, 0x3d, 0x8f, 0xb8, 0xed, 0x9b, 0x67, 0x6d, 0xcd, 0x58, 0x31, 0x16, 0x82, 0x5e, 0xce,
	0x3a, 0x99, 0x3d, 0x85, 0xd1, 0x14, 0x4e, 0xd3, 0x00, 0x46, 0x28, 0x0b, 0xd3, 0x84, 0x63, 0x9e,
	0x70, 0xf5, 0x7c, 0x02, 0x97, 0xec, 0x22, 0x94, 0x59, 0x06, 0xde, 0x91, 0xe1, 0x20, 0x20, 0x1b,
	0xd7, 0x27, 0x54, 0x1d, 0x43, 0xe5, 0xf1, 0x10, 0x51, 0x15, 0x0a, 0xab, 0xd0, 0xe3, 0xbb, 0x24,
	0x5b, 0x71, 0x89, 0x14, 0x28, 0xf1, 0x5e, 0xb7, 0x1b, 0x22, 0x3e, 0xd0, 0x4b, 0x90, 0xb7, 0x17,
	0x5c, 0xd0, 0xed, 0x6a, 0x1c, 0x0a, 0xa0, 0x4f, 0x55, 0x35, 0xb5, 0x4d, 0x06, 0xf7, 0xc4, 0x56,
	0xbd, 0x81, 0x93, 0x8c, 0xb9, 0x64, 0xe4, 0x57, 0xa1, 0xe0, 0xbb, 0xc9, 0x7e, 0xc6, 0x65, 0x8c,
	0x78, 0xec, 0xeb, 0x36, 0x35, 0x2e, 0x55, 0x13, 0x94, 0xac, 0x11, 0xfc, 0x9b, 0xdb, 0xd5, 0x83,
	0x04, 0x72, 0xfa, 0x5e, 0xd0, 0x29, 0x20, 0x3c, 0xc1, 0x83, 0x91, 0x3d, 0xfa, 0x38, 0xc4, 0x76,
	0x7f, 0x30, 0xd1, 0xcd, 0xbe, 0x51, 0xcd, 0xfd, 0x86, 0x8f, 0x07, 0x37, 0x83, 0xdb, 0xf7, 0x83,
	0xaa, 0x84, 0x5e, 0xc1, 0xf9, 0x1e, 0xde, 0xed, 0xe9, 0xc3, 0x11, 0xb6, 0xec, 0xae, 0x85, 0xf5,
	0x11, 0x36, 0xaa, 0xf9, 0x3f, 0x9c, 0x1b, 0xd8, 0xc4, 0xf1, 0x79, 0x01, 0x5d, 0x40, 0x3d, 0xe3,
	0xdc, 0xd4, 0x3b, 0xd8, 0xb4, 0xc7, 0x43, 0x83, 0xbb, 0x14, 0xd1, 0x25, 0x34, 0xb2, 0x58, 0x58,
	0x37, 0xb0, 0x95, 0xd2, 0x4a, 0x1d, 0x1f, 0x1a, 0x7e, 0xe8, 0x68, 0xdf, 0x98, 0xb7, 0x88, 0xe6,
	0x24, 0xcc, 0xdc, 0x9e, 0x0e, 0xf0, 0x66, 0x87, 0xf1, 0x83, 0x1c, 0x4a, 0x9f, 0xc4, 0x2f, 0x61,
	0xb9, 0x6e, 0x7f, 0xcf, 0x17, 0xee, 0xf0, 0x87, 0x1f, 0x79, 0xe5, 0x6e, 0x27, 0xc2, 0x42, 0x34,
	0x69, 0xff, 0xdc, 0x87, 0xef, 0x05, 0x7c, 0x3f, 0x69, 0x4f, 0xff, 0xe3, 0xef, 0xfa, 0xed, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xca, 0xe5, 0xbc, 0xf8, 0x04, 0x00, 0x00,
}

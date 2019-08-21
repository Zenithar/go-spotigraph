// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spotigraph/chapter/v1/chapter_api.proto

package chapterv1

import (
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "go.zenithar.org/spotigraph/pkg/gen/go/spotigraph/system/v1"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	LeaderId             string   `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}

func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}

func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}

func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}

func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *CreateRequest) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}

func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}

func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}

func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}

func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRequest struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                *types.StringValue `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	LeaderId             *types.StringValue `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}

func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}

func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}

func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}

func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetLabel() *types.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *UpdateRequest) GetLeaderId() *types.StringValue {
	if m != nil {
		return m.LeaderId
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}

func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}

func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}

func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}

func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchRequest struct {
	Page                 uint32             `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32             `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Sorts                []string           `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Cursor               *types.StringValue `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	ChapterId            *types.StringValue `protobuf:"bytes,5,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
	Label                *types.StringValue `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{4}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}

func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}

func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}

func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}

func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchRequest) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *SearchRequest) GetSorts() []string {
	if m != nil {
		return m.Sorts
	}
	return nil
}

func (m *SearchRequest) GetCursor() *types.StringValue {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *SearchRequest) GetChapterId() *types.StringValue {
	if m != nil {
		return m.ChapterId
	}
	return nil
}

func (m *SearchRequest) GetLabel() *types.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

type CreateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Chapter  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{5}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}

func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}

func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}

func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}

func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CreateResponse) GetEntity() *Chapter {
	if m != nil {
		return m.Entity
	}
	return nil
}

type GetResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Chapter  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}

func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}

func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}

func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}

func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetResponse) GetEntity() *Chapter {
	if m != nil {
		return m.Entity
	}
	return nil
}

type UpdateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Chapter  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}

func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}

func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}

func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}

func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *UpdateResponse) GetEntity() *Chapter {
	if m != nil {
		return m.Entity
	}
	return nil
}

type DeleteResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}

func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}

func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}

func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}

func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SearchResponse struct {
	Error                *v1.Error          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Total                uint32             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PerPage              uint32             `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Count                uint32             `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CurrentPage          uint32             `protobuf:"varint,5,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	NextCursor           *types.StringValue `protobuf:"bytes,6,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	Members              []*Chapter         `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{9}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}

func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}

func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}

func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}

func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *SearchResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *SearchResponse) GetPerPage() uint32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *SearchResponse) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchResponse) GetCurrentPage() uint32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *SearchResponse) GetNextCursor() *types.StringValue {
	if m != nil {
		return m.NextCursor
	}
	return nil
}

func (m *SearchResponse) GetMembers() []*Chapter {
	if m != nil {
		return m.Members
	}
	return nil
}

type JoinRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	ChapterId            string   `protobuf:"bytes,2,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{10}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}

func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}

func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}

func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}

func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func (m *JoinRequest) GetChapterId() string {
	if m != nil {
		return m.ChapterId
	}
	return ""
}

type JoinResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{11}
}

func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}

func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}

func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}

func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}

func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LeaveRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	ChapterId            string   `protobuf:"bytes,2,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{12}
}

func (m *LeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRequest.Unmarshal(m, b)
}

func (m *LeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRequest.Marshal(b, m, deterministic)
}

func (m *LeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRequest.Merge(m, src)
}

func (m *LeaveRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveRequest.Size(m)
}

func (m *LeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRequest proto.InternalMessageInfo

func (m *LeaveRequest) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func (m *LeaveRequest) GetChapterId() string {
	if m != nil {
		return m.ChapterId
	}
	return ""
}

type LeaveResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LeaveResponse) Reset()         { *m = LeaveResponse{} }
func (m *LeaveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaveResponse) ProtoMessage()    {}
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abb19bb3b46a1265, []int{13}
}

func (m *LeaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveResponse.Unmarshal(m, b)
}

func (m *LeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveResponse.Marshal(b, m, deterministic)
}

func (m *LeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveResponse.Merge(m, src)
}

func (m *LeaveResponse) XXX_Size() int {
	return xxx_messageInfo_LeaveResponse.Size(m)
}

func (m *LeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveResponse proto.InternalMessageInfo

func (m *LeaveResponse) GetError() *v1.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "spotigraph.chapter.v1.CreateRequest")
	proto.RegisterType((*GetRequest)(nil), "spotigraph.chapter.v1.GetRequest")
	proto.RegisterType((*UpdateRequest)(nil), "spotigraph.chapter.v1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "spotigraph.chapter.v1.DeleteRequest")
	proto.RegisterType((*SearchRequest)(nil), "spotigraph.chapter.v1.SearchRequest")
	proto.RegisterType((*CreateResponse)(nil), "spotigraph.chapter.v1.CreateResponse")
	proto.RegisterType((*GetResponse)(nil), "spotigraph.chapter.v1.GetResponse")
	proto.RegisterType((*UpdateResponse)(nil), "spotigraph.chapter.v1.UpdateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "spotigraph.chapter.v1.DeleteResponse")
	proto.RegisterType((*SearchResponse)(nil), "spotigraph.chapter.v1.SearchResponse")
	proto.RegisterType((*JoinRequest)(nil), "spotigraph.chapter.v1.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "spotigraph.chapter.v1.JoinResponse")
	proto.RegisterType((*LeaveRequest)(nil), "spotigraph.chapter.v1.LeaveRequest")
	proto.RegisterType((*LeaveResponse)(nil), "spotigraph.chapter.v1.LeaveResponse")
}

func init() {
	proto.RegisterFile("spotigraph/chapter/v1/chapter_api.proto", fileDescriptor_abb19bb3b46a1265)
}

var fileDescriptor_abb19bb3b46a1265 = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0xd7, 0xf1, 0x26, 0x3e, 0x8e, 0x8d, 0x34, 0x34, 0x92, 0x71, 0x51, 0xe5, 0x6c, 0x02,
	0x18, 0x90, 0xd7, 0xd8, 0xd0, 0x42, 0x25, 0xb8, 0x88, 0x0d, 0xaa, 0x82, 0x2a, 0x11, 0xad, 0xe5,
	0x08, 0x35, 0xa5, 0xd6, 0xd8, 0x7b, 0x58, 0xaf, 0xb4, 0xd9, 0x59, 0x66, 0xc7, 0x26, 0x69, 0x05,
	0x57, 0xbc, 0x02, 0x12, 0xd7, 0x5c, 0x70, 0xc1, 0x63, 0x70, 0xc9, 0x2b, 0xf0, 0x26, 0xbd, 0x42,
	0xbb, 0x33, 0xa3, 0xac, 0xd3, 0xae, 0x63, 0x19, 0x29, 0xdc, 0xcd, 0xcf, 0xf7, 0x9d, 0x99, 0x73,
	0xe6, 0xfb, 0xce, 0xc0, 0x7b, 0x71, 0xc4, 0x84, 0xef, 0x71, 0x1a, 0xcd, 0x3a, 0xd3, 0x19, 0x8d,
	0x04, 0xf2, 0xce, 0xa2, 0xab, 0x87, 0x63, 0x1a, 0xf9, 0x76, 0xc4, 0x99, 0x60, 0x64, 0xef, 0x0a,
	0x68, 0xab, 0x5d, 0x7b, 0xd1, 0x6d, 0x1c, 0x79, 0xbe, 0x98, 0xcd, 0x27, 0xf6, 0x94, 0x9d, 0x77,
	0x30, 0x5c, 0xb0, 0xcb, 0x88, 0xb3, 0x8b, 0xcb, 0x4e, 0xca, 0x99, 0xb6, 0x3d, 0x0c, 0xdb, 0x0b,
	0x1a, 0xf8, 0x2e, 0x15, 0xd8, 0x79, 0x65, 0x20, 0x23, 0x37, 0xee, 0x79, 0x8c, 0x79, 0x01, 0x4a,
	0xce, 0x64, 0xfe, 0x7d, 0xe7, 0x47, 0x4e, 0xa3, 0x08, 0x79, 0xac, 0xf6, 0x0f, 0x56, 0x5e, 0x51,
	0x81, 0xf6, 0x33, 0xa0, 0xf8, 0x32, 0x16, 0x78, 0x9e, 0x60, 0xe4, 0x48, 0x42, 0xac, 0x5f, 0x0a,
	0x50, 0x1d, 0x70, 0xa4, 0x02, 0x1d, 0xfc, 0x61, 0x8e, 0xb1, 0x20, 0x9f, 0x43, 0x29, 0xa0, 0x13,
	0x0c, 0xea, 0x85, 0x66, 0xa1, 0x55, 0xee, 0xbf, 0xfb, 0xb2, 0x7f, 0xc0, 0xf7, 0x9b, 0x46, 0xab,
	0xd7, 0xbb, 0xfb, 0xec, 0xec, 0xa8, 0xfd, 0x84, 0xb6, 0x9f, 0x7f, 0xf7, 0x61, 0xeb, 0xac, 0xf9,
	0x51, 0xfb, 0xa1, 0x9e, 0xbd, 0xff, 0xc1, 0xa1, 0x23, 0x49, 0xe4, 0x3e, 0x94, 0x03, 0xa4, 0x2e,
	0xf2, 0xb1, 0xef, 0xd6, 0x8d, 0x34, 0x42, 0xfd, 0x65, 0x7f, 0x8f, 0xbf, 0xd9, 0xab, 0x3d, 0x3b,
	0xcb, 0x50, 0x0e, 0x7f, 0x2b, 0x34, 0x9d, 0x1d, 0x09, 0x3d, 0x76, 0xad, 0x07, 0x00, 0x8f, 0x50,
	0xe8, 0x2b, 0xb4, 0xc0, 0xf0, 0x5d, 0x75, 0x7e, 0x3e, 0xdb, 0xf0, 0x5d, 0xeb, 0xaf, 0x02, 0x54,
	0x47, 0x91, 0x9b, 0xb9, 0xfe, 0xda, 0x5c, 0xd2, 0xd3, 0x89, 0x26, 0xd7, 0xac, 0xf4, 0xde, 0xb6,
	0x65, 0xc9, 0x6d, 0x5d, 0x72, 0x7b, 0x28, 0xb8, 0x1f, 0x7a, 0xa7, 0x34, 0x98, 0xa3, 0x4e, 0x6f,
	0x98, 0x4d, 0xaf, 0x78, 0x33, 0x6f, 0xad, 0xe4, 0x1f, 0x42, 0xf5, 0x4b, 0x0c, 0x70, 0x83, 0x1c,
	0xac, 0x5f, 0x0d, 0xa8, 0x0e, 0x91, 0xf2, 0xe9, 0x4c, 0x73, 0x09, 0x6c, 0x45, 0xd4, 0xc3, 0x94,
	0x5d, 0x75, 0xd2, 0x31, 0x79, 0x0b, 0x76, 0x22, 0xe4, 0xe3, 0x74, 0xdd, 0x48, 0xd7, 0xb7, 0x23,
	0xe4, 0x27, 0xc9, 0xd6, 0x1d, 0x28, 0xc5, 0x8c, 0x8b, 0xb8, 0x5e, 0x6c, 0x16, 0x5b, 0x65, 0x47,
	0x4e, 0xc8, 0x27, 0x60, 0x4e, 0xe7, 0x3c, 0x66, 0xbc, 0xbe, 0xb5, 0x46, 0x6d, 0x14, 0x96, 0x8c,
	0x00, 0xb4, 0x45, 0x7c, 0xb7, 0x5e, 0xfa, 0x4f, 0xd5, 0x29, 0xab, 0x48, 0xc7, 0x99, 0x77, 0x32,
	0xd7, 0x7e, 0x27, 0xeb, 0x05, 0xd4, 0xb4, 0xaa, 0xe3, 0x88, 0x85, 0x31, 0x92, 0x2e, 0x94, 0x90,
	0x73, 0xc6, 0xd3, 0xc2, 0x54, 0x7a, 0x77, 0xed, 0x8c, 0x75, 0x95, 0x23, 0x16, 0x5d, 0xfb, 0xab,
	0x04, 0xe2, 0x48, 0x24, 0x79, 0x00, 0x26, 0x86, 0xc2, 0x17, 0x97, 0x4a, 0x21, 0xf7, 0xec, 0xd7,
	0xda, 0xdd, 0x1e, 0xc8, 0xa1, 0xa3, 0xd0, 0xd6, 0x05, 0x54, 0x52, 0x31, 0xdf, 0xfe, 0xc9, 0x2f,
	0xa0, 0xa6, 0xdd, 0x70, 0xfb, 0x87, 0x0f, 0xa0, 0xa6, 0x65, 0xbc, 0xf1, 0xe1, 0xd6, 0x1f, 0x06,
	0xd4, 0xb4, 0xa0, 0x37, 0x4f, 0xe1, 0x0e, 0x94, 0x04, 0x13, 0x34, 0x50, 0x6a, 0x97, 0x93, 0x25,
	0x1b, 0x14, 0x5f, 0xb1, 0xc1, 0x94, 0xcd, 0x43, 0x91, 0xea, 0xbd, 0xea, 0xc8, 0x09, 0xd9, 0x87,
	0xdd, 0xe9, 0x9c, 0x73, 0x0c, 0x85, 0x24, 0x95, 0xd2, 0xcd, 0x8a, 0x5a, 0x4b, 0x89, 0x5f, 0x40,
	0x25, 0xc4, 0x0b, 0x31, 0x56, 0x76, 0x59, 0x47, 0xa2, 0x90, 0x10, 0x06, 0xd2, 0x32, 0x9f, 0xc1,
	0xf6, 0x39, 0x9e, 0x4f, 0x90, 0xc7, 0xf5, 0xed, 0x66, 0x71, 0x8d, 0x62, 0x6b, 0xb8, 0xf5, 0x13,
	0x54, 0xbe, 0x66, 0x7e, 0xa8, 0x6d, 0x7f, 0x1f, 0xca, 0xc9, 0xef, 0xc0, 0xc2, 0xf1, 0x1a, 0x9d,
	0x63, 0x47, 0x42, 0x8f, 0x5d, 0xf2, 0xe9, 0x92, 0x65, 0x6f, 0xea, 0xd7, 0x57, 0xa6, 0xb4, 0x8e,
	0x60, 0x57, 0x1e, 0xbf, 0xf9, 0x53, 0xff, 0x0c, 0xbb, 0x8f, 0x91, 0x2e, 0xf0, 0xff, 0x4a, 0xa1,
	0x0f, 0x55, 0x75, 0xfe, 0xc6, 0x39, 0xf4, 0xfe, 0xd9, 0x02, 0x50, 0x4f, 0x73, 0x74, 0x72, 0x4c,
	0x46, 0x60, 0xca, 0xb6, 0x43, 0x0e, 0xf3, 0xde, 0x31, 0xfb, 0xd7, 0x36, 0xde, 0xb9, 0x01, 0xa5,
	0x2e, 0xf6, 0x18, 0x8a, 0x8f, 0x50, 0x90, 0xfd, 0x1c, 0xf4, 0xd5, 0xcf, 0xd9, 0xb0, 0x56, 0x41,
	0x54, 0xb4, 0x11, 0x98, 0xb2, 0x49, 0xe4, 0x5e, 0x72, 0xe9, 0x47, 0xcd, 0xbd, 0xe4, 0xb5, 0x4e,
	0x33, 0x02, 0x53, 0xda, 0x3f, 0x37, 0xec, 0xd2, 0x27, 0x97, 0x1b, 0xf6, 0x5a, 0x0f, 0x19, 0x81,
	0x29, 0xfb, 0x41, 0x6e, 0xd8, 0xa5, 0xff, 0x2f, 0x37, 0xec, 0xb5, 0xa6, 0xf2, 0x0d, 0x6c, 0x25,
	0xfa, 0x25, 0x79, 0x05, 0xcb, 0x78, 0xab, 0x71, 0xb0, 0x12, 0xa3, 0x02, 0x3a, 0x50, 0x4a, 0xd5,
	0x44, 0xf2, 0xd0, 0x59, 0xad, 0x37, 0x0e, 0x57, 0x83, 0x64, 0xcc, 0xfe, 0x05, 0x58, 0x8c, 0x7b,
	0xf6, 0x73, 0x0c, 0x7d, 0x31, 0xa3, 0xfc, 0xf5, 0x9c, 0xfe, 0x1b, 0x5a, 0x80, 0x91, 0x7f, 0x92,
	0xf4, 0x9b, 0x93, 0xc2, 0x13, 0xad, 0xf2, 0x45, 0xf7, 0x77, 0xa3, 0x38, 0x1c, 0x7c, 0xfb, 0xa7,
	0xb1, 0x37, 0xbc, 0xe2, 0x2a, 0xbc, 0x7d, 0xda, 0xfd, 0x3b, 0xbb, 0xfe, 0x54, 0xad, 0x3f, 0x3d,
	0xed, 0x4e, 0xcc, 0xb4, 0x73, 0x7d, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x8c, 0xe1,
	0x65, 0x0a, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChapterAPIClient is the client API for ChapterAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChapterAPIClient interface {
	// Create a chapter.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get a chapter by id.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update chapter attributes.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a chapter by id.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Search for chapters.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Join a chapter.
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	// Leave a chapter.
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
}

type chapterAPIClient struct {
	cc *grpc.ClientConn
}

func NewChapterAPIClient(cc *grpc.ClientConn) ChapterAPIClient {
	return &chapterAPIClient{cc}
}

func (c *chapterAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chapterAPIClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.chapter.v1.ChapterAPI/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChapterAPIServer is the server API for ChapterAPI service.
type ChapterAPIServer interface {
	// Create a chapter.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get a chapter by id.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update chapter attributes.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a chapter by id.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Search for chapters.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Join a chapter.
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	// Leave a chapter.
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
}

func RegisterChapterAPIServer(s *grpc.Server, srv ChapterAPIServer) {
	s.RegisterService(&_ChapterAPI_serviceDesc, srv)
}

func _ChapterAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChapterAPI_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChapterAPIServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.chapter.v1.ChapterAPI/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChapterAPIServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChapterAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotigraph.chapter.v1.ChapterAPI",
	HandlerType: (*ChapterAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChapterAPI_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ChapterAPI_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ChapterAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChapterAPI_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ChapterAPI_Search_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _ChapterAPI_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _ChapterAPI_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotigraph/chapter/v1/chapter_api.proto",
}

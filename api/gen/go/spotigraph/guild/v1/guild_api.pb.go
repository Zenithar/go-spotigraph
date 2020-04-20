// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spotigraph/guild/v1/guild_api.proto

package guildv1

import (
	context "context"
	fmt "fmt"

	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	v1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/system/v1"
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

type CreateRequest struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{0}
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
	return fileDescriptor_4379a0d23097e9ba, []int{1}
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
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                *wrappers.StringValue `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{2}
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

func (m *UpdateRequest) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
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
	return fileDescriptor_4379a0d23097e9ba, []int{3}
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
	Page                 uint32                `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              uint32                `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Sorts                []string              `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Cursor               *wrappers.StringValue `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	GuildId              *wrappers.StringValue `protobuf:"bytes,5,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	Label                *wrappers.StringValue `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{4}
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

func (m *SearchRequest) GetCursor() *wrappers.StringValue {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *SearchRequest) GetGuildId() *wrappers.StringValue {
	if m != nil {
		return m.GuildId
	}
	return nil
}

func (m *SearchRequest) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

type CreateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Guild    `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{5}
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

func (m *CreateResponse) GetEntity() *Guild {
	if m != nil {
		return m.Entity
	}
	return nil
}

type GetResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Guild    `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{6}
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

func (m *GetResponse) GetEntity() *Guild {
	if m != nil {
		return m.Entity
	}
	return nil
}

type UpdateResponse struct {
	Error                *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity               *Guild    `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{7}
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

func (m *UpdateResponse) GetEntity() *Guild {
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
	return fileDescriptor_4379a0d23097e9ba, []int{8}
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
	Error                *v1.Error             `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Total                uint32                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PerPage              uint32                `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Count                uint32                `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CurrentPage          uint32                `protobuf:"varint,5,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	NextCursor           *wrappers.StringValue `protobuf:"bytes,6,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	Members              []*Guild              `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{9}
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

func (m *SearchResponse) GetNextCursor() *wrappers.StringValue {
	if m != nil {
		return m.NextCursor
	}
	return nil
}

func (m *SearchResponse) GetMembers() []*Guild {
	if m != nil {
		return m.Members
	}
	return nil
}

type JoinRequest struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	GuildId              string   `protobuf:"bytes,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{10}
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

func (m *JoinRequest) GetGuildId() string {
	if m != nil {
		return m.GuildId
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
	return fileDescriptor_4379a0d23097e9ba, []int{11}
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
	GuildId              string   `protobuf:"bytes,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4379a0d23097e9ba, []int{12}
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

func (m *LeaveRequest) GetGuildId() string {
	if m != nil {
		return m.GuildId
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
	return fileDescriptor_4379a0d23097e9ba, []int{13}
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
	proto.RegisterType((*CreateRequest)(nil), "spotigraph.guild.v1.CreateRequest")
	proto.RegisterType((*GetRequest)(nil), "spotigraph.guild.v1.GetRequest")
	proto.RegisterType((*UpdateRequest)(nil), "spotigraph.guild.v1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "spotigraph.guild.v1.DeleteRequest")
	proto.RegisterType((*SearchRequest)(nil), "spotigraph.guild.v1.SearchRequest")
	proto.RegisterType((*CreateResponse)(nil), "spotigraph.guild.v1.CreateResponse")
	proto.RegisterType((*GetResponse)(nil), "spotigraph.guild.v1.GetResponse")
	proto.RegisterType((*UpdateResponse)(nil), "spotigraph.guild.v1.UpdateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "spotigraph.guild.v1.DeleteResponse")
	proto.RegisterType((*SearchResponse)(nil), "spotigraph.guild.v1.SearchResponse")
	proto.RegisterType((*JoinRequest)(nil), "spotigraph.guild.v1.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "spotigraph.guild.v1.JoinResponse")
	proto.RegisterType((*LeaveRequest)(nil), "spotigraph.guild.v1.LeaveRequest")
	proto.RegisterType((*LeaveResponse)(nil), "spotigraph.guild.v1.LeaveResponse")
}

func init() {
	proto.RegisterFile("spotigraph/guild/v1/guild_api.proto", fileDescriptor_4379a0d23097e9ba)
}

var fileDescriptor_4379a0d23097e9ba = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x16, 0x06, 0x0c, 0x39, 0xc4, 0x5c, 0x4c, 0xb6, 0x12, 0x65, 0xab, 0xae, 0x71, 0x56, 0x15,
	0x6d, 0x85, 0x29, 0xde, 0x6d, 0xa5, 0x95, 0xda, 0x0b, 0x48, 0x2b, 0x9a, 0xfe, 0x09, 0x8c, 0x1a,
	0x55, 0x9b, 0x6d, 0x90, 0x81, 0xa9, 0xb1, 0x64, 0x3c, 0xee, 0x78, 0x4c, 0x42, 0xde, 0xa2, 0x77,
	0xbd, 0xce, 0x65, 0xdf, 0xa1, 0x2f, 0xd0, 0x47, 0xca, 0x55, 0xe5, 0x19, 0xbb, 0x31, 0x09, 0x26,
	0x88, 0x8b, 0xf4, 0x6e, 0x3c, 0xf3, 0x7d, 0xdf, 0x39, 0x3e, 0xf3, 0x9d, 0x33, 0x70, 0x1c, 0xf8,
	0x84, 0x39, 0x36, 0xb5, 0xfc, 0x79, 0xdb, 0x0e, 0x1d, 0x77, 0xd6, 0x5e, 0x76, 0xc4, 0x62, 0x6c,
	0xf9, 0x8e, 0xee, 0x53, 0xc2, 0x08, 0x3a, 0xba, 0x03, 0xe9, 0xfc, 0x4c, 0x5f, 0x76, 0xea, 0x5d,
	0xdb, 0x61, 0xf3, 0x70, 0xa2, 0x4f, 0xc9, 0xa2, 0x8d, 0xbd, 0x25, 0x59, 0xf9, 0x94, 0x5c, 0xad,
	0xda, 0x9c, 0x31, 0x6d, 0xd9, 0xd8, 0x6b, 0x2d, 0x2d, 0xd7, 0x99, 0x59, 0x0c, 0xb7, 0x1f, 0x2c,
	0x84, 0x6e, 0xfd, 0x43, 0x9b, 0x10, 0xdb, 0xc5, 0x82, 0x33, 0x09, 0x7f, 0x6b, 0x5f, 0x52, 0xcb,
	0xf7, 0x31, 0x0d, 0xe2, 0xf3, 0x17, 0x99, 0xc9, 0xc5, 0x80, 0x46, 0x0a, 0x10, 0xac, 0x02, 0x86,
	0x17, 0x11, 0x42, 0xac, 0x04, 0x44, 0xfb, 0x11, 0x94, 0x13, 0x8a, 0x2d, 0x86, 0x4d, 0xfc, 0x7b,
	0x88, 0x03, 0x86, 0xbe, 0x84, 0xa2, 0x6b, 0x4d, 0xb0, 0x5b, 0xcb, 0xa9, 0xb9, 0xe6, 0x41, 0xef,
	0xa3, 0xdb, 0xde, 0x31, 0x6d, 0xa8, 0x52, 0xd3, 0x30, 0x9e, 0x5f, 0x9c, 0x77, 0x5b, 0x6f, 0xad,
	0xd6, 0xf5, 0xaf, 0x9f, 0x36, 0xcf, 0xd5, 0xcf, 0x5a, 0x6f, 0x92, 0xaf, 0x8f, 0x3f, 0x79, 0x69,
	0x0a, 0x92, 0xf6, 0x05, 0x40, 0x1f, 0xb3, 0x44, 0xab, 0x09, 0x92, 0x33, 0x8b, 0x85, 0x6a, 0xb7,
	0xbd, 0xf7, 0xe8, 0x91, 0x51, 0xbd, 0x38, 0x4f, 0x31, 0x5f, 0xfe, 0x99, 0x53, 0x4d, 0xc9, 0x99,
	0x69, 0x0b, 0x50, 0x7e, 0xf6, 0x67, 0xa9, 0x34, 0x76, 0xa6, 0x22, 0x23, 0x49, 0x58, 0x52, 0x73,
	0xcd, 0x8a, 0xf1, 0x81, 0x2e, 0xaa, 0xa6, 0x27, 0x55, 0xd3, 0x47, 0x8c, 0x3a, 0x9e, 0x7d, 0x66,
	0xb9, 0x21, 0x4e, 0xd2, 0x7c, 0x03, 0xca, 0xd7, 0xd8, 0xc5, 0x7b, 0x84, 0xd3, 0xfe, 0x90, 0x40,
	0x19, 0x61, 0x8b, 0x4e, 0xe7, 0x09, 0x17, 0x41, 0xc1, 0xb7, 0x6c, 0xcc, 0xd9, 0x8a, 0xc9, 0xd7,
	0xe8, 0x7d, 0x28, 0xfb, 0x98, 0x8e, 0xf9, 0xbe, 0xc4, 0xf7, 0x4b, 0x3e, 0xa6, 0x83, 0xe8, 0xe8,
	0x19, 0x14, 0x03, 0x42, 0x59, 0x50, 0xcb, 0xab, 0xf9, 0xe6, 0x81, 0x29, 0x3e, 0xd0, 0x6b, 0x90,
	0xa7, 0x21, 0x0d, 0x08, 0xad, 0x15, 0x76, 0xf8, 0x8d, 0x18, 0x8b, 0x86, 0x50, 0x16, 0x66, 0x74,
	0x66, 0xb5, 0xe2, 0xe3, 0xbc, 0x2d, 0xbf, 0x56, 0xe2, 0x3a, 0xa7, 0xa9, 0x72, 0xca, 0xbb, 0x97,
	0xf3, 0x12, 0xaa, 0x89, 0x89, 0x02, 0x9f, 0x78, 0x01, 0x46, 0x1d, 0x28, 0x62, 0x4a, 0x09, 0xe5,
	0x45, 0xa9, 0x18, 0xcf, 0xf5, 0x54, 0x8b, 0xc4, 0xfe, 0x5b, 0x76, 0xf4, 0x6f, 0x22, 0x88, 0x29,
	0x90, 0xc8, 0x00, 0x19, 0x7b, 0xcc, 0x61, 0xab, 0xf8, 0x22, 0xeb, 0xfa, 0x86, 0xb6, 0xd2, 0xfb,
	0xd1, 0xc2, 0x8c, 0x91, 0x1a, 0x83, 0x0a, 0xb7, 0xdb, 0xd3, 0x46, 0xbd, 0x84, 0x6a, 0x62, 0xd6,
	0xa7, 0x0d, 0x7c, 0x02, 0xd5, 0xc4, 0xb6, 0x7b, 0x07, 0xd6, 0x6e, 0x24, 0xa8, 0x26, 0x06, 0xde,
	0x3f, 0xfd, 0x67, 0x50, 0x64, 0x84, 0x59, 0x6e, 0xec, 0x6e, 0xf1, 0xb1, 0x66, 0xfb, 0xfc, 0x03,
	0xdb, 0x4f, 0x49, 0xe8, 0x31, 0xee, 0x6f, 0xc5, 0x14, 0x1f, 0xa8, 0x01, 0x87, 0xd3, 0x90, 0x52,
	0xec, 0x31, 0x41, 0x2a, 0xf2, 0xc3, 0x4a, 0xbc, 0xc7, 0x89, 0x5f, 0x41, 0xc5, 0xc3, 0x57, 0x6c,
	0x1c, 0xb7, 0xc7, 0x2e, 0xb6, 0x84, 0x88, 0x70, 0x22, 0x5a, 0xe4, 0x35, 0x94, 0x16, 0x78, 0x31,
	0xc1, 0x34, 0xa8, 0x95, 0xd4, 0xfc, 0x23, 0x85, 0x4e, 0xa0, 0xda, 0x0a, 0x2a, 0xdf, 0x11, 0xc7,
	0x4b, 0x5a, 0xfc, 0x73, 0x38, 0x88, 0xe6, 0x2e, 0xf1, 0xc6, 0x3b, 0x4c, 0x89, 0xb2, 0x80, 0x9e,
	0xce, 0xd0, 0xab, 0x54, 0x7b, 0x4a, 0x8f, 0xb0, 0x92, 0x06, 0xd4, 0xba, 0x70, 0x28, 0x42, 0xef,
	0x7f, 0xc5, 0xd7, 0x70, 0xf8, 0x03, 0xb6, 0x96, 0xf8, 0xff, 0x48, 0xbf, 0x07, 0x4a, 0x1c, 0x7b,
	0xef, 0xfc, 0x8d, 0xbf, 0x0b, 0x50, 0xe6, 0x17, 0xd2, 0x1d, 0x9c, 0xa2, 0x21, 0xc8, 0x62, 0xb8,
	0x20, 0x6d, 0xe3, 0xcd, 0xad, 0x3d, 0x5f, 0xf5, 0xe3, 0xad, 0x98, 0x38, 0xa5, 0x6f, 0x21, 0xdf,
	0xc7, 0x0c, 0xbd, 0xd8, 0xec, 0x84, 0xff, 0xde, 0xaf, 0xba, 0x9a, 0x0d, 0x88, 0x95, 0x86, 0x20,
	0x8b, 0x51, 0x90, 0x91, 0xdc, 0xda, 0xa3, 0x96, 0x91, 0xdc, 0xbd, 0x59, 0x32, 0x04, 0x59, 0x34,
	0x79, 0x86, 0xe4, 0xda, 0xc3, 0x95, 0x21, 0x79, 0x6f, 0x4a, 0x0c, 0x41, 0x16, 0x1d, 0x9f, 0x21,
	0xb9, 0xf6, 0x9e, 0x65, 0x48, 0xde, 0x1b, 0x19, 0xdf, 0x43, 0x21, 0x72, 0x29, 0xda, 0x5c, 0xa2,
	0x54, 0xef, 0xd4, 0x1b, 0x5b, 0x10, 0xb1, 0xd8, 0x4f, 0x50, 0xe4, 0x9e, 0x41, 0x9b, 0xb1, 0x69,
	0x2f, 0xd7, 0xb5, 0x6d, 0x10, 0xa1, 0xd7, 0xf3, 0x40, 0x25, 0xd4, 0xd6, 0xaf, 0xb1, 0xe7, 0xb0,
	0xb9, 0x45, 0x37, 0x31, 0x7a, 0x8a, 0x30, 0x98, 0xef, 0x0c, 0xa2, 0x09, 0x32, 0xc8, 0xbd, 0x15,
	0x0e, 0x5e, 0x76, 0x6e, 0xa4, 0xfc, 0xa8, 0xff, 0xcb, 0x5f, 0xd2, 0xd1, 0xe8, 0x8e, 0xc5, 0xb1,
	0xfa, 0x59, 0xe7, 0x9f, 0xf4, 0xee, 0x3b, 0xbe, 0xfb, 0xee, 0xac, 0x33, 0x91, 0xf9, 0x14, 0x7a,
	0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0xf2, 0x0a, 0x5c, 0x2e, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GuildAPIClient is the client API for GuildAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuildAPIClient interface {
	// Create a guild.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get a guild by id.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update guild attributes.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a guild by id.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Search for guilds.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Join a guild.
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	// Leave a guild.
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
}

type guildAPIClient struct {
	cc *grpc.ClientConn
}

func NewGuildAPIClient(cc *grpc.ClientConn) GuildAPIClient {
	return &guildAPIClient{cc}
}

func (c *guildAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildAPIClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, "/spotigraph.guild.v1.GuildAPI/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuildAPIServer is the server API for GuildAPI service.
type GuildAPIServer interface {
	// Create a guild.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get a guild by id.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update guild attributes.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a guild by id.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Search for guilds.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Join a guild.
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	// Leave a guild.
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
}

// UnimplementedGuildAPIServer can be embedded to have forward compatible implementations.
type UnimplementedGuildAPIServer struct {
}

func (*UnimplementedGuildAPIServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (*UnimplementedGuildAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (*UnimplementedGuildAPIServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (*UnimplementedGuildAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (*UnimplementedGuildAPIServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func (*UnimplementedGuildAPIServer) Join(ctx context.Context, req *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}

func (*UnimplementedGuildAPIServer) Leave(ctx context.Context, req *LeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}

func RegisterGuildAPIServer(s *grpc.Server, srv GuildAPIServer) {
	s.RegisterService(&_GuildAPI_serviceDesc, srv)
}

func _GuildAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildAPI_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildAPIServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotigraph.guild.v1.GuildAPI/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildAPIServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuildAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spotigraph.guild.v1.GuildAPI",
	HandlerType: (*GuildAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GuildAPI_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GuildAPI_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GuildAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GuildAPI_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _GuildAPI_Search_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _GuildAPI_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _GuildAPI_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotigraph/guild/v1/guild_api.proto",
}

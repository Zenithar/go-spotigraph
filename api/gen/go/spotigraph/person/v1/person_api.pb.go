// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: spotigraph/person/v1/person_api.proto

package personv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	v1 "zntr.io/spotigraph/api/gen/go/spotigraph/system/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetFirstName() *wrapperspb.StringValue {
	if x != nil {
		return x.FirstName
	}
	return nil
}

func (x *UpdateRequest) GetLastName() *wrapperspb.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage   uint32                  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Sorts     []string                `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Cursor    *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	PersonId  *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	Principal *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{4}
}

func (x *SearchRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRequest) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *SearchRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *SearchRequest) GetCursor() *wrapperspb.StringValue {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *SearchRequest) GetPersonId() *wrapperspb.StringValue {
	if x != nil {
		return x.PersonId
	}
	return nil
}

func (x *SearchRequest) GetPrincipal() *wrapperspb.StringValue {
	if x != nil {
		return x.Principal
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreateResponse) GetEntity() *Person {
	if x != nil {
		return x.Entity
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetResponse) GetEntity() *Person {
	if x != nil {
		return x.Entity
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Entity *Person   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *UpdateResponse) GetEntity() *Person {
	if x != nil {
		return x.Entity
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *v1.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error       *v1.Error               `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Total       uint32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	PerPage     uint32                  `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Count       uint32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	CurrentPage uint32                  `protobuf:"varint,5,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	NextCursor  *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	Members     []*Person               `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_person_v1_person_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_spotigraph_person_v1_person_api_proto_rawDescGZIP(), []int{9}
}

func (x *SearchResponse) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SearchResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchResponse) GetPerPage() uint32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *SearchResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchResponse) GetCurrentPage() uint32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *SearchResponse) GetNextCursor() *wrapperspb.StringValue {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

func (x *SearchResponse) GetMembers() []*Person {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_spotigraph_person_v1_person_api_proto protoreflect.FileDescriptor

var file_spotigraph_person_v1_person_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72,
	0x05, 0x20, 0x03, 0x28, 0x80, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x22, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15,
	0x72, 0x13, 0x32, 0x0e, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d,
	0x2b, 0x24, 0x98, 0x01, 0x20, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x32, 0x0e,
	0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x24, 0x98, 0x01,
	0x20, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x03,
	0x28, 0x80, 0x02, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x03, 0x28, 0x80, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x32, 0x0e, 0x5e, 0x5b, 0x30, 0x2d, 0x39,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x24, 0x98, 0x01, 0x20, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xa7, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x53, 0x0a,
	0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x18,
	0xfa, 0x42, 0x15, 0x72, 0x13, 0x32, 0x0e, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x5a, 0x61,
	0x2d, 0x7a, 0x5d, 0x2b, 0x24, 0x98, 0x01, 0x20, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x46, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x03, 0x28, 0x80, 0x02, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22, 0x79, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70,
	0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x76, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x79, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x6f, 0x74,
	0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa4, 0x02,
	0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x32, 0xab, 0x03, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41,
	0x50, 0x49, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73,
	0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20,
	0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x23, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xa1, 0x01, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x7a, 0x6e, 0x74, 0x72, 0x2e, 0x73,
	0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x7a, 0x6e, 0x74, 0x72, 0x2e, 0x69, 0x6f, 0x2f, 0x73,
	0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x50, 0x58, 0xaa, 0x02, 0x14, 0x53, 0x70, 0x6f, 0x74, 0x69,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x14, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spotigraph_person_v1_person_api_proto_rawDescOnce sync.Once
	file_spotigraph_person_v1_person_api_proto_rawDescData = file_spotigraph_person_v1_person_api_proto_rawDesc
)

func file_spotigraph_person_v1_person_api_proto_rawDescGZIP() []byte {
	file_spotigraph_person_v1_person_api_proto_rawDescOnce.Do(func() {
		file_spotigraph_person_v1_person_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_spotigraph_person_v1_person_api_proto_rawDescData)
	})
	return file_spotigraph_person_v1_person_api_proto_rawDescData
}

var file_spotigraph_person_v1_person_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_spotigraph_person_v1_person_api_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),          // 0: spotigraph.person.v1.CreateRequest
	(*GetRequest)(nil),             // 1: spotigraph.person.v1.GetRequest
	(*UpdateRequest)(nil),          // 2: spotigraph.person.v1.UpdateRequest
	(*DeleteRequest)(nil),          // 3: spotigraph.person.v1.DeleteRequest
	(*SearchRequest)(nil),          // 4: spotigraph.person.v1.SearchRequest
	(*CreateResponse)(nil),         // 5: spotigraph.person.v1.CreateResponse
	(*GetResponse)(nil),            // 6: spotigraph.person.v1.GetResponse
	(*UpdateResponse)(nil),         // 7: spotigraph.person.v1.UpdateResponse
	(*DeleteResponse)(nil),         // 8: spotigraph.person.v1.DeleteResponse
	(*SearchResponse)(nil),         // 9: spotigraph.person.v1.SearchResponse
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*v1.Error)(nil),               // 11: spotigraph.system.v1.Error
	(*Person)(nil),                 // 12: spotigraph.person.v1.Person
}
var file_spotigraph_person_v1_person_api_proto_depIdxs = []int32{
	10, // 0: spotigraph.person.v1.UpdateRequest.first_name:type_name -> google.protobuf.StringValue
	10, // 1: spotigraph.person.v1.UpdateRequest.last_name:type_name -> google.protobuf.StringValue
	10, // 2: spotigraph.person.v1.SearchRequest.cursor:type_name -> google.protobuf.StringValue
	10, // 3: spotigraph.person.v1.SearchRequest.person_id:type_name -> google.protobuf.StringValue
	10, // 4: spotigraph.person.v1.SearchRequest.principal:type_name -> google.protobuf.StringValue
	11, // 5: spotigraph.person.v1.CreateResponse.error:type_name -> spotigraph.system.v1.Error
	12, // 6: spotigraph.person.v1.CreateResponse.entity:type_name -> spotigraph.person.v1.Person
	11, // 7: spotigraph.person.v1.GetResponse.error:type_name -> spotigraph.system.v1.Error
	12, // 8: spotigraph.person.v1.GetResponse.entity:type_name -> spotigraph.person.v1.Person
	11, // 9: spotigraph.person.v1.UpdateResponse.error:type_name -> spotigraph.system.v1.Error
	12, // 10: spotigraph.person.v1.UpdateResponse.entity:type_name -> spotigraph.person.v1.Person
	11, // 11: spotigraph.person.v1.DeleteResponse.error:type_name -> spotigraph.system.v1.Error
	11, // 12: spotigraph.person.v1.SearchResponse.error:type_name -> spotigraph.system.v1.Error
	10, // 13: spotigraph.person.v1.SearchResponse.next_cursor:type_name -> google.protobuf.StringValue
	12, // 14: spotigraph.person.v1.SearchResponse.members:type_name -> spotigraph.person.v1.Person
	0,  // 15: spotigraph.person.v1.PersonAPI.Create:input_type -> spotigraph.person.v1.CreateRequest
	1,  // 16: spotigraph.person.v1.PersonAPI.Get:input_type -> spotigraph.person.v1.GetRequest
	2,  // 17: spotigraph.person.v1.PersonAPI.Update:input_type -> spotigraph.person.v1.UpdateRequest
	3,  // 18: spotigraph.person.v1.PersonAPI.Delete:input_type -> spotigraph.person.v1.DeleteRequest
	4,  // 19: spotigraph.person.v1.PersonAPI.Search:input_type -> spotigraph.person.v1.SearchRequest
	5,  // 20: spotigraph.person.v1.PersonAPI.Create:output_type -> spotigraph.person.v1.CreateResponse
	6,  // 21: spotigraph.person.v1.PersonAPI.Get:output_type -> spotigraph.person.v1.GetResponse
	7,  // 22: spotigraph.person.v1.PersonAPI.Update:output_type -> spotigraph.person.v1.UpdateResponse
	8,  // 23: spotigraph.person.v1.PersonAPI.Delete:output_type -> spotigraph.person.v1.DeleteResponse
	9,  // 24: spotigraph.person.v1.PersonAPI.Search:output_type -> spotigraph.person.v1.SearchResponse
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_spotigraph_person_v1_person_api_proto_init() }
func file_spotigraph_person_v1_person_api_proto_init() {
	if File_spotigraph_person_v1_person_api_proto != nil {
		return
	}
	file_spotigraph_person_v1_person_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spotigraph_person_v1_person_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spotigraph_person_v1_person_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spotigraph_person_v1_person_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spotigraph_person_v1_person_api_proto_goTypes,
		DependencyIndexes: file_spotigraph_person_v1_person_api_proto_depIdxs,
		MessageInfos:      file_spotigraph_person_v1_person_api_proto_msgTypes,
	}.Build()
	File_spotigraph_person_v1_person_api_proto = out.File
	file_spotigraph_person_v1_person_api_proto_rawDesc = nil
	file_spotigraph_person_v1_person_api_proto_goTypes = nil
	file_spotigraph_person_v1_person_api_proto_depIdxs = nil
}

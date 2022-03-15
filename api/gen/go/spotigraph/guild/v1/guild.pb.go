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
// source: spotigraph/guild/v1/guild.proto

package guildv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Guild repesents a collection of user of the same expertise.
type Guild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Urn   string `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *Guild) Reset() {
	*x = Guild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_guild_v1_guild_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guild) ProtoMessage() {}

func (x *Guild) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_guild_v1_guild_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guild.ProtoReflect.Descriptor instead.
func (*Guild) Descriptor() ([]byte, []int) {
	return file_spotigraph_guild_v1_guild_proto_rawDescGZIP(), []int{0}
}

func (x *Guild) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Guild) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Guild) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

var File_spotigraph_guild_v1_guild_proto protoreflect.FileDescriptor

var file_spotigraph_guild_v1_guild_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x3f, 0x0a, 0x05, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x42, 0x9d, 0x01, 0x0a, 0x20, 0x6f, 0x72, 0x67, 0x2e,
	0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x61, 0x72, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x7a, 0x6e, 0x74, 0x72,
	0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x47, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x70,
	0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x13, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spotigraph_guild_v1_guild_proto_rawDescOnce sync.Once
	file_spotigraph_guild_v1_guild_proto_rawDescData = file_spotigraph_guild_v1_guild_proto_rawDesc
)

func file_spotigraph_guild_v1_guild_proto_rawDescGZIP() []byte {
	file_spotigraph_guild_v1_guild_proto_rawDescOnce.Do(func() {
		file_spotigraph_guild_v1_guild_proto_rawDescData = protoimpl.X.CompressGZIP(file_spotigraph_guild_v1_guild_proto_rawDescData)
	})
	return file_spotigraph_guild_v1_guild_proto_rawDescData
}

var file_spotigraph_guild_v1_guild_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spotigraph_guild_v1_guild_proto_goTypes = []interface{}{
	(*Guild)(nil), // 0: spotigraph.guild.v1.Guild
}
var file_spotigraph_guild_v1_guild_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spotigraph_guild_v1_guild_proto_init() }
func file_spotigraph_guild_v1_guild_proto_init() {
	if File_spotigraph_guild_v1_guild_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spotigraph_guild_v1_guild_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guild); i {
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
			RawDescriptor: file_spotigraph_guild_v1_guild_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spotigraph_guild_v1_guild_proto_goTypes,
		DependencyIndexes: file_spotigraph_guild_v1_guild_proto_depIdxs,
		MessageInfos:      file_spotigraph_guild_v1_guild_proto_msgTypes,
	}.Build()
	File_spotigraph_guild_v1_guild_proto = out.File
	file_spotigraph_guild_v1_guild_proto_rawDesc = nil
	file_spotigraph_guild_v1_guild_proto_goTypes = nil
	file_spotigraph_guild_v1_guild_proto_depIdxs = nil
}

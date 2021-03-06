// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: readtoken.proto

package readtoken_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	keytransparency_go_proto "github.com/google/keytransparency/core/api/v1/keytransparency_go_proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ReadToken can be serialized and handed to users for pagination.
type ReadToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// slice_index identifies the source for reading.
	SliceIndex int64 `protobuf:"varint,1,opt,name=slice_index,json=sliceIndex,proto3" json:"slice_index,omitempty"`
	// start_watermark identifies the lowest (exclusive) row to return.
	StartWatermark uint64 `protobuf:"varint,4,opt,name=start_watermark,json=startWatermark,proto3" json:"start_watermark,omitempty"`
}

func (x *ReadToken) Reset() {
	*x = ReadToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readtoken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadToken) ProtoMessage() {}

func (x *ReadToken) ProtoReflect() protoreflect.Message {
	mi := &file_readtoken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadToken.ProtoReflect.Descriptor instead.
func (*ReadToken) Descriptor() ([]byte, []int) {
	return file_readtoken_proto_rawDescGZIP(), []int{0}
}

func (x *ReadToken) GetSliceIndex() int64 {
	if x != nil {
		return x.SliceIndex
	}
	return 0
}

func (x *ReadToken) GetStartWatermark() uint64 {
	if x != nil {
		return x.StartWatermark
	}
	return 0
}

// ListUserRevisions token can be serialized and handed to users for pagination
// when listing revisions.
type ListUserRevisionsToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request is the query being paginated over, used for validation of
	// subsequent requests. Fields that are allowed to change between requests
	// (such as page_token or last_verified_tree_size) will not be validated and
	// should be omitted for brevity.
	Request *keytransparency_go_proto.ListUserRevisionsRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// revisions_returned is a running tally of the number of revisions that have
	// been returned across paginated requests in this query.
	RevisionsReturned int64 `protobuf:"varint,2,opt,name=revisions_returned,json=revisionsReturned,proto3" json:"revisions_returned,omitempty"`
}

func (x *ListUserRevisionsToken) Reset() {
	*x = ListUserRevisionsToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_readtoken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRevisionsToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRevisionsToken) ProtoMessage() {}

func (x *ListUserRevisionsToken) ProtoReflect() protoreflect.Message {
	mi := &file_readtoken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRevisionsToken.ProtoReflect.Descriptor instead.
func (*ListUserRevisionsToken) Descriptor() ([]byte, []int) {
	return file_readtoken_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserRevisionsToken) GetRequest() *keytransparency_go_proto.ListUserRevisionsRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ListUserRevisionsToken) GetRevisionsReturned() int64 {
	if x != nil {
		return x.RevisionsReturned
	}
	return 0
}

var File_readtoken_proto protoreflect.FileDescriptor

var file_readtoken_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x61, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x76, 0x31,
	0x2f, 0x6b, 0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6b,
	0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x65, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_readtoken_proto_rawDescOnce sync.Once
	file_readtoken_proto_rawDescData = file_readtoken_proto_rawDesc
)

func file_readtoken_proto_rawDescGZIP() []byte {
	file_readtoken_proto_rawDescOnce.Do(func() {
		file_readtoken_proto_rawDescData = protoimpl.X.CompressGZIP(file_readtoken_proto_rawDescData)
	})
	return file_readtoken_proto_rawDescData
}

var file_readtoken_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_readtoken_proto_goTypes = []interface{}{
	(*ReadToken)(nil),              // 0: google.keytransparency.v1.ReadToken
	(*ListUserRevisionsToken)(nil), // 1: google.keytransparency.v1.ListUserRevisionsToken
	(*keytransparency_go_proto.ListUserRevisionsRequest)(nil), // 2: google.keytransparency.v1.ListUserRevisionsRequest
}
var file_readtoken_proto_depIdxs = []int32{
	2, // 0: google.keytransparency.v1.ListUserRevisionsToken.request:type_name -> google.keytransparency.v1.ListUserRevisionsRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_readtoken_proto_init() }
func file_readtoken_proto_init() {
	if File_readtoken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_readtoken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadToken); i {
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
		file_readtoken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRevisionsToken); i {
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
			RawDescriptor: file_readtoken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_readtoken_proto_goTypes,
		DependencyIndexes: file_readtoken_proto_depIdxs,
		MessageInfos:      file_readtoken_proto_msgTypes,
	}.Build()
	File_readtoken_proto = out.File
	file_readtoken_proto_rawDesc = nil
	file_readtoken_proto_goTypes = nil
	file_readtoken_proto_depIdxs = nil
}

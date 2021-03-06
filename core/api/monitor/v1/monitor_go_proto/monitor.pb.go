// Copyright 2016 Google Inc. All Rights Reserved.
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
// source: monitor/v1/monitor.proto

// Monitor Service
//
// The Key Transparency monitor server service consists of APIs to fetch
// monitor results queried using the mutations API.

package monitor_go_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	trillian "github.com/google/trillian"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// GetStateRequest requests the verification state of a keytransparency
// directory for a particular point in time.
type GetStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kt_url is the URL of the keytransparency server for which the monitoring
	// result will be returned.
	KtUrl string `protobuf:"bytes,2,opt,name=kt_url,json=ktUrl,proto3" json:"kt_url,omitempty"`
	// directory_id identifies the merkle tree being monitored.
	DirectoryId string `protobuf:"bytes,3,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// revision specifies the revision for which the monitoring results will
	// be returned (revisions start at 0).
	Revision int64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *GetStateRequest) Reset() {
	*x = GetStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_v1_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRequest) ProtoMessage() {}

func (x *GetStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_v1_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRequest.ProtoReflect.Descriptor instead.
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return file_monitor_v1_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *GetStateRequest) GetKtUrl() string {
	if x != nil {
		return x.KtUrl
	}
	return ""
}

func (x *GetStateRequest) GetDirectoryId() string {
	if x != nil {
		return x.DirectoryId
	}
	return ""
}

func (x *GetStateRequest) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

// State represents the monitor's evaluation of a Key Transparency directory
// at a particular revision.
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// smr contains the map root for the sparse Merkle Tree signed with the
	// monitor's key on success. If the checks were not successful the
	// smr will be empty. The revisions are encoded into the smr map_revision.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,1,opt,name=smr,proto3" json:"smr,omitempty"`
	// seen_time contains the time when this particular signed map root was
	// retrieved and processed.
	SeenTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=seen_time,json=seenTime,proto3" json:"seen_time,omitempty"`
	// errors contains a list of errors representing the verification checks
	// that failed while monitoring the key-transparency server.
	Errors []*status.Status `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_v1_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_v1_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_monitor_v1_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *State) GetSmr() *trillian.SignedMapRoot {
	if x != nil {
		return x.Smr
	}
	return nil
}

func (x *State) GetSeenTime() *timestamp.Timestamp {
	if x != nil {
		return x.SeenTime
	}
	return nil
}

func (x *State) GetErrors() []*status.Status {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_monitor_v1_monitor_proto protoreflect.FileDescriptor

var file_monitor_v1_monitor_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x97,
	0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x73, 0x6d, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x03,
	0x73, 0x6d, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0x8b, 0x03, 0x0a, 0x07, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x12, 0xb7, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6b,
	0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x12, 0x45, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6b, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x7d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0xc5,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x79, 0x52, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6b,
	0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x12, 0x49, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x6b, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x7d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_v1_monitor_proto_rawDescOnce sync.Once
	file_monitor_v1_monitor_proto_rawDescData = file_monitor_v1_monitor_proto_rawDesc
)

func file_monitor_v1_monitor_proto_rawDescGZIP() []byte {
	file_monitor_v1_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_v1_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_v1_monitor_proto_rawDescData)
	})
	return file_monitor_v1_monitor_proto_rawDescData
}

var file_monitor_v1_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_monitor_v1_monitor_proto_goTypes = []interface{}{
	(*GetStateRequest)(nil),        // 0: google.keytransparency.monitor.v1.GetStateRequest
	(*State)(nil),                  // 1: google.keytransparency.monitor.v1.State
	(*trillian.SignedMapRoot)(nil), // 2: trillian.SignedMapRoot
	(*timestamp.Timestamp)(nil),    // 3: google.protobuf.Timestamp
	(*status.Status)(nil),          // 4: google.rpc.Status
}
var file_monitor_v1_monitor_proto_depIdxs = []int32{
	2, // 0: google.keytransparency.monitor.v1.State.smr:type_name -> trillian.SignedMapRoot
	3, // 1: google.keytransparency.monitor.v1.State.seen_time:type_name -> google.protobuf.Timestamp
	4, // 2: google.keytransparency.monitor.v1.State.errors:type_name -> google.rpc.Status
	0, // 3: google.keytransparency.monitor.v1.Monitor.GetState:input_type -> google.keytransparency.monitor.v1.GetStateRequest
	0, // 4: google.keytransparency.monitor.v1.Monitor.GetStateByRevision:input_type -> google.keytransparency.monitor.v1.GetStateRequest
	1, // 5: google.keytransparency.monitor.v1.Monitor.GetState:output_type -> google.keytransparency.monitor.v1.State
	1, // 6: google.keytransparency.monitor.v1.Monitor.GetStateByRevision:output_type -> google.keytransparency.monitor.v1.State
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_monitor_v1_monitor_proto_init() }
func file_monitor_v1_monitor_proto_init() {
	if File_monitor_v1_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_v1_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateRequest); i {
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
		file_monitor_v1_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
			RawDescriptor: file_monitor_v1_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_v1_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_v1_monitor_proto_depIdxs,
		MessageInfos:      file_monitor_v1_monitor_proto_msgTypes,
	}.Build()
	File_monitor_v1_monitor_proto = out.File
	file_monitor_v1_monitor_proto_rawDesc = nil
	file_monitor_v1_monitor_proto_goTypes = nil
	file_monitor_v1_monitor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest revision the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current revision it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified revision the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current revision it won't sign the map
	// root and additional data will be provided to reproduce the failure.
	GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServer is the server API for Monitor service.
type MonitorServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest revision the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current revision it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(context.Context, *GetStateRequest) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified revision the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current revision it won't sign the map
	// root and additional data will be provided to reproduce the failure.
	GetStateByRevision(context.Context, *GetStateRequest) (*State, error)
}

// UnimplementedMonitorServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (*UnimplementedMonitorServer) GetState(context.Context, *GetStateRequest) (*State, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedMonitorServer) GetStateByRevision(context.Context, *GetStateRequest) (*State, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetStateByRevision not implemented")
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_GetStateByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetStateByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetStateByRevision(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.monitor.v1.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _Monitor_GetState_Handler,
		},
		{
			MethodName: "GetStateByRevision",
			Handler:    _Monitor_GetStateByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor/v1/monitor.proto",
}

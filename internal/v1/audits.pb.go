// Copyright 2020 StrongDM Inc
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
//
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: audits.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// GetAccessRulesRequest specifies which AccessRule to retrieve.
type GetAccessRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Fuzzy timestamp to recover the list of access rule in that point in time.
	When string `protobuf:"bytes,2,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *GetAccessRulesRequest) Reset() {
	*x = GetAccessRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audits_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRulesRequest) ProtoMessage() {}

func (x *GetAccessRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audits_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRulesRequest.ProtoReflect.Descriptor instead.
func (*GetAccessRulesRequest) Descriptor() ([]byte, []int) {
	return file_audits_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccessRulesRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetAccessRulesRequest) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

// GetAccessRulesResponse returns the requested AccessRule in that point in
// time.
type GetAccessRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested AccessRule.
	AccessRules []*AccessRule `protobuf:"bytes,2,rep,name=access_rules,json=accessRules,proto3" json:"access_rules,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *GetAccessRulesResponse) Reset() {
	*x = GetAccessRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audits_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRulesResponse) ProtoMessage() {}

func (x *GetAccessRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audits_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRulesResponse.ProtoReflect.Descriptor instead.
func (*GetAccessRulesResponse) Descriptor() ([]byte, []int) {
	return file_audits_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessRulesResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetAccessRulesResponse) GetAccessRules() []*AccessRule {
	if x != nil {
		return x.AccessRules
	}
	return nil
}

func (x *GetAccessRulesResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

var File_audits_proto protoreflect.FileDescriptor

var file_audits_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x3a, 0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xd0, 0xf3, 0xb3,
	0x07, 0x01, 0x22, 0xf2, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x0a, 0xf2, 0xf8,
	0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x14, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0x90, 0xf4, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x3a, 0x14, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xfa, 0xf8, 0xb3,
	0x07, 0x05, 0xd0, 0xf3, 0xb3, 0x07, 0x01, 0x32, 0x9b, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x73, 0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xf9, 0xb3,
	0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x1b, 0xaa,
	0xf3, 0xb3, 0x07, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x19, 0xca, 0xf9, 0xb3, 0x07,
	0x0a, 0xc2, 0xf9, 0xb3, 0x07, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0xca, 0xf9, 0xb3, 0x07, 0x05,
	0xc8, 0xf9, 0xb3, 0x07, 0x01, 0x42, 0x62, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72,
	0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75,
	0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x50, 0x6c, 0x75,
	0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e,
	0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_audits_proto_rawDescOnce sync.Once
	file_audits_proto_rawDescData = file_audits_proto_rawDesc
)

func file_audits_proto_rawDescGZIP() []byte {
	file_audits_proto_rawDescOnce.Do(func() {
		file_audits_proto_rawDescData = protoimpl.X.CompressGZIP(file_audits_proto_rawDescData)
	})
	return file_audits_proto_rawDescData
}

var file_audits_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_audits_proto_goTypes = []interface{}{
	(*GetAccessRulesRequest)(nil),  // 0: v1.GetAccessRulesRequest
	(*GetAccessRulesResponse)(nil), // 1: v1.GetAccessRulesResponse
	(*GetRequestMetadata)(nil),     // 2: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),    // 3: v1.GetResponseMetadata
	(*AccessRule)(nil),             // 4: v1.AccessRule
	(*RateLimitMetadata)(nil),      // 5: v1.RateLimitMetadata
}
var file_audits_proto_depIdxs = []int32{
	2, // 0: v1.GetAccessRulesRequest.meta:type_name -> v1.GetRequestMetadata
	3, // 1: v1.GetAccessRulesResponse.meta:type_name -> v1.GetResponseMetadata
	4, // 2: v1.GetAccessRulesResponse.access_rules:type_name -> v1.AccessRule
	5, // 3: v1.GetAccessRulesResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0, // 4: v1.Audits.GetAccessRules:input_type -> v1.GetAccessRulesRequest
	1, // 5: v1.Audits.GetAccessRules:output_type -> v1.GetAccessRulesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_audits_proto_init() }
func file_audits_proto_init() {
	if File_audits_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	file_access_rules_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_audits_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRulesRequest); i {
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
		file_audits_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRulesResponse); i {
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
			RawDescriptor: file_audits_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audits_proto_goTypes,
		DependencyIndexes: file_audits_proto_depIdxs,
		MessageInfos:      file_audits_proto_msgTypes,
	}.Build()
	File_audits_proto = out.File
	file_audits_proto_rawDesc = nil
	file_audits_proto_goTypes = nil
	file_audits_proto_depIdxs = nil
}

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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: remote_identity_groups.proto

package v1

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

// RemoteIdentityGroupGetRequest specifies which RemoteIdentityGroup to retrieve.
type RemoteIdentityGroupGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RemoteIdentityGroup to retrieve.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoteIdentityGroupGetRequest) Reset() {
	*x = RemoteIdentityGroupGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_identity_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteIdentityGroupGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteIdentityGroupGetRequest) ProtoMessage() {}

func (x *RemoteIdentityGroupGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_identity_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteIdentityGroupGetRequest.ProtoReflect.Descriptor instead.
func (*RemoteIdentityGroupGetRequest) Descriptor() ([]byte, []int) {
	return file_remote_identity_groups_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteIdentityGroupGetRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RemoteIdentityGroupGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// RemoteIdentityGroupGetResponse returns a requested RemoteIdentityGroup.
type RemoteIdentityGroupGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested RemoteIdentityGroup.
	RemoteIdentityGroup *RemoteIdentityGroup `protobuf:"bytes,2,opt,name=remote_identity_group,json=remoteIdentityGroup,proto3" json:"remote_identity_group,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RemoteIdentityGroupGetResponse) Reset() {
	*x = RemoteIdentityGroupGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_identity_groups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteIdentityGroupGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteIdentityGroupGetResponse) ProtoMessage() {}

func (x *RemoteIdentityGroupGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_identity_groups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteIdentityGroupGetResponse.ProtoReflect.Descriptor instead.
func (*RemoteIdentityGroupGetResponse) Descriptor() ([]byte, []int) {
	return file_remote_identity_groups_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteIdentityGroupGetResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RemoteIdentityGroupGetResponse) GetRemoteIdentityGroup() *RemoteIdentityGroup {
	if x != nil {
		return x.RemoteIdentityGroup
	}
	return nil
}

func (x *RemoteIdentityGroupGetResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// RemoteIdentityGroupListRequest specifies criteria for retrieving a list of RemoteIdentityGroups.
type RemoteIdentityGroupListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *RemoteIdentityGroupListRequest) Reset() {
	*x = RemoteIdentityGroupListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_identity_groups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteIdentityGroupListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteIdentityGroupListRequest) ProtoMessage() {}

func (x *RemoteIdentityGroupListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_identity_groups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteIdentityGroupListRequest.ProtoReflect.Descriptor instead.
func (*RemoteIdentityGroupListRequest) Descriptor() ([]byte, []int) {
	return file_remote_identity_groups_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteIdentityGroupListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RemoteIdentityGroupListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// RemoteIdentityGroupListResponse returns a list of RemoteIdentityGroups that meet the criteria of a
// RemoteIdentityGroupListRequest.
type RemoteIdentityGroupListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	RemoteIdentityGroups []*RemoteIdentityGroup `protobuf:"bytes,2,rep,name=remote_identity_groups,json=remoteIdentityGroups,proto3" json:"remote_identity_groups,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RemoteIdentityGroupListResponse) Reset() {
	*x = RemoteIdentityGroupListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_identity_groups_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteIdentityGroupListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteIdentityGroupListResponse) ProtoMessage() {}

func (x *RemoteIdentityGroupListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_identity_groups_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteIdentityGroupListResponse.ProtoReflect.Descriptor instead.
func (*RemoteIdentityGroupListResponse) Descriptor() ([]byte, []int) {
	return file_remote_identity_groups_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteIdentityGroupListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RemoteIdentityGroupListResponse) GetRemoteIdentityGroups() []*RemoteIdentityGroup {
	if x != nil {
		return x.RemoteIdentityGroups
	}
	return nil
}

func (x *RemoteIdentityGroupListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// A RemoteIdentityGroup defines a group of remote identities.
type RemoteIdentityGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the RemoteIdentityGroup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique human-readable name of the RemoteIdentityGroup.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemoteIdentityGroup) Reset() {
	*x = RemoteIdentityGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_identity_groups_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteIdentityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteIdentityGroup) ProtoMessage() {}

func (x *RemoteIdentityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_remote_identity_groups_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteIdentityGroup.ProtoReflect.Descriptor instead.
func (*RemoteIdentityGroup) Descriptor() ([]byte, []int) {
	return file_remote_identity_groups_proto_rawDescGZIP(), []int{4}
}

func (x *RemoteIdentityGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoteIdentityGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_remote_identity_groups_proto protoreflect.FileDescriptor

var file_remote_identity_groups_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x1d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa2, 0x02, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x57, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x62, 0x0a, 0x0a, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3,
	0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a,
	0x0a, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x71, 0x0a, 0x1e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8e,
	0x02, 0x0a, 0x1f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x59, 0x0a, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x62, 0x0a, 0x0a, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x06, 0xb2, 0xf4, 0xb3, 0x07, 0x01, 0x2a, 0xf2, 0xf8, 0xb3,
	0x07, 0x12, 0xb2, 0xf4, 0xb3, 0x07, 0x0d, 0x21, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x9d, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3,
	0x07, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x45, 0xfa, 0xf8, 0xb3, 0x07, 0x40, 0xa8,
	0xf3, 0xb3, 0x07, 0x01, 0xc2, 0xf3, 0xb3, 0x07, 0x36, 0xaa, 0xf3, 0xb3, 0x07, 0x31, 0x74, 0x66,
	0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x78, 0x74, 0x32,
	0xce, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3,
	0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x24, 0xaa, 0xf3, 0xb3, 0x07, 0x1f,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x82, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x31, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67, 0x65, 0x74,
	0x82, 0xf9, 0xb3, 0x07, 0x1f, 0xaa, 0xf3, 0xb3, 0x07, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x1a, 0x2a, 0xca, 0xf9, 0xb3, 0x07, 0x18, 0xc2, 0xf9, 0xb3, 0x07, 0x13,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0xca, 0xf9, 0xb3, 0x07, 0x08, 0xd2, 0xf9, 0xb3, 0x07, 0x03, 0x69, 0x67, 0x2d,
	0x42, 0x70, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x42, 0x1c, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x76, 0x33, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remote_identity_groups_proto_rawDescOnce sync.Once
	file_remote_identity_groups_proto_rawDescData = file_remote_identity_groups_proto_rawDesc
)

func file_remote_identity_groups_proto_rawDescGZIP() []byte {
	file_remote_identity_groups_proto_rawDescOnce.Do(func() {
		file_remote_identity_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_identity_groups_proto_rawDescData)
	})
	return file_remote_identity_groups_proto_rawDescData
}

var file_remote_identity_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_remote_identity_groups_proto_goTypes = []interface{}{
	(*RemoteIdentityGroupGetRequest)(nil),   // 0: v1.RemoteIdentityGroupGetRequest
	(*RemoteIdentityGroupGetResponse)(nil),  // 1: v1.RemoteIdentityGroupGetResponse
	(*RemoteIdentityGroupListRequest)(nil),  // 2: v1.RemoteIdentityGroupListRequest
	(*RemoteIdentityGroupListResponse)(nil), // 3: v1.RemoteIdentityGroupListResponse
	(*RemoteIdentityGroup)(nil),             // 4: v1.RemoteIdentityGroup
	(*GetRequestMetadata)(nil),              // 5: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),             // 6: v1.GetResponseMetadata
	(*RateLimitMetadata)(nil),               // 7: v1.RateLimitMetadata
	(*ListRequestMetadata)(nil),             // 8: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),            // 9: v1.ListResponseMetadata
}
var file_remote_identity_groups_proto_depIdxs = []int32{
	5,  // 0: v1.RemoteIdentityGroupGetRequest.meta:type_name -> v1.GetRequestMetadata
	6,  // 1: v1.RemoteIdentityGroupGetResponse.meta:type_name -> v1.GetResponseMetadata
	4,  // 2: v1.RemoteIdentityGroupGetResponse.remote_identity_group:type_name -> v1.RemoteIdentityGroup
	7,  // 3: v1.RemoteIdentityGroupGetResponse.rate_limit:type_name -> v1.RateLimitMetadata
	8,  // 4: v1.RemoteIdentityGroupListRequest.meta:type_name -> v1.ListRequestMetadata
	9,  // 5: v1.RemoteIdentityGroupListResponse.meta:type_name -> v1.ListResponseMetadata
	4,  // 6: v1.RemoteIdentityGroupListResponse.remote_identity_groups:type_name -> v1.RemoteIdentityGroup
	7,  // 7: v1.RemoteIdentityGroupListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0,  // 8: v1.RemoteIdentityGroups.Get:input_type -> v1.RemoteIdentityGroupGetRequest
	2,  // 9: v1.RemoteIdentityGroups.List:input_type -> v1.RemoteIdentityGroupListRequest
	1,  // 10: v1.RemoteIdentityGroups.Get:output_type -> v1.RemoteIdentityGroupGetResponse
	3,  // 11: v1.RemoteIdentityGroups.List:output_type -> v1.RemoteIdentityGroupListResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_remote_identity_groups_proto_init() }
func file_remote_identity_groups_proto_init() {
	if File_remote_identity_groups_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remote_identity_groups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteIdentityGroupGetRequest); i {
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
		file_remote_identity_groups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteIdentityGroupGetResponse); i {
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
		file_remote_identity_groups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteIdentityGroupListRequest); i {
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
		file_remote_identity_groups_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteIdentityGroupListResponse); i {
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
		file_remote_identity_groups_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteIdentityGroup); i {
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
			RawDescriptor: file_remote_identity_groups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_identity_groups_proto_goTypes,
		DependencyIndexes: file_remote_identity_groups_proto_depIdxs,
		MessageInfos:      file_remote_identity_groups_proto_msgTypes,
	}.Build()
	File_remote_identity_groups_proto = out.File
	file_remote_identity_groups_proto_rawDesc = nil
	file_remote_identity_groups_proto_goTypes = nil
	file_remote_identity_groups_proto_depIdxs = nil
}

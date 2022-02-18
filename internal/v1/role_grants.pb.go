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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: role_grants.proto

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

// RoleGrantCreateRequest specifies what kind of RoleGrants should be registered in
// the organizations fleet.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new RoleGrant.
	RoleGrant *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
}

func (x *RoleGrantCreateRequest) Reset() {
	*x = RoleGrantCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantCreateRequest) ProtoMessage() {}

func (x *RoleGrantCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantCreateRequest.ProtoReflect.Descriptor instead.
func (*RoleGrantCreateRequest) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{0}
}

func (x *RoleGrantCreateRequest) GetMeta() *CreateRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantCreateRequest) GetRoleGrant() *RoleGrant {
	if x != nil {
		return x.RoleGrant
	}
	return nil
}

// RoleGrantCreateResponse reports how the RoleGrants were created in the system.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created RoleGrant.
	RoleGrant *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RoleGrantCreateResponse) Reset() {
	*x = RoleGrantCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantCreateResponse) ProtoMessage() {}

func (x *RoleGrantCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantCreateResponse.ProtoReflect.Descriptor instead.
func (*RoleGrantCreateResponse) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{1}
}

func (x *RoleGrantCreateResponse) GetMeta() *CreateResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantCreateResponse) GetRoleGrant() *RoleGrant {
	if x != nil {
		return x.RoleGrant
	}
	return nil
}

func (x *RoleGrantCreateResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// RoleGrantGetRequest specifies which RoleGrant to retrieve.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleGrant to retrieve.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleGrantGetRequest) Reset() {
	*x = RoleGrantGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantGetRequest) ProtoMessage() {}

func (x *RoleGrantGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantGetRequest.ProtoReflect.Descriptor instead.
func (*RoleGrantGetRequest) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{2}
}

func (x *RoleGrantGetRequest) GetMeta() *GetRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// RoleGrantGetResponse returns a requested RoleGrant.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested RoleGrant.
	RoleGrant *RoleGrant `protobuf:"bytes,2,opt,name=role_grant,json=roleGrant,proto3" json:"role_grant,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RoleGrantGetResponse) Reset() {
	*x = RoleGrantGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantGetResponse) ProtoMessage() {}

func (x *RoleGrantGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantGetResponse.ProtoReflect.Descriptor instead.
func (*RoleGrantGetResponse) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{3}
}

func (x *RoleGrantGetResponse) GetMeta() *GetResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantGetResponse) GetRoleGrant() *RoleGrant {
	if x != nil {
		return x.RoleGrant
	}
	return nil
}

func (x *RoleGrantGetResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// RoleGrantDeleteRequest identifies a RoleGrant by ID to delete.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the RoleGrant to delete.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RoleGrantDeleteRequest) Reset() {
	*x = RoleGrantDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantDeleteRequest) ProtoMessage() {}

func (x *RoleGrantDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantDeleteRequest.ProtoReflect.Descriptor instead.
func (*RoleGrantDeleteRequest) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{4}
}

func (x *RoleGrantDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// RoleGrantDeleteResponse returns information about a RoleGrant that was deleted.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RoleGrantDeleteResponse) Reset() {
	*x = RoleGrantDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantDeleteResponse) ProtoMessage() {}

func (x *RoleGrantDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantDeleteResponse.ProtoReflect.Descriptor instead.
func (*RoleGrantDeleteResponse) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{5}
}

func (x *RoleGrantDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// RoleGrantListRequest specifies criteria for retrieving a list of RoleGrants.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *RoleGrantListRequest) Reset() {
	*x = RoleGrantListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantListRequest) ProtoMessage() {}

func (x *RoleGrantListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantListRequest.ProtoReflect.Descriptor instead.
func (*RoleGrantListRequest) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{6}
}

func (x *RoleGrantListRequest) GetMeta() *ListRequestMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// RoleGrantListResponse returns a list of RoleGrants that meet the criteria of a
// RoleGrantListRequest.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrantListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	RoleGrants []*RoleGrant `protobuf:"bytes,2,rep,name=role_grants,json=roleGrants,proto3" json:"role_grants,omitempty"`
	// Rate limit information.
	RateLimit *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
}

func (x *RoleGrantListResponse) Reset() {
	*x = RoleGrantListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrantListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrantListResponse) ProtoMessage() {}

func (x *RoleGrantListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrantListResponse.ProtoReflect.Descriptor instead.
func (*RoleGrantListResponse) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{7}
}

func (x *RoleGrantListResponse) GetMeta() *ListResponseMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RoleGrantListResponse) GetRoleGrants() []*RoleGrant {
	if x != nil {
		return x.RoleGrants
	}
	return nil
}

func (x *RoleGrantListResponse) GetRateLimit() *RateLimitMetadata {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

// A RoleGrant connects a resource to a role, granting members of the role access to that resource.
//
// Deprecated: use access rules instead.
//
// Deprecated: Do not use.
type RoleGrant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the RoleGrant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the resource of this RoleGrant.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The id of the attached role of this RoleGrant.
	RoleId string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RoleGrant) Reset() {
	*x = RoleGrant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_grants_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrant) ProtoMessage() {}

func (x *RoleGrant) ProtoReflect() protoreflect.Message {
	mi := &file_role_grants_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrant.ProtoReflect.Descriptor instead.
func (*RoleGrant) Descriptor() ([]byte, []int) {
	return file_role_grants_proto_rawDescGZIP(), []int{8}
}

func (x *RoleGrant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleGrant) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *RoleGrant) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_role_grants_proto protoreflect.FileDescriptor

var file_role_grants_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0a,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72, 0x6f, 0x6c,
	0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x52,
	0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0a,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x14, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3,
	0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0x90, 0xf4, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0c, 0x18, 0x01, 0xfa, 0xf8, 0xb3, 0x07,
	0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x61, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xe3, 0x01, 0x0a, 0x14, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05,
	0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x0a, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x14, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07,
	0x05, 0x90, 0xf4, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x3a, 0x0c, 0x18, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22,
	0x67, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xaf, 0x01, 0x0a, 0x17, 0x52, 0x6f, 0x6c,
	0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a,
	0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x4a, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x14, 0xf2, 0xf8, 0xb3,
	0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0x90, 0xf4, 0xb3, 0x07,
	0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x0c, 0x18, 0x01,
	0xfa, 0xf8, 0xb3, 0x07, 0x05, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0x22, 0x6b, 0x0a, 0x14, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xd1, 0x01, 0x0a, 0x15, 0x52, 0x6f, 0x6c, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x3a, 0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb8, 0xf3, 0xb3, 0x07, 0x01, 0x52,
	0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x14, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3,
	0x07, 0x01, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0x90, 0xf4, 0xb3, 0x07, 0x01, 0x52, 0x09, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xe9, 0x01, 0x0a, 0x09,
	0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xf2, 0xf8, 0xb3, 0x07, 0x05, 0xb0, 0xf3, 0xb3, 0x07,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07,
	0x0a, 0xb0, 0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xf8, 0xb3, 0x07, 0x0a, 0xb0,
	0xf3, 0xb3, 0x07, 0x01, 0xc0, 0xf3, 0xb3, 0x07, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x3a, 0x64, 0x18, 0x01, 0xfa, 0xf8, 0xb3, 0x07, 0x5d, 0xa8, 0xf3, 0xb3, 0x07, 0x01, 0xc2,
	0xf3, 0xb3, 0x07, 0x53, 0xa2, 0xf3, 0xb3, 0x07, 0x23, 0x74, 0x66, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x74, 0x78, 0x74, 0xaa, 0xf3, 0xb3, 0x07,
	0x26, 0x74, 0x66, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x74, 0x78, 0x74, 0x32, 0xba, 0x04, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x88, 0x02, 0x01, 0x82,
	0xf9, 0xb3, 0x07, 0x09, 0xa2, 0xf3, 0xb3, 0x07, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x82, 0xf9, 0xb3,
	0x07, 0x14, 0xaa, 0xf3, 0xb3, 0x07, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x82, 0xf9, 0xb3, 0x07, 0x0f, 0xb2, 0xf3, 0xb3, 0x07, 0x0a,
	0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x33, 0x2d, 0x33, 0x31, 0x12, 0x7c, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x88, 0x02, 0x01, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3,
	0xb3, 0x07, 0x03, 0x67, 0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x19, 0xaa, 0xf3, 0xb3, 0x07, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x82, 0xf9, 0xb3, 0x07, 0x0f, 0xb2, 0xf3, 0xb3, 0x07, 0x0a, 0x32, 0x30,
	0x32, 0x32, 0x2d, 0x30, 0x33, 0x2d, 0x33, 0x31, 0x12, 0x88, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x88, 0x02,
	0x01, 0x82, 0xf9, 0xb3, 0x07, 0x0b, 0xa2, 0xf3, 0xb3, 0x07, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x82, 0xf9, 0xb3, 0x07, 0x19, 0xaa, 0xf3, 0xb3, 0x07, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x82,
	0xf9, 0xb3, 0x07, 0x0f, 0xb2, 0xf3, 0xb3, 0x07, 0x0a, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x33,
	0x2d, 0x33, 0x31, 0x12, 0x7a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3d, 0x88, 0x02, 0x01, 0x82, 0xf9, 0xb3, 0x07, 0x08, 0xa2, 0xf3, 0xb3, 0x07, 0x03, 0x67,
	0x65, 0x74, 0x82, 0xf9, 0xb3, 0x07, 0x14, 0xaa, 0xf3, 0xb3, 0x07, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2d, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x82, 0xf9, 0xb3, 0x07, 0x0f,
	0xb2, 0xf3, 0xb3, 0x07, 0x0a, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x33, 0x2d, 0x33, 0x31, 0x1a,
	0x23, 0x88, 0x02, 0x01, 0xca, 0xf9, 0xb3, 0x07, 0x0e, 0xc2, 0xf9, 0xb3, 0x07, 0x09, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0xca, 0xf9, 0xb3, 0x07, 0x08, 0xd2, 0xf9, 0xb3, 0x07,
	0x03, 0x72, 0x67, 0x2d, 0x42, 0x69, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x6d,
	0x62, 0x69, 0x6e, 0x67, 0x42, 0x12, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x50, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_grants_proto_rawDescOnce sync.Once
	file_role_grants_proto_rawDescData = file_role_grants_proto_rawDesc
)

func file_role_grants_proto_rawDescGZIP() []byte {
	file_role_grants_proto_rawDescOnce.Do(func() {
		file_role_grants_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_grants_proto_rawDescData)
	})
	return file_role_grants_proto_rawDescData
}

var file_role_grants_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_role_grants_proto_goTypes = []interface{}{
	(*RoleGrantCreateRequest)(nil),  // 0: v1.RoleGrantCreateRequest
	(*RoleGrantCreateResponse)(nil), // 1: v1.RoleGrantCreateResponse
	(*RoleGrantGetRequest)(nil),     // 2: v1.RoleGrantGetRequest
	(*RoleGrantGetResponse)(nil),    // 3: v1.RoleGrantGetResponse
	(*RoleGrantDeleteRequest)(nil),  // 4: v1.RoleGrantDeleteRequest
	(*RoleGrantDeleteResponse)(nil), // 5: v1.RoleGrantDeleteResponse
	(*RoleGrantListRequest)(nil),    // 6: v1.RoleGrantListRequest
	(*RoleGrantListResponse)(nil),   // 7: v1.RoleGrantListResponse
	(*RoleGrant)(nil),               // 8: v1.RoleGrant
	(*CreateRequestMetadata)(nil),   // 9: v1.CreateRequestMetadata
	(*CreateResponseMetadata)(nil),  // 10: v1.CreateResponseMetadata
	(*RateLimitMetadata)(nil),       // 11: v1.RateLimitMetadata
	(*GetRequestMetadata)(nil),      // 12: v1.GetRequestMetadata
	(*GetResponseMetadata)(nil),     // 13: v1.GetResponseMetadata
	(*DeleteRequestMetadata)(nil),   // 14: v1.DeleteRequestMetadata
	(*DeleteResponseMetadata)(nil),  // 15: v1.DeleteResponseMetadata
	(*ListRequestMetadata)(nil),     // 16: v1.ListRequestMetadata
	(*ListResponseMetadata)(nil),    // 17: v1.ListResponseMetadata
}
var file_role_grants_proto_depIdxs = []int32{
	9,  // 0: v1.RoleGrantCreateRequest.meta:type_name -> v1.CreateRequestMetadata
	8,  // 1: v1.RoleGrantCreateRequest.role_grant:type_name -> v1.RoleGrant
	10, // 2: v1.RoleGrantCreateResponse.meta:type_name -> v1.CreateResponseMetadata
	8,  // 3: v1.RoleGrantCreateResponse.role_grant:type_name -> v1.RoleGrant
	11, // 4: v1.RoleGrantCreateResponse.rate_limit:type_name -> v1.RateLimitMetadata
	12, // 5: v1.RoleGrantGetRequest.meta:type_name -> v1.GetRequestMetadata
	13, // 6: v1.RoleGrantGetResponse.meta:type_name -> v1.GetResponseMetadata
	8,  // 7: v1.RoleGrantGetResponse.role_grant:type_name -> v1.RoleGrant
	11, // 8: v1.RoleGrantGetResponse.rate_limit:type_name -> v1.RateLimitMetadata
	14, // 9: v1.RoleGrantDeleteRequest.meta:type_name -> v1.DeleteRequestMetadata
	15, // 10: v1.RoleGrantDeleteResponse.meta:type_name -> v1.DeleteResponseMetadata
	11, // 11: v1.RoleGrantDeleteResponse.rate_limit:type_name -> v1.RateLimitMetadata
	16, // 12: v1.RoleGrantListRequest.meta:type_name -> v1.ListRequestMetadata
	17, // 13: v1.RoleGrantListResponse.meta:type_name -> v1.ListResponseMetadata
	8,  // 14: v1.RoleGrantListResponse.role_grants:type_name -> v1.RoleGrant
	11, // 15: v1.RoleGrantListResponse.rate_limit:type_name -> v1.RateLimitMetadata
	0,  // 16: v1.RoleGrants.Create:input_type -> v1.RoleGrantCreateRequest
	2,  // 17: v1.RoleGrants.Get:input_type -> v1.RoleGrantGetRequest
	4,  // 18: v1.RoleGrants.Delete:input_type -> v1.RoleGrantDeleteRequest
	6,  // 19: v1.RoleGrants.List:input_type -> v1.RoleGrantListRequest
	1,  // 20: v1.RoleGrants.Create:output_type -> v1.RoleGrantCreateResponse
	3,  // 21: v1.RoleGrants.Get:output_type -> v1.RoleGrantGetResponse
	5,  // 22: v1.RoleGrants.Delete:output_type -> v1.RoleGrantDeleteResponse
	7,  // 23: v1.RoleGrants.List:output_type -> v1.RoleGrantListResponse
	20, // [20:24] is the sub-list for method output_type
	16, // [16:20] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_role_grants_proto_init() }
func file_role_grants_proto_init() {
	if File_role_grants_proto != nil {
		return
	}
	file_options_proto_init()
	file_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_grants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantCreateRequest); i {
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
		file_role_grants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantCreateResponse); i {
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
		file_role_grants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantGetRequest); i {
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
		file_role_grants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantGetResponse); i {
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
		file_role_grants_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantDeleteRequest); i {
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
		file_role_grants_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantDeleteResponse); i {
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
		file_role_grants_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantListRequest); i {
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
		file_role_grants_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrantListResponse); i {
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
		file_role_grants_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrant); i {
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
			RawDescriptor: file_role_grants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_grants_proto_goTypes,
		DependencyIndexes: file_role_grants_proto_depIdxs,
		MessageInfos:      file_role_grants_proto_msgTypes,
	}.Build()
	File_role_grants_proto = out.File
	file_role_grants_proto_rawDesc = nil
	file_role_grants_proto_goTypes = nil
	file_role_grants_proto_depIdxs = nil
}

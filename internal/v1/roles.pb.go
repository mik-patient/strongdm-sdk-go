/*
Copyright 2020 StrongDM Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roles.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RoleCreateRequest specifies what kind of Roles that should be registered in
// the organizations fleet.
type RoleCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new Role.
	Role                 *Role    `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleCreateRequest) Reset()         { *m = RoleCreateRequest{} }
func (m *RoleCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RoleCreateRequest) ProtoMessage()    {}
func (*RoleCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{0}
}

func (m *RoleCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleCreateRequest.Unmarshal(m, b)
}
func (m *RoleCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleCreateRequest.Marshal(b, m, deterministic)
}
func (m *RoleCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleCreateRequest.Merge(m, src)
}
func (m *RoleCreateRequest) XXX_Size() int {
	return xxx_messageInfo_RoleCreateRequest.Size(m)
}
func (m *RoleCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleCreateRequest proto.InternalMessageInfo

func (m *RoleCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleCreateRequest) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

// RoleCreateResponse reports how the Roles were created in the system. It can
// communicate partial successes or failures.
type RoleCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created Role.
	Role *Role `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleCreateResponse) Reset()         { *m = RoleCreateResponse{} }
func (m *RoleCreateResponse) String() string { return proto.CompactTextString(m) }
func (*RoleCreateResponse) ProtoMessage()    {}
func (*RoleCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{1}
}

func (m *RoleCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleCreateResponse.Unmarshal(m, b)
}
func (m *RoleCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleCreateResponse.Marshal(b, m, deterministic)
}
func (m *RoleCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleCreateResponse.Merge(m, src)
}
func (m *RoleCreateResponse) XXX_Size() int {
	return xxx_messageInfo_RoleCreateResponse.Size(m)
}
func (m *RoleCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleCreateResponse proto.InternalMessageInfo

func (m *RoleCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleCreateResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleGetRequest specifies which Role to retrieve.
type RoleGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the Role to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleGetRequest) Reset()         { *m = RoleGetRequest{} }
func (m *RoleGetRequest) String() string { return proto.CompactTextString(m) }
func (*RoleGetRequest) ProtoMessage()    {}
func (*RoleGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{2}
}

func (m *RoleGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGetRequest.Unmarshal(m, b)
}
func (m *RoleGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGetRequest.Marshal(b, m, deterministic)
}
func (m *RoleGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGetRequest.Merge(m, src)
}
func (m *RoleGetRequest) XXX_Size() int {
	return xxx_messageInfo_RoleGetRequest.Size(m)
}
func (m *RoleGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGetRequest proto.InternalMessageInfo

func (m *RoleGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleGetResponse returns a requested Role.
type RoleGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested Role.
	Role *Role `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleGetResponse) Reset()         { *m = RoleGetResponse{} }
func (m *RoleGetResponse) String() string { return proto.CompactTextString(m) }
func (*RoleGetResponse) ProtoMessage()    {}
func (*RoleGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{3}
}

func (m *RoleGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleGetResponse.Unmarshal(m, b)
}
func (m *RoleGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleGetResponse.Marshal(b, m, deterministic)
}
func (m *RoleGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleGetResponse.Merge(m, src)
}
func (m *RoleGetResponse) XXX_Size() int {
	return xxx_messageInfo_RoleGetResponse.Size(m)
}
func (m *RoleGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleGetResponse proto.InternalMessageInfo

func (m *RoleGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleGetResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleUpdateRequest identifies a Role by ID and provides fields to update on
// that Role record.
type RoleUpdateRequest struct {
	// Reserved for future use.
	Meta *UpdateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the Role to update. If an ID is already
	// specified in the `role` field, this field is not required. If an ID is
	// specified in both places, they must match.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Parameters to overwrite the specified Role.
	Role                 *Role    `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleUpdateRequest) Reset()         { *m = RoleUpdateRequest{} }
func (m *RoleUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateRequest) ProtoMessage()    {}
func (*RoleUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{4}
}

func (m *RoleUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleUpdateRequest.Unmarshal(m, b)
}
func (m *RoleUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleUpdateRequest.Marshal(b, m, deterministic)
}
func (m *RoleUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleUpdateRequest.Merge(m, src)
}
func (m *RoleUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_RoleUpdateRequest.Size(m)
}
func (m *RoleUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleUpdateRequest proto.InternalMessageInfo

func (m *RoleUpdateRequest) GetMeta() *UpdateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleUpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RoleUpdateRequest) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

// RoleUpdateResponse returns the fields of a Role after it has been updated by
// a RoleUpdateRequest.
type RoleUpdateResponse struct {
	// Reserved for future use.
	Meta *UpdateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The updated Role.
	Role *Role `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleUpdateResponse) Reset()         { *m = RoleUpdateResponse{} }
func (m *RoleUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*RoleUpdateResponse) ProtoMessage()    {}
func (*RoleUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{5}
}

func (m *RoleUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleUpdateResponse.Unmarshal(m, b)
}
func (m *RoleUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleUpdateResponse.Marshal(b, m, deterministic)
}
func (m *RoleUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleUpdateResponse.Merge(m, src)
}
func (m *RoleUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_RoleUpdateResponse.Size(m)
}
func (m *RoleUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleUpdateResponse proto.InternalMessageInfo

func (m *RoleUpdateResponse) GetMeta() *UpdateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleUpdateResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleUpdateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleDeleteRequest identifies a Role by ID to delete.
type RoleDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the Role to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleDeleteRequest) Reset()         { *m = RoleDeleteRequest{} }
func (m *RoleDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteRequest) ProtoMessage()    {}
func (*RoleDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{6}
}

func (m *RoleDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleDeleteRequest.Unmarshal(m, b)
}
func (m *RoleDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleDeleteRequest.Marshal(b, m, deterministic)
}
func (m *RoleDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleDeleteRequest.Merge(m, src)
}
func (m *RoleDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_RoleDeleteRequest.Size(m)
}
func (m *RoleDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleDeleteRequest proto.InternalMessageInfo

func (m *RoleDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// RoleDeleteResponse returns information about a Role that was deleted.
type RoleDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleDeleteResponse) Reset()         { *m = RoleDeleteResponse{} }
func (m *RoleDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*RoleDeleteResponse) ProtoMessage()    {}
func (*RoleDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{7}
}

func (m *RoleDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleDeleteResponse.Unmarshal(m, b)
}
func (m *RoleDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleDeleteResponse.Marshal(b, m, deterministic)
}
func (m *RoleDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleDeleteResponse.Merge(m, src)
}
func (m *RoleDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_RoleDeleteResponse.Size(m)
}
func (m *RoleDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleDeleteResponse proto.InternalMessageInfo

func (m *RoleDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// RoleListRequest specifies criteria for retrieving a list of Roles.
type RoleListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleListRequest) Reset()         { *m = RoleListRequest{} }
func (m *RoleListRequest) String() string { return proto.CompactTextString(m) }
func (*RoleListRequest) ProtoMessage()    {}
func (*RoleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{8}
}

func (m *RoleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleListRequest.Unmarshal(m, b)
}
func (m *RoleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleListRequest.Marshal(b, m, deterministic)
}
func (m *RoleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleListRequest.Merge(m, src)
}
func (m *RoleListRequest) XXX_Size() int {
	return xxx_messageInfo_RoleListRequest.Size(m)
}
func (m *RoleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleListRequest proto.InternalMessageInfo

func (m *RoleListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// RoleListResponse returns a list of Roles that meet the criteria of a
// RoleListRequest.
type RoleListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	Roles []*Role `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RoleListResponse) Reset()         { *m = RoleListResponse{} }
func (m *RoleListResponse) String() string { return proto.CompactTextString(m) }
func (*RoleListResponse) ProtoMessage()    {}
func (*RoleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{9}
}

func (m *RoleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleListResponse.Unmarshal(m, b)
}
func (m *RoleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleListResponse.Marshal(b, m, deterministic)
}
func (m *RoleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleListResponse.Merge(m, src)
}
func (m *RoleListResponse) XXX_Size() int {
	return xxx_messageInfo_RoleListResponse.Size(m)
}
func (m *RoleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoleListResponse proto.InternalMessageInfo

func (m *RoleListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *RoleListResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *RoleListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// A Role grants users access to a set of resources. Composite roles have no
// resource associations of their own, but instead grant access to the combined
// resources of their child roles.
type Role struct {
	// Unique identifier of the Role.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique human-readable name of the Role.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// True if the Role is a composite role.
	Composite            bool     `protobuf:"varint,3,opt,name=composite,proto3" json:"composite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_b96358c61fe6d5ae, []int{10}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetComposite() bool {
	if m != nil {
		return m.Composite
	}
	return false
}

func init() {
	proto.RegisterType((*RoleCreateRequest)(nil), "v1.RoleCreateRequest")
	proto.RegisterType((*RoleCreateResponse)(nil), "v1.RoleCreateResponse")
	proto.RegisterType((*RoleGetRequest)(nil), "v1.RoleGetRequest")
	proto.RegisterType((*RoleGetResponse)(nil), "v1.RoleGetResponse")
	proto.RegisterType((*RoleUpdateRequest)(nil), "v1.RoleUpdateRequest")
	proto.RegisterType((*RoleUpdateResponse)(nil), "v1.RoleUpdateResponse")
	proto.RegisterType((*RoleDeleteRequest)(nil), "v1.RoleDeleteRequest")
	proto.RegisterType((*RoleDeleteResponse)(nil), "v1.RoleDeleteResponse")
	proto.RegisterType((*RoleListRequest)(nil), "v1.RoleListRequest")
	proto.RegisterType((*RoleListResponse)(nil), "v1.RoleListResponse")
	proto.RegisterType((*Role)(nil), "v1.Role")
}

func init() { proto.RegisterFile("roles.proto", fileDescriptor_b96358c61fe6d5ae) }

var fileDescriptor_b96358c61fe6d5ae = []byte{
	// 842 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xbf, 0x4f, 0x23, 0x47,
	0x14, 0xd6, 0xae, 0x8d, 0x8d, 0x07, 0xe1, 0xc0, 0x38, 0x10, 0xb3, 0xa2, 0x58, 0x0d, 0x91, 0x82,
	0x1c, 0xec, 0x15, 0x44, 0x0a, 0x92, 0xab, 0x60, 0x50, 0x50, 0x24, 0xf2, 0x43, 0x8e, 0x90, 0x52,
	0x25, 0x1a, 0x7b, 0x27, 0xcb, 0x28, 0xbb, 0x3b, 0x9b, 0xdd, 0xc1, 0x4e, 0x44, 0x68, 0xd2, 0xa7,
	0x21, 0x7d, 0x8a, 0x48, 0x27, 0x5d, 0x75, 0x3a, 0xc9, 0xcd, 0xd5, 0x54, 0x27, 0xda, 0xfb, 0x17,
	0xee, 0x2f, 0x80, 0xe6, 0xa0, 0x3a, 0xcd, 0xec, 0x2f, 0x8f, 0xcf, 0x70, 0x87, 0x68, 0xa8, 0x60,
	0x67, 0xbe, 0xf9, 0xbe, 0xf7, 0xbe, 0xf7, 0xde, 0x8c, 0xc1, 0x5c, 0xc8, 0x5c, 0x12, 0xb5, 0x82,
	0x90, 0x71, 0x06, 0xf5, 0xc1, 0xa6, 0xb1, 0xea, 0x30, 0xe6, 0xb8, 0xc4, 0xc2, 0x01, 0xb5, 0xb0,
	0xef, 0x33, 0x8e, 0x39, 0x65, 0x7e, 0x82, 0x30, 0x36, 0xe4, 0x9f, 0x7e, 0xd3, 0x21, 0x7e, 0x33,
	0x1a, 0x62, 0xc7, 0x21, 0xa1, 0xc5, 0x02, 0x89, 0x98, 0x82, 0x9e, 0x4f, 0xb6, 0x92, 0x4f, 0x10,
	0x05, 0xa4, 0x1f, 0xff, 0x8f, 0x5c, 0xb0, 0xd8, 0x65, 0x2e, 0xd9, 0x0d, 0x09, 0xe6, 0xa4, 0x4b,
	0x7e, 0x3f, 0x26, 0x11, 0x87, 0x4d, 0x50, 0xf4, 0x08, 0xc7, 0x75, 0xcd, 0xd4, 0xd6, 0xe7, 0xb6,
	0x56, 0x5a, 0x83, 0xcd, 0x96, 0x02, 0xf8, 0x96, 0x70, 0x6c, 0x63, 0x8e, 0xbb, 0x12, 0x06, 0xd7,
	0x41, 0x51, 0x44, 0x5f, 0xd7, 0x25, 0x7c, 0x56, 0xc0, 0x05, 0x67, 0x07, 0x5c, 0xbe, 0x19, 0x95,
	0x67, 0x9e, 0x5f, 0x8d, 0xca, 0x5a, 0x57, 0x22, 0xd0, 0x4b, 0x0d, 0xc0, 0x71, 0xb9, 0x28, 0x60,
	0x7e, 0x44, 0x60, 0x5b, 0xd1, 0x33, 0xc6, 0xf5, 0x62, 0x44, 0x2a, 0xa8, 0x52, 0xde, 0x4f, 0x1c,
	0x7e, 0x05, 0x40, 0x88, 0x39, 0xf9, 0xc5, 0xa5, 0x1e, 0xe5, 0xf5, 0x82, 0xc4, 0x2f, 0x49, 0x3c,
	0xe6, 0xe4, 0x40, 0x2c, 0x4e, 0x95, 0xa9, 0x84, 0xe9, 0x76, 0x1b, 0xdc, 0x88, 0xe5, 0xa7, 0x62,
	0x19, 0xfd, 0x04, 0xaa, 0x42, 0x67, 0x9f, 0xf0, 0xd4, 0xb5, 0x86, 0x92, 0xc5, 0xb2, 0x60, 0xce,
	0x77, 0x27, 0x2c, 0x33, 0x80, 0x4e, 0x6d, 0x19, 0x73, 0x45, 0x11, 0xd3, 0xa9, 0x8d, 0xce, 0x35,
	0xf0, 0x51, 0x46, 0x9d, 0x38, 0xb4, 0xad, 0x70, 0x7f, 0x92, 0x71, 0x3f, 0x5a, 0x7b, 0xfe, 0x8a,
	0xfb, 0xea, 0x30, 0xb0, 0xef, 0xee, 0x2b, 0x05, 0x30, 0x61, 0x52, 0x35, 0x37, 0x49, 0x18, 0x93,
	0xe5, 0x52, 0xf8, 0xe0, 0x3e, 0x4b, 0xd9, 0x6f, 0xef, 0x33, 0x15, 0xf1, 0x08, 0x8d, 0xfc, 0x39,
	0x36, 0x72, 0x8f, 0xb8, 0xe4, 0x4e, 0x23, 0x15, 0xc0, 0x3d, 0xba, 0xed, 0xbf, 0xc4, 0xaa, 0xf4,
	0xfc, 0xed, 0x56, 0xa9, 0x88, 0x3b, 0xac, 0x52, 0x0d, 0xd0, 0x1f, 0x68, 0x40, 0x2f, 0x9e, 0x86,
	0x03, 0x1a, 0x65, 0x93, 0xf6, 0xf9, 0xbb, 0xd3, 0x30, 0xb6, 0x3d, 0x91, 0x3c, 0x02, 0xa5, 0x5f,
	0xa9, 0xcb, 0x49, 0x38, 0xc5, 0x80, 0x64, 0x07, 0x3d, 0xd3, 0xc0, 0x42, 0x2e, 0x92, 0x58, 0xb0,
	0xa1, 0xa8, 0xd4, 0x73, 0x15, 0xd5, 0x80, 0x44, 0xa6, 0x01, 0x66, 0xe4, 0x15, 0x5e, 0xd7, 0xcd,
	0xc2, 0x94, 0x06, 0x79, 0x21, 0xf5, 0x62, 0xc8, 0xc3, 0x3b, 0x04, 0xfd, 0xa3, 0x83, 0xa2, 0x60,
	0x87, 0x0d, 0x59, 0x5a, 0x4d, 0x66, 0x66, 0x08, 0xec, 0xd2, 0xff, 0x57, 0xa3, 0xb2, 0xfe, 0xcd,
	0x9e, 0x3c, 0x72, 0x71, 0x35, 0x2a, 0x4b, 0x9c, 0x9c, 0x9f, 0xcf, 0x40, 0xd1, 0xc7, 0x1e, 0x49,
	0x7c, 0xa8, 0x09, 0x74, 0x55, 0xa0, 0x8b, 0xdf, 0x61, 0x8f, 0x24, 0x05, 0x14, 0x00, 0xf8, 0x25,
	0xa8, 0xf4, 0x99, 0x17, 0xb0, 0x88, 0xf2, 0x78, 0xda, 0x66, 0x3b, 0x75, 0x81, 0xae, 0x09, 0x74,
	0x65, 0x37, 0xdd, 0x4a, 0xa2, 0xca, 0xa0, 0xed, 0x3f, 0xce, 0x76, 0x7a, 0x8d, 0x36, 0x28, 0xef,
	0x98, 0x42, 0xb3, 0x05, 0xad, 0x23, 0xce, 0x83, 0xa8, 0x6d, 0x59, 0xc3, 0xe1, 0xb0, 0x15, 0xf1,
	0x90, 0xf9, 0x8e, 0xed, 0xb5, 0xfa, 0xcc, 0xb3, 0x6c, 0xd6, 0x8f, 0xe4, 0xd3, 0x46, 0x7c, 0x4e,
	0x39, 0x25, 0xd1, 0x9a, 0x38, 0xb1, 0xf5, 0x29, 0x44, 0x27, 0x26, 0xa2, 0x36, 0x6a, 0x9b, 0x68,
	0x1b, 0x6d, 0x98, 0x48, 0x84, 0x23, 0x3e, 0x8e, 0x70, 0x10, 0xfc, 0xd9, 0x74, 0x18, 0xe6, 0xe8,
	0xf4, 0x26, 0x8b, 0x5b, 0x9c, 0x91, 0x4d, 0xb2, 0x75, 0x59, 0x00, 0x33, 0x5d, 0xe9, 0xed, 0x13,
	0x0d, 0x94, 0xe2, 0xc7, 0x03, 0x2e, 0xa5, 0x35, 0x50, 0x1e, 0x2f, 0x63, 0x79, 0x72, 0x39, 0xae,
	0x27, 0x72, 0xcf, 0x76, 0x0e, 0xd1, 0x8f, 0xa0, 0x7e, 0x40, 0x70, 0xe8, 0x9b, 0x47, 0x6c, 0x68,
	0x72, 0x66, 0x7a, 0xf8, 0x37, 0x62, 0x62, 0x99, 0x0f, 0xdc, 0x7e, 0x7f, 0x3a, 0x11, 0x09, 0x07,
	0xb4, 0x4f, 0x22, 0x4b, 0x06, 0xb3, 0x16, 0x2b, 0xfc, 0xfd, 0xea, 0xf5, 0xbf, 0x7a, 0x15, 0x55,
	0xac, 0xc1, 0xa6, 0x25, 0x1b, 0xa0, 0xad, 0x35, 0xe0, 0xd7, 0xa0, 0xb0, 0x4f, 0x38, 0x84, 0x69,
	0x30, 0xf9, 0x53, 0x61, 0xd4, 0x94, 0xb5, 0x24, 0xba, 0x65, 0xc9, 0xb2, 0x00, 0xab, 0x19, 0x8b,
	0x75, 0x42, 0xed, 0x53, 0xd8, 0x05, 0xa5, 0xf8, 0x0e, 0xcb, 0xd3, 0x55, 0xee, 0xd4, 0x3c, 0x5d,
	0xf5, 0xaa, 0x43, 0x2b, 0x92, 0xb0, 0x66, 0x4c, 0x10, 0x8a, 0xd8, 0xbe, 0x07, 0xa5, 0x78, 0xd8,
	0x73, 0x4e, 0xe5, 0x7a, 0xc9, 0x39, 0xd5, 0x3b, 0x21, 0x0d, 0xb2, 0x31, 0x19, 0xe4, 0x1e, 0x28,
	0x8a, 0xd1, 0x81, 0x59, 0x66, 0x63, 0xe3, 0x6a, 0x7c, 0xac, 0x2e, 0x26, 0x54, 0x8b, 0x92, 0x6a,
	0x0e, 0xe6, 0xae, 0x19, 0xd5, 0x8b, 0xeb, 0x51, 0xb9, 0x72, 0x7e, 0x9d, 0x54, 0xbe, 0xd3, 0x04,
	0xab, 0x7d, 0xe6, 0xe5, 0x95, 0xc0, 0x01, 0x15, 0x54, 0x81, 0x7b, 0xec, 0xf5, 0xa8, 0xef, 0x74,
	0xe6, 0x65, 0x11, 0x7e, 0x48, 0x3e, 0x7b, 0x25, 0xf9, 0x8b, 0xe7, 0x8b, 0xb7, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0x2f, 0xb1, 0xe5, 0x6b, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RolesClient interface {
	// Create registers a new Role.
	Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	// Get reads one Role by ID.
	Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error)
	// Update patches a Role by ID.
	Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error)
	// Delete removes a Role by ID.
	Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	// List gets a list of Roles matching a given set of criteria.
	List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
}

type rolesClient struct {
	cc *grpc.ClientConn
}

func NewRolesClient(cc *grpc.ClientConn) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) Create(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Get(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error) {
	out := new(RoleGetResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Update(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error) {
	out := new(RoleUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Delete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) List(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, "/v1.Roles/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
type RolesServer interface {
	// Create registers a new Role.
	Create(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	// Get reads one Role by ID.
	Get(context.Context, *RoleGetRequest) (*RoleGetResponse, error)
	// Update patches a Role by ID.
	Update(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error)
	// Delete removes a Role by ID.
	Delete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	// List gets a list of Roles matching a given set of criteria.
	List(context.Context, *RoleListRequest) (*RoleListResponse, error)
}

// UnimplementedRolesServer can be embedded to have forward compatible implementations.
type UnimplementedRolesServer struct {
}

func (*UnimplementedRolesServer) Create(ctx context.Context, req *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRolesServer) Get(ctx context.Context, req *RoleGetRequest) (*RoleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRolesServer) Update(ctx context.Context, req *RoleUpdateRequest) (*RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRolesServer) Delete(ctx context.Context, req *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRolesServer) List(ctx context.Context, req *RoleListRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRolesServer(s *grpc.Server, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Create(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Get(ctx, req.(*RoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Update(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Delete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Roles/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).List(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Roles_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Roles_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Roles_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Roles_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Roles_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roles.proto",
}

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
// source: account_grants.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// AccountGrantCreateRequest specifies what kind of AccountGrants should be registered in
// the organizations fleet.
type AccountGrantCreateRequest struct {
	// Reserved for future use.
	Meta *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Parameters to define the new AccountGrant.
	AccountGrant         *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AccountGrantCreateRequest) Reset()         { *m = AccountGrantCreateRequest{} }
func (m *AccountGrantCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantCreateRequest) ProtoMessage()    {}
func (*AccountGrantCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{0}
}

func (m *AccountGrantCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantCreateRequest.Unmarshal(m, b)
}
func (m *AccountGrantCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantCreateRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantCreateRequest.Merge(m, src)
}
func (m *AccountGrantCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantCreateRequest.Size(m)
}
func (m *AccountGrantCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantCreateRequest proto.InternalMessageInfo

func (m *AccountGrantCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantCreateRequest) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

// AccountGrantCreateResponse reports how the AccountGrants were created in the system.
type AccountGrantCreateResponse struct {
	// Reserved for future use.
	Meta *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The created AccountGrant.
	AccountGrant *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantCreateResponse) Reset()         { *m = AccountGrantCreateResponse{} }
func (m *AccountGrantCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantCreateResponse) ProtoMessage()    {}
func (*AccountGrantCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{1}
}

func (m *AccountGrantCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantCreateResponse.Unmarshal(m, b)
}
func (m *AccountGrantCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantCreateResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantCreateResponse.Merge(m, src)
}
func (m *AccountGrantCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantCreateResponse.Size(m)
}
func (m *AccountGrantCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantCreateResponse proto.InternalMessageInfo

func (m *AccountGrantCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantCreateResponse) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

func (m *AccountGrantCreateResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantGetRequest specifies which AccountGrant to retrieve.
type AccountGrantGetRequest struct {
	// Reserved for future use.
	Meta *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountGrant to retrieve.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantGetRequest) Reset()         { *m = AccountGrantGetRequest{} }
func (m *AccountGrantGetRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantGetRequest) ProtoMessage()    {}
func (*AccountGrantGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{2}
}

func (m *AccountGrantGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantGetRequest.Unmarshal(m, b)
}
func (m *AccountGrantGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantGetRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantGetRequest.Merge(m, src)
}
func (m *AccountGrantGetRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantGetRequest.Size(m)
}
func (m *AccountGrantGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantGetRequest proto.InternalMessageInfo

func (m *AccountGrantGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountGrantGetResponse returns a requested AccountGrant.
type AccountGrantGetResponse struct {
	// Reserved for future use.
	Meta *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The requested AccountGrant.
	AccountGrant *AccountGrant `protobuf:"bytes,2,opt,name=account_grant,json=accountGrant,proto3" json:"account_grant,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantGetResponse) Reset()         { *m = AccountGrantGetResponse{} }
func (m *AccountGrantGetResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantGetResponse) ProtoMessage()    {}
func (*AccountGrantGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{3}
}

func (m *AccountGrantGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantGetResponse.Unmarshal(m, b)
}
func (m *AccountGrantGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantGetResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantGetResponse.Merge(m, src)
}
func (m *AccountGrantGetResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantGetResponse.Size(m)
}
func (m *AccountGrantGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantGetResponse proto.InternalMessageInfo

func (m *AccountGrantGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantGetResponse) GetAccountGrant() *AccountGrant {
	if m != nil {
		return m.AccountGrant
	}
	return nil
}

func (m *AccountGrantGetResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantDeleteRequest identifies a AccountGrant by ID to delete.
type AccountGrantDeleteRequest struct {
	// Reserved for future use.
	Meta *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The unique identifier of the AccountGrant to delete.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantDeleteRequest) Reset()         { *m = AccountGrantDeleteRequest{} }
func (m *AccountGrantDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantDeleteRequest) ProtoMessage()    {}
func (*AccountGrantDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{4}
}

func (m *AccountGrantDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantDeleteRequest.Unmarshal(m, b)
}
func (m *AccountGrantDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantDeleteRequest.Merge(m, src)
}
func (m *AccountGrantDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantDeleteRequest.Size(m)
}
func (m *AccountGrantDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantDeleteRequest proto.InternalMessageInfo

func (m *AccountGrantDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AccountGrantDeleteResponse returns information about a AccountGrant that was deleted.
type AccountGrantDeleteResponse struct {
	// Reserved for future use.
	Meta *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,2,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantDeleteResponse) Reset()         { *m = AccountGrantDeleteResponse{} }
func (m *AccountGrantDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantDeleteResponse) ProtoMessage()    {}
func (*AccountGrantDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{5}
}

func (m *AccountGrantDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantDeleteResponse.Unmarshal(m, b)
}
func (m *AccountGrantDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantDeleteResponse.Merge(m, src)
}
func (m *AccountGrantDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantDeleteResponse.Size(m)
}
func (m *AccountGrantDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantDeleteResponse proto.InternalMessageInfo

func (m *AccountGrantDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantDeleteResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// AccountGrantListRequest specifies criteria for retrieving a list of AccountGrants.
type AccountGrantListRequest struct {
	// Paging parameters for the query.
	Meta *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A human-readable filter query string.
	Filter               string   `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGrantListRequest) Reset()         { *m = AccountGrantListRequest{} }
func (m *AccountGrantListRequest) String() string { return proto.CompactTextString(m) }
func (*AccountGrantListRequest) ProtoMessage()    {}
func (*AccountGrantListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{6}
}

func (m *AccountGrantListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantListRequest.Unmarshal(m, b)
}
func (m *AccountGrantListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantListRequest.Marshal(b, m, deterministic)
}
func (m *AccountGrantListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantListRequest.Merge(m, src)
}
func (m *AccountGrantListRequest) XXX_Size() int {
	return xxx_messageInfo_AccountGrantListRequest.Size(m)
}
func (m *AccountGrantListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantListRequest proto.InternalMessageInfo

func (m *AccountGrantListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// AccountGrantListResponse returns a list of AccountGrants that meet the criteria of a
// AccountGrantListRequest.
type AccountGrantListResponse struct {
	// Paging information for the query.
	Meta *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// A single page of results matching the list request criteria.
	AccountGrants []*AccountGrant `protobuf:"bytes,2,rep,name=account_grants,json=accountGrants,proto3" json:"account_grants,omitempty"`
	// Rate limit information.
	RateLimit            *RateLimitMetadata `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountGrantListResponse) Reset()         { *m = AccountGrantListResponse{} }
func (m *AccountGrantListResponse) String() string { return proto.CompactTextString(m) }
func (*AccountGrantListResponse) ProtoMessage()    {}
func (*AccountGrantListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{7}
}

func (m *AccountGrantListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrantListResponse.Unmarshal(m, b)
}
func (m *AccountGrantListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrantListResponse.Marshal(b, m, deterministic)
}
func (m *AccountGrantListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrantListResponse.Merge(m, src)
}
func (m *AccountGrantListResponse) XXX_Size() int {
	return xxx_messageInfo_AccountGrantListResponse.Size(m)
}
func (m *AccountGrantListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrantListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrantListResponse proto.InternalMessageInfo

func (m *AccountGrantListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AccountGrantListResponse) GetAccountGrants() []*AccountGrant {
	if m != nil {
		return m.AccountGrants
	}
	return nil
}

func (m *AccountGrantListResponse) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// A AccountGrant connects a composite role to another role, granting members
// of the composite role the permissions granted to the attached role.
type AccountGrant struct {
	// Unique identifier of the AccountGrant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the composite role of this AccountGrant.
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The id of the attached role of this AccountGrant.
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The timestamp when the resource will be granted. Optional. Both start_at
	// and end_at must be defined together, or not defined at all.
	StartFrom *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
	// The timestamp when the resource grant will expire. Optional. Both
	// start_at and end_at must be defined together, or not defined at all.
	ValidUntil           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AccountGrant) Reset()         { *m = AccountGrant{} }
func (m *AccountGrant) String() string { return proto.CompactTextString(m) }
func (*AccountGrant) ProtoMessage()    {}
func (*AccountGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a322caa55cfb217, []int{8}
}

func (m *AccountGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGrant.Unmarshal(m, b)
}
func (m *AccountGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGrant.Marshal(b, m, deterministic)
}
func (m *AccountGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGrant.Merge(m, src)
}
func (m *AccountGrant) XXX_Size() int {
	return xxx_messageInfo_AccountGrant.Size(m)
}
func (m *AccountGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGrant.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGrant proto.InternalMessageInfo

func (m *AccountGrant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountGrant) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AccountGrant) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccountGrant) GetStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.StartFrom
	}
	return nil
}

func (m *AccountGrant) GetValidUntil() *timestamp.Timestamp {
	if m != nil {
		return m.ValidUntil
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountGrantCreateRequest)(nil), "v1.AccountGrantCreateRequest")
	proto.RegisterType((*AccountGrantCreateResponse)(nil), "v1.AccountGrantCreateResponse")
	proto.RegisterType((*AccountGrantGetRequest)(nil), "v1.AccountGrantGetRequest")
	proto.RegisterType((*AccountGrantGetResponse)(nil), "v1.AccountGrantGetResponse")
	proto.RegisterType((*AccountGrantDeleteRequest)(nil), "v1.AccountGrantDeleteRequest")
	proto.RegisterType((*AccountGrantDeleteResponse)(nil), "v1.AccountGrantDeleteResponse")
	proto.RegisterType((*AccountGrantListRequest)(nil), "v1.AccountGrantListRequest")
	proto.RegisterType((*AccountGrantListResponse)(nil), "v1.AccountGrantListResponse")
	proto.RegisterType((*AccountGrant)(nil), "v1.AccountGrant")
}

func init() { proto.RegisterFile("account_grants.proto", fileDescriptor_5a322caa55cfb217) }

var fileDescriptor_5a322caa55cfb217 = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0x3a, 0xae, 0x2d, 0x4f, 0x12, 0xa0, 0xd3, 0xd6, 0x5d, 0x6f, 0x43, 0x58, 0x4d, 0x85,
	0x84, 0x4c, 0xec, 0xc5, 0xc1, 0xa8, 0x92, 0x25, 0x44, 0xed, 0x46, 0x44, 0x81, 0x20, 0x55, 0x2e,
	0xed, 0x09, 0x61, 0x26, 0xbb, 0x93, 0xed, 0x80, 0x77, 0x67, 0x99, 0x19, 0xdb, 0x87, 0xaa, 0x17,
	0xc4, 0x15, 0x09, 0x95, 0x3f, 0x00, 0x89, 0x03, 0x82, 0x03, 0x12, 0x92, 0x2f, 0x9c, 0x39, 0xf6,
	0x00, 0x07, 0xfe, 0x04, 0xb8, 0x70, 0x25, 0x17, 0xda, 0x13, 0x9a, 0xfd, 0x61, 0xef, 0xd8, 0x4e,
	0x2a, 0x14, 0x2e, 0x9c, 0xec, 0x9d, 0xf7, 0xbe, 0xf7, 0xbd, 0xf9, 0xe6, 0x9b, 0x79, 0xe0, 0x32,
	0x76, 0x5d, 0x36, 0x0a, 0xe5, 0xc0, 0xe7, 0x38, 0x94, 0xa2, 0x19, 0x71, 0x26, 0x19, 0x2c, 0x8c,
	0x5b, 0xd6, 0x96, 0xcf, 0x98, 0x3f, 0x24, 0x0e, 0x8e, 0xa8, 0x83, 0xc3, 0x90, 0x49, 0x2c, 0x29,
	0x0b, 0xd3, 0x0c, 0x6b, 0x27, 0xfe, 0x71, 0x1b, 0x3e, 0x09, 0x1b, 0x62, 0x82, 0x7d, 0x9f, 0x70,
	0x87, 0x45, 0x71, 0xc6, 0x8a, 0xec, 0x97, 0xd2, 0x5a, 0xf1, 0xd7, 0xd1, 0xe8, 0xd8, 0x91, 0x34,
	0x20, 0x42, 0xe2, 0x20, 0x4a, 0x13, 0x36, 0x53, 0x6c, 0xfa, 0x09, 0x44, 0x44, 0xdc, 0xe4, 0x3f,
	0xfa, 0xc2, 0x00, 0xb5, 0x6e, 0xd2, 0xe4, 0xbe, 0xea, 0xf1, 0x16, 0x27, 0x58, 0x92, 0x3e, 0xf9,
	0x74, 0x44, 0x84, 0x84, 0x0d, 0x50, 0x0c, 0x88, 0xc4, 0xa6, 0x61, 0x1b, 0xaf, 0xac, 0xef, 0xd6,
	0x9a, 0xe3, 0x56, 0x53, 0x4b, 0x78, 0x8f, 0x48, 0xec, 0x61, 0x89, 0xfb, 0x71, 0x1a, 0xec, 0x82,
	0x4d, 0x6d, 0xc3, 0x66, 0x21, 0xc6, 0xbd, 0xa0, 0x70, 0x79, 0x92, 0x1e, 0xf8, 0xeb, 0xef, 0x69,
	0xf9, 0xc2, 0x8f, 0x27, 0xd3, 0xb2, 0xd1, 0xdf, 0xc0, 0xb9, 0x08, 0xfa, 0xd3, 0x00, 0xd6, 0xaa,
	0x7e, 0x44, 0xc4, 0x42, 0x41, 0x60, 0x47, 0x6b, 0xc8, 0xca, 0x37, 0x94, 0x64, 0x64, 0x1d, 0x69,
	0x14, 0xff, 0x55, 0x77, 0xf0, 0x26, 0x00, 0x1c, 0x4b, 0x32, 0x18, 0xd2, 0x80, 0x4a, 0x73, 0x2d,
	0xc6, 0x5f, 0x51, 0xf8, 0x3e, 0x96, 0xe4, 0x50, 0x2d, 0xae, 0xe4, 0xaf, 0xf0, 0x2c, 0xdc, 0x01,
	0x4f, 0xd5, 0xf2, 0x77, 0x6a, 0x19, 0x7d, 0x04, 0xaa, 0x79, 0xde, 0x7d, 0x22, 0x33, 0xdd, 0xeb,
	0xda, 0x36, 0xab, 0x8a, 0x61, 0x1e, 0x5d, 0x10, 0xdd, 0x02, 0x05, 0xea, 0xc5, 0x7b, 0xa9, 0x68,
	0xa4, 0x05, 0xea, 0xa1, 0xdf, 0x0d, 0x70, 0x75, 0x89, 0x22, 0x95, 0xf2, 0x86, 0xc6, 0x71, 0x75,
	0xc6, 0xf1, 0xff, 0xd3, 0xf1, 0x58, 0xb7, 0xf0, 0x1e, 0x19, 0x92, 0x33, 0x2d, 0xac, 0x25, 0xfc,
	0x0b, 0x35, 0xbf, 0x5d, 0xf0, 0x66, 0x56, 0xe7, 0x74, 0x6f, 0xea, 0x19, 0x67, 0x68, 0xaa, 0x0b,
	0x52, 0x38, 0xa7, 0x20, 0x1f, 0xeb, 0xa7, 0x7e, 0x48, 0xc5, 0xcc, 0x59, 0xaf, 0x2e, 0x9f, 0x7a,
	0x2e, 0xbc, 0x20, 0x06, 0x02, 0xa5, 0x63, 0x3a, 0x94, 0x84, 0xaf, 0x10, 0x24, 0x8d, 0xa0, 0x5f,
	0x0d, 0x60, 0x2e, 0x93, 0xa5, 0x92, 0xec, 0x68, 0x6c, 0xe6, 0x9c, 0x4d, 0x17, 0x24, 0xa5, 0xbb,
	0x05, 0x9e, 0xd3, 0xdf, 0x4b, 0xb3, 0x60, 0xaf, 0x9d, 0xe1, 0xac, 0x9f, 0xe2, 0x46, 0x36, 0xf3,
	0xce, 0x12, 0xe7, 0xb7, 0x16, 0xfa, 0xbe, 0x08, 0x36, 0xf2, 0x6c, 0xf0, 0xb5, 0xd8, 0x13, 0x46,
	0x2c, 0x81, 0xad, 0x30, 0xd7, 0xbe, 0x39, 0x99, 0x96, 0x0b, 0x07, 0x7b, 0x31, 0xf4, 0xf1, 0xc9,
	0xb4, 0x0c, 0x6e, 0x13, 0x1e, 0x50, 0x21, 0x28, 0x0b, 0x95, 0x53, 0xe0, 0xbb, 0x60, 0x9d, 0x13,
	0xc1, 0x46, 0xdc, 0x25, 0x83, 0x99, 0x9d, 0xea, 0x0a, 0xfa, 0xb2, 0x82, 0x6e, 0xec, 0x61, 0x89,
	0x93, 0x70, 0xbe, 0xc8, 0x7c, 0xb5, 0x0f, 0x32, 0xf8, 0x81, 0x07, 0xdf, 0x04, 0x20, 0x93, 0x85,
	0x7a, 0xf1, 0x8e, 0x2a, 0xbd, 0x6d, 0x55, 0xab, 0xa6, 0x6a, 0x95, 0xee, 0x0a, 0xc2, 0x73, 0x55,
	0x8a, 0xea, 0xbb, 0x5f, 0x49, 0x11, 0x07, 0x1e, 0xbc, 0x03, 0x80, 0x90, 0x98, 0xcb, 0xc1, 0x31,
	0x67, 0x81, 0x59, 0x4c, 0xcd, 0x99, 0x8c, 0x8c, 0x66, 0x36, 0x32, 0x9a, 0xef, 0x67, 0x23, 0xa3,
	0x67, 0xaa, 0xd2, 0x97, 0x54, 0xe9, 0xca, 0x1d, 0x85, 0x7b, 0x9b, 0xb3, 0x20, 0xd5, 0x48, 0x64,
	0xdf, 0xf0, 0x1e, 0x58, 0x1f, 0xe3, 0x21, 0xf5, 0x06, 0xa3, 0x50, 0xd2, 0xa1, 0x79, 0xe1, 0x99,
	0x55, 0x6b, 0xaa, 0xea, 0x65, 0x55, 0x15, 0xdc, 0x53, 0xc0, 0xbb, 0x0a, 0x97, 0x94, 0x05, 0xe3,
	0xd9, 0x42, 0xe7, 0x07, 0xe3, 0x51, 0xf7, 0x4b, 0xa3, 0xfe, 0x0e, 0x78, 0xbe, 0x6b, 0xe7, 0xcf,
	0xa0, 0x09, 0x6f, 0xdc, 0x97, 0x32, 0x12, 0x1d, 0xc7, 0x99, 0x4c, 0x26, 0x4d, 0x21, 0x39, 0x0b,
	0x7d, 0x2f, 0x68, 0xba, 0x2c, 0x70, 0x3c, 0xe6, 0x8a, 0x78, 0x8e, 0x92, 0x50, 0x52, 0x49, 0x89,
	0xb8, 0x9e, 0x47, 0xee, 0xf6, 0xe0, 0xcd, 0x07, 0x36, 0xa2, 0x1e, 0xea, 0xd8, 0x68, 0xe4, 0x37,
	0x76, 0xdb, 0x6d, 0xb4, 0x63, 0xa3, 0xdc, 0xf1, 0xa8, 0x00, 0x17, 0x8d, 0x56, 0xab, 0xa5, 0x02,
	0x73, 0xa9, 0x51, 0x07, 0xe1, 0x46, 0xbb, 0xdd, 0x46, 0x0f, 0x9f, 0xce, 0x7a, 0x9f, 0x9f, 0x73,
	0x7c, 0xd3, 0x76, 0xbf, 0x2e, 0x82, 0xcd, 0xae, 0xe6, 0xbf, 0x5f, 0x0c, 0x50, 0x4a, 0x46, 0x12,
	0x7c, 0x71, 0xd1, 0xb7, 0xda, 0xec, 0xb4, 0xb6, 0x4f, 0x0b, 0x27, 0x97, 0x03, 0x7d, 0x6e, 0x3c,
	0xea, 0x12, 0xe4, 0x02, 0xfb, 0x90, 0x60, 0x1e, 0xda, 0xf7, 0xd9, 0xc4, 0x96, 0xcc, 0x0e, 0xf0,
	0x27, 0xc4, 0xc6, 0x9a, 0x38, 0xf0, 0xad, 0x67, 0x6b, 0x23, 0x08, 0x1f, 0x53, 0x97, 0x08, 0x47,
	0xeb, 0xf6, 0x7a, 0x42, 0xf9, 0xd9, 0x6f, 0x7f, 0x7c, 0x55, 0xa8, 0xa2, 0x8b, 0xce, 0xb8, 0xe5,
	0x68, 0xb7, 0xa9, 0x63, 0xd4, 0xe1, 0x87, 0x60, 0x6d, 0x9f, 0x48, 0x68, 0x2d, 0x76, 0x3b, 0x1f,
	0x48, 0xd6, 0xb5, 0x95, 0xb1, 0x74, 0x1b, 0xdb, 0x71, 0x75, 0x13, 0x56, 0x97, 0xaa, 0x3b, 0x0f,
	0xa8, 0xf7, 0x10, 0xfa, 0xa0, 0x94, 0x3c, 0x93, 0xcb, 0x7a, 0x69, 0x0f, 0xf5, 0xb2, 0x5e, 0xfa,
	0xeb, 0x9a, 0x11, 0xd5, 0x4f, 0x23, 0xfa, 0x00, 0x14, 0xd5, 0xe3, 0x03, 0x97, 0xba, 0xcd, 0x3d,
	0x80, 0xd6, 0xd6, 0xea, 0x60, 0x4a, 0x51, 0x8b, 0x29, 0x2e, 0xc1, 0x65, 0xa5, 0xac, 0xea, 0xe3,
	0x27, 0xd3, 0xf2, 0xc5, 0x9f, 0x9f, 0x4c, 0xcb, 0xda, 0xe3, 0xd1, 0x7b, 0x03, 0x6c, 0xb9, 0x2c,
	0x98, 0x9f, 0x0a, 0x8e, 0xa8, 0xa2, 0x88, 0x86, 0xa3, 0xe0, 0x88, 0x86, 0x7e, 0xef, 0x8a, 0x76,
	0x20, 0xb7, 0xd3, 0xe5, 0xa3, 0x52, 0x7c, 0x87, 0x5e, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0xf0, 0x43, 0xad, 0x42, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountGrantsClient is the client API for AccountGrants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountGrantsClient interface {
	// Create registers a new AccountGrant.
	Create(ctx context.Context, in *AccountGrantCreateRequest, opts ...grpc.CallOption) (*AccountGrantCreateResponse, error)
	// Get reads one AccountGrant by ID.
	Get(ctx context.Context, in *AccountGrantGetRequest, opts ...grpc.CallOption) (*AccountGrantGetResponse, error)
	// Delete removes a AccountGrant by ID.
	Delete(ctx context.Context, in *AccountGrantDeleteRequest, opts ...grpc.CallOption) (*AccountGrantDeleteResponse, error)
	// List gets a list of AccountGrants matching a given set of criteria.
	List(ctx context.Context, in *AccountGrantListRequest, opts ...grpc.CallOption) (*AccountGrantListResponse, error)
}

type accountGrantsClient struct {
	cc *grpc.ClientConn
}

func NewAccountGrantsClient(cc *grpc.ClientConn) AccountGrantsClient {
	return &accountGrantsClient{cc}
}

func (c *accountGrantsClient) Create(ctx context.Context, in *AccountGrantCreateRequest, opts ...grpc.CallOption) (*AccountGrantCreateResponse, error) {
	out := new(AccountGrantCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) Get(ctx context.Context, in *AccountGrantGetRequest, opts ...grpc.CallOption) (*AccountGrantGetResponse, error) {
	out := new(AccountGrantGetResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) Delete(ctx context.Context, in *AccountGrantDeleteRequest, opts ...grpc.CallOption) (*AccountGrantDeleteResponse, error) {
	out := new(AccountGrantDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountGrantsClient) List(ctx context.Context, in *AccountGrantListRequest, opts ...grpc.CallOption) (*AccountGrantListResponse, error) {
	out := new(AccountGrantListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccountGrants/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountGrantsServer is the server API for AccountGrants service.
type AccountGrantsServer interface {
	// Create registers a new AccountGrant.
	Create(context.Context, *AccountGrantCreateRequest) (*AccountGrantCreateResponse, error)
	// Get reads one AccountGrant by ID.
	Get(context.Context, *AccountGrantGetRequest) (*AccountGrantGetResponse, error)
	// Delete removes a AccountGrant by ID.
	Delete(context.Context, *AccountGrantDeleteRequest) (*AccountGrantDeleteResponse, error)
	// List gets a list of AccountGrants matching a given set of criteria.
	List(context.Context, *AccountGrantListRequest) (*AccountGrantListResponse, error)
}

// UnimplementedAccountGrantsServer can be embedded to have forward compatible implementations.
type UnimplementedAccountGrantsServer struct {
}

func (*UnimplementedAccountGrantsServer) Create(ctx context.Context, req *AccountGrantCreateRequest) (*AccountGrantCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountGrantsServer) Get(ctx context.Context, req *AccountGrantGetRequest) (*AccountGrantGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountGrantsServer) Delete(ctx context.Context, req *AccountGrantDeleteRequest) (*AccountGrantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAccountGrantsServer) List(ctx context.Context, req *AccountGrantListRequest) (*AccountGrantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterAccountGrantsServer(s *grpc.Server, srv AccountGrantsServer) {
	s.RegisterService(&_AccountGrants_serviceDesc, srv)
}

func _AccountGrants_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Create(ctx, req.(*AccountGrantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Get(ctx, req.(*AccountGrantGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).Delete(ctx, req.(*AccountGrantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountGrants_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGrantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountGrantsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccountGrants/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountGrantsServer).List(ctx, req.(*AccountGrantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountGrants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccountGrants",
	HandlerType: (*AccountGrantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccountGrants_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountGrants_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountGrants_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountGrants_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account_grants.proto",
}

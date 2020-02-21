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
// source: spec.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// AlreadyExistsError is used when an entity already exists in the system
type AlreadyExistsError struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlreadyExistsError) Reset()         { *m = AlreadyExistsError{} }
func (m *AlreadyExistsError) String() string { return proto.CompactTextString(m) }
func (*AlreadyExistsError) ProtoMessage()    {}
func (*AlreadyExistsError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{0}
}

func (m *AlreadyExistsError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlreadyExistsError.Unmarshal(m, b)
}
func (m *AlreadyExistsError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlreadyExistsError.Marshal(b, m, deterministic)
}
func (m *AlreadyExistsError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlreadyExistsError.Merge(m, src)
}
func (m *AlreadyExistsError) XXX_Size() int {
	return xxx_messageInfo_AlreadyExistsError.Size(m)
}
func (m *AlreadyExistsError) XXX_DiscardUnknown() {
	xxx_messageInfo_AlreadyExistsError.DiscardUnknown(m)
}

var xxx_messageInfo_AlreadyExistsError proto.InternalMessageInfo

func (m *AlreadyExistsError) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

// NotFoundError is used when an entity does not exist in the system
type NotFoundError struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotFoundError) Reset()         { *m = NotFoundError{} }
func (m *NotFoundError) String() string { return proto.CompactTextString(m) }
func (*NotFoundError) ProtoMessage()    {}
func (*NotFoundError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{1}
}

func (m *NotFoundError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotFoundError.Unmarshal(m, b)
}
func (m *NotFoundError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotFoundError.Marshal(b, m, deterministic)
}
func (m *NotFoundError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotFoundError.Merge(m, src)
}
func (m *NotFoundError) XXX_Size() int {
	return xxx_messageInfo_NotFoundError.Size(m)
}
func (m *NotFoundError) XXX_DiscardUnknown() {
	xxx_messageInfo_NotFoundError.DiscardUnknown(m)
}

var xxx_messageInfo_NotFoundError proto.InternalMessageInfo

func (m *NotFoundError) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

// BadRequestError identifies a bad request sent by the client
type BadRequestError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BadRequestError) Reset()         { *m = BadRequestError{} }
func (m *BadRequestError) String() string { return proto.CompactTextString(m) }
func (*BadRequestError) ProtoMessage()    {}
func (*BadRequestError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{2}
}

func (m *BadRequestError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequestError.Unmarshal(m, b)
}
func (m *BadRequestError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequestError.Marshal(b, m, deterministic)
}
func (m *BadRequestError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequestError.Merge(m, src)
}
func (m *BadRequestError) XXX_Size() int {
	return xxx_messageInfo_BadRequestError.Size(m)
}
func (m *BadRequestError) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequestError.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequestError proto.InternalMessageInfo

// AuthenticationError is used to specify an authentication failure condition
type AuthenticationError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationError) Reset()         { *m = AuthenticationError{} }
func (m *AuthenticationError) String() string { return proto.CompactTextString(m) }
func (*AuthenticationError) ProtoMessage()    {}
func (*AuthenticationError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{3}
}

func (m *AuthenticationError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationError.Unmarshal(m, b)
}
func (m *AuthenticationError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationError.Marshal(b, m, deterministic)
}
func (m *AuthenticationError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationError.Merge(m, src)
}
func (m *AuthenticationError) XXX_Size() int {
	return xxx_messageInfo_AuthenticationError.Size(m)
}
func (m *AuthenticationError) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationError.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationError proto.InternalMessageInfo

// PermissionError is used to specify a permissions violation
type PermissionError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermissionError) Reset()         { *m = PermissionError{} }
func (m *PermissionError) String() string { return proto.CompactTextString(m) }
func (*PermissionError) ProtoMessage()    {}
func (*PermissionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{4}
}

func (m *PermissionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermissionError.Unmarshal(m, b)
}
func (m *PermissionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermissionError.Marshal(b, m, deterministic)
}
func (m *PermissionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionError.Merge(m, src)
}
func (m *PermissionError) XXX_Size() int {
	return xxx_messageInfo_PermissionError.Size(m)
}
func (m *PermissionError) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionError.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionError proto.InternalMessageInfo

// InternalError is used to specify an internal system error
type InternalError struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalError) Reset()         { *m = InternalError{} }
func (m *InternalError) String() string { return proto.CompactTextString(m) }
func (*InternalError) ProtoMessage()    {}
func (*InternalError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{5}
}

func (m *InternalError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalError.Unmarshal(m, b)
}
func (m *InternalError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalError.Marshal(b, m, deterministic)
}
func (m *InternalError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalError.Merge(m, src)
}
func (m *InternalError) XXX_Size() int {
	return xxx_messageInfo_InternalError.Size(m)
}
func (m *InternalError) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalError.DiscardUnknown(m)
}

var xxx_messageInfo_InternalError proto.InternalMessageInfo

// RateLimitError is used for rate limit excess condition
type RateLimitError struct {
	RateLimit            *RateLimitMetadata `protobuf:"bytes,1,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RateLimitError) Reset()         { *m = RateLimitError{} }
func (m *RateLimitError) String() string { return proto.CompactTextString(m) }
func (*RateLimitError) ProtoMessage()    {}
func (*RateLimitError) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{6}
}

func (m *RateLimitError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitError.Unmarshal(m, b)
}
func (m *RateLimitError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitError.Marshal(b, m, deterministic)
}
func (m *RateLimitError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitError.Merge(m, src)
}
func (m *RateLimitError) XXX_Size() int {
	return xxx_messageInfo_RateLimitError.Size(m)
}
func (m *RateLimitError) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitError.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitError proto.InternalMessageInfo

func (m *RateLimitError) GetRateLimit() *RateLimitMetadata {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

// CreateRequestMetadata is reserved for future use.
type CreateRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequestMetadata) Reset()         { *m = CreateRequestMetadata{} }
func (m *CreateRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateRequestMetadata) ProtoMessage()    {}
func (*CreateRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{7}
}

func (m *CreateRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequestMetadata.Unmarshal(m, b)
}
func (m *CreateRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequestMetadata.Marshal(b, m, deterministic)
}
func (m *CreateRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequestMetadata.Merge(m, src)
}
func (m *CreateRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateRequestMetadata.Size(m)
}
func (m *CreateRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequestMetadata proto.InternalMessageInfo

// CreateResponseMetadata is reserved for future use.
type CreateResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponseMetadata) Reset()         { *m = CreateResponseMetadata{} }
func (m *CreateResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateResponseMetadata) ProtoMessage()    {}
func (*CreateResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{8}
}

func (m *CreateResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponseMetadata.Unmarshal(m, b)
}
func (m *CreateResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponseMetadata.Marshal(b, m, deterministic)
}
func (m *CreateResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponseMetadata.Merge(m, src)
}
func (m *CreateResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateResponseMetadata.Size(m)
}
func (m *CreateResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponseMetadata proto.InternalMessageInfo

// GetRequestMetadata is reserved for future use.
type GetRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequestMetadata) Reset()         { *m = GetRequestMetadata{} }
func (m *GetRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*GetRequestMetadata) ProtoMessage()    {}
func (*GetRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{9}
}

func (m *GetRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequestMetadata.Unmarshal(m, b)
}
func (m *GetRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequestMetadata.Marshal(b, m, deterministic)
}
func (m *GetRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequestMetadata.Merge(m, src)
}
func (m *GetRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_GetRequestMetadata.Size(m)
}
func (m *GetRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequestMetadata proto.InternalMessageInfo

// GetResponseMetadata is reserved for future use.
type GetResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponseMetadata) Reset()         { *m = GetResponseMetadata{} }
func (m *GetResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*GetResponseMetadata) ProtoMessage()    {}
func (*GetResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{10}
}

func (m *GetResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponseMetadata.Unmarshal(m, b)
}
func (m *GetResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponseMetadata.Marshal(b, m, deterministic)
}
func (m *GetResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponseMetadata.Merge(m, src)
}
func (m *GetResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_GetResponseMetadata.Size(m)
}
func (m *GetResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponseMetadata proto.InternalMessageInfo

// UpdateRequestMetadata is reserved for future use.
type UpdateRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequestMetadata) Reset()         { *m = UpdateRequestMetadata{} }
func (m *UpdateRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateRequestMetadata) ProtoMessage()    {}
func (*UpdateRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{11}
}

func (m *UpdateRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequestMetadata.Unmarshal(m, b)
}
func (m *UpdateRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequestMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequestMetadata.Merge(m, src)
}
func (m *UpdateRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateRequestMetadata.Size(m)
}
func (m *UpdateRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequestMetadata proto.InternalMessageInfo

// UpdateResponseMetadata is reserved for future use.
type UpdateResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponseMetadata) Reset()         { *m = UpdateResponseMetadata{} }
func (m *UpdateResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateResponseMetadata) ProtoMessage()    {}
func (*UpdateResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{12}
}

func (m *UpdateResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponseMetadata.Unmarshal(m, b)
}
func (m *UpdateResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponseMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponseMetadata.Merge(m, src)
}
func (m *UpdateResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateResponseMetadata.Size(m)
}
func (m *UpdateResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponseMetadata proto.InternalMessageInfo

// DeleteRequestMetadata is reserved for future use.
type DeleteRequestMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequestMetadata) Reset()         { *m = DeleteRequestMetadata{} }
func (m *DeleteRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteRequestMetadata) ProtoMessage()    {}
func (*DeleteRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{13}
}

func (m *DeleteRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequestMetadata.Unmarshal(m, b)
}
func (m *DeleteRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequestMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequestMetadata.Merge(m, src)
}
func (m *DeleteRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteRequestMetadata.Size(m)
}
func (m *DeleteRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequestMetadata proto.InternalMessageInfo

// DeleteResponseMetadata is reserved for future use.
type DeleteResponseMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponseMetadata) Reset()         { *m = DeleteResponseMetadata{} }
func (m *DeleteResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteResponseMetadata) ProtoMessage()    {}
func (*DeleteResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{14}
}

func (m *DeleteResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponseMetadata.Unmarshal(m, b)
}
func (m *DeleteResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponseMetadata.Marshal(b, m, deterministic)
}
func (m *DeleteResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponseMetadata.Merge(m, src)
}
func (m *DeleteResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteResponseMetadata.Size(m)
}
func (m *DeleteResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponseMetadata proto.InternalMessageInfo

// ListRequestMetadata specifies paging parameters for listing entities. If this
// metadata is not provided, the default behavior is to return the first page of
// entities, along with a cursor which can be used to fetch the remaining pages.
type ListRequestMetadata struct {
	// The cursor specifies where to start fetching entities in the total list
	// of all entities. If the cursor is non-empty, the page and limit
	// parameters are ignored. See ListResponseMetadata.next_cursor.
	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// The page number to fetch. Use of this parameter is not recommended. Use
	// the cursor instead.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// The number of entities to fetch in a single page. If not specified, a
	// default value will be used.
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequestMetadata) Reset()         { *m = ListRequestMetadata{} }
func (m *ListRequestMetadata) String() string { return proto.CompactTextString(m) }
func (*ListRequestMetadata) ProtoMessage()    {}
func (*ListRequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{15}
}

func (m *ListRequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequestMetadata.Unmarshal(m, b)
}
func (m *ListRequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequestMetadata.Marshal(b, m, deterministic)
}
func (m *ListRequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequestMetadata.Merge(m, src)
}
func (m *ListRequestMetadata) XXX_Size() int {
	return xxx_messageInfo_ListRequestMetadata.Size(m)
}
func (m *ListRequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequestMetadata proto.InternalMessageInfo

func (m *ListRequestMetadata) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *ListRequestMetadata) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequestMetadata) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// ListResponseMetadata contains paging information about the remaining entities
// in a list request.
type ListResponseMetadata struct {
	// A cursor to fetch the next page. If the cursor is an empty string, there
	// are no more entities to fetch. If the cursor is non-empty, make another
	// list request and pass the cursor value in the metadata.
	NextCursor string `protobuf:"bytes,1,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// The total count of all entities matching the criteria of a list request.
	// Note that this value may change between page requests.
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponseMetadata) Reset()         { *m = ListResponseMetadata{} }
func (m *ListResponseMetadata) String() string { return proto.CompactTextString(m) }
func (*ListResponseMetadata) ProtoMessage()    {}
func (*ListResponseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{16}
}

func (m *ListResponseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponseMetadata.Unmarshal(m, b)
}
func (m *ListResponseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponseMetadata.Marshal(b, m, deterministic)
}
func (m *ListResponseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponseMetadata.Merge(m, src)
}
func (m *ListResponseMetadata) XXX_Size() int {
	return xxx_messageInfo_ListResponseMetadata.Size(m)
}
func (m *ListResponseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponseMetadata proto.InternalMessageInfo

func (m *ListResponseMetadata) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *ListResponseMetadata) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// RateLimitMetadata contains information about remaining requests avaialable
// to the user over some timeframe.
type RateLimitMetadata struct {
	// How many total requests the user/token is authorized to make before being
	// rate limited.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// How many remaining requests out of the limit are still avaialable.
	Remaining int64 `protobuf:"varint,2,opt,name=remaining,proto3" json:"remaining,omitempty"`
	// The time when remaining will be reset to limit.
	ResetAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=reset_at,json=resetAt,proto3" json:"reset_at,omitempty"`
	// The bucket this user/token is associated with, which may be shared between
	// multiple users/tokens.
	Bucket               string   `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitMetadata) Reset()         { *m = RateLimitMetadata{} }
func (m *RateLimitMetadata) String() string { return proto.CompactTextString(m) }
func (*RateLimitMetadata) ProtoMessage()    {}
func (*RateLimitMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_423806180556987f, []int{17}
}

func (m *RateLimitMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitMetadata.Unmarshal(m, b)
}
func (m *RateLimitMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitMetadata.Marshal(b, m, deterministic)
}
func (m *RateLimitMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitMetadata.Merge(m, src)
}
func (m *RateLimitMetadata) XXX_Size() int {
	return xxx_messageInfo_RateLimitMetadata.Size(m)
}
func (m *RateLimitMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitMetadata proto.InternalMessageInfo

func (m *RateLimitMetadata) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RateLimitMetadata) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *RateLimitMetadata) GetResetAt() *timestamp.Timestamp {
	if m != nil {
		return m.ResetAt
	}
	return nil
}

func (m *RateLimitMetadata) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func init() {
	proto.RegisterType((*AlreadyExistsError)(nil), "v1.AlreadyExistsError")
	proto.RegisterType((*NotFoundError)(nil), "v1.NotFoundError")
	proto.RegisterType((*BadRequestError)(nil), "v1.BadRequestError")
	proto.RegisterType((*AuthenticationError)(nil), "v1.AuthenticationError")
	proto.RegisterType((*PermissionError)(nil), "v1.PermissionError")
	proto.RegisterType((*InternalError)(nil), "v1.InternalError")
	proto.RegisterType((*RateLimitError)(nil), "v1.RateLimitError")
	proto.RegisterType((*CreateRequestMetadata)(nil), "v1.CreateRequestMetadata")
	proto.RegisterType((*CreateResponseMetadata)(nil), "v1.CreateResponseMetadata")
	proto.RegisterType((*GetRequestMetadata)(nil), "v1.GetRequestMetadata")
	proto.RegisterType((*GetResponseMetadata)(nil), "v1.GetResponseMetadata")
	proto.RegisterType((*UpdateRequestMetadata)(nil), "v1.UpdateRequestMetadata")
	proto.RegisterType((*UpdateResponseMetadata)(nil), "v1.UpdateResponseMetadata")
	proto.RegisterType((*DeleteRequestMetadata)(nil), "v1.DeleteRequestMetadata")
	proto.RegisterType((*DeleteResponseMetadata)(nil), "v1.DeleteResponseMetadata")
	proto.RegisterType((*ListRequestMetadata)(nil), "v1.ListRequestMetadata")
	proto.RegisterType((*ListResponseMetadata)(nil), "v1.ListResponseMetadata")
	proto.RegisterType((*RateLimitMetadata)(nil), "v1.RateLimitMetadata")
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor_423806180556987f) }

var fileDescriptor_423806180556987f = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xea, 0x46,
	0x14, 0x16, 0x21, 0x40, 0x72, 0x22, 0xda, 0x74, 0x02, 0x09, 0x45, 0xf9, 0x31, 0x56, 0x17, 0x59,
	0x14, 0xbb, 0xa4, 0xab, 0xb2, 0x0a, 0xe4, 0xa7, 0x8a, 0x94, 0x54, 0x11, 0x4d, 0xd5, 0x4d, 0xa5,
	0x68, 0x30, 0x27, 0xc6, 0x2a, 0x9e, 0x71, 0x66, 0x8e, 0x43, 0xf2, 0x1a, 0x7d, 0x84, 0xae, 0xba,
	0xac, 0xc4, 0x93, 0xb4, 0x8f, 0x72, 0xb3, 0x89, 0xee, 0xea, 0x6a, 0x6c, 0x9c, 0x0b, 0x5c, 0x94,
	0x2b, 0xdd, 0x15, 0xcc, 0xf7, 0x3b, 0x3e, 0x63, 0x0f, 0x80, 0x8e, 0xd0, 0x73, 0x22, 0x25, 0x49,
	0xb2, 0x95, 0x87, 0x56, 0xbd, 0x2c, 0x23, 0x0a, 0xa4, 0xd0, 0x29, 0x54, 0xff, 0x3e, 0xf9, 0xf1,
	0x9a, 0x3e, 0x8a, 0xa6, 0x1e, 0x73, 0xdf, 0x47, 0xe5, 0x4e, 0x15, 0x2e, 0x17, 0x42, 0x12, 0x9f,
	0x55, 0x1f, 0xf8, 0x52, 0xfa, 0x23, 0x74, 0x93, 0x55, 0x3f, 0xbe, 0x73, 0x29, 0x08, 0x51, 0x13,
	0x0f, 0xa3, 0x54, 0x60, 0xdf, 0x01, 0xeb, 0x8c, 0x14, 0xf2, 0xc1, 0xd3, 0xd9, 0x63, 0xa0, 0x49,
	0x9f, 0x29, 0x25, 0x15, 0x6b, 0x42, 0x11, 0x05, 0x05, 0xf4, 0x54, 0xcb, 0x59, 0xb9, 0xc3, 0xf5,
	0x6e, 0xf5, 0xdd, 0xcb, 0xa4, 0xb4, 0xf9, 0xf7, 0xf3, 0xa4, 0x54, 0x3c, 0x4b, 0xf0, 0x7f, 0x9f,
	0x27, 0xa5, 0x5c, 0x6f, 0x2a, 0x6a, 0x37, 0xde, 0xbf, 0x4c, 0x4a, 0xbb, 0x86, 0x5e, 0x12, 0x66,
	0xa4, 0x45, 0xfb, 0x0f, 0x28, 0xff, 0x22, 0xe9, 0x5c, 0xc6, 0x62, 0xf0, 0x45, 0x15, 0xbb, 0xa6,
	0x62, 0xc7, 0xd0, 0xf3, 0x39, 0x46, 0x55, 0xb0, 0x5b, 0xf0, 0x75, 0x97, 0x0f, 0x7a, 0x78, 0x1f,
	0xa3, 0xa6, 0x04, 0x6f, 0xef, 0x1b, 0xc3, 0xb7, 0xc6, 0xb0, 0x48, 0x19, 0x4b, 0xde, 0xfe, 0x09,
	0xb6, 0x3a, 0x31, 0x0d, 0x4d, 0xbc, 0x97, 0x8c, 0x2c, 0xb5, 0xd9, 0xc6, 0xb6, 0x67, 0x6c, 0xcb,
	0x68, 0x63, 0xdd, 0x34, 0x6d, 0xd7, 0xa8, 0xc2, 0x40, 0xeb, 0x57, 0xdb, 0x4c, 0xdb, 0x02, 0x65,
	0x2c, 0x25, 0xbb, 0x09, 0xe5, 0x0b, 0x41, 0xa8, 0x04, 0x1f, 0xa5, 0x86, 0x99, 0xe7, 0x99, 0x23,
	0x8c, 0xbc, 0x6c, 0xdf, 0xc3, 0x57, 0x3d, 0x4e, 0x78, 0x19, 0x84, 0x41, 0xba, 0x67, 0x76, 0x0c,
	0xa0, 0x38, 0xe1, 0xed, 0xc8, 0x40, 0xc9, 0xc8, 0x36, 0x8e, 0xaa, 0xce, 0x43, 0xcb, 0x79, 0xd5,
	0x5d, 0x21, 0xf1, 0x01, 0x27, 0xde, 0x05, 0x33, 0xc9, 0x42, 0x3a, 0xbe, 0x75, 0x95, 0xd1, 0xed,
	0x3d, 0xd3, 0x58, 0x33, 0x8d, 0x0b, 0xd9, 0x46, 0xb9, 0x66, 0xef, 0x40, 0xf5, 0x44, 0x21, 0x27,
	0x9c, 0x8e, 0x2a, 0x8b, 0xb3, 0xbf, 0x83, 0xed, 0x8c, 0xd0, 0x91, 0x14, 0x1a, 0x33, 0xa6, 0x0d,
	0x26, 0xb1, 0xf0, 0x8f, 0x29, 0xb2, 0x2b, 0xc0, 0x7e, 0x46, 0x5a, 0xf4, 0x36, 0x60, 0x2b, 0x41,
	0xdf, 0x30, 0xee, 0x40, 0xf5, 0xb7, 0x68, 0xb0, 0xbc, 0x37, 0x23, 0xde, 0xb6, 0x9f, 0xe2, 0x08,
	0x97, 0xda, 0x33, 0xe2, 0x0d, 0xfb, 0xef, 0xb0, 0x75, 0x19, 0xe8, 0xc5, 0x7d, 0xb3, 0x6d, 0x28,
	0x7a, 0xb1, 0xd2, 0x52, 0xa5, 0x2f, 0x67, 0x6f, 0xba, 0x62, 0x0c, 0x56, 0x23, 0xee, 0x63, 0x6d,
	0xc5, 0xca, 0x1d, 0x16, 0x7a, 0xc9, 0x7f, 0x56, 0x81, 0x42, 0x7a, 0x28, 0xf9, 0x04, 0x4c, 0x17,
	0xf6, 0x15, 0x54, 0xd2, 0xe0, 0xf9, 0x72, 0x76, 0x00, 0x1b, 0x02, 0x1f, 0xe9, 0x76, 0x2e, 0x1e,
	0x0c, 0x74, 0x92, 0x56, 0x54, 0xa0, 0x40, 0x92, 0xf8, 0x68, 0xda, 0x91, 0x2e, 0xec, 0xff, 0x73,
	0xf0, 0xcd, 0x27, 0x27, 0xcd, 0xac, 0xac, 0xda, 0xc4, 0xe4, 0xe7, 0x0e, 0x3e, 0x25, 0xd8, 0x21,
	0xac, 0x2b, 0x0c, 0x79, 0x20, 0x02, 0xe1, 0x27, 0x89, 0xf9, 0x85, 0xd7, 0x23, 0x23, 0x59, 0x07,
	0xd6, 0x14, 0x6a, 0xa4, 0x5b, 0x9e, 0x3e, 0xc9, 0xc6, 0x51, 0xdd, 0x49, 0x2f, 0x0f, 0x27, 0xbb,
	0x3c, 0x9c, 0x9b, 0xec, 0xf2, 0x98, 0x0b, 0x29, 0x25, 0xbe, 0x0e, 0x31, 0x1b, 0x8a, 0xfd, 0xd8,
	0xfb, 0x13, 0xa9, 0xb6, 0x9a, 0x7c, 0xd2, 0xb3, 0xa2, 0x29, 0x33, 0x3b, 0xfc, 0xee, 0x73, 0xee,
	0xaf, 0xce, 0x7f, 0x39, 0x16, 0xc3, 0xda, 0xaf, 0xa4, 0xa4, 0xf0, 0x4f, 0xaf, 0xd8, 0xfe, 0xcd,
	0x10, 0xad, 0x40, 0xdc, 0x29, 0xae, 0x49, 0xc5, 0x1e, 0xc5, 0x0a, 0x2d, 0xee, 0x79, 0xa8, 0xb5,
	0xd5, 0xb9, 0xbe, 0x70, 0xec, 0xd3, 0x19, 0xad, 0x3d, 0x24, 0x8a, 0x74, 0xdb, 0x75, 0xc7, 0xe3,
	0xb1, 0xa3, 0x13, 0x74, 0x10, 0x3a, 0x9e, 0x0c, 0xdd, 0x81, 0xf4, 0xb4, 0xcb, 0xa3, 0xc0, 0xad,
	0x57, 0x74, 0x1c, 0x45, 0x52, 0xd1, 0xf1, 0x2c, 0x7f, 0x94, 0x6f, 0x39, 0x3f, 0xd4, 0x37, 0x79,
	0x14, 0xcc, 0xd9, 0xd4, 0x39, 0x34, 0x2e, 0x91, 0x2b, 0x61, 0x85, 0xd2, 0xd4, 0xf6, 0x65, 0x4c,
	0x16, 0x0d, 0xd1, 0xd2, 0xd3, 0x46, 0xb3, 0x05, 0xd6, 0xf8, 0x6c, 0x2b, 0xec, 0x7a, 0x32, 0xfc,
	0x48, 0x99, 0x9a, 0x87, 0x96, 0x13, 0x8d, 0xe2, 0xb0, 0x1f, 0x08, 0xbf, 0x5f, 0x4c, 0xc6, 0xf9,
	0xe3, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x84, 0x88, 0x0c, 0xe8, 0x05, 0x00, 0x00,
}

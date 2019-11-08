// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attachments.proto

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

// AttachmentCreateRequest specifies what kind of Attachments that should be registered in
// the organizations fleet. Note that a Attachment must be either a Relay or a
// Gateway.
type AttachmentCreateRequest struct {
	Meta                 *CreateRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Attachments          []*Attachment          `protobuf:"bytes,2,rep,name=attachments,proto3" json:"attachments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AttachmentCreateRequest) Reset()         { *m = AttachmentCreateRequest{} }
func (m *AttachmentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*AttachmentCreateRequest) ProtoMessage()    {}
func (*AttachmentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{0}
}

func (m *AttachmentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentCreateRequest.Unmarshal(m, b)
}
func (m *AttachmentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentCreateRequest.Marshal(b, m, deterministic)
}
func (m *AttachmentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentCreateRequest.Merge(m, src)
}
func (m *AttachmentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_AttachmentCreateRequest.Size(m)
}
func (m *AttachmentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentCreateRequest proto.InternalMessageInfo

func (m *AttachmentCreateRequest) GetMeta() *CreateRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentCreateRequest) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

// AttachmentCreateResponse reports how the Attachments were created in the system. It can
// communicate partial successes or failures.
type AttachmentCreateResponse struct {
	Meta                 *CreateResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Attachments          []*Attachment           `protobuf:"bytes,2,rep,name=attachments,proto3" json:"attachments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AttachmentCreateResponse) Reset()         { *m = AttachmentCreateResponse{} }
func (m *AttachmentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*AttachmentCreateResponse) ProtoMessage()    {}
func (*AttachmentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{1}
}

func (m *AttachmentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentCreateResponse.Unmarshal(m, b)
}
func (m *AttachmentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentCreateResponse.Marshal(b, m, deterministic)
}
func (m *AttachmentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentCreateResponse.Merge(m, src)
}
func (m *AttachmentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_AttachmentCreateResponse.Size(m)
}
func (m *AttachmentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentCreateResponse proto.InternalMessageInfo

func (m *AttachmentCreateResponse) GetMeta() *CreateResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentCreateResponse) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

// AttachmentGetRequest specifies which Attachment to retrieve.
type AttachmentGetRequest struct {
	Meta                 *GetRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Id                   string              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AttachmentGetRequest) Reset()         { *m = AttachmentGetRequest{} }
func (m *AttachmentGetRequest) String() string { return proto.CompactTextString(m) }
func (*AttachmentGetRequest) ProtoMessage()    {}
func (*AttachmentGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{2}
}

func (m *AttachmentGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentGetRequest.Unmarshal(m, b)
}
func (m *AttachmentGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentGetRequest.Marshal(b, m, deterministic)
}
func (m *AttachmentGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentGetRequest.Merge(m, src)
}
func (m *AttachmentGetRequest) XXX_Size() int {
	return xxx_messageInfo_AttachmentGetRequest.Size(m)
}
func (m *AttachmentGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentGetRequest proto.InternalMessageInfo

func (m *AttachmentGetRequest) GetMeta() *GetRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AttachmentGetResponse returns a requested Attachment.
type AttachmentGetResponse struct {
	Meta                 *GetResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Attachment           *Attachment          `protobuf:"bytes,2,opt,name=attachment,proto3" json:"attachment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AttachmentGetResponse) Reset()         { *m = AttachmentGetResponse{} }
func (m *AttachmentGetResponse) String() string { return proto.CompactTextString(m) }
func (*AttachmentGetResponse) ProtoMessage()    {}
func (*AttachmentGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{3}
}

func (m *AttachmentGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentGetResponse.Unmarshal(m, b)
}
func (m *AttachmentGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentGetResponse.Marshal(b, m, deterministic)
}
func (m *AttachmentGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentGetResponse.Merge(m, src)
}
func (m *AttachmentGetResponse) XXX_Size() int {
	return xxx_messageInfo_AttachmentGetResponse.Size(m)
}
func (m *AttachmentGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentGetResponse proto.InternalMessageInfo

func (m *AttachmentGetResponse) GetMeta() *GetResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentGetResponse) GetAttachment() *Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

// AttachmentDeleteRequest identifies a Attachment by ID to delete.
type AttachmentDeleteRequest struct {
	Meta                 *DeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Id                   string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AttachmentDeleteRequest) Reset()         { *m = AttachmentDeleteRequest{} }
func (m *AttachmentDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AttachmentDeleteRequest) ProtoMessage()    {}
func (*AttachmentDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{4}
}

func (m *AttachmentDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentDeleteRequest.Unmarshal(m, b)
}
func (m *AttachmentDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AttachmentDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentDeleteRequest.Merge(m, src)
}
func (m *AttachmentDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AttachmentDeleteRequest.Size(m)
}
func (m *AttachmentDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentDeleteRequest proto.InternalMessageInfo

func (m *AttachmentDeleteRequest) GetMeta() *DeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentDeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// AttachmentDeleteResponse returns information about a Attachment that was deleted.
type AttachmentDeleteResponse struct {
	Meta                 *DeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AttachmentDeleteResponse) Reset()         { *m = AttachmentDeleteResponse{} }
func (m *AttachmentDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AttachmentDeleteResponse) ProtoMessage()    {}
func (*AttachmentDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{5}
}

func (m *AttachmentDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentDeleteResponse.Unmarshal(m, b)
}
func (m *AttachmentDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AttachmentDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentDeleteResponse.Merge(m, src)
}
func (m *AttachmentDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AttachmentDeleteResponse.Size(m)
}
func (m *AttachmentDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentDeleteResponse proto.InternalMessageInfo

func (m *AttachmentDeleteResponse) GetMeta() *DeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

// AttachmentListRequest specifies criteria for retrieving a list of Attachments.
type AttachmentListRequest struct {
	Meta                 *ListRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Filter               string               `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AttachmentListRequest) Reset()         { *m = AttachmentListRequest{} }
func (m *AttachmentListRequest) String() string { return proto.CompactTextString(m) }
func (*AttachmentListRequest) ProtoMessage()    {}
func (*AttachmentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{6}
}

func (m *AttachmentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentListRequest.Unmarshal(m, b)
}
func (m *AttachmentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentListRequest.Marshal(b, m, deterministic)
}
func (m *AttachmentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentListRequest.Merge(m, src)
}
func (m *AttachmentListRequest) XXX_Size() int {
	return xxx_messageInfo_AttachmentListRequest.Size(m)
}
func (m *AttachmentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentListRequest proto.InternalMessageInfo

func (m *AttachmentListRequest) GetMeta() *ListRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// AttachmentListResponse returns a list of Attachments that meet the criteria of a
// AttachmentListRequest.
type AttachmentListResponse struct {
	Meta                 *ListResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Attachments          []*Attachment         `protobuf:"bytes,2,rep,name=attachments,proto3" json:"attachments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AttachmentListResponse) Reset()         { *m = AttachmentListResponse{} }
func (m *AttachmentListResponse) String() string { return proto.CompactTextString(m) }
func (*AttachmentListResponse) ProtoMessage()    {}
func (*AttachmentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{7}
}

func (m *AttachmentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentListResponse.Unmarshal(m, b)
}
func (m *AttachmentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentListResponse.Marshal(b, m, deterministic)
}
func (m *AttachmentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentListResponse.Merge(m, src)
}
func (m *AttachmentListResponse) XXX_Size() int {
	return xxx_messageInfo_AttachmentListResponse.Size(m)
}
func (m *AttachmentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentListResponse proto.InternalMessageInfo

func (m *AttachmentListResponse) GetMeta() *ListResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentListResponse) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

// AttachmentBatchDeleteRequest specifies a list of Attachments to delete by ID.
type AttachmentBatchDeleteRequest struct {
	Meta                 *BatchDeleteRequestMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Ids                  []string                    `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AttachmentBatchDeleteRequest) Reset()         { *m = AttachmentBatchDeleteRequest{} }
func (m *AttachmentBatchDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*AttachmentBatchDeleteRequest) ProtoMessage()    {}
func (*AttachmentBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{8}
}

func (m *AttachmentBatchDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentBatchDeleteRequest.Unmarshal(m, b)
}
func (m *AttachmentBatchDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentBatchDeleteRequest.Marshal(b, m, deterministic)
}
func (m *AttachmentBatchDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentBatchDeleteRequest.Merge(m, src)
}
func (m *AttachmentBatchDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_AttachmentBatchDeleteRequest.Size(m)
}
func (m *AttachmentBatchDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentBatchDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentBatchDeleteRequest proto.InternalMessageInfo

func (m *AttachmentBatchDeleteRequest) GetMeta() *BatchDeleteRequestMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *AttachmentBatchDeleteRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

// AttachmentBatchDeleteResponse returns information about Attachments deleted via a
// AttachmentBatchDeleteRequest.
type AttachmentBatchDeleteResponse struct {
	Meta                 *BatchDeleteResponseMetadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AttachmentBatchDeleteResponse) Reset()         { *m = AttachmentBatchDeleteResponse{} }
func (m *AttachmentBatchDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*AttachmentBatchDeleteResponse) ProtoMessage()    {}
func (*AttachmentBatchDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{9}
}

func (m *AttachmentBatchDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachmentBatchDeleteResponse.Unmarshal(m, b)
}
func (m *AttachmentBatchDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachmentBatchDeleteResponse.Marshal(b, m, deterministic)
}
func (m *AttachmentBatchDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachmentBatchDeleteResponse.Merge(m, src)
}
func (m *AttachmentBatchDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_AttachmentBatchDeleteResponse.Size(m)
}
func (m *AttachmentBatchDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachmentBatchDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttachmentBatchDeleteResponse proto.InternalMessageInfo

func (m *AttachmentBatchDeleteResponse) GetMeta() *BatchDeleteResponseMetadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

// Attachment is a domain object --
type Attachment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	CompositeRoleId      string   `protobuf:"bytes,3,opt,name=composite_role_id,json=compositeRoleId,proto3" json:"composite_role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a97d0ef1337fe45, []int{10}
}

func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (m *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(m, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Attachment) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *Attachment) GetCompositeRoleId() string {
	if m != nil {
		return m.CompositeRoleId
	}
	return ""
}

func init() {
	proto.RegisterType((*AttachmentCreateRequest)(nil), "v1.AttachmentCreateRequest")
	proto.RegisterType((*AttachmentCreateResponse)(nil), "v1.AttachmentCreateResponse")
	proto.RegisterType((*AttachmentGetRequest)(nil), "v1.AttachmentGetRequest")
	proto.RegisterType((*AttachmentGetResponse)(nil), "v1.AttachmentGetResponse")
	proto.RegisterType((*AttachmentDeleteRequest)(nil), "v1.AttachmentDeleteRequest")
	proto.RegisterType((*AttachmentDeleteResponse)(nil), "v1.AttachmentDeleteResponse")
	proto.RegisterType((*AttachmentListRequest)(nil), "v1.AttachmentListRequest")
	proto.RegisterType((*AttachmentListResponse)(nil), "v1.AttachmentListResponse")
	proto.RegisterType((*AttachmentBatchDeleteRequest)(nil), "v1.AttachmentBatchDeleteRequest")
	proto.RegisterType((*AttachmentBatchDeleteResponse)(nil), "v1.AttachmentBatchDeleteResponse")
	proto.RegisterType((*Attachment)(nil), "v1.Attachment")
}

func init() { proto.RegisterFile("attachments.proto", fileDescriptor_1a97d0ef1337fe45) }

var fileDescriptor_1a97d0ef1337fe45 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xd3, 0x48,
	0x14, 0x96, 0x9d, 0x6e, 0xa2, 0x7d, 0xd9, 0x6d, 0x36, 0xa3, 0xfe, 0x48, 0xbd, 0xd9, 0xae, 0x3b,
	0xbd, 0x54, 0xd9, 0x26, 0x56, 0xb3, 0x5a, 0x75, 0x37, 0xea, 0x82, 0x42, 0x2b, 0x55, 0x45, 0xe5,
	0x92, 0x03, 0xe2, 0x44, 0x35, 0xb5, 0x87, 0xc4, 0x6a, 0xe2, 0x31, 0xf6, 0x90, 0x1c, 0xaa, 0x72,
	0xe0, 0x80, 0xc4, 0xb5, 0x08, 0xc1, 0x11, 0x71, 0xe2, 0x88, 0xd4, 0x0b, 0x77, 0xc4, 0x5f, 0x80,
	0xf8, 0x0f, 0xf8, 0x0b, 0xe8, 0x05, 0x71, 0x42, 0x63, 0x3b, 0xb1, 0xc7, 0x49, 0x53, 0x2a, 0x71,
	0x4a, 0x32, 0xef, 0x7b, 0xef, 0xfb, 0xbe, 0x37, 0x6f, 0x66, 0x02, 0x45, 0xc2, 0x39, 0x31, 0x3b,
	0x3d, 0xea, 0x70, 0xbf, 0xe6, 0x7a, 0x8c, 0x33, 0xa4, 0xf6, 0x37, 0xb4, 0x72, 0x9b, 0xb1, 0x76,
	0x97, 0x1a, 0xc4, 0xb5, 0x0d, 0xe2, 0x38, 0x8c, 0x13, 0x6e, 0x33, 0x27, 0x42, 0x68, 0xeb, 0xc1,
	0x87, 0x59, 0x6d, 0x53, 0xa7, 0xea, 0x0f, 0x48, 0xbb, 0x4d, 0x3d, 0x83, 0xb9, 0x01, 0x62, 0x02,
	0xfa, 0xd7, 0x28, 0x14, 0xfd, 0x04, 0xdf, 0xa5, 0x66, 0xf8, 0x1d, 0x3f, 0x56, 0x60, 0xb1, 0x39,
	0x12, 0xb0, 0xed, 0x51, 0xc2, 0x69, 0x8b, 0xde, 0x7f, 0x40, 0x7d, 0x8e, 0xaa, 0x30, 0xd3, 0xa3,
	0x9c, 0x94, 0x14, 0x5d, 0x59, 0xcb, 0xd7, 0x97, 0x6a, 0xfd, 0x8d, 0x9a, 0x04, 0xb8, 0x45, 0x39,
	0xb1, 0x08, 0x27, 0xad, 0x00, 0x86, 0xb6, 0x20, 0x9f, 0xb0, 0x52, 0x52, 0xf5, 0xcc, 0x5a, 0xbe,
	0x3e, 0x2b, 0xb2, 0x62, 0x82, 0x1b, 0xf0, 0xf9, 0xcb, 0x59, 0xee, 0xa7, 0x37, 0xe7, 0x67, 0x39,
	0xa5, 0x95, 0x84, 0xe3, 0x97, 0x0a, 0x94, 0xc6, 0x85, 0xf8, 0x2e, 0x73, 0x7c, 0x8a, 0x1a, 0x92,
	0x12, 0x2d, 0xa9, 0x24, 0x44, 0x0c, 0xa5, 0x48, 0xf5, 0x7f, 0x80, 0xac, 0x06, 0x7c, 0x15, 0x81,
	0xd7, 0x22, 0x80, 0xef, 0xc2, 0x5c, 0x9c, 0xb2, 0x4b, 0xf9, 0xb0, 0x4f, 0x15, 0x49, 0xdd, 0x82,
	0x28, 0x1d, 0x47, 0x53, 0x4d, 0xd2, 0x40, 0xb5, 0xad, 0x92, 0xaa, 0x2b, 0x6b, 0x3f, 0x4b, 0xa4,
	0xaa, 0x6d, 0xe1, 0x17, 0x0a, 0xcc, 0xa7, 0x08, 0x22, 0xff, 0x9b, 0x12, 0xc3, 0xe2, 0x88, 0xe1,
	0x52, 0xf3, 0x0d, 0x80, 0xd8, 0x4d, 0x40, 0x3b, 0xdd, 0x7b, 0x02, 0x2d, 0x59, 0xb7, 0x92, 0x53,
	0xb2, 0x43, 0xbb, 0x74, 0xea, 0x94, 0x48, 0x80, 0x2b, 0x34, 0xe0, 0x30, 0x39, 0x02, 0xc3, 0x22,
	0x17, 0x8f, 0x80, 0x8c, 0xb8, 0xb8, 0x0b, 0x92, 0x93, 0x4e, 0xb2, 0xc7, 0xfb, 0xb6, 0x3f, 0xda,
	0xc5, 0xbf, 0xc6, 0x7b, 0x9c, 0x08, 0xa7, 0x5c, 0x60, 0xc8, 0xde, 0xb3, 0xbb, 0x9c, 0x7a, 0x13,
	0x9c, 0x44, 0x11, 0xfc, 0x4c, 0x81, 0x85, 0x34, 0x55, 0x64, 0x66, 0x5d, 0xe2, 0x2a, 0xc5, 0x5c,
	0xb2, 0x95, 0x88, 0xac, 0xf9, 0x3d, 0x13, 0x5c, 0x10, 0x0a, 0x20, 0x50, 0xf0, 0x76, 0xfa, 0x18,
	0xbb, 0x50, 0x4e, 0xe4, 0x11, 0x6e, 0x76, 0xe4, 0x0d, 0xad, 0x4b, 0xe2, 0x96, 0x05, 0xcf, 0x38,
	0x2a, 0x25, 0xb1, 0x0c, 0x19, 0xdb, 0x0a, 0xa5, 0xc9, 0xcd, 0x10, 0xcb, 0xb8, 0x0b, 0x7f, 0x5c,
	0xc0, 0x18, 0xf5, 0xe3, 0xba, 0x44, 0xf9, 0xe7, 0x18, 0xe5, 0x95, 0x76, 0xf8, 0xbd, 0x0a, 0x10,
	0xd3, 0xa1, 0x95, 0x60, 0xe0, 0x94, 0x60, 0x9b, 0x8a, 0x22, 0xf1, 0x97, 0x57, 0xe7, 0x67, 0x39,
	0x75, 0x6f, 0x67, 0x34, 0x77, 0xa8, 0x06, 0x39, 0x8f, 0x75, 0xe9, 0xc1, 0x68, 0x30, 0xe7, 0x05,
	0xee, 0x37, 0x81, 0xcb, 0xb6, 0x58, 0x97, 0x0e, 0xb1, 0x59, 0x81, 0xda, 0xb3, 0xd0, 0x4d, 0x28,
	0x9a, 0xac, 0xe7, 0x32, 0xdf, 0xe6, 0xf4, 0x60, 0x98, 0x99, 0x09, 0x32, 0x97, 0x45, 0xe6, 0x92,
	0xc8, 0x2c, 0x6c, 0x0f, 0x21, 0xc9, 0x12, 0x05, 0x53, 0x5a, 0xb5, 0x1a, 0xcf, 0x95, 0xd3, 0xe6,
	0x13, 0xa5, 0xb2, 0x0b, 0xb3, 0x4d, 0x47, 0x8f, 0x65, 0xd7, 0xd0, 0x3f, 0x1d, 0xce, 0x5d, 0xbf,
	0x61, 0x18, 0x83, 0xc1, 0xa0, 0xe6, 0x73, 0x8f, 0x39, 0x6d, 0xab, 0x57, 0x33, 0x59, 0xcf, 0xb0,
	0x98, 0xe9, 0x07, 0xef, 0x02, 0x75, 0xb8, 0xcd, 0x6d, 0xea, 0xaf, 0xc6, 0x79, 0xf5, 0x6b, 0x68,
	0xeb, 0x58, 0xc7, 0xb6, 0x85, 0x1b, 0x3a, 0x26, 0xd5, 0x4d, 0xbc, 0xae, 0xe3, 0x48, 0xa1, 0x58,
	0xf1, 0xaa, 0xff, 0x8a, 0x95, 0x31, 0xf5, 0x61, 0xec, 0x3f, 0x7c, 0x12, 0xf7, 0xb1, 0xfe, 0x71,
	0x06, 0xf2, 0x71, 0x61, 0x1f, 0xbd, 0x53, 0x20, 0x1b, 0xde, 0xba, 0xe8, 0x77, 0x79, 0xf8, 0xa4,
	0x57, 0x41, 0x2b, 0x4f, 0x0e, 0x86, 0x7b, 0x88, 0x1f, 0x9e, 0x36, 0x09, 0x3e, 0x80, 0xe5, 0x7d,
	0x4a, 0x3c, 0x47, 0xef, 0xb0, 0x81, 0xce, 0x99, 0xde, 0x23, 0x47, 0x54, 0x27, 0x09, 0xfb, 0xe8,
	0xff, 0xcb, 0xdd, 0xfb, 0xd4, 0xeb, 0xdb, 0x26, 0xf5, 0x8d, 0x84, 0xc8, 0xd5, 0x90, 0xed, 0xd1,
	0x87, 0x4f, 0x4f, 0xd5, 0x39, 0x5c, 0x30, 0xfa, 0x1b, 0x46, 0xf2, 0x18, 0x28, 0x15, 0x74, 0x07,
	0x32, 0xbb, 0x94, 0xa3, 0x92, 0x2c, 0x32, 0xbe, 0xaf, 0xb5, 0xa5, 0x09, 0x91, 0x48, 0x7b, 0x39,
	0xa8, 0xbb, 0x80, 0xe6, 0x52, 0x75, 0x8d, 0x63, 0xdb, 0x3a, 0x41, 0x04, 0xb2, 0xe1, 0xbc, 0xa6,
	0xdb, 0x23, 0x1d, 0x9c, 0x74, 0x7b, 0xe4, 0x11, 0x1f, 0x52, 0x54, 0x26, 0x53, 0xdc, 0x86, 0x19,
	0x71, 0x4f, 0xa0, 0x94, 0xc6, 0xc4, 0x3d, 0xa5, 0x69, 0x93, 0x42, 0x51, 0xf1, 0xc5, 0xa0, 0x78,
	0x11, 0xa5, 0xfb, 0x82, 0x8e, 0x20, 0x9f, 0x38, 0x6f, 0x48, 0x4f, 0xdd, 0x2d, 0x63, 0xa7, 0x5f,
	0x5b, 0x99, 0x82, 0x90, 0xc9, 0x2a, 0x69, 0xb2, 0xc3, 0x6c, 0xf0, 0xcf, 0xe3, 0xef, 0x6f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0xdf, 0x40, 0x23, 0xf9, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AttachmentsClient is the client API for Attachments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AttachmentsClient interface {
	// Create registers a new attachment.
	Create(ctx context.Context, in *AttachmentCreateRequest, opts ...grpc.CallOption) (*AttachmentCreateResponse, error)
	// Get reads one attachment by ID.
	Get(ctx context.Context, in *AttachmentGetRequest, opts ...grpc.CallOption) (*AttachmentGetResponse, error)
	// Delete removes a Attachment by ID.
	Delete(ctx context.Context, in *AttachmentDeleteRequest, opts ...grpc.CallOption) (*AttachmentDeleteResponse, error)
	// List is a batched Get call.
	List(ctx context.Context, in *AttachmentListRequest, opts ...grpc.CallOption) (*AttachmentListResponse, error)
	// BatchDelete is a batched Delete call.
	BatchDelete(ctx context.Context, in *AttachmentBatchDeleteRequest, opts ...grpc.CallOption) (*AttachmentBatchDeleteResponse, error)
}

type attachmentsClient struct {
	cc *grpc.ClientConn
}

func NewAttachmentsClient(cc *grpc.ClientConn) AttachmentsClient {
	return &attachmentsClient{cc}
}

func (c *attachmentsClient) Create(ctx context.Context, in *AttachmentCreateRequest, opts ...grpc.CallOption) (*AttachmentCreateResponse, error) {
	out := new(AttachmentCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.Attachments/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) Get(ctx context.Context, in *AttachmentGetRequest, opts ...grpc.CallOption) (*AttachmentGetResponse, error) {
	out := new(AttachmentGetResponse)
	err := c.cc.Invoke(ctx, "/v1.Attachments/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) Delete(ctx context.Context, in *AttachmentDeleteRequest, opts ...grpc.CallOption) (*AttachmentDeleteResponse, error) {
	out := new(AttachmentDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.Attachments/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) List(ctx context.Context, in *AttachmentListRequest, opts ...grpc.CallOption) (*AttachmentListResponse, error) {
	out := new(AttachmentListResponse)
	err := c.cc.Invoke(ctx, "/v1.Attachments/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attachmentsClient) BatchDelete(ctx context.Context, in *AttachmentBatchDeleteRequest, opts ...grpc.CallOption) (*AttachmentBatchDeleteResponse, error) {
	out := new(AttachmentBatchDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.Attachments/BatchDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttachmentsServer is the server API for Attachments service.
type AttachmentsServer interface {
	// Create registers a new attachment.
	Create(context.Context, *AttachmentCreateRequest) (*AttachmentCreateResponse, error)
	// Get reads one attachment by ID.
	Get(context.Context, *AttachmentGetRequest) (*AttachmentGetResponse, error)
	// Delete removes a Attachment by ID.
	Delete(context.Context, *AttachmentDeleteRequest) (*AttachmentDeleteResponse, error)
	// List is a batched Get call.
	List(context.Context, *AttachmentListRequest) (*AttachmentListResponse, error)
	// BatchDelete is a batched Delete call.
	BatchDelete(context.Context, *AttachmentBatchDeleteRequest) (*AttachmentBatchDeleteResponse, error)
}

// UnimplementedAttachmentsServer can be embedded to have forward compatible implementations.
type UnimplementedAttachmentsServer struct {
}

func (*UnimplementedAttachmentsServer) Create(ctx context.Context, req *AttachmentCreateRequest) (*AttachmentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAttachmentsServer) Get(ctx context.Context, req *AttachmentGetRequest) (*AttachmentGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAttachmentsServer) Delete(ctx context.Context, req *AttachmentDeleteRequest) (*AttachmentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAttachmentsServer) List(ctx context.Context, req *AttachmentListRequest) (*AttachmentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAttachmentsServer) BatchDelete(ctx context.Context, req *AttachmentBatchDeleteRequest) (*AttachmentBatchDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}

func RegisterAttachmentsServer(s *grpc.Server, srv AttachmentsServer) {
	s.RegisterService(&_Attachments_serviceDesc, srv)
}

func _Attachments_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Attachments/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).Create(ctx, req.(*AttachmentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Attachments/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).Get(ctx, req.(*AttachmentGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Attachments/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).Delete(ctx, req.(*AttachmentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Attachments/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).List(ctx, req.(*AttachmentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Attachments_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachmentBatchDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttachmentsServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Attachments/BatchDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttachmentsServer).BatchDelete(ctx, req.(*AttachmentBatchDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Attachments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Attachments",
	HandlerType: (*AttachmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Attachments_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Attachments_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Attachments_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Attachments_List_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _Attachments_BatchDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attachments.proto",
}

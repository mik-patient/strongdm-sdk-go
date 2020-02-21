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
// source: options.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type FieldOptions struct {
	Name              string `protobuf:"bytes,1941300,opt,name=name,proto3" json:"name,omitempty"`
	SqlNullable       bool   `protobuf:"varint,1941301,opt,name=sql_nullable,json=sqlNullable,proto3" json:"sql_nullable,omitempty"`
	ExposeAsPorcelain bool   `protobuf:"varint,1941302,opt,name=expose_as_porcelain,json=exposeAsPorcelain,proto3" json:"expose_as_porcelain,omitempty"`
	// The iterable flag is only valid for repeated fields on RPC output types.
	// If true, the generated RPC method will return the contents of this field
	// as an iterable/generator object. Only one field can have this flag set in
	// the RPC output type.
	Iterable             bool     `protobuf:"varint,1941303,opt,name=iterable,proto3" json:"iterable,omitempty"`
	Required             bool     `protobuf:"varint,1941304,opt,name=required,proto3" json:"required,omitempty"`
	IdType               string   `protobuf:"bytes,1941305,opt,name=id_type,json=idType,proto3" json:"id_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldOptions) Reset()         { *m = FieldOptions{} }
func (m *FieldOptions) String() string { return proto.CompactTextString(m) }
func (*FieldOptions) ProtoMessage()    {}
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *FieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOptions.Unmarshal(m, b)
}
func (m *FieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOptions.Marshal(b, m, deterministic)
}
func (m *FieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOptions.Merge(m, src)
}
func (m *FieldOptions) XXX_Size() int {
	return xxx_messageInfo_FieldOptions.Size(m)
}
func (m *FieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOptions proto.InternalMessageInfo

func (m *FieldOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldOptions) GetSqlNullable() bool {
	if m != nil {
		return m.SqlNullable
	}
	return false
}

func (m *FieldOptions) GetExposeAsPorcelain() bool {
	if m != nil {
		return m.ExposeAsPorcelain
	}
	return false
}

func (m *FieldOptions) GetIterable() bool {
	if m != nil {
		return m.Iterable
	}
	return false
}

func (m *FieldOptions) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *FieldOptions) GetIdType() string {
	if m != nil {
		return m.IdType
	}
	return ""
}

type MessageOptions struct {
	ModelName string `protobuf:"bytes,1941300,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Porcelain bool   `protobuf:"varint,1941301,opt,name=porcelain,proto3" json:"porcelain,omitempty"`
	// Non-zero if the message is an error type. This corresponds to the gRPC status code.
	Error int32 `protobuf:"varint,1941302,opt,name=error,proto3" json:"error,omitempty"`
	// Set this option on an RPC request message to specify which field holds
	// the "options" for that RPC method.
	OptionsField         string   `protobuf:"bytes,1941303,opt,name=options_field,json=optionsField,proto3" json:"options_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOptions) Reset()         { *m = MessageOptions{} }
func (m *MessageOptions) String() string { return proto.CompactTextString(m) }
func (*MessageOptions) ProtoMessage()    {}
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *MessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOptions.Unmarshal(m, b)
}
func (m *MessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOptions.Marshal(b, m, deterministic)
}
func (m *MessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOptions.Merge(m, src)
}
func (m *MessageOptions) XXX_Size() int {
	return xxx_messageInfo_MessageOptions.Size(m)
}
func (m *MessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOptions proto.InternalMessageInfo

func (m *MessageOptions) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *MessageOptions) GetPorcelain() bool {
	if m != nil {
		return m.Porcelain
	}
	return false
}

func (m *MessageOptions) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *MessageOptions) GetOptionsField() string {
	if m != nil {
		return m.OptionsField
	}
	return ""
}

type OneofOptions struct {
	ModelName            string   `protobuf:"bytes,1941380,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneofOptions) Reset()         { *m = OneofOptions{} }
func (m *OneofOptions) String() string { return proto.CompactTextString(m) }
func (*OneofOptions) ProtoMessage()    {}
func (*OneofOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *OneofOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneofOptions.Unmarshal(m, b)
}
func (m *OneofOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneofOptions.Marshal(b, m, deterministic)
}
func (m *OneofOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneofOptions.Merge(m, src)
}
func (m *OneofOptions) XXX_Size() int {
	return xxx_messageInfo_OneofOptions.Size(m)
}
func (m *OneofOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_OneofOptions.DiscardUnknown(m)
}

var xxx_messageInfo_OneofOptions proto.InternalMessageInfo

func (m *OneofOptions) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

type ServiceOptions struct {
	MainNoun             string   `protobuf:"bytes,1941400,opt,name=main_noun,json=mainNoun,proto3" json:"main_noun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceOptions) Reset()         { *m = ServiceOptions{} }
func (m *ServiceOptions) String() string { return proto.CompactTextString(m) }
func (*ServiceOptions) ProtoMessage()    {}
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{3}
}

func (m *ServiceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceOptions.Unmarshal(m, b)
}
func (m *ServiceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceOptions.Marshal(b, m, deterministic)
}
func (m *ServiceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceOptions.Merge(m, src)
}
func (m *ServiceOptions) XXX_Size() int {
	return xxx_messageInfo_ServiceOptions.Size(m)
}
func (m *ServiceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceOptions proto.InternalMessageInfo

func (m *ServiceOptions) GetMainNoun() string {
	if m != nil {
		return m.MainNoun
	}
	return ""
}

var E_FieldOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldOptions)(nil),
	Field:         1941390,
	Name:          "v1.field_options",
	Tag:           "bytes,1941390,opt,name=field_options",
	Filename:      "options.proto",
}

var E_MessageOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MessageOptions)(nil),
	Field:         1941391,
	Name:          "v1.message_options",
	Tag:           "bytes,1941391,opt,name=message_options",
	Filename:      "options.proto",
}

var E_OneofOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.OneofOptions)(nil),
	ExtensionType: (*OneofOptions)(nil),
	Field:         1941381,
	Name:          "v1.oneof_options",
	Tag:           "bytes,1941381,opt,name=oneof_options",
	Filename:      "options.proto",
}

var E_ServiceOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ServiceOptions)(nil),
	Field:         1941401,
	Name:          "v1.service_options",
	Tag:           "bytes,1941401,opt,name=service_options",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterType((*FieldOptions)(nil), "v1.FieldOptions")
	proto.RegisterType((*MessageOptions)(nil), "v1.MessageOptions")
	proto.RegisterType((*OneofOptions)(nil), "v1.OneofOptions")
	proto.RegisterType((*ServiceOptions)(nil), "v1.ServiceOptions")
	proto.RegisterExtension(E_FieldOptions)
	proto.RegisterExtension(E_MessageOptions)
	proto.RegisterExtension(E_OneofOptions)
	proto.RegisterExtension(E_ServiceOptions)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x10, 0x2d, 0xc9, 0xd4, 0x0d, 0xb0, 0x20, 0xb4, 0xa0, 0xd2, 0x56, 0x41, 0x48, 0x9c,
	0x1c, 0x0a, 0xb7, 0xde, 0xe0, 0xc0, 0x8d, 0x14, 0x19, 0xc4, 0x11, 0xcb, 0x89, 0x27, 0xd6, 0x4a,
	0xeb, 0x9d, 0xcd, 0xae, 0x1d, 0xd1, 0x3b, 0x5c, 0x81, 0x23, 0xfc, 0x0f, 0x3e, 0x2f, 0xf0, 0x1f,
	0xf8, 0x45, 0x68, 0xd7, 0xeb, 0xd4, 0x06, 0xd4, 0xe3, 0xbc, 0x37, 0xfb, 0xf6, 0xbd, 0x37, 0xb0,
	0x4b, 0xba, 0x12, 0xa4, 0x6c, 0xac, 0x0d, 0x55, 0xc4, 0x2e, 0xac, 0x8f, 0x6e, 0x1d, 0x16, 0x44,
	0x85, 0xc4, 0xa9, 0x47, 0xe6, 0xf5, 0x72, 0x9a, 0xa3, 0x5d, 0x18, 0xa1, 0x2b, 0x32, 0xcd, 0xd6,
	0xe4, 0xf7, 0x00, 0xa2, 0x27, 0x02, 0x65, 0x7e, 0xd2, 0x3c, 0x66, 0xd7, 0xe1, 0xa2, 0xca, 0x4a,
	0xe4, 0x9f, 0x7f, 0xae, 0x0f, 0x07, 0xf7, 0x46, 0x89, 0x9f, 0xd8, 0x1d, 0x88, 0xec, 0x4a, 0xa6,
	0xaa, 0x96, 0x32, 0x9b, 0x4b, 0xe4, 0x5f, 0x3c, 0x3b, 0x4c, 0x76, 0xec, 0x4a, 0xce, 0x02, 0xc8,
	0xee, 0xc3, 0x35, 0x7c, 0xad, 0xc9, 0x62, 0x9a, 0xd9, 0x54, 0x93, 0x59, 0xa0, 0xcc, 0x84, 0xe2,
	0x5f, 0xc3, 0xee, 0xd5, 0x86, 0x7c, 0x64, 0x9f, 0xb5, 0x14, 0xdb, 0x83, 0xa1, 0xa8, 0xd0, 0x78,
	0xc9, 0x6f, 0x61, 0x6d, 0x83, 0x38, 0xd6, 0xe0, 0xaa, 0x16, 0x06, 0x73, 0xfe, 0xbd, 0x65, 0x5b,
	0x84, 0xdd, 0x84, 0x4b, 0x22, 0x4f, 0xab, 0x53, 0x8d, 0xfc, 0x47, 0xf0, 0xba, 0x2d, 0xf2, 0x17,
	0xa7, 0x1a, 0x27, 0x1f, 0x06, 0x30, 0x7e, 0x8a, 0xd6, 0x66, 0x05, 0xb6, 0xb1, 0x0e, 0x00, 0x4a,
	0xca, 0x51, 0xa6, 0xbd, 0x70, 0x23, 0x8f, 0xcd, 0x5c, 0xc2, 0x7d, 0x18, 0x9d, 0x59, 0x6e, 0xe3,
	0x9d, 0x41, 0xec, 0x06, 0x6c, 0xa1, 0x31, 0x64, 0x42, 0x9c, 0xad, 0xa4, 0x19, 0xd9, 0xdd, 0x4d,
	0xef, 0xe9, 0xd2, 0xf5, 0x18, 0x72, 0x8c, 0x92, 0x28, 0xc0, 0xbe, 0xdd, 0xc9, 0x14, 0xa2, 0x13,
	0x85, 0xb4, 0xfc, 0xbf, 0x9f, 0x37, 0xbf, 0xfe, 0xf6, 0x33, 0x99, 0xc2, 0xf8, 0x39, 0x9a, 0xb5,
	0x58, 0x6c, 0x22, 0xdc, 0x86, 0x51, 0x99, 0x09, 0x95, 0x2a, 0xaa, 0x15, 0xff, 0x18, 0x5e, 0x0c,
	0x1d, 0x34, 0xa3, 0x5a, 0x1d, 0xbf, 0x84, 0x5d, 0x6f, 0x20, 0xa5, 0x76, 0x3f, 0x6e, 0xae, 0x1f,
	0xb7, 0xd7, 0x8f, 0xbb, 0x87, 0xe6, 0xef, 0xbc, 0xc2, 0xce, 0x83, 0x2b, 0xf1, 0xfa, 0xa8, 0xc7,
	0x24, 0xd1, 0xb2, 0x33, 0x1d, 0xbf, 0x82, 0xcb, 0x65, 0xd3, 0xe5, 0x46, 0xf9, 0xe0, 0x1f, 0xe5,
	0x7e, 0xdb, 0xfc, 0x7d, 0xd0, 0x66, 0x4e, 0xbb, 0xcf, 0x25, 0xe3, 0xb2, 0x37, 0x3b, 0xdf, 0xe4,
	0x9a, 0x39, 0xc7, 0x77, 0xb7, 0x39, 0xfe, 0xb6, 0xeb, 0xbb, 0xcb, 0x24, 0x11, 0x75, 0x26, 0xe7,
	0xdb, 0x36, 0x05, 0x9e, 0xe3, 0xbb, 0x5f, 0x31, 0xff, 0xd4, 0xf5, 0xdd, 0xe7, 0x92, 0xb1, 0xed,
	0xcd, 0x8f, 0xf7, 0x61, 0x6f, 0x41, 0x65, 0x6c, 0x2b, 0x43, 0xaa, 0xc8, 0xcb, 0x38, 0xd3, 0xc2,
	0xbd, 0xd2, 0xb2, 0x2e, 0xe7, 0x42, 0x15, 0xf3, 0x6d, 0xff, 0xc9, 0xc3, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x56, 0xd8, 0xad, 0x97, 0x03, 0x00, 0x00,
}

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
	OptionsField         string         `protobuf:"bytes,1941303,opt,name=options_field,json=optionsField,proto3" json:"options_field,omitempty"`
	TerraformDocs        *TerraformDocs `protobuf:"bytes,1941304,opt,name=terraform_docs,json=terraformDocs,proto3" json:"terraform_docs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *MessageOptions) GetTerraformDocs() *TerraformDocs {
	if m != nil {
		return m.TerraformDocs
	}
	return nil
}

type TerraformDocs struct {
	ResourceExamplePath   string   `protobuf:"bytes,1941300,opt,name=resource_example_path,json=resourceExamplePath,proto3" json:"resource_example_path,omitempty"`
	DataSourceExamplePath string   `protobuf:"bytes,1941301,opt,name=data_source_example_path,json=dataSourceExamplePath,proto3" json:"data_source_example_path,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *TerraformDocs) Reset()         { *m = TerraformDocs{} }
func (m *TerraformDocs) String() string { return proto.CompactTextString(m) }
func (*TerraformDocs) ProtoMessage()    {}
func (*TerraformDocs) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *TerraformDocs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerraformDocs.Unmarshal(m, b)
}
func (m *TerraformDocs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerraformDocs.Marshal(b, m, deterministic)
}
func (m *TerraformDocs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerraformDocs.Merge(m, src)
}
func (m *TerraformDocs) XXX_Size() int {
	return xxx_messageInfo_TerraformDocs.Size(m)
}
func (m *TerraformDocs) XXX_DiscardUnknown() {
	xxx_messageInfo_TerraformDocs.DiscardUnknown(m)
}

var xxx_messageInfo_TerraformDocs proto.InternalMessageInfo

func (m *TerraformDocs) GetResourceExamplePath() string {
	if m != nil {
		return m.ResourceExamplePath
	}
	return ""
}

func (m *TerraformDocs) GetDataSourceExamplePath() string {
	if m != nil {
		return m.DataSourceExamplePath
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
	return fileDescriptor_110d40819f1994f9, []int{3}
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
	return fileDescriptor_110d40819f1994f9, []int{4}
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
	proto.RegisterType((*TerraformDocs)(nil), "v1.TerraformDocs")
	proto.RegisterType((*OneofOptions)(nil), "v1.OneofOptions")
	proto.RegisterType((*ServiceOptions)(nil), "v1.ServiceOptions")
	proto.RegisterExtension(E_FieldOptions)
	proto.RegisterExtension(E_MessageOptions)
	proto.RegisterExtension(E_OneofOptions)
	proto.RegisterExtension(E_ServiceOptions)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x10, 0x2d, 0xc9, 0xd4, 0x09, 0x74, 0x4b, 0x91, 0x41, 0xa5, 0xad, 0x82, 0x90, 0x38,
	0x39, 0x14, 0xb8, 0x90, 0x1b, 0x08, 0xb8, 0x91, 0x56, 0x6e, 0xc5, 0x91, 0xd5, 0xc6, 0x9e, 0xa4,
	0x2b, 0xd9, 0x1e, 0x67, 0xd7, 0x8e, 0xda, 0x1b, 0x07, 0xb8, 0xc2, 0x15, 0xfe, 0x07, 0xe5, 0xe3,
	0x42, 0xff, 0x00, 0x27, 0x7e, 0x11, 0x5a, 0x7b, 0x9d, 0xda, 0x6d, 0xd5, 0xe3, 0x7b, 0x6f, 0x66,
	0xfd, 0xde, 0xf3, 0x40, 0x97, 0xd2, 0x4c, 0x52, 0xa2, 0xbd, 0x54, 0x51, 0x46, 0xec, 0xda, 0x7c,
	0xe7, 0xde, 0xf6, 0x94, 0x68, 0x1a, 0xe1, 0xa0, 0x60, 0xc6, 0xf9, 0x64, 0x10, 0xa2, 0x0e, 0x94,
	0x4c, 0x33, 0x52, 0xe5, 0x54, 0xff, 0x5f, 0x0b, 0x9c, 0x37, 0x12, 0xa3, 0x70, 0xb7, 0x5c, 0x66,
	0xb7, 0xe1, 0x7a, 0x22, 0x62, 0x74, 0xbf, 0xff, 0x99, 0x6f, 0xb7, 0x1e, 0x75, 0xfc, 0x02, 0xb1,
	0x07, 0xe0, 0xe8, 0x59, 0xc4, 0x93, 0x3c, 0x8a, 0xc4, 0x38, 0x42, 0xf7, 0xa4, 0x50, 0xdb, 0xfe,
	0x8a, 0x9e, 0x45, 0x23, 0x4b, 0xb2, 0xc7, 0xb0, 0x86, 0x47, 0x29, 0x69, 0xe4, 0x42, 0xf3, 0x94,
	0x54, 0x80, 0x91, 0x90, 0x89, 0xfb, 0xc3, 0xce, 0xae, 0x96, 0xe2, 0x0b, 0xbd, 0x57, 0x49, 0x6c,
	0x03, 0xda, 0x32, 0x43, 0x55, 0x3c, 0xf9, 0xd3, 0x8e, 0x2d, 0x18, 0xa3, 0x2a, 0x9c, 0xe5, 0x52,
	0x61, 0xe8, 0xfe, 0xaa, 0xd4, 0x8a, 0x61, 0x77, 0xe1, 0x86, 0x0c, 0x79, 0x76, 0x9c, 0xa2, 0xfb,
	0xdb, 0x7a, 0x5d, 0x96, 0xe1, 0xc1, 0x71, 0x8a, 0xfd, 0xbf, 0x2d, 0xe8, 0xbd, 0x45, 0xad, 0xc5,
	0x14, 0xab, 0x58, 0x5b, 0x00, 0x31, 0x85, 0x18, 0xf1, 0x46, 0xb8, 0x4e, 0xc1, 0x8d, 0x4c, 0xc2,
	0x4d, 0xe8, 0x9c, 0x59, 0xae, 0xe2, 0x9d, 0x51, 0xec, 0x0e, 0x2c, 0xa1, 0x52, 0xa4, 0x6c, 0x9c,
	0x25, 0xbf, 0x84, 0xec, 0xe1, 0xa2, 0x77, 0x3e, 0x31, 0x3d, 0xda, 0x1c, 0x1d, 0xdf, 0xb1, 0x74,
	0xd1, 0x2e, 0x1b, 0x42, 0x2f, 0x43, 0xa5, 0xc4, 0x84, 0x54, 0xcc, 0x43, 0x0a, 0xb4, 0x4d, 0xb4,
	0xf2, 0x64, 0xd5, 0x9b, 0xef, 0x78, 0x07, 0x95, 0xf6, 0x8a, 0x02, 0xed, 0x77, 0xb3, 0x3a, 0xec,
	0x7f, 0x68, 0x41, 0xb7, 0x31, 0xc0, 0x9e, 0xc1, 0xba, 0x42, 0x4d, 0xb9, 0x0a, 0x90, 0xe3, 0x91,
	0x88, 0xd3, 0x08, 0x79, 0x2a, 0xb2, 0xc3, 0x45, 0xb0, 0xb5, 0x4a, 0x7e, 0x5d, 0xaa, 0x7b, 0x22,
	0x3b, 0x64, 0xcf, 0xc1, 0x0d, 0x45, 0x26, 0xf8, 0x65, 0x8b, 0x27, 0x76, 0x71, 0xdd, 0x4c, 0xec,
	0x9f, 0x5f, 0xed, 0x0f, 0xc0, 0xd9, 0x4d, 0x90, 0x26, 0x97, 0xd7, 0xf9, 0xf1, 0xf4, 0x7c, 0x9d,
	0xfd, 0x01, 0xf4, 0xf6, 0x51, 0xcd, 0x65, 0xb0, 0xf8, 0x03, 0xf7, 0xa1, 0x13, 0x0b, 0x99, 0xf0,
	0x84, 0xf2, 0xc4, 0xfd, 0x6a, 0x37, 0xda, 0x86, 0x1a, 0x51, 0x9e, 0x0c, 0xdf, 0x41, 0xb7, 0xe8,
	0x8f, 0x53, 0x35, 0xef, 0x95, 0xc7, 0xeb, 0x55, 0xc7, 0xeb, 0xd5, 0xef, 0xd4, 0xfd, 0x7c, 0x5a,
	0xd6, 0x77, 0xcb, 0xd4, 0x57, 0x57, 0x7c, 0x67, 0x52, 0x43, 0xc3, 0xf7, 0x70, 0x33, 0x2e, 0x4f,
	0x61, 0xf1, 0xf2, 0xd6, 0x85, 0x97, 0x9b, 0xc7, 0xe2, 0x7e, 0xb1, 0x6f, 0x33, 0xf3, 0x76, 0x53,
	0xf3, 0x7b, 0x71, 0x03, 0x1b, 0xdf, 0x64, 0x9a, 0xb9, 0xc2, 0x77, 0xbd, 0x39, 0xf7, 0x53, 0xdd,
	0x77, 0x5d, 0xf1, 0x1d, 0xaa, 0x21, 0xe3, 0x5b, 0x97, 0x05, 0x5e, 0xe1, 0xbb, 0x59, 0xb1, 0xfb,
	0xad, 0xee, 0xbb, 0xa9, 0xf9, 0x3d, 0xdd, 0xc0, 0x2f, 0x37, 0x61, 0x23, 0xa0, 0xd8, 0xd3, 0x99,
	0xa2, 0x64, 0x1a, 0xc6, 0x9e, 0x48, 0xa5, 0xd9, 0x4a, 0xa3, 0x3c, 0x1e, 0xcb, 0x64, 0x3a, 0x5e,
	0x2e, 0x3e, 0xf2, 0xf4, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x96, 0x5c, 0xdb, 0x56, 0x04,
	0x00, 0x00,
}

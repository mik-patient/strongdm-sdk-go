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
	Iterable bool   `protobuf:"varint,1941303,opt,name=iterable,proto3" json:"iterable,omitempty"`
	Required bool   `protobuf:"varint,1941304,opt,name=required,proto3" json:"required,omitempty"`
	IdType   string `protobuf:"bytes,1941305,opt,name=id_type,json=idType,proto3" json:"id_type,omitempty"`
	// sdk_only fields are not mapped directly to fields in our internal models.
	SdkOnly bool `protobuf:"varint,1941306,opt,name=sdk_only,json=sdkOnly,proto3" json:"sdk_only,omitempty"`
	// computed exposes fields to Terraform even if they are read_only.
	Computed bool `protobuf:"varint,1941307,opt,name=computed,proto3" json:"computed,omitempty"`
	// force_new forces Terraform to delete and recreate the object if the field changes.
	ForceNew bool `protobuf:"varint,1941308,opt,name=force_new,json=forceNew,proto3" json:"force_new,omitempty"`
	// custom_converter allows you to write custom "to porcelain" and "to plumbing"
	// conversion functions for this field in the SDK templates.
	CustomConverter string `protobuf:"bytes,1941309,opt,name=custom_converter,json=customConverter,proto3" json:"custom_converter,omitempty"`
	// custom_go_porcelain_type allows you to customize the porcelain field type in go
	CustomGoPorcelainType string `protobuf:"bytes,1941310,opt,name=custom_go_porcelain_type,json=customGoPorcelainType,proto3" json:"custom_go_porcelain_type,omitempty"`
	// custom_java_porcelain_type allows you to customize the porcelain field type in java
	CustomJavaPorcelainType string   `protobuf:"bytes,1941311,opt,name=custom_java_porcelain_type,json=customJavaPorcelainType,proto3" json:"custom_java_porcelain_type,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
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

func (m *FieldOptions) GetSdkOnly() bool {
	if m != nil {
		return m.SdkOnly
	}
	return false
}

func (m *FieldOptions) GetComputed() bool {
	if m != nil {
		return m.Computed
	}
	return false
}

func (m *FieldOptions) GetForceNew() bool {
	if m != nil {
		return m.ForceNew
	}
	return false
}

func (m *FieldOptions) GetCustomConverter() string {
	if m != nil {
		return m.CustomConverter
	}
	return ""
}

func (m *FieldOptions) GetCustomGoPorcelainType() string {
	if m != nil {
		return m.CustomGoPorcelainType
	}
	return ""
}

func (m *FieldOptions) GetCustomJavaPorcelainType() string {
	if m != nil {
		return m.CustomJavaPorcelainType
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
	CommonFields         []string `protobuf:"bytes,1941381,rep,name=common_fields,json=commonFields,proto3" json:"common_fields,omitempty"`
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

func (m *OneofOptions) GetCommonFields() []string {
	if m != nil {
		return m.CommonFields
	}
	return nil
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
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x4f, 0xd4, 0x4c,
	0x14, 0x0e, 0x2f, 0x2f, 0xb0, 0x3b, 0xec, 0x2e, 0x30, 0xbc, 0xbc, 0x8e, 0xc8, 0x57, 0xd6, 0x90,
	0x98, 0x98, 0x14, 0x51, 0x6f, 0x24, 0xf1, 0xc2, 0xef, 0xc4, 0xc4, 0x85, 0x14, 0xc2, 0xa5, 0xcd,
	0x6c, 0x3b, 0xbb, 0x54, 0xda, 0x39, 0x65, 0xa6, 0x2d, 0xec, 0x9d, 0x17, 0x7a, 0xab, 0xb7, 0xfa,
	0x3f, 0xc4, 0x6f, 0xd9, 0x3f, 0xe0, 0x9d, 0x7f, 0xc8, 0x4c, 0x67, 0x5a, 0x5a, 0x20, 0x5c, 0x9e,
	0xe7, 0x39, 0xcf, 0xe9, 0xe9, 0xf3, 0x9c, 0x41, 0x4d, 0x88, 0x62, 0x1f, 0xb8, 0xb4, 0x22, 0x01,
	0x31, 0xe0, 0x7f, 0xd2, 0xf5, 0xf9, 0x95, 0x3e, 0x40, 0x3f, 0x60, 0x6b, 0x19, 0xd2, 0x4d, 0x7a,
	0x6b, 0x1e, 0x93, 0xae, 0xf0, 0xa3, 0x18, 0x84, 0xee, 0x6a, 0xff, 0x19, 0x45, 0x8d, 0xa7, 0x3e,
	0x0b, 0xbc, 0x4d, 0x2d, 0xc6, 0xff, 0xa1, 0x7f, 0x39, 0x0d, 0x19, 0xf9, 0x74, 0x92, 0xae, 0x8c,
	0xdc, 0xa8, 0xdb, 0x59, 0x85, 0xaf, 0xa3, 0x86, 0x3c, 0x08, 0x1c, 0x9e, 0x04, 0x01, 0xed, 0x06,
	0x8c, 0x1c, 0x67, 0x6c, 0xcd, 0x9e, 0x94, 0x07, 0x41, 0xc7, 0x80, 0xf8, 0x16, 0x9a, 0x65, 0x47,
	0x11, 0x48, 0xe6, 0x50, 0xe9, 0x44, 0x20, 0x5c, 0x16, 0x50, 0x9f, 0x93, 0xcf, 0xa6, 0x77, 0x46,
	0x93, 0x0f, 0xe4, 0x56, 0x4e, 0xe1, 0x05, 0x54, 0xf3, 0x63, 0x26, 0xb2, 0x91, 0x5f, 0x4c, 0x5b,
	0x81, 0x28, 0x56, 0xb0, 0x83, 0xc4, 0x17, 0xcc, 0x23, 0x5f, 0x73, 0x36, 0x47, 0xf0, 0x55, 0x34,
	0xe1, 0x7b, 0x4e, 0x3c, 0x88, 0x18, 0xf9, 0x66, 0x76, 0x1d, 0xf7, 0xbd, 0x9d, 0x41, 0xc4, 0xf0,
	0x35, 0x54, 0x93, 0xde, 0xbe, 0x03, 0x3c, 0x18, 0x90, 0xef, 0x46, 0x38, 0x21, 0xbd, 0xfd, 0x4d,
	0x1e, 0x0c, 0xd4, 0x54, 0x17, 0xc2, 0x28, 0x89, 0x99, 0x47, 0x7e, 0xe4, 0x53, 0x73, 0x04, 0x2f,
	0xa2, 0x7a, 0x4f, 0xad, 0xe7, 0x70, 0x76, 0x48, 0x7e, 0xe6, 0x74, 0x06, 0x75, 0xd8, 0x21, 0xbe,
	0x89, 0xa6, 0xdd, 0x44, 0xc6, 0x10, 0x3a, 0x2e, 0xf0, 0x94, 0x89, 0x98, 0x09, 0xf2, 0xcb, 0x7c,
	0x7d, 0x4a, 0x33, 0x8f, 0x72, 0x02, 0xdf, 0x43, 0xc4, 0x34, 0xf7, 0xe1, 0xd4, 0x0f, 0xbd, 0xf2,
	0x89, 0x11, 0xcd, 0xe9, 0x8e, 0x67, 0x50, 0x98, 0x92, 0xfd, 0xc1, 0x7d, 0x34, 0x6f, 0xa4, 0xaf,
	0x68, 0x4a, 0xcf, 0x8a, 0x87, 0x46, 0x7c, 0x45, 0xf7, 0x3c, 0xa7, 0x29, 0xad, 0xc8, 0xdb, 0xbf,
	0x47, 0x50, 0xeb, 0x05, 0x93, 0x92, 0xf6, 0x59, 0x9e, 0xeb, 0x32, 0x42, 0x21, 0x78, 0x2c, 0x70,
	0x2a, 0xe9, 0xd6, 0x33, 0xac, 0xa3, 0x22, 0x5e, 0x42, 0xf5, 0xd3, 0xcc, 0xf2, 0x7c, 0x4f, 0x21,
	0xfc, 0x3f, 0x1a, 0x63, 0x42, 0x80, 0x30, 0x79, 0x8e, 0xd9, 0xba, 0xc4, 0xab, 0xc5, 0xe1, 0x39,
	0x3d, 0x75, 0x48, 0x26, 0xc8, 0xba, 0xdd, 0x30, 0x70, 0x76, 0x5e, 0x78, 0x03, 0xb5, 0x62, 0x26,
	0x04, 0xed, 0x81, 0x08, 0x1d, 0x0f, 0x5c, 0x69, 0x22, 0x9d, 0xbc, 0x3d, 0x63, 0xa5, 0xeb, 0xd6,
	0x4e, 0xce, 0x3d, 0x06, 0x57, 0xda, 0xcd, 0xb8, 0x5c, 0xb6, 0x5f, 0x8f, 0xa0, 0x66, 0xa5, 0x01,
	0xdf, 0x45, 0x73, 0x82, 0x49, 0x48, 0x54, 0x52, 0xec, 0x88, 0x86, 0x51, 0xc0, 0x9c, 0x88, 0xc6,
	0x7b, 0xc5, 0x8f, 0xcd, 0xe6, 0xf4, 0x13, 0xcd, 0x6e, 0xd1, 0x78, 0x4f, 0x05, 0xe2, 0xd1, 0x98,
	0x3a, 0x17, 0x09, 0x8f, 0xf3, 0x40, 0x54, 0xc7, 0xf6, 0x59, 0x69, 0x7b, 0x17, 0x35, 0x36, 0x39,
	0x83, 0xde, 0xc5, 0x76, 0xbe, 0x19, 0x9e, 0xb3, 0x73, 0x15, 0x35, 0x5d, 0x08, 0x43, 0xe0, 0xda,
	0x15, 0x49, 0xde, 0x0e, 0xd3, 0x95, 0x51, 0x65, 0x8b, 0x86, 0x33, 0x57, 0x64, 0x7b, 0x0d, 0xb5,
	0xb6, 0x99, 0x48, 0x7d, 0xb7, 0x08, 0x6a, 0x11, 0xd5, 0x43, 0x95, 0x34, 0x87, 0x84, 0x93, 0x0f,
	0x66, 0x70, 0x4d, 0x41, 0x1d, 0x48, 0xf8, 0xc6, 0x2e, 0x6a, 0x66, 0x03, 0x1d, 0xc8, 0xfb, 0x2d,
	0xfd, 0xc8, 0xad, 0xfc, 0x91, 0x5b, 0xe5, 0xf7, 0x4c, 0xde, 0x0d, 0xb5, 0xcb, 0xd3, 0xca, 0xe5,
	0x32, 0x63, 0x37, 0x7a, 0xa5, 0x6a, 0xe3, 0x25, 0x9a, 0x0a, 0xf5, 0xc5, 0x14, 0x93, 0x97, 0xcf,
	0x4d, 0xae, 0xde, 0x14, 0x79, 0x6f, 0x66, 0x63, 0x35, 0xbb, 0xca, 0xd9, 0xad, 0xb0, 0x52, 0xab,
	0xbd, 0x41, 0x19, 0x78, 0xc9, 0xde, 0x65, 0x83, 0x33, 0xbb, 0x8a, 0xbd, 0xcb, 0x8c, 0xdd, 0x80,
	0x52, 0xa5, 0xf6, 0x96, 0xda, 0xc0, 0x4b, 0xf6, 0xae, 0x5a, 0x4c, 0x3e, 0x96, 0xf7, 0xae, 0x72,
	0x76, 0x4b, 0x56, 0xea, 0x87, 0x4b, 0x68, 0xc1, 0x85, 0xd0, 0x92, 0xb1, 0x00, 0xde, 0xf7, 0x42,
	0x8b, 0x46, 0xbe, 0x52, 0x45, 0x41, 0x12, 0x76, 0x7d, 0xde, 0xef, 0x8e, 0x67, 0x1f, 0xb9, 0xf3,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xe0, 0x7f, 0xcc, 0x7e, 0x05, 0x00, 0x00,
}

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
// source: options.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	ForceNew  bool `protobuf:"varint,1941308,opt,name=force_new,json=forceNew,proto3" json:"force_new,omitempty"`
	WriteOnly bool `protobuf:"varint,1941309,opt,name=write_only,json=writeOnly,proto3" json:"write_only,omitempty"`
	// sensitive determines whether the field should be marked as sensitive in Terraform.
	Sensitive           bool   `protobuf:"varint,1941310,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	CliName             string `protobuf:"bytes,1941311,opt,name=cli_name,json=cliName,proto3" json:"cli_name,omitempty"`
	JsonName            string `protobuf:"bytes,1941312,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	JsonGatewayName     string `protobuf:"bytes,1941313,opt,name=json_gateway_name,json=jsonGatewayName,proto3" json:"json_gateway_name,omitempty"`
	HideFromJsonGateway bool   `protobuf:"varint,1941314,opt,name=hide_from_json_gateway,json=hideFromJsonGateway,proto3" json:"hide_from_json_gateway,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{0}
}

func (x *FieldOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FieldOptions) GetSqlNullable() bool {
	if x != nil {
		return x.SqlNullable
	}
	return false
}

func (x *FieldOptions) GetExposeAsPorcelain() bool {
	if x != nil {
		return x.ExposeAsPorcelain
	}
	return false
}

func (x *FieldOptions) GetIterable() bool {
	if x != nil {
		return x.Iterable
	}
	return false
}

func (x *FieldOptions) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FieldOptions) GetIdType() string {
	if x != nil {
		return x.IdType
	}
	return ""
}

func (x *FieldOptions) GetSdkOnly() bool {
	if x != nil {
		return x.SdkOnly
	}
	return false
}

func (x *FieldOptions) GetComputed() bool {
	if x != nil {
		return x.Computed
	}
	return false
}

func (x *FieldOptions) GetForceNew() bool {
	if x != nil {
		return x.ForceNew
	}
	return false
}

func (x *FieldOptions) GetWriteOnly() bool {
	if x != nil {
		return x.WriteOnly
	}
	return false
}

func (x *FieldOptions) GetSensitive() bool {
	if x != nil {
		return x.Sensitive
	}
	return false
}

func (x *FieldOptions) GetCliName() string {
	if x != nil {
		return x.CliName
	}
	return ""
}

func (x *FieldOptions) GetJsonName() string {
	if x != nil {
		return x.JsonName
	}
	return ""
}

func (x *FieldOptions) GetJsonGatewayName() string {
	if x != nil {
		return x.JsonGatewayName
	}
	return ""
}

func (x *FieldOptions) GetHideFromJsonGateway() bool {
	if x != nil {
		return x.HideFromJsonGateway
	}
	return false
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName string `protobuf:"bytes,1941300,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	Porcelain bool   `protobuf:"varint,1941301,opt,name=porcelain,proto3" json:"porcelain,omitempty"`
	// Non-zero if the message is an error type. This corresponds to the gRPC status code.
	Error int32 `protobuf:"varint,1941302,opt,name=error,proto3" json:"error,omitempty"`
	// Set this option on an RPC request message to specify which field holds
	// the "options" for that RPC method.
	OptionsField        string                         `protobuf:"bytes,1941303,opt,name=options_field,json=optionsField,proto3" json:"options_field,omitempty"`
	TerraformDocs       *TerraformDocs                 `protobuf:"bytes,1941304,opt,name=terraform_docs,json=terraformDocs,proto3" json:"terraform_docs,omitempty"`
	Custom              *CustomPorcelainMessageOptions `protobuf:"bytes,1941305,opt,name=custom,proto3" json:"custom,omitempty"`
	PrivateSdk          bool                           `protobuf:"varint,1941306,opt,name=private_sdk,json=privateSdk,proto3" json:"private_sdk,omitempty"`
	CliName             string                         `protobuf:"bytes,1941307,opt,name=cli_name,json=cliName,proto3" json:"cli_name,omitempty"`
	JsonName            string                         `protobuf:"bytes,1941308,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	JsonGatewayName     string                         `protobuf:"bytes,1941309,opt,name=json_gateway_name,json=jsonGatewayName,proto3" json:"json_gateway_name,omitempty"`
	HideFromJsonGateway bool                           `protobuf:"varint,1941310,opt,name=hide_from_json_gateway,json=hideFromJsonGateway,proto3" json:"hide_from_json_gateway,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{1}
}

func (x *MessageOptions) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *MessageOptions) GetPorcelain() bool {
	if x != nil {
		return x.Porcelain
	}
	return false
}

func (x *MessageOptions) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *MessageOptions) GetOptionsField() string {
	if x != nil {
		return x.OptionsField
	}
	return ""
}

func (x *MessageOptions) GetTerraformDocs() *TerraformDocs {
	if x != nil {
		return x.TerraformDocs
	}
	return nil
}

func (x *MessageOptions) GetCustom() *CustomPorcelainMessageOptions {
	if x != nil {
		return x.Custom
	}
	return nil
}

func (x *MessageOptions) GetPrivateSdk() bool {
	if x != nil {
		return x.PrivateSdk
	}
	return false
}

func (x *MessageOptions) GetCliName() string {
	if x != nil {
		return x.CliName
	}
	return ""
}

func (x *MessageOptions) GetJsonName() string {
	if x != nil {
		return x.JsonName
	}
	return ""
}

func (x *MessageOptions) GetJsonGatewayName() string {
	if x != nil {
		return x.JsonGatewayName
	}
	return ""
}

func (x *MessageOptions) GetHideFromJsonGateway() bool {
	if x != nil {
		return x.HideFromJsonGateway
	}
	return false
}

// CustomPorcelainMessageOptions allows you to create a message type that
// is converted from the underlying proto message into whatever representation
// is most appropriate in the target languages.
type CustomPorcelainMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// converter allows you to write custom "to porcelain" and "to plumbing"
	// conversion functions for this message in the SDK templates.
	Converter string `protobuf:"bytes,1941309,opt,name=converter,proto3" json:"converter,omitempty"`
	// go_porcelain_type allows you to customize the porcelain message type in go
	GoPorcelainType string `protobuf:"bytes,1941310,opt,name=go_porcelain_type,json=goPorcelainType,proto3" json:"go_porcelain_type,omitempty"`
	// java_porcelain_type allows you to customize the porcelain message type in java
	JavaPorcelainType string `protobuf:"bytes,1941311,opt,name=java_porcelain_type,json=javaPorcelainType,proto3" json:"java_porcelain_type,omitempty"`
	// terraform_porcelain_type allows you to customize the porcelain message type in terraform
	TerraformPorcelainType string `protobuf:"bytes,1941312,opt,name=terraform_porcelain_type,json=terraformPorcelainType,proto3" json:"terraform_porcelain_type,omitempty"`
	// openapi_porcelain_type allows you to customize the porcelain message type in openapi spec
	OpenapiPorcelainType string `protobuf:"bytes,1941313,opt,name=openapi_porcelain_type,json=openapiPorcelainType,proto3" json:"openapi_porcelain_type,omitempty"`
}

func (x *CustomPorcelainMessageOptions) Reset() {
	*x = CustomPorcelainMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPorcelainMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPorcelainMessageOptions) ProtoMessage() {}

func (x *CustomPorcelainMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPorcelainMessageOptions.ProtoReflect.Descriptor instead.
func (*CustomPorcelainMessageOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{2}
}

func (x *CustomPorcelainMessageOptions) GetConverter() string {
	if x != nil {
		return x.Converter
	}
	return ""
}

func (x *CustomPorcelainMessageOptions) GetGoPorcelainType() string {
	if x != nil {
		return x.GoPorcelainType
	}
	return ""
}

func (x *CustomPorcelainMessageOptions) GetJavaPorcelainType() string {
	if x != nil {
		return x.JavaPorcelainType
	}
	return ""
}

func (x *CustomPorcelainMessageOptions) GetTerraformPorcelainType() string {
	if x != nil {
		return x.TerraformPorcelainType
	}
	return ""
}

func (x *CustomPorcelainMessageOptions) GetOpenapiPorcelainType() string {
	if x != nil {
		return x.OpenapiPorcelainType
	}
	return ""
}

type TerraformDocs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceExamplePath   string `protobuf:"bytes,1941300,opt,name=resource_example_path,json=resourceExamplePath,proto3" json:"resource_example_path,omitempty"`
	DataSourceExamplePath string `protobuf:"bytes,1941301,opt,name=data_source_example_path,json=dataSourceExamplePath,proto3" json:"data_source_example_path,omitempty"`
}

func (x *TerraformDocs) Reset() {
	*x = TerraformDocs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerraformDocs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerraformDocs) ProtoMessage() {}

func (x *TerraformDocs) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerraformDocs.ProtoReflect.Descriptor instead.
func (*TerraformDocs) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{3}
}

func (x *TerraformDocs) GetResourceExamplePath() string {
	if x != nil {
		return x.ResourceExamplePath
	}
	return ""
}

func (x *TerraformDocs) GetDataSourceExamplePath() string {
	if x != nil {
		return x.DataSourceExamplePath
	}
	return ""
}

type OneofOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName    string   `protobuf:"bytes,1941380,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	CommonFields []string `protobuf:"bytes,1941381,rep,name=common_fields,json=commonFields,proto3" json:"common_fields,omitempty"`
}

func (x *OneofOptions) Reset() {
	*x = OneofOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofOptions) ProtoMessage() {}

func (x *OneofOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofOptions.ProtoReflect.Descriptor instead.
func (*OneofOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{4}
}

func (x *OneofOptions) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *OneofOptions) GetCommonFields() []string {
	if x != nil {
		return x.CommonFields
	}
	return nil
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainNoun   string `protobuf:"bytes,1941400,opt,name=main_noun,json=mainNoun,proto3" json:"main_noun,omitempty"`
	PrivateSdk bool   `protobuf:"varint,1941401,opt,name=private_sdk,json=privateSdk,proto3" json:"private_sdk,omitempty"`
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_options_proto_rawDescGZIP(), []int{5}
}

func (x *ServiceOptions) GetMainNoun() string {
	if x != nil {
		return x.MainNoun
	}
	return ""
}

func (x *ServiceOptions) GetPrivateSdk() bool {
	if x != nil {
		return x.PrivateSdk
	}
	return false
}

var file_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1941390,
		Name:          "v1.field_options",
		Tag:           "bytes,1941390,opt,name=field_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1941391,
		Name:          "v1.message_options",
		Tag:           "bytes,1941391,opt,name=message_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.OneofOptions)(nil),
		ExtensionType: (*OneofOptions)(nil),
		Field:         1941381,
		Name:          "v1.oneof_options",
		Tag:           "bytes,1941381,opt,name=oneof_options",
		Filename:      "options.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*ServiceOptions)(nil),
		Field:         1941401,
		Name:          "v1.service_options",
		Tag:           "bytes,1941401,opt,name=service_options",
		Filename:      "options.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional v1.FieldOptions field_options = 1941390;
	E_FieldOptions = &file_options_proto_extTypes[0]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional v1.MessageOptions message_options = 1941391;
	E_MessageOptions = &file_options_proto_extTypes[1]
)

// Extension fields to descriptor.OneofOptions.
var (
	// optional v1.OneofOptions oneof_options = 1941381;
	E_OneofOptions = &file_options_proto_extTypes[2]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional v1.ServiceOptions service_options = 1941401;
	E_ServiceOptions = &file_options_proto_extTypes[3]
)

var File_options_proto protoreflect.FileDescriptor

var file_options_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x04, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xb4,
	0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0xb5, 0xbe, 0x76,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x71, 0x6c, 0x4e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x5f, 0x70,
	0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0xb6, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x41, 0x73, 0x50, 0x6f, 0x72, 0x63, 0x65, 0x6c,
	0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0xb7, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0xb8, 0xbe,
	0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x19, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xb9, 0xbe, 0x76, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x73, 0x64,
	0x6b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0xba, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x64, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x64, 0x18, 0xbb, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x65, 0x77, 0x18, 0xbc, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x4e, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0xbd, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x18, 0xbe, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0xbf, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0xc0, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc1, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6a, 0x73, 0x6f, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x16, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0xc2, 0xbe, 0x76, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x68, 0x69, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xcd, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xb4, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x09, 0x70, 0x6f,
	0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0xb5, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x70, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0xb6, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0xb7, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x64, 0x6f, 0x63, 0x73, 0x18, 0xb8, 0xbe, 0x76, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x44, 0x6f, 0x63, 0x73, 0x52, 0x0d, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x44, 0x6f, 0x63, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18,
	0xb9, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x64,
	0x6b, 0x18, 0xba, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x53, 0x64, 0x6b, 0x12, 0x1b, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0xbb, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0xbc, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xbd, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6a, 0x73, 0x6f, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x16, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0xbe, 0xbe, 0x76, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x68, 0x69, 0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x93, 0x02, 0x0a, 0x1d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x50, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x18, 0xbd, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x67, 0x6f, 0x5f, 0x70,
	0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xbe, 0xbe,
	0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x6f, 0x50, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x70,
	0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xbf, 0xbe,
	0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x6f, 0x72, 0x63, 0x65,
	0x6c, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x70, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0xc0, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xc1,
	0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x50,
	0x6f, 0x72, 0x63, 0x65, 0x6c, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x0d, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x44, 0x6f, 0x63, 0x73, 0x12, 0x34,
	0x0a, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0xb4, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x18, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0xb5, 0xbe, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x56, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x84, 0xbf,
	0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x85, 0xbf, 0x76, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x6e, 0x6f, 0x75, 0x6e, 0x18, 0x98, 0xbf, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x6f, 0x75, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x5f, 0x73, 0x64, 0x6b, 0x18, 0x99, 0xbf, 0x76, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x64, 0x6b, 0x3a, 0x56, 0x0a, 0x0d, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8e, 0xbf, 0x76, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x5e, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8f, 0xbf, 0x76, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x56, 0x0a, 0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x85, 0xbf, 0x76, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x5e, 0x0a, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x99, 0xbf, 0x76, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x52, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6e, 0x67, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x64, 0x6d, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_proto_rawDescOnce sync.Once
	file_options_proto_rawDescData = file_options_proto_rawDesc
)

func file_options_proto_rawDescGZIP() []byte {
	file_options_proto_rawDescOnce.Do(func() {
		file_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_proto_rawDescData)
	})
	return file_options_proto_rawDescData
}

var file_options_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_options_proto_goTypes = []interface{}{
	(*FieldOptions)(nil),                  // 0: v1.FieldOptions
	(*MessageOptions)(nil),                // 1: v1.MessageOptions
	(*CustomPorcelainMessageOptions)(nil), // 2: v1.CustomPorcelainMessageOptions
	(*TerraformDocs)(nil),                 // 3: v1.TerraformDocs
	(*OneofOptions)(nil),                  // 4: v1.OneofOptions
	(*ServiceOptions)(nil),                // 5: v1.ServiceOptions
	(*descriptor.FieldOptions)(nil),       // 6: google.protobuf.FieldOptions
	(*descriptor.MessageOptions)(nil),     // 7: google.protobuf.MessageOptions
	(*descriptor.OneofOptions)(nil),       // 8: google.protobuf.OneofOptions
	(*descriptor.ServiceOptions)(nil),     // 9: google.protobuf.ServiceOptions
}
var file_options_proto_depIdxs = []int32{
	3,  // 0: v1.MessageOptions.terraform_docs:type_name -> v1.TerraformDocs
	2,  // 1: v1.MessageOptions.custom:type_name -> v1.CustomPorcelainMessageOptions
	6,  // 2: v1.field_options:extendee -> google.protobuf.FieldOptions
	7,  // 3: v1.message_options:extendee -> google.protobuf.MessageOptions
	8,  // 4: v1.oneof_options:extendee -> google.protobuf.OneofOptions
	9,  // 5: v1.service_options:extendee -> google.protobuf.ServiceOptions
	0,  // 6: v1.field_options:type_name -> v1.FieldOptions
	1,  // 7: v1.message_options:type_name -> v1.MessageOptions
	4,  // 8: v1.oneof_options:type_name -> v1.OneofOptions
	5,  // 9: v1.service_options:type_name -> v1.ServiceOptions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	6,  // [6:10] is the sub-list for extension type_name
	2,  // [2:6] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_options_proto_init() }
func file_options_proto_init() {
	if File_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
		file_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPorcelainMessageOptions); i {
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
		file_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerraformDocs); i {
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
		file_options_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofOptions); i {
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
		file_options_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
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
			RawDescriptor: file_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_options_proto_goTypes,
		DependencyIndexes: file_options_proto_depIdxs,
		MessageInfos:      file_options_proto_msgTypes,
		ExtensionInfos:    file_options_proto_extTypes,
	}.Build()
	File_options_proto = out.File
	file_options_proto_rawDesc = nil
	file_options_proto_goTypes = nil
	file_options_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drivers.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Driver struct {
	// Types that are valid to be assigned to Driver:
	//	*Driver_Mysql
	//	*Driver_Athena
	Driver               isDriver_Driver `protobuf_oneof:"driver"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Driver) Reset()         { *m = Driver{} }
func (m *Driver) String() string { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()    {}
func (*Driver) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{0}
}

func (m *Driver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Driver.Unmarshal(m, b)
}
func (m *Driver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Driver.Marshal(b, m, deterministic)
}
func (m *Driver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Driver.Merge(m, src)
}
func (m *Driver) XXX_Size() int {
	return xxx_messageInfo_Driver.Size(m)
}
func (m *Driver) XXX_DiscardUnknown() {
	xxx_messageInfo_Driver.DiscardUnknown(m)
}

var xxx_messageInfo_Driver proto.InternalMessageInfo

type isDriver_Driver interface {
	isDriver_Driver()
}

type Driver_Mysql struct {
	Mysql *Mysql `protobuf:"bytes,1,opt,name=mysql,proto3,oneof"`
}

type Driver_Athena struct {
	Athena *Athena `protobuf:"bytes,2,opt,name=athena,proto3,oneof"`
}

func (*Driver_Mysql) isDriver_Driver() {}

func (*Driver_Athena) isDriver_Driver() {}

func (m *Driver) GetDriver() isDriver_Driver {
	if m != nil {
		return m.Driver
	}
	return nil
}

func (m *Driver) GetMysql() *Mysql {
	if x, ok := m.GetDriver().(*Driver_Mysql); ok {
		return x.Mysql
	}
	return nil
}

func (m *Driver) GetAthena() *Athena {
	if x, ok := m.GetDriver().(*Driver_Athena); ok {
		return x.Athena
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Driver) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Driver_Mysql)(nil),
		(*Driver_Athena)(nil),
	}
}

type Mysql struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Port                 int32    `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mysql) Reset()         { *m = Mysql{} }
func (m *Mysql) String() string { return proto.CompactTextString(m) }
func (*Mysql) ProtoMessage()    {}
func (*Mysql) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{1}
}

func (m *Mysql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mysql.Unmarshal(m, b)
}
func (m *Mysql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mysql.Marshal(b, m, deterministic)
}
func (m *Mysql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mysql.Merge(m, src)
}
func (m *Mysql) XXX_Size() int {
	return xxx_messageInfo_Mysql.Size(m)
}
func (m *Mysql) XXX_DiscardUnknown() {
	xxx_messageInfo_Mysql.DiscardUnknown(m)
}

var xxx_messageInfo_Mysql proto.InternalMessageInfo

func (m *Mysql) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Mysql) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Mysql) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Mysql) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Mysql) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Athena struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	AccessKey            string   `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretAccessKey      string   `protobuf:"bytes,3,opt,name=secretAccessKey,proto3" json:"secretAccessKey,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Output               string   `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Athena) Reset()         { *m = Athena{} }
func (m *Athena) String() string { return proto.CompactTextString(m) }
func (*Athena) ProtoMessage()    {}
func (*Athena) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{2}
}

func (m *Athena) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Athena.Unmarshal(m, b)
}
func (m *Athena) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Athena.Marshal(b, m, deterministic)
}
func (m *Athena) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Athena.Merge(m, src)
}
func (m *Athena) XXX_Size() int {
	return xxx_messageInfo_Athena.Size(m)
}
func (m *Athena) XXX_DiscardUnknown() {
	xxx_messageInfo_Athena.DiscardUnknown(m)
}

var xxx_messageInfo_Athena proto.InternalMessageInfo

func (m *Athena) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Athena) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Athena) GetSecretAccessKey() string {
	if m != nil {
		return m.SecretAccessKey
	}
	return ""
}

func (m *Athena) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Athena) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*Driver)(nil), "v1.Driver")
	proto.RegisterType((*Mysql)(nil), "v1.Mysql")
	proto.RegisterType((*Athena)(nil), "v1.Athena")
}

func init() { proto.RegisterFile("drivers.proto", fileDescriptor_81dfd49b5b303fb4) }

var fileDescriptor_81dfd49b5b303fb4 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x9b, 0xda, 0xa4, 0xed, 0x93, 0x52, 0x98, 0x55, 0x10, 0x91, 0x1a, 0x44, 0xea, 0x26,
	0x52, 0x75, 0xe5, 0xae, 0xc5, 0x45, 0x41, 0x04, 0xe9, 0x05, 0x64, 0x9a, 0x0c, 0x69, 0xb0, 0xc9,
	0x8c, 0xf3, 0x26, 0x95, 0xde, 0xc8, 0xa5, 0xd0, 0xa3, 0x78, 0x06, 0x2f, 0xa0, 0x1b, 0x71, 0x25,
	0x33, 0x93, 0x88, 0xfd, 0xb3, 0x70, 0xfb, 0xfd, 0x7e, 0x2f, 0xc9, 0x97, 0xf7, 0xa0, 0x13, 0xcb,
	0x74, 0xc1, 0x24, 0x86, 0x42, 0x72, 0xc5, 0x49, 0x7d, 0x31, 0x38, 0xe8, 0x70, 0xa1, 0x52, 0x9e,
	0x97, 0x51, 0x90, 0x80, 0x77, 0x63, 0x1c, 0x72, 0x0c, 0x6e, 0xb6, 0xc4, 0xa7, 0xb9, 0xef, 0xf4,
	0x9c, 0xfe, 0xfe, 0x45, 0x3b, 0x5c, 0x0c, 0xc2, 0x3b, 0x1d, 0x8c, 0x6b, 0x13, 0x4b, 0xc8, 0x09,
	0x78, 0x54, 0xcd, 0x58, 0x4e, 0xfd, 0xba, 0x71, 0x40, 0x3b, 0x43, 0x93, 0x8c, 0x6b, 0x93, 0x92,
	0x5d, 0xc3, 0xf7, 0xd7, 0xaa, 0xe9, 0xbe, 0x7c, 0xae, 0x9a, 0xce, 0xa8, 0x05, 0x9e, 0xfd, 0x84,
	0xe0, 0xcd, 0x01, 0xd7, 0x3c, 0x8e, 0x9c, 0x42, 0x6b, 0xc6, 0x51, 0xe5, 0x34, 0x63, 0xe6, 0x5d,
	0xed, 0x11, 0x7c, 0xe8, 0x89, 0x57, 0x3d, 0x31, 0xf9, 0x65, 0xda, 0x2b, 0x90, 0x49, 0xe3, 0xd5,
	0xb7, 0xbd, 0x8a, 0x69, 0x4f, 0x50, 0xc4, 0x67, 0x2e, 0x63, 0x7f, 0x6f, 0xdb, 0xab, 0x98, 0xf6,
	0x62, 0xaa, 0xe8, 0x94, 0x22, 0xf3, 0x1b, 0xdb, 0x5e, 0xc5, 0xc8, 0x11, 0x34, 0x04, 0x97, 0xca,
	0x77, 0x7b, 0x4e, 0xdf, 0x5d, 0x73, 0x4c, 0xfe, 0xb7, 0x5f, 0xf0, 0xee, 0x80, 0x67, 0x7f, 0xc0,
	0xbf, 0x6b, 0x9d, 0x01, 0xd0, 0x28, 0x62, 0x88, 0x0f, 0x8f, 0x6c, 0xb9, 0xa3, 0x58, 0xdb, 0xd2,
	0x5b, 0xb6, 0x24, 0x57, 0xd0, 0x45, 0x16, 0x49, 0xa6, 0x86, 0x55, 0xb4, 0xa3, 0xe0, 0xa6, 0x42,
	0x02, 0xf0, 0x24, 0x4b, 0x52, 0x9e, 0xef, 0x68, 0x59, 0x12, 0xed, 0xf0, 0x42, 0x89, 0xc2, 0xb6,
	0xdc, 0x70, 0x2c, 0x59, 0xdb, 0xe3, 0x39, 0x1c, 0x46, 0x3c, 0x0b, 0x51, 0x49, 0x9e, 0x27, 0x71,
	0x16, 0x52, 0x91, 0xea, 0xdd, 0x8b, 0x79, 0x91, 0x4d, 0xd3, 0x3c, 0x19, 0x75, 0xed, 0x11, 0xe1,
	0x7d, 0x19, 0x4c, 0x3d, 0x73, 0x5e, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0xfa, 0x76,
	0xac, 0x82, 0x02, 0x00, 0x00,
}

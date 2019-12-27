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
	//	*Driver_AuroraMysql
	//	*Driver_Clustrix
	//	*Driver_Maria
	//	*Driver_Memsql
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

type Driver_AuroraMysql struct {
	AuroraMysql *AuroraMysql `protobuf:"bytes,2,opt,name=aurora_mysql,json=auroraMysql,proto3,oneof"`
}

type Driver_Clustrix struct {
	Clustrix *Clustrix `protobuf:"bytes,3,opt,name=clustrix,proto3,oneof"`
}

type Driver_Maria struct {
	Maria *Maria `protobuf:"bytes,4,opt,name=maria,proto3,oneof"`
}

type Driver_Memsql struct {
	Memsql *Memsql `protobuf:"bytes,5,opt,name=memsql,proto3,oneof"`
}

type Driver_Athena struct {
	Athena *Athena `protobuf:"bytes,6,opt,name=athena,proto3,oneof"`
}

func (*Driver_Mysql) isDriver_Driver() {}

func (*Driver_AuroraMysql) isDriver_Driver() {}

func (*Driver_Clustrix) isDriver_Driver() {}

func (*Driver_Maria) isDriver_Driver() {}

func (*Driver_Memsql) isDriver_Driver() {}

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

func (m *Driver) GetAuroraMysql() *AuroraMysql {
	if x, ok := m.GetDriver().(*Driver_AuroraMysql); ok {
		return x.AuroraMysql
	}
	return nil
}

func (m *Driver) GetClustrix() *Clustrix {
	if x, ok := m.GetDriver().(*Driver_Clustrix); ok {
		return x.Clustrix
	}
	return nil
}

func (m *Driver) GetMaria() *Maria {
	if x, ok := m.GetDriver().(*Driver_Maria); ok {
		return x.Maria
	}
	return nil
}

func (m *Driver) GetMemsql() *Memsql {
	if x, ok := m.GetDriver().(*Driver_Memsql); ok {
		return x.Memsql
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
		(*Driver_AuroraMysql)(nil),
		(*Driver_Clustrix)(nil),
		(*Driver_Maria)(nil),
		(*Driver_Memsql)(nil),
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

type AuroraMysql struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Port                 int32    `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuroraMysql) Reset()         { *m = AuroraMysql{} }
func (m *AuroraMysql) String() string { return proto.CompactTextString(m) }
func (*AuroraMysql) ProtoMessage()    {}
func (*AuroraMysql) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{2}
}

func (m *AuroraMysql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuroraMysql.Unmarshal(m, b)
}
func (m *AuroraMysql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuroraMysql.Marshal(b, m, deterministic)
}
func (m *AuroraMysql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuroraMysql.Merge(m, src)
}
func (m *AuroraMysql) XXX_Size() int {
	return xxx_messageInfo_AuroraMysql.Size(m)
}
func (m *AuroraMysql) XXX_DiscardUnknown() {
	xxx_messageInfo_AuroraMysql.DiscardUnknown(m)
}

var xxx_messageInfo_AuroraMysql proto.InternalMessageInfo

func (m *AuroraMysql) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *AuroraMysql) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuroraMysql) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuroraMysql) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *AuroraMysql) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Clustrix struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Port                 int32    `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Clustrix) Reset()         { *m = Clustrix{} }
func (m *Clustrix) String() string { return proto.CompactTextString(m) }
func (*Clustrix) ProtoMessage()    {}
func (*Clustrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{3}
}

func (m *Clustrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clustrix.Unmarshal(m, b)
}
func (m *Clustrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clustrix.Marshal(b, m, deterministic)
}
func (m *Clustrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clustrix.Merge(m, src)
}
func (m *Clustrix) XXX_Size() int {
	return xxx_messageInfo_Clustrix.Size(m)
}
func (m *Clustrix) XXX_DiscardUnknown() {
	xxx_messageInfo_Clustrix.DiscardUnknown(m)
}

var xxx_messageInfo_Clustrix proto.InternalMessageInfo

func (m *Clustrix) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Clustrix) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Clustrix) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Clustrix) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Clustrix) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Maria struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Port                 int32    `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Maria) Reset()         { *m = Maria{} }
func (m *Maria) String() string { return proto.CompactTextString(m) }
func (*Maria) ProtoMessage()    {}
func (*Maria) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{4}
}

func (m *Maria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maria.Unmarshal(m, b)
}
func (m *Maria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maria.Marshal(b, m, deterministic)
}
func (m *Maria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maria.Merge(m, src)
}
func (m *Maria) XXX_Size() int {
	return xxx_messageInfo_Maria.Size(m)
}
func (m *Maria) XXX_DiscardUnknown() {
	xxx_messageInfo_Maria.DiscardUnknown(m)
}

var xxx_messageInfo_Maria proto.InternalMessageInfo

func (m *Maria) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Maria) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Maria) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Maria) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Maria) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Memsql struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database             string   `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	Port                 int32    `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memsql) Reset()         { *m = Memsql{} }
func (m *Memsql) String() string { return proto.CompactTextString(m) }
func (*Memsql) ProtoMessage()    {}
func (*Memsql) Descriptor() ([]byte, []int) {
	return fileDescriptor_81dfd49b5b303fb4, []int{5}
}

func (m *Memsql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memsql.Unmarshal(m, b)
}
func (m *Memsql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memsql.Marshal(b, m, deterministic)
}
func (m *Memsql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memsql.Merge(m, src)
}
func (m *Memsql) XXX_Size() int {
	return xxx_messageInfo_Memsql.Size(m)
}
func (m *Memsql) XXX_DiscardUnknown() {
	xxx_messageInfo_Memsql.DiscardUnknown(m)
}

var xxx_messageInfo_Memsql proto.InternalMessageInfo

func (m *Memsql) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Memsql) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Memsql) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Memsql) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *Memsql) GetPort() int32 {
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
	return fileDescriptor_81dfd49b5b303fb4, []int{6}
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
	proto.RegisterType((*AuroraMysql)(nil), "v1.AuroraMysql")
	proto.RegisterType((*Clustrix)(nil), "v1.Clustrix")
	proto.RegisterType((*Maria)(nil), "v1.Maria")
	proto.RegisterType((*Memsql)(nil), "v1.Memsql")
	proto.RegisterType((*Athena)(nil), "v1.Athena")
}

func init() { proto.RegisterFile("drivers.proto", fileDescriptor_81dfd49b5b303fb4) }

var fileDescriptor_81dfd49b5b303fb4 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x95, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0xf1, 0x66, 0x6e, 0x7a, 0x7a, 0x2f, 0x85, 0x59, 0x05, 0x11, 0xd1, 0x20, 0xa2,
	0x2e, 0x22, 0x57, 0xef, 0xca, 0x5d, 0xab, 0x8b, 0x82, 0x08, 0x92, 0x17, 0xb8, 0x4c, 0x93, 0xa1,
	0x37, 0xd8, 0x64, 0xe2, 0xcc, 0xa4, 0xda, 0xd7, 0xf0, 0x29, 0x5c, 0x0a, 0x7d, 0x14, 0x45, 0xf0,
	0x01, 0x7c, 0x01, 0xdd, 0x88, 0x2b, 0x99, 0x33, 0x49, 0xff, 0xc4, 0x2e, 0xee, 0x3a, 0xbb, 0xf6,
	0xfb, 0x7e, 0x87, 0x9c, 0xef, 0x64, 0x72, 0x06, 0xce, 0x33, 0x95, 0xaf, 0x84, 0xd2, 0x71, 0xa5,
	0xa4, 0x91, 0xcc, 0x5b, 0x5d, 0xdc, 0x3e, 0x97, 0x95, 0xc9, 0x65, 0xd9, 0x48, 0xd1, 0x27, 0x0f,
	0xe8, 0x2b, 0x84, 0xd8, 0x7d, 0xf0, 0x8b, 0xb5, 0x7e, 0xbf, 0x0c, 0xc9, 0x3d, 0xf2, 0x68, 0xf4,
	0x6c, 0x18, 0xaf, 0x2e, 0xe2, 0x37, 0x56, 0x98, 0x0d, 0x12, 0xe7, 0xb0, 0x4b, 0x38, 0xe3, 0xb5,
	0x92, 0x8a, 0x5f, 0x39, 0xd2, 0x43, 0x72, 0x6c, 0xc9, 0x09, 0xea, 0x2d, 0x3f, 0xe2, 0xbb, 0xbf,
	0xec, 0x09, 0x04, 0xe9, 0xb2, 0xd6, 0x46, 0xe5, 0x1f, 0xc3, 0x5b, 0x58, 0x71, 0x66, 0x2b, 0x5e,
	0x36, 0xda, 0x6c, 0x90, 0x6c, 0x7d, 0x6c, 0x82, 0xab, 0x9c, 0x87, 0x27, 0x7b, 0x4d, 0x58, 0x01,
	0x9b, 0xb0, 0x3f, 0xd8, 0x03, 0xa0, 0x85, 0x28, 0xec, 0xe3, 0x7d, 0x64, 0x00, 0x19, 0x54, 0x66,
	0x83, 0xa4, 0xf1, 0x2c, 0xc5, 0xcd, 0xb5, 0x28, 0x79, 0x48, 0x77, 0xd4, 0x04, 0x15, 0x4b, 0x39,
	0xef, 0x05, 0xfc, 0xfd, 0xb3, 0x39, 0xf5, 0x3f, 0xff, 0xde, 0x9c, 0x92, 0x69, 0x00, 0xd4, 0x8d,
	0x2b, 0xfa, 0x4a, 0xc0, 0x77, 0xad, 0x3f, 0x84, 0xe0, 0x5a, 0x6a, 0x53, 0xf2, 0x42, 0xe0, 0x58,
	0x86, 0x53, 0xf8, 0x65, 0x2b, 0xbe, 0xd8, 0x8a, 0x64, 0xeb, 0x59, 0xae, 0xd6, 0x42, 0x21, 0xe7,
	0xfd, 0xcf, 0xb5, 0x9e, 0xe5, 0x2a, 0xae, 0xf5, 0x07, 0xa9, 0x32, 0x1c, 0x45, 0x87, 0x6b, 0x3d,
	0xcb, 0x65, 0xdc, 0xf0, 0x39, 0xd7, 0x02, 0x27, 0xd1, 0xe1, 0x5a, 0x8f, 0xdd, 0x85, 0x93, 0x4a,
	0x2a, 0x83, 0x93, 0xf0, 0x0f, 0x18, 0xd4, 0xf7, 0xf3, 0x45, 0x3f, 0x08, 0x8c, 0xf6, 0xde, 0x52,
	0xaf, 0xb2, 0x7d, 0x27, 0x10, 0xb4, 0xe7, 0xa9, 0x57, 0xc1, 0xf0, 0x28, 0xe2, 0xb1, 0xef, 0x53,
	0xaa, 0x6f, 0x04, 0xa8, 0xfb, 0x62, 0x7b, 0x15, 0xeb, 0x27, 0x01, 0xea, 0x56, 0xcc, 0x8d, 0x63,
	0x3d, 0x06, 0xe0, 0x69, 0x2a, 0xb4, 0xbe, 0x7a, 0x27, 0xd6, 0x47, 0x82, 0x0d, 0x9d, 0xfb, 0x5a,
	0xac, 0xd9, 0x25, 0x8c, 0xb5, 0x48, 0x95, 0x30, 0x93, 0x56, 0x3a, 0x12, 0xb0, 0x8b, 0xb0, 0x08,
	0xa8, 0x12, 0x8b, 0x5c, 0x96, 0x47, 0x52, 0x36, 0x8e, 0x65, 0x64, 0x6d, 0xaa, 0xda, 0xa5, 0xec,
	0x30, 0xce, 0x39, 0xd8, 0x94, 0x4f, 0xe1, 0x4e, 0x2a, 0x8b, 0x58, 0x1b, 0x25, 0xcb, 0x45, 0x56,
	0xc4, 0xbc, 0xca, 0xed, 0x76, 0xad, 0x96, 0x75, 0x31, 0xcf, 0xcb, 0xc5, 0x74, 0xec, 0x6e, 0x14,
	0xfd, 0xb6, 0x11, 0xe6, 0x14, 0x2f, 0x9b, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xb1,
	0x90, 0xb0, 0x90, 0x06, 0x00, 0x00,
}

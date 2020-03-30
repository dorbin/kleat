// Code generated by protoc-gen-go. DO NOT EDIT.
// source: http_artifact_provider.proto

package client

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

// Configuration for the HTTP artifact provider.
type HttpArtifactProvider struct {
	// Whether the HTTP artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured HTTP accounts.
	Accounts             []*HttpArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HttpArtifactProvider) Reset()         { *m = HttpArtifactProvider{} }
func (m *HttpArtifactProvider) String() string { return proto.CompactTextString(m) }
func (*HttpArtifactProvider) ProtoMessage()    {}
func (*HttpArtifactProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8493c6bd644e46, []int{0}
}

func (m *HttpArtifactProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpArtifactProvider.Unmarshal(m, b)
}
func (m *HttpArtifactProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpArtifactProvider.Marshal(b, m, deterministic)
}
func (m *HttpArtifactProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpArtifactProvider.Merge(m, src)
}
func (m *HttpArtifactProvider) XXX_Size() int {
	return xxx_messageInfo_HttpArtifactProvider.Size(m)
}
func (m *HttpArtifactProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpArtifactProvider.DiscardUnknown(m)
}

var xxx_messageInfo_HttpArtifactProvider proto.InternalMessageInfo

func (m *HttpArtifactProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *HttpArtifactProvider) GetAccounts() []*HttpArtifactAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Configuration for an HTTP artifact account. Either `username` and `password`
// or `usernamePasswordFile` should be specified as means of authentication.
type HttpArtifactAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A username for HTTP basic auth.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// A password for HTTP basic auth.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// The path to a file containing the username and password for HTTP basic
	// auth. Must be in the format `${username}:${password}`.
	UsernamePasswordFile string   `protobuf:"bytes,4,opt,name=usernamePasswordFile,proto3" json:"usernamePasswordFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpArtifactAccount) Reset()         { *m = HttpArtifactAccount{} }
func (m *HttpArtifactAccount) String() string { return proto.CompactTextString(m) }
func (*HttpArtifactAccount) ProtoMessage()    {}
func (*HttpArtifactAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8493c6bd644e46, []int{1}
}

func (m *HttpArtifactAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpArtifactAccount.Unmarshal(m, b)
}
func (m *HttpArtifactAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpArtifactAccount.Marshal(b, m, deterministic)
}
func (m *HttpArtifactAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpArtifactAccount.Merge(m, src)
}
func (m *HttpArtifactAccount) XXX_Size() int {
	return xxx_messageInfo_HttpArtifactAccount.Size(m)
}
func (m *HttpArtifactAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpArtifactAccount.DiscardUnknown(m)
}

var xxx_messageInfo_HttpArtifactAccount proto.InternalMessageInfo

func (m *HttpArtifactAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HttpArtifactAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *HttpArtifactAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *HttpArtifactAccount) GetUsernamePasswordFile() string {
	if m != nil {
		return m.UsernamePasswordFile
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpArtifactProvider)(nil), "proto.HttpArtifactProvider")
	proto.RegisterType((*HttpArtifactAccount)(nil), "proto.HttpArtifactAccount")
}

func init() { proto.RegisterFile("http_artifact_provider.proto", fileDescriptor_ee8493c6bd644e46) }

var fileDescriptor_ee8493c6bd644e46 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb6, 0x40, 0x38, 0x36, 0xd3, 0xc1, 0xaa, 0x18, 0xaa, 0x4a, 0x88, 0x4e, 0x89,
	0x54, 0x24, 0xf6, 0x32, 0x20, 0xc6, 0x2a, 0x23, 0x4b, 0x75, 0x71, 0x0e, 0x62, 0xd5, 0xb5, 0x2d,
	0xfb, 0x02, 0xbf, 0x84, 0xff, 0x8b, 0x70, 0x92, 0xaa, 0x43, 0xa7, 0xbb, 0xf7, 0xde, 0xa7, 0x77,
	0x3a, 0x78, 0x68, 0x99, 0xfd, 0x1e, 0x03, 0xeb, 0x4f, 0x54, 0xbc, 0xf7, 0xc1, 0x7d, 0xeb, 0x86,
	0x42, 0xe1, 0x83, 0x63, 0x27, 0xae, 0xd2, 0x58, 0xb5, 0x30, 0x7f, 0x67, 0xf6, 0xdb, 0x81, 0xda,
	0x0d, 0x90, 0x90, 0x70, 0x43, 0x16, 0x6b, 0x43, 0x8d, 0xcc, 0x96, 0xd9, 0x3a, 0xaf, 0x46, 0x29,
	0x5e, 0x20, 0x47, 0xa5, 0x5c, 0x67, 0x39, 0xca, 0xc9, 0x72, 0xba, 0xbe, 0xdb, 0x2c, 0xfa, 0xca,
	0xe2, 0xbc, 0x68, 0xdb, 0x23, 0xd5, 0x89, 0x5d, 0xfd, 0x66, 0x70, 0x7f, 0x81, 0x10, 0x02, 0x66,
	0x16, 0x8f, 0x94, 0xce, 0xdc, 0x56, 0x69, 0x17, 0x0b, 0xc8, 0xbb, 0x48, 0x21, 0xf9, 0x93, 0xe4,
	0x9f, 0xf4, 0x7f, 0xe6, 0x31, 0xc6, 0x1f, 0x17, 0x1a, 0x39, 0xed, 0xb3, 0x51, 0x8b, 0x0d, 0xcc,
	0x47, 0x6e, 0x37, 0x78, 0x6f, 0xda, 0x90, 0x9c, 0x25, 0xee, 0x62, 0xf6, 0xfa, 0xf4, 0xf1, 0xf8,
	0xa5, 0xb9, 0xed, 0xea, 0x42, 0xb9, 0x63, 0x19, 0xbd, 0xb6, 0x16, 0x0f, 0x14, 0xca, 0x83, 0x21,
	0xe4, 0x12, 0xbd, 0x2e, 0x95, 0xd1, 0x64, 0xb9, 0xbe, 0x4e, 0x5f, 0x3e, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x25, 0x2b, 0xf2, 0x65, 0x58, 0x01, 0x00, 0x00,
}
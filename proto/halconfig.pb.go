// Code generated by protoc-gen-go. DO NOT EDIT.
// source: halconfig.proto

package proto

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

type HalConfig struct {
	PersistentStorage    *HalConfig_PersistentStorage `protobuf:"bytes,1,opt,name=persistentStorage,proto3" json:"persistentStorage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *HalConfig) Reset()         { *m = HalConfig{} }
func (m *HalConfig) String() string { return proto.CompactTextString(m) }
func (*HalConfig) ProtoMessage()    {}
func (*HalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0}
}

func (m *HalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig.Unmarshal(m, b)
}
func (m *HalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig.Marshal(b, m, deterministic)
}
func (m *HalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig.Merge(m, src)
}
func (m *HalConfig) XXX_Size() int {
	return xxx_messageInfo_HalConfig.Size(m)
}
func (m *HalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig proto.InternalMessageInfo

func (m *HalConfig) GetPersistentStorage() *HalConfig_PersistentStorage {
	if m != nil {
		return m.PersistentStorage
	}
	return nil
}

type HalConfig_PersistentStorage struct {
	PersistentStoreType  string                           `protobuf:"bytes,1,opt,name=persistentStoreType,proto3" json:"persistentStoreType,omitempty"`
	Gcs                  *HalConfig_PersistentStorage_GCS `protobuf:"bytes,2,opt,name=gcs,proto3" json:"gcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *HalConfig_PersistentStorage) Reset()         { *m = HalConfig_PersistentStorage{} }
func (m *HalConfig_PersistentStorage) String() string { return proto.CompactTextString(m) }
func (*HalConfig_PersistentStorage) ProtoMessage()    {}
func (*HalConfig_PersistentStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0, 0}
}

func (m *HalConfig_PersistentStorage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig_PersistentStorage.Unmarshal(m, b)
}
func (m *HalConfig_PersistentStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig_PersistentStorage.Marshal(b, m, deterministic)
}
func (m *HalConfig_PersistentStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig_PersistentStorage.Merge(m, src)
}
func (m *HalConfig_PersistentStorage) XXX_Size() int {
	return xxx_messageInfo_HalConfig_PersistentStorage.Size(m)
}
func (m *HalConfig_PersistentStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig_PersistentStorage.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig_PersistentStorage proto.InternalMessageInfo

func (m *HalConfig_PersistentStorage) GetPersistentStoreType() string {
	if m != nil {
		return m.PersistentStoreType
	}
	return ""
}

func (m *HalConfig_PersistentStorage) GetGcs() *HalConfig_PersistentStorage_GCS {
	if m != nil {
		return m.Gcs
	}
	return nil
}

type HalConfig_PersistentStorage_GCS struct {
	JsonPath string `protobuf:"bytes,1,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	Project  string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The name of a storage bucket that your specified account has access to. If not specified, a
	// random name will be chosen. If you specify a globally unique bucket name that doesn’t exist
	// yet, Halyard will create that bucket for you.
	Bucket               string   `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	RootFolder           string   `protobuf:"bytes,4,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	BucketLocation       string   `protobuf:"bytes,5,opt,name=bucketLocation,proto3" json:"bucketLocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HalConfig_PersistentStorage_GCS) Reset()         { *m = HalConfig_PersistentStorage_GCS{} }
func (m *HalConfig_PersistentStorage_GCS) String() string { return proto.CompactTextString(m) }
func (*HalConfig_PersistentStorage_GCS) ProtoMessage()    {}
func (*HalConfig_PersistentStorage_GCS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0, 0, 0}
}

func (m *HalConfig_PersistentStorage_GCS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig_PersistentStorage_GCS.Unmarshal(m, b)
}
func (m *HalConfig_PersistentStorage_GCS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig_PersistentStorage_GCS.Marshal(b, m, deterministic)
}
func (m *HalConfig_PersistentStorage_GCS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig_PersistentStorage_GCS.Merge(m, src)
}
func (m *HalConfig_PersistentStorage_GCS) XXX_Size() int {
	return xxx_messageInfo_HalConfig_PersistentStorage_GCS.Size(m)
}
func (m *HalConfig_PersistentStorage_GCS) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig_PersistentStorage_GCS.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig_PersistentStorage_GCS proto.InternalMessageInfo

func (m *HalConfig_PersistentStorage_GCS) GetJsonPath() string {
	if m != nil {
		return m.JsonPath
	}
	return ""
}

func (m *HalConfig_PersistentStorage_GCS) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *HalConfig_PersistentStorage_GCS) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *HalConfig_PersistentStorage_GCS) GetRootFolder() string {
	if m != nil {
		return m.RootFolder
	}
	return ""
}

func (m *HalConfig_PersistentStorage_GCS) GetBucketLocation() string {
	if m != nil {
		return m.BucketLocation
	}
	return ""
}

func init() {
	proto.RegisterType((*HalConfig)(nil), "proto.HalConfig")
	proto.RegisterType((*HalConfig_PersistentStorage)(nil), "proto.HalConfig.PersistentStorage")
	proto.RegisterType((*HalConfig_PersistentStorage_GCS)(nil), "proto.HalConfig.PersistentStorage.GCS")
}

func init() { proto.RegisterFile("halconfig.proto", fileDescriptor_cda1892e07435feb) }

var fileDescriptor_cda1892e07435feb = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x48, 0xcc, 0x49,
	0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0x7f, 0x99, 0xb8, 0x38, 0x3d, 0x12, 0x73, 0x9c, 0xc1, 0x52, 0x42, 0x01, 0x5c, 0x82, 0x05, 0xa9,
	0x45, 0xc5, 0x99, 0xc5, 0x25, 0xa9, 0x79, 0x25, 0xc1, 0x25, 0xf9, 0x45, 0x89, 0xe9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x4a, 0x10, 0x7d, 0x7a, 0x70, 0xc5, 0x7a, 0x01, 0xe8, 0x2a,
	0x83, 0x30, 0x35, 0x4b, 0xcd, 0x65, 0xe2, 0x12, 0xc4, 0x50, 0x28, 0x64, 0xc0, 0x25, 0x8c, 0xaa,
	0x34, 0x35, 0xa4, 0xb2, 0x00, 0x62, 0x13, 0x67, 0x10, 0x36, 0x29, 0x21, 0x0b, 0x2e, 0xe6, 0xf4,
	0xe4, 0x62, 0x09, 0x26, 0xb0, 0x5b, 0xd4, 0x08, 0xbb, 0x45, 0xcf, 0xdd, 0x39, 0x38, 0x08, 0xa4,
	0x45, 0x6a, 0x36, 0x23, 0x17, 0xb3, 0xbb, 0x73, 0xb0, 0x90, 0x14, 0x17, 0x47, 0x56, 0x71, 0x7e,
	0x5e, 0x40, 0x62, 0x49, 0x06, 0xd4, 0x22, 0x38, 0x5f, 0x48, 0x82, 0x8b, 0xbd, 0xa0, 0x28, 0x3f,
	0x2b, 0x35, 0xb9, 0x04, 0x6c, 0x03, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x96, 0x54, 0x9a,
	0x9c, 0x9d, 0x5a, 0x22, 0xc1, 0x0c, 0x96, 0x80, 0xf2, 0x84, 0xe4, 0xb8, 0xb8, 0x8a, 0xf2, 0xf3,
	0x4b, 0xdc, 0xf2, 0x73, 0x52, 0x52, 0x8b, 0x24, 0x58, 0xc0, 0x72, 0x48, 0x22, 0x42, 0x6a, 0x5c,
	0x7c, 0x10, 0x95, 0x3e, 0xf9, 0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xac, 0x60, 0x35, 0x68,
	0xa2, 0x49, 0x6c, 0x60, 0x9f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xeb, 0xc6, 0x62,
	0xa0, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcos_provider.proto

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

// Configuration for the DC/OS (Distributed Cloud Operating System) provider.
type DcosProvider struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*DcosAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// The list of configured clusters.
	Clusters             []*DcosCluster `protobuf:"bytes,4,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DcosProvider) Reset()         { *m = DcosProvider{} }
func (m *DcosProvider) String() string { return proto.CompactTextString(m) }
func (*DcosProvider) ProtoMessage()    {}
func (*DcosProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{0}
}

func (m *DcosProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosProvider.Unmarshal(m, b)
}
func (m *DcosProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosProvider.Marshal(b, m, deterministic)
}
func (m *DcosProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosProvider.Merge(m, src)
}
func (m *DcosProvider) XXX_Size() int {
	return xxx_messageInfo_DcosProvider.Size(m)
}
func (m *DcosProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosProvider.DiscardUnknown(m)
}

var xxx_messageInfo_DcosProvider proto.InternalMessageInfo

func (m *DcosProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DcosProvider) GetAccounts() []*DcosAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *DcosProvider) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

func (m *DcosProvider) GetClusters() []*DcosCluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Credentials to authenticate against one or more DC/OS clusters.
type DcosAccount struct {
	// (Required) The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The clusters against which this account will authenticate.
	Clusters []*DcosAccountCluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	// (Required) The list of Docker registries to use with this DC/OS account.
	DockerRegistries []*DcosAccountDockerRegistry `protobuf:"bytes,4,rep,name=dockerRegistries,proto3" json:"dockerRegistries,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,6,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *DcosAccount) Reset()         { *m = DcosAccount{} }
func (m *DcosAccount) String() string { return proto.CompactTextString(m) }
func (*DcosAccount) ProtoMessage()    {}
func (*DcosAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{1}
}

func (m *DcosAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosAccount.Unmarshal(m, b)
}
func (m *DcosAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosAccount.Marshal(b, m, deterministic)
}
func (m *DcosAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosAccount.Merge(m, src)
}
func (m *DcosAccount) XXX_Size() int {
	return xxx_messageInfo_DcosAccount.Size(m)
}
func (m *DcosAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DcosAccount proto.InternalMessageInfo

func (m *DcosAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DcosAccount) GetClusters() []*DcosAccountCluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *DcosAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *DcosAccount) GetDockerRegistries() []*DcosAccountDockerRegistry {
	if m != nil {
		return m.DockerRegistries
	}
	return nil
}

func (m *DcosAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *DcosAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

// Configuration for a DC/OS cluster associated with a `DcosAccount`.
type DcosAccountCluster struct {
	// (Required) The name of the cluster. Must match the name of a
	// `DcosCluster` defined for this provider.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) User or service account identifier.
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// Path to a file containing the secret key for service account
	// authentication. If set, `password` should not be set.
	ServiceKeyFile string `protobuf:"bytes,3,opt,name=serviceKeyFile,proto3" json:"serviceKeyFile,omitempty"`
	// Password for a user account. If set, `serviceKeyFile` should not be set.
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcosAccountCluster) Reset()         { *m = DcosAccountCluster{} }
func (m *DcosAccountCluster) String() string { return proto.CompactTextString(m) }
func (*DcosAccountCluster) ProtoMessage()    {}
func (*DcosAccountCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{2}
}

func (m *DcosAccountCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosAccountCluster.Unmarshal(m, b)
}
func (m *DcosAccountCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosAccountCluster.Marshal(b, m, deterministic)
}
func (m *DcosAccountCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosAccountCluster.Merge(m, src)
}
func (m *DcosAccountCluster) XXX_Size() int {
	return xxx_messageInfo_DcosAccountCluster.Size(m)
}
func (m *DcosAccountCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosAccountCluster.DiscardUnknown(m)
}

var xxx_messageInfo_DcosAccountCluster proto.InternalMessageInfo

func (m *DcosAccountCluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DcosAccountCluster) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DcosAccountCluster) GetServiceKeyFile() string {
	if m != nil {
		return m.ServiceKeyFile
	}
	return ""
}

func (m *DcosAccountCluster) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Configuration for a DC/OS cluster.
type DcosCluster struct {
	// (Required) The name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Root certificate file to trust for connections to the cluster.
	CaCertFile string `protobuf:"bytes,2,opt,name=caCertFile,proto3" json:"caCertFile,omitempty"`
	// (Required) URL of the endpoint for the DC/OS cluster's admin router.
	DcosUrl string `protobuf:"bytes,3,opt,name=dcosUrl,proto3" json:"dcosUrl,omitempty"`
	// Configuration for a DC/OS load balancer.
	LoadBalancer *DcosClusterLoadBalancer `protobuf:"bytes,4,opt,name=loadBalancer,proto3" json:"loadBalancer,omitempty"`
	// If `true`, disables verification of certificates from the cluster
	// (insecure).
	InsecureSkipTlsVerify bool     `protobuf:"varint,5,opt,name=insecureSkipTlsVerify,proto3" json:"insecureSkipTlsVerify,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *DcosCluster) Reset()         { *m = DcosCluster{} }
func (m *DcosCluster) String() string { return proto.CompactTextString(m) }
func (*DcosCluster) ProtoMessage()    {}
func (*DcosCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{3}
}

func (m *DcosCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosCluster.Unmarshal(m, b)
}
func (m *DcosCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosCluster.Marshal(b, m, deterministic)
}
func (m *DcosCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosCluster.Merge(m, src)
}
func (m *DcosCluster) XXX_Size() int {
	return xxx_messageInfo_DcosCluster.Size(m)
}
func (m *DcosCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosCluster.DiscardUnknown(m)
}

var xxx_messageInfo_DcosCluster proto.InternalMessageInfo

func (m *DcosCluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DcosCluster) GetCaCertFile() string {
	if m != nil {
		return m.CaCertFile
	}
	return ""
}

func (m *DcosCluster) GetDcosUrl() string {
	if m != nil {
		return m.DcosUrl
	}
	return ""
}

func (m *DcosCluster) GetLoadBalancer() *DcosClusterLoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *DcosCluster) GetInsecureSkipTlsVerify() bool {
	if m != nil {
		return m.InsecureSkipTlsVerify
	}
	return false
}

// Configuration for a DC/OS load balancer.
type DcosClusterLoadBalancer struct {
	// Marathon-lb image to use when creating a load balancer with Spinnaker.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Name of the secret to use for allowing marathon-lb to authenticate with
	// the cluster. Only necessary for clusters with strict or permissive
	// security.
	ServiceAccountSecret string   `protobuf:"bytes,2,opt,name=serviceAccountSecret,proto3" json:"serviceAccountSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcosClusterLoadBalancer) Reset()         { *m = DcosClusterLoadBalancer{} }
func (m *DcosClusterLoadBalancer) String() string { return proto.CompactTextString(m) }
func (*DcosClusterLoadBalancer) ProtoMessage()    {}
func (*DcosClusterLoadBalancer) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{4}
}

func (m *DcosClusterLoadBalancer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosClusterLoadBalancer.Unmarshal(m, b)
}
func (m *DcosClusterLoadBalancer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosClusterLoadBalancer.Marshal(b, m, deterministic)
}
func (m *DcosClusterLoadBalancer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosClusterLoadBalancer.Merge(m, src)
}
func (m *DcosClusterLoadBalancer) XXX_Size() int {
	return xxx_messageInfo_DcosClusterLoadBalancer.Size(m)
}
func (m *DcosClusterLoadBalancer) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosClusterLoadBalancer.DiscardUnknown(m)
}

var xxx_messageInfo_DcosClusterLoadBalancer proto.InternalMessageInfo

func (m *DcosClusterLoadBalancer) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DcosClusterLoadBalancer) GetServiceAccountSecret() string {
	if m != nil {
		return m.ServiceAccountSecret
	}
	return ""
}

// Configuration for a Docker registry associated with a `DcosAccount`.
type DcosAccountDockerRegistry struct {
	// The name of the Docker registry. Must be the name of an account
	// configured with the Docker registry provider.
	AccountName          string   `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DcosAccountDockerRegistry) Reset()         { *m = DcosAccountDockerRegistry{} }
func (m *DcosAccountDockerRegistry) String() string { return proto.CompactTextString(m) }
func (*DcosAccountDockerRegistry) ProtoMessage()    {}
func (*DcosAccountDockerRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_589956762bdc777b, []int{5}
}

func (m *DcosAccountDockerRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DcosAccountDockerRegistry.Unmarshal(m, b)
}
func (m *DcosAccountDockerRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DcosAccountDockerRegistry.Marshal(b, m, deterministic)
}
func (m *DcosAccountDockerRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DcosAccountDockerRegistry.Merge(m, src)
}
func (m *DcosAccountDockerRegistry) XXX_Size() int {
	return xxx_messageInfo_DcosAccountDockerRegistry.Size(m)
}
func (m *DcosAccountDockerRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DcosAccountDockerRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DcosAccountDockerRegistry proto.InternalMessageInfo

func (m *DcosAccountDockerRegistry) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func init() {
	proto.RegisterType((*DcosProvider)(nil), "proto.DcosProvider")
	proto.RegisterType((*DcosAccount)(nil), "proto.DcosAccount")
	proto.RegisterType((*DcosAccountCluster)(nil), "proto.DcosAccountCluster")
	proto.RegisterType((*DcosCluster)(nil), "proto.DcosCluster")
	proto.RegisterType((*DcosClusterLoadBalancer)(nil), "proto.DcosClusterLoadBalancer")
	proto.RegisterType((*DcosAccountDockerRegistry)(nil), "proto.DcosAccountDockerRegistry")
}

func init() { proto.RegisterFile("dcos_provider.proto", fileDescriptor_589956762bdc777b) }

var fileDescriptor_589956762bdc777b = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xa4, 0x24, 0x93, 0x0a, 0x95, 0xa5, 0x08, 0xb7, 0x87, 0xca, 0xb2, 0x04, 0xe4,
	0x94, 0x48, 0xa1, 0x5c, 0x90, 0x38, 0x90, 0x56, 0x70, 0xa0, 0xa0, 0x6a, 0x0b, 0x1c, 0xb8, 0xa0,
	0xcd, 0x7a, 0x48, 0x57, 0xb1, 0x77, 0xcd, 0xac, 0x1d, 0x94, 0x0b, 0xbf, 0x89, 0x33, 0xff, 0x83,
	0xff, 0x83, 0xfc, 0x95, 0x6e, 0x71, 0xc2, 0xc9, 0x3b, 0xf3, 0x66, 0xdf, 0xbc, 0x9d, 0x37, 0x86,
	0x87, 0x91, 0x34, 0xf6, 0x6b, 0x4a, 0x66, 0xa5, 0x22, 0xa4, 0x71, 0x4a, 0x26, 0x33, 0xac, 0x57,
	0x7e, 0x4e, 0x1e, 0xa4, 0x48, 0x89, 0xb2, 0x56, 0x19, 0x6d, 0x2b, 0x24, 0xfc, 0xe5, 0xc1, 0xc1,
	0x85, 0x34, 0xf6, 0xaa, 0xbe, 0xc0, 0x7c, 0xb8, 0x87, 0x5a, 0xcc, 0x63, 0x8c, 0x7c, 0x2f, 0xf0,
	0x46, 0x7d, 0xde, 0x84, 0x6c, 0x0c, 0x7d, 0x21, 0xa5, 0xc9, 0x75, 0x66, 0xfd, 0xbd, 0xa0, 0x33,
	0x1a, 0x4e, 0x59, 0x45, 0x32, 0x2e, 0x08, 0x5e, 0x57, 0x10, 0xdf, 0xd4, 0xb0, 0xa7, 0x70, 0x3f,
	0x25, 0x95, 0x08, 0x5a, 0xd7, 0x98, 0xdf, 0x09, 0xbc, 0xd1, 0x80, 0xff, 0x93, 0x2d, 0x78, 0x65,
	0x9c, 0xdb, 0x0c, 0xc9, 0xfa, 0xdd, 0x16, 0xef, 0x79, 0x05, 0xf1, 0x4d, 0x4d, 0xf8, 0x7b, 0x0f,
	0x86, 0x4e, 0x47, 0xc6, 0xa0, 0xab, 0x45, 0x82, 0xa5, 0xdc, 0x01, 0x2f, 0xcf, 0xec, 0x85, 0xc3,
	0x59, 0x69, 0x3d, 0x6e, 0x6b, 0x6d, 0x51, 0xb3, 0x00, 0x86, 0xa8, 0x57, 0x8a, 0x8c, 0x4e, 0x70,
	0xa3, 0xd7, 0x4d, 0xb1, 0x4b, 0x38, 0x8c, 0x8c, 0x5c, 0x22, 0x71, 0x5c, 0x28, 0x9b, 0x91, 0xc2,
	0x46, 0x74, 0xd0, 0x6e, 0x70, 0xe1, 0x56, 0xae, 0x79, 0xeb, 0x26, 0x3b, 0x83, 0xa1, 0x63, 0x89,
	0xdf, 0x0b, 0x3c, 0xe7, 0xf5, 0x57, 0xb7, 0x08, 0x77, 0xcb, 0xd8, 0x4b, 0xf0, 0x09, 0xbf, 0xe7,
	0x8a, 0x30, 0x7a, 0x4b, 0x26, 0x4f, 0xdf, 0x63, 0x32, 0x47, 0xb2, 0x37, 0x2a, 0xb5, 0xfe, 0x7e,
	0xd0, 0x19, 0x0d, 0xf8, 0x4e, 0x3c, 0xfc, 0x09, 0xac, 0x3d, 0x81, 0xad, 0x23, 0x3c, 0x84, 0x4e,
	0xae, 0x22, 0x7f, 0xaf, 0x4c, 0x15, 0xc7, 0xc2, 0x50, 0x8b, 0xb4, 0x52, 0x12, 0xdf, 0xe1, 0xfa,
	0x8d, 0x8a, 0xb1, 0x31, 0xf4, 0x6e, 0x96, 0x9d, 0x40, 0x3f, 0x15, 0xd6, 0xfe, 0x30, 0x14, 0xf9,
	0xdd, 0xb2, 0x62, 0x13, 0x87, 0x7f, 0xbc, 0xca, 0xbc, 0xff, 0x75, 0x3e, 0x05, 0x90, 0xe2, 0x1c,
	0x29, 0x2b, 0x7b, 0x54, 0x02, 0x9c, 0x4c, 0xb1, 0xa2, 0xc5, 0x92, 0x7f, 0xa2, 0xb8, 0x16, 0xd0,
	0x84, 0x6c, 0x06, 0x07, 0xb1, 0x11, 0xd1, 0x4c, 0xc4, 0x42, 0x4b, 0xa4, 0xb2, 0xfb, 0x70, 0x7a,
	0xda, 0x5e, 0xa7, 0x4b, 0xa7, 0x8a, 0xdf, 0xb9, 0xc3, 0xce, 0xe0, 0x91, 0xd2, 0x16, 0x65, 0x4e,
	0x78, 0xbd, 0x54, 0xe9, 0xc7, 0xd8, 0x7e, 0x46, 0x52, 0xdf, 0xd6, 0xa5, 0x3b, 0x7d, 0xbe, 0x1d,
	0x0c, 0x25, 0x3c, 0xde, 0x41, 0xcf, 0x8e, 0xa0, 0xa7, 0x12, 0xb1, 0x68, 0xde, 0x58, 0x05, 0x6c,
	0x0a, 0x47, 0xf5, 0xd8, 0x6a, 0x2f, 0xae, 0x51, 0x12, 0x66, 0xf5, 0x73, 0xb7, 0x62, 0xe1, 0x2b,
	0x38, 0xde, 0xb9, 0x5d, 0xc5, 0xee, 0xd6, 0xbf, 0xde, 0x87, 0xdb, 0x81, 0xba, 0xa9, 0xd9, 0xb3,
	0x2f, 0x4f, 0x16, 0x2a, 0xbb, 0xc9, 0xe7, 0x63, 0x69, 0x92, 0x89, 0x4d, 0x95, 0xd6, 0x62, 0x89,
	0x34, 0x59, 0xc6, 0x28, 0xb2, 0x89, 0x48, 0xd5, 0x44, 0xc6, 0x0a, 0x75, 0x36, 0xdf, 0x2f, 0xe7,
	0xf5, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0xcd, 0xb2, 0xb6, 0x4c, 0x04, 0x00, 0x00,
}
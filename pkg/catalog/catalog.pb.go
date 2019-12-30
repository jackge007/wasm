// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/catalog/catalog.proto

package catalog

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

type ExtensionSpec struct {
	// The name of the extension. Must be unique within a repository.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Module logo URL. optional
	LogoUrl string `protobuf:"bytes,2,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	// Short description of the extension.
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	// Long description of the extension.
	// TODO support colocated markdown file OR link to markdown file
	LongDescription string `protobuf:"bytes,4,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	// Link to documentation for the extension.
	DocumentationUrl string `protobuf:"bytes,5,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	// Link to repository containing the extension's source code.
	RepositoryUrl string `protobuf:"bytes,6,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty"`
	// Link to compiled wasm module for this extension.
	ExtensionRef string `protobuf:"bytes,7,opt,name=extension_ref,json=extensionRef,proto3" json:"extension_ref,omitempty"`
	// The name of the entity that authored the extension.
	CreatorName string `protobuf:"bytes,8,opt,name=creator_name,json=creatorName,proto3" json:"creator_name,omitempty"`
	// Link to the entity that authored the extension.
	CreatorUrl string `protobuf:"bytes,9,opt,name=creator_url,json=creatorUrl,proto3" json:"creator_url,omitempty"`
	// Install instructions for this extension.
	InstallInstructions *InstallInstructions `protobuf:"bytes,10,opt,name=install_instructions,json=installInstructions,proto3" json:"install_instructions,omitempty"`
	// List of tags to help identify the extension.
	// TODO joekelley do not merge -- this should be a separate file, meta.yaml or similar, not per-version?
	Tags []string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	// Whether this extension is coming soon
	// TODO should this also come from a separate file?
	ComingSoon           bool     `protobuf:"varint,12,opt,name=coming_soon,json=comingSoon,proto3" json:"coming_soon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionSpec) Reset()         { *m = ExtensionSpec{} }
func (m *ExtensionSpec) String() string { return proto.CompactTextString(m) }
func (*ExtensionSpec) ProtoMessage()    {}
func (*ExtensionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb76b5cc14a8263, []int{0}
}

func (m *ExtensionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionSpec.Unmarshal(m, b)
}
func (m *ExtensionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionSpec.Marshal(b, m, deterministic)
}
func (m *ExtensionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionSpec.Merge(m, src)
}
func (m *ExtensionSpec) XXX_Size() int {
	return xxx_messageInfo_ExtensionSpec.Size(m)
}
func (m *ExtensionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionSpec proto.InternalMessageInfo

func (m *ExtensionSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtensionSpec) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *ExtensionSpec) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *ExtensionSpec) GetLongDescription() string {
	if m != nil {
		return m.LongDescription
	}
	return ""
}

func (m *ExtensionSpec) GetDocumentationUrl() string {
	if m != nil {
		return m.DocumentationUrl
	}
	return ""
}

func (m *ExtensionSpec) GetRepositoryUrl() string {
	if m != nil {
		return m.RepositoryUrl
	}
	return ""
}

func (m *ExtensionSpec) GetExtensionRef() string {
	if m != nil {
		return m.ExtensionRef
	}
	return ""
}

func (m *ExtensionSpec) GetCreatorName() string {
	if m != nil {
		return m.CreatorName
	}
	return ""
}

func (m *ExtensionSpec) GetCreatorUrl() string {
	if m != nil {
		return m.CreatorUrl
	}
	return ""
}

func (m *ExtensionSpec) GetInstallInstructions() *InstallInstructions {
	if m != nil {
		return m.InstallInstructions
	}
	return nil
}

func (m *ExtensionSpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ExtensionSpec) GetComingSoon() bool {
	if m != nil {
		return m.ComingSoon
	}
	return false
}

// Set of install instructions for each supported envoy flavor
type InstallInstructions struct {
	Gloo                 *InstallInstructions_Gloo         `protobuf:"bytes,1,opt,name=gloo,proto3" json:"gloo,omitempty"`
	VanillaEnvoy         *InstallInstructions_VanillaEnvoy `protobuf:"bytes,3,opt,name=vanilla_envoy,json=vanillaEnvoy,proto3" json:"vanilla_envoy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *InstallInstructions) Reset()         { *m = InstallInstructions{} }
func (m *InstallInstructions) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions) ProtoMessage()    {}
func (*InstallInstructions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb76b5cc14a8263, []int{1}
}

func (m *InstallInstructions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions.Unmarshal(m, b)
}
func (m *InstallInstructions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions.Marshal(b, m, deterministic)
}
func (m *InstallInstructions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions.Merge(m, src)
}
func (m *InstallInstructions) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions.Size(m)
}
func (m *InstallInstructions) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions proto.InternalMessageInfo

func (m *InstallInstructions) GetGloo() *InstallInstructions_Gloo {
	if m != nil {
		return m.Gloo
	}
	return nil
}

func (m *InstallInstructions) GetVanillaEnvoy() *InstallInstructions_VanillaEnvoy {
	if m != nil {
		return m.VanillaEnvoy
	}
	return nil
}

type InstallInstructions_Gloo struct {
	Glooctl              string   `protobuf:"bytes,1,opt,name=glooctl,proto3" json:"glooctl,omitempty"`
	GatewayYaml          string   `protobuf:"bytes,2,opt,name=gateway_yaml,json=gatewayYaml,proto3" json:"gateway_yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallInstructions_Gloo) Reset()         { *m = InstallInstructions_Gloo{} }
func (m *InstallInstructions_Gloo) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions_Gloo) ProtoMessage()    {}
func (*InstallInstructions_Gloo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb76b5cc14a8263, []int{1, 0}
}

func (m *InstallInstructions_Gloo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions_Gloo.Unmarshal(m, b)
}
func (m *InstallInstructions_Gloo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions_Gloo.Marshal(b, m, deterministic)
}
func (m *InstallInstructions_Gloo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions_Gloo.Merge(m, src)
}
func (m *InstallInstructions_Gloo) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions_Gloo.Size(m)
}
func (m *InstallInstructions_Gloo) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions_Gloo.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions_Gloo proto.InternalMessageInfo

func (m *InstallInstructions_Gloo) GetGlooctl() string {
	if m != nil {
		return m.Glooctl
	}
	return ""
}

func (m *InstallInstructions_Gloo) GetGatewayYaml() string {
	if m != nil {
		return m.GatewayYaml
	}
	return ""
}

type InstallInstructions_VanillaEnvoy struct {
	Config               string   `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallInstructions_VanillaEnvoy) Reset()         { *m = InstallInstructions_VanillaEnvoy{} }
func (m *InstallInstructions_VanillaEnvoy) String() string { return proto.CompactTextString(m) }
func (*InstallInstructions_VanillaEnvoy) ProtoMessage()    {}
func (*InstallInstructions_VanillaEnvoy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb76b5cc14a8263, []int{1, 1}
}

func (m *InstallInstructions_VanillaEnvoy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Unmarshal(m, b)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Marshal(b, m, deterministic)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallInstructions_VanillaEnvoy.Merge(m, src)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_Size() int {
	return xxx_messageInfo_InstallInstructions_VanillaEnvoy.Size(m)
}
func (m *InstallInstructions_VanillaEnvoy) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallInstructions_VanillaEnvoy.DiscardUnknown(m)
}

var xxx_messageInfo_InstallInstructions_VanillaEnvoy proto.InternalMessageInfo

func (m *InstallInstructions_VanillaEnvoy) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtensionSpec)(nil), "ExtensionSpec")
	proto.RegisterType((*InstallInstructions)(nil), "InstallInstructions")
	proto.RegisterType((*InstallInstructions_Gloo)(nil), "InstallInstructions.Gloo")
	proto.RegisterType((*InstallInstructions_VanillaEnvoy)(nil), "InstallInstructions.VanillaEnvoy")
}

func init() { proto.RegisterFile("pkg/catalog/catalog.proto", fileDescriptor_2fb76b5cc14a8263) }

var fileDescriptor_2fb76b5cc14a8263 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x24, 0xe4, 0xcf, 0xec, 0x06, 0x5a, 0xb7, 0x42, 0x9b, 0x5e, 0x48, 0x8b, 0x40,
	0x41, 0x15, 0x41, 0x0a, 0x6f, 0x00, 0x94, 0xaa, 0x17, 0x0e, 0x5b, 0x81, 0x04, 0x97, 0x95, 0xd9,
	0x3a, 0xc6, 0xc2, 0xeb, 0x59, 0xd9, 0x4e, 0x20, 0x8f, 0xc9, 0xe3, 0x70, 0x43, 0x33, 0xd9, 0xb4,
	0xa9, 0x9a, 0x93, 0x3d, 0xdf, 0xfc, 0x76, 0x3e, 0xef, 0xcc, 0xc0, 0xb8, 0xf9, 0xa5, 0xdf, 0x56,
	0x32, 0x4a, 0x8b, 0xb7, 0xe7, 0xac, 0xf1, 0x18, 0xf1, 0xec, 0x6f, 0x07, 0x46, 0x17, 0x7f, 0xa2,
	0x72, 0xc1, 0xa0, 0xbb, 0x6e, 0x54, 0x25, 0x04, 0x74, 0x9d, 0xac, 0x55, 0x9e, 0x4c, 0x92, 0xe9,
	0xb0, 0xe0, 0xbb, 0x18, 0xc3, 0xc0, 0xa2, 0xc6, 0x72, 0xe9, 0x6d, 0xfe, 0x88, 0xf5, 0x3e, 0xc5,
	0x5f, 0xbc, 0x15, 0xe7, 0x70, 0x18, 0x7e, 0xa2, 0x8f, 0xe5, 0x8d, 0x0a, 0x95, 0x37, 0x4d, 0x34,
	0xe8, 0xf2, 0x0e, 0x33, 0x07, 0x9c, 0xf8, 0x78, 0xa7, 0x8b, 0xd7, 0x70, 0x60, 0xd1, 0xe9, 0x7b,
	0x6c, 0x97, 0xd9, 0xa7, 0xa4, 0xef, 0xa2, 0xe7, 0x70, 0x78, 0x83, 0xd5, 0xb2, 0x56, 0x2e, 0x4a,
	0x12, 0xd8, 0xfb, 0xf1, 0xa6, 0xee, 0xbd, 0x04, 0x3d, 0xe2, 0x25, 0x3c, 0xf1, 0xaa, 0xc1, 0x60,
	0x22, 0xfa, 0x35, 0x93, 0x3d, 0x26, 0x47, 0x77, 0x2a, 0x61, 0x2f, 0x60, 0xa4, 0xb6, 0xff, 0x5a,
	0x7a, 0xb5, 0xc8, 0xfb, 0x4c, 0x65, 0xb7, 0x62, 0xa1, 0x16, 0xe2, 0x14, 0xb2, 0xca, 0x2b, 0x19,
	0xd1, 0x97, 0xdc, 0x87, 0x01, 0x33, 0x69, 0xab, 0x7d, 0xa6, 0x76, 0x3c, 0x87, 0x6d, 0xc8, 0x5e,
	0x43, 0x26, 0xa0, 0x95, 0xc8, 0xe8, 0x12, 0x8e, 0x8d, 0x0b, 0x51, 0x5a, 0x5b, 0xd2, 0xe9, 0x97,
	0x15, 0xbd, 0x34, 0xe4, 0x30, 0x49, 0xa6, 0xe9, 0xfc, 0x78, 0x76, 0xb5, 0x49, 0x5e, 0xed, 0xe4,
	0x8a, 0x23, 0xf3, 0x50, 0xa4, 0x61, 0x44, 0xa9, 0x43, 0x9e, 0x4e, 0x3a, 0x34, 0x0c, 0xba, 0xb3,
	0x3b, 0xd6, 0xc6, 0xe9, 0x32, 0x20, 0xba, 0x3c, 0x9b, 0x24, 0xd3, 0x41, 0x01, 0x1b, 0xe9, 0x1a,
	0xd1, 0x9d, 0xfd, 0x4b, 0xe0, 0x68, 0x8f, 0x83, 0x78, 0x03, 0x5d, 0x6d, 0x11, 0x79, 0xb2, 0xe9,
	0x7c, 0xbc, 0xef, 0x15, 0xb3, 0x4b, 0x8b, 0x58, 0x30, 0x26, 0x3e, 0xc1, 0x68, 0x25, 0x9d, 0xb1,
	0x56, 0x96, 0xca, 0xad, 0x70, 0xcd, 0x53, 0x4d, 0xe7, 0xa7, 0x7b, 0xbf, 0xfb, 0xba, 0x21, 0x2f,
	0x08, 0x2c, 0xb2, 0xd5, 0x4e, 0x74, 0xf2, 0x01, 0xba, 0x54, 0x55, 0xe4, 0xd0, 0xa7, 0xba, 0x55,
	0xb4, 0xed, 0x6e, 0x6d, 0x43, 0x6a, 0xb9, 0x96, 0x51, 0xfd, 0x96, 0xeb, 0x72, 0x2d, 0xeb, 0xed,
	0x8a, 0xa5, 0xad, 0xf6, 0x4d, 0xd6, 0xf6, 0xe4, 0x15, 0x64, 0xbb, 0x16, 0xe2, 0x19, 0xf4, 0x2a,
	0x74, 0x0b, 0xa3, 0xdb, 0x5a, 0x6d, 0xf4, 0x7e, 0xf8, 0xbd, 0xdf, 0x2e, 0xf8, 0x8f, 0x1e, 0x6f,
	0xf8, 0xbb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x55, 0x0b, 0x59, 0xfe, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: config.proto

package api

import (
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

type Kind int32

const (
	Kind_UNKNOWN     Kind = 0
	Kind_HELM        Kind = 1
	Kind_CHARTMUSEUM Kind = 2
	Kind_HARBOR      Kind = 3
	Kind_OCI         Kind = 4
	Kind_LOCAL       Kind = 5
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "HELM",
		2: "CHARTMUSEUM",
		3: "HARBOR",
		4: "OCI",
		5: "LOCAL",
	}
	Kind_value = map[string]int32{
		"UNKNOWN":     0,
		"HELM":        1,
		"CHARTMUSEUM": 2,
		"HARBOR":      3,
		"OCI":         4,
		"LOCAL":       5,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

// Config file structure
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target *Target `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// Helm Charts to include during sync
	Charts []string `protobuf:"bytes,3,rep,name=charts,proto3" json:"charts,omitempty"`
	// Opposite of charts property. It indicates the list of charts to skip during sync
	SkipCharts              []string `protobuf:"bytes,5,rep,name=skip_charts,json=skipCharts,proto3" json:"skip_charts,omitempty"`
	RelocateContainerImages bool     `protobuf:"varint,4,opt,name=relocate_container_images,json=relocateContainerImages,proto3" json:"relocate_container_images,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Config) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Config) GetCharts() []string {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *Config) GetSkipCharts() []string {
	if x != nil {
		return x.SkipCharts
	}
	return nil
}

func (x *Config) GetRelocateContainerImages() bool {
	if x != nil {
		return x.RelocateContainerImages
	}
	return false
}

// SourceRepo contains the required information of the source chart repository
type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Spec:
	//
	//	*Source_Repo
	//	*Source_IntermediateBundlesPath
	Spec isSource_Spec `protobuf_oneof:"spec"`
	// Ignored if the repo is an intermediate bundle since the images are inside the bundle
	Containers        *Containers `protobuf:"bytes,3,opt,name=containers,proto3" json:"containers,omitempty"`
	TrustedSourceDeps []*Repo     `protobuf:"bytes,6,rep,name=trustedSourceDeps,proto3" json:"trustedSourceDeps,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (m *Source) GetSpec() isSource_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *Source) GetRepo() *Repo {
	if x, ok := x.GetSpec().(*Source_Repo); ok {
		return x.Repo
	}
	return nil
}

func (x *Source) GetIntermediateBundlesPath() string {
	if x, ok := x.GetSpec().(*Source_IntermediateBundlesPath); ok {
		return x.IntermediateBundlesPath
	}
	return ""
}

func (x *Source) GetContainers() *Containers {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Source) GetTrustedSourceDeps() []*Repo {
	if x != nil {
		return x.TrustedSourceDeps
	}
	return nil
}

type isSource_Spec interface {
	isSource_Spec()
}

type Source_Repo struct {
	Repo *Repo `protobuf:"bytes,1,opt,name=repo,proto3,oneof"`
}

type Source_IntermediateBundlesPath struct {
	IntermediateBundlesPath string `protobuf:"bytes,2,opt,name=intermediate_bundles_path,json=intermediateBundlesPath,proto3,oneof"`
}

func (*Source_Repo) isSource_Spec() {}

func (*Source_IntermediateBundlesPath) isSource_Spec() {}

type Containers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *Containers_ContainerAuth `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Containers) Reset() {
	*x = Containers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Containers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Containers) ProtoMessage() {}

func (x *Containers) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Containers.ProtoReflect.Descriptor instead.
func (*Containers) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *Containers) GetAuth() *Containers_ContainerAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// TargetRepo contains the required information of the target chart repository
type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Spec:
	//
	//	*Target_Repo
	//	*Target_IntermediateBundlesPath
	Spec                  isTarget_Spec `protobuf_oneof:"spec"`
	ContainerRegistry     string        `protobuf:"bytes,2,opt,name=container_registry,json=containerRegistry,proto3" json:"container_registry,omitempty"`
	ContainerRepository   string        `protobuf:"bytes,3,opt,name=container_repository,json=containerRepository,proto3" json:"container_repository,omitempty"`
	RepoName              string        `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	Containers            *Containers   `protobuf:"bytes,6,opt,name=containers,proto3" json:"containers,omitempty"`
	ReplaceDependencyRepo bool          `protobuf:"varint,7,opt,name=replaceDependencyRepo,proto3" json:"replaceDependencyRepo,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (m *Target) GetSpec() isTarget_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *Target) GetRepo() *Repo {
	if x, ok := x.GetSpec().(*Target_Repo); ok {
		return x.Repo
	}
	return nil
}

func (x *Target) GetIntermediateBundlesPath() string {
	if x, ok := x.GetSpec().(*Target_IntermediateBundlesPath); ok {
		return x.IntermediateBundlesPath
	}
	return ""
}

func (x *Target) GetContainerRegistry() string {
	if x != nil {
		return x.ContainerRegistry
	}
	return ""
}

func (x *Target) GetContainerRepository() string {
	if x != nil {
		return x.ContainerRepository
	}
	return ""
}

func (x *Target) GetRepoName() string {
	if x != nil {
		return x.RepoName
	}
	return ""
}

func (x *Target) GetContainers() *Containers {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *Target) GetReplaceDependencyRepo() bool {
	if x != nil {
		return x.ReplaceDependencyRepo
	}
	return false
}

type isTarget_Spec interface {
	isTarget_Spec()
}

type Target_Repo struct {
	Repo *Repo `protobuf:"bytes,1,opt,name=repo,proto3,oneof"`
}

type Target_IntermediateBundlesPath struct {
	IntermediateBundlesPath string `protobuf:"bytes,5,opt,name=intermediate_bundles_path,json=intermediateBundlesPath,proto3,oneof"`
}

func (*Target_Repo) isTarget_Spec() {}

func (*Target_IntermediateBundlesPath) isTarget_Spec() {}

// Generic repo representation
type Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Kind Kind   `protobuf:"varint,2,opt,name=kind,proto3,enum=api.Kind" json:"kind,omitempty"`
	Auth *Auth  `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	// The path where the repo stores charts. Useful for LOCAL kind only
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// The OCI reference where the index of charts is located
	// Example: my.oci.domain/index:latest
	ChartsIndex string `protobuf:"bytes,5,opt,name=charts_index,json=chartsIndex,proto3" json:"charts_index,omitempty"`
	// Whether to use a charts index to find charts
	//
	// Deprecated: Do not use.
	UseChartsIndex     bool `protobuf:"varint,6,opt,name=use_charts_index,json=useChartsIndex,proto3" json:"use_charts_index,omitempty"`
	DisableChartsIndex bool `protobuf:"varint,7,opt,name=disable_charts_index,json=disableChartsIndex,proto3" json:"disable_charts_index,omitempty"`
}

func (x *Repo) Reset() {
	*x = Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repo) ProtoMessage() {}

func (x *Repo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repo.ProtoReflect.Descriptor instead.
func (*Repo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *Repo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Repo) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_UNKNOWN
}

func (x *Repo) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Repo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Repo) GetChartsIndex() string {
	if x != nil {
		return x.ChartsIndex
	}
	return ""
}

// Deprecated: Do not use.
func (x *Repo) GetUseChartsIndex() bool {
	if x != nil {
		return x.UseChartsIndex
	}
	return false
}

func (x *Repo) GetDisableChartsIndex() bool {
	if x != nil {
		return x.DisableChartsIndex
	}
	return false
}

// Auth contains credentials to login to a chart repository
type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *Auth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Auth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// ContainerAuth defines the authentication parameters required to access the source/target
// OCI registries during container image relocation
type Containers_ContainerAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Registry string `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *Containers_ContainerAuth) Reset() {
	*x = Containers_ContainerAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Containers_ContainerAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Containers_ContainerAuth) ProtoMessage() {}

func (x *Containers_ContainerAuth) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Containers_ContainerAuth.ProtoReflect.Descriptor instead.
func (*Containers_ContainerAuth) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Containers_ContainerAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Containers_ContainerAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Containers_ContainerAuth) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0xc7, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x73, 0x12, 0x3a, 0x0a, 0x19, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x72, 0x65, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0xd9, 0x01,
	0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x48, 0x00, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x3c, 0x0a, 0x19, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x17,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x11, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x70, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x11,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x70,
	0x73, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x63, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x22, 0xd5, 0x02, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x3c, 0x0a, 0x19,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x42, 0x06, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xed, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x70,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3e, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2a, 0x4e, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x48, 0x45, 0x4c, 0x4d, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x52, 0x54,
	0x4d, 0x55, 0x53, 0x45, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x41, 0x52, 0x42,
	0x4f, 0x52, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x43, 0x49, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x69, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_proto_goTypes = []interface{}{
	(Kind)(0),                        // 0: api.Kind
	(*Config)(nil),                   // 1: api.Config
	(*Source)(nil),                   // 2: api.Source
	(*Containers)(nil),               // 3: api.Containers
	(*Target)(nil),                   // 4: api.Target
	(*Repo)(nil),                     // 5: api.Repo
	(*Auth)(nil),                     // 6: api.Auth
	(*Containers_ContainerAuth)(nil), // 7: api.Containers.ContainerAuth
}
var file_config_proto_depIdxs = []int32{
	2,  // 0: api.Config.source:type_name -> api.Source
	4,  // 1: api.Config.target:type_name -> api.Target
	5,  // 2: api.Source.repo:type_name -> api.Repo
	3,  // 3: api.Source.containers:type_name -> api.Containers
	5,  // 4: api.Source.trustedSourceDeps:type_name -> api.Repo
	7,  // 5: api.Containers.auth:type_name -> api.Containers.ContainerAuth
	5,  // 6: api.Target.repo:type_name -> api.Repo
	3,  // 7: api.Target.containers:type_name -> api.Containers
	0,  // 8: api.Repo.kind:type_name -> api.Kind
	6,  // 9: api.Repo.auth:type_name -> api.Auth
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Containers); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repo); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Containers_ContainerAuth); i {
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
	file_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Source_Repo)(nil),
		(*Source_IntermediateBundlesPath)(nil),
	}
	file_config_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Target_Repo)(nil),
		(*Target_IntermediateBundlesPath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}

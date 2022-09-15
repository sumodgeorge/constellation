// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: init.proto

package initproto

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

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated string autoscaling_node_groups = 1; removed
	MasterSecret           []byte        `protobuf:"bytes,2,opt,name=master_secret,json=masterSecret,proto3" json:"master_secret,omitempty"`
	KmsUri                 string        `protobuf:"bytes,3,opt,name=kms_uri,json=kmsUri,proto3" json:"kms_uri,omitempty"`
	StorageUri             string        `protobuf:"bytes,4,opt,name=storage_uri,json=storageUri,proto3" json:"storage_uri,omitempty"`
	KeyEncryptionKeyId     string        `protobuf:"bytes,5,opt,name=key_encryption_key_id,json=keyEncryptionKeyId,proto3" json:"key_encryption_key_id,omitempty"`
	UseExistingKek         bool          `protobuf:"varint,6,opt,name=use_existing_kek,json=useExistingKek,proto3" json:"use_existing_kek,omitempty"`
	CloudServiceAccountUri string        `protobuf:"bytes,7,opt,name=cloud_service_account_uri,json=cloudServiceAccountUri,proto3" json:"cloud_service_account_uri,omitempty"`
	KubernetesVersion      string        `protobuf:"bytes,8,opt,name=kubernetes_version,json=kubernetesVersion,proto3" json:"kubernetes_version,omitempty"`
	SshUserKeys            []*SSHUserKey `protobuf:"bytes,9,rep,name=ssh_user_keys,json=sshUserKeys,proto3" json:"ssh_user_keys,omitempty"`
	Salt                   []byte        `protobuf:"bytes,10,opt,name=salt,proto3" json:"salt,omitempty"`
	HelmDeployments        []byte        `protobuf:"bytes,11,opt,name=helm_deployments,json=helmDeployments,proto3" json:"helm_deployments,omitempty"`
	EnforcedPcrs           []uint32      `protobuf:"varint,12,rep,packed,name=enforced_pcrs,json=enforcedPcrs,proto3" json:"enforced_pcrs,omitempty"`
	EnforceIdkeydigest     bool          `protobuf:"varint,13,opt,name=enforce_idkeydigest,json=enforceIdkeydigest,proto3" json:"enforce_idkeydigest,omitempty"`
	ConformanceMode        bool          `protobuf:"varint,14,opt,name=conformance_mode,json=conformanceMode,proto3" json:"conformance_mode,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetMasterSecret() []byte {
	if x != nil {
		return x.MasterSecret
	}
	return nil
}

func (x *InitRequest) GetKmsUri() string {
	if x != nil {
		return x.KmsUri
	}
	return ""
}

func (x *InitRequest) GetStorageUri() string {
	if x != nil {
		return x.StorageUri
	}
	return ""
}

func (x *InitRequest) GetKeyEncryptionKeyId() string {
	if x != nil {
		return x.KeyEncryptionKeyId
	}
	return ""
}

func (x *InitRequest) GetUseExistingKek() bool {
	if x != nil {
		return x.UseExistingKek
	}
	return false
}

func (x *InitRequest) GetCloudServiceAccountUri() string {
	if x != nil {
		return x.CloudServiceAccountUri
	}
	return ""
}

func (x *InitRequest) GetKubernetesVersion() string {
	if x != nil {
		return x.KubernetesVersion
	}
	return ""
}

func (x *InitRequest) GetSshUserKeys() []*SSHUserKey {
	if x != nil {
		return x.SshUserKeys
	}
	return nil
}

func (x *InitRequest) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *InitRequest) GetHelmDeployments() []byte {
	if x != nil {
		return x.HelmDeployments
	}
	return nil
}

func (x *InitRequest) GetEnforcedPcrs() []uint32 {
	if x != nil {
		return x.EnforcedPcrs
	}
	return nil
}

func (x *InitRequest) GetEnforceIdkeydigest() bool {
	if x != nil {
		return x.EnforceIdkeydigest
	}
	return false
}

func (x *InitRequest) GetConformanceMode() bool {
	if x != nil {
		return x.ConformanceMode
	}
	return false
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubeconfig []byte `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	OwnerId    []byte `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	ClusterId  []byte `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetKubeconfig() []byte {
	if x != nil {
		return x.Kubeconfig
	}
	return nil
}

func (x *InitResponse) GetOwnerId() []byte {
	if x != nil {
		return x.OwnerId
	}
	return nil
}

func (x *InitResponse) GetClusterId() []byte {
	if x != nil {
		return x.ClusterId
	}
	return nil
}

type SSHUserKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SSHUserKey) Reset() {
	*x = SSHUserKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_init_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHUserKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHUserKey) ProtoMessage() {}

func (x *SSHUserKey) ProtoReflect() protoreflect.Message {
	mi := &file_init_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHUserKey.ProtoReflect.Descriptor instead.
func (*SSHUserKey) Descriptor() ([]byte, []int) {
	return file_init_proto_rawDescGZIP(), []int{2}
}

func (x *SSHUserKey) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SSHUserKey) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_init_proto protoreflect.FileDescriptor

var file_init_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e,
	0x69, 0x74, 0x22, 0xa9, 0x04, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x6d, 0x73, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x6d, 0x73, 0x55, 0x72, 0x69,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x69, 0x12, 0x31, 0x0a, 0x15, 0x6b, 0x65, 0x79, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x6b, 0x65, 0x79, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x75, 0x73, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x6b, 0x12, 0x39,
	0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0d, 0x73, 0x73, 0x68, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x53, 0x53, 0x48, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x52, 0x0b, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61,
	0x6c, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x68, 0x65,
	0x6c, 0x6d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x5f, 0x70, 0x63, 0x72, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x50, 0x63,
	0x72, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x6b, 0x65, 0x79, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x49, 0x64, 0x6b, 0x65, 0x79, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x68,
	0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0a, 0x53, 0x53, 0x48, 0x55,
	0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x32, 0x34, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x11, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x73, 0x79,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x69,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_init_proto_rawDescOnce sync.Once
	file_init_proto_rawDescData = file_init_proto_rawDesc
)

func file_init_proto_rawDescGZIP() []byte {
	file_init_proto_rawDescOnce.Do(func() {
		file_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_init_proto_rawDescData)
	})
	return file_init_proto_rawDescData
}

var file_init_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_init_proto_goTypes = []interface{}{
	(*InitRequest)(nil),  // 0: init.InitRequest
	(*InitResponse)(nil), // 1: init.InitResponse
	(*SSHUserKey)(nil),   // 2: init.SSHUserKey
}
var file_init_proto_depIdxs = []int32{
	2, // 0: init.InitRequest.ssh_user_keys:type_name -> init.SSHUserKey
	0, // 1: init.API.Init:input_type -> init.InitRequest
	1, // 2: init.API.Init:output_type -> init.InitResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_init_proto_init() }
func file_init_proto_init() {
	if File_init_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_init_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHUserKey); i {
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
			RawDescriptor: file_init_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_init_proto_goTypes,
		DependencyIndexes: file_init_proto_depIdxs,
		MessageInfos:      file_init_proto_msgTypes,
	}.Build()
	File_init_proto = out.File
	file_init_proto_rawDesc = nil
	file_init_proto_goTypes = nil
	file_init_proto_depIdxs = nil
}

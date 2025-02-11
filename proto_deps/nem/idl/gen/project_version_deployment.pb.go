// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: project_version_deployment.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProjectVersionDeploymentStatus int32

const (
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_INVALID     ProjectVersionDeploymentStatus = 0
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_NOT_STARTED ProjectVersionDeploymentStatus = 1
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_STARTED     ProjectVersionDeploymentStatus = 2
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_COMPLETED   ProjectVersionDeploymentStatus = 3
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_FAILED      ProjectVersionDeploymentStatus = 4
	ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_ROLLED_BACK ProjectVersionDeploymentStatus = 5
)

// Enum value maps for ProjectVersionDeploymentStatus.
var (
	ProjectVersionDeploymentStatus_name = map[int32]string{
		0: "PROJECT_VERSION_DEPLOYMENT_STATUS_INVALID",
		1: "PROJECT_VERSION_DEPLOYMENT_STATUS_NOT_STARTED",
		2: "PROJECT_VERSION_DEPLOYMENT_STATUS_STARTED",
		3: "PROJECT_VERSION_DEPLOYMENT_STATUS_COMPLETED",
		4: "PROJECT_VERSION_DEPLOYMENT_STATUS_FAILED",
		5: "PROJECT_VERSION_DEPLOYMENT_STATUS_ROLLED_BACK",
	}
	ProjectVersionDeploymentStatus_value = map[string]int32{
		"PROJECT_VERSION_DEPLOYMENT_STATUS_INVALID":     0,
		"PROJECT_VERSION_DEPLOYMENT_STATUS_NOT_STARTED": 1,
		"PROJECT_VERSION_DEPLOYMENT_STATUS_STARTED":     2,
		"PROJECT_VERSION_DEPLOYMENT_STATUS_COMPLETED":   3,
		"PROJECT_VERSION_DEPLOYMENT_STATUS_FAILED":      4,
		"PROJECT_VERSION_DEPLOYMENT_STATUS_ROLLED_BACK": 5,
	}
)

func (x ProjectVersionDeploymentStatus) Enum() *ProjectVersionDeploymentStatus {
	p := new(ProjectVersionDeploymentStatus)
	*p = x
	return p
}

func (x ProjectVersionDeploymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectVersionDeploymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_version_deployment_proto_enumTypes[0].Descriptor()
}

func (ProjectVersionDeploymentStatus) Type() protoreflect.EnumType {
	return &file_project_version_deployment_proto_enumTypes[0]
}

func (x ProjectVersionDeploymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectVersionDeploymentStatus.Descriptor instead.
func (ProjectVersionDeploymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_version_deployment_proto_rawDescGZIP(), []int{0}
}

type ProjectVersionDeployment struct {
	state          protoimpl.MessageState         `protogen:"open.v1"`
	EnviromentUuid string                         `protobuf:"bytes,1,opt,name=enviroment_uuid,json=enviromentUuid,proto3" json:"enviroment_uuid,omitempty"`
	Status         ProjectVersionDeploymentStatus `protobuf:"varint,2,opt,name=status,proto3,enum=nem.ProjectVersionDeploymentStatus" json:"status,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ProjectVersionDeployment) Reset() {
	*x = ProjectVersionDeployment{}
	mi := &file_project_version_deployment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectVersionDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectVersionDeployment) ProtoMessage() {}

func (x *ProjectVersionDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_project_version_deployment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectVersionDeployment.ProtoReflect.Descriptor instead.
func (*ProjectVersionDeployment) Descriptor() ([]byte, []int) {
	return file_project_version_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectVersionDeployment) GetEnviromentUuid() string {
	if x != nil {
		return x.EnviromentUuid
	}
	return ""
}

func (x *ProjectVersionDeployment) GetStatus() ProjectVersionDeploymentStatus {
	if x != nil {
		return x.Status
	}
	return ProjectVersionDeploymentStatus_PROJECT_VERSION_DEPLOYMENT_STATUS_INVALID
}

var File_project_version_deployment_proto protoreflect.FileDescriptor

var file_project_version_deployment_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0x80, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x3b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e,
	0x6e, 0x65, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0xc3, 0x02, 0x0a, 0x1e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a,
	0x29, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x31, 0x0a, 0x2d,
	0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x2d, 0x0a, 0x29, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x2f,
	0x0a, 0x2b, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x2c, 0x0a, 0x28, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x31, 0x0a,
	0x2d, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x05,
	0x42, 0x2e, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_project_version_deployment_proto_rawDescOnce sync.Once
	file_project_version_deployment_proto_rawDescData []byte
)

func file_project_version_deployment_proto_rawDescGZIP() []byte {
	file_project_version_deployment_proto_rawDescOnce.Do(func() {
		file_project_version_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_version_deployment_proto_rawDesc), len(file_project_version_deployment_proto_rawDesc)))
	})
	return file_project_version_deployment_proto_rawDescData
}

var file_project_version_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_version_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_version_deployment_proto_goTypes = []any{
	(ProjectVersionDeploymentStatus)(0), // 0: nem.ProjectVersionDeploymentStatus
	(*ProjectVersionDeployment)(nil),    // 1: nem.ProjectVersionDeployment
}
var file_project_version_deployment_proto_depIdxs = []int32{
	0, // 0: nem.ProjectVersionDeployment.status:type_name -> nem.ProjectVersionDeploymentStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_project_version_deployment_proto_init() }
func file_project_version_deployment_proto_init() {
	if File_project_version_deployment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_version_deployment_proto_rawDesc), len(file_project_version_deployment_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_version_deployment_proto_goTypes,
		DependencyIndexes: file_project_version_deployment_proto_depIdxs,
		EnumInfos:         file_project_version_deployment_proto_enumTypes,
		MessageInfos:      file_project_version_deployment_proto_msgTypes,
	}.Build()
	File_project_version_deployment_proto = out.File
	file_project_version_deployment_proto_goTypes = nil
	file_project_version_deployment_proto_depIdxs = nil
}

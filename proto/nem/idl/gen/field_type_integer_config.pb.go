// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: field_type_integer_config.proto

package gen

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

type FieldTypeIntegerConfigSize int32

const (
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_INVALID          FieldTypeIntegerConfigSize = 0
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_ONE_BIT          FieldTypeIntegerConfigSize = 1
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_EIGHT_BITS       FieldTypeIntegerConfigSize = 2
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTEEN_BITS     FieldTypeIntegerConfigSize = 3
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_TWENTY_FOUR_BITS FieldTypeIntegerConfigSize = 4
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_THIRTY_TWO_BITS  FieldTypeIntegerConfigSize = 5
	FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTY_FOUR_BITS  FieldTypeIntegerConfigSize = 6
)

// Enum value maps for FieldTypeIntegerConfigSize.
var (
	FieldTypeIntegerConfigSize_name = map[int32]string{
		0: "FIELD_TYPE_INTEGER_CONFIG_SIZE_INVALID",
		1: "FIELD_TYPE_INTEGER_CONFIG_SIZE_ONE_BIT",
		2: "FIELD_TYPE_INTEGER_CONFIG_SIZE_EIGHT_BITS",
		3: "FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTEEN_BITS",
		4: "FIELD_TYPE_INTEGER_CONFIG_SIZE_TWENTY_FOUR_BITS",
		5: "FIELD_TYPE_INTEGER_CONFIG_SIZE_THIRTY_TWO_BITS",
		6: "FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTY_FOUR_BITS",
	}
	FieldTypeIntegerConfigSize_value = map[string]int32{
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_INVALID":          0,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_ONE_BIT":          1,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_EIGHT_BITS":       2,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTEEN_BITS":     3,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_TWENTY_FOUR_BITS": 4,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_THIRTY_TWO_BITS":  5,
		"FIELD_TYPE_INTEGER_CONFIG_SIZE_SIXTY_FOUR_BITS":  6,
	}
)

func (x FieldTypeIntegerConfigSize) Enum() *FieldTypeIntegerConfigSize {
	p := new(FieldTypeIntegerConfigSize)
	*p = x
	return p
}

func (x FieldTypeIntegerConfigSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldTypeIntegerConfigSize) Descriptor() protoreflect.EnumDescriptor {
	return file_field_type_integer_config_proto_enumTypes[0].Descriptor()
}

func (FieldTypeIntegerConfigSize) Type() protoreflect.EnumType {
	return &file_field_type_integer_config_proto_enumTypes[0]
}

func (x FieldTypeIntegerConfigSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldTypeIntegerConfigSize.Descriptor instead.
func (FieldTypeIntegerConfigSize) EnumDescriptor() ([]byte, []int) {
	return file_field_type_integer_config_proto_rawDescGZIP(), []int{0}
}

type FieldTypeIntegerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size              FieldTypeIntegerConfigSize `protobuf:"varint,1,opt,name=size,proto3,enum=nem.FieldTypeIntegerConfigSize" json:"size,omitempty"`
	MinValue          int64                      `protobuf:"varint,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	MinValueInclusive bool                       `protobuf:"varint,3,opt,name=min_value_inclusive,json=minValueInclusive,proto3" json:"min_value_inclusive,omitempty"`
	MaxValue          int64                      `protobuf:"varint,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MaxValueInclusive bool                       `protobuf:"varint,5,opt,name=max_value_inclusive,json=maxValueInclusive,proto3" json:"max_value_inclusive,omitempty"`
	AllowNegatives    bool                       `protobuf:"varint,6,opt,name=allow_negatives,json=allowNegatives,proto3" json:"allow_negatives,omitempty"`
	EnableLimits      bool                       `protobuf:"varint,7,opt,name=enable_limits,json=enableLimits,proto3" json:"enable_limits,omitempty"`
}

func (x *FieldTypeIntegerConfig) Reset() {
	*x = FieldTypeIntegerConfig{}
	mi := &file_field_type_integer_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeIntegerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeIntegerConfig) ProtoMessage() {}

func (x *FieldTypeIntegerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_integer_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeIntegerConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeIntegerConfig) Descriptor() ([]byte, []int) {
	return file_field_type_integer_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeIntegerConfig) GetSize() FieldTypeIntegerConfigSize {
	if x != nil {
		return x.Size
	}
	return FieldTypeIntegerConfigSize_FIELD_TYPE_INTEGER_CONFIG_SIZE_INVALID
}

func (x *FieldTypeIntegerConfig) GetMinValue() int64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *FieldTypeIntegerConfig) GetMinValueInclusive() bool {
	if x != nil {
		return x.MinValueInclusive
	}
	return false
}

func (x *FieldTypeIntegerConfig) GetMaxValue() int64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *FieldTypeIntegerConfig) GetMaxValueInclusive() bool {
	if x != nil {
		return x.MaxValueInclusive
	}
	return false
}

func (x *FieldTypeIntegerConfig) GetAllowNegatives() bool {
	if x != nil {
		return x.AllowNegatives
	}
	return false
}

func (x *FieldTypeIntegerConfig) GetEnableLimits() bool {
	if x != nil {
		return x.EnableLimits
	}
	return false
}

var File_field_type_integer_config_proto protoreflect.FileDescriptor

var file_field_type_integer_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x22, 0xb5, 0x02, 0x0a, 0x16, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x69, 0x7a, 0x65,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6d,
	0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x2a, 0xf1,
	0x02, 0x0a, 0x1a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a,
	0x26, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4f, 0x4e, 0x45, 0x5f,
	0x42, 0x49, 0x54, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x42, 0x49,
	0x54, 0x53, 0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x53, 0x49, 0x58, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x42,
	0x49, 0x54, 0x53, 0x10, 0x03, 0x12, 0x33, 0x0a, 0x2f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x54, 0x57, 0x45, 0x4e, 0x54, 0x59, 0x5f, 0x46,
	0x4f, 0x55, 0x52, 0x5f, 0x42, 0x49, 0x54, 0x53, 0x10, 0x04, 0x12, 0x32, 0x0a, 0x2e, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x54, 0x48, 0x49,
	0x52, 0x54, 0x59, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x42, 0x49, 0x54, 0x53, 0x10, 0x05, 0x12, 0x32,
	0x0a, 0x2e, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54,
	0x45, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x49, 0x5a, 0x45,
	0x5f, 0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x5f, 0x42, 0x49, 0x54, 0x53,
	0x10, 0x06, 0x42, 0x2c, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x16, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_field_type_integer_config_proto_rawDescOnce sync.Once
	file_field_type_integer_config_proto_rawDescData = file_field_type_integer_config_proto_rawDesc
)

func file_field_type_integer_config_proto_rawDescGZIP() []byte {
	file_field_type_integer_config_proto_rawDescOnce.Do(func() {
		file_field_type_integer_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_type_integer_config_proto_rawDescData)
	})
	return file_field_type_integer_config_proto_rawDescData
}

var file_field_type_integer_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_type_integer_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_integer_config_proto_goTypes = []any{
	(FieldTypeIntegerConfigSize)(0), // 0: nem.FieldTypeIntegerConfigSize
	(*FieldTypeIntegerConfig)(nil),  // 1: nem.FieldTypeIntegerConfig
}
var file_field_type_integer_config_proto_depIdxs = []int32{
	0, // 0: nem.FieldTypeIntegerConfig.size:type_name -> nem.FieldTypeIntegerConfigSize
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_field_type_integer_config_proto_init() }
func file_field_type_integer_config_proto_init() {
	if File_field_type_integer_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_type_integer_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_integer_config_proto_goTypes,
		DependencyIndexes: file_field_type_integer_config_proto_depIdxs,
		EnumInfos:         file_field_type_integer_config_proto_enumTypes,
		MessageInfos:      file_field_type_integer_config_proto_msgTypes,
	}.Build()
	File_field_type_integer_config_proto = out.File
	file_field_type_integer_config_proto_rawDesc = nil
	file_field_type_integer_config_proto_goTypes = nil
	file_field_type_integer_config_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: field_type_float_config.proto

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

type FieldTypeFloatConfigSeparator int32

const (
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID FieldTypeFloatConfigSeparator = 0
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT   FieldTypeFloatConfigSeparator = 1
	FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA   FieldTypeFloatConfigSeparator = 2
)

// Enum value maps for FieldTypeFloatConfigSeparator.
var (
	FieldTypeFloatConfigSeparator_name = map[int32]string{
		0: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID",
		1: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT",
		2: "FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA",
	}
	FieldTypeFloatConfigSeparator_value = map[string]int32{
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID": 0,
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_POINT":   1,
		"FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_COMMA":   2,
	}
)

func (x FieldTypeFloatConfigSeparator) Enum() *FieldTypeFloatConfigSeparator {
	p := new(FieldTypeFloatConfigSeparator)
	*p = x
	return p
}

func (x FieldTypeFloatConfigSeparator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldTypeFloatConfigSeparator) Descriptor() protoreflect.EnumDescriptor {
	return file_field_type_float_config_proto_enumTypes[0].Descriptor()
}

func (FieldTypeFloatConfigSeparator) Type() protoreflect.EnumType {
	return &file_field_type_float_config_proto_enumTypes[0]
}

func (x FieldTypeFloatConfigSeparator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldTypeFloatConfigSeparator.Descriptor instead.
func (FieldTypeFloatConfigSeparator) EnumDescriptor() ([]byte, []int) {
	return file_field_type_float_config_proto_rawDescGZIP(), []int{0}
}

type FieldTypeFloatConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinValue          float64                       `protobuf:"fixed64,1,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	MinValueInclusive bool                          `protobuf:"varint,2,opt,name=min_value_inclusive,json=minValueInclusive,proto3" json:"min_value_inclusive,omitempty"`
	MaxValue          float64                       `protobuf:"fixed64,3,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MaxValueInclusive bool                          `protobuf:"varint,4,opt,name=max_value_inclusive,json=maxValueInclusive,proto3" json:"max_value_inclusive,omitempty"`
	AllowNegatives    bool                          `protobuf:"varint,5,opt,name=allow_negatives,json=allowNegatives,proto3" json:"allow_negatives,omitempty"`
	NumberOfDecimals  int64                         `protobuf:"varint,6,opt,name=number_of_decimals,json=numberOfDecimals,proto3" json:"number_of_decimals,omitempty"`
	Separator         FieldTypeFloatConfigSeparator `protobuf:"varint,7,opt,name=separator,proto3,enum=nem.FieldTypeFloatConfigSeparator" json:"separator,omitempty"`
	EnableLimits      bool                          `protobuf:"varint,8,opt,name=enable_limits,json=enableLimits,proto3" json:"enable_limits,omitempty"`
}

func (x *FieldTypeFloatConfig) Reset() {
	*x = FieldTypeFloatConfig{}
	mi := &file_field_type_float_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeFloatConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeFloatConfig) ProtoMessage() {}

func (x *FieldTypeFloatConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_float_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeFloatConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeFloatConfig) Descriptor() ([]byte, []int) {
	return file_field_type_float_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeFloatConfig) GetMinValue() float64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetMinValueInclusive() bool {
	if x != nil {
		return x.MinValueInclusive
	}
	return false
}

func (x *FieldTypeFloatConfig) GetMaxValue() float64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetMaxValueInclusive() bool {
	if x != nil {
		return x.MaxValueInclusive
	}
	return false
}

func (x *FieldTypeFloatConfig) GetAllowNegatives() bool {
	if x != nil {
		return x.AllowNegatives
	}
	return false
}

func (x *FieldTypeFloatConfig) GetNumberOfDecimals() int64 {
	if x != nil {
		return x.NumberOfDecimals
	}
	return 0
}

func (x *FieldTypeFloatConfig) GetSeparator() FieldTypeFloatConfigSeparator {
	if x != nil {
		return x.Separator
	}
	return FieldTypeFloatConfigSeparator_FIELD_TYPE_FLOAT_CONFIG_SEPARATOR_INVALID
}

func (x *FieldTypeFloatConfig) GetEnableLimits() bool {
	if x != nil {
		return x.EnableLimits
	}
	return false
}

var File_field_type_float_config_proto protoreflect.FileDescriptor

var file_field_type_float_config_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6e, 0x65, 0x6d, 0x22, 0xee, 0x02, 0x0a, 0x14, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x40,
	0x0a, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x2a, 0xa8, 0x01, 0x0a, 0x1d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x29, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x5f, 0x53, 0x45, 0x50, 0x41, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x2b, 0x0a, 0x27, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x5f, 0x53, 0x45, 0x50, 0x41, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x49, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53,
	0x45, 0x50, 0x41, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x10, 0x02,
	0x42, 0x2a, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x14, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01, 0x5a,
	0x0b, 0x6e, 0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_field_type_float_config_proto_rawDescOnce sync.Once
	file_field_type_float_config_proto_rawDescData = file_field_type_float_config_proto_rawDesc
)

func file_field_type_float_config_proto_rawDescGZIP() []byte {
	file_field_type_float_config_proto_rawDescOnce.Do(func() {
		file_field_type_float_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_type_float_config_proto_rawDescData)
	})
	return file_field_type_float_config_proto_rawDescData
}

var file_field_type_float_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_field_type_float_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_float_config_proto_goTypes = []any{
	(FieldTypeFloatConfigSeparator)(0), // 0: nem.FieldTypeFloatConfigSeparator
	(*FieldTypeFloatConfig)(nil),       // 1: nem.FieldTypeFloatConfig
}
var file_field_type_float_config_proto_depIdxs = []int32{
	0, // 0: nem.FieldTypeFloatConfig.separator:type_name -> nem.FieldTypeFloatConfigSeparator
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_field_type_float_config_proto_init() }
func file_field_type_float_config_proto_init() {
	if File_field_type_float_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_type_float_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_float_config_proto_goTypes,
		DependencyIndexes: file_field_type_float_config_proto_depIdxs,
		EnumInfos:         file_field_type_float_config_proto_enumTypes,
		MessageInfos:      file_field_type_float_config_proto_msgTypes,
	}.Build()
	File_field_type_float_config_proto = out.File
	file_field_type_float_config_proto_rawDesc = nil
	file_field_type_float_config_proto_goTypes = nil
	file_field_type_float_config_proto_depIdxs = nil
}
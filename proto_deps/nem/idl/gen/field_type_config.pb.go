// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: field_type_config.proto

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

type FieldTypeConfig struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Uuid          string                    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Integer       *FieldTypeIntegerConfig   `protobuf:"bytes,2,opt,name=integer,proto3" json:"integer,omitempty"`
	Float         *FieldTypeFloatConfig     `protobuf:"bytes,3,opt,name=float,proto3" json:"float,omitempty"`
	Decimal       *FieldTypeDecimalConfig   `protobuf:"bytes,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	Boolean       string                    `protobuf:"bytes,5,opt,name=boolean,proto3" json:"boolean,omitempty"`
	Char          *FieldTypeCharConfig      `protobuf:"bytes,6,opt,name=char,proto3" json:"char,omitempty"`
	Varchar       *FieldTypeVarcharConfig   `protobuf:"bytes,7,opt,name=varchar,proto3" json:"varchar,omitempty"`
	Text          *FieldTypeTextConfig      `protobuf:"bytes,8,opt,name=text,proto3" json:"text,omitempty"`
	Encrypted     *FieldTypeEncryptedConfig `protobuf:"bytes,9,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Email         *FieldTypeEmailConfig     `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Phone         *FieldTypePhoneConfig     `protobuf:"bytes,11,opt,name=phone,proto3" json:"phone,omitempty"`
	Url           *FieldTypeURLConfig       `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
	Location      string                    `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	Color         string                    `protobuf:"bytes,14,opt,name=color,proto3" json:"color,omitempty"`
	RichText      string                    `protobuf:"bytes,15,opt,name=rich_text,json=richText,proto3" json:"rich_text,omitempty"`
	Code          string                    `protobuf:"bytes,16,opt,name=code,proto3" json:"code,omitempty"`
	Markdown      string                    `protobuf:"bytes,17,opt,name=markdown,proto3" json:"markdown,omitempty"`
	File          *FieldTypeFileConfig      `protobuf:"bytes,18,opt,name=file,proto3" json:"file,omitempty"`
	Image         *FieldTypeFileConfig      `protobuf:"bytes,19,opt,name=image,proto3" json:"image,omitempty"`
	Video         *FieldTypeFileConfig      `protobuf:"bytes,20,opt,name=video,proto3" json:"video,omitempty"`
	Audio         *FieldTypeFileConfig      `protobuf:"bytes,21,opt,name=audio,proto3" json:"audio,omitempty"`
	Enum          *FieldTypeEnumConfig      `protobuf:"bytes,22,opt,name=enum,proto3" json:"enum,omitempty"`
	Json          *FieldTypeJSONConfig      `protobuf:"bytes,23,opt,name=json,proto3" json:"json,omitempty"`
	Array         *FieldTypeArrayConfig     `protobuf:"bytes,24,opt,name=array,proto3" json:"array,omitempty"`
	Date          *FieldTypeDateConfig      `protobuf:"bytes,25,opt,name=date,proto3" json:"date,omitempty"`
	Datetime      *FieldTypeDatetimeConfig  `protobuf:"bytes,26,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Time          string                    `protobuf:"bytes,27,opt,name=time,proto3" json:"time,omitempty"`
	Slug          *FieldTypeSlugConfig      `protobuf:"bytes,28,opt,name=slug,proto3" json:"slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldTypeConfig) Reset() {
	*x = FieldTypeConfig{}
	mi := &file_field_type_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldTypeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTypeConfig) ProtoMessage() {}

func (x *FieldTypeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_field_type_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTypeConfig.ProtoReflect.Descriptor instead.
func (*FieldTypeConfig) Descriptor() ([]byte, []int) {
	return file_field_type_config_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTypeConfig) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *FieldTypeConfig) GetInteger() *FieldTypeIntegerConfig {
	if x != nil {
		return x.Integer
	}
	return nil
}

func (x *FieldTypeConfig) GetFloat() *FieldTypeFloatConfig {
	if x != nil {
		return x.Float
	}
	return nil
}

func (x *FieldTypeConfig) GetDecimal() *FieldTypeDecimalConfig {
	if x != nil {
		return x.Decimal
	}
	return nil
}

func (x *FieldTypeConfig) GetBoolean() string {
	if x != nil {
		return x.Boolean
	}
	return ""
}

func (x *FieldTypeConfig) GetChar() *FieldTypeCharConfig {
	if x != nil {
		return x.Char
	}
	return nil
}

func (x *FieldTypeConfig) GetVarchar() *FieldTypeVarcharConfig {
	if x != nil {
		return x.Varchar
	}
	return nil
}

func (x *FieldTypeConfig) GetText() *FieldTypeTextConfig {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *FieldTypeConfig) GetEncrypted() *FieldTypeEncryptedConfig {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

func (x *FieldTypeConfig) GetEmail() *FieldTypeEmailConfig {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *FieldTypeConfig) GetPhone() *FieldTypePhoneConfig {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *FieldTypeConfig) GetUrl() *FieldTypeURLConfig {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *FieldTypeConfig) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *FieldTypeConfig) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *FieldTypeConfig) GetRichText() string {
	if x != nil {
		return x.RichText
	}
	return ""
}

func (x *FieldTypeConfig) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FieldTypeConfig) GetMarkdown() string {
	if x != nil {
		return x.Markdown
	}
	return ""
}

func (x *FieldTypeConfig) GetFile() *FieldTypeFileConfig {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FieldTypeConfig) GetImage() *FieldTypeFileConfig {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *FieldTypeConfig) GetVideo() *FieldTypeFileConfig {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *FieldTypeConfig) GetAudio() *FieldTypeFileConfig {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *FieldTypeConfig) GetEnum() *FieldTypeEnumConfig {
	if x != nil {
		return x.Enum
	}
	return nil
}

func (x *FieldTypeConfig) GetJson() *FieldTypeJSONConfig {
	if x != nil {
		return x.Json
	}
	return nil
}

func (x *FieldTypeConfig) GetArray() *FieldTypeArrayConfig {
	if x != nil {
		return x.Array
	}
	return nil
}

func (x *FieldTypeConfig) GetDate() *FieldTypeDateConfig {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *FieldTypeConfig) GetDatetime() *FieldTypeDatetimeConfig {
	if x != nil {
		return x.Datetime
	}
	return nil
}

func (x *FieldTypeConfig) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *FieldTypeConfig) GetSlug() *FieldTypeSlugConfig {
	if x != nil {
		return x.Slug
	}
	return nil
}

var File_field_type_config_proto protoreflect.FileDescriptor

var file_field_type_config_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x6d, 0x1a, 0x1d,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x72,
	0x63, 0x68, 0x61, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xaf, 0x09, 0x0a, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6d,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x12, 0x2f, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x12, 0x35, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x63, 0x68, 0x61, 0x72,
	0x12, 0x35, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x56, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07,
	0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x2f, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x52, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x69, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x2c,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e,
	0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x2e, 0x0a, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65,
	0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x2c, 0x0a, 0x04,
	0x65, 0x6e, 0x75, 0x6d, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x04, 0x6a, 0x73,
	0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x6d, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x1c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x42, 0x25, 0x0a, 0x03, 0x6e, 0x65, 0x6d, 0x42, 0x0f, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x01, 0x5a, 0x0b, 0x6e,
	0x65, 0x6d, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_field_type_config_proto_rawDescOnce sync.Once
	file_field_type_config_proto_rawDescData = file_field_type_config_proto_rawDesc
)

func file_field_type_config_proto_rawDescGZIP() []byte {
	file_field_type_config_proto_rawDescOnce.Do(func() {
		file_field_type_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_type_config_proto_rawDescData)
	})
	return file_field_type_config_proto_rawDescData
}

var file_field_type_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_field_type_config_proto_goTypes = []any{
	(*FieldTypeConfig)(nil),          // 0: nem.FieldTypeConfig
	(*FieldTypeIntegerConfig)(nil),   // 1: nem.FieldTypeIntegerConfig
	(*FieldTypeFloatConfig)(nil),     // 2: nem.FieldTypeFloatConfig
	(*FieldTypeDecimalConfig)(nil),   // 3: nem.FieldTypeDecimalConfig
	(*FieldTypeCharConfig)(nil),      // 4: nem.FieldTypeCharConfig
	(*FieldTypeVarcharConfig)(nil),   // 5: nem.FieldTypeVarcharConfig
	(*FieldTypeTextConfig)(nil),      // 6: nem.FieldTypeTextConfig
	(*FieldTypeEncryptedConfig)(nil), // 7: nem.FieldTypeEncryptedConfig
	(*FieldTypeEmailConfig)(nil),     // 8: nem.FieldTypeEmailConfig
	(*FieldTypePhoneConfig)(nil),     // 9: nem.FieldTypePhoneConfig
	(*FieldTypeURLConfig)(nil),       // 10: nem.FieldTypeURLConfig
	(*FieldTypeFileConfig)(nil),      // 11: nem.FieldTypeFileConfig
	(*FieldTypeEnumConfig)(nil),      // 12: nem.FieldTypeEnumConfig
	(*FieldTypeJSONConfig)(nil),      // 13: nem.FieldTypeJSONConfig
	(*FieldTypeArrayConfig)(nil),     // 14: nem.FieldTypeArrayConfig
	(*FieldTypeDateConfig)(nil),      // 15: nem.FieldTypeDateConfig
	(*FieldTypeDatetimeConfig)(nil),  // 16: nem.FieldTypeDatetimeConfig
	(*FieldTypeSlugConfig)(nil),      // 17: nem.FieldTypeSlugConfig
}
var file_field_type_config_proto_depIdxs = []int32{
	1,  // 0: nem.FieldTypeConfig.integer:type_name -> nem.FieldTypeIntegerConfig
	2,  // 1: nem.FieldTypeConfig.float:type_name -> nem.FieldTypeFloatConfig
	3,  // 2: nem.FieldTypeConfig.decimal:type_name -> nem.FieldTypeDecimalConfig
	4,  // 3: nem.FieldTypeConfig.char:type_name -> nem.FieldTypeCharConfig
	5,  // 4: nem.FieldTypeConfig.varchar:type_name -> nem.FieldTypeVarcharConfig
	6,  // 5: nem.FieldTypeConfig.text:type_name -> nem.FieldTypeTextConfig
	7,  // 6: nem.FieldTypeConfig.encrypted:type_name -> nem.FieldTypeEncryptedConfig
	8,  // 7: nem.FieldTypeConfig.email:type_name -> nem.FieldTypeEmailConfig
	9,  // 8: nem.FieldTypeConfig.phone:type_name -> nem.FieldTypePhoneConfig
	10, // 9: nem.FieldTypeConfig.url:type_name -> nem.FieldTypeURLConfig
	11, // 10: nem.FieldTypeConfig.file:type_name -> nem.FieldTypeFileConfig
	11, // 11: nem.FieldTypeConfig.image:type_name -> nem.FieldTypeFileConfig
	11, // 12: nem.FieldTypeConfig.video:type_name -> nem.FieldTypeFileConfig
	11, // 13: nem.FieldTypeConfig.audio:type_name -> nem.FieldTypeFileConfig
	12, // 14: nem.FieldTypeConfig.enum:type_name -> nem.FieldTypeEnumConfig
	13, // 15: nem.FieldTypeConfig.json:type_name -> nem.FieldTypeJSONConfig
	14, // 16: nem.FieldTypeConfig.array:type_name -> nem.FieldTypeArrayConfig
	15, // 17: nem.FieldTypeConfig.date:type_name -> nem.FieldTypeDateConfig
	16, // 18: nem.FieldTypeConfig.datetime:type_name -> nem.FieldTypeDatetimeConfig
	17, // 19: nem.FieldTypeConfig.slug:type_name -> nem.FieldTypeSlugConfig
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_field_type_config_proto_init() }
func file_field_type_config_proto_init() {
	if File_field_type_config_proto != nil {
		return
	}
	file_field_type_array_config_proto_init()
	file_field_type_char_config_proto_init()
	file_field_type_date_config_proto_init()
	file_field_type_datetime_config_proto_init()
	file_field_type_decimal_config_proto_init()
	file_field_type_email_config_proto_init()
	file_field_type_encrypted_config_proto_init()
	file_field_type_enum_config_proto_init()
	file_field_type_file_config_proto_init()
	file_field_type_float_config_proto_init()
	file_field_type_integer_config_proto_init()
	file_field_type_json_config_proto_init()
	file_field_type_phone_config_proto_init()
	file_field_type_slug_config_proto_init()
	file_field_type_text_config_proto_init()
	file_field_type_url_config_proto_init()
	file_field_type_varchar_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_field_type_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_field_type_config_proto_goTypes,
		DependencyIndexes: file_field_type_config_proto_depIdxs,
		MessageInfos:      file_field_type_config_proto_msgTypes,
	}.Build()
	File_field_type_config_proto = out.File
	file_field_type_config_proto_rawDesc = nil
	file_field_type_config_proto_goTypes = nil
	file_field_type_config_proto_depIdxs = nil
}

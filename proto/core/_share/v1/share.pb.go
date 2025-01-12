// proto/core/_share/v1/share.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: proto/core/_share/v1/share.proto

package sharepb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ShareQID (uuid, revision) optional
type ShareQID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kind
	Kind string `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// key
	Key string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	// uid
	Uid string `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	// uuid
	Uuid string `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// rev - revision string
	Rev string `protobuf:"bytes,8,opt,name=rev,proto3" json:"rev,omitempty"`
}

func (x *ShareQID) Reset() {
	*x = ShareQID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core__share_v1_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareQID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareQID) ProtoMessage() {}

func (x *ShareQID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core__share_v1_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareQID.ProtoReflect.Descriptor instead.
func (*ShareQID) Descriptor() ([]byte, []int) {
	return file_proto_core__share_v1_share_proto_rawDescGZIP(), []int{0}
}

func (x *ShareQID) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ShareQID) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ShareQID) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ShareQID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ShareQID) GetRev() string {
	if x != nil {
		return x.Rev
	}
	return ""
}

// MetaDescription - title, topic, label, description
type MetaDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// title - (0-256 characters) optional
	Title *string `protobuf:"bytes,4,opt,name=title,proto3,oneof" json:"title,omitempty"`
	// topic - (0-256 characters) optional
	Topic *string `protobuf:"bytes,5,opt,name=topic,proto3,oneof" json:"topic,omitempty"`
	// label - (0-256 characters) optional
	Label *string `protobuf:"bytes,6,opt,name=label,proto3,oneof" json:"label,omitempty"`
	// description - (0-1024 characters) optional
	Description *string `protobuf:"bytes,7,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *MetaDescription) Reset() {
	*x = MetaDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core__share_v1_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDescription) ProtoMessage() {}

func (x *MetaDescription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core__share_v1_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDescription.ProtoReflect.Descriptor instead.
func (*MetaDescription) Descriptor() ([]byte, []int) {
	return file_proto_core__share_v1_share_proto_rawDescGZIP(), []int{1}
}

func (x *MetaDescription) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *MetaDescription) GetTopic() string {
	if x != nil && x.Topic != nil {
		return *x.Topic
	}
	return ""
}

func (x *MetaDescription) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *MetaDescription) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

// MetaMultiDescription - []labels, []tags
type MetaMultiDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// labels - (0-32 elements, unique, tag max_len 32) optional
	Labels []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	// tags - (0-32 elements, unique, tag max_len 32) optional
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *MetaMultiDescription) Reset() {
	*x = MetaMultiDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core__share_v1_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaMultiDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaMultiDescription) ProtoMessage() {}

func (x *MetaMultiDescription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core__share_v1_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaMultiDescription.ProtoReflect.Descriptor instead.
func (*MetaMultiDescription) Descriptor() ([]byte, []int) {
	return file_proto_core__share_v1_share_proto_rawDescGZIP(), []int{2}
}

func (x *MetaMultiDescription) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MetaMultiDescription) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_proto_core__share_v1_share_proto protoreflect.FileDescriptor

var file_proto_core__share_v1_share_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe0,
	0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x01,
	0xd0, 0x01, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0x18, 0x80, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe0, 0x41, 0x01,
	0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x01, 0xd0, 0x01,
	0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xe0, 0x41, 0x01, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0xfa,
	0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x01, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x03, 0x72, 0x65, 0x76, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xe0,
	0x41, 0x01, 0xe0, 0x41, 0x05, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x01, 0xd0, 0x01, 0x01,
	0x52, 0x03, 0x72, 0x65, 0x76, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x89, 0x02, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe0, 0x41, 0x01, 0xfa,
	0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x02, 0xd0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe0, 0x41, 0x01, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18,
	0x80, 0x02, 0xd0, 0x01, 0x01, 0x48, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xe0, 0x41, 0x01, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x02, 0xd0, 0x01, 0x01,
	0x48, 0x02, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0xe0, 0x41, 0x01, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x18, 0x80, 0x08, 0xd0, 0x01,
	0x01, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x15, 0xe0, 0x41,
	0x01, 0xfa, 0x42, 0x0f, 0x92, 0x01, 0x0c, 0x10, 0x20, 0x18, 0x01, 0x22, 0x04, 0x72, 0x02, 0x18,
	0x20, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x15, 0xe0, 0x41, 0x01, 0xfa, 0x42,
	0x0f, 0x92, 0x01, 0x0c, 0x10, 0x20, 0x18, 0x01, 0x22, 0x04, 0x72, 0x02, 0x18, 0x20, 0x28, 0x01,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04,
	0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x42, 0x24, 0x5a, 0x22, 0x78, 0x2d,
	0x67, 0x77, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_core__share_v1_share_proto_rawDescOnce sync.Once
	file_proto_core__share_v1_share_proto_rawDescData = file_proto_core__share_v1_share_proto_rawDesc
)

func file_proto_core__share_v1_share_proto_rawDescGZIP() []byte {
	file_proto_core__share_v1_share_proto_rawDescOnce.Do(func() {
		file_proto_core__share_v1_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_core__share_v1_share_proto_rawDescData)
	})
	return file_proto_core__share_v1_share_proto_rawDescData
}

var file_proto_core__share_v1_share_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_core__share_v1_share_proto_goTypes = []interface{}{
	(*ShareQID)(nil),             // 0: proto.core._share.v1.ShareQID
	(*MetaDescription)(nil),      // 1: proto.core._share.v1.MetaDescription
	(*MetaMultiDescription)(nil), // 2: proto.core._share.v1.MetaMultiDescription
}
var file_proto_core__share_v1_share_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_core__share_v1_share_proto_init() }
func file_proto_core__share_v1_share_proto_init() {
	if File_proto_core__share_v1_share_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_core__share_v1_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareQID); i {
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
		file_proto_core__share_v1_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaDescription); i {
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
		file_proto_core__share_v1_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaMultiDescription); i {
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
	file_proto_core__share_v1_share_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_core__share_v1_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_core__share_v1_share_proto_goTypes,
		DependencyIndexes: file_proto_core__share_v1_share_proto_depIdxs,
		MessageInfos:      file_proto_core__share_v1_share_proto_msgTypes,
	}.Build()
	File_proto_core__share_v1_share_proto = out.File
	file_proto_core__share_v1_share_proto_rawDesc = nil
	file_proto_core__share_v1_share_proto_goTypes = nil
	file_proto_core__share_v1_share_proto_depIdxs = nil
}

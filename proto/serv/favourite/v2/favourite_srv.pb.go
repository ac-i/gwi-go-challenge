// proto/serv/favourite/v2/favourite_srv.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: proto/serv/favourite/v2/favourite_srv.proto

package favouritepbapiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v11 "x-gwi/proto/core/_share/v1"
	v1 "x-gwi/proto/core/favourite/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_serv_favourite_v2_favourite_srv_proto protoreflect.FileDescriptor

var file_proto_serv_favourite_v2_favourite_srv_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x66, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x9d, 0x06, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x1a, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x5a, 0x1a, 0x3a, 0x01, 0x2a,
	0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b,
	0x6b, 0x65, 0x79, 0x7d, 0x12, 0x4e, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x74, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x1a, 0x26, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22,
	0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x32, 0x2c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x71, 0x69, 0x64, 0x2e, 0x6b, 0x65, 0x79, 0x7d, 0x2f, 0x7b,
	0x71, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x76, 0x7d, 0x12, 0x78, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51,
	0x49, 0x44, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x6b, 0x65,
	0x79, 0x7d, 0x12, 0x9a, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x1a, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x40, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3a, 0x3a, 0x01, 0x2a, 0x5a, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x66,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x42,
	0x30, 0x5a, 0x2e, 0x78, 0x2d, 0x67, 0x77, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x32,
	0x3b, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x69, 0x74, 0x65, 0x70, 0x62, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_serv_favourite_v2_favourite_srv_proto_goTypes = []interface{}{
	(*v1.FavouriteCore)(nil), // 0: proto.core.favourite.v1.FavouriteCore
	(*v11.ShareQID)(nil),     // 1: proto.core._share.v1.ShareQID
}
var file_proto_serv_favourite_v2_favourite_srv_proto_depIdxs = []int32{
	0, // 0: proto.serv.favourite.v2.FavouriteService.Create:input_type -> proto.core.favourite.v1.FavouriteCore
	1, // 1: proto.serv.favourite.v2.FavouriteService.Get:input_type -> proto.core._share.v1.ShareQID
	1, // 2: proto.serv.favourite.v2.FavouriteService.Gett:input_type -> proto.core._share.v1.ShareQID
	0, // 3: proto.serv.favourite.v2.FavouriteService.Update:input_type -> proto.core.favourite.v1.FavouriteCore
	1, // 4: proto.serv.favourite.v2.FavouriteService.Delete:input_type -> proto.core._share.v1.ShareQID
	0, // 5: proto.serv.favourite.v2.FavouriteService.List:input_type -> proto.core.favourite.v1.FavouriteCore
	0, // 6: proto.serv.favourite.v2.FavouriteService.Create:output_type -> proto.core.favourite.v1.FavouriteCore
	0, // 7: proto.serv.favourite.v2.FavouriteService.Get:output_type -> proto.core.favourite.v1.FavouriteCore
	0, // 8: proto.serv.favourite.v2.FavouriteService.Gett:output_type -> proto.core.favourite.v1.FavouriteCore
	0, // 9: proto.serv.favourite.v2.FavouriteService.Update:output_type -> proto.core.favourite.v1.FavouriteCore
	0, // 10: proto.serv.favourite.v2.FavouriteService.Delete:output_type -> proto.core.favourite.v1.FavouriteCore
	0, // 11: proto.serv.favourite.v2.FavouriteService.List:output_type -> proto.core.favourite.v1.FavouriteCore
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_serv_favourite_v2_favourite_srv_proto_init() }
func file_proto_serv_favourite_v2_favourite_srv_proto_init() {
	if File_proto_serv_favourite_v2_favourite_srv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_serv_favourite_v2_favourite_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_serv_favourite_v2_favourite_srv_proto_goTypes,
		DependencyIndexes: file_proto_serv_favourite_v2_favourite_srv_proto_depIdxs,
	}.Build()
	File_proto_serv_favourite_v2_favourite_srv_proto = out.File
	file_proto_serv_favourite_v2_favourite_srv_proto_rawDesc = nil
	file_proto_serv_favourite_v2_favourite_srv_proto_goTypes = nil
	file_proto_serv_favourite_v2_favourite_srv_proto_depIdxs = nil
}

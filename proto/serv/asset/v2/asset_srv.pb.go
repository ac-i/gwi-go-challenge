// proto/serv/asset/v2/asset_srv.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: proto/serv/asset/v2/asset_srv.proto

package assetpbapiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v11 "x-gwi/proto/core/_share/v1"
	v1 "x-gwi/proto/core/asset/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_serv_asset_v2_asset_srv_proto protoreflect.FileDescriptor

var file_proto_serv_asset_v2_asset_srv_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xae, 0x04, 0x0a, 0x0c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01,
	0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x7e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x37,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x5a, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x17,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x12, 0x46, 0x0a, 0x04, 0x47, 0x65, 0x74, 0x74, 0x12,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x1a,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x12,
	0x7d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2d, 0x3a, 0x01, 0x2a, 0x32, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x71, 0x69, 0x64, 0x2e,
	0x6b, 0x65, 0x79, 0x7d, 0x2f, 0x7b, 0x71, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x76, 0x7d, 0x12, 0x6c,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x51, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x2a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x28, 0x5a, 0x26,
	0x78, 0x2d, 0x67, 0x77, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x70,
	0x62, 0x61, 0x70, 0x69, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_serv_asset_v2_asset_srv_proto_goTypes = []interface{}{
	(*v1.AssetCore)(nil), // 0: proto.core.asset.v1.AssetCore
	(*v11.ShareQID)(nil), // 1: proto.core._share.v1.ShareQID
}
var file_proto_serv_asset_v2_asset_srv_proto_depIdxs = []int32{
	0, // 0: proto.serv.asset.v2.AssetService.Create:input_type -> proto.core.asset.v1.AssetCore
	1, // 1: proto.serv.asset.v2.AssetService.Get:input_type -> proto.core._share.v1.ShareQID
	1, // 2: proto.serv.asset.v2.AssetService.Gett:input_type -> proto.core._share.v1.ShareQID
	0, // 3: proto.serv.asset.v2.AssetService.Update:input_type -> proto.core.asset.v1.AssetCore
	1, // 4: proto.serv.asset.v2.AssetService.Delete:input_type -> proto.core._share.v1.ShareQID
	0, // 5: proto.serv.asset.v2.AssetService.Create:output_type -> proto.core.asset.v1.AssetCore
	0, // 6: proto.serv.asset.v2.AssetService.Get:output_type -> proto.core.asset.v1.AssetCore
	0, // 7: proto.serv.asset.v2.AssetService.Gett:output_type -> proto.core.asset.v1.AssetCore
	0, // 8: proto.serv.asset.v2.AssetService.Update:output_type -> proto.core.asset.v1.AssetCore
	0, // 9: proto.serv.asset.v2.AssetService.Delete:output_type -> proto.core.asset.v1.AssetCore
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_serv_asset_v2_asset_srv_proto_init() }
func file_proto_serv_asset_v2_asset_srv_proto_init() {
	if File_proto_serv_asset_v2_asset_srv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_serv_asset_v2_asset_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_serv_asset_v2_asset_srv_proto_goTypes,
		DependencyIndexes: file_proto_serv_asset_v2_asset_srv_proto_depIdxs,
	}.Build()
	File_proto_serv_asset_v2_asset_srv_proto = out.File
	file_proto_serv_asset_v2_asset_srv_proto_rawDesc = nil
	file_proto_serv_asset_v2_asset_srv_proto_goTypes = nil
	file_proto_serv_asset_v2_asset_srv_proto_depIdxs = nil
}

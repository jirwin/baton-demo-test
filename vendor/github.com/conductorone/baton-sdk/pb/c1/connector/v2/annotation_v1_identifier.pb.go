// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: c1/connector/v2/annotation_v1_identifier.proto

package v2

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

type V1Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *V1Identifier) Reset() {
	*x = V1Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c1_connector_v2_annotation_v1_identifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1Identifier) ProtoMessage() {}

func (x *V1Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_c1_connector_v2_annotation_v1_identifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1Identifier.ProtoReflect.Descriptor instead.
func (*V1Identifier) Descriptor() ([]byte, []int) {
	return file_c1_connector_v2_annotation_v1_identifier_proto_rawDescGZIP(), []int{0}
}

func (x *V1Identifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_c1_connector_v2_annotation_v1_identifier_proto protoreflect.FileDescriptor

var file_c1_connector_v2_annotation_v1_identifier_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x63, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x32, 0x22, 0x1e, 0x0a, 0x0c, 0x56, 0x31, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x6f, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x74,
	0x6f, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_c1_connector_v2_annotation_v1_identifier_proto_rawDescOnce sync.Once
	file_c1_connector_v2_annotation_v1_identifier_proto_rawDescData = file_c1_connector_v2_annotation_v1_identifier_proto_rawDesc
)

func file_c1_connector_v2_annotation_v1_identifier_proto_rawDescGZIP() []byte {
	file_c1_connector_v2_annotation_v1_identifier_proto_rawDescOnce.Do(func() {
		file_c1_connector_v2_annotation_v1_identifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_c1_connector_v2_annotation_v1_identifier_proto_rawDescData)
	})
	return file_c1_connector_v2_annotation_v1_identifier_proto_rawDescData
}

var file_c1_connector_v2_annotation_v1_identifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_c1_connector_v2_annotation_v1_identifier_proto_goTypes = []interface{}{
	(*V1Identifier)(nil), // 0: c1.connector.v2.V1Identifier
}
var file_c1_connector_v2_annotation_v1_identifier_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_c1_connector_v2_annotation_v1_identifier_proto_init() }
func file_c1_connector_v2_annotation_v1_identifier_proto_init() {
	if File_c1_connector_v2_annotation_v1_identifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c1_connector_v2_annotation_v1_identifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1Identifier); i {
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
			RawDescriptor: file_c1_connector_v2_annotation_v1_identifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c1_connector_v2_annotation_v1_identifier_proto_goTypes,
		DependencyIndexes: file_c1_connector_v2_annotation_v1_identifier_proto_depIdxs,
		MessageInfos:      file_c1_connector_v2_annotation_v1_identifier_proto_msgTypes,
	}.Build()
	File_c1_connector_v2_annotation_v1_identifier_proto = out.File
	file_c1_connector_v2_annotation_v1_identifier_proto_rawDesc = nil
	file_c1_connector_v2_annotation_v1_identifier_proto_goTypes = nil
	file_c1_connector_v2_annotation_v1_identifier_proto_depIdxs = nil
}

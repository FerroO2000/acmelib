// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: acmelib/v1/signal_type.proto

package acmelibv1

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

type SignalType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *SignalType) Reset() {
	*x = SignalType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acmelib_v1_signal_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalType) ProtoMessage() {}

func (x *SignalType) ProtoReflect() protoreflect.Message {
	mi := &file_acmelib_v1_signal_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalType.ProtoReflect.Descriptor instead.
func (*SignalType) Descriptor() ([]byte, []int) {
	return file_acmelib_v1_signal_type_proto_rawDescGZIP(), []int{0}
}

func (x *SignalType) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_acmelib_v1_signal_type_proto protoreflect.FileDescriptor

var file_acmelib_v1_signal_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x61, 0x63, 0x6d, 0x65,
	0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x80, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x14, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa,
	0x02, 0x0a, 0x41, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x41,
	0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x41, 0x63, 0x6d, 0x65,
	0x6c, 0x69, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x41, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acmelib_v1_signal_type_proto_rawDescOnce sync.Once
	file_acmelib_v1_signal_type_proto_rawDescData = file_acmelib_v1_signal_type_proto_rawDesc
)

func file_acmelib_v1_signal_type_proto_rawDescGZIP() []byte {
	file_acmelib_v1_signal_type_proto_rawDescOnce.Do(func() {
		file_acmelib_v1_signal_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_acmelib_v1_signal_type_proto_rawDescData)
	})
	return file_acmelib_v1_signal_type_proto_rawDescData
}

var file_acmelib_v1_signal_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_acmelib_v1_signal_type_proto_goTypes = []interface{}{
	(*SignalType)(nil), // 0: acmelib.v1.SignalType
	(*Entity)(nil),     // 1: acmelib.v1.Entity
}
var file_acmelib_v1_signal_type_proto_depIdxs = []int32{
	1, // 0: acmelib.v1.SignalType.entity:type_name -> acmelib.v1.Entity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acmelib_v1_signal_type_proto_init() }
func file_acmelib_v1_signal_type_proto_init() {
	if File_acmelib_v1_signal_type_proto != nil {
		return
	}
	file_acmelib_v1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_acmelib_v1_signal_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalType); i {
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
			RawDescriptor: file_acmelib_v1_signal_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acmelib_v1_signal_type_proto_goTypes,
		DependencyIndexes: file_acmelib_v1_signal_type_proto_depIdxs,
		MessageInfos:      file_acmelib_v1_signal_type_proto_msgTypes,
	}.Build()
	File_acmelib_v1_signal_type_proto = out.File
	file_acmelib_v1_signal_type_proto_rawDesc = nil
	file_acmelib_v1_signal_type_proto_goTypes = nil
	file_acmelib_v1_signal_type_proto_depIdxs = nil
}
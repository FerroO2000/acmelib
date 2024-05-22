// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: acmelib/v1/signal_payload.proto

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

type SignalPayloadRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalEntityId string `protobuf:"bytes,1,opt,name=signal_entity_id,json=signalEntityId,proto3" json:"signal_entity_id,omitempty"`
	StartBit       uint32 `protobuf:"varint,2,opt,name=start_bit,json=startBit,proto3" json:"start_bit,omitempty"`
}

func (x *SignalPayloadRef) Reset() {
	*x = SignalPayloadRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acmelib_v1_signal_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalPayloadRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalPayloadRef) ProtoMessage() {}

func (x *SignalPayloadRef) ProtoReflect() protoreflect.Message {
	mi := &file_acmelib_v1_signal_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalPayloadRef.ProtoReflect.Descriptor instead.
func (*SignalPayloadRef) Descriptor() ([]byte, []int) {
	return file_acmelib_v1_signal_payload_proto_rawDescGZIP(), []int{0}
}

func (x *SignalPayloadRef) GetSignalEntityId() string {
	if x != nil {
		return x.SignalEntityId
	}
	return ""
}

func (x *SignalPayloadRef) GetStartBit() uint32 {
	if x != nil {
		return x.StartBit
	}
	return 0
}

type SignalPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*SignalPayloadRef `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *SignalPayload) Reset() {
	*x = SignalPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acmelib_v1_signal_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalPayload) ProtoMessage() {}

func (x *SignalPayload) ProtoReflect() protoreflect.Message {
	mi := &file_acmelib_v1_signal_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalPayload.ProtoReflect.Descriptor instead.
func (*SignalPayload) Descriptor() ([]byte, []int) {
	return file_acmelib_v1_signal_payload_proto_rawDescGZIP(), []int{1}
}

func (x *SignalPayload) GetRefs() []*SignalPayloadRef {
	if x != nil {
		return x.Refs
	}
	return nil
}

var File_acmelib_v1_signal_payload_proto protoreflect.FileDescriptor

var file_acmelib_v1_signal_payload_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x59, 0x0a,
	0x10, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x66, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x69, 0x74, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x65, 0x66,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x66, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x42, 0x83, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x41, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a,
	0x41, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x41, 0x63, 0x6d,
	0x65, 0x6c, 0x69, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x41, 0x63, 0x6d, 0x65, 0x6c, 0x69, 0x62, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acmelib_v1_signal_payload_proto_rawDescOnce sync.Once
	file_acmelib_v1_signal_payload_proto_rawDescData = file_acmelib_v1_signal_payload_proto_rawDesc
)

func file_acmelib_v1_signal_payload_proto_rawDescGZIP() []byte {
	file_acmelib_v1_signal_payload_proto_rawDescOnce.Do(func() {
		file_acmelib_v1_signal_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_acmelib_v1_signal_payload_proto_rawDescData)
	})
	return file_acmelib_v1_signal_payload_proto_rawDescData
}

var file_acmelib_v1_signal_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_acmelib_v1_signal_payload_proto_goTypes = []interface{}{
	(*SignalPayloadRef)(nil), // 0: acmelib.v1.SignalPayloadRef
	(*SignalPayload)(nil),    // 1: acmelib.v1.SignalPayload
}
var file_acmelib_v1_signal_payload_proto_depIdxs = []int32{
	0, // 0: acmelib.v1.SignalPayload.refs:type_name -> acmelib.v1.SignalPayloadRef
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acmelib_v1_signal_payload_proto_init() }
func file_acmelib_v1_signal_payload_proto_init() {
	if File_acmelib_v1_signal_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acmelib_v1_signal_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalPayloadRef); i {
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
		file_acmelib_v1_signal_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalPayload); i {
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
			RawDescriptor: file_acmelib_v1_signal_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acmelib_v1_signal_payload_proto_goTypes,
		DependencyIndexes: file_acmelib_v1_signal_payload_proto_depIdxs,
		MessageInfos:      file_acmelib_v1_signal_payload_proto_msgTypes,
	}.Build()
	File_acmelib_v1_signal_payload_proto = out.File
	file_acmelib_v1_signal_payload_proto_rawDesc = nil
	file_acmelib_v1_signal_payload_proto_goTypes = nil
	file_acmelib_v1_signal_payload_proto_depIdxs = nil
}

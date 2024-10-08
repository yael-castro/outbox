// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: purchase.proto

package pb

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

type Purchase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId int64 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *Purchase) Reset() {
	*x = Purchase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Purchase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Purchase) ProtoMessage() {}

func (x *Purchase) ProtoReflect() protoreflect.Message {
	mi := &file_purchase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Purchase.ProtoReflect.Descriptor instead.
func (*Purchase) Descriptor() ([]byte, []int) {
	return file_purchase_proto_rawDescGZIP(), []int{0}
}

func (x *Purchase) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Purchase) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_purchase_proto protoreflect.FileDescriptor

var file_purchase_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchase_proto_rawDescOnce sync.Once
	file_purchase_proto_rawDescData = file_purchase_proto_rawDesc
)

func file_purchase_proto_rawDescGZIP() []byte {
	file_purchase_proto_rawDescOnce.Do(func() {
		file_purchase_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchase_proto_rawDescData)
	})
	return file_purchase_proto_rawDescData
}

var file_purchase_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_purchase_proto_goTypes = []any{
	(*Purchase)(nil), // 0: Purchase
}
var file_purchase_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_purchase_proto_init() }
func file_purchase_proto_init() {
	if File_purchase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_purchase_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Purchase); i {
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
			RawDescriptor: file_purchase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_purchase_proto_goTypes,
		DependencyIndexes: file_purchase_proto_depIdxs,
		MessageInfos:      file_purchase_proto_msgTypes,
	}.Build()
	File_purchase_proto = out.File
	file_purchase_proto_rawDesc = nil
	file_purchase_proto_goTypes = nil
	file_purchase_proto_depIdxs = nil
}

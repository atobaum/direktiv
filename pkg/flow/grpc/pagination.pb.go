// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: pkg/flow/grpc/pagination.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasNextPage     bool   `protobuf:"varint,1,opt,name=hasNextPage,proto3" json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `protobuf:"varint,2,opt,name=hasPreviousPage,proto3" json:"hasPreviousPage,omitempty"`
	StartCursor     string `protobuf:"bytes,3,opt,name=startCursor,proto3" json:"startCursor,omitempty"`
	EndCursor       string `protobuf:"bytes,4,opt,name=endCursor,proto3" json:"endCursor,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PageInfo) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

func (x *PageInfo) GetHasPreviousPage() bool {
	if x != nil {
		return x.HasPreviousPage
	}
	return false
}

func (x *PageInfo) GetStartCursor() string {
	if x != nil {
		return x.StartCursor
	}
	return ""
}

func (x *PageInfo) GetEndCursor() string {
	if x != nil {
		return x.EndCursor
	}
	return ""
}

type PageOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field     string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *PageOrder) Reset() {
	*x = PageOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageOrder) ProtoMessage() {}

func (x *PageOrder) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageOrder.ProtoReflect.Descriptor instead.
func (*PageOrder) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PageOrder) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *PageOrder) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type PageFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Val   string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *PageFilter) Reset() {
	*x = PageFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageFilter) ProtoMessage() {}

func (x *PageFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageFilter.ProtoReflect.Descriptor instead.
func (*PageFilter) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *PageFilter) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *PageFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PageFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After  string      `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	First  int32       `protobuf:"varint,2,opt,name=first,proto3" json:"first,omitempty"`
	Before string      `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	Last   int32       `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	Order  *PageOrder  `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Filter *PageFilter `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_pagination_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_pagination_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *Pagination) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *Pagination) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *Pagination) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *Pagination) GetOrder() *PageOrder {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *Pagination) GetFilter() *PageFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_pkg_flow_grpc_pagination_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_pagination_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x22,
	0x96, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x68, 0x61, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0xc7, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74,
	0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_pagination_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_pagination_proto_rawDescData = file_pkg_flow_grpc_pagination_proto_rawDesc
)

func file_pkg_flow_grpc_pagination_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_pagination_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_pagination_proto_rawDescData)
	})
	return file_pkg_flow_grpc_pagination_proto_rawDescData
}

var file_pkg_flow_grpc_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_flow_grpc_pagination_proto_goTypes = []interface{}{
	(*PageInfo)(nil),   // 0: direktiv_flow.PageInfo
	(*PageOrder)(nil),  // 1: direktiv_flow.PageOrder
	(*PageFilter)(nil), // 2: direktiv_flow.PageFilter
	(*Pagination)(nil), // 3: direktiv_flow.Pagination
}
var file_pkg_flow_grpc_pagination_proto_depIdxs = []int32{
	1, // 0: direktiv_flow.Pagination.order:type_name -> direktiv_flow.PageOrder
	2, // 1: direktiv_flow.Pagination.filter:type_name -> direktiv_flow.PageFilter
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_pagination_proto_init() }
func file_pkg_flow_grpc_pagination_proto_init() {
	if File_pkg_flow_grpc_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_pkg_flow_grpc_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageOrder); i {
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
		file_pkg_flow_grpc_pagination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageFilter); i {
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
		file_pkg_flow_grpc_pagination_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
			RawDescriptor: file_pkg_flow_grpc_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_pagination_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_pagination_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_pagination_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_pagination_proto = out.File
	file_pkg_flow_grpc_pagination_proto_rawDesc = nil
	file_pkg_flow_grpc_pagination_proto_goTypes = nil
	file_pkg_flow_grpc_pagination_proto_depIdxs = nil
}
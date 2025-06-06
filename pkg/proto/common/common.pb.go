// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: common/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListParamsSt struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Page           int64                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize       int64                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	WithTotalCount bool                   `protobuf:"varint,4,opt,name=with_total_count,json=withTotalCount,proto3" json:"with_total_count,omitempty"`
	OnlyCount      bool                   `protobuf:"varint,5,opt,name=only_count,json=onlyCount,proto3" json:"only_count,omitempty"`
	SortName       string                 `protobuf:"bytes,6,opt,name=sort_name,json=sortName,proto3" json:"sort_name,omitempty"`
	Sort           []string               `protobuf:"bytes,7,rep,name=sort,proto3" json:"sort,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListParamsSt) Reset() {
	*x = ListParamsSt{}
	mi := &file_common_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListParamsSt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParamsSt) ProtoMessage() {}

func (x *ListParamsSt) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParamsSt.ProtoReflect.Descriptor instead.
func (*ListParamsSt) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *ListParamsSt) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListParamsSt) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListParamsSt) GetWithTotalCount() bool {
	if x != nil {
		return x.WithTotalCount
	}
	return false
}

func (x *ListParamsSt) GetOnlyCount() bool {
	if x != nil {
		return x.OnlyCount
	}
	return false
}

func (x *ListParamsSt) GetSortName() string {
	if x != nil {
		return x.SortName
	}
	return ""
}

func (x *ListParamsSt) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

type PaginationInfoSt struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int64                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalCount    int64                  `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationInfoSt) Reset() {
	*x = PaginationInfoSt{}
	mi := &file_common_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationInfoSt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationInfoSt) ProtoMessage() {}

func (x *PaginationInfoSt) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationInfoSt.ProtoReflect.Descriptor instead.
func (*PaginationInfoSt) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationInfoSt) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationInfoSt) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationInfoSt) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type I18NFieldSt struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kz            string                 `protobuf:"bytes,1,opt,name=kz,proto3" json:"kz,omitempty"`
	Ru            string                 `protobuf:"bytes,2,opt,name=ru,proto3" json:"ru,omitempty"`
	En            string                 `protobuf:"bytes,3,opt,name=en,proto3" json:"en,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *I18NFieldSt) Reset() {
	*x = I18NFieldSt{}
	mi := &file_common_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *I18NFieldSt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I18NFieldSt) ProtoMessage() {}

func (x *I18NFieldSt) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I18NFieldSt.ProtoReflect.Descriptor instead.
func (*I18NFieldSt) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *I18NFieldSt) GetKz() string {
	if x != nil {
		return x.Kz
	}
	return ""
}

func (x *I18NFieldSt) GetRu() string {
	if x != nil {
		return x.Ru
	}
	return ""
}

func (x *I18NFieldSt) GetEn() string {
	if x != nil {
		return x.En
	}
	return ""
}

type ErrorRep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Fields        map[string]string      `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorRep) Reset() {
	*x = ErrorRep{}
	mi := &file_common_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorRep) ProtoMessage() {}

func (x *ErrorRep) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorRep.ProtoReflect.Descriptor instead.
func (*ErrorRep) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorRep) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorRep) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorRep) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Extension     string                 `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	MimeType      string                 `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Data          []byte                 `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_common_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *File) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_common_proto protoreflect.FileDescriptor

const file_common_common_proto_rawDesc = "" +
	"\n" +
	"\x13common/common.proto\x12\x06common\"\xb9\x01\n" +
	"\fListParamsSt\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x03R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x03R\bpageSize\x12(\n" +
	"\x10with_total_count\x18\x04 \x01(\bR\x0ewithTotalCount\x12\x1d\n" +
	"\n" +
	"only_count\x18\x05 \x01(\bR\tonlyCount\x12\x1b\n" +
	"\tsort_name\x18\x06 \x01(\tR\bsortName\x12\x12\n" +
	"\x04sort\x18\a \x03(\tR\x04sort\"d\n" +
	"\x10PaginationInfoSt\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x03R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1f\n" +
	"\vtotal_count\x18\x03 \x01(\x03R\n" +
	"totalCount\"=\n" +
	"\vI18nFieldSt\x12\x0e\n" +
	"\x02kz\x18\x01 \x01(\tR\x02kz\x12\x0e\n" +
	"\x02ru\x18\x02 \x01(\tR\x02ru\x12\x0e\n" +
	"\x02en\x18\x03 \x01(\tR\x02en\"\xa9\x01\n" +
	"\bErrorRep\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x124\n" +
	"\x06fields\x18\x03 \x03(\v2\x1c.common.ErrorRep.FieldsEntryR\x06fields\x1a9\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"i\n" +
	"\x04File\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1c\n" +
	"\textension\x18\x02 \x01(\tR\textension\x12\x1b\n" +
	"\tmime_type\x18\x03 \x01(\tR\bmimeType\x12\x12\n" +
	"\x04data\x18\x04 \x01(\fR\x04dataB/Z-github.com/mechta-market/nsi/pkg/proto/commonb\x06proto3"

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData []byte
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)))
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_common_proto_goTypes = []any{
	(*ListParamsSt)(nil),     // 0: common.ListParamsSt
	(*PaginationInfoSt)(nil), // 1: common.PaginationInfoSt
	(*I18NFieldSt)(nil),      // 2: common.I18nFieldSt
	(*ErrorRep)(nil),         // 3: common.ErrorRep
	(*File)(nil),             // 4: common.File
	nil,                      // 5: common.ErrorRep.FieldsEntry
}
var file_common_common_proto_depIdxs = []int32{
	5, // 0: common.ErrorRep.fields:type_name -> common.ErrorRep.FieldsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}

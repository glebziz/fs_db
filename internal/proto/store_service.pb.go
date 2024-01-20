// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: store_service.proto

package store

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

type ChunkSize int32

const (
	ChunkSize_EMPTY ChunkSize = 0
	ChunkSize_MAX   ChunkSize = 2048
)

// Enum value maps for ChunkSize.
var (
	ChunkSize_name = map[int32]string{
		0:    "EMPTY",
		2048: "MAX",
	}
	ChunkSize_value = map[string]int32{
		"EMPTY": 0,
		"MAX":   2048,
	}
)

func (x ChunkSize) Enum() *ChunkSize {
	p := new(ChunkSize)
	*p = x
	return p
}

func (x ChunkSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChunkSize) Descriptor() protoreflect.EnumDescriptor {
	return file_store_service_proto_enumTypes[0].Descriptor()
}

func (ChunkSize) Type() protoreflect.EnumType {
	return &file_store_service_proto_enumTypes[0]
}

func (x ChunkSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChunkSize.Descriptor instead.
func (ChunkSize) EnumDescriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{0}
}

type TxIsoLevel int32

const (
	TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED TxIsoLevel = 0
	TxIsoLevel_ISO_LEVEL_READ_COMMITTED   TxIsoLevel = 1
	TxIsoLevel_ISO_LEVEL_REPEATABLE_READ  TxIsoLevel = 2
	TxIsoLevel_ISO_LEVEL_SERIALIZABLE     TxIsoLevel = 3
)

// Enum value maps for TxIsoLevel.
var (
	TxIsoLevel_name = map[int32]string{
		0: "ISO_LEVEL_READ_UNCOMMITTED",
		1: "ISO_LEVEL_READ_COMMITTED",
		2: "ISO_LEVEL_REPEATABLE_READ",
		3: "ISO_LEVEL_SERIALIZABLE",
	}
	TxIsoLevel_value = map[string]int32{
		"ISO_LEVEL_READ_UNCOMMITTED": 0,
		"ISO_LEVEL_READ_COMMITTED":   1,
		"ISO_LEVEL_REPEATABLE_READ":  2,
		"ISO_LEVEL_SERIALIZABLE":     3,
	}
)

func (x TxIsoLevel) Enum() *TxIsoLevel {
	p := new(TxIsoLevel)
	*p = x
	return p
}

func (x TxIsoLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxIsoLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_store_service_proto_enumTypes[1].Descriptor()
}

func (TxIsoLevel) Type() protoreflect.EnumType {
	return &file_store_service_proto_enumTypes[1]
}

func (x TxIsoLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxIsoLevel.Descriptor instead.
func (TxIsoLevel) EnumDescriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{1}
}

type FileHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Size uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FileHeader) Reset() {
	*x = FileHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHeader) ProtoMessage() {}

func (x *FileHeader) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHeader.ProtoReflect.Descriptor instead.
func (*FileHeader) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileHeader) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FileHeader) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*SetFileRequest_Header
	//	*SetFileRequest_Chunk
	Data isSetFileRequest_Data `protobuf_oneof:"data"`
}

func (x *SetFileRequest) Reset() {
	*x = SetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFileRequest) ProtoMessage() {}

func (x *SetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFileRequest.ProtoReflect.Descriptor instead.
func (*SetFileRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{1}
}

func (m *SetFileRequest) GetData() isSetFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SetFileRequest) GetHeader() *FileHeader {
	if x, ok := x.GetData().(*SetFileRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (x *SetFileRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*SetFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isSetFileRequest_Data interface {
	isSetFileRequest_Data()
}

type SetFileRequest_Header struct {
	Header *FileHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type SetFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*SetFileRequest_Header) isSetFileRequest_Data() {}

func (*SetFileRequest_Chunk) isSetFileRequest_Data() {}

type SetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetFileResponse) Reset() {
	*x = SetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFileResponse) ProtoMessage() {}

func (x *SetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFileResponse.ProtoReflect.Descriptor instead.
func (*SetFileResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{2}
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*GetFileResponse_Chunk
	//	*GetFileResponse_Header
	Data isGetFileResponse_Data `protobuf_oneof:"data"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{4}
}

func (m *GetFileResponse) GetData() isGetFileResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *GetFileResponse) GetChunk() []byte {
	if x, ok := x.GetData().(*GetFileResponse_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *GetFileResponse) GetHeader() *FileHeader {
	if x, ok := x.GetData().(*GetFileResponse_Header); ok {
		return x.Header
	}
	return nil
}

type isGetFileResponse_Data interface {
	isGetFileResponse_Data()
}

type GetFileResponse_Chunk struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3,oneof"`
}

type GetFileResponse_Header struct {
	Header *FileHeader `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

func (*GetFileResponse_Chunk) isGetFileResponse_Data() {}

func (*GetFileResponse_Header) isGetFileResponse_Data() {}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFileRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{6}
}

type BeginTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsoLevel TxIsoLevel `protobuf:"varint,1,opt,name=isoLevel,proto3,enum=store.TxIsoLevel" json:"isoLevel,omitempty"`
}

func (x *BeginTxRequest) Reset() {
	*x = BeginTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginTxRequest) ProtoMessage() {}

func (x *BeginTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginTxRequest.ProtoReflect.Descriptor instead.
func (*BeginTxRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{7}
}

func (x *BeginTxRequest) GetIsoLevel() TxIsoLevel {
	if x != nil {
		return x.IsoLevel
	}
	return TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED
}

type BeginTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BeginTxResponse) Reset() {
	*x = BeginTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginTxResponse) ProtoMessage() {}

func (x *BeginTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginTxResponse.ProtoReflect.Descriptor instead.
func (*BeginTxResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{8}
}

func (x *BeginTxResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommitTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitTxRequest) Reset() {
	*x = CommitTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTxRequest) ProtoMessage() {}

func (x *CommitTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTxRequest.ProtoReflect.Descriptor instead.
func (*CommitTxRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{9}
}

type CommitTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitTxResponse) Reset() {
	*x = CommitTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitTxResponse) ProtoMessage() {}

func (x *CommitTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitTxResponse.ProtoReflect.Descriptor instead.
func (*CommitTxResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{10}
}

type RollbackTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RollbackTxRequest) Reset() {
	*x = RollbackTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackTxRequest) ProtoMessage() {}

func (x *RollbackTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackTxRequest.ProtoReflect.Descriptor instead.
func (*RollbackTxRequest) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{11}
}

type RollbackTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RollbackTxResponse) Reset() {
	*x = RollbackTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackTxResponse) ProtoMessage() {}

func (x *RollbackTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackTxResponse.ProtoReflect.Descriptor instead.
func (*RollbackTxResponse) Descriptor() ([]byte, []int) {
	return file_store_service_proto_rawDescGZIP(), []int{12}
}

var File_store_service_proto protoreflect.FileDescriptor

var file_store_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x32, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x5d, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x0e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x69, 0x73, 0x6f, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x54, 0x78, 0x49, 0x73, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x69, 0x73, 0x6f, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x21, 0x0a, 0x0f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x20, 0x0a, 0x09, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10, 0x80, 0x10, 0x2a, 0x85, 0x01, 0x0a, 0x0a,
	0x54, 0x78, 0x49, 0x73, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x53,
	0x4f, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x55, 0x4e, 0x43,
	0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x53,
	0x4f, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x53, 0x4f, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x53, 0x4f, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x03, 0x32, 0x8a, 0x03, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x12,
	0x3c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x07, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x12, 0x15, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x78, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x78, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x13, 0x5a, 0x11, 0x66, 0x73, 0x5f, 0x66, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_service_proto_rawDescOnce sync.Once
	file_store_service_proto_rawDescData = file_store_service_proto_rawDesc
)

func file_store_service_proto_rawDescGZIP() []byte {
	file_store_service_proto_rawDescOnce.Do(func() {
		file_store_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_service_proto_rawDescData)
	})
	return file_store_service_proto_rawDescData
}

var file_store_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_store_service_proto_goTypes = []interface{}{
	(ChunkSize)(0),             // 0: store.ChunkSize
	(TxIsoLevel)(0),            // 1: store.TxIsoLevel
	(*FileHeader)(nil),         // 2: store.FileHeader
	(*SetFileRequest)(nil),     // 3: store.SetFileRequest
	(*SetFileResponse)(nil),    // 4: store.SetFileResponse
	(*GetFileRequest)(nil),     // 5: store.GetFileRequest
	(*GetFileResponse)(nil),    // 6: store.GetFileResponse
	(*DeleteFileRequest)(nil),  // 7: store.DeleteFileRequest
	(*DeleteFileResponse)(nil), // 8: store.DeleteFileResponse
	(*BeginTxRequest)(nil),     // 9: store.BeginTxRequest
	(*BeginTxResponse)(nil),    // 10: store.BeginTxResponse
	(*CommitTxRequest)(nil),    // 11: store.CommitTxRequest
	(*CommitTxResponse)(nil),   // 12: store.CommitTxResponse
	(*RollbackTxRequest)(nil),  // 13: store.RollbackTxRequest
	(*RollbackTxResponse)(nil), // 14: store.RollbackTxResponse
}
var file_store_service_proto_depIdxs = []int32{
	2,  // 0: store.SetFileRequest.header:type_name -> store.FileHeader
	2,  // 1: store.GetFileResponse.header:type_name -> store.FileHeader
	1,  // 2: store.BeginTxRequest.isoLevel:type_name -> store.TxIsoLevel
	3,  // 3: store.StoreV1.SetFile:input_type -> store.SetFileRequest
	5,  // 4: store.StoreV1.GetFile:input_type -> store.GetFileRequest
	7,  // 5: store.StoreV1.DeleteFile:input_type -> store.DeleteFileRequest
	9,  // 6: store.StoreV1.BeginTx:input_type -> store.BeginTxRequest
	11, // 7: store.StoreV1.CommitTx:input_type -> store.CommitTxRequest
	13, // 8: store.StoreV1.RollbackTx:input_type -> store.RollbackTxRequest
	4,  // 9: store.StoreV1.SetFile:output_type -> store.SetFileResponse
	6,  // 10: store.StoreV1.GetFile:output_type -> store.GetFileResponse
	8,  // 11: store.StoreV1.DeleteFile:output_type -> store.DeleteFileResponse
	10, // 12: store.StoreV1.BeginTx:output_type -> store.BeginTxResponse
	12, // 13: store.StoreV1.CommitTx:output_type -> store.CommitTxResponse
	14, // 14: store.StoreV1.RollbackTx:output_type -> store.RollbackTxResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_store_service_proto_init() }
func file_store_service_proto_init() {
	if File_store_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHeader); i {
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
		file_store_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFileRequest); i {
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
		file_store_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFileResponse); i {
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
		file_store_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
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
		file_store_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
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
		file_store_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
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
		file_store_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
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
		file_store_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginTxRequest); i {
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
		file_store_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginTxResponse); i {
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
		file_store_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTxRequest); i {
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
		file_store_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitTxResponse); i {
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
		file_store_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackTxRequest); i {
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
		file_store_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackTxResponse); i {
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
	file_store_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SetFileRequest_Header)(nil),
		(*SetFileRequest_Chunk)(nil),
	}
	file_store_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetFileResponse_Chunk)(nil),
		(*GetFileResponse_Header)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_service_proto_goTypes,
		DependencyIndexes: file_store_service_proto_depIdxs,
		EnumInfos:         file_store_service_proto_enumTypes,
		MessageInfos:      file_store_service_proto_msgTypes,
	}.Build()
	File_store_service_proto = out.File
	file_store_service_proto_rawDesc = nil
	file_store_service_proto_goTypes = nil
	file_store_service_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: v1/contracts.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LogLevel) Reset() {
	*x = LogLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contracts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLevel) ProtoMessage() {}

func (x *LogLevel) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contracts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLevel.ProtoReflect.Descriptor instead.
func (*LogLevel) Descriptor() ([]byte, []int) {
	return file_v1_contracts_proto_rawDescGZIP(), []int{0}
}

func (x *LogLevel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogLevel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LevelId int32  `protobuf:"varint,2,opt,name=levelId,proto3" json:"levelId,omitempty"`
	Date    string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Source  string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contracts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contracts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_v1_contracts_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Log) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *Log) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Log) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLogsLevelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogsLevels []*LogLevel `protobuf:"bytes,1,rep,name=logs_levels,json=logsLevels,proto3" json:"logs_levels,omitempty"`
}

func (x *GetLogsLevelsResponse) Reset() {
	*x = GetLogsLevelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contracts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsLevelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsLevelsResponse) ProtoMessage() {}

func (x *GetLogsLevelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contracts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsLevelsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsLevelsResponse) Descriptor() ([]byte, []int) {
	return file_v1_contracts_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogsLevelsResponse) GetLogsLevels() []*LogLevel {
	if x != nil {
		return x.LogsLevels
	}
	return nil
}

type GetLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelIds string `protobuf:"bytes,1,opt,name=level_ids,json=levelIds,proto3" json:"level_ids,omitempty"`
	Source   string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	DateFrom string `protobuf:"bytes,4,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo   string `protobuf:"bytes,5,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *GetLogsRequest) Reset() {
	*x = GetLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contracts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRequest) ProtoMessage() {}

func (x *GetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contracts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRequest.ProtoReflect.Descriptor instead.
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_contracts_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogsRequest) GetLevelIds() string {
	if x != nil {
		return x.LevelIds
	}
	return ""
}

func (x *GetLogsRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetLogsRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetLogsRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *GetLogsRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type GetLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetLogsResponse) Reset() {
	*x = GetLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_contracts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsResponse) ProtoMessage() {}

func (x *GetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_contracts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return file_v1_contracts_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogsResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_v1_contracts_proto protoreflect.FileDescriptor

var file_v1_contracts_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x6c,
	0x6f, 0x67, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x32, 0xbb, 0x01, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x60, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x51, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_contracts_proto_rawDescOnce sync.Once
	file_v1_contracts_proto_rawDescData = file_v1_contracts_proto_rawDesc
)

func file_v1_contracts_proto_rawDescGZIP() []byte {
	file_v1_contracts_proto_rawDescOnce.Do(func() {
		file_v1_contracts_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_contracts_proto_rawDescData)
	})
	return file_v1_contracts_proto_rawDescData
}

var file_v1_contracts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_contracts_proto_goTypes = []interface{}{
	(*LogLevel)(nil),              // 0: service.v1.LogLevel
	(*Log)(nil),                   // 1: service.v1.Log
	(*GetLogsLevelsResponse)(nil), // 2: service.v1.GetLogsLevelsResponse
	(*GetLogsRequest)(nil),        // 3: service.v1.GetLogsRequest
	(*GetLogsResponse)(nil),       // 4: service.v1.GetLogsResponse
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_v1_contracts_proto_depIdxs = []int32{
	0, // 0: service.v1.GetLogsLevelsResponse.logs_levels:type_name -> service.v1.LogLevel
	1, // 1: service.v1.GetLogsResponse.logs:type_name -> service.v1.Log
	5, // 2: service.v1.Logs.GetLogsLevels:input_type -> google.protobuf.Empty
	3, // 3: service.v1.Logs.GetLogs:input_type -> service.v1.GetLogsRequest
	2, // 4: service.v1.Logs.GetLogsLevels:output_type -> service.v1.GetLogsLevelsResponse
	4, // 5: service.v1.Logs.GetLogs:output_type -> service.v1.GetLogsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_contracts_proto_init() }
func file_v1_contracts_proto_init() {
	if File_v1_contracts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_contracts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLevel); i {
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
		file_v1_contracts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_v1_contracts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsLevelsResponse); i {
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
		file_v1_contracts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRequest); i {
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
		file_v1_contracts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsResponse); i {
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
			RawDescriptor: file_v1_contracts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_contracts_proto_goTypes,
		DependencyIndexes: file_v1_contracts_proto_depIdxs,
		MessageInfos:      file_v1_contracts_proto_msgTypes,
	}.Build()
	File_v1_contracts_proto = out.File
	file_v1_contracts_proto_rawDesc = nil
	file_v1_contracts_proto_goTypes = nil
	file_v1_contracts_proto_depIdxs = nil
}

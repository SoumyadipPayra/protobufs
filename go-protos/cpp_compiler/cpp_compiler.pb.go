// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: protos/cpp_compiler/cpp_compiler.proto

package cpp_compiler

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_protos_cpp_compiler_cpp_compiler_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_protos_cpp_compiler_cpp_compiler_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CompileAndRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileData []byte `protobuf:"bytes,3,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
}

func (x *CompileAndRunRequest) Reset() {
	*x = CompileAndRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileAndRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileAndRunRequest) ProtoMessage() {}

func (x *CompileAndRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileAndRunRequest.ProtoReflect.Descriptor instead.
func (*CompileAndRunRequest) Descriptor() ([]byte, []int) {
	return file_protos_cpp_compiler_cpp_compiler_proto_rawDescGZIP(), []int{2}
}

func (x *CompileAndRunRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CompileAndRunRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *CompileAndRunRequest) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

type CompileAndRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []byte `protobuf:"bytes,2,opt,name=logs,proto3" json:"logs,omitempty"`
}

func (x *CompileAndRunResponse) Reset() {
	*x = CompileAndRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileAndRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileAndRunResponse) ProtoMessage() {}

func (x *CompileAndRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileAndRunResponse.ProtoReflect.Descriptor instead.
func (*CompileAndRunResponse) Descriptor() ([]byte, []int) {
	return file_protos_cpp_compiler_cpp_compiler_proto_rawDescGZIP(), []int{3}
}

func (x *CompileAndRunResponse) GetLogs() []byte {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_protos_cpp_compiler_cpp_compiler_proto protoreflect.FileDescriptor

var file_protos_cpp_compiler_cpp_compiler_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x70, 0x70, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x6d, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x32, 0xea, 0x01, 0x0a, 0x0b, 0x43, 0x50, 0x50, 0x43, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67,
	0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x70, 0x70, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x70, 0x70,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x41, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x70, 0x70, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x41, 0x6e,
	0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x70,
	0x70, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x41, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x6f, 0x75, 0x6d, 0x79, 0x61, 0x64, 0x69, 0x70, 0x50, 0x61, 0x79, 0x72, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x63, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_cpp_compiler_cpp_compiler_proto_rawDescOnce sync.Once
	file_protos_cpp_compiler_cpp_compiler_proto_rawDescData = file_protos_cpp_compiler_cpp_compiler_proto_rawDesc
)

func file_protos_cpp_compiler_cpp_compiler_proto_rawDescGZIP() []byte {
	file_protos_cpp_compiler_cpp_compiler_proto_rawDescOnce.Do(func() {
		file_protos_cpp_compiler_cpp_compiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_cpp_compiler_cpp_compiler_proto_rawDescData)
	})
	return file_protos_cpp_compiler_cpp_compiler_proto_rawDescData
}

var file_protos_cpp_compiler_cpp_compiler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_cpp_compiler_cpp_compiler_proto_goTypes = []interface{}{
	(*PingRequest)(nil),           // 0: protobufs.proto.cppcompiler.PingRequest
	(*PingResponse)(nil),          // 1: protobufs.proto.cppcompiler.PingResponse
	(*CompileAndRunRequest)(nil),  // 2: protobufs.proto.cppcompiler.CompileAndRunRequest
	(*CompileAndRunResponse)(nil), // 3: protobufs.proto.cppcompiler.CompileAndRunResponse
}
var file_protos_cpp_compiler_cpp_compiler_proto_depIdxs = []int32{
	0, // 0: protobufs.proto.cppcompiler.CPPCompiler.PingPong:input_type -> protobufs.proto.cppcompiler.PingRequest
	2, // 1: protobufs.proto.cppcompiler.CPPCompiler.CompileAndRun:input_type -> protobufs.proto.cppcompiler.CompileAndRunRequest
	1, // 2: protobufs.proto.cppcompiler.CPPCompiler.PingPong:output_type -> protobufs.proto.cppcompiler.PingResponse
	3, // 3: protobufs.proto.cppcompiler.CPPCompiler.CompileAndRun:output_type -> protobufs.proto.cppcompiler.CompileAndRunResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_cpp_compiler_cpp_compiler_proto_init() }
func file_protos_cpp_compiler_cpp_compiler_proto_init() {
	if File_protos_cpp_compiler_cpp_compiler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileAndRunRequest); i {
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
		file_protos_cpp_compiler_cpp_compiler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileAndRunResponse); i {
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
			RawDescriptor: file_protos_cpp_compiler_cpp_compiler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_cpp_compiler_cpp_compiler_proto_goTypes,
		DependencyIndexes: file_protos_cpp_compiler_cpp_compiler_proto_depIdxs,
		MessageInfos:      file_protos_cpp_compiler_cpp_compiler_proto_msgTypes,
	}.Build()
	File_protos_cpp_compiler_cpp_compiler_proto = out.File
	file_protos_cpp_compiler_cpp_compiler_proto_rawDesc = nil
	file_protos_cpp_compiler_cpp_compiler_proto_goTypes = nil
	file_protos_cpp_compiler_cpp_compiler_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m3.proto

package test_a_2

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

type M3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M3) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_a_2_m3_proto_messageTypes[0].MessageOf(m)
}
func (m *M3) Reset()         { *m = M3{} }
func (m *M3) String() string { return protoimpl.X.MessageStringOf(m) }
func (*M3) ProtoMessage()    {}

// Deprecated: Use M3.ProtoReflect.Type instead.
func (*M3) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_2_m3_proto_rawdesc_gzipped, []int{0}
}

var xxx_File_imports_test_a_2_m3_proto_rawdesc = []byte{
	// 126 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x32, 0x2f, 0x6d, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x33, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_imports_test_a_2_m3_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_imports_test_a_2_m3_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_imports_test_a_2_m3_proto protoreflect.FileDescriptor

var xxx_File_imports_test_a_2_m3_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_imports_test_a_2_m3_proto_goTypes = []interface{}{
	(*M3)(nil), // 0: test.a.M3
}
var xxx_File_imports_test_a_2_m3_proto_depIdxs = []int32{}

func init() { xxx_File_imports_test_a_2_m3_proto_init() }
func xxx_File_imports_test_a_2_m3_proto_init() {
	if File_imports_test_a_2_m3_proto != nil {
		return
	}
	File_imports_test_a_2_m3_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_imports_test_a_2_m3_proto_rawdesc,
		GoTypes:            xxx_File_imports_test_a_2_m3_proto_goTypes,
		DependencyIndexes:  xxx_File_imports_test_a_2_m3_proto_depIdxs,
		MessageOutputTypes: xxx_File_imports_test_a_2_m3_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_imports_test_a_2_m3_proto_goTypes = nil
	xxx_File_imports_test_a_2_m3_proto_depIdxs = nil
}

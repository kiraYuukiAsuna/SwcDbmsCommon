// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: Service/Service.proto

package service

import (
	_ "DBMS/SwcDbmsCommon/Generated/go/proto/message"
	request "DBMS/SwcDbmsCommon/Generated/go/proto/request"
	response "DBMS/SwcDbmsCommon/Generated/go/proto/response"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

var File_Service_Service_proto protoreflect.FileDescriptor

var file_Service_Service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xee, 0x16, 0x0a, 0x04, 0x44, 0x42, 0x4d, 0x53, 0x12, 0x43,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x20, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x77,
	0x63, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x77, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x77, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x77, 0x63, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x77, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x77, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x77, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x77, 0x63,
	0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x77,
	0x63, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63,
	0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x4e, 0x6f,
	0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x77, 0x63, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63,
	0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x46, 0x75, 0x6c, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x82, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x53, 0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x77, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77,
	0x63, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x46, 0x75,
	0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x64, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x44, 0x42, 0x4d, 0x53, 0x2f, 0x53,
	0x77, 0x63, 0x44, 0x62, 0x6d, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_Service_Service_proto_goTypes = []interface{}{
	(*request.CreateUserRequest)(nil),                        // 0: proto.CreateUserRequest
	(*request.DeleteUserRequest)(nil),                        // 1: proto.DeleteUserRequest
	(*request.UpdateUserRequest)(nil),                        // 2: proto.UpdateUserRequest
	(*request.GetUserRequest)(nil),                           // 3: proto.GetUserRequest
	(*request.GetAllUserRequest)(nil),                        // 4: proto.GetAllUserRequest
	(*request.UserLoginRequest)(nil),                         // 5: proto.UserLoginRequest
	(*request.UserLogoutRequest)(nil),                        // 6: proto.UserLogoutRequest
	(*request.UserOnlineHeartBeatNotification)(nil),          // 7: proto.UserOnlineHeartBeatNotification
	(*request.GetUserPermissionGroupRequest)(nil),            // 8: proto.GetUserPermissionGroupRequest
	(*request.GetPermissionGroupRequest)(nil),                // 9: proto.GetPermissionGroupRequest
	(*request.GetAllPermissionGroupRequest)(nil),             // 10: proto.GetAllPermissionGroupRequest
	(*request.ChangeUserPermissionGroupRequest)(nil),         // 11: proto.ChangeUserPermissionGroupRequest
	(*request.CreateProjectRequest)(nil),                     // 12: proto.CreateProjectRequest
	(*request.DeleteProjectRequest)(nil),                     // 13: proto.DeleteProjectRequest
	(*request.UpdateProjectRequest)(nil),                     // 14: proto.UpdateProjectRequest
	(*request.GetProjectRequest)(nil),                        // 15: proto.GetProjectRequest
	(*request.GetAllProjectRequest)(nil),                     // 16: proto.GetAllProjectRequest
	(*request.CreateSwcRequest)(nil),                         // 17: proto.CreateSwcRequest
	(*request.DeleteSwcRequest)(nil),                         // 18: proto.DeleteSwcRequest
	(*request.UpdateSwcRequest)(nil),                         // 19: proto.UpdateSwcRequest
	(*request.GetSwcMetaInfoRequest)(nil),                    // 20: proto.GetSwcMetaInfoRequest
	(*request.GetAllSwcMetaInfoRequest)(nil),                 // 21: proto.GetAllSwcMetaInfoRequest
	(*request.CreateSwcNodeDataRequest)(nil),                 // 22: proto.CreateSwcNodeDataRequest
	(*request.DeleteSwcNodeDataRequest)(nil),                 // 23: proto.DeleteSwcNodeDataRequest
	(*request.UpdateSwcNodeDataRequest)(nil),                 // 24: proto.UpdateSwcNodeDataRequest
	(*request.GetSwcNodeDataRequest)(nil),                    // 25: proto.GetSwcNodeDataRequest
	(*request.GetSwcFullNodeDataRequest)(nil),                // 26: proto.GetSwcFullNodeDataRequest
	(*request.GetSwcNodeDataListByTimeAndUserRequest)(nil),   // 27: proto.GetSwcNodeDataListByTimeAndUserRequest
	(*request.BackupFullDatabaseRequest)(nil),                // 28: proto.BackupFullDatabaseRequest
	(*request.CreateDailyStatisticsRequest)(nil),             // 29: proto.CreateDailyStatisticsRequest
	(*request.DeleteDailyStatisticsRequest)(nil),             // 30: proto.DeleteDailyStatisticsRequest
	(*request.UpdateDailyStatisticsRequest)(nil),             // 31: proto.UpdateDailyStatisticsRequest
	(*request.GetDailyStatisticsRequest)(nil),                // 32: proto.GetDailyStatisticsRequest
	(*request.GetAllDailyStatisticsRequest)(nil),             // 33: proto.GetAllDailyStatisticsRequest
	(*response.CreateUserResponse)(nil),                      // 34: proto.CreateUserResponse
	(*response.DeleteUserResponse)(nil),                      // 35: proto.DeleteUserResponse
	(*response.UpdateUserResponse)(nil),                      // 36: proto.UpdateUserResponse
	(*response.GetUserResponse)(nil),                         // 37: proto.GetUserResponse
	(*response.GetAllUserResponse)(nil),                      // 38: proto.GetAllUserResponse
	(*response.UserLoginResponse)(nil),                       // 39: proto.UserLoginResponse
	(*response.UserLogoutResponse)(nil),                      // 40: proto.UserLogoutResponse
	(*response.UserOnlineHeartBeatResponse)(nil),             // 41: proto.UserOnlineHeartBeatResponse
	(*response.GetUserPermissionGroupResponse)(nil),          // 42: proto.GetUserPermissionGroupResponse
	(*response.GetPermissionGroupResponse)(nil),              // 43: proto.GetPermissionGroupResponse
	(*response.GetAllPermissionGroupResponse)(nil),           // 44: proto.GetAllPermissionGroupResponse
	(*response.ChangeUserPermissionGroupResponse)(nil),       // 45: proto.ChangeUserPermissionGroupResponse
	(*response.CreateProjectResponse)(nil),                   // 46: proto.CreateProjectResponse
	(*response.DeleteProjectResponse)(nil),                   // 47: proto.DeleteProjectResponse
	(*response.UpdateProjectResponse)(nil),                   // 48: proto.UpdateProjectResponse
	(*response.GetProjectResponse)(nil),                      // 49: proto.GetProjectResponse
	(*response.GetAllProjectResponse)(nil),                   // 50: proto.GetAllProjectResponse
	(*response.CreateSwcResponse)(nil),                       // 51: proto.CreateSwcResponse
	(*response.DeleteSwcResponse)(nil),                       // 52: proto.DeleteSwcResponse
	(*response.UpdateSwcResponse)(nil),                       // 53: proto.UpdateSwcResponse
	(*response.GetSwcMetaInfoResponse)(nil),                  // 54: proto.GetSwcMetaInfoResponse
	(*response.GetAllSwcMetaInfoResponse)(nil),               // 55: proto.GetAllSwcMetaInfoResponse
	(*response.CreateSwcNodeDataResponse)(nil),               // 56: proto.CreateSwcNodeDataResponse
	(*response.DeleteSwcNodeDataResponse)(nil),               // 57: proto.DeleteSwcNodeDataResponse
	(*response.UpdateSwcNodeDataResponse)(nil),               // 58: proto.UpdateSwcNodeDataResponse
	(*response.GetSwcNodeDataResponse)(nil),                  // 59: proto.GetSwcNodeDataResponse
	(*response.GetSwcFullNodeDataResponse)(nil),              // 60: proto.GetSwcFullNodeDataResponse
	(*response.GetSwcNodeDataListByTimeAndUserResponse)(nil), // 61: proto.GetSwcNodeDataListByTimeAndUserResponse
	(*response.BackupFullDatabaseResponse)(nil),              // 62: proto.BackupFullDatabaseResponse
	(*response.CreateDailyStatisticsResponse)(nil),           // 63: proto.CreateDailyStatisticsResponse
	(*response.DeleteDailyStatisticsResponse)(nil),           // 64: proto.DeleteDailyStatisticsResponse
	(*response.UpdateDailyStatisticsResponse)(nil),           // 65: proto.UpdateDailyStatisticsResponse
	(*response.GetDailyStatisticsResponse)(nil),              // 66: proto.GetDailyStatisticsResponse
	(*response.GetAllDailyStatisticsResponse)(nil),           // 67: proto.GetAllDailyStatisticsResponse
}
var file_Service_Service_proto_depIdxs = []int32{
	0,  // 0: proto.DBMS.CreateUser:input_type -> proto.CreateUserRequest
	1,  // 1: proto.DBMS.DeleteUser:input_type -> proto.DeleteUserRequest
	2,  // 2: proto.DBMS.UpdateUser:input_type -> proto.UpdateUserRequest
	3,  // 3: proto.DBMS.GetUser:input_type -> proto.GetUserRequest
	4,  // 4: proto.DBMS.GetAllUser:input_type -> proto.GetAllUserRequest
	5,  // 5: proto.DBMS.UserLogin:input_type -> proto.UserLoginRequest
	6,  // 6: proto.DBMS.UserLogout:input_type -> proto.UserLogoutRequest
	7,  // 7: proto.DBMS.UserOnlineHeartBeatNotifications:input_type -> proto.UserOnlineHeartBeatNotification
	8,  // 8: proto.DBMS.GetUserPermissionGroup:input_type -> proto.GetUserPermissionGroupRequest
	9,  // 9: proto.DBMS.GetPermissionGroup:input_type -> proto.GetPermissionGroupRequest
	10, // 10: proto.DBMS.GetAllPermissionGroup:input_type -> proto.GetAllPermissionGroupRequest
	11, // 11: proto.DBMS.ChangeUserPermissionGroup:input_type -> proto.ChangeUserPermissionGroupRequest
	12, // 12: proto.DBMS.CreateProject:input_type -> proto.CreateProjectRequest
	13, // 13: proto.DBMS.DeleteProject:input_type -> proto.DeleteProjectRequest
	14, // 14: proto.DBMS.UpdateProject:input_type -> proto.UpdateProjectRequest
	15, // 15: proto.DBMS.GetProject:input_type -> proto.GetProjectRequest
	16, // 16: proto.DBMS.GetAllProject:input_type -> proto.GetAllProjectRequest
	17, // 17: proto.DBMS.CreateSwc:input_type -> proto.CreateSwcRequest
	18, // 18: proto.DBMS.DeleteSwc:input_type -> proto.DeleteSwcRequest
	19, // 19: proto.DBMS.UpdateSwc:input_type -> proto.UpdateSwcRequest
	20, // 20: proto.DBMS.GetSwcMetaInfo:input_type -> proto.GetSwcMetaInfoRequest
	21, // 21: proto.DBMS.GetAllSwcMetaInfo:input_type -> proto.GetAllSwcMetaInfoRequest
	22, // 22: proto.DBMS.CreateSwcNodeData:input_type -> proto.CreateSwcNodeDataRequest
	23, // 23: proto.DBMS.DeleteSwcNodeData:input_type -> proto.DeleteSwcNodeDataRequest
	24, // 24: proto.DBMS.UpdateSwcNodeData:input_type -> proto.UpdateSwcNodeDataRequest
	25, // 25: proto.DBMS.GetSwcNodeData:input_type -> proto.GetSwcNodeDataRequest
	26, // 26: proto.DBMS.GetSwcFullNodeData:input_type -> proto.GetSwcFullNodeDataRequest
	27, // 27: proto.DBMS.GetSwcNodeDataListByTimeAndUser:input_type -> proto.GetSwcNodeDataListByTimeAndUserRequest
	28, // 28: proto.DBMS.BackupFullDatabase:input_type -> proto.BackupFullDatabaseRequest
	29, // 29: proto.DBMS.CreateDailyStatistics:input_type -> proto.CreateDailyStatisticsRequest
	30, // 30: proto.DBMS.DeleteDailyStatistics:input_type -> proto.DeleteDailyStatisticsRequest
	31, // 31: proto.DBMS.UpdateDailyStatistics:input_type -> proto.UpdateDailyStatisticsRequest
	32, // 32: proto.DBMS.GetDailyStatistics:input_type -> proto.GetDailyStatisticsRequest
	33, // 33: proto.DBMS.GetAllDailyStatistics:input_type -> proto.GetAllDailyStatisticsRequest
	34, // 34: proto.DBMS.CreateUser:output_type -> proto.CreateUserResponse
	35, // 35: proto.DBMS.DeleteUser:output_type -> proto.DeleteUserResponse
	36, // 36: proto.DBMS.UpdateUser:output_type -> proto.UpdateUserResponse
	37, // 37: proto.DBMS.GetUser:output_type -> proto.GetUserResponse
	38, // 38: proto.DBMS.GetAllUser:output_type -> proto.GetAllUserResponse
	39, // 39: proto.DBMS.UserLogin:output_type -> proto.UserLoginResponse
	40, // 40: proto.DBMS.UserLogout:output_type -> proto.UserLogoutResponse
	41, // 41: proto.DBMS.UserOnlineHeartBeatNotifications:output_type -> proto.UserOnlineHeartBeatResponse
	42, // 42: proto.DBMS.GetUserPermissionGroup:output_type -> proto.GetUserPermissionGroupResponse
	43, // 43: proto.DBMS.GetPermissionGroup:output_type -> proto.GetPermissionGroupResponse
	44, // 44: proto.DBMS.GetAllPermissionGroup:output_type -> proto.GetAllPermissionGroupResponse
	45, // 45: proto.DBMS.ChangeUserPermissionGroup:output_type -> proto.ChangeUserPermissionGroupResponse
	46, // 46: proto.DBMS.CreateProject:output_type -> proto.CreateProjectResponse
	47, // 47: proto.DBMS.DeleteProject:output_type -> proto.DeleteProjectResponse
	48, // 48: proto.DBMS.UpdateProject:output_type -> proto.UpdateProjectResponse
	49, // 49: proto.DBMS.GetProject:output_type -> proto.GetProjectResponse
	50, // 50: proto.DBMS.GetAllProject:output_type -> proto.GetAllProjectResponse
	51, // 51: proto.DBMS.CreateSwc:output_type -> proto.CreateSwcResponse
	52, // 52: proto.DBMS.DeleteSwc:output_type -> proto.DeleteSwcResponse
	53, // 53: proto.DBMS.UpdateSwc:output_type -> proto.UpdateSwcResponse
	54, // 54: proto.DBMS.GetSwcMetaInfo:output_type -> proto.GetSwcMetaInfoResponse
	55, // 55: proto.DBMS.GetAllSwcMetaInfo:output_type -> proto.GetAllSwcMetaInfoResponse
	56, // 56: proto.DBMS.CreateSwcNodeData:output_type -> proto.CreateSwcNodeDataResponse
	57, // 57: proto.DBMS.DeleteSwcNodeData:output_type -> proto.DeleteSwcNodeDataResponse
	58, // 58: proto.DBMS.UpdateSwcNodeData:output_type -> proto.UpdateSwcNodeDataResponse
	59, // 59: proto.DBMS.GetSwcNodeData:output_type -> proto.GetSwcNodeDataResponse
	60, // 60: proto.DBMS.GetSwcFullNodeData:output_type -> proto.GetSwcFullNodeDataResponse
	61, // 61: proto.DBMS.GetSwcNodeDataListByTimeAndUser:output_type -> proto.GetSwcNodeDataListByTimeAndUserResponse
	62, // 62: proto.DBMS.BackupFullDatabase:output_type -> proto.BackupFullDatabaseResponse
	63, // 63: proto.DBMS.CreateDailyStatistics:output_type -> proto.CreateDailyStatisticsResponse
	64, // 64: proto.DBMS.DeleteDailyStatistics:output_type -> proto.DeleteDailyStatisticsResponse
	65, // 65: proto.DBMS.UpdateDailyStatistics:output_type -> proto.UpdateDailyStatisticsResponse
	66, // 66: proto.DBMS.GetDailyStatistics:output_type -> proto.GetDailyStatisticsResponse
	67, // 67: proto.DBMS.GetAllDailyStatistics:output_type -> proto.GetAllDailyStatisticsResponse
	34, // [34:68] is the sub-list for method output_type
	0,  // [0:34] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_Service_Service_proto_init() }
func file_Service_Service_proto_init() {
	if File_Service_Service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Service_Service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Service_Service_proto_goTypes,
		DependencyIndexes: file_Service_Service_proto_depIdxs,
	}.Build()
	File_Service_Service_proto = out.File
	file_Service_Service_proto_rawDesc = nil
	file_Service_Service_proto_goTypes = nil
	file_Service_Service_proto_depIdxs = nil
}
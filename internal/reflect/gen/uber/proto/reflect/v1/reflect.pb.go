// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uber/proto/reflect/v1/reflect.proto

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package uber.proto.reflect.v1 contains structures to represent information
// about Protobuf definitions on a per-package basis. This differs from
// the structures in descriptor.proto which represent Protobuf definitions
// on a per-file basis.
//
// This package does not represent all information about a Protobuf package
// at this time, be added in the future.
//
// A non-comprehensive list of excluded items:
//
// - Source code information.
// - All options.
// - Reserved ranges and names.
// - Message field default values.
// - Message field oneof indexes.
// - Message field JSON names.
// - Anything to do with extensions.
//
// Excluded items that should not be relevant at a package level:
//
// - Public vs non-public dependencies are all grouped into the same list.
// - Syntax information.
// - Fields vs extenstions.
// - Extension ranges.

package reflectv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Label is the label of the message field.
//
// The numbers match the FieldDescriptorProto.Label numbers.
//
// Note that a map field will come up as repeated, with type TYPE_MESSAGE,
// and the type_name will be *Entry. We may parse this from FieldDescriptorProtos
// in the future and add a LABEL_MAP.
type MessageField_Label int32

const (
	MessageField_LABEL_INVALID  MessageField_Label = 0
	MessageField_LABEL_OPTIONAL MessageField_Label = 1
	MessageField_LABEL_REQUIRED MessageField_Label = 2
	MessageField_LABEL_REPEATED MessageField_Label = 3
)

var MessageField_Label_name = map[int32]string{
	0: "LABEL_INVALID",
	1: "LABEL_OPTIONAL",
	2: "LABEL_REQUIRED",
	3: "LABEL_REPEATED",
}

var MessageField_Label_value = map[string]int32{
	"LABEL_INVALID":  0,
	"LABEL_OPTIONAL": 1,
	"LABEL_REQUIRED": 2,
	"LABEL_REPEATED": 3,
}

func (x MessageField_Label) String() string {
	return proto.EnumName(MessageField_Label_name, int32(x))
}

func (MessageField_Label) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{5, 0}
}

// Type is the type of the message field.
//
// The numbers match the FieldDescriptorProto.Type numbers.
type MessageField_Type int32

const (
	MessageField_TYPE_INVALID  MessageField_Type = 0
	MessageField_TYPE_DOUBLE   MessageField_Type = 1
	MessageField_TYPE_FLOAT    MessageField_Type = 2
	MessageField_TYPE_INT64    MessageField_Type = 3
	MessageField_TYPE_UINT64   MessageField_Type = 4
	MessageField_TYPE_INT32    MessageField_Type = 5
	MessageField_TYPE_FIXED64  MessageField_Type = 6
	MessageField_TYPE_FIXED32  MessageField_Type = 7
	MessageField_TYPE_BOOL     MessageField_Type = 8
	MessageField_TYPE_STRING   MessageField_Type = 9
	MessageField_TYPE_GROUP    MessageField_Type = 10
	MessageField_TYPE_MESSAGE  MessageField_Type = 11
	MessageField_TYPE_BYTES    MessageField_Type = 12
	MessageField_TYPE_UINT32   MessageField_Type = 13
	MessageField_TYPE_ENUM     MessageField_Type = 14
	MessageField_TYPE_SFIXED32 MessageField_Type = 15
	MessageField_TYPE_SFIXED64 MessageField_Type = 16
	MessageField_TYPE_SINT32   MessageField_Type = 17
	MessageField_TYPE_SINT64   MessageField_Type = 18
)

var MessageField_Type_name = map[int32]string{
	0:  "TYPE_INVALID",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}

var MessageField_Type_value = map[string]int32{
	"TYPE_INVALID":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x MessageField_Type) String() string {
	return proto.EnumName(MessageField_Type_name, int32(x))
}

func (MessageField_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{5, 1}
}

// PackageSet is a set of Packages.
type PackageSet struct {
	// packages contains the set of Packages.
	//
	// These will all be unique, and can be referenced by their name field.
	// These will also be sorted by name.
	Packages             []*Package `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PackageSet) Reset()         { *m = PackageSet{} }
func (m *PackageSet) String() string { return proto.CompactTextString(m) }
func (*PackageSet) ProtoMessage()    {}
func (*PackageSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{0}
}

func (m *PackageSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageSet.Unmarshal(m, b)
}
func (m *PackageSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageSet.Marshal(b, m, deterministic)
}
func (m *PackageSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageSet.Merge(m, src)
}
func (m *PackageSet) XXX_Size() int {
	return xxx_messageInfo_PackageSet.Size(m)
}
func (m *PackageSet) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageSet.DiscardUnknown(m)
}

var xxx_messageInfo_PackageSet proto.InternalMessageInfo

func (m *PackageSet) GetPackages() []*Package {
	if m != nil {
		return m.Packages
	}
	return nil
}

// Package describes a Protobuf package, constructed from a set of Protobuf
// files that have the same "package" value.
type Package struct {
	// name is the fully-qualified name of the package.
	//
	// The package name can be thought of as the unique identifier for a
	// Package message, and is used as a reference.
	//
	// This does not include the prefix '.' found in the traditional package
	// fully-qualified name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// dependency_names contains the names of this packages' dependencies.
	//
	// These will be sorted.
	// If you have a set of Package messages, this will correspond to the name
	// field of other Packages in the set.
	DependencyNames []string `protobuf:"bytes,2,rep,name=dependency_names,json=dependencyNames,proto3" json:"dependency_names,omitempty"`
	// enums contains the top-level enums within this package.
	//
	// These will be sorted by name.
	// Nested enums will be within Messages.
	Enums []*Enum `protobuf:"bytes,3,rep,name=enums,proto3" json:"enums,omitempty"`
	// messages contains the top-level messages within this package.
	//
	// These will be sorted by name.
	// Nested messages will be within Messages.
	Messages []*Message `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	// services contains the services within this package.
	//
	// These will be sorted by name.
	Services             []*Service `protobuf:"bytes,5,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{1}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetDependencyNames() []string {
	if m != nil {
		return m.DependencyNames
	}
	return nil
}

func (m *Package) GetEnums() []*Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Package) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Package) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// Enum describes a Protobuf enum.
type Enum struct {
	// name is the name of the enum.
	//
	// If this is a nested enum, this will not contain the name of the
	// encapsulating message.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// enum_values contains the enum values.
	//
	// These will be sorted by number.
	EnumValues           []*EnumValue `protobuf:"bytes,2,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Enum) Reset()         { *m = Enum{} }
func (m *Enum) String() string { return proto.CompactTextString(m) }
func (*Enum) ProtoMessage()    {}
func (*Enum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{2}
}

func (m *Enum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enum.Unmarshal(m, b)
}
func (m *Enum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enum.Marshal(b, m, deterministic)
}
func (m *Enum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enum.Merge(m, src)
}
func (m *Enum) XXX_Size() int {
	return xxx_messageInfo_Enum.Size(m)
}
func (m *Enum) XXX_DiscardUnknown() {
	xxx_messageInfo_Enum.DiscardUnknown(m)
}

var xxx_messageInfo_Enum proto.InternalMessageInfo

func (m *Enum) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Enum) GetEnumValues() []*EnumValue {
	if m != nil {
		return m.EnumValues
	}
	return nil
}

// EnumValue describes a Protobuf enum value.
type EnumValue struct {
	// name contains the value name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// number contains the value number.
	Number               int32    `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{3}
}

func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (m *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(m, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnumValue) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// Message describes a Protobuf message.
type Message struct {
	// name is the name of the message.
	//
	// If this is a nested message, this will not contain the name of the
	// encapsulating message.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// message_fields contains the message fields.
	//
	// This does not include any extended fields.
	// These will be sorted by number.
	MessageFields []*MessageField `protobuf:"bytes,2,rep,name=message_fields,json=messageFields,proto3" json:"message_fields,omitempty"`
	// message_oneofs contains the oneofs.
	//
	// These will be sorted by name.
	MessageOneofs []*MessageOneof `protobuf:"bytes,3,rep,name=message_oneofs,json=messageOneofs,proto3" json:"message_oneofs,omitempty"`
	// nested_message contains the messages directly nested on this message.
	//
	// These will be sorted by name.
	NestedMessages []*Message `protobuf:"bytes,4,rep,name=nested_messages,json=nestedMessages,proto3" json:"nested_messages,omitempty"`
	// nested_enums contains the enums directed nested on this message.
	//
	// These will be sorted by name.
	NestedEnums          []*Enum  `protobuf:"bytes,5,rep,name=nested_enums,json=nestedEnums,proto3" json:"nested_enums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{4}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetMessageFields() []*MessageField {
	if m != nil {
		return m.MessageFields
	}
	return nil
}

func (m *Message) GetMessageOneofs() []*MessageOneof {
	if m != nil {
		return m.MessageOneofs
	}
	return nil
}

func (m *Message) GetNestedMessages() []*Message {
	if m != nil {
		return m.NestedMessages
	}
	return nil
}

func (m *Message) GetNestedEnums() []*Enum {
	if m != nil {
		return m.NestedEnums
	}
	return nil
}

// MessageField describes a Protobuf message field.
type MessageField struct {
	// name is the name of the message field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// number is the number of the message field.
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// label is the label of the message field.
	Label MessageField_Label `protobuf:"varint,3,opt,name=label,proto3,enum=uber.proto.reflect.v1.MessageField_Label" json:"label,omitempty"`
	// type is the type of the message field.
	Type MessageField_Type `protobuf:"varint,4,opt,name=type,proto3,enum=uber.proto.reflect.v1.MessageField_Type" json:"type,omitempty"`
	// type_name is the fully-qualified name of the type for message and enum
	// fields.
	//
	// This does not include the prefix '.' found in the traditional package
	// fully-qualified name. If this is a nested message, the parent messages
	// will be part of the name.
	TypeName             string   `protobuf:"bytes,5,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageField) Reset()         { *m = MessageField{} }
func (m *MessageField) String() string { return proto.CompactTextString(m) }
func (*MessageField) ProtoMessage()    {}
func (*MessageField) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{5}
}

func (m *MessageField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageField.Unmarshal(m, b)
}
func (m *MessageField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageField.Marshal(b, m, deterministic)
}
func (m *MessageField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageField.Merge(m, src)
}
func (m *MessageField) XXX_Size() int {
	return xxx_messageInfo_MessageField.Size(m)
}
func (m *MessageField) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageField.DiscardUnknown(m)
}

var xxx_messageInfo_MessageField proto.InternalMessageInfo

func (m *MessageField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageField) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *MessageField) GetLabel() MessageField_Label {
	if m != nil {
		return m.Label
	}
	return MessageField_LABEL_INVALID
}

func (m *MessageField) GetType() MessageField_Type {
	if m != nil {
		return m.Type
	}
	return MessageField_TYPE_INVALID
}

func (m *MessageField) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

// MessageOneof describes a Protobuf message oneof.
type MessageOneof struct {
	// name is the name of the message oneof.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// field_numbers are the field numbers in the encapsulating type that make up
	// this message oneof.
	//
	// This will be sorted.
	FieldNumbers         []int32  `protobuf:"varint,2,rep,packed,name=field_numbers,json=fieldNumbers,proto3" json:"field_numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOneof) Reset()         { *m = MessageOneof{} }
func (m *MessageOneof) String() string { return proto.CompactTextString(m) }
func (*MessageOneof) ProtoMessage()    {}
func (*MessageOneof) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{6}
}

func (m *MessageOneof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOneof.Unmarshal(m, b)
}
func (m *MessageOneof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOneof.Marshal(b, m, deterministic)
}
func (m *MessageOneof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOneof.Merge(m, src)
}
func (m *MessageOneof) XXX_Size() int {
	return xxx_messageInfo_MessageOneof.Size(m)
}
func (m *MessageOneof) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOneof.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOneof proto.InternalMessageInfo

func (m *MessageOneof) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageOneof) GetFieldNumbers() []int32 {
	if m != nil {
		return m.FieldNumbers
	}
	return nil
}

// Service describes a Protobuf service.
type Service struct {
	// name is the name of the service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// service_methods are the service methods.
	//
	// These will be sorted by name.
	ServiceMethods       []*ServiceMethod `protobuf:"bytes,2,rep,name=service_methods,json=serviceMethods,proto3" json:"service_methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{7}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetServiceMethods() []*ServiceMethod {
	if m != nil {
		return m.ServiceMethods
	}
	return nil
}

// ServiceMethod describes a Protobuf service method.
type ServiceMethod struct {
	// name is the name of the service method.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// request_type_name is the fully-qualified name of request message type.
	//
	// This does not include the prefix '.' found in the traditional package
	// fully-qualified name. If this is a nested message, the parent messages
	// will be part of the name.
	RequestTypeName string `protobuf:"bytes,2,opt,name=request_type_name,json=requestTypeName,proto3" json:"request_type_name,omitempty"`
	// response_type_name is the fully-qualified name of response message type.
	//
	// This does not include the prefix '.' found in the traditional package
	// fully-qualified name. If this is a nested message, the parent messages
	// will be part of the name.
	ResponseTypeName string `protobuf:"bytes,3,opt,name=response_type_name,json=responseTypeName,proto3" json:"response_type_name,omitempty"`
	// client_streaming representing whether this is a client-side streaming
	// method
	ClientStreaming bool `protobuf:"varint,4,opt,name=client_streaming,json=clientStreaming,proto3" json:"client_streaming,omitempty"`
	// server_streaming representing whether this is a server-side streaming
	// method
	ServerStreaming      bool     `protobuf:"varint,5,opt,name=server_streaming,json=serverStreaming,proto3" json:"server_streaming,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceMethod) Reset()         { *m = ServiceMethod{} }
func (m *ServiceMethod) String() string { return proto.CompactTextString(m) }
func (*ServiceMethod) ProtoMessage()    {}
func (*ServiceMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826d1b778a478a3, []int{8}
}

func (m *ServiceMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMethod.Unmarshal(m, b)
}
func (m *ServiceMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMethod.Marshal(b, m, deterministic)
}
func (m *ServiceMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMethod.Merge(m, src)
}
func (m *ServiceMethod) XXX_Size() int {
	return xxx_messageInfo_ServiceMethod.Size(m)
}
func (m *ServiceMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMethod proto.InternalMessageInfo

func (m *ServiceMethod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceMethod) GetRequestTypeName() string {
	if m != nil {
		return m.RequestTypeName
	}
	return ""
}

func (m *ServiceMethod) GetResponseTypeName() string {
	if m != nil {
		return m.ResponseTypeName
	}
	return ""
}

func (m *ServiceMethod) GetClientStreaming() bool {
	if m != nil {
		return m.ClientStreaming
	}
	return false
}

func (m *ServiceMethod) GetServerStreaming() bool {
	if m != nil {
		return m.ServerStreaming
	}
	return false
}

func init() {
	proto.RegisterEnum("uber.proto.reflect.v1.MessageField_Label", MessageField_Label_name, MessageField_Label_value)
	proto.RegisterEnum("uber.proto.reflect.v1.MessageField_Type", MessageField_Type_name, MessageField_Type_value)
	proto.RegisterType((*PackageSet)(nil), "uber.proto.reflect.v1.PackageSet")
	proto.RegisterType((*Package)(nil), "uber.proto.reflect.v1.Package")
	proto.RegisterType((*Enum)(nil), "uber.proto.reflect.v1.Enum")
	proto.RegisterType((*EnumValue)(nil), "uber.proto.reflect.v1.EnumValue")
	proto.RegisterType((*Message)(nil), "uber.proto.reflect.v1.Message")
	proto.RegisterType((*MessageField)(nil), "uber.proto.reflect.v1.MessageField")
	proto.RegisterType((*MessageOneof)(nil), "uber.proto.reflect.v1.MessageOneof")
	proto.RegisterType((*Service)(nil), "uber.proto.reflect.v1.Service")
	proto.RegisterType((*ServiceMethod)(nil), "uber.proto.reflect.v1.ServiceMethod")
}

func init() {
	proto.RegisterFile("uber/proto/reflect/v1/reflect.proto", fileDescriptor_4826d1b778a478a3)
}

var fileDescriptor_4826d1b778a478a3 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xeb, 0x44,
	0x14, 0xc5, 0x4e, 0xdc, 0x36, 0x37, 0x5f, 0xd3, 0x91, 0x1e, 0x32, 0x7a, 0x12, 0x8a, 0x5c, 0x16,
	0x29, 0x42, 0xa9, 0x92, 0x56, 0x45, 0x42, 0x08, 0x94, 0x28, 0x6e, 0x08, 0xca, 0x87, 0x99, 0x38,
	0x11, 0x0f, 0x15, 0x59, 0x69, 0x3a, 0x7d, 0x54, 0xd8, 0x4e, 0xf0, 0x38, 0x91, 0xfa, 0x77, 0x58,
	0xf2, 0x23, 0x58, 0x23, 0xb6, 0xfc, 0x14, 0x56, 0xec, 0xd0, 0xcc, 0xd8, 0x53, 0x07, 0x99, 0xbc,
	0xd7, 0x95, 0x67, 0xce, 0x3d, 0xe7, 0xde, 0x33, 0xe3, 0x7b, 0x07, 0xce, 0xb6, 0x77, 0x34, 0xba,
	0xd8, 0x44, 0xeb, 0x78, 0x7d, 0x11, 0xd1, 0x07, 0x9f, 0xae, 0xe2, 0x8b, 0x5d, 0x3b, 0x5d, 0xb6,
	0x44, 0x00, 0xbf, 0xe2, 0x24, 0xb9, 0x6e, 0xa5, 0x91, 0x5d, 0xdb, 0xfa, 0x06, 0xc0, 0x59, 0xae,
	0x7e, 0x5e, 0xbe, 0xa5, 0x33, 0x1a, 0xe3, 0x2f, 0xe0, 0x64, 0x23, 0x77, 0xcc, 0xd4, 0x1a, 0x85,
	0x66, 0xb9, 0xf3, 0x71, 0x2b, 0x57, 0xd7, 0x4a, 0x44, 0x44, 0xf1, 0xad, 0xbf, 0x35, 0x38, 0x4e,
	0x50, 0x8c, 0xa1, 0x18, 0x2e, 0x03, 0x6a, 0x6a, 0x0d, 0xad, 0x59, 0x22, 0x62, 0x8d, 0xcf, 0x01,
	0xdd, 0xd3, 0x0d, 0x0d, 0xef, 0x69, 0xb8, 0x7a, 0xf2, 0x38, 0xc4, 0x4c, 0xbd, 0x51, 0x68, 0x96,
	0x48, 0xfd, 0x19, 0x9f, 0x70, 0x18, 0xb7, 0xc1, 0xa0, 0xe1, 0x36, 0x60, 0x66, 0x41, 0x78, 0x78,
	0xfd, 0x3f, 0x1e, 0xec, 0x70, 0x1b, 0x10, 0xc9, 0xe4, 0xce, 0x03, 0xca, 0x98, 0x70, 0x5e, 0x3c,
	0xe8, 0x7c, 0x2c, 0x69, 0x44, 0xf1, 0xb9, 0x96, 0xd1, 0x68, 0xf7, 0xb8, 0xa2, 0xcc, 0x34, 0x0e,
	0x6a, 0x67, 0x92, 0x46, 0x14, 0xdf, 0xfa, 0x11, 0x8a, 0xdc, 0x46, 0xee, 0x89, 0xbb, 0x50, 0xe6,
	0xe6, 0xbc, 0xdd, 0xd2, 0xdf, 0x26, 0x87, 0x2d, 0x77, 0x1a, 0x07, 0x0e, 0xb3, 0xe0, 0x44, 0x02,
	0x34, 0x5d, 0x32, 0xeb, 0x73, 0x28, 0xa9, 0x40, 0x6e, 0x8d, 0x0f, 0xe1, 0x28, 0xdc, 0x06, 0x77,
	0x34, 0x32, 0xf5, 0x86, 0xd6, 0x34, 0x48, 0xb2, 0xb3, 0x7e, 0xd7, 0xe1, 0x38, 0x39, 0x69, 0xae,
	0xee, 0x5b, 0xa8, 0x25, 0xe7, 0xf7, 0x1e, 0x1e, 0xa9, 0x7f, 0x9f, 0xda, 0x3b, 0x3b, 0x7c, 0x6b,
	0x37, 0x9c, 0x4b, 0xaa, 0x41, 0x66, 0xc7, 0xb2, 0xb9, 0xd6, 0x21, 0x5d, 0x3f, 0xa4, 0xff, 0xed,
	0x1d, 0xb9, 0xa6, 0x9c, 0xab, 0x72, 0x89, 0x1d, 0xc3, 0x03, 0xa8, 0x87, 0x94, 0xc5, 0xf4, 0xde,
	0x7b, 0xe1, 0xef, 0xac, 0x49, 0xd9, 0x38, 0xfd, 0xa9, 0x5f, 0x41, 0x25, 0x49, 0x24, 0x5b, 0xc9,
	0x78, 0x77, 0x2b, 0x95, 0xa5, 0x80, 0xaf, 0x99, 0xf5, 0x4f, 0x11, 0x2a, 0xd9, 0x43, 0xbf, 0xe4,
	0xf6, 0xf1, 0xd7, 0x60, 0xf8, 0xcb, 0x3b, 0xea, 0x9b, 0x85, 0x86, 0xd6, 0xac, 0x75, 0xce, 0xdf,
	0xe3, 0x52, 0x5b, 0x23, 0x2e, 0x20, 0x52, 0x87, 0xbf, 0x84, 0x62, 0xfc, 0xb4, 0xa1, 0x66, 0x51,
	0xe8, 0x9b, 0xef, 0xa3, 0x77, 0x9f, 0x36, 0x94, 0x08, 0x15, 0x7e, 0x0d, 0x25, 0xfe, 0x15, 0x43,
	0x66, 0x1a, 0xc2, 0xef, 0x09, 0x07, 0xf8, 0x74, 0x59, 0x0b, 0x30, 0x44, 0x29, 0x7c, 0x0a, 0xd5,
	0x51, 0xb7, 0x67, 0x8f, 0xbc, 0xe1, 0x64, 0xd1, 0x1d, 0x0d, 0xfb, 0xe8, 0x03, 0x8c, 0xa1, 0x26,
	0xa1, 0xa9, 0xe3, 0x0e, 0xa7, 0x93, 0xee, 0x08, 0x69, 0xcf, 0x18, 0xb1, 0xbf, 0x9b, 0x0f, 0x89,
	0xdd, 0x47, 0x7a, 0x16, 0x73, 0xec, 0xae, 0x6b, 0xf7, 0x51, 0xc1, 0xfa, 0x43, 0x87, 0x22, 0xf7,
	0x80, 0x11, 0x54, 0xdc, 0x37, 0x8e, 0x9d, 0x49, 0x5b, 0x87, 0xb2, 0x40, 0xfa, 0xd3, 0x79, 0x6f,
	0x64, 0x23, 0x0d, 0xd7, 0x00, 0x04, 0x70, 0x33, 0x9a, 0x76, 0x5d, 0xa4, 0xab, 0xfd, 0x70, 0xe2,
	0x5e, 0x5f, 0xa1, 0x82, 0x12, 0xcc, 0x25, 0x50, 0xcc, 0x12, 0x2e, 0x3b, 0xc8, 0x50, 0x35, 0x6e,
	0x86, 0xdf, 0xdb, 0xfd, 0xeb, 0x2b, 0x74, 0xb4, 0x8f, 0x5c, 0x76, 0xd0, 0x31, 0xae, 0x42, 0x49,
	0x20, 0xbd, 0xe9, 0x74, 0x84, 0x4e, 0x54, 0xce, 0x99, 0x4b, 0x86, 0x93, 0x01, 0x2a, 0xa9, 0x9c,
	0x03, 0x32, 0x9d, 0x3b, 0x08, 0x54, 0x86, 0xb1, 0x3d, 0x9b, 0x75, 0x07, 0x36, 0x2a, 0x2b, 0x46,
	0xef, 0x8d, 0x6b, 0xcf, 0x50, 0x65, 0xcf, 0xd6, 0x65, 0x07, 0x55, 0x55, 0x09, 0x7b, 0x32, 0x1f,
	0xa3, 0x1a, 0xbf, 0x51, 0x59, 0x22, 0x35, 0x51, 0xff, 0x0f, 0x74, 0x7d, 0x85, 0xd0, 0xb3, 0x11,
	0x99, 0xe5, 0x74, 0x0f, 0xb8, 0xbe, 0x42, 0xd8, 0x1a, 0xa8, 0xd6, 0x13, 0x53, 0x91, 0xdb, 0x7a,
	0x67, 0x50, 0x15, 0x83, 0xeb, 0xc9, 0x96, 0x93, 0xf3, 0x6b, 0x90, 0x8a, 0x00, 0x27, 0x12, 0xb3,
	0x7c, 0x38, 0x4e, 0x9e, 0xac, 0xdc, 0x1c, 0x63, 0xa8, 0x27, 0x0f, 0x99, 0x17, 0xd0, 0xf8, 0xa7,
	0xb5, 0x7a, 0x05, 0x3e, 0x39, 0xfc, 0xfe, 0x8d, 0x05, 0x99, 0xd4, 0x58, 0x76, 0xcb, 0xac, 0xbf,
	0x34, 0xa8, 0xee, 0x31, 0x72, 0x8b, 0x7e, 0x0a, 0xa7, 0x11, 0xfd, 0x65, 0x4b, 0x59, 0xec, 0x3d,
	0x37, 0xa9, 0x2e, 0x08, 0xf5, 0x24, 0xe0, 0x26, 0xbd, 0x8a, 0x3f, 0x03, 0x1c, 0x51, 0xb6, 0x59,
	0x87, 0x8c, 0x66, 0xc8, 0x05, 0x41, 0x46, 0x69, 0x44, 0xb1, 0xcf, 0x01, 0xad, 0xfc, 0x47, 0x1a,
	0xc6, 0x1e, 0x8b, 0x23, 0xba, 0x0c, 0x1e, 0xc3, 0xb7, 0x62, 0x80, 0x4e, 0x48, 0x5d, 0xe2, 0xb3,
	0x14, 0xe6, 0x54, 0x6e, 0x9e, 0x46, 0x19, 0xaa, 0x21, 0xa9, 0x12, 0x57, 0xd4, 0x9e, 0x0f, 0x1f,
	0xad, 0xd6, 0x41, 0xfe, 0x85, 0xf4, 0x2a, 0x44, 0xae, 0x1d, 0x1e, 0x70, 0xb4, 0x1f, 0x4a, 0x49,
	0x6c, 0xd7, 0xfe, 0x55, 0x2f, 0xcc, 0x1d, 0xf2, 0x9b, 0xfe, 0x6a, 0xce, 0x85, 0x22, 0xde, 0x4a,
	0xc8, 0xad, 0x45, 0xfb, 0x4f, 0x89, 0xdf, 0x0a, 0xfc, 0x36, 0xc1, 0x6f, 0x17, 0xed, 0xbb, 0x23,
	0x51, 0xe2, 0xf2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x32, 0x37, 0xc1, 0xd4, 0x07, 0x00,
	0x00,
}

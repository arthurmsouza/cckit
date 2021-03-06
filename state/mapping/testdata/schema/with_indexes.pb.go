// Code generated by protoc-gen-go. DO NOT EDIT.
// source: with_indexes.proto

package schema

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

// EntityWithIndexes
type EntityWithIndexes struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds  []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value                int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityWithIndexes) Reset()         { *m = EntityWithIndexes{} }
func (m *EntityWithIndexes) String() string { return proto.CompactTextString(m) }
func (*EntityWithIndexes) ProtoMessage()    {}
func (*EntityWithIndexes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22c1b92282e067, []int{0}
}

func (m *EntityWithIndexes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityWithIndexes.Unmarshal(m, b)
}
func (m *EntityWithIndexes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityWithIndexes.Marshal(b, m, deterministic)
}
func (m *EntityWithIndexes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityWithIndexes.Merge(m, src)
}
func (m *EntityWithIndexes) XXX_Size() int {
	return xxx_messageInfo_EntityWithIndexes.Size(m)
}
func (m *EntityWithIndexes) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityWithIndexes.DiscardUnknown(m)
}

var xxx_messageInfo_EntityWithIndexes proto.InternalMessageInfo

func (m *EntityWithIndexes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EntityWithIndexes) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *EntityWithIndexes) GetRequiredExternalIds() []string {
	if m != nil {
		return m.RequiredExternalIds
	}
	return nil
}

func (m *EntityWithIndexes) GetOptionalExternalIds() []string {
	if m != nil {
		return m.OptionalExternalIds
	}
	return nil
}

func (m *EntityWithIndexes) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// EntityWithIndexesList
type EntityWithIndexesList struct {
	Items                []*EntityWithIndexes `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EntityWithIndexesList) Reset()         { *m = EntityWithIndexesList{} }
func (m *EntityWithIndexesList) String() string { return proto.CompactTextString(m) }
func (*EntityWithIndexesList) ProtoMessage()    {}
func (*EntityWithIndexesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22c1b92282e067, []int{1}
}

func (m *EntityWithIndexesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityWithIndexesList.Unmarshal(m, b)
}
func (m *EntityWithIndexesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityWithIndexesList.Marshal(b, m, deterministic)
}
func (m *EntityWithIndexesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityWithIndexesList.Merge(m, src)
}
func (m *EntityWithIndexesList) XXX_Size() int {
	return xxx_messageInfo_EntityWithIndexesList.Size(m)
}
func (m *EntityWithIndexesList) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityWithIndexesList.DiscardUnknown(m)
}

var xxx_messageInfo_EntityWithIndexesList proto.InternalMessageInfo

func (m *EntityWithIndexesList) GetItems() []*EntityWithIndexes {
	if m != nil {
		return m.Items
	}
	return nil
}

// CreateEntityWithIndexes
type CreateEntityWithIndexes struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds  []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value                int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityWithIndexes) Reset()         { *m = CreateEntityWithIndexes{} }
func (m *CreateEntityWithIndexes) String() string { return proto.CompactTextString(m) }
func (*CreateEntityWithIndexes) ProtoMessage()    {}
func (*CreateEntityWithIndexes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22c1b92282e067, []int{2}
}

func (m *CreateEntityWithIndexes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityWithIndexes.Unmarshal(m, b)
}
func (m *CreateEntityWithIndexes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityWithIndexes.Marshal(b, m, deterministic)
}
func (m *CreateEntityWithIndexes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityWithIndexes.Merge(m, src)
}
func (m *CreateEntityWithIndexes) XXX_Size() int {
	return xxx_messageInfo_CreateEntityWithIndexes.Size(m)
}
func (m *CreateEntityWithIndexes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityWithIndexes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityWithIndexes proto.InternalMessageInfo

func (m *CreateEntityWithIndexes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateEntityWithIndexes) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *CreateEntityWithIndexes) GetRequiredExternalIds() []string {
	if m != nil {
		return m.RequiredExternalIds
	}
	return nil
}

func (m *CreateEntityWithIndexes) GetOptionalExternalIds() []string {
	if m != nil {
		return m.OptionalExternalIds
	}
	return nil
}

func (m *CreateEntityWithIndexes) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// UpdateEntityEntityWithIndexes
type UpdateEntityWithIndexes struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds  []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value                int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEntityWithIndexes) Reset()         { *m = UpdateEntityWithIndexes{} }
func (m *UpdateEntityWithIndexes) String() string { return proto.CompactTextString(m) }
func (*UpdateEntityWithIndexes) ProtoMessage()    {}
func (*UpdateEntityWithIndexes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22c1b92282e067, []int{3}
}

func (m *UpdateEntityWithIndexes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEntityWithIndexes.Unmarshal(m, b)
}
func (m *UpdateEntityWithIndexes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEntityWithIndexes.Marshal(b, m, deterministic)
}
func (m *UpdateEntityWithIndexes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEntityWithIndexes.Merge(m, src)
}
func (m *UpdateEntityWithIndexes) XXX_Size() int {
	return xxx_messageInfo_UpdateEntityWithIndexes.Size(m)
}
func (m *UpdateEntityWithIndexes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEntityWithIndexes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEntityWithIndexes proto.InternalMessageInfo

func (m *UpdateEntityWithIndexes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateEntityWithIndexes) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *UpdateEntityWithIndexes) GetRequiredExternalIds() []string {
	if m != nil {
		return m.RequiredExternalIds
	}
	return nil
}

func (m *UpdateEntityWithIndexes) GetOptionalExternalIds() []string {
	if m != nil {
		return m.OptionalExternalIds
	}
	return nil
}

func (m *UpdateEntityWithIndexes) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*EntityWithIndexes)(nil), "schema.EntityWithIndexes")
	proto.RegisterType((*EntityWithIndexesList)(nil), "schema.EntityWithIndexesList")
	proto.RegisterType((*CreateEntityWithIndexes)(nil), "schema.CreateEntityWithIndexes")
	proto.RegisterType((*UpdateEntityWithIndexes)(nil), "schema.UpdateEntityWithIndexes")
}

func init() { proto.RegisterFile("with_indexes.proto", fileDescriptor_fd22c1b92282e067) }

var fileDescriptor_fd22c1b92282e067 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x49, 0x6b, 0x07, 0xfb, 0x06, 0x82, 0xd1, 0x61, 0x3c, 0x59, 0x7a, 0xca, 0xa9, 0xc2,
	0x7c, 0x04, 0x19, 0x38, 0xf0, 0x14, 0x10, 0x8f, 0x25, 0x9a, 0x0f, 0xfa, 0xc1, 0xd6, 0xd4, 0xe4,
	0x9b, 0xce, 0xd7, 0xf3, 0xe4, 0x63, 0x49, 0x17, 0x2a, 0x93, 0x3d, 0x81, 0x1e, 0x93, 0xff, 0xef,
	0x77, 0xf8, 0x85, 0x80, 0x7c, 0x27, 0x6e, 0x1b, 0xea, 0x1c, 0xee, 0x30, 0xd6, 0x7d, 0xf0, 0xec,
	0xe5, 0x24, 0xbe, 0xb4, 0xb8, 0xb1, 0xd5, 0xa7, 0x80, 0xb3, 0x65, 0xc7, 0xc4, 0x1f, 0x4f, 0xc4,
	0xed, 0x2a, 0x31, 0xf2, 0x14, 0x32, 0x72, 0x4a, 0x94, 0x42, 0x4f, 0x4d, 0x46, 0x4e, 0x5e, 0xc3,
	0x0c, 0x77, 0x8c, 0xa1, 0xb3, 0xeb, 0x86, 0x9c, 0xca, 0xf6, 0x03, 0x8c, 0x57, 0x2b, 0x27, 0x17,
	0x30, 0x0f, 0xf8, 0xba, 0xa5, 0x80, 0xae, 0x39, 0x20, 0xa3, 0xca, 0xcb, 0x5c, 0x4f, 0xcd, 0xf9,
	0x38, 0x2e, 0x7f, 0x94, 0x38, 0x38, 0xbe, 0x67, 0xf2, 0x03, 0xfa, 0xcb, 0x39, 0x49, 0xce, 0x38,
	0x1e, 0x3a, 0x17, 0x50, 0xbc, 0xd9, 0xf5, 0x16, 0x55, 0x51, 0x0a, 0x5d, 0x98, 0x74, 0xa8, 0xee,
	0x61, 0x7e, 0xd4, 0xf0, 0x40, 0x91, 0xe5, 0x0d, 0x14, 0xc4, 0xb8, 0x89, 0x4a, 0x94, 0xb9, 0x9e,
	0x2d, 0xae, 0xea, 0x54, 0x5d, 0x1f, 0xd1, 0x26, 0x71, 0xd5, 0x97, 0x80, 0xcb, 0xbb, 0x80, 0x96,
	0xf1, 0xcf, 0x3f, 0xca, 0x90, 0xf2, 0xd8, 0xbb, 0x7f, 0x90, 0xf2, 0x3c, 0xd9, 0xff, 0xd9, 0xdb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x94, 0x9b, 0x51, 0xc9, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/databags.proto

package response

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

type DataBags struct {
	// List of data bags item.
	DataBags             []*DataBagListItem `protobuf:"bytes,2,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataBags) Reset()         { *m = DataBags{} }
func (m *DataBags) String() string { return proto.CompactTextString(m) }
func (*DataBags) ProtoMessage()    {}
func (*DataBags) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{0}
}

func (m *DataBags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBags.Unmarshal(m, b)
}
func (m *DataBags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBags.Marshal(b, m, deterministic)
}
func (m *DataBags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBags.Merge(m, src)
}
func (m *DataBags) XXX_Size() int {
	return xxx_messageInfo_DataBags.Size(m)
}
func (m *DataBags) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBags.DiscardUnknown(m)
}

var xxx_messageInfo_DataBags proto.InternalMessageInfo

func (m *DataBags) GetDataBags() []*DataBagListItem {
	if m != nil {
		return m.DataBags
	}
	return nil
}

type DataBagListItem struct {
	// Name of the data bag item.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataBagListItem) Reset()         { *m = DataBagListItem{} }
func (m *DataBagListItem) String() string { return proto.CompactTextString(m) }
func (*DataBagListItem) ProtoMessage()    {}
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{1}
}

func (m *DataBagListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBagListItem.Unmarshal(m, b)
}
func (m *DataBagListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBagListItem.Marshal(b, m, deterministic)
}
func (m *DataBagListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBagListItem.Merge(m, src)
}
func (m *DataBagListItem) XXX_Size() int {
	return xxx_messageInfo_DataBagListItem.Size(m)
}
func (m *DataBagListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBagListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DataBagListItem proto.InternalMessageInfo

func (m *DataBagListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataBag struct {
	// Stringified json of the data bag item.
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataBag) Reset()         { *m = DataBag{} }
func (m *DataBag) String() string { return proto.CompactTextString(m) }
func (*DataBag) ProtoMessage()    {}
func (*DataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{2}
}

func (m *DataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBag.Unmarshal(m, b)
}
func (m *DataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBag.Marshal(b, m, deterministic)
}
func (m *DataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBag.Merge(m, src)
}
func (m *DataBag) XXX_Size() int {
	return xxx_messageInfo_DataBag.Size(m)
}
func (m *DataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBag.DiscardUnknown(m)
}

var xxx_messageInfo_DataBag proto.InternalMessageInfo

func (m *DataBag) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*DataBags)(nil), "chef.automate.api.infra_proxy.response.DataBags")
	proto.RegisterType((*DataBagListItem)(nil), "chef.automate.api.infra_proxy.response.DataBagListItem")
	proto.RegisterType((*DataBag)(nil), "chef.automate.api.infra_proxy.response.DataBag")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/databags.proto", fileDescriptor_e518f553399bb7b1)
}

var fileDescriptor_e518f553399bb7b1 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3d, 0x4b, 0xc5, 0x40,
	0x10, 0x45, 0x79, 0x2a, 0xfa, 0xde, 0x5a, 0x08, 0x5b, 0xa5, 0x11, 0x42, 0x40, 0x49, 0x35, 0x2b,
	0x5a, 0x08, 0x62, 0x25, 0x36, 0x82, 0x55, 0xb0, 0xb2, 0x89, 0x93, 0x64, 0x92, 0x2c, 0xb8, 0x1f,
	0xec, 0x4e, 0x20, 0xfe, 0x7b, 0xd9, 0x90, 0x80, 0xd8, 0x68, 0x37, 0x70, 0xe7, 0x1c, 0x2e, 0x57,
	0xdc, 0xa0, 0xd7, 0x8a, 0x66, 0xa6, 0x60, 0xf1, 0x53, 0x69, 0xdb, 0x07, 0xac, 0x7d, 0x70, 0xf3,
	0x97, 0x0a, 0x14, 0xbd, 0xb3, 0x91, 0x54, 0x87, 0x8c, 0x0d, 0x0e, 0x11, 0x7c, 0x70, 0xec, 0xe4,
	0x75, 0x3b, 0x52, 0x0f, 0x38, 0xb1, 0x33, 0xc8, 0x04, 0xe8, 0x35, 0xfc, 0xc0, 0x60, 0xc3, 0x8a,
	0x0f, 0xb1, 0x7f, 0x46, 0xc6, 0x27, 0x1c, 0xa2, 0x7c, 0x13, 0x87, 0x64, 0xa9, 0x93, 0x26, 0x3b,
	0xca, 0x8f, 0xcb, 0xf3, 0xdb, 0x7b, 0xf8, 0x9f, 0x07, 0x56, 0xc9, 0xab, 0x8e, 0xfc, 0xc2, 0x64,
	0xaa, 0x7d, 0xb7, 0x5a, 0x8b, 0x2b, 0x71, 0xf1, 0x2b, 0x94, 0x52, 0x9c, 0x58, 0x34, 0x94, 0xed,
	0xf2, 0x5d, 0x79, 0xa8, 0x96, 0xbb, 0xb8, 0x14, 0x67, 0xeb, 0x5b, 0x8a, 0x13, 0xbd, 0xc5, 0x8b,
	0xe9, 0xf1, 0xfd, 0x61, 0xd0, 0x3c, 0x4e, 0x0d, 0xb4, 0xce, 0xa8, 0x54, 0x4a, 0x6d, 0xa5, 0xd4,
	0x9f, 0xe3, 0x34, 0xa7, 0xcb, 0x28, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22, 0x38, 0xcc,
	0xc3, 0x48, 0x01, 0x00, 0x00,
}
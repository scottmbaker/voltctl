// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bbf_fiber_multicast_gemport_body.proto

package bbf_fiber

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/opencord/voltha/protos/go/common"
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

type MulticastGemportsConfigData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GemportId            uint32   `protobuf:"varint,3,opt,name=gemport_id,json=gemportId,proto3" json:"gemport_id,omitempty"`
	ItfRef               string   `protobuf:"bytes,4,opt,name=itf_ref,json=itfRef,proto3" json:"itf_ref,omitempty"`
	TrafficClass         uint32   `protobuf:"varint,5,opt,name=traffic_class,json=trafficClass,proto3" json:"traffic_class,omitempty"`
	IsBroadcast          bool     `protobuf:"varint,6,opt,name=is_broadcast,json=isBroadcast,proto3" json:"is_broadcast,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MulticastGemportsConfigData) Reset()         { *m = MulticastGemportsConfigData{} }
func (m *MulticastGemportsConfigData) String() string { return proto.CompactTextString(m) }
func (*MulticastGemportsConfigData) ProtoMessage()    {}
func (*MulticastGemportsConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_707ac0af9d681d27, []int{0}
}

func (m *MulticastGemportsConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MulticastGemportsConfigData.Unmarshal(m, b)
}
func (m *MulticastGemportsConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MulticastGemportsConfigData.Marshal(b, m, deterministic)
}
func (m *MulticastGemportsConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulticastGemportsConfigData.Merge(m, src)
}
func (m *MulticastGemportsConfigData) XXX_Size() int {
	return xxx_messageInfo_MulticastGemportsConfigData.Size(m)
}
func (m *MulticastGemportsConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_MulticastGemportsConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_MulticastGemportsConfigData proto.InternalMessageInfo

func (m *MulticastGemportsConfigData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MulticastGemportsConfigData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MulticastGemportsConfigData) GetGemportId() uint32 {
	if m != nil {
		return m.GemportId
	}
	return 0
}

func (m *MulticastGemportsConfigData) GetItfRef() string {
	if m != nil {
		return m.ItfRef
	}
	return ""
}

func (m *MulticastGemportsConfigData) GetTrafficClass() uint32 {
	if m != nil {
		return m.TrafficClass
	}
	return 0
}

func (m *MulticastGemportsConfigData) GetIsBroadcast() bool {
	if m != nil {
		return m.IsBroadcast
	}
	return false
}

func init() {
	proto.RegisterType((*MulticastGemportsConfigData)(nil), "bbf_fiber.MulticastGemportsConfigData")
}

func init() {
	proto.RegisterFile("bbf_fiber_multicast_gemport_body.proto", fileDescriptor_707ac0af9d681d27)
}

var fileDescriptor_707ac0af9d681d27 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xdd, 0xad, 0x36, 0xee, 0x5e, 0x02, 0x62, 0x50, 0x16, 0xaa, 0x82, 0xf4, 0xd4,
	0x0a, 0xbe, 0x41, 0x57, 0x10, 0x0f, 0x5e, 0x7a, 0xf4, 0x12, 0x92, 0x26, 0xe9, 0x0e, 0xb4, 0x4d,
	0x49, 0x66, 0x05, 0xdf, 0xd0, 0x8b, 0xef, 0xe1, 0x13, 0x78, 0x96, 0x76, 0xb7, 0x7b, 0x1b, 0xbe,
	0x99, 0x8f, 0x99, 0xf9, 0xe9, 0xa3, 0x52, 0x56, 0x58, 0x50, 0xc6, 0x8b, 0x6e, 0xdf, 0x22, 0xd4,
	0x32, 0xa0, 0x68, 0x4c, 0x37, 0x38, 0x8f, 0x42, 0x39, 0xfd, 0x95, 0x0f, 0xde, 0xa1, 0x63, 0xc9,
	0x69, 0xee, 0x86, 0x76, 0x06, 0xe5, 0x01, 0xdf, 0x7f, 0x13, 0x7a, 0xfb, 0x3e, 0x7b, 0xaf, 0x07,
	0x2d, 0x6c, 0x5d, 0x6f, 0xa1, 0x79, 0x91, 0x28, 0xd9, 0x15, 0x8d, 0x40, 0x73, 0x92, 0x92, 0x2c,
	0x29, 0x97, 0xbf, 0x7f, 0x3f, 0x1b, 0x52, 0x45, 0xa0, 0x19, 0xa3, 0x8b, 0x5e, 0x76, 0x86, 0x47,
	0x63, 0xa3, 0x9a, 0x6a, 0xb6, 0xa1, 0x74, 0xde, 0x0b, 0x9a, 0x9f, 0xa5, 0x24, 0x5b, 0x57, 0xc9,
	0x91, 0xbc, 0x69, 0x76, 0x4d, 0xcf, 0x01, 0xad, 0xf0, 0xc6, 0xf2, 0xc5, 0x64, 0xc5, 0x80, 0xb6,
	0x32, 0x96, 0x3d, 0xd0, 0x35, 0x7a, 0x69, 0x2d, 0xd4, 0xa2, 0x6e, 0x65, 0x08, 0x7c, 0x39, 0xa9,
	0xab, 0x23, 0xdc, 0x8e, 0x8c, 0xdd, 0xd1, 0x15, 0x04, 0xa1, 0xbc, 0x93, 0x7a, 0xbc, 0x94, 0xc7,
	0x29, 0xc9, 0x2e, 0xaa, 0x4b, 0x08, 0xe5, 0x8c, 0xca, 0xa7, 0x8f, 0xbc, 0x01, 0xdc, 0xed, 0x55,
	0x5e, 0xbb, 0xae, 0x70, 0x83, 0xe9, 0x6b, 0xe7, 0x75, 0xf1, 0xe9, 0x5a, 0xdc, 0xc9, 0x62, 0x7a,
	0x37, 0x14, 0x8d, 0x2b, 0x4e, 0x41, 0xa8, 0x78, 0x82, 0xcf, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xab, 0xc1, 0xcb, 0xfa, 0x44, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: omci_alarm_db.proto

package omci

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

type AlarmOpenOmciEventType_OpenOmciEventType int32

const (
	AlarmOpenOmciEventType_state_change AlarmOpenOmciEventType_OpenOmciEventType = 0
)

var AlarmOpenOmciEventType_OpenOmciEventType_name = map[int32]string{
	0: "state_change",
}

var AlarmOpenOmciEventType_OpenOmciEventType_value = map[string]int32{
	"state_change": 0,
}

func (x AlarmOpenOmciEventType_OpenOmciEventType) String() string {
	return proto.EnumName(AlarmOpenOmciEventType_OpenOmciEventType_name, int32(x))
}

func (AlarmOpenOmciEventType_OpenOmciEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{6, 0}
}

type AlarmAttributeData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmAttributeData) Reset()         { *m = AlarmAttributeData{} }
func (m *AlarmAttributeData) String() string { return proto.CompactTextString(m) }
func (*AlarmAttributeData) ProtoMessage()    {}
func (*AlarmAttributeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{0}
}

func (m *AlarmAttributeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmAttributeData.Unmarshal(m, b)
}
func (m *AlarmAttributeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmAttributeData.Marshal(b, m, deterministic)
}
func (m *AlarmAttributeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmAttributeData.Merge(m, src)
}
func (m *AlarmAttributeData) XXX_Size() int {
	return xxx_messageInfo_AlarmAttributeData.Size(m)
}
func (m *AlarmAttributeData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmAttributeData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmAttributeData proto.InternalMessageInfo

func (m *AlarmAttributeData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlarmAttributeData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AlarmInstanceData struct {
	InstanceId           uint32                `protobuf:"varint,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Created              string                `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Modified             string                `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	Attributes           []*AlarmAttributeData `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AlarmInstanceData) Reset()         { *m = AlarmInstanceData{} }
func (m *AlarmInstanceData) String() string { return proto.CompactTextString(m) }
func (*AlarmInstanceData) ProtoMessage()    {}
func (*AlarmInstanceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{1}
}

func (m *AlarmInstanceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmInstanceData.Unmarshal(m, b)
}
func (m *AlarmInstanceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmInstanceData.Marshal(b, m, deterministic)
}
func (m *AlarmInstanceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmInstanceData.Merge(m, src)
}
func (m *AlarmInstanceData) XXX_Size() int {
	return xxx_messageInfo_AlarmInstanceData.Size(m)
}
func (m *AlarmInstanceData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmInstanceData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmInstanceData proto.InternalMessageInfo

func (m *AlarmInstanceData) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *AlarmInstanceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *AlarmInstanceData) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *AlarmInstanceData) GetAttributes() []*AlarmAttributeData {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type AlarmClassData struct {
	ClassId              uint32               `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Instances            []*AlarmInstanceData `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AlarmClassData) Reset()         { *m = AlarmClassData{} }
func (m *AlarmClassData) String() string { return proto.CompactTextString(m) }
func (*AlarmClassData) ProtoMessage()    {}
func (*AlarmClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{2}
}

func (m *AlarmClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmClassData.Unmarshal(m, b)
}
func (m *AlarmClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmClassData.Marshal(b, m, deterministic)
}
func (m *AlarmClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmClassData.Merge(m, src)
}
func (m *AlarmClassData) XXX_Size() int {
	return xxx_messageInfo_AlarmClassData.Size(m)
}
func (m *AlarmClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmClassData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmClassData proto.InternalMessageInfo

func (m *AlarmClassData) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *AlarmClassData) GetInstances() []*AlarmInstanceData {
	if m != nil {
		return m.Instances
	}
	return nil
}

type AlarmManagedEntity struct {
	ClassId              uint32   `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmManagedEntity) Reset()         { *m = AlarmManagedEntity{} }
func (m *AlarmManagedEntity) String() string { return proto.CompactTextString(m) }
func (*AlarmManagedEntity) ProtoMessage()    {}
func (*AlarmManagedEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{3}
}

func (m *AlarmManagedEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmManagedEntity.Unmarshal(m, b)
}
func (m *AlarmManagedEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmManagedEntity.Marshal(b, m, deterministic)
}
func (m *AlarmManagedEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmManagedEntity.Merge(m, src)
}
func (m *AlarmManagedEntity) XXX_Size() int {
	return xxx_messageInfo_AlarmManagedEntity.Size(m)
}
func (m *AlarmManagedEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmManagedEntity.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmManagedEntity proto.InternalMessageInfo

func (m *AlarmManagedEntity) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *AlarmManagedEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AlarmMessageType struct {
	MessageType          uint32   `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMessageType) Reset()         { *m = AlarmMessageType{} }
func (m *AlarmMessageType) String() string { return proto.CompactTextString(m) }
func (*AlarmMessageType) ProtoMessage()    {}
func (*AlarmMessageType) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{4}
}

func (m *AlarmMessageType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMessageType.Unmarshal(m, b)
}
func (m *AlarmMessageType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMessageType.Marshal(b, m, deterministic)
}
func (m *AlarmMessageType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMessageType.Merge(m, src)
}
func (m *AlarmMessageType) XXX_Size() int {
	return xxx_messageInfo_AlarmMessageType.Size(m)
}
func (m *AlarmMessageType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMessageType.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMessageType proto.InternalMessageInfo

func (m *AlarmMessageType) GetMessageType() uint32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type AlarmDeviceData struct {
	DeviceId             string                `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Created              string                `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	LastAlarmSequence    uint32                `protobuf:"varint,3,opt,name=last_alarm_sequence,json=lastAlarmSequence,proto3" json:"last_alarm_sequence,omitempty"`
	LastSyncTime         string                `protobuf:"bytes,4,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
	Version              uint32                `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	Classes              []*AlarmClassData     `protobuf:"bytes,6,rep,name=classes,proto3" json:"classes,omitempty"`
	ManagedEntities      []*AlarmManagedEntity `protobuf:"bytes,7,rep,name=managed_entities,json=managedEntities,proto3" json:"managed_entities,omitempty"`
	MessageTypes         []*AlarmMessageType   `protobuf:"bytes,8,rep,name=message_types,json=messageTypes,proto3" json:"message_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AlarmDeviceData) Reset()         { *m = AlarmDeviceData{} }
func (m *AlarmDeviceData) String() string { return proto.CompactTextString(m) }
func (*AlarmDeviceData) ProtoMessage()    {}
func (*AlarmDeviceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{5}
}

func (m *AlarmDeviceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmDeviceData.Unmarshal(m, b)
}
func (m *AlarmDeviceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmDeviceData.Marshal(b, m, deterministic)
}
func (m *AlarmDeviceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmDeviceData.Merge(m, src)
}
func (m *AlarmDeviceData) XXX_Size() int {
	return xxx_messageInfo_AlarmDeviceData.Size(m)
}
func (m *AlarmDeviceData) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmDeviceData.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmDeviceData proto.InternalMessageInfo

func (m *AlarmDeviceData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *AlarmDeviceData) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *AlarmDeviceData) GetLastAlarmSequence() uint32 {
	if m != nil {
		return m.LastAlarmSequence
	}
	return 0
}

func (m *AlarmDeviceData) GetLastSyncTime() string {
	if m != nil {
		return m.LastSyncTime
	}
	return ""
}

func (m *AlarmDeviceData) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AlarmDeviceData) GetClasses() []*AlarmClassData {
	if m != nil {
		return m.Classes
	}
	return nil
}

func (m *AlarmDeviceData) GetManagedEntities() []*AlarmManagedEntity {
	if m != nil {
		return m.ManagedEntities
	}
	return nil
}

func (m *AlarmDeviceData) GetMessageTypes() []*AlarmMessageType {
	if m != nil {
		return m.MessageTypes
	}
	return nil
}

type AlarmOpenOmciEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmOpenOmciEventType) Reset()         { *m = AlarmOpenOmciEventType{} }
func (m *AlarmOpenOmciEventType) String() string { return proto.CompactTextString(m) }
func (*AlarmOpenOmciEventType) ProtoMessage()    {}
func (*AlarmOpenOmciEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{6}
}

func (m *AlarmOpenOmciEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOpenOmciEventType.Unmarshal(m, b)
}
func (m *AlarmOpenOmciEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOpenOmciEventType.Marshal(b, m, deterministic)
}
func (m *AlarmOpenOmciEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOpenOmciEventType.Merge(m, src)
}
func (m *AlarmOpenOmciEventType) XXX_Size() int {
	return xxx_messageInfo_AlarmOpenOmciEventType.Size(m)
}
func (m *AlarmOpenOmciEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOpenOmciEventType.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOpenOmciEventType proto.InternalMessageInfo

type AlarmOpenOmciEvent struct {
	Type                 AlarmOpenOmciEventType_OpenOmciEventType `protobuf:"varint,1,opt,name=type,proto3,enum=alarm.AlarmOpenOmciEventType_OpenOmciEventType" json:"type,omitempty"`
	Data                 string                                   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *AlarmOpenOmciEvent) Reset()         { *m = AlarmOpenOmciEvent{} }
func (m *AlarmOpenOmciEvent) String() string { return proto.CompactTextString(m) }
func (*AlarmOpenOmciEvent) ProtoMessage()    {}
func (*AlarmOpenOmciEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_941261b445a6569b, []int{7}
}

func (m *AlarmOpenOmciEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmOpenOmciEvent.Unmarshal(m, b)
}
func (m *AlarmOpenOmciEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmOpenOmciEvent.Marshal(b, m, deterministic)
}
func (m *AlarmOpenOmciEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmOpenOmciEvent.Merge(m, src)
}
func (m *AlarmOpenOmciEvent) XXX_Size() int {
	return xxx_messageInfo_AlarmOpenOmciEvent.Size(m)
}
func (m *AlarmOpenOmciEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmOpenOmciEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmOpenOmciEvent proto.InternalMessageInfo

func (m *AlarmOpenOmciEvent) GetType() AlarmOpenOmciEventType_OpenOmciEventType {
	if m != nil {
		return m.Type
	}
	return AlarmOpenOmciEventType_state_change
}

func (m *AlarmOpenOmciEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterEnum("alarm.AlarmOpenOmciEventType_OpenOmciEventType", AlarmOpenOmciEventType_OpenOmciEventType_name, AlarmOpenOmciEventType_OpenOmciEventType_value)
	proto.RegisterType((*AlarmAttributeData)(nil), "alarm.AlarmAttributeData")
	proto.RegisterType((*AlarmInstanceData)(nil), "alarm.AlarmInstanceData")
	proto.RegisterType((*AlarmClassData)(nil), "alarm.AlarmClassData")
	proto.RegisterType((*AlarmManagedEntity)(nil), "alarm.AlarmManagedEntity")
	proto.RegisterType((*AlarmMessageType)(nil), "alarm.AlarmMessageType")
	proto.RegisterType((*AlarmDeviceData)(nil), "alarm.AlarmDeviceData")
	proto.RegisterType((*AlarmOpenOmciEventType)(nil), "alarm.AlarmOpenOmciEventType")
	proto.RegisterType((*AlarmOpenOmciEvent)(nil), "alarm.AlarmOpenOmciEvent")
}

func init() { proto.RegisterFile("omci_alarm_db.proto", fileDescriptor_941261b445a6569b) }

var fileDescriptor_941261b445a6569b = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xae, 0x13, 0x3b, 0x71, 0xc6, 0x76, 0xe2, 0x6c, 0xfa, 0xb3, 0x09, 0x04, 0x82, 0x68, 0x8b,
	0x4b, 0x41, 0x82, 0xf4, 0x1a, 0x68, 0xe3, 0xd8, 0x05, 0x17, 0x4a, 0xa8, 0x92, 0x53, 0x2f, 0x62,
	0x2d, 0x4d, 0xed, 0x05, 0xef, 0xca, 0xd5, 0xae, 0x0d, 0xbe, 0xf4, 0xd6, 0x87, 0xea, 0x43, 0xe4,
	0x25, 0x7a, 0xea, 0x13, 0xe4, 0x5c, 0x34, 0x92, 0x6c, 0x19, 0x41, 0xe9, 0x4d, 0xdf, 0x37, 0x33,
	0xdf, 0xcc, 0xce, 0x37, 0x08, 0x4e, 0x62, 0x15, 0xca, 0x40, 0xcc, 0x44, 0xa2, 0x82, 0x68, 0xec,
	0xce, 0x93, 0xd8, 0xc6, 0xac, 0x41, 0xf8, 0x0c, 0x14, 0x5a, 0x91, 0x51, 0xce, 0x10, 0xd8, 0x75,
	0x4a, 0x5e, 0x5b, 0x9b, 0xc8, 0xf1, 0xc2, 0xe2, 0x40, 0x58, 0xc1, 0x4e, 0xa1, 0xae, 0x85, 0x42,
	0x5e, 0xbb, 0xa8, 0xf5, 0x0e, 0xfa, 0x8d, 0x3f, 0x8f, 0x0f, 0xe7, 0x35, 0x9f, 0x28, 0xf6, 0x14,
	0x1a, 0x4b, 0x31, 0x5b, 0x20, 0xdf, 0x49, 0x63, 0x7e, 0x06, 0x9c, 0x5f, 0x35, 0x38, 0x26, 0x9d,
	0x91, 0x36, 0x56, 0xe8, 0x30, 0x93, 0x79, 0x0d, 0x2d, 0x99, 0xe3, 0x40, 0x46, 0xa4, 0xd6, 0x29,
	0xd4, 0xa0, 0x88, 0x8c, 0x22, 0xc6, 0x61, 0x3f, 0x4c, 0x50, 0x58, 0x8c, 0x72, 0xd5, 0x02, 0xb2,
	0x33, 0x68, 0xaa, 0x38, 0x92, 0xdf, 0x24, 0x46, 0x7c, 0x97, 0x42, 0x6b, 0xcc, 0x3e, 0x02, 0x88,
	0x62, 0x6a, 0xc3, 0xeb, 0x17, 0xbb, 0xbd, 0xd6, 0xe5, 0xa9, 0x4b, 0x4f, 0x74, 0xab, 0x6f, 0xea,
	0xb7, 0x7e, 0x3f, 0x3e, 0x9c, 0xef, 0x65, 0x0f, 0xf3, 0x4b, 0x95, 0xce, 0x0f, 0x38, 0xa4, 0xf4,
	0x9b, 0x99, 0x30, 0x86, 0xe6, 0xbe, 0x80, 0x66, 0x98, 0x82, 0xca, 0xd0, 0xfb, 0x44, 0x8f, 0x22,
	0xf6, 0x09, 0x0e, 0x8a, 0xf9, 0x0d, 0xdf, 0xa1, 0xd6, 0xbc, 0xdc, 0xba, 0xbc, 0x86, 0x3e, 0x4b,
	0x3b, 0x77, 0xb6, 0x76, 0xe1, 0x6f, 0xca, 0x9d, 0x2f, 0xb9, 0x05, 0x9f, 0x85, 0x16, 0x13, 0x8c,
	0x86, 0xda, 0x4a, 0xbb, 0xfa, 0x8f, 0x19, 0x0a, 0x93, 0x76, 0x2a, 0x26, 0x39, 0x57, 0xd0, 0xcd,
	0x24, 0xd1, 0x18, 0x31, 0xc1, 0xfb, 0xd5, 0x1c, 0x59, 0x0f, 0xda, 0x2a, 0x83, 0x81, 0x5d, 0xcd,
	0x71, 0x5b, 0xb4, 0xa5, 0x36, 0x99, 0xce, 0xcf, 0x5d, 0x38, 0xa2, 0xf2, 0x01, 0x2e, 0x65, 0x6e,
	0xa5, 0x03, 0x07, 0x11, 0xa1, 0x62, 0x9e, 0x75, 0xc7, 0x66, 0xc6, 0xff, 0xd3, 0x46, 0x17, 0x4e,
	0x66, 0xc2, 0xd8, 0xfc, 0x1e, 0x0d, 0x7e, 0x5f, 0xa0, 0x0e, 0x91, 0x1c, 0xed, 0xf8, 0xc7, 0x69,
	0x88, 0xfa, 0xdd, 0xe5, 0x01, 0xf6, 0x12, 0x0e, 0x29, 0xdf, 0xac, 0x74, 0x18, 0x58, 0xa9, 0x90,
	0xd7, 0x49, 0xb0, 0x9d, 0xb2, 0x77, 0x2b, 0x1d, 0xde, 0x4b, 0x85, 0x69, 0xbf, 0x25, 0x26, 0x46,
	0xc6, 0x9a, 0x37, 0x48, 0xa9, 0x80, 0xec, 0x03, 0x64, 0x5b, 0x42, 0xc3, 0xf7, 0xc8, 0x9c, 0x67,
	0x65, 0x73, 0xd6, 0x46, 0xf7, 0x8f, 0x52, 0x67, 0x60, 0xb3, 0x69, 0xbf, 0x28, 0x63, 0x03, 0xe8,
	0xaa, 0xcc, 0x8f, 0x00, 0x53, 0x43, 0x24, 0x1a, 0xbe, 0x5f, 0x3d, 0xb1, 0x2d, 0xcf, 0xfc, 0x23,
	0x55, 0x82, 0x12, 0x0d, 0xbb, 0x82, 0x4e, 0x79, 0xe7, 0x86, 0x37, 0x49, 0xe2, 0xc5, 0x96, 0xc4,
	0x66, 0xf3, 0x7e, 0xbb, 0x64, 0x83, 0x71, 0xde, 0xc3, 0x73, 0xca, 0xb8, 0x9d, 0xa3, 0xbe, 0x55,
	0xa1, 0x1c, 0x2e, 0x51, 0x5b, 0x72, 0xe8, 0x15, 0x1c, 0x57, 0x48, 0xd6, 0x85, 0xb6, 0xb1, 0xc2,
	0x62, 0x10, 0x4e, 0x85, 0x9e, 0x60, 0xf7, 0x89, 0xa3, 0xf2, 0xcb, 0xda, 0xca, 0x65, 0x37, 0x50,
	0x5f, 0x1f, 0xc0, 0xe1, 0xa5, 0x57, 0x9e, 0xa5, 0x22, 0xea, 0x56, 0x18, 0x9f, 0x8a, 0x19, 0x83,
	0x7a, 0x24, 0xac, 0xc8, 0x8d, 0xa6, 0xef, 0xfe, 0xdb, 0xaf, 0x6f, 0x26, 0xd2, 0x4e, 0x17, 0x63,
	0x37, 0x8c, 0x95, 0x17, 0xcf, 0x51, 0x87, 0x71, 0x12, 0x79, 0xcb, 0x78, 0x66, 0xa7, 0xc2, 0xa3,
	0xff, 0x8d, 0xf1, 0x26, 0xb1, 0x97, 0xfe, 0x9a, 0xc6, 0x7b, 0x84, 0xdf, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x2a, 0x42, 0x5c, 0xa9, 0x04, 0x00, 0x00,
}

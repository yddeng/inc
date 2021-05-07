// Code generated by protoc-gen-go. DO NOT EDIT.
// source: intun.proto

package net

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

// 内网服务器注册到中心服务器
type RegisterReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ID                   uint32   `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RegisterResp) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// 控制台连接内网服务器
type ConnectionReq struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionReq) Reset()         { *m = ConnectionReq{} }
func (m *ConnectionReq) String() string { return proto.CompactTextString(m) }
func (*ConnectionReq) ProtoMessage()    {}
func (*ConnectionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{2}
}

func (m *ConnectionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionReq.Unmarshal(m, b)
}
func (m *ConnectionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionReq.Marshal(b, m, deterministic)
}
func (m *ConnectionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionReq.Merge(m, src)
}
func (m *ConnectionReq) XXX_Size() int {
	return xxx_messageInfo_ConnectionReq.Size(m)
}
func (m *ConnectionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionReq proto.InternalMessageInfo

func (m *ConnectionReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ConnectionResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ID                   uint32   `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionResp) Reset()         { *m = ConnectionResp{} }
func (m *ConnectionResp) String() string { return proto.CompactTextString(m) }
func (*ConnectionResp) ProtoMessage()    {}
func (*ConnectionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{3}
}

func (m *ConnectionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResp.Unmarshal(m, b)
}
func (m *ConnectionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResp.Marshal(b, m, deterministic)
}
func (m *ConnectionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResp.Merge(m, src)
}
func (m *ConnectionResp) XXX_Size() int {
	return xxx_messageInfo_ConnectionResp.Size(m)
}
func (m *ConnectionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResp proto.InternalMessageInfo

func (m *ConnectionResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ConnectionResp) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// 控制台命令到中心服务器
type CommandReq struct {
	Cmd                  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandReq) Reset()         { *m = CommandReq{} }
func (m *CommandReq) String() string { return proto.CompactTextString(m) }
func (*CommandReq) ProtoMessage()    {}
func (*CommandReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{4}
}

func (m *CommandReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandReq.Unmarshal(m, b)
}
func (m *CommandReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandReq.Marshal(b, m, deterministic)
}
func (m *CommandReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandReq.Merge(m, src)
}
func (m *CommandReq) XXX_Size() int {
	return xxx_messageInfo_CommandReq.Size(m)
}
func (m *CommandReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandReq.DiscardUnknown(m)
}

var xxx_messageInfo_CommandReq proto.InternalMessageInfo

func (m *CommandReq) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

type CommandResp struct {
	Ret                  string   `protobuf:"bytes,2,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResp) Reset()         { *m = CommandResp{} }
func (m *CommandResp) String() string { return proto.CompactTextString(m) }
func (*CommandResp) ProtoMessage()    {}
func (*CommandResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{5}
}

func (m *CommandResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResp.Unmarshal(m, b)
}
func (m *CommandResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResp.Marshal(b, m, deterministic)
}
func (m *CommandResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResp.Merge(m, src)
}
func (m *CommandResp) XXX_Size() int {
	return xxx_messageInfo_CommandResp.Size(m)
}
func (m *CommandResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResp.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResp proto.InternalMessageInfo

func (m *CommandResp) GetRet() string {
	if m != nil {
		return m.Ret
	}
	return ""
}

// 心跳
type Heartbeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{6}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

// 隧道消息
type TunMsg struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TunMsg) Reset()         { *m = TunMsg{} }
func (m *TunMsg) String() string { return proto.CompactTextString(m) }
func (*TunMsg) ProtoMessage()    {}
func (*TunMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_76f4cf0dc0f704f9, []int{7}
}

func (m *TunMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunMsg.Unmarshal(m, b)
}
func (m *TunMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunMsg.Marshal(b, m, deterministic)
}
func (m *TunMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunMsg.Merge(m, src)
}
func (m *TunMsg) XXX_Size() int {
	return xxx_messageInfo_TunMsg.Size(m)
}
func (m *TunMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TunMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TunMsg proto.InternalMessageInfo

func (m *TunMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "register_req")
	proto.RegisterType((*RegisterResp)(nil), "register_resp")
	proto.RegisterType((*ConnectionReq)(nil), "connection_req")
	proto.RegisterType((*ConnectionResp)(nil), "connection_resp")
	proto.RegisterType((*CommandReq)(nil), "command_req")
	proto.RegisterType((*CommandResp)(nil), "command_resp")
	proto.RegisterType((*Heartbeat)(nil), "heartbeat")
	proto.RegisterType((*TunMsg)(nil), "tun_msg")
}

func init() { proto.RegisterFile("intun.proto", fileDescriptor_76f4cf0dc0f704f9) }

var fileDescriptor_76f4cf0dc0f704f9 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0x47, 0xb1, 0x9d, 0xfe, 0xf1, 0xd9, 0x49, 0x83, 0xa6, 0x50, 0x28, 0x35, 0x37, 0x65, 0x28,
	0x85, 0x92, 0x6f, 0x50, 0xb2, 0x64, 0xd5, 0xd8, 0x25, 0x28, 0xf2, 0x91, 0x7a, 0xd0, 0xc9, 0x95,
	0xce, 0xf4, 0xeb, 0x17, 0xa9, 0xee, 0xbf, 0xad, 0xdb, 0x13, 0x3c, 0xbd, 0x1f, 0x1c, 0x34, 0x03,
	0xcb, 0xc4, 0x8f, 0x63, 0xf0, 0xe2, 0x11, 0xa1, 0x0d, 0x74, 0x1e, 0xa2, 0x50, 0x38, 0x06, 0x7a,
	0x53, 0x0a, 0x16, 0x6c, 0x1c, 0x6d, 0x8a, 0xae, 0xd8, 0xd6, 0x3a, 0x33, 0x3e, 0xc1, 0xf2, 0x97,
	0x13, 0x47, 0xb5, 0x86, 0xca, 0xc5, 0xf3, 0xec, 0x24, 0x54, 0x2b, 0x28, 0x0f, 0xfb, 0x4d, 0xd9,
	0x15, 0xdb, 0xa5, 0x2e, 0x0f, 0x7b, 0x7c, 0x80, 0x95, 0xf5, 0xcc, 0x64, 0x65, 0xf0, 0x9c, 0xc3,
	0xb7, 0x70, 0x3d, 0x9a, 0x18, 0xdf, 0x7d, 0xe8, 0xe7, 0x8f, 0xdf, 0x6f, 0xdc, 0xc1, 0xcd, 0x1f,
	0xfb, 0x5f, 0x13, 0xf7, 0xd0, 0x58, 0xef, 0x9c, 0xe1, 0x3e, 0xf7, 0xd7, 0x50, 0x59, 0xf7, 0x95,
	0x4e, 0x88, 0x1d, 0xb4, 0x3f, 0xc2, 0x67, 0x32, 0x90, 0xe4, 0x42, 0xad, 0x13, 0x62, 0x03, 0xf5,
	0x2b, 0x99, 0x20, 0x27, 0x32, 0x82, 0x77, 0x70, 0x25, 0x13, 0x1f, 0xd3, 0x94, 0x82, 0x45, 0x6f,
	0xc4, 0xe4, 0x58, 0xab, 0x33, 0x3f, 0x5f, 0xbc, 0x54, 0x4c, 0x72, 0xba, 0xcc, 0x67, 0xdb, 0x7d,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x82, 0x14, 0x8a, 0x45, 0x01, 0x00, 0x00,
}
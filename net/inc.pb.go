// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inc.proto

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

// 内网服务器注册
type LoginReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LoginResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LoginResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 命令行客户端连接
type AuthReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{2}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return xxx_messageInfo_AuthReq.Size(m)
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResp) Reset()         { *m = AuthResp{} }
func (m *AuthResp) String() string { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()    {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{3}
}

func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResp.Unmarshal(m, b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
}
func (m *AuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResp.Merge(m, src)
}
func (m *AuthResp) XXX_Size() int {
	return xxx_messageInfo_AuthResp.Size(m)
}
func (m *AuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResp proto.InternalMessageInfo

func (m *AuthResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AuthResp) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 代理映射
type Mapping struct {
	InternalIp           string   `protobuf:"bytes,1,opt,name=internal_ip,json=internalIp,proto3" json:"internal_ip,omitempty"`
	InternalPort         int32    `protobuf:"varint,2,opt,name=internal_port,json=internalPort,proto3" json:"internal_port,omitempty"`
	ExternalPort         int32    `protobuf:"varint,3,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MapId                uint32   `protobuf:"varint,5,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	SlaveId              uint32   `protobuf:"varint,6,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mapping) Reset()         { *m = Mapping{} }
func (m *Mapping) String() string { return proto.CompactTextString(m) }
func (*Mapping) ProtoMessage()    {}
func (*Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{4}
}

func (m *Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mapping.Unmarshal(m, b)
}
func (m *Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mapping.Marshal(b, m, deterministic)
}
func (m *Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mapping.Merge(m, src)
}
func (m *Mapping) XXX_Size() int {
	return xxx_messageInfo_Mapping.Size(m)
}
func (m *Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_Mapping proto.InternalMessageInfo

func (m *Mapping) GetInternalIp() string {
	if m != nil {
		return m.InternalIp
	}
	return ""
}

func (m *Mapping) GetInternalPort() int32 {
	if m != nil {
		return m.InternalPort
	}
	return 0
}

func (m *Mapping) GetExternalPort() int32 {
	if m != nil {
		return m.ExternalPort
	}
	return 0
}

func (m *Mapping) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Mapping) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *Mapping) GetSlaveId() uint32 {
	if m != nil {
		return m.SlaveId
	}
	return 0
}

// 注册代理映射
type RegisterReq struct {
	Maps                 *Mapping `protobuf:"bytes,1,opt,name=maps,proto3" json:"maps,omitempty"`
	SlaveId              uint32   `protobuf:"varint,2,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{5}
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

func (m *RegisterReq) GetMaps() *Mapping {
	if m != nil {
		return m.Maps
	}
	return nil
}

func (m *RegisterReq) GetSlaveId() uint32 {
	if m != nil {
		return m.SlaveId
	}
	return 0
}

type RegisterResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{6}
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

// 注销代理映射
type UnregisterReq struct {
	MapId                uint32   `protobuf:"varint,1,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterReq) Reset()         { *m = UnregisterReq{} }
func (m *UnregisterReq) String() string { return proto.CompactTextString(m) }
func (*UnregisterReq) ProtoMessage()    {}
func (*UnregisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{7}
}

func (m *UnregisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterReq.Unmarshal(m, b)
}
func (m *UnregisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterReq.Marshal(b, m, deterministic)
}
func (m *UnregisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterReq.Merge(m, src)
}
func (m *UnregisterReq) XXX_Size() int {
	return xxx_messageInfo_UnregisterReq.Size(m)
}
func (m *UnregisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterReq proto.InternalMessageInfo

func (m *UnregisterReq) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type UnregisterResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterResp) Reset()         { *m = UnregisterResp{} }
func (m *UnregisterResp) String() string { return proto.CompactTextString(m) }
func (*UnregisterResp) ProtoMessage()    {}
func (*UnregisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{8}
}

func (m *UnregisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterResp.Unmarshal(m, b)
}
func (m *UnregisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterResp.Marshal(b, m, deterministic)
}
func (m *UnregisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterResp.Merge(m, src)
}
func (m *UnregisterResp) XXX_Size() int {
	return xxx_messageInfo_UnregisterResp.Size(m)
}
func (m *UnregisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterResp proto.InternalMessageInfo

func (m *UnregisterResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 创建映射连接器
type CreateDialerReq struct {
	MapId                uint32   `protobuf:"varint,1,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDialerReq) Reset()         { *m = CreateDialerReq{} }
func (m *CreateDialerReq) String() string { return proto.CompactTextString(m) }
func (*CreateDialerReq) ProtoMessage()    {}
func (*CreateDialerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{9}
}

func (m *CreateDialerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDialerReq.Unmarshal(m, b)
}
func (m *CreateDialerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDialerReq.Marshal(b, m, deterministic)
}
func (m *CreateDialerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDialerReq.Merge(m, src)
}
func (m *CreateDialerReq) XXX_Size() int {
	return xxx_messageInfo_CreateDialerReq.Size(m)
}
func (m *CreateDialerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDialerReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDialerReq proto.InternalMessageInfo

func (m *CreateDialerReq) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *CreateDialerReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateDialerResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDialerResp) Reset()         { *m = CreateDialerResp{} }
func (m *CreateDialerResp) String() string { return proto.CompactTextString(m) }
func (*CreateDialerResp) ProtoMessage()    {}
func (*CreateDialerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{10}
}

func (m *CreateDialerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDialerResp.Unmarshal(m, b)
}
func (m *CreateDialerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDialerResp.Marshal(b, m, deterministic)
}
func (m *CreateDialerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDialerResp.Merge(m, src)
}
func (m *CreateDialerResp) XXX_Size() int {
	return xxx_messageInfo_CreateDialerResp.Size(m)
}
func (m *CreateDialerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDialerResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDialerResp proto.InternalMessageInfo

func (m *CreateDialerResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 销毁映射连接器
type DestroyDialerReq struct {
	MapId                uint32   `protobuf:"varint,1,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyDialerReq) Reset()         { *m = DestroyDialerReq{} }
func (m *DestroyDialerReq) String() string { return proto.CompactTextString(m) }
func (*DestroyDialerReq) ProtoMessage()    {}
func (*DestroyDialerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{11}
}

func (m *DestroyDialerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyDialerReq.Unmarshal(m, b)
}
func (m *DestroyDialerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyDialerReq.Marshal(b, m, deterministic)
}
func (m *DestroyDialerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyDialerReq.Merge(m, src)
}
func (m *DestroyDialerReq) XXX_Size() int {
	return xxx_messageInfo_DestroyDialerReq.Size(m)
}
func (m *DestroyDialerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyDialerReq.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyDialerReq proto.InternalMessageInfo

func (m *DestroyDialerReq) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type DestroyDialerResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyDialerResp) Reset()         { *m = DestroyDialerResp{} }
func (m *DestroyDialerResp) String() string { return proto.CompactTextString(m) }
func (*DestroyDialerResp) ProtoMessage()    {}
func (*DestroyDialerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{12}
}

func (m *DestroyDialerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyDialerResp.Unmarshal(m, b)
}
func (m *DestroyDialerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyDialerResp.Marshal(b, m, deterministic)
}
func (m *DestroyDialerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyDialerResp.Merge(m, src)
}
func (m *DestroyDialerResp) XXX_Size() int {
	return xxx_messageInfo_DestroyDialerResp.Size(m)
}
func (m *DestroyDialerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyDialerResp.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyDialerResp proto.InternalMessageInfo

func (m *DestroyDialerResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 开启通信隧道
type OpenChannelReq struct {
	ChannelId            uint32   `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MapId                uint32   `protobuf:"varint,2,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenChannelReq) Reset()         { *m = OpenChannelReq{} }
func (m *OpenChannelReq) String() string { return proto.CompactTextString(m) }
func (*OpenChannelReq) ProtoMessage()    {}
func (*OpenChannelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{13}
}

func (m *OpenChannelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelReq.Unmarshal(m, b)
}
func (m *OpenChannelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelReq.Marshal(b, m, deterministic)
}
func (m *OpenChannelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelReq.Merge(m, src)
}
func (m *OpenChannelReq) XXX_Size() int {
	return xxx_messageInfo_OpenChannelReq.Size(m)
}
func (m *OpenChannelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelReq.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelReq proto.InternalMessageInfo

func (m *OpenChannelReq) GetChannelId() uint32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *OpenChannelReq) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type OpenChannelResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenChannelResp) Reset()         { *m = OpenChannelResp{} }
func (m *OpenChannelResp) String() string { return proto.CompactTextString(m) }
func (*OpenChannelResp) ProtoMessage()    {}
func (*OpenChannelResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{14}
}

func (m *OpenChannelResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenChannelResp.Unmarshal(m, b)
}
func (m *OpenChannelResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenChannelResp.Marshal(b, m, deterministic)
}
func (m *OpenChannelResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenChannelResp.Merge(m, src)
}
func (m *OpenChannelResp) XXX_Size() int {
	return xxx_messageInfo_OpenChannelResp.Size(m)
}
func (m *OpenChannelResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenChannelResp.DiscardUnknown(m)
}

var xxx_messageInfo_OpenChannelResp proto.InternalMessageInfo

func (m *OpenChannelResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 关闭通信隧道
type CloseChannelReq struct {
	ChannelId            uint32   `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseChannelReq) Reset()         { *m = CloseChannelReq{} }
func (m *CloseChannelReq) String() string { return proto.CompactTextString(m) }
func (*CloseChannelReq) ProtoMessage()    {}
func (*CloseChannelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{15}
}

func (m *CloseChannelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseChannelReq.Unmarshal(m, b)
}
func (m *CloseChannelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseChannelReq.Marshal(b, m, deterministic)
}
func (m *CloseChannelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseChannelReq.Merge(m, src)
}
func (m *CloseChannelReq) XXX_Size() int {
	return xxx_messageInfo_CloseChannelReq.Size(m)
}
func (m *CloseChannelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseChannelReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseChannelReq proto.InternalMessageInfo

func (m *CloseChannelReq) GetChannelId() uint32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

type CloseChannelResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseChannelResp) Reset()         { *m = CloseChannelResp{} }
func (m *CloseChannelResp) String() string { return proto.CompactTextString(m) }
func (*CloseChannelResp) ProtoMessage()    {}
func (*CloseChannelResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{16}
}

func (m *CloseChannelResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseChannelResp.Unmarshal(m, b)
}
func (m *CloseChannelResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseChannelResp.Marshal(b, m, deterministic)
}
func (m *CloseChannelResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseChannelResp.Merge(m, src)
}
func (m *CloseChannelResp) XXX_Size() int {
	return xxx_messageInfo_CloseChannelResp.Size(m)
}
func (m *CloseChannelResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseChannelResp.DiscardUnknown(m)
}

var xxx_messageInfo_CloseChannelResp proto.InternalMessageInfo

func (m *CloseChannelResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 隧道消息
type ChannelMessageReq struct {
	ChannelId            uint32   `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelMessageReq) Reset()         { *m = ChannelMessageReq{} }
func (m *ChannelMessageReq) String() string { return proto.CompactTextString(m) }
func (*ChannelMessageReq) ProtoMessage()    {}
func (*ChannelMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{17}
}

func (m *ChannelMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelMessageReq.Unmarshal(m, b)
}
func (m *ChannelMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelMessageReq.Marshal(b, m, deterministic)
}
func (m *ChannelMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelMessageReq.Merge(m, src)
}
func (m *ChannelMessageReq) XXX_Size() int {
	return xxx_messageInfo_ChannelMessageReq.Size(m)
}
func (m *ChannelMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelMessageReq proto.InternalMessageInfo

func (m *ChannelMessageReq) GetChannelId() uint32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *ChannelMessageReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ChannelMessageResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelMessageResp) Reset()         { *m = ChannelMessageResp{} }
func (m *ChannelMessageResp) String() string { return proto.CompactTextString(m) }
func (*ChannelMessageResp) ProtoMessage()    {}
func (*ChannelMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{18}
}

func (m *ChannelMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelMessageResp.Unmarshal(m, b)
}
func (m *ChannelMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelMessageResp.Marshal(b, m, deterministic)
}
func (m *ChannelMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelMessageResp.Merge(m, src)
}
func (m *ChannelMessageResp) XXX_Size() int {
	return xxx_messageInfo_ChannelMessageResp.Size(m)
}
func (m *ChannelMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelMessageResp proto.InternalMessageInfo

func (m *ChannelMessageResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// command line
type ClientCmdReq struct {
	Cmd                  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCmdReq) Reset()         { *m = ClientCmdReq{} }
func (m *ClientCmdReq) String() string { return proto.CompactTextString(m) }
func (*ClientCmdReq) ProtoMessage()    {}
func (*ClientCmdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{19}
}

func (m *ClientCmdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCmdReq.Unmarshal(m, b)
}
func (m *ClientCmdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCmdReq.Marshal(b, m, deterministic)
}
func (m *ClientCmdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCmdReq.Merge(m, src)
}
func (m *ClientCmdReq) XXX_Size() int {
	return xxx_messageInfo_ClientCmdReq.Size(m)
}
func (m *ClientCmdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCmdReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCmdReq proto.InternalMessageInfo

func (m *ClientCmdReq) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

type ClientCmdResp struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientCmdResp) Reset()         { *m = ClientCmdResp{} }
func (m *ClientCmdResp) String() string { return proto.CompactTextString(m) }
func (*ClientCmdResp) ProtoMessage()    {}
func (*ClientCmdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{20}
}

func (m *ClientCmdResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientCmdResp.Unmarshal(m, b)
}
func (m *ClientCmdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientCmdResp.Marshal(b, m, deterministic)
}
func (m *ClientCmdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientCmdResp.Merge(m, src)
}
func (m *ClientCmdResp) XXX_Size() int {
	return xxx_messageInfo_ClientCmdResp.Size(m)
}
func (m *ClientCmdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientCmdResp.DiscardUnknown(m)
}

var xxx_messageInfo_ClientCmdResp proto.InternalMessageInfo

func (m *ClientCmdResp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "login_req")
	proto.RegisterType((*LoginResp)(nil), "login_resp")
	proto.RegisterType((*AuthReq)(nil), "auth_req")
	proto.RegisterType((*AuthResp)(nil), "auth_resp")
	proto.RegisterType((*Mapping)(nil), "mapping")
	proto.RegisterType((*RegisterReq)(nil), "register_req")
	proto.RegisterType((*RegisterResp)(nil), "register_resp")
	proto.RegisterType((*UnregisterReq)(nil), "unregister_req")
	proto.RegisterType((*UnregisterResp)(nil), "unregister_resp")
	proto.RegisterType((*CreateDialerReq)(nil), "create_dialer_req")
	proto.RegisterType((*CreateDialerResp)(nil), "create_dialer_resp")
	proto.RegisterType((*DestroyDialerReq)(nil), "destroy_dialer_req")
	proto.RegisterType((*DestroyDialerResp)(nil), "destroy_dialer_resp")
	proto.RegisterType((*OpenChannelReq)(nil), "open_channel_req")
	proto.RegisterType((*OpenChannelResp)(nil), "open_channel_resp")
	proto.RegisterType((*CloseChannelReq)(nil), "close_channel_req")
	proto.RegisterType((*CloseChannelResp)(nil), "close_channel_resp")
	proto.RegisterType((*ChannelMessageReq)(nil), "channel_message_req")
	proto.RegisterType((*ChannelMessageResp)(nil), "channel_message_resp")
	proto.RegisterType((*ClientCmdReq)(nil), "client_cmd_req")
	proto.RegisterType((*ClientCmdResp)(nil), "client_cmd_resp")
}

func init() { proto.RegisterFile("inc.proto", fileDescriptor_a9ee4de5949f8e83) }

var fileDescriptor_a9ee4de5949f8e83 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x9c, 0x3d, 0x39, 0x34, 0xd9, 0x16, 0xc9, 0x48, 0x20, 0xc2, 0x56, 0xd0, 0x48,
	0x88, 0x20, 0x95, 0x37, 0x40, 0x48, 0x90, 0x3b, 0xe4, 0x4b, 0x6e, 0xac, 0xc5, 0x3b, 0x72, 0x57,
	0x78, 0x0f, 0xec, 0x6e, 0x11, 0xbc, 0x21, 0x8f, 0x85, 0xb2, 0xb5, 0x53, 0x3b, 0x31, 0x50, 0xee,
	0x76, 0x7e, 0xff, 0xf3, 0xfd, 0x33, 0x73, 0x61, 0x88, 0x85, 0xca, 0xb7, 0xc6, 0x6a, 0xaf, 0xe9,
	0x1b, 0x88, 0x4b, 0x5d, 0x08, 0x95, 0x59, 0xfc, 0x46, 0x08, 0x0c, 0x14, 0x93, 0x98, 0x44, 0xeb,
	0x68, 0x13, 0xa7, 0xe1, 0x4d, 0x16, 0xd0, 0x13, 0x3c, 0xe9, 0xad, 0xa3, 0xcd, 0x3c, 0xed, 0x09,
	0x4e, 0xb7, 0x00, 0x75, 0x83, 0x33, 0x64, 0x09, 0x7d, 0xe9, 0x8a, 0xaa, 0x61, 0xff, 0x3c, 0xf1,
	0xaf, 0x61, 0xc2, 0x6e, 0xfd, 0x4d, 0xe0, 0x5f, 0xc0, 0xd0, 0xeb, 0xaf, 0xa8, 0x2a, 0xff, 0x5d,
	0x41, 0x5f, 0x43, 0x5c, 0x39, 0x1e, 0x04, 0xfc, 0x15, 0xc1, 0x58, 0x32, 0x63, 0x84, 0x2a, 0xc8,
	0x33, 0x98, 0x0a, 0xe5, 0xd1, 0x2a, 0x56, 0x66, 0xc2, 0x54, 0x5d, 0x50, 0x4b, 0x3b, 0x43, 0x2e,
	0x61, 0x7e, 0x30, 0x18, 0x6d, 0x7d, 0xe0, 0x0c, 0xd3, 0x59, 0x2d, 0x7e, 0xd2, 0xd6, 0xef, 0x4d,
	0xf8, 0xa3, 0x69, 0xea, 0xdf, 0x99, 0x6a, 0x31, 0x98, 0xd6, 0x30, 0xe5, 0xe8, 0x72, 0x2b, 0x8c,
	0x17, 0x5a, 0x25, 0x83, 0x10, 0xd5, 0x94, 0xc8, 0x23, 0x18, 0x49, 0x66, 0x32, 0xc1, 0x93, 0x61,
	0x18, 0x76, 0x28, 0x99, 0xd9, 0x71, 0xf2, 0x18, 0x26, 0xae, 0x64, 0xdf, 0x71, 0xff, 0x61, 0x14,
	0x3e, 0x8c, 0x43, 0xbd, 0xe3, 0xf4, 0x03, 0xcc, 0x2c, 0x16, 0xc2, 0x79, 0xb4, 0xe1, 0x3e, 0x4f,
	0x60, 0x20, 0x99, 0x71, 0x61, 0x8f, 0xe9, 0xf5, 0x64, 0x5b, 0xad, 0x99, 0x06, 0xb5, 0x05, 0xea,
	0xb5, 0x41, 0xcf, 0x61, 0xde, 0x00, 0x75, 0x9d, 0x91, 0x5e, 0xc1, 0xe2, 0x56, 0xb5, 0xd2, 0xee,
	0xe7, 0x8d, 0x1a, 0xf3, 0xd2, 0x4b, 0x38, 0x6b, 0x19, 0x3b, 0x69, 0xef, 0x61, 0x95, 0x5b, 0x64,
	0x1e, 0x33, 0x2e, 0x58, 0xf9, 0x57, 0x20, 0x49, 0x60, 0xcc, 0x38, 0xb7, 0xe8, 0x5c, 0x18, 0x3b,
	0x4e, 0xeb, 0x92, 0xbe, 0x04, 0x72, 0x4c, 0xe9, 0x4c, 0x7b, 0x05, 0x84, 0xa3, 0xf3, 0x56, 0xff,
	0xfc, 0x77, 0x1c, 0xbd, 0x82, 0xf3, 0x13, 0x73, 0x27, 0xf5, 0x23, 0x2c, 0xb5, 0x41, 0x95, 0xe5,
	0x37, 0x4c, 0x29, 0x2c, 0x03, 0xf3, 0x29, 0x40, 0x5d, 0x1e, 0xb8, 0x71, 0xa5, 0xec, 0x78, 0x23,
	0xb2, 0xd7, 0x8c, 0x7c, 0x01, 0xab, 0x23, 0x52, 0x67, 0xe0, 0x35, 0xac, 0xf2, 0x52, 0x3b, 0xfc,
	0x8f, 0xc4, 0x70, 0xa2, 0xa3, 0x9e, 0x3f, 0x2c, 0x73, 0x5e, 0x3b, 0x24, 0x3a, 0xc7, 0x0a, 0x7c,
	0xc8, 0x3e, 0x04, 0x06, 0x9c, 0x79, 0x16, 0xb6, 0x99, 0xa5, 0xe1, 0x4d, 0x37, 0x70, 0x71, 0x4a,
	0xea, 0xcc, 0xa4, 0xb0, 0xc8, 0x4b, 0x81, 0xca, 0x67, 0xb9, 0xe4, 0x21, 0x6e, 0x09, 0xfd, 0x5c,
	0xf2, 0xda, 0x93, 0xcb, 0xfd, 0x69, 0xce, 0x5a, 0x1e, 0x67, 0x0e, 0xa1, 0xd1, 0x7d, 0xe8, 0xbb,
	0xe1, 0xe7, 0xbe, 0x42, 0xff, 0x65, 0x14, 0x7e, 0x4a, 0x6f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x02, 0xf9, 0xb1, 0x5a, 0xa1, 0x04, 0x00, 0x00,
}

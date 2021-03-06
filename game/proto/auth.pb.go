// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UUID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Session struct {
	AccessToken          *UUID    `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	UserUID              *UUID    `protobuf:"bytes,2,opt,name=userUID,proto3" json:"userUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetAccessToken() *UUID {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *Session) GetUserUID() *UUID {
	if m != nil {
		return m.UserUID
	}
	return nil
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

type User struct {
	Uid                  *UUID    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Points               int32    `protobuf:"varint,5,opt,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUid() *UUID {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetPoints() int32 {
	if m != nil {
		return m.Points
	}
	return 0
}

type UserSignUp struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSignUp) Reset()         { *m = UserSignUp{} }
func (m *UserSignUp) String() string { return proto.CompactTextString(m) }
func (*UserSignUp) ProtoMessage()    {}
func (*UserSignUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *UserSignUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSignUp.Unmarshal(m, b)
}
func (m *UserSignUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSignUp.Marshal(b, m, deterministic)
}
func (m *UserSignUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSignUp.Merge(m, src)
}
func (m *UserSignUp) XXX_Size() int {
	return xxx_messageInfo_UserSignUp.Size(m)
}
func (m *UserSignUp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSignUp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSignUp proto.InternalMessageInfo

func (m *UserSignUp) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserSignUp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserSignUp) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserSignIn struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSignIn) Reset()         { *m = UserSignIn{} }
func (m *UserSignIn) String() string { return proto.CompactTextString(m) }
func (*UserSignIn) ProtoMessage()    {}
func (*UserSignIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *UserSignIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSignIn.Unmarshal(m, b)
}
func (m *UserSignIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSignIn.Marshal(b, m, deterministic)
}
func (m *UserSignIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSignIn.Merge(m, src)
}
func (m *UserSignIn) XXX_Size() int {
	return xxx_messageInfo_UserSignIn.Size(m)
}
func (m *UserSignIn) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSignIn.DiscardUnknown(m)
}

var xxx_messageInfo_UserSignIn proto.InternalMessageInfo

func (m *UserSignIn) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserSignIn) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LeadersRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeadersRequest) Reset()         { *m = LeadersRequest{} }
func (m *LeadersRequest) String() string { return proto.CompactTextString(m) }
func (*LeadersRequest) ProtoMessage()    {}
func (*LeadersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *LeadersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeadersRequest.Unmarshal(m, b)
}
func (m *LeadersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeadersRequest.Marshal(b, m, deterministic)
}
func (m *LeadersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeadersRequest.Merge(m, src)
}
func (m *LeadersRequest) XXX_Size() int {
	return xxx_messageInfo_LeadersRequest.Size(m)
}
func (m *LeadersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeadersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeadersRequest proto.InternalMessageInfo

func (m *LeadersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *LeadersRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type LeadersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeadersResponse) Reset()         { *m = LeadersResponse{} }
func (m *LeadersResponse) String() string { return proto.CompactTextString(m) }
func (*LeadersResponse) ProtoMessage()    {}
func (*LeadersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *LeadersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeadersResponse.Unmarshal(m, b)
}
func (m *LeadersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeadersResponse.Marshal(b, m, deterministic)
}
func (m *LeadersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeadersResponse.Merge(m, src)
}
func (m *LeadersResponse) XXX_Size() int {
	return xxx_messageInfo_LeadersResponse.Size(m)
}
func (m *LeadersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeadersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeadersResponse proto.InternalMessageInfo

func (m *LeadersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *LeadersResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type AddAvatarRequest struct {
	UserUID              *UUID    `protobuf:"bytes,1,opt,name=userUID,proto3" json:"userUID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAvatarRequest) Reset()         { *m = AddAvatarRequest{} }
func (m *AddAvatarRequest) String() string { return proto.CompactTextString(m) }
func (*AddAvatarRequest) ProtoMessage()    {}
func (*AddAvatarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *AddAvatarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAvatarRequest.Unmarshal(m, b)
}
func (m *AddAvatarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAvatarRequest.Marshal(b, m, deterministic)
}
func (m *AddAvatarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAvatarRequest.Merge(m, src)
}
func (m *AddAvatarRequest) XXX_Size() int {
	return xxx_messageInfo_AddAvatarRequest.Size(m)
}
func (m *AddAvatarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAvatarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAvatarRequest proto.InternalMessageInfo

func (m *AddAvatarRequest) GetUserUID() *UUID {
	if m != nil {
		return m.UserUID
	}
	return nil
}

func (m *AddAvatarRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*UUID)(nil), "proto.UUID")
	proto.RegisterType((*Session)(nil), "proto.Session")
	proto.RegisterType((*Nothing)(nil), "proto.Nothing")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserSignUp)(nil), "proto.UserSignUp")
	proto.RegisterType((*UserSignIn)(nil), "proto.UserSignIn")
	proto.RegisterType((*LeadersRequest)(nil), "proto.LeadersRequest")
	proto.RegisterType((*LeadersResponse)(nil), "proto.LeadersResponse")
	proto.RegisterType((*AddAvatarRequest)(nil), "proto.AddAvatarRequest")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x51, 0x8f, 0xd2, 0x4c,
	0x14, 0xa5, 0x40, 0xe9, 0xf6, 0x92, 0x8f, 0x4f, 0x27, 0xba, 0x36, 0x44, 0x13, 0x1c, 0x63, 0x24,
	0x1b, 0x97, 0x07, 0x4c, 0x7c, 0xf0, 0xc1, 0x04, 0xe4, 0xa5, 0x46, 0x7d, 0x28, 0xf6, 0xc5, 0x17,
	0x33, 0xd2, 0xbb, 0xbb, 0x93, 0x85, 0x99, 0xda, 0x69, 0x31, 0xfc, 0x02, 0x7f, 0x8e, 0x3f, 0xc1,
	0xbf, 0x66, 0x3a, 0x33, 0xd4, 0x6e, 0x61, 0xd7, 0x27, 0x38, 0xf7, 0x9e, 0x9e, 0x7b, 0xee, 0x9c,
	0x0b, 0xc0, 0x8a, 0xfc, 0x6a, 0x92, 0x66, 0x32, 0x97, 0xc4, 0xd5, 0x3f, 0xf4, 0x31, 0x74, 0xe3,
	0x38, 0x5c, 0x90, 0x07, 0xe0, 0x6e, 0xd9, 0xba, 0xc0, 0xc0, 0x19, 0x39, 0x63, 0x3f, 0x32, 0x80,
	0x7e, 0x05, 0x6f, 0x89, 0x4a, 0x71, 0x29, 0xc8, 0x39, 0xf4, 0xd9, 0x6a, 0x85, 0x4a, 0x7d, 0x96,
	0xd7, 0x28, 0x34, 0xad, 0x3f, 0xed, 0x1b, 0xb1, 0x49, 0x29, 0x11, 0xd5, 0xfb, 0xe4, 0x39, 0x78,
	0x85, 0xc2, 0x2c, 0x0e, 0x17, 0x41, 0xfb, 0x90, 0xba, 0xef, 0x51, 0x1f, 0xbc, 0x4f, 0x32, 0xbf,
	0xe2, 0xe2, 0x92, 0xfe, 0x74, 0xa0, 0x1b, 0x2b, 0xcc, 0xc8, 0x13, 0xe8, 0x14, 0x3c, 0x39, 0x36,
	0xa1, 0xac, 0x93, 0x21, 0x9c, 0x08, 0xbe, 0xba, 0x16, 0x6c, 0x83, 0x5a, 0xda, 0x8f, 0x2a, 0x5c,
	0x6e, 0x81, 0x1b, 0xc6, 0xd7, 0x41, 0xc7, 0x6c, 0xa1, 0x01, 0x39, 0x85, 0x1e, 0xdb, 0xb2, 0x9c,
	0x65, 0x41, 0x57, 0x97, 0x2d, 0x2a, 0xeb, 0xa9, 0xe4, 0x22, 0x57, 0x81, 0x3b, 0x72, 0xc6, 0x6e,
	0x64, 0x11, 0xfd, 0x02, 0x50, 0x1a, 0x59, 0xf2, 0x4b, 0x11, 0xa7, 0x37, 0xe6, 0x39, 0xb7, 0xcd,
	0x6b, 0xd7, 0xe7, 0x0d, 0xe1, 0x24, 0x65, 0x4a, 0xfd, 0x90, 0x59, 0x62, 0x8d, 0x54, 0x98, 0x2e,
	0xfe, 0x6a, 0x87, 0xe2, 0x4e, 0xed, 0xba, 0x4a, 0xbb, 0xa1, 0xf2, 0x16, 0x06, 0x1f, 0x90, 0x25,
	0x98, 0xa9, 0x08, 0xbf, 0x17, 0xa8, 0xf2, 0xd2, 0xc9, 0x9a, 0x6f, 0x78, 0xae, 0x65, 0xdc, 0xc8,
	0x80, 0x72, 0x43, 0x79, 0x71, 0xa1, 0x30, 0xd7, 0x0a, 0x6e, 0x64, 0x11, 0x7d, 0x0f, 0xff, 0x57,
	0xdf, 0xab, 0x54, 0x0a, 0x85, 0xe4, 0x29, 0xb8, 0x65, 0x28, 0x2a, 0x70, 0x46, 0x9d, 0xfa, 0xbb,
	0x2b, 0xcc, 0x22, 0xd3, 0x29, 0x67, 0xbc, 0x93, 0x85, 0xd8, 0x8b, 0x19, 0x40, 0x3f, 0xc2, 0xbd,
	0x59, 0x92, 0xcc, 0xf4, 0x93, 0xee, 0xdd, 0xd4, 0xd2, 0x77, 0x6e, 0x4f, 0x9f, 0x10, 0xe8, 0xd6,
	0x62, 0xd4, 0xff, 0xa7, 0xbf, 0x1d, 0x18, 0xd8, 0x9b, 0x5b, 0x62, 0xb6, 0xe5, 0x2b, 0x24, 0x63,
	0xf0, 0xe6, 0x3b, 0x73, 0x56, 0x75, 0x9d, 0xe1, 0xc0, 0x02, 0x4b, 0xa7, 0x2d, 0x72, 0x06, 0xfe,
	0x7c, 0x17, 0x5b, 0xf5, 0x7f, 0x70, 0x5f, 0x40, 0x67, 0x96, 0x24, 0xa4, 0xd1, 0xa8, 0x88, 0xfb,
	0xb3, 0x6c, 0x91, 0x09, 0xfc, 0x17, 0xe1, 0x46, 0x6e, 0xf1, 0x4e, 0x13, 0x15, 0x7f, 0xfa, 0xab,
	0x0d, 0x7d, 0x9d, 0xb1, 0xb5, 0xff, 0x06, 0x3c, 0xfb, 0xd8, 0xe4, 0xa1, 0x25, 0xdf, 0x0c, 0x6f,
	0x78, 0xda, 0x2c, 0x9b, 0x4c, 0x68, 0x8b, 0xbc, 0x84, 0xc1, 0x7c, 0x67, 0xad, 0x1d, 0x19, 0x5e,
	0x4f, 0x89, 0xb6, 0xc8, 0x33, 0x70, 0xe7, 0xbb, 0x83, 0xd5, 0x1b, 0xa4, 0x73, 0xe8, 0xd9, 0xcb,
	0xbe, 0x5f, 0x6b, 0x98, 0xd2, 0x91, 0xed, 0xcf, 0x0c, 0x3d, 0x14, 0x07, 0xf4, 0x50, 0x34, 0xa5,
	0x5f, 0x83, 0x5f, 0x9d, 0x02, 0x79, 0x64, 0x7b, 0xcd, 0xe3, 0x38, 0x9c, 0xf1, 0xad, 0xa7, 0x0b,
	0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x25, 0x5e, 0x72, 0xa9, 0xa0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	ByToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Session, error)
	ByUserUID(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Session, error)
	Add(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Nothing, error)
	RemoveByToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Nothing, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) ByToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/proto.SessionService/ByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) ByUserUID(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/proto.SessionService/ByUserUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) Add(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.SessionService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) RemoveByToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.SessionService/RemoveByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	ByToken(context.Context, *UUID) (*Session, error)
	ByUserUID(context.Context, *UUID) (*Session, error)
	Add(context.Context, *Session) (*Nothing, error)
	RemoveByToken(context.Context, *UUID) (*Nothing, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_ByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/ByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ByToken(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_ByUserUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ByUserUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/ByUserUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ByUserUID(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Add(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_RemoveByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).RemoveByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/RemoveByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).RemoveByToken(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByToken",
			Handler:    _SessionService_ByToken_Handler,
		},
		{
			MethodName: "ByUserUID",
			Handler:    _SessionService_ByUserUID_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _SessionService_Add_Handler,
		},
		{
			MethodName: "RemoveByToken",
			Handler:    _SessionService_RemoveByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Leaders(ctx context.Context, in *LeadersRequest, opts ...grpc.CallOption) (*LeadersResponse, error)
	BySessionToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*User, error)
	ByUID(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*User, error)
	SignUp(ctx context.Context, in *UserSignUp, opts ...grpc.CallOption) (*Nothing, error)
	SignIn(ctx context.Context, in *UserSignIn, opts ...grpc.CallOption) (*User, error)
	AddAvatar(ctx context.Context, in *AddAvatarRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Leaders(ctx context.Context, in *LeadersRequest, opts ...grpc.CallOption) (*LeadersResponse, error) {
	out := new(LeadersResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Leaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BySessionToken(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserService/BySessionToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ByUID(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserService/ByUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignUp(ctx context.Context, in *UserSignUp, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignIn(ctx context.Context, in *UserSignIn, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAvatar(ctx context.Context, in *AddAvatarRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/proto.UserService/AddAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Leaders(context.Context, *LeadersRequest) (*LeadersResponse, error)
	BySessionToken(context.Context, *UUID) (*User, error)
	ByUID(context.Context, *UUID) (*User, error)
	SignUp(context.Context, *UserSignUp) (*Nothing, error)
	SignIn(context.Context, *UserSignIn) (*User, error)
	AddAvatar(context.Context, *AddAvatarRequest) (*Nothing, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Leaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Leaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Leaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Leaders(ctx, req.(*LeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BySessionToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BySessionToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/BySessionToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BySessionToken(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ByUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ByUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ByUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ByUID(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*UserSignUp))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignIn(ctx, req.(*UserSignIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/AddAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAvatar(ctx, req.(*AddAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Leaders",
			Handler:    _UserService_Leaders_Handler,
		},
		{
			MethodName: "BySessionToken",
			Handler:    _UserService_BySessionToken_Handler,
		},
		{
			MethodName: "ByUID",
			Handler:    _UserService_ByUID_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UserService_SignIn_Handler,
		},
		{
			MethodName: "AddAvatar",
			Handler:    _UserService_AddAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

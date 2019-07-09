// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/pkg/proto/auth.proto

package authProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Role                 string   `protobuf:"bytes,5,opt,name=role" json:"role,omitempty"`
	RefreshToken         string   `protobuf:"bytes,6,opt,name=refreshToken" json:"refreshToken,omitempty"`
	PeopleNumber         int64    `protobuf:"varint,7,opt,name=peopleNumber" json:"peopleNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *User) GetPeopleNumber() int64 {
	if m != nil {
		return m.PeopleNumber
	}
	return 0
}

type Developer struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	PrimarySkill         string   `protobuf:"bytes,4,opt,name=primarySkill" json:"primarySkill,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Developer) Reset()         { *m = Developer{} }
func (m *Developer) String() string { return proto.CompactTextString(m) }
func (*Developer) ProtoMessage()    {}
func (*Developer) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{1}
}
func (m *Developer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Developer.Unmarshal(m, b)
}
func (m *Developer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Developer.Marshal(b, m, deterministic)
}
func (dst *Developer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Developer.Merge(dst, src)
}
func (m *Developer) XXX_Size() int {
	return xxx_messageInfo_Developer.Size(m)
}
func (m *Developer) XXX_DiscardUnknown() {
	xxx_messageInfo_Developer.DiscardUnknown(m)
}

var xxx_messageInfo_Developer proto.InternalMessageInfo

func (m *Developer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Developer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Developer) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Developer) GetPrimarySkill() string {
	if m != nil {
		return m.PrimarySkill
	}
	return ""
}

type CreateDeveloperRequest struct {
	Developer            *Developer `protobuf:"bytes,1,opt,name=developer" json:"developer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateDeveloperRequest) Reset()         { *m = CreateDeveloperRequest{} }
func (m *CreateDeveloperRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeveloperRequest) ProtoMessage()    {}
func (*CreateDeveloperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{2}
}
func (m *CreateDeveloperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeveloperRequest.Unmarshal(m, b)
}
func (m *CreateDeveloperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeveloperRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDeveloperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeveloperRequest.Merge(dst, src)
}
func (m *CreateDeveloperRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeveloperRequest.Size(m)
}
func (m *CreateDeveloperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeveloperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeveloperRequest proto.InternalMessageInfo

func (m *CreateDeveloperRequest) GetDeveloper() *Developer {
	if m != nil {
		return m.Developer
	}
	return nil
}

type CreateDeveloperResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDeveloperResponse) Reset()         { *m = CreateDeveloperResponse{} }
func (m *CreateDeveloperResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeveloperResponse) ProtoMessage()    {}
func (*CreateDeveloperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{3}
}
func (m *CreateDeveloperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeveloperResponse.Unmarshal(m, b)
}
func (m *CreateDeveloperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeveloperResponse.Marshal(b, m, deterministic)
}
func (dst *CreateDeveloperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeveloperResponse.Merge(dst, src)
}
func (m *CreateDeveloperResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDeveloperResponse.Size(m)
}
func (m *CreateDeveloperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeveloperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeveloperResponse proto.InternalMessageInfo

func (m *CreateDeveloperResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadDeveloperRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadDeveloperRequest) Reset()         { *m = ReadDeveloperRequest{} }
func (m *ReadDeveloperRequest) String() string { return proto.CompactTextString(m) }
func (*ReadDeveloperRequest) ProtoMessage()    {}
func (*ReadDeveloperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{4}
}
func (m *ReadDeveloperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadDeveloperRequest.Unmarshal(m, b)
}
func (m *ReadDeveloperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadDeveloperRequest.Marshal(b, m, deterministic)
}
func (dst *ReadDeveloperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadDeveloperRequest.Merge(dst, src)
}
func (m *ReadDeveloperRequest) XXX_Size() int {
	return xxx_messageInfo_ReadDeveloperRequest.Size(m)
}
func (m *ReadDeveloperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadDeveloperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadDeveloperRequest proto.InternalMessageInfo

func (m *ReadDeveloperRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadDeveloperResponse struct {
	Developer            *Developer `protobuf:"bytes,1,opt,name=developer" json:"developer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadDeveloperResponse) Reset()         { *m = ReadDeveloperResponse{} }
func (m *ReadDeveloperResponse) String() string { return proto.CompactTextString(m) }
func (*ReadDeveloperResponse) ProtoMessage()    {}
func (*ReadDeveloperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{5}
}
func (m *ReadDeveloperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadDeveloperResponse.Unmarshal(m, b)
}
func (m *ReadDeveloperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadDeveloperResponse.Marshal(b, m, deterministic)
}
func (dst *ReadDeveloperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadDeveloperResponse.Merge(dst, src)
}
func (m *ReadDeveloperResponse) XXX_Size() int {
	return xxx_messageInfo_ReadDeveloperResponse.Size(m)
}
func (m *ReadDeveloperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadDeveloperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadDeveloperResponse proto.InternalMessageInfo

func (m *ReadDeveloperResponse) GetDeveloper() *Developer {
	if m != nil {
		return m.Developer
	}
	return nil
}

type UpdateDeveloperRequest struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Developer            *Developer `protobuf:"bytes,2,opt,name=developer" json:"developer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateDeveloperRequest) Reset()         { *m = UpdateDeveloperRequest{} }
func (m *UpdateDeveloperRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeveloperRequest) ProtoMessage()    {}
func (*UpdateDeveloperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{6}
}
func (m *UpdateDeveloperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeveloperRequest.Unmarshal(m, b)
}
func (m *UpdateDeveloperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeveloperRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateDeveloperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeveloperRequest.Merge(dst, src)
}
func (m *UpdateDeveloperRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeveloperRequest.Size(m)
}
func (m *UpdateDeveloperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeveloperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeveloperRequest proto.InternalMessageInfo

func (m *UpdateDeveloperRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateDeveloperRequest) GetDeveloper() *Developer {
	if m != nil {
		return m.Developer
	}
	return nil
}

type UpdateDeveloperResponse struct {
	Developer            *Developer `protobuf:"bytes,1,opt,name=developer" json:"developer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateDeveloperResponse) Reset()         { *m = UpdateDeveloperResponse{} }
func (m *UpdateDeveloperResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDeveloperResponse) ProtoMessage()    {}
func (*UpdateDeveloperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{7}
}
func (m *UpdateDeveloperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeveloperResponse.Unmarshal(m, b)
}
func (m *UpdateDeveloperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeveloperResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateDeveloperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeveloperResponse.Merge(dst, src)
}
func (m *UpdateDeveloperResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDeveloperResponse.Size(m)
}
func (m *UpdateDeveloperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeveloperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeveloperResponse proto.InternalMessageInfo

func (m *UpdateDeveloperResponse) GetDeveloper() *Developer {
	if m != nil {
		return m.Developer
	}
	return nil
}

type DeleteDeveloperRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeveloperRequest) Reset()         { *m = DeleteDeveloperRequest{} }
func (m *DeleteDeveloperRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeveloperRequest) ProtoMessage()    {}
func (*DeleteDeveloperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{8}
}
func (m *DeleteDeveloperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeveloperRequest.Unmarshal(m, b)
}
func (m *DeleteDeveloperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeveloperRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDeveloperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeveloperRequest.Merge(dst, src)
}
func (m *DeleteDeveloperRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeveloperRequest.Size(m)
}
func (m *DeleteDeveloperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeveloperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeveloperRequest proto.InternalMessageInfo

func (m *DeleteDeveloperRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteDeveloperResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeveloperResponse) Reset()         { *m = DeleteDeveloperResponse{} }
func (m *DeleteDeveloperResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteDeveloperResponse) ProtoMessage()    {}
func (*DeleteDeveloperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{9}
}
func (m *DeleteDeveloperResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeveloperResponse.Unmarshal(m, b)
}
func (m *DeleteDeveloperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeveloperResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteDeveloperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeveloperResponse.Merge(dst, src)
}
func (m *DeleteDeveloperResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteDeveloperResponse.Size(m)
}
func (m *DeleteDeveloperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeveloperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeveloperResponse proto.InternalMessageInfo

type ReadAllDevelopersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllDevelopersRequest) Reset()         { *m = ReadAllDevelopersRequest{} }
func (m *ReadAllDevelopersRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllDevelopersRequest) ProtoMessage()    {}
func (*ReadAllDevelopersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{10}
}
func (m *ReadAllDevelopersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllDevelopersRequest.Unmarshal(m, b)
}
func (m *ReadAllDevelopersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllDevelopersRequest.Marshal(b, m, deterministic)
}
func (dst *ReadAllDevelopersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllDevelopersRequest.Merge(dst, src)
}
func (m *ReadAllDevelopersRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllDevelopersRequest.Size(m)
}
func (m *ReadAllDevelopersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllDevelopersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllDevelopersRequest proto.InternalMessageInfo

type ReadAllDevelopersResponse struct {
	Developers           []*Developer `protobuf:"bytes,1,rep,name=developers" json:"developers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadAllDevelopersResponse) Reset()         { *m = ReadAllDevelopersResponse{} }
func (m *ReadAllDevelopersResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllDevelopersResponse) ProtoMessage()    {}
func (*ReadAllDevelopersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{11}
}
func (m *ReadAllDevelopersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllDevelopersResponse.Unmarshal(m, b)
}
func (m *ReadAllDevelopersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllDevelopersResponse.Marshal(b, m, deterministic)
}
func (dst *ReadAllDevelopersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllDevelopersResponse.Merge(dst, src)
}
func (m *ReadAllDevelopersResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllDevelopersResponse.Size(m)
}
func (m *ReadAllDevelopersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllDevelopersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllDevelopersResponse proto.InternalMessageInfo

func (m *ReadAllDevelopersResponse) GetDevelopers() []*Developer {
	if m != nil {
		return m.Developers
	}
	return nil
}

type SignInRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{12}
}
func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (dst *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(dst, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken" json:"accessToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{13}
}
func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (dst *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(dst, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *SignInResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type SignUpRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{14}
}
func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (dst *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(dst, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SignUpResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken" json:"accessToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refreshToken" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d04ca531c53671f3, []int{15}
}
func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (dst *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(dst, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *SignUpResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "authProto.User")
	proto.RegisterType((*Developer)(nil), "authProto.Developer")
	proto.RegisterType((*CreateDeveloperRequest)(nil), "authProto.CreateDeveloperRequest")
	proto.RegisterType((*CreateDeveloperResponse)(nil), "authProto.CreateDeveloperResponse")
	proto.RegisterType((*ReadDeveloperRequest)(nil), "authProto.ReadDeveloperRequest")
	proto.RegisterType((*ReadDeveloperResponse)(nil), "authProto.ReadDeveloperResponse")
	proto.RegisterType((*UpdateDeveloperRequest)(nil), "authProto.UpdateDeveloperRequest")
	proto.RegisterType((*UpdateDeveloperResponse)(nil), "authProto.UpdateDeveloperResponse")
	proto.RegisterType((*DeleteDeveloperRequest)(nil), "authProto.DeleteDeveloperRequest")
	proto.RegisterType((*DeleteDeveloperResponse)(nil), "authProto.DeleteDeveloperResponse")
	proto.RegisterType((*ReadAllDevelopersRequest)(nil), "authProto.ReadAllDevelopersRequest")
	proto.RegisterType((*ReadAllDevelopersResponse)(nil), "authProto.ReadAllDevelopersResponse")
	proto.RegisterType((*SignInRequest)(nil), "authProto.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "authProto.SignInResponse")
	proto.RegisterType((*SignUpRequest)(nil), "authProto.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "authProto.SignUpResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	ReadAllDevelopers(ctx context.Context, in *ReadAllDevelopersRequest, opts ...grpc.CallOption) (*ReadAllDevelopersResponse, error)
	CreateDeveloper(ctx context.Context, in *CreateDeveloperRequest, opts ...grpc.CallOption) (*CreateDeveloperResponse, error)
	ReadDeveloper(ctx context.Context, in *ReadDeveloperRequest, opts ...grpc.CallOption) (*ReadDeveloperResponse, error)
	UpdateDeveloper(ctx context.Context, in *UpdateDeveloperRequest, opts ...grpc.CallOption) (*UpdateDeveloperResponse, error)
	DeleteDeveloper(ctx context.Context, in *DeleteDeveloperRequest, opts ...grpc.CallOption) (*DeleteDeveloperResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) ReadAllDevelopers(ctx context.Context, in *ReadAllDevelopersRequest, opts ...grpc.CallOption) (*ReadAllDevelopersResponse, error) {
	out := new(ReadAllDevelopersResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/ReadAllDevelopers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateDeveloper(ctx context.Context, in *CreateDeveloperRequest, opts ...grpc.CallOption) (*CreateDeveloperResponse, error) {
	out := new(CreateDeveloperResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/CreateDeveloper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ReadDeveloper(ctx context.Context, in *ReadDeveloperRequest, opts ...grpc.CallOption) (*ReadDeveloperResponse, error) {
	out := new(ReadDeveloperResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/ReadDeveloper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateDeveloper(ctx context.Context, in *UpdateDeveloperRequest, opts ...grpc.CallOption) (*UpdateDeveloperResponse, error) {
	out := new(UpdateDeveloperResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/UpdateDeveloper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteDeveloper(ctx context.Context, in *DeleteDeveloperRequest, opts ...grpc.CallOption) (*DeleteDeveloperResponse, error) {
	out := new(DeleteDeveloperResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/DeleteDeveloper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/SignIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := grpc.Invoke(ctx, "/authProto.Auth/SignUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	ReadAllDevelopers(context.Context, *ReadAllDevelopersRequest) (*ReadAllDevelopersResponse, error)
	CreateDeveloper(context.Context, *CreateDeveloperRequest) (*CreateDeveloperResponse, error)
	ReadDeveloper(context.Context, *ReadDeveloperRequest) (*ReadDeveloperResponse, error)
	UpdateDeveloper(context.Context, *UpdateDeveloperRequest) (*UpdateDeveloperResponse, error)
	DeleteDeveloper(context.Context, *DeleteDeveloperRequest) (*DeleteDeveloperResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_ReadAllDevelopers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllDevelopersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ReadAllDevelopers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/ReadAllDevelopers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ReadAllDevelopers(ctx, req.(*ReadAllDevelopersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/CreateDeveloper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateDeveloper(ctx, req.(*CreateDeveloperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ReadDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ReadDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/ReadDeveloper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ReadDeveloper(ctx, req.(*ReadDeveloperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/UpdateDeveloper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateDeveloper(ctx, req.(*UpdateDeveloperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/DeleteDeveloper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteDeveloper(ctx, req.(*DeleteDeveloperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authProto.Auth/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authProto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadAllDevelopers",
			Handler:    _Auth_ReadAllDevelopers_Handler,
		},
		{
			MethodName: "CreateDeveloper",
			Handler:    _Auth_CreateDeveloper_Handler,
		},
		{
			MethodName: "ReadDeveloper",
			Handler:    _Auth_ReadDeveloper_Handler,
		},
		{
			MethodName: "UpdateDeveloper",
			Handler:    _Auth_UpdateDeveloper_Handler,
		},
		{
			MethodName: "DeleteDeveloper",
			Handler:    _Auth_DeleteDeveloper_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Auth_SignIn_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Auth_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/pkg/proto/auth.proto",
}

func init() { proto.RegisterFile("rpc/pkg/proto/auth.proto", fileDescriptor_auth_d04ca531c53671f3) }

var fileDescriptor_auth_d04ca531c53671f3 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x97, 0x34, 0xeb, 0xfb, 0xf6, 0x8c, 0x6d, 0x60, 0x95, 0x36, 0x8d, 0x90, 0x28, 0x1e,
	0x42, 0xe5, 0xa6, 0x95, 0xca, 0xee, 0x51, 0xc5, 0x6e, 0x10, 0x7f, 0x04, 0xd9, 0xca, 0x05, 0xe2,
	0x02, 0xaf, 0x39, 0xa4, 0x51, 0xd3, 0xd8, 0xc4, 0x29, 0x88, 0xaf, 0xc6, 0xe7, 0xe0, 0x03, 0xa1,
	0x38, 0x69, 0x1a, 0x27, 0xe9, 0x40, 0xc0, 0xdd, 0xc9, 0x39, 0x4f, 0x1e, 0xff, 0x1c, 0xfb, 0x69,
	0xc1, 0x8e, 0xc5, 0x62, 0x22, 0x56, 0xfe, 0x44, 0xc4, 0x3c, 0xe1, 0x13, 0xb6, 0x49, 0x96, 0x63,
	0x55, 0x92, 0x4e, 0x5a, 0xbf, 0x49, 0x4b, 0xe7, 0x9e, 0xcf, 0xb9, 0x1f, 0xe2, 0x84, 0x89, 0x60,
	0xc2, 0xa2, 0x88, 0x27, 0x2c, 0x09, 0x78, 0x24, 0x33, 0x21, 0xfd, 0x6e, 0x80, 0x35, 0x97, 0x18,
	0x93, 0x13, 0x30, 0x03, 0xcf, 0x36, 0x86, 0xc6, 0xa8, 0xe5, 0x9a, 0x81, 0x47, 0xba, 0x70, 0x18,
	0x72, 0x3f, 0x88, 0x6c, 0x73, 0x68, 0x8c, 0x3a, 0x6e, 0xf6, 0x90, 0x76, 0x71, 0xcd, 0x82, 0xd0,
	0x6e, 0x65, 0x5d, 0xf5, 0x40, 0x1c, 0xf8, 0x5f, 0x30, 0x29, 0xbf, 0xf2, 0xd8, 0xb3, 0x2d, 0x35,
	0x28, 0x9e, 0x09, 0x01, 0x2b, 0xe6, 0x21, 0xda, 0x87, 0xaa, 0xaf, 0x6a, 0x42, 0xe1, 0x56, 0x8c,
	0x9f, 0x62, 0x94, 0xcb, 0x2b, 0xbe, 0xc2, 0xc8, 0x6e, 0xab, 0x99, 0xd6, 0x4b, 0x35, 0x02, 0xb9,
	0x08, 0xf1, 0xf5, 0x66, 0x7d, 0x8d, 0xb1, 0xfd, 0x9f, 0x22, 0xd3, 0x7a, 0x14, 0xa1, 0x73, 0x81,
	0x5f, 0x30, 0xe4, 0xa2, 0x61, 0x03, 0x04, 0xac, 0x88, 0xad, 0x31, 0xe7, 0x57, 0x35, 0xb9, 0x0d,
	0x2d, 0xe6, 0xa3, 0x82, 0x6f, 0xb9, 0x69, 0xa9, 0x96, 0x89, 0x83, 0x35, 0x8b, 0xbf, 0x5d, 0xae,
	0x82, 0x30, 0xcc, 0xf1, 0xb5, 0x1e, 0x7d, 0x09, 0xbd, 0x67, 0x31, 0xb2, 0x04, 0x8b, 0xc5, 0x5c,
	0xfc, 0xbc, 0x41, 0x99, 0x90, 0x29, 0x74, 0xbc, 0x6d, 0x4f, 0x2d, 0x7d, 0x34, 0xed, 0x8e, 0x8b,
	0x4f, 0x3f, 0xde, 0xe9, 0x77, 0x32, 0xfa, 0x18, 0xfa, 0x35, 0x37, 0x29, 0x78, 0x24, 0xb1, 0xba,
	0x05, 0xfa, 0x08, 0xba, 0x2e, 0x32, 0xaf, 0xb6, 0x6c, 0x55, 0xf7, 0x02, 0xee, 0x56, 0x74, 0xb9,
	0xe1, 0x9f, 0xf0, 0x7d, 0x80, 0xde, 0x5c, 0x78, 0x4d, 0xbb, 0xad, 0x7e, 0x61, 0xcd, 0xdd, 0xfc,
	0x3d, 0xf7, 0x57, 0xd0, 0xaf, 0xb9, 0xff, 0x05, 0xec, 0x08, 0x7a, 0x17, 0x18, 0xe2, 0xaf, 0x61,
	0xe9, 0x00, 0xfa, 0x35, 0x65, 0xb6, 0x30, 0x75, 0xc0, 0x4e, 0x3f, 0xdf, 0x2c, 0x0c, 0x8b, 0x99,
	0xcc, 0x6d, 0xe8, 0x5b, 0x18, 0x34, 0xcc, 0x72, 0xe2, 0x73, 0x80, 0x02, 0x45, 0xda, 0xc6, 0xb0,
	0xb5, 0x17, 0xb9, 0xa4, 0xa3, 0x33, 0x38, 0xbe, 0x0c, 0xfc, 0xe8, 0x79, 0xb4, 0x45, 0x2d, 0xa2,
	0x66, 0x94, 0xa3, 0x56, 0x0e, 0x95, 0xa9, 0x87, 0x8a, 0xbe, 0x83, 0x93, 0xad, 0x45, 0x8e, 0x32,
	0x84, 0x23, 0xb6, 0x58, 0xa0, 0x94, 0x59, 0xa2, 0x32, 0xa7, 0x72, 0xab, 0x16, 0x3a, 0xb3, 0x1e,
	0x3a, 0x7a, 0x9e, 0xa1, 0xcd, 0xc5, 0x16, 0xed, 0x0c, 0xac, 0x8d, 0x2c, 0x8e, 0xe3, 0xb4, 0xb4,
	0xb7, 0xf4, 0x47, 0xc3, 0x55, 0xc3, 0x2d, 0x4d, 0xfa, 0xd6, 0xbf, 0xa4, 0x99, 0xfe, 0xb0, 0xc0,
	0x9a, 0x6d, 0x92, 0x25, 0xf9, 0x08, 0x77, 0x6a, 0x87, 0x40, 0xce, 0x4a, 0x30, 0xfb, 0x8e, 0xcf,
	0x79, 0x78, 0xb3, 0x28, 0xbf, 0x00, 0x07, 0xe4, 0x3d, 0x9c, 0x56, 0x42, 0x49, 0x1e, 0x94, 0x5e,
	0x6d, 0x8e, 0xbf, 0x43, 0x6f, 0x92, 0x14, 0xde, 0x57, 0x70, 0xac, 0xa5, 0x93, 0xdc, 0xaf, 0x40,
	0xd5, 0x7c, 0x87, 0xfb, 0x05, 0x65, 0xe2, 0x4a, 0x90, 0x34, 0xe2, 0xe6, 0x08, 0x6b, 0xc4, 0x7b,
	0x72, 0x98, 0x79, 0x57, 0xb2, 0xa2, 0x79, 0x37, 0x27, 0x4e, 0xf3, 0xde, 0x17, 0xb5, 0x03, 0xf2,
	0x14, 0xda, 0xd9, 0xd5, 0x25, 0x76, 0x49, 0xaf, 0x05, 0xc2, 0x19, 0x34, 0x4c, 0xaa, 0x06, 0x73,
	0x51, 0x33, 0x28, 0xae, 0x6d, 0xcd, 0x60, 0x77, 0x35, 0xe9, 0xc1, 0x75, 0x5b, 0xfd, 0xf3, 0x3d,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xcc, 0x69, 0xd7, 0x3e, 0x07, 0x00, 0x00,
}

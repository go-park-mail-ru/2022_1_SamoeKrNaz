// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: handler/user.proto

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative handler/user.proto

package handler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Username struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	USERNAME string `protobuf:"bytes,1,opt,name=USERNAME,proto3" json:"USERNAME,omitempty"`
}

func (x *Username) Reset() {
	*x = Username{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Username) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Username) ProtoMessage() {}

func (x *Username) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Username.ProtoReflect.Descriptor instead.
func (*Username) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{0}
}

func (x *Username) GetUSERNAME() string {
	if x != nil {
		return x.USERNAME
	}
	return ""
}

type IdUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDU uint64 `protobuf:"varint,1,opt,name=IDU,proto3" json:"IDU,omitempty"`
}

func (x *IdUser) Reset() {
	*x = IdUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdUser) ProtoMessage() {}

func (x *IdUser) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdUser.ProtoReflect.Descriptor instead.
func (*IdUser) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{1}
}

func (x *IdUser) GetIDU() uint64 {
	if x != nil {
		return x.IDU
	}
	return 0
}

type IdBoard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDB uint64 `protobuf:"varint,1,opt,name=IDB,proto3" json:"IDB,omitempty"`
}

func (x *IdBoard) Reset() {
	*x = IdBoard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdBoard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdBoard) ProtoMessage() {}

func (x *IdBoard) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdBoard.ProtoReflect.Descriptor instead.
func (*IdBoard) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{2}
}

func (x *IdBoard) GetIDB() uint64 {
	if x != nil {
		return x.IDB
	}
	return 0
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDB *IdBoard `protobuf:"bytes,1,opt,name=IDB,proto3" json:"IDB,omitempty"`
	IDU *IdUser  `protobuf:"bytes,2,opt,name=IDU,proto3" json:"IDU,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{3}
}

func (x *Ids) GetIDB() *IdBoard {
	if x != nil {
		return x.IDB
	}
	return nil
}

func (x *Ids) GetIDU() *IdUser {
	if x != nil {
		return x.IDU
	}
	return nil
}

type CheckLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uname *Username `protobuf:"bytes,1,opt,name=uname,proto3" json:"uname,omitempty"`
	Pass  string    `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *CheckLog) Reset() {
	*x = CheckLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckLog) ProtoMessage() {}

func (x *CheckLog) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckLog.ProtoReflect.Descriptor instead.
func (*CheckLog) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{4}
}

func (x *CheckLog) GetUname() *Username {
	if x != nil {
		return x.Uname
	}
	return nil
}

func (x *CheckLog) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDU      *IdUser   `protobuf:"bytes,1,opt,name=IDU,proto3" json:"IDU,omitempty"`
	UserData *CheckLog `protobuf:"bytes,2,opt,name=userData,proto3" json:"userData,omitempty"`
	IMG      string    `protobuf:"bytes,3,opt,name=IMG,proto3" json:"IMG,omitempty"`
	BOARDS   []byte    `protobuf:"bytes,4,opt,name=BOARDS,proto3" json:"BOARDS,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetIDU() *IdUser {
	if x != nil {
		return x.IDU
	}
	return nil
}

func (x *User) GetUserData() *CheckLog {
	if x != nil {
		return x.UserData
	}
	return nil
}

func (x *User) GetIMG() string {
	if x != nil {
		return x.IMG
	}
	return ""
}

func (x *User) GetBOARDS() []byte {
	if x != nil {
		return x.BOARDS
	}
	return nil
}

type NothingSec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *NothingSec) Reset() {
	*x = NothingSec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NothingSec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NothingSec) ProtoMessage() {}

func (x *NothingSec) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NothingSec.ProtoReflect.Descriptor instead.
func (*NothingSec) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{6}
}

func (x *NothingSec) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	USERS []byte `protobuf:"bytes,1,opt,name=USERS,proto3" json:"USERS,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_handler_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_handler_user_proto_rawDescGZIP(), []int{7}
}

func (x *Users) GetUSERS() []byte {
	if x != nil {
		return x.USERS
	}
	return nil
}

var File_handler_user_proto protoreflect.FileDescriptor

var file_handler_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x26, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x53, 0x45,
	0x52, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x53, 0x45,
	0x52, 0x4e, 0x41, 0x4d, 0x45, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x49, 0x44, 0x55, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x49, 0x44,
	0x55, 0x22, 0x1b, 0x0a, 0x07, 0x49, 0x64, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x49, 0x44, 0x42, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x49, 0x44, 0x42, 0x22, 0x4c,
	0x0a, 0x03, 0x49, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x03, 0x49, 0x44, 0x42, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x03, 0x49, 0x44, 0x42, 0x12, 0x21, 0x0a, 0x03, 0x49, 0x44, 0x55,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x03, 0x49, 0x44, 0x55, 0x22, 0x47, 0x0a, 0x08,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x75, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x03, 0x49, 0x44, 0x55, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x03, 0x49, 0x44,
	0x55, 0x12, 0x2d, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x10, 0x0a, 0x03, 0x49, 0x4d, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49,
	0x4d, 0x47, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x53, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x53, 0x22, 0x22, 0x0a, 0x0a, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x1d,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x53, 0x45, 0x52, 0x53,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x55, 0x53, 0x45, 0x52, 0x53, 0x32, 0xac, 0x03,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x49, 0x64, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x13, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x49, 0x73, 0x41,
	0x62, 0x6c, 0x65, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x1a, 0x13, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x63, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0c, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x49, 0x64, 0x73, 0x1a, 0x13, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x0d, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0f, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0d, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x1a, 0x13, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x63, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x0e, 0x2e, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_handler_user_proto_rawDescOnce sync.Once
	file_handler_user_proto_rawDescData = file_handler_user_proto_rawDesc
)

func file_handler_user_proto_rawDescGZIP() []byte {
	file_handler_user_proto_rawDescOnce.Do(func() {
		file_handler_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_user_proto_rawDescData)
	})
	return file_handler_user_proto_rawDescData
}

var file_handler_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_handler_user_proto_goTypes = []interface{}{
	(*Username)(nil),   // 0: handler.Username
	(*IdUser)(nil),     // 1: handler.IdUser
	(*IdBoard)(nil),    // 2: handler.IdBoard
	(*Ids)(nil),        // 3: handler.Ids
	(*CheckLog)(nil),   // 4: handler.CheckLog
	(*User)(nil),       // 5: handler.User
	(*NothingSec)(nil), // 6: handler.NothingSec
	(*Users)(nil),      // 7: handler.Users
}
var file_handler_user_proto_depIdxs = []int32{
	2,  // 0: handler.Ids.IDB:type_name -> handler.IdBoard
	1,  // 1: handler.Ids.IDU:type_name -> handler.IdUser
	0,  // 2: handler.CheckLog.uname:type_name -> handler.Username
	1,  // 3: handler.User.IDU:type_name -> handler.IdUser
	4,  // 4: handler.User.userData:type_name -> handler.CheckLog
	5,  // 5: handler.UserService.Create:input_type -> handler.User
	5,  // 6: handler.UserService.Update:input_type -> handler.User
	4,  // 7: handler.UserService.IsAbleToLogin:input_type -> handler.CheckLog
	3,  // 8: handler.UserService.AddUserToBoard:input_type -> handler.Ids
	0,  // 9: handler.UserService.GetUserByLogin:input_type -> handler.Username
	1,  // 10: handler.UserService.GetUserById:input_type -> handler.IdUser
	0,  // 11: handler.UserService.IsExist:input_type -> handler.Username
	0,  // 12: handler.UserService.GetUsersLike:input_type -> handler.Username
	1,  // 13: handler.UserService.Create:output_type -> handler.IdUser
	6,  // 14: handler.UserService.Update:output_type -> handler.NothingSec
	6,  // 15: handler.UserService.IsAbleToLogin:output_type -> handler.NothingSec
	6,  // 16: handler.UserService.AddUserToBoard:output_type -> handler.NothingSec
	5,  // 17: handler.UserService.GetUserByLogin:output_type -> handler.User
	5,  // 18: handler.UserService.GetUserById:output_type -> handler.User
	6,  // 19: handler.UserService.IsExist:output_type -> handler.NothingSec
	7,  // 20: handler.UserService.GetUsersLike:output_type -> handler.Users
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_handler_user_proto_init() }
func file_handler_user_proto_init() {
	if File_handler_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Username); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdBoard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NothingSec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handler_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_handler_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_user_proto_goTypes,
		DependencyIndexes: file_handler_user_proto_depIdxs,
		MessageInfos:      file_handler_user_proto_msgTypes,
	}.Build()
	File_handler_user_proto = out.File
	file_handler_user_proto_rawDesc = nil
	file_handler_user_proto_goTypes = nil
	file_handler_user_proto_depIdxs = nil
}

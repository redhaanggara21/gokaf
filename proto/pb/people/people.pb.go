// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: people.proto

package people

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	house "pb/house"
	pagination "pb/pagination"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenderGender int32

const (
	Gender_UNKNOWN GenderGender = 0
	Gender_MALE    GenderGender = 1
	Gender_FEMALE  GenderGender = 2
)

// Enum value maps for GenderGender.
var (
	GenderGender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	GenderGender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x GenderGender) Enum() *GenderGender {
	p := new(GenderGender)
	*p = x
	return p
}

func (x GenderGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenderGender) Descriptor() protoreflect.EnumDescriptor {
	return file_people_proto_enumTypes[0].Descriptor()
}

func (GenderGender) Type() protoreflect.EnumType {
	return &file_people_proto_enumTypes[0]
}

func (x GenderGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenderGender.Descriptor instead.
func (GenderGender) EnumDescriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{5, 0}
}

type GetPeopleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPeopleRequest) Reset() {
	*x = GetPeopleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeopleRequest) ProtoMessage() {}

func (x *GetPeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeopleRequest.ProtoReflect.Descriptor instead.
func (*GetPeopleRequest) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{0}
}

func (x *GetPeopleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePeopleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePeopleResponse) Reset() {
	*x = DeletePeopleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePeopleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePeopleResponse) ProtoMessage() {}

func (x *DeletePeopleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePeopleResponse.ProtoReflect.Descriptor instead.
func (*DeletePeopleResponse) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{1}
}

func (x *DeletePeopleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPeopleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPeopleRequest) Reset() {
	*x = ListPeopleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeopleRequest) ProtoMessage() {}

func (x *ListPeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeopleRequest.ProtoReflect.Descriptor instead.
func (*ListPeopleRequest) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{2}
}

type Peoples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *pagination.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Data       []*People              `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Peoples) Reset() {
	*x = Peoples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peoples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peoples) ProtoMessage() {}

func (x *Peoples) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peoples.ProtoReflect.Descriptor instead.
func (*Peoples) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{3}
}

func (x *Peoples) GetPagination() *pagination.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *Peoples) GetData() []*People {
	if x != nil {
		return x.Data
	}
	return nil
}

type People struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age         uint32                 `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Loan        float64                `protobuf:"fixed64,4,opt,name=loan,proto3" json:"loan,omitempty"`
	Gender      *Gender                `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	House       *house.House           `protobuf:"bytes,6,opt,name=house,proto3" json:"house,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *People) Reset() {
	*x = People{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *People) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*People) ProtoMessage() {}

func (x *People) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use People.ProtoReflect.Descriptor instead.
func (*People) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{4}
}

func (x *People) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *People) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *People) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *People) GetLoan() float64 {
	if x != nil {
		return x.Loan
	}
	return 0
}

func (x *People) GetGender() *Gender {
	if x != nil {
		return x.Gender
	}
	return nil
}

func (x *People) GetHouse() *house.House {
	if x != nil {
		return x.House
	}
	return nil
}

func (x *People) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type Gender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name GenderGender `protobuf:"varint,2,opt,name=name,proto3,enum=goprotobuf.GenderGender" json:"name,omitempty"`
}

func (x *Gender) Reset() {
	*x = Gender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gender) ProtoMessage() {}

func (x *Gender) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gender.ProtoReflect.Descriptor instead.
func (*Gender) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{5}
}

func (x *Gender) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Gender) GetName() GenderGender {
	if x != nil {
		return x.Name
	}
	return Gender_UNKNOWN
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_people_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_people_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_people_proto_rawDescGZIP(), []int{6}
}

var File_people_proto protoreflect.FileDescriptor

var file_people_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x07, 0x50,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe6, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x74, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x02, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xd8,
	0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x37, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x73, 0x30, 0x01, 0x12,
	0x37, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x12, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x62, 0x2f,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_people_proto_rawDescOnce sync.Once
	file_people_proto_rawDescData = file_people_proto_rawDesc
)

func file_people_proto_rawDescGZIP() []byte {
	file_people_proto_rawDescOnce.Do(func() {
		file_people_proto_rawDescData = protoimpl.X.CompressGZIP(file_people_proto_rawDescData)
	})
	return file_people_proto_rawDescData
}

var file_people_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_people_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_people_proto_goTypes = []interface{}{
	(GenderGender)(0),             // 0: goprotobuf.Gender.gender
	(*GetPeopleRequest)(nil),      // 1: goprotobuf.GetPeopleRequest
	(*DeletePeopleResponse)(nil),  // 2: goprotobuf.DeletePeopleResponse
	(*ListPeopleRequest)(nil),     // 3: goprotobuf.ListPeopleRequest
	(*Peoples)(nil),               // 4: goprotobuf.Peoples
	(*People)(nil),                // 5: goprotobuf.People
	(*Gender)(nil),                // 6: goprotobuf.Gender
	(*Empty)(nil),                 // 7: goprotobuf.Empty
	(*pagination.Pagination)(nil), // 8: goprotobuf.Pagination
	(*house.House)(nil),           // 9: goprotobuf.House
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_people_proto_depIdxs = []int32{
	8,  // 0: goprotobuf.Peoples.pagination:type_name -> goprotobuf.Pagination
	5,  // 1: goprotobuf.Peoples.data:type_name -> goprotobuf.People
	6,  // 2: goprotobuf.People.gender:type_name -> goprotobuf.Gender
	9,  // 3: goprotobuf.People.house:type_name -> goprotobuf.House
	10, // 4: goprotobuf.People.last_updated:type_name -> google.protobuf.Timestamp
	0,  // 5: goprotobuf.Gender.name:type_name -> goprotobuf.Gender.gender
	5,  // 6: goprotobuf.ProductService.CreateProduct:input_type -> goprotobuf.People
	1,  // 7: goprotobuf.ProductService.GetProduct:input_type -> goprotobuf.GetPeopleRequest
	3,  // 8: goprotobuf.ProductService.GetProducts:input_type -> goprotobuf.ListPeopleRequest
	5,  // 9: goprotobuf.ProductService.UpdateProduct:input_type -> goprotobuf.People
	1,  // 10: goprotobuf.ProductService.DeleteProduct:input_type -> goprotobuf.GetPeopleRequest
	5,  // 11: goprotobuf.ProductService.CreateProduct:output_type -> goprotobuf.People
	5,  // 12: goprotobuf.ProductService.GetProduct:output_type -> goprotobuf.People
	4,  // 13: goprotobuf.ProductService.GetProducts:output_type -> goprotobuf.Peoples
	5,  // 14: goprotobuf.ProductService.UpdateProduct:output_type -> goprotobuf.People
	2,  // 15: goprotobuf.ProductService.DeleteProduct:output_type -> goprotobuf.DeletePeopleResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_people_proto_init() }
func file_people_proto_init() {
	if File_people_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_people_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeopleRequest); i {
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
		file_people_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePeopleResponse); i {
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
		file_people_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPeopleRequest); i {
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
		file_people_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peoples); i {
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
		file_people_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*People); i {
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
		file_people_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gender); i {
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
		file_people_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_people_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_people_proto_goTypes,
		DependencyIndexes: file_people_proto_depIdxs,
		EnumInfos:         file_people_proto_enumTypes,
		MessageInfos:      file_people_proto_msgTypes,
	}.Build()
	File_people_proto = out.File
	file_people_proto_rawDesc = nil
	file_people_proto_goTypes = nil
	file_people_proto_depIdxs = nil
}

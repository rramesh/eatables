// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: items.proto

package items

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllRequest) Reset() {
	*x = ListAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllRequest) ProtoMessage() {}

func (x *ListAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllRequest.ProtoReflect.Descriptor instead.
func (*ListAllRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{0}
}

type ItemsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemDetails `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemsListResponse) Reset() {
	*x = ItemsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsListResponse) ProtoMessage() {}

func (x *ItemsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsListResponse.ProtoReflect.Descriptor instead.
func (*ItemsListResponse) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{1}
}

func (x *ItemsListResponse) GetItems() []*ItemDetails {
	if x != nil {
		return x.Items
	}
	return nil
}

type IDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDRequest) Reset() {
	*x = IDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDRequest) ProtoMessage() {}

func (x *IDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDRequest.ProtoReflect.Descriptor instead.
func (*IDRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{2}
}

func (x *IDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *UUIDRequest) Reset() {
	*x = UUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDRequest) ProtoMessage() {}

func (x *UUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDRequest.ProtoReflect.Descriptor instead.
func (*UUIDRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{3}
}

func (x *UUIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ItemDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sku               string       `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	VendorCode        string       `protobuf:"bytes,3,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	Name              string       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description       string       `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Price             float32      `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	NonVegetarian     bool         `protobuf:"varint,7,opt,name=nonVegetarian,proto3" json:"nonVegetarian,omitempty"`
	Cuisine           string       `protobuf:"bytes,8,opt,name=cuisine,proto3" json:"cuisine,omitempty"`
	Category          []string     `protobuf:"bytes,9,rep,name=category,proto3" json:"category,omitempty"`
	Customizable      bool         `protobuf:"varint,10,opt,name=customizable,proto3" json:"customizable,omitempty"`
	AvailableTimes    []*TimeRange `protobuf:"bytes,11,rep,name=availableTimes,proto3" json:"availableTimes,omitempty"`
	Tags              []string     `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	DontMakeItAnymore bool         `protobuf:"varint,13,opt,name=dontMakeItAnymore,proto3" json:"dontMakeItAnymore,omitempty"`
}

func (x *ItemDetails) Reset() {
	*x = ItemDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDetails) ProtoMessage() {}

func (x *ItemDetails) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDetails.ProtoReflect.Descriptor instead.
func (*ItemDetails) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{4}
}

func (x *ItemDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemDetails) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ItemDetails) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *ItemDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ItemDetails) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemDetails) GetNonVegetarian() bool {
	if x != nil {
		return x.NonVegetarian
	}
	return false
}

func (x *ItemDetails) GetCuisine() string {
	if x != nil {
		return x.Cuisine
	}
	return ""
}

func (x *ItemDetails) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ItemDetails) GetCustomizable() bool {
	if x != nil {
		return x.Customizable
	}
	return false
}

func (x *ItemDetails) GetAvailableTimes() []*TimeRange {
	if x != nil {
		return x.AvailableTimes
	}
	return nil
}

func (x *ItemDetails) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ItemDetails) GetDontMakeItAnymore() bool {
	if x != nil {
		return x.DontMakeItAnymore
	}
	return false
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{5}
}

func (x *TimeRange) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *TimeRange) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

type CreateOrUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku            string       `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	VendorCode     string       `protobuf:"bytes,2,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	Name           string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description    string       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price          float64      `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	NonVegetarian  bool         `protobuf:"varint,6,opt,name=nonVegetarian,proto3" json:"nonVegetarian,omitempty"`
	Cuisine        string       `protobuf:"bytes,7,opt,name=cuisine,proto3" json:"cuisine,omitempty"`
	Category       []string     `protobuf:"bytes,8,rep,name=category,proto3" json:"category,omitempty"`
	Customizable   bool         `protobuf:"varint,9,opt,name=customizable,proto3" json:"customizable,omitempty"`
	AvailableTimes []*TimeRange `protobuf:"bytes,10,rep,name=availableTimes,proto3" json:"availableTimes,omitempty"`
	Tags           []string     `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateOrUpdateRequest) Reset() {
	*x = CreateOrUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateRequest) ProtoMessage() {}

func (x *CreateOrUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOrUpdateRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *CreateOrUpdateRequest) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *CreateOrUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateOrUpdateRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrUpdateRequest) GetNonVegetarian() bool {
	if x != nil {
		return x.NonVegetarian
	}
	return false
}

func (x *CreateOrUpdateRequest) GetCuisine() string {
	if x != nil {
		return x.Cuisine
	}
	return ""
}

func (x *CreateOrUpdateRequest) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *CreateOrUpdateRequest) GetCustomizable() bool {
	if x != nil {
		return x.Customizable
	}
	return false
}

func (x *CreateOrUpdateRequest) GetAvailableTimes() []*TimeRange {
	if x != nil {
		return x.AvailableTimes
	}
	return nil
}

func (x *CreateOrUpdateRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GenericResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{7}
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_items_proto protoreflect.FileDescriptor

var file_items_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x37, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x91, 0x03, 0x0a, 0x0b, 0x49, 0x74, 0x65,
	0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x6e, 0x56, 0x65, 0x67,
	0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6e,
	0x6f, 0x6e, 0x56, 0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x69, 0x73, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x75, 0x69, 0x73, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2c,
	0x0a, 0x11, 0x64, 0x6f, 0x6e, 0x74, 0x4d, 0x61, 0x6b, 0x65, 0x49, 0x74, 0x41, 0x6e, 0x79, 0x6d,
	0x6f, 0x72, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x6f, 0x6e, 0x74, 0x4d,
	0x61, 0x6b, 0x65, 0x49, 0x74, 0x41, 0x6e, 0x79, 0x6d, 0x6f, 0x72, 0x65, 0x22, 0x2f, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xdd, 0x02,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x6e, 0x56, 0x65, 0x67, 0x65,
	0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6e, 0x6f,
	0x6e, 0x56, 0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x75, 0x69, 0x73, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75,
	0x69, 0x73, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x2b, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd5, 0x02, 0x0a, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x0a, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x4b, 0x55, 0x12, 0x0c, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0c, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_items_proto_rawDescOnce sync.Once
	file_items_proto_rawDescData = file_items_proto_rawDesc
)

func file_items_proto_rawDescGZIP() []byte {
	file_items_proto_rawDescOnce.Do(func() {
		file_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_items_proto_rawDescData)
	})
	return file_items_proto_rawDescData
}

var file_items_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_items_proto_goTypes = []interface{}{
	(*ListAllRequest)(nil),        // 0: ListAllRequest
	(*ItemsListResponse)(nil),     // 1: ItemsListResponse
	(*IDRequest)(nil),             // 2: IDRequest
	(*UUIDRequest)(nil),           // 3: UUIDRequest
	(*ItemDetails)(nil),           // 4: ItemDetails
	(*TimeRange)(nil),             // 5: TimeRange
	(*CreateOrUpdateRequest)(nil), // 6: CreateOrUpdateRequest
	(*GenericResponse)(nil),       // 7: GenericResponse
}
var file_items_proto_depIdxs = []int32{
	4,  // 0: ItemsListResponse.items:type_name -> ItemDetails
	5,  // 1: ItemDetails.availableTimes:type_name -> TimeRange
	5,  // 2: CreateOrUpdateRequest.availableTimes:type_name -> TimeRange
	0,  // 3: Items.ListAll:input_type -> ListAllRequest
	2,  // 4: Items.ListByID:input_type -> IDRequest
	3,  // 5: Items.ListBySKU:input_type -> UUIDRequest
	3,  // 6: Items.ListByVendorCode:input_type -> UUIDRequest
	6,  // 7: Items.Add:input_type -> CreateOrUpdateRequest
	6,  // 8: Items.Update:input_type -> CreateOrUpdateRequest
	2,  // 9: Items.Delete:input_type -> IDRequest
	1,  // 10: Items.ListAll:output_type -> ItemsListResponse
	1,  // 11: Items.ListByID:output_type -> ItemsListResponse
	1,  // 12: Items.ListBySKU:output_type -> ItemsListResponse
	1,  // 13: Items.ListByVendorCode:output_type -> ItemsListResponse
	7,  // 14: Items.Add:output_type -> GenericResponse
	7,  // 15: Items.Update:output_type -> GenericResponse
	7,  // 16: Items.Delete:output_type -> GenericResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_items_proto_init() }
func file_items_proto_init() {
	if File_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllRequest); i {
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
		file_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemsListResponse); i {
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
		file_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDRequest); i {
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
		file_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDRequest); i {
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
		file_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDetails); i {
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
		file_items_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
		file_items_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateRequest); i {
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
		file_items_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResponse); i {
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
			RawDescriptor: file_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_items_proto_goTypes,
		DependencyIndexes: file_items_proto_depIdxs,
		MessageInfos:      file_items_proto_msgTypes,
	}.Build()
	File_items_proto = out.File
	file_items_proto_rawDesc = nil
	file_items_proto_goTypes = nil
	file_items_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ItemsClient is the client API for Items service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemsClient interface {
	ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ItemsListResponse, error)
	ListByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error)
	ListBySKU(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error)
	ListByVendorCode(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error)
	Add(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	Update(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	Delete(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GenericResponse, error)
}

type itemsClient struct {
	cc grpc.ClientConnInterface
}

func NewItemsClient(cc grpc.ClientConnInterface) ItemsClient {
	return &itemsClient{cc}
}

func (c *itemsClient) ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ItemsListResponse, error) {
	out := new(ItemsListResponse)
	err := c.cc.Invoke(ctx, "/Items/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) ListByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error) {
	out := new(ItemsListResponse)
	err := c.cc.Invoke(ctx, "/Items/ListByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) ListBySKU(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error) {
	out := new(ItemsListResponse)
	err := c.cc.Invoke(ctx, "/Items/ListBySKU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) ListByVendorCode(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*ItemsListResponse, error) {
	out := new(ItemsListResponse)
	err := c.cc.Invoke(ctx, "/Items/ListByVendorCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) Add(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/Items/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) Update(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/Items/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsClient) Delete(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/Items/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemsServer is the server API for Items service.
type ItemsServer interface {
	ListAll(context.Context, *ListAllRequest) (*ItemsListResponse, error)
	ListByID(context.Context, *IDRequest) (*ItemsListResponse, error)
	ListBySKU(context.Context, *UUIDRequest) (*ItemsListResponse, error)
	ListByVendorCode(context.Context, *UUIDRequest) (*ItemsListResponse, error)
	Add(context.Context, *CreateOrUpdateRequest) (*GenericResponse, error)
	Update(context.Context, *CreateOrUpdateRequest) (*GenericResponse, error)
	Delete(context.Context, *IDRequest) (*GenericResponse, error)
}

// UnimplementedItemsServer can be embedded to have forward compatible implementations.
type UnimplementedItemsServer struct {
}

func (*UnimplementedItemsServer) ListAll(context.Context, *ListAllRequest) (*ItemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (*UnimplementedItemsServer) ListByID(context.Context, *IDRequest) (*ItemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByID not implemented")
}
func (*UnimplementedItemsServer) ListBySKU(context.Context, *UUIDRequest) (*ItemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBySKU not implemented")
}
func (*UnimplementedItemsServer) ListByVendorCode(context.Context, *UUIDRequest) (*ItemsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByVendorCode not implemented")
}
func (*UnimplementedItemsServer) Add(context.Context, *CreateOrUpdateRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedItemsServer) Update(context.Context, *CreateOrUpdateRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedItemsServer) Delete(context.Context, *IDRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterItemsServer(s *grpc.Server, srv ItemsServer) {
	s.RegisterService(&_Items_serviceDesc, srv)
}

func _Items_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).ListAll(ctx, req.(*ListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_ListByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).ListByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/ListByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).ListByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_ListBySKU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).ListBySKU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/ListBySKU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).ListBySKU(ctx, req.(*UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_ListByVendorCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).ListByVendorCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/ListByVendorCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).ListByVendorCode(ctx, req.(*UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).Add(ctx, req.(*CreateOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).Update(ctx, req.(*CreateOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Items_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Items/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServer).Delete(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Items_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Items",
	HandlerType: (*ItemsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAll",
			Handler:    _Items_ListAll_Handler,
		},
		{
			MethodName: "ListByID",
			Handler:    _Items_ListByID_Handler,
		},
		{
			MethodName: "ListBySKU",
			Handler:    _Items_ListBySKU_Handler,
		},
		{
			MethodName: "ListByVendorCode",
			Handler:    _Items_ListByVendorCode_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Items_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Items_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Items_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "items.proto",
}

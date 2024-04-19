// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: finances/finances.proto

package finances_v1

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

type Type int32

const (
	Type_week  Type = 0
	Type_month Type = 1
	Type_year  Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "week",
		1: "month",
		2: "year",
	}
	Type_value = map[string]int32{
		"week":  0,
		"month": 1,
		"year":  2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_finances_finances_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_finances_finances_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{0}
}

type ExpenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Who         string `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
	Category    string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Date        string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ExpenseRequest) Reset() {
	*x = ExpenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseRequest) ProtoMessage() {}

func (x *ExpenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseRequest.ProtoReflect.Descriptor instead.
func (*ExpenseRequest) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{0}
}

func (x *ExpenseRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ExpenseRequest) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *ExpenseRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ExpenseRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ExpenseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ExpenseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ExpenseResponse) Reset() {
	*x = ExpenseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpenseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpenseResponse) ProtoMessage() {}

func (x *ExpenseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpenseResponse.ProtoReflect.Descriptor instead.
func (*ExpenseResponse) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{1}
}

func (x *ExpenseResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ExpensesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datefrom string `protobuf:"bytes,1,opt,name=datefrom,proto3" json:"datefrom,omitempty"`
	Dateto   string `protobuf:"bytes,2,opt,name=dateto,proto3" json:"dateto,omitempty"`
}

func (x *ExpensesListRequest) Reset() {
	*x = ExpensesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpensesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpensesListRequest) ProtoMessage() {}

func (x *ExpensesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpensesListRequest.ProtoReflect.Descriptor instead.
func (*ExpensesListRequest) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{2}
}

func (x *ExpensesListRequest) GetDatefrom() string {
	if x != nil {
		return x.Datefrom
	}
	return ""
}

func (x *ExpensesListRequest) GetDateto() string {
	if x != nil {
		return x.Dateto
	}
	return ""
}

type ExpensesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expenses []*Expense `protobuf:"bytes,1,rep,name=expenses,proto3" json:"expenses,omitempty"`
}

func (x *ExpensesListResponse) Reset() {
	*x = ExpensesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpensesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpensesListResponse) ProtoMessage() {}

func (x *ExpensesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpensesListResponse.ProtoReflect.Descriptor instead.
func (*ExpensesListResponse) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{3}
}

func (x *ExpensesListResponse) GetExpenses() []*Expense {
	if x != nil {
		return x.Expenses
	}
	return nil
}

type Expense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Who      string `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Date     string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Expense) Reset() {
	*x = Expense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Expense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expense) ProtoMessage() {}

func (x *Expense) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expense.ProtoReflect.Descriptor instead.
func (*Expense) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{4}
}

func (x *Expense) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Expense) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *Expense) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Expense) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCategoryResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type CategoriesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datefrom string `protobuf:"bytes,1,opt,name=datefrom,proto3" json:"datefrom,omitempty"`
	Dateto   string `protobuf:"bytes,2,opt,name=dateto,proto3" json:"dateto,omitempty"`
}

func (x *CategoriesListRequest) Reset() {
	*x = CategoriesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesListRequest) ProtoMessage() {}

func (x *CategoriesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesListRequest.ProtoReflect.Descriptor instead.
func (*CategoriesListRequest) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{7}
}

func (x *CategoriesListRequest) GetDatefrom() string {
	if x != nil {
		return x.Datefrom
	}
	return ""
}

func (x *CategoriesListRequest) GetDateto() string {
	if x != nil {
		return x.Dateto
	}
	return ""
}

type CategoriesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoriesListResponse) Reset() {
	*x = CategoriesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesListResponse) ProtoMessage() {}

func (x *CategoriesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesListResponse.ProtoReflect.Descriptor instead.
func (*CategoriesListResponse) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{8}
}

func (x *CategoriesListResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{9}
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Category) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=finances.Type" json:"type,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{10}
}

func (x *ReportRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_week
}

type ReportCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Percent int64  `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReportCategory) Reset() {
	*x = ReportCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportCategory) ProtoMessage() {}

func (x *ReportCategory) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportCategory.ProtoReflect.Descriptor instead.
func (*ReportCategory) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{11}
}

func (x *ReportCategory) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReportCategory) GetPercent() int64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

func (x *ReportCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total      int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Categories []*ReportCategory `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finances_finances_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_finances_finances_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_finances_finances_proto_rawDescGZIP(), []int{12}
}

func (x *ReportResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ReportResponse) GetCategories() []*ReportCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_finances_finances_proto protoreflect.FileDescriptor

var file_finances_finances_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x49, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f,
	0x22, 0x45, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x68,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x22, 0x4b, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x74, 0x6f,
	0x22, 0x4c, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x4a,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x0d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x56, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x38, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2a, 0x25, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x10, 0x02,
	0x32, 0x80, 0x03, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a,
	0x07, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x0c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e,
	0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x17, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_finances_finances_proto_rawDescOnce sync.Once
	file_finances_finances_proto_rawDescData = file_finances_finances_proto_rawDesc
)

func file_finances_finances_proto_rawDescGZIP() []byte {
	file_finances_finances_proto_rawDescOnce.Do(func() {
		file_finances_finances_proto_rawDescData = protoimpl.X.CompressGZIP(file_finances_finances_proto_rawDescData)
	})
	return file_finances_finances_proto_rawDescData
}

var file_finances_finances_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_finances_finances_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_finances_finances_proto_goTypes = []interface{}{
	(Type)(0),                      // 0: finances.Type
	(*ExpenseRequest)(nil),         // 1: finances.ExpenseRequest
	(*ExpenseResponse)(nil),        // 2: finances.ExpenseResponse
	(*ExpensesListRequest)(nil),    // 3: finances.ExpensesListRequest
	(*ExpensesListResponse)(nil),   // 4: finances.ExpensesListResponse
	(*Expense)(nil),                // 5: finances.Expense
	(*CreateCategoryRequest)(nil),  // 6: finances.CreateCategoryRequest
	(*CreateCategoryResponse)(nil), // 7: finances.CreateCategoryResponse
	(*CategoriesListRequest)(nil),  // 8: finances.CategoriesListRequest
	(*CategoriesListResponse)(nil), // 9: finances.CategoriesListResponse
	(*Category)(nil),               // 10: finances.Category
	(*ReportRequest)(nil),          // 11: finances.ReportRequest
	(*ReportCategory)(nil),         // 12: finances.ReportCategory
	(*ReportResponse)(nil),         // 13: finances.ReportResponse
}
var file_finances_finances_proto_depIdxs = []int32{
	5,  // 0: finances.ExpensesListResponse.expenses:type_name -> finances.Expense
	10, // 1: finances.CategoriesListResponse.categories:type_name -> finances.Category
	0,  // 2: finances.ReportRequest.type:type_name -> finances.Type
	12, // 3: finances.ReportResponse.categories:type_name -> finances.ReportCategory
	1,  // 4: finances.Finances.Expense:input_type -> finances.ExpenseRequest
	3,  // 5: finances.Finances.ExpensesList:input_type -> finances.ExpensesListRequest
	6,  // 6: finances.Finances.CreateCategory:input_type -> finances.CreateCategoryRequest
	8,  // 7: finances.Finances.CategoriesList:input_type -> finances.CategoriesListRequest
	11, // 8: finances.Finances.Report:input_type -> finances.ReportRequest
	2,  // 9: finances.Finances.Expense:output_type -> finances.ExpenseResponse
	4,  // 10: finances.Finances.ExpensesList:output_type -> finances.ExpensesListResponse
	7,  // 11: finances.Finances.CreateCategory:output_type -> finances.CreateCategoryResponse
	9,  // 12: finances.Finances.CategoriesList:output_type -> finances.CategoriesListResponse
	13, // 13: finances.Finances.Report:output_type -> finances.ReportResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_finances_finances_proto_init() }
func file_finances_finances_proto_init() {
	if File_finances_finances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_finances_finances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpenseRequest); i {
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
		file_finances_finances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpenseResponse); i {
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
		file_finances_finances_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpensesListRequest); i {
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
		file_finances_finances_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpensesListResponse); i {
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
		file_finances_finances_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Expense); i {
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
		file_finances_finances_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
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
		file_finances_finances_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_finances_finances_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesListRequest); i {
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
		file_finances_finances_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesListResponse); i {
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
		file_finances_finances_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_finances_finances_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_finances_finances_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportCategory); i {
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
		file_finances_finances_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
			RawDescriptor: file_finances_finances_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_finances_finances_proto_goTypes,
		DependencyIndexes: file_finances_finances_proto_depIdxs,
		EnumInfos:         file_finances_finances_proto_enumTypes,
		MessageInfos:      file_finances_finances_proto_msgTypes,
	}.Build()
	File_finances_finances_proto = out.File
	file_finances_finances_proto_rawDesc = nil
	file_finances_finances_proto_goTypes = nil
	file_finances_finances_proto_depIdxs = nil
}

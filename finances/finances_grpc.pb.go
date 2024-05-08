// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: finances/finances.proto

package finances_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Finances_Expense_FullMethodName        = "/finances.Finances/Expense"
	Finances_ExpenseEdit_FullMethodName    = "/finances.Finances/ExpenseEdit"
	Finances_ExpensesList_FullMethodName   = "/finances.Finances/ExpensesList"
	Finances_CreateCategory_FullMethodName = "/finances.Finances/CreateCategory"
	Finances_CategoriesList_FullMethodName = "/finances.Finances/CategoriesList"
	Finances_Report_FullMethodName         = "/finances.Finances/Report"
	Finances_MassiveReport_FullMethodName  = "/finances.Finances/MassiveReport"
)

// FinancesClient is the client API for Finances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinancesClient interface {
	// Expense adds an expense.
	Expense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	ExpenseEdit(ctx context.Context, in *ExpenseEditRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	ExpensesList(ctx context.Context, in *ExpensesListRequest, opts ...grpc.CallOption) (*ExpensesListResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	CategoriesList(ctx context.Context, in *CategoriesListRequest, opts ...grpc.CallOption) (*CategoriesListResponse, error)
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	MassiveReport(ctx context.Context, in *MassiveReportRequest, opts ...grpc.CallOption) (*MassiveReportResponse, error)
}

type financesClient struct {
	cc grpc.ClientConnInterface
}

func NewFinancesClient(cc grpc.ClientConnInterface) FinancesClient {
	return &financesClient{cc}
}

func (c *financesClient) Expense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, Finances_Expense_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) ExpenseEdit(ctx context.Context, in *ExpenseEditRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, Finances_ExpenseEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) ExpensesList(ctx context.Context, in *ExpensesListRequest, opts ...grpc.CallOption) (*ExpensesListResponse, error) {
	out := new(ExpensesListResponse)
	err := c.cc.Invoke(ctx, Finances_ExpensesList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, Finances_CreateCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) CategoriesList(ctx context.Context, in *CategoriesListRequest, opts ...grpc.CallOption) (*CategoriesListResponse, error) {
	out := new(CategoriesListResponse)
	err := c.cc.Invoke(ctx, Finances_CategoriesList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, Finances_Report_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financesClient) MassiveReport(ctx context.Context, in *MassiveReportRequest, opts ...grpc.CallOption) (*MassiveReportResponse, error) {
	out := new(MassiveReportResponse)
	err := c.cc.Invoke(ctx, Finances_MassiveReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinancesServer is the server API for Finances service.
// All implementations must embed UnimplementedFinancesServer
// for forward compatibility
type FinancesServer interface {
	// Expense adds an expense.
	Expense(context.Context, *ExpenseRequest) (*ExpenseResponse, error)
	ExpenseEdit(context.Context, *ExpenseEditRequest) (*ExpenseResponse, error)
	ExpensesList(context.Context, *ExpensesListRequest) (*ExpensesListResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	CategoriesList(context.Context, *CategoriesListRequest) (*CategoriesListResponse, error)
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	MassiveReport(context.Context, *MassiveReportRequest) (*MassiveReportResponse, error)
	mustEmbedUnimplementedFinancesServer()
}

// UnimplementedFinancesServer must be embedded to have forward compatible implementations.
type UnimplementedFinancesServer struct {
}

func (UnimplementedFinancesServer) Expense(context.Context, *ExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expense not implemented")
}
func (UnimplementedFinancesServer) ExpenseEdit(context.Context, *ExpenseEditRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpenseEdit not implemented")
}
func (UnimplementedFinancesServer) ExpensesList(context.Context, *ExpensesListRequest) (*ExpensesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpensesList not implemented")
}
func (UnimplementedFinancesServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedFinancesServer) CategoriesList(context.Context, *CategoriesListRequest) (*CategoriesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoriesList not implemented")
}
func (UnimplementedFinancesServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedFinancesServer) MassiveReport(context.Context, *MassiveReportRequest) (*MassiveReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MassiveReport not implemented")
}
func (UnimplementedFinancesServer) mustEmbedUnimplementedFinancesServer() {}

// UnsafeFinancesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinancesServer will
// result in compilation errors.
type UnsafeFinancesServer interface {
	mustEmbedUnimplementedFinancesServer()
}

func RegisterFinancesServer(s grpc.ServiceRegistrar, srv FinancesServer) {
	s.RegisterService(&Finances_ServiceDesc, srv)
}

func _Finances_Expense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).Expense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_Expense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).Expense(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_ExpenseEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).ExpenseEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_ExpenseEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).ExpenseEdit(ctx, req.(*ExpenseEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_ExpensesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpensesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).ExpensesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_ExpensesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).ExpensesList(ctx, req.(*ExpensesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_CategoriesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoriesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).CategoriesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_CategoriesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).CategoriesList(ctx, req.(*CategoriesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_Report_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finances_MassiveReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MassiveReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinancesServer).MassiveReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Finances_MassiveReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinancesServer).MassiveReport(ctx, req.(*MassiveReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Finances_ServiceDesc is the grpc.ServiceDesc for Finances service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Finances_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finances.Finances",
	HandlerType: (*FinancesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expense",
			Handler:    _Finances_Expense_Handler,
		},
		{
			MethodName: "ExpenseEdit",
			Handler:    _Finances_ExpenseEdit_Handler,
		},
		{
			MethodName: "ExpensesList",
			Handler:    _Finances_ExpensesList_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Finances_CreateCategory_Handler,
		},
		{
			MethodName: "CategoriesList",
			Handler:    _Finances_CategoriesList_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _Finances_Report_Handler,
		},
		{
			MethodName: "MassiveReport",
			Handler:    _Finances_MassiveReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finances/finances.proto",
}

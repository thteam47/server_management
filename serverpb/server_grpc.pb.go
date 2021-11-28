// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package serverpb

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

// ServerServiceClient is the client API for ServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerServiceClient interface {
	Login(ctx context.Context, in *LoginServer, opts ...grpc.CallOption) (*ResultLogin, error)
	GetUser(ctx context.Context, in *InfoUser, opts ...grpc.CallOption) (*User, error)
	GetListUser(ctx context.Context, in *GetListUser, opts ...grpc.CallOption) (*ListUser, error)
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	Connect(ctx context.Context, in *LoginServer, opts ...grpc.CallOption) (*MessResponse, error)
	Index(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*ListServer, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ListServer, error)
	CheckServerName(ctx context.Context, in *CheckServerNameRequest, opts ...grpc.CallOption) (*CheckServerNameResponse, error)
	AddServer(ctx context.Context, in *Server, opts ...grpc.CallOption) (*Server, error)
	UpdateServer(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Server, error)
	DetailsServer(ctx context.Context, in *DetailsServer, opts ...grpc.CallOption) (*DetailsServerResponse, error)
	DeleteServer(ctx context.Context, in *DelServer, opts ...grpc.CallOption) (*DeleteServerResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*MessResponse, error)
	ChangeActionUser(ctx context.Context, in *ChangeActionUser, opts ...grpc.CallOption) (*MessResponse, error)
	CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error)
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
	Logout(ctx context.Context, in *Logout, opts ...grpc.CallOption) (*MessResponse, error)
	Disconnect(ctx context.Context, in *Disconnect, opts ...grpc.CallOption) (*MessResponse, error)
	UpdateUser(ctx context.Context, in *ChangeUser, opts ...grpc.CallOption) (*UserResponse, error)
	ChangePassUser(ctx context.Context, in *ChangePasswordUser, opts ...grpc.CallOption) (*MessResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUser, opts ...grpc.CallOption) (*MessResponse, error)
}

type serverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerServiceClient(cc grpc.ClientConnInterface) ServerServiceClient {
	return &serverServiceClient{cc}
}

func (c *serverServiceClient) Login(ctx context.Context, in *LoginServer, opts ...grpc.CallOption) (*ResultLogin, error) {
	out := new(ResultLogin)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) GetUser(ctx context.Context, in *InfoUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) GetListUser(ctx context.Context, in *GetListUser, opts ...grpc.CallOption) (*ListUser, error) {
	out := new(ListUser)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/getListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/addUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Connect(ctx context.Context, in *LoginServer, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Index(ctx context.Context, in *PaginationRequest, opts ...grpc.CallOption) (*ListServer, error) {
	out := new(ListServer)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ListServer, error) {
	out := new(ListServer)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) CheckServerName(ctx context.Context, in *CheckServerNameRequest, opts ...grpc.CallOption) (*CheckServerNameResponse, error) {
	out := new(CheckServerNameResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/checkServerName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) AddServer(ctx context.Context, in *Server, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/addServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) UpdateServer(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Server, error) {
	out := new(Server)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/updateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) DetailsServer(ctx context.Context, in *DetailsServer, opts ...grpc.CallOption) (*DetailsServerResponse, error) {
	out := new(DetailsServerResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/detailsServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) DeleteServer(ctx context.Context, in *DelServer, opts ...grpc.CallOption) (*DeleteServerResponse, error) {
	out := new(DeleteServerResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/deleteServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/changePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ChangeActionUser(ctx context.Context, in *ChangeActionUser, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/changeActionUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/checkStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Logout(ctx context.Context, in *Logout, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) Disconnect(ctx context.Context, in *Disconnect, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) UpdateUser(ctx context.Context, in *ChangeUser, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ChangePassUser(ctx context.Context, in *ChangePasswordUser, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/changePassUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) DeleteUser(ctx context.Context, in *DeleteUser, opts ...grpc.CallOption) (*MessResponse, error) {
	out := new(MessResponse)
	err := c.cc.Invoke(ctx, "/server.management.ServerService/deleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServiceServer is the server API for ServerService service.
// All implementations must embed UnimplementedServerServiceServer
// for forward compatibility
type ServerServiceServer interface {
	Login(context.Context, *LoginServer) (*ResultLogin, error)
	GetUser(context.Context, *InfoUser) (*User, error)
	GetListUser(context.Context, *GetListUser) (*ListUser, error)
	AddUser(context.Context, *User) (*User, error)
	Connect(context.Context, *LoginServer) (*MessResponse, error)
	Index(context.Context, *PaginationRequest) (*ListServer, error)
	Search(context.Context, *SearchRequest) (*ListServer, error)
	CheckServerName(context.Context, *CheckServerNameRequest) (*CheckServerNameResponse, error)
	AddServer(context.Context, *Server) (*Server, error)
	UpdateServer(context.Context, *UpdateRequest) (*Server, error)
	DetailsServer(context.Context, *DetailsServer) (*DetailsServerResponse, error)
	DeleteServer(context.Context, *DelServer) (*DeleteServerResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*MessResponse, error)
	ChangeActionUser(context.Context, *ChangeActionUser) (*MessResponse, error)
	CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error)
	Export(context.Context, *ExportRequest) (*ExportResponse, error)
	Logout(context.Context, *Logout) (*MessResponse, error)
	Disconnect(context.Context, *Disconnect) (*MessResponse, error)
	UpdateUser(context.Context, *ChangeUser) (*UserResponse, error)
	ChangePassUser(context.Context, *ChangePasswordUser) (*MessResponse, error)
	DeleteUser(context.Context, *DeleteUser) (*MessResponse, error)
	mustEmbedUnimplementedServerServiceServer()
}

// UnimplementedServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerServiceServer struct {
}

func (UnimplementedServerServiceServer) Login(context.Context, *LoginServer) (*ResultLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedServerServiceServer) GetUser(context.Context, *InfoUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedServerServiceServer) GetListUser(context.Context, *GetListUser) (*ListUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUser not implemented")
}
func (UnimplementedServerServiceServer) AddUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedServerServiceServer) Connect(context.Context, *LoginServer) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedServerServiceServer) Index(context.Context, *PaginationRequest) (*ListServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedServerServiceServer) Search(context.Context, *SearchRequest) (*ListServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedServerServiceServer) CheckServerName(context.Context, *CheckServerNameRequest) (*CheckServerNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckServerName not implemented")
}
func (UnimplementedServerServiceServer) AddServer(context.Context, *Server) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddServer not implemented")
}
func (UnimplementedServerServiceServer) UpdateServer(context.Context, *UpdateRequest) (*Server, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (UnimplementedServerServiceServer) DetailsServer(context.Context, *DetailsServer) (*DetailsServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailsServer not implemented")
}
func (UnimplementedServerServiceServer) DeleteServer(context.Context, *DelServer) (*DeleteServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedServerServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedServerServiceServer) ChangeActionUser(context.Context, *ChangeActionUser) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeActionUser not implemented")
}
func (UnimplementedServerServiceServer) CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedServerServiceServer) Export(context.Context, *ExportRequest) (*ExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedServerServiceServer) Logout(context.Context, *Logout) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedServerServiceServer) Disconnect(context.Context, *Disconnect) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedServerServiceServer) UpdateUser(context.Context, *ChangeUser) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedServerServiceServer) ChangePassUser(context.Context, *ChangePasswordUser) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassUser not implemented")
}
func (UnimplementedServerServiceServer) DeleteUser(context.Context, *DeleteUser) (*MessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedServerServiceServer) mustEmbedUnimplementedServerServiceServer() {}

// UnsafeServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServiceServer will
// result in compilation errors.
type UnsafeServerServiceServer interface {
	mustEmbedUnimplementedServerServiceServer()
}

func RegisterServerServiceServer(s grpc.ServiceRegistrar, srv ServerServiceServer) {
	s.RegisterService(&ServerService_ServiceDesc, srv)
}

func _ServerService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Login(ctx, req.(*LoginServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetUser(ctx, req.(*InfoUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_GetListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/getListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetListUser(ctx, req.(*GetListUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/addUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Connect(ctx, req.(*LoginServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Index(ctx, req.(*PaginationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_CheckServerName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckServerNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).CheckServerName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/checkServerName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).CheckServerName(ctx, req.(*CheckServerNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_AddServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Server)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).AddServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/addServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).AddServer(ctx, req.(*Server))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/updateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).UpdateServer(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_DetailsServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailsServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).DetailsServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/detailsServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).DetailsServer(ctx, req.(*DetailsServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/deleteServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).DeleteServer(ctx, req.(*DelServer))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/changePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ChangeActionUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeActionUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ChangeActionUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/changeActionUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ChangeActionUser(ctx, req.(*ChangeActionUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/checkStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).CheckStatus(ctx, req.(*CheckStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Export(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Logout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Logout(ctx, req.(*Logout))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Disconnect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).Disconnect(ctx, req.(*Disconnect))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/updateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).UpdateUser(ctx, req.(*ChangeUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ChangePassUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ChangePassUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/changePassUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ChangePassUser(ctx, req.(*ChangePasswordUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.management.ServerService/deleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).DeleteUser(ctx, req.(*DeleteUser))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerService_ServiceDesc is the grpc.ServiceDesc for ServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.management.ServerService",
	HandlerType: (*ServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _ServerService_Login_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _ServerService_GetUser_Handler,
		},
		{
			MethodName: "getListUser",
			Handler:    _ServerService_GetListUser_Handler,
		},
		{
			MethodName: "addUser",
			Handler:    _ServerService_AddUser_Handler,
		},
		{
			MethodName: "connect",
			Handler:    _ServerService_Connect_Handler,
		},
		{
			MethodName: "index",
			Handler:    _ServerService_Index_Handler,
		},
		{
			MethodName: "search",
			Handler:    _ServerService_Search_Handler,
		},
		{
			MethodName: "checkServerName",
			Handler:    _ServerService_CheckServerName_Handler,
		},
		{
			MethodName: "addServer",
			Handler:    _ServerService_AddServer_Handler,
		},
		{
			MethodName: "updateServer",
			Handler:    _ServerService_UpdateServer_Handler,
		},
		{
			MethodName: "detailsServer",
			Handler:    _ServerService_DetailsServer_Handler,
		},
		{
			MethodName: "deleteServer",
			Handler:    _ServerService_DeleteServer_Handler,
		},
		{
			MethodName: "changePassword",
			Handler:    _ServerService_ChangePassword_Handler,
		},
		{
			MethodName: "changeActionUser",
			Handler:    _ServerService_ChangeActionUser_Handler,
		},
		{
			MethodName: "checkStatus",
			Handler:    _ServerService_CheckStatus_Handler,
		},
		{
			MethodName: "export",
			Handler:    _ServerService_Export_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _ServerService_Logout_Handler,
		},
		{
			MethodName: "disconnect",
			Handler:    _ServerService_Disconnect_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _ServerService_UpdateUser_Handler,
		},
		{
			MethodName: "changePassUser",
			Handler:    _ServerService_ChangePassUser_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _ServerService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server.proto",
}

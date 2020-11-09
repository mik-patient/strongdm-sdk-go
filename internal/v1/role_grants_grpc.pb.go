// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RoleGrantsClient is the client API for RoleGrants service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleGrantsClient interface {
	// Create registers a new RoleGrant.
	Create(ctx context.Context, in *RoleGrantCreateRequest, opts ...grpc.CallOption) (*RoleGrantCreateResponse, error)
	// Get reads one RoleGrant by ID.
	Get(ctx context.Context, in *RoleGrantGetRequest, opts ...grpc.CallOption) (*RoleGrantGetResponse, error)
	// Delete removes a RoleGrant by ID.
	Delete(ctx context.Context, in *RoleGrantDeleteRequest, opts ...grpc.CallOption) (*RoleGrantDeleteResponse, error)
	// List gets a list of RoleGrants matching a given set of criteria.
	List(ctx context.Context, in *RoleGrantListRequest, opts ...grpc.CallOption) (*RoleGrantListResponse, error)
}

type roleGrantsClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleGrantsClient(cc grpc.ClientConnInterface) RoleGrantsClient {
	return &roleGrantsClient{cc}
}

func (c *roleGrantsClient) Create(ctx context.Context, in *RoleGrantCreateRequest, opts ...grpc.CallOption) (*RoleGrantCreateResponse, error) {
	out := new(RoleGrantCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) Get(ctx context.Context, in *RoleGrantGetRequest, opts ...grpc.CallOption) (*RoleGrantGetResponse, error) {
	out := new(RoleGrantGetResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) Delete(ctx context.Context, in *RoleGrantDeleteRequest, opts ...grpc.CallOption) (*RoleGrantDeleteResponse, error) {
	out := new(RoleGrantDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleGrantsClient) List(ctx context.Context, in *RoleGrantListRequest, opts ...grpc.CallOption) (*RoleGrantListResponse, error) {
	out := new(RoleGrantListResponse)
	err := c.cc.Invoke(ctx, "/v1.RoleGrants/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleGrantsServer is the server API for RoleGrants service.
// All implementations must embed UnimplementedRoleGrantsServer
// for forward compatibility
type RoleGrantsServer interface {
	// Create registers a new RoleGrant.
	Create(context.Context, *RoleGrantCreateRequest) (*RoleGrantCreateResponse, error)
	// Get reads one RoleGrant by ID.
	Get(context.Context, *RoleGrantGetRequest) (*RoleGrantGetResponse, error)
	// Delete removes a RoleGrant by ID.
	Delete(context.Context, *RoleGrantDeleteRequest) (*RoleGrantDeleteResponse, error)
	// List gets a list of RoleGrants matching a given set of criteria.
	List(context.Context, *RoleGrantListRequest) (*RoleGrantListResponse, error)
	mustEmbedUnimplementedRoleGrantsServer()
}

// UnimplementedRoleGrantsServer must be embedded to have forward compatible implementations.
type UnimplementedRoleGrantsServer struct {
}

func (UnimplementedRoleGrantsServer) Create(context.Context, *RoleGrantCreateRequest) (*RoleGrantCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoleGrantsServer) Get(context.Context, *RoleGrantGetRequest) (*RoleGrantGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoleGrantsServer) Delete(context.Context, *RoleGrantDeleteRequest) (*RoleGrantDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRoleGrantsServer) List(context.Context, *RoleGrantListRequest) (*RoleGrantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRoleGrantsServer) mustEmbedUnimplementedRoleGrantsServer() {}

// UnsafeRoleGrantsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleGrantsServer will
// result in compilation errors.
type UnsafeRoleGrantsServer interface {
	mustEmbedUnimplementedRoleGrantsServer()
}

func RegisterRoleGrantsServer(s *grpc.Server, srv RoleGrantsServer) {
	s.RegisterService(&_RoleGrants_serviceDesc, srv)
}

func _RoleGrants_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Create(ctx, req.(*RoleGrantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Get(ctx, req.(*RoleGrantGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).Delete(ctx, req.(*RoleGrantDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleGrants_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGrantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleGrantsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RoleGrants/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleGrantsServer).List(ctx, req.(*RoleGrantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleGrants_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RoleGrants",
	HandlerType: (*RoleGrantsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoleGrants_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoleGrants_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleGrants_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RoleGrants_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_grants.proto",
}

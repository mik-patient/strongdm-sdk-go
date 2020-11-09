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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	// Create registers a new Account.
	Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error)
	// Get reads one Account by ID.
	Get(ctx context.Context, in *AccountGetRequest, opts ...grpc.CallOption) (*AccountGetResponse, error)
	// Update patches a Account by ID.
	Update(ctx context.Context, in *AccountUpdateRequest, opts ...grpc.CallOption) (*AccountUpdateResponse, error)
	// Delete removes a Account by ID.
	Delete(ctx context.Context, in *AccountDeleteRequest, opts ...grpc.CallOption) (*AccountDeleteResponse, error)
	// List gets a list of Accounts matching a given set of criteria.
	List(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*AccountListResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error) {
	out := new(AccountCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Get(ctx context.Context, in *AccountGetRequest, opts ...grpc.CallOption) (*AccountGetResponse, error) {
	out := new(AccountGetResponse)
	err := c.cc.Invoke(ctx, "/v1.Accounts/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Update(ctx context.Context, in *AccountUpdateRequest, opts ...grpc.CallOption) (*AccountUpdateResponse, error) {
	out := new(AccountUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.Accounts/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Delete(ctx context.Context, in *AccountDeleteRequest, opts ...grpc.CallOption) (*AccountDeleteResponse, error) {
	out := new(AccountDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.Accounts/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) List(ctx context.Context, in *AccountListRequest, opts ...grpc.CallOption) (*AccountListResponse, error) {
	out := new(AccountListResponse)
	err := c.cc.Invoke(ctx, "/v1.Accounts/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// Create registers a new Account.
	Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error)
	// Get reads one Account by ID.
	Get(context.Context, *AccountGetRequest) (*AccountGetResponse, error)
	// Update patches a Account by ID.
	Update(context.Context, *AccountUpdateRequest) (*AccountUpdateResponse, error)
	// Delete removes a Account by ID.
	Delete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error)
	// List gets a list of Accounts matching a given set of criteria.
	List(context.Context, *AccountListRequest) (*AccountListResponse, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServer) Get(context.Context, *AccountGetRequest) (*AccountGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccountsServer) Update(context.Context, *AccountUpdateRequest) (*AccountUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccountsServer) Delete(context.Context, *AccountDeleteRequest) (*AccountDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountsServer) List(context.Context, *AccountListRequest) (*AccountListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Accounts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Create(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Accounts/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Get(ctx, req.(*AccountGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Accounts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Update(ctx, req.(*AccountUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Accounts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Delete(ctx, req.(*AccountDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.Accounts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).List(ctx, req.(*AccountListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Accounts_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Accounts_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Accounts_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Accounts_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Accounts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}

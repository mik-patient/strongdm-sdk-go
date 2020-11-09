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

// AccessRulesClient is the client API for AccessRules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessRulesClient interface {
	// Plan registers a new AccessRule.
	Plan(ctx context.Context, in *AccessRulePlanRequest, opts ...grpc.CallOption) (*AccessRulePlanResponse, error)
	// Create registers a new AccessRule.
	Create(ctx context.Context, in *AccessRuleCreateRequest, opts ...grpc.CallOption) (*AccessRuleCreateResponse, error)
	// Get reads one AccessRule by ID.
	Get(ctx context.Context, in *AccessRuleGetRequest, opts ...grpc.CallOption) (*AccessRuleGetResponse, error)
	// Update patches a AccessRule by ID.
	Update(ctx context.Context, in *AccessRuleUpdateRequest, opts ...grpc.CallOption) (*AccessRuleUpdateResponse, error)
	// Delete removes a AccessRule by ID.
	Delete(ctx context.Context, in *AccessRuleDeleteRequest, opts ...grpc.CallOption) (*AccessRuleDeleteResponse, error)
	// List gets a list of Access Rules matching a given set of criteria.
	List(ctx context.Context, in *AccessRuleListRequest, opts ...grpc.CallOption) (*AccessRuleListResponse, error)
}

type accessRulesClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessRulesClient(cc grpc.ClientConnInterface) AccessRulesClient {
	return &accessRulesClient{cc}
}

func (c *accessRulesClient) Plan(ctx context.Context, in *AccessRulePlanRequest, opts ...grpc.CallOption) (*AccessRulePlanResponse, error) {
	out := new(AccessRulePlanResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/Plan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRulesClient) Create(ctx context.Context, in *AccessRuleCreateRequest, opts ...grpc.CallOption) (*AccessRuleCreateResponse, error) {
	out := new(AccessRuleCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRulesClient) Get(ctx context.Context, in *AccessRuleGetRequest, opts ...grpc.CallOption) (*AccessRuleGetResponse, error) {
	out := new(AccessRuleGetResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRulesClient) Update(ctx context.Context, in *AccessRuleUpdateRequest, opts ...grpc.CallOption) (*AccessRuleUpdateResponse, error) {
	out := new(AccessRuleUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRulesClient) Delete(ctx context.Context, in *AccessRuleDeleteRequest, opts ...grpc.CallOption) (*AccessRuleDeleteResponse, error) {
	out := new(AccessRuleDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessRulesClient) List(ctx context.Context, in *AccessRuleListRequest, opts ...grpc.CallOption) (*AccessRuleListResponse, error) {
	out := new(AccessRuleListResponse)
	err := c.cc.Invoke(ctx, "/v1.AccessRules/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessRulesServer is the server API for AccessRules service.
// All implementations must embed UnimplementedAccessRulesServer
// for forward compatibility
type AccessRulesServer interface {
	// Plan registers a new AccessRule.
	Plan(context.Context, *AccessRulePlanRequest) (*AccessRulePlanResponse, error)
	// Create registers a new AccessRule.
	Create(context.Context, *AccessRuleCreateRequest) (*AccessRuleCreateResponse, error)
	// Get reads one AccessRule by ID.
	Get(context.Context, *AccessRuleGetRequest) (*AccessRuleGetResponse, error)
	// Update patches a AccessRule by ID.
	Update(context.Context, *AccessRuleUpdateRequest) (*AccessRuleUpdateResponse, error)
	// Delete removes a AccessRule by ID.
	Delete(context.Context, *AccessRuleDeleteRequest) (*AccessRuleDeleteResponse, error)
	// List gets a list of Access Rules matching a given set of criteria.
	List(context.Context, *AccessRuleListRequest) (*AccessRuleListResponse, error)
	mustEmbedUnimplementedAccessRulesServer()
}

// UnimplementedAccessRulesServer must be embedded to have forward compatible implementations.
type UnimplementedAccessRulesServer struct {
}

func (UnimplementedAccessRulesServer) Plan(context.Context, *AccessRulePlanRequest) (*AccessRulePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plan not implemented")
}
func (UnimplementedAccessRulesServer) Create(context.Context, *AccessRuleCreateRequest) (*AccessRuleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccessRulesServer) Get(context.Context, *AccessRuleGetRequest) (*AccessRuleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccessRulesServer) Update(context.Context, *AccessRuleUpdateRequest) (*AccessRuleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAccessRulesServer) Delete(context.Context, *AccessRuleDeleteRequest) (*AccessRuleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccessRulesServer) List(context.Context, *AccessRuleListRequest) (*AccessRuleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccessRulesServer) mustEmbedUnimplementedAccessRulesServer() {}

// UnsafeAccessRulesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessRulesServer will
// result in compilation errors.
type UnsafeAccessRulesServer interface {
	mustEmbedUnimplementedAccessRulesServer()
}

func RegisterAccessRulesServer(s *grpc.Server, srv AccessRulesServer) {
	s.RegisterService(&_AccessRules_serviceDesc, srv)
}

func _AccessRules_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRulePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/Plan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).Plan(ctx, req.(*AccessRulePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRules_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).Create(ctx, req.(*AccessRuleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRules_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).Get(ctx, req.(*AccessRuleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRules_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).Update(ctx, req.(*AccessRuleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRules_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).Delete(ctx, req.(*AccessRuleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessRules_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRuleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessRulesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AccessRules/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessRulesServer).List(ctx, req.(*AccessRuleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccessRules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AccessRules",
	HandlerType: (*AccessRulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plan",
			Handler:    _AccessRules_Plan_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccessRules_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccessRules_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccessRules_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccessRules_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccessRules_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access_rules.proto",
}

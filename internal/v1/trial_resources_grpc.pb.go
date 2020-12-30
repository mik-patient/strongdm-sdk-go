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

// TrialResourcesClient is the client API for TrialResources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrialResourcesClient interface {
	// Create
	Create(ctx context.Context, in *TrialResourceCreateRequest, opts ...grpc.CallOption) (*TrialResourceCreateResponse, error)
	// ListForOrganization gets a list of TrialResources in your organization
	// matching a given set of criteria. This operation can be done by account
	// administrators.
	ListForOrganization(ctx context.Context, in *TrialResourceListForOrganizationRequest, opts ...grpc.CallOption) (*TrialResourceListForOrganizationResponse, error)
	// ListAll gets a list of TrialResources across all orgs matching a given
	// set of criteria. This operation can only be done by operators and the
	// trial provisioner.
	ListAll(ctx context.Context, in *TrialResourceListAllRequest, opts ...grpc.CallOption) (*TrialResourceListAllResponse, error)
	// Update updates a TrialResource.
	Update(ctx context.Context, in *TrialResourceUpdateRequest, opts ...grpc.CallOption) (*TrialResourceUpdateResponse, error)
}

type trialResourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewTrialResourcesClient(cc grpc.ClientConnInterface) TrialResourcesClient {
	return &trialResourcesClient{cc}
}

func (c *trialResourcesClient) Create(ctx context.Context, in *TrialResourceCreateRequest, opts ...grpc.CallOption) (*TrialResourceCreateResponse, error) {
	out := new(TrialResourceCreateResponse)
	err := c.cc.Invoke(ctx, "/v1.TrialResources/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialResourcesClient) ListForOrganization(ctx context.Context, in *TrialResourceListForOrganizationRequest, opts ...grpc.CallOption) (*TrialResourceListForOrganizationResponse, error) {
	out := new(TrialResourceListForOrganizationResponse)
	err := c.cc.Invoke(ctx, "/v1.TrialResources/ListForOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialResourcesClient) ListAll(ctx context.Context, in *TrialResourceListAllRequest, opts ...grpc.CallOption) (*TrialResourceListAllResponse, error) {
	out := new(TrialResourceListAllResponse)
	err := c.cc.Invoke(ctx, "/v1.TrialResources/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trialResourcesClient) Update(ctx context.Context, in *TrialResourceUpdateRequest, opts ...grpc.CallOption) (*TrialResourceUpdateResponse, error) {
	out := new(TrialResourceUpdateResponse)
	err := c.cc.Invoke(ctx, "/v1.TrialResources/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrialResourcesServer is the server API for TrialResources service.
// All implementations must embed UnimplementedTrialResourcesServer
// for forward compatibility
type TrialResourcesServer interface {
	// Create
	Create(context.Context, *TrialResourceCreateRequest) (*TrialResourceCreateResponse, error)
	// ListForOrganization gets a list of TrialResources in your organization
	// matching a given set of criteria. This operation can be done by account
	// administrators.
	ListForOrganization(context.Context, *TrialResourceListForOrganizationRequest) (*TrialResourceListForOrganizationResponse, error)
	// ListAll gets a list of TrialResources across all orgs matching a given
	// set of criteria. This operation can only be done by operators and the
	// trial provisioner.
	ListAll(context.Context, *TrialResourceListAllRequest) (*TrialResourceListAllResponse, error)
	// Update updates a TrialResource.
	Update(context.Context, *TrialResourceUpdateRequest) (*TrialResourceUpdateResponse, error)
	mustEmbedUnimplementedTrialResourcesServer()
}

// UnimplementedTrialResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedTrialResourcesServer struct {
}

func (UnimplementedTrialResourcesServer) Create(context.Context, *TrialResourceCreateRequest) (*TrialResourceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTrialResourcesServer) ListForOrganization(context.Context, *TrialResourceListForOrganizationRequest) (*TrialResourceListForOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForOrganization not implemented")
}
func (UnimplementedTrialResourcesServer) ListAll(context.Context, *TrialResourceListAllRequest) (*TrialResourceListAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (UnimplementedTrialResourcesServer) Update(context.Context, *TrialResourceUpdateRequest) (*TrialResourceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTrialResourcesServer) mustEmbedUnimplementedTrialResourcesServer() {}

// UnsafeTrialResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrialResourcesServer will
// result in compilation errors.
type UnsafeTrialResourcesServer interface {
	mustEmbedUnimplementedTrialResourcesServer()
}

func RegisterTrialResourcesServer(s grpc.ServiceRegistrar, srv TrialResourcesServer) {
	s.RegisterService(&_TrialResources_serviceDesc, srv)
}

func _TrialResources_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrialResourceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialResourcesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TrialResources/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialResourcesServer).Create(ctx, req.(*TrialResourceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrialResources_ListForOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrialResourceListForOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialResourcesServer).ListForOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TrialResources/ListForOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialResourcesServer).ListForOrganization(ctx, req.(*TrialResourceListForOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrialResources_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrialResourceListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialResourcesServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TrialResources/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialResourcesServer).ListAll(ctx, req.(*TrialResourceListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrialResources_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrialResourceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrialResourcesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TrialResources/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrialResourcesServer).Update(ctx, req.(*TrialResourceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrialResources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TrialResources",
	HandlerType: (*TrialResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TrialResources_Create_Handler,
		},
		{
			MethodName: "ListForOrganization",
			Handler:    _TrialResources_ListForOrganization_Handler,
		},
		{
			MethodName: "ListAll",
			Handler:    _TrialResources_ListAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TrialResources_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trial_resources.proto",
}

// Copyright 2020 StrongDM Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sdm

// This file was generated by protogen. DO NOT EDIT.

// Error is a generic RPC error indicating something went wrong at the
// transport layer. Use Code() and Unwrap() to inspect the actual failed
// condition.
type Error interface {
	// Code returns the gRPC error code
	Code() int
	error
}

// UnknownError is a generic wrapper that indicates an unknown internal error in the SDK.
type UnknownError struct {
	Wrapped error
}

func (e *UnknownError) Error() string {
	return e.Wrapped.Error()
}

func (e *UnknownError) Unwrap() error {
	return e.Wrapped
}

func (e *UnknownError) Code() int {
	return 2 // Unknown
}

// DeadlineExceededError indicates a timeout occurred.
type DeadlineExceededError struct {
	Wrapped error
}

func (e *DeadlineExceededError) Error() string {
	return "deadline exceeded"
}

func (e *DeadlineExceededError) Unwrap() error {
	return e.Wrapped
}

func (e *DeadlineExceededError) Code() int {
	return 4
}

// ContextCanceledError indicates an operation was canceled.
type ContextCanceledError struct {
	Wrapped error
}

func (e *ContextCanceledError) Error() string {
	return "context canceled"
}

func (e *ContextCanceledError) Unwrap() error {
	return e.Wrapped
}

func (e *ContextCanceledError) Code() int {
	return 1
}

// AlreadyExistsError is used when an entity already exists in the system
type AlreadyExistsError struct {
	Message string
}

func (e AlreadyExistsError) Error() string {
	return e.Message
}

func (e AlreadyExistsError) Code() int {
	return 6
}

// NotFoundError is used when an entity does not exist in the system
type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e NotFoundError) Code() int {
	return 5
}

// BadRequestError identifies a bad request sent by the client
type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

func (e BadRequestError) Code() int {
	return 3
}

// AuthenticationError is used to specify an authentication failure condition
type AuthenticationError struct {
	Message string
}

func (e AuthenticationError) Error() string {
	return e.Message
}

func (e AuthenticationError) Code() int {
	return 16
}

// PermissionError is used to specify a permissions violation
type PermissionError struct {
	Message string
}

func (e PermissionError) Error() string {
	return e.Message
}

func (e PermissionError) Code() int {
	return 7
}

// InternalError is used to specify an internal system error
type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

func (e InternalError) Code() int {
	return 13
}

// RateLimitError is used for rate limit excess condition
type RateLimitError struct {
	Message   string
	RateLimit *RateLimitMetadata
}

func (e RateLimitError) Error() string {
	return e.Message
}

func (e RateLimitError) Code() int {
	return 8
}

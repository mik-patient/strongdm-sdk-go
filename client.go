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

// This file was generated by protogen. DO NOT EDIT.

/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	plumbing "github.com/strongdm/strongdm-sdk-go/v2/internal/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	defaultAPIHost   = "api.strongdm.com:443"
	apiVersion       = "2021-08-23"
	defaultUserAgent = "strongdm-sdk-go/0.9.40"
)

var _ = metadata.Pairs

// Client is the strongDM API client implementation.
type Client struct {
	testOptionsMu sync.RWMutex
	testOptions   map[string]interface{}

	apiHost              string
	apiToken             string
	apiSecret            []byte
	apiInsecureTransport bool
	userAgent            string

	grpcConn *grpc.ClientConn

	maxRetries         int
	baseRetryDelay     time.Duration
	maxRetryDelay      time.Duration
	accountAttachments *AccountAttachments
	accountGrants      *AccountGrants
	accounts           *Accounts
	controlPanel       *ControlPanel
	nodes              *Nodes
	resources          *Resources
	roleAttachments    *RoleAttachments
	roleGrants         *RoleGrants
	roles              *Roles
	secretStores       *SecretStores
}

// New creates a new strongDM API client.
func New(token, secret string, opts ...ClientOption) (*Client, error) {
	token = strings.TrimSpace(token)
	secret = strings.TrimSpace(secret)
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("invalid secret: %w", err))
	}

	client := &Client{
		apiHost:        defaultAPIHost,
		maxRetries:     defaultMaxRetries,
		baseRetryDelay: defaultBaseRetryDelay,
		maxRetryDelay:  defaultMaxRetryDelay,
		testOptions:    map[string]interface{}{},
		apiToken:       token,
		apiSecret:      decodedSecret,
		userAgent:      defaultUserAgent,
	}

	for _, opt := range opts {
		opt(client)
	}

	var dialOpt grpc.DialOption
	if client.apiInsecureTransport {
		dialOpt = grpc.WithInsecure()
	} else {
		dialOpt = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: false,
		}))
	}
	cc, err := grpc.Dial(client.apiHost, dialOpt)
	if err != nil {
		return nil, convertErrorToPorcelain(fmt.Errorf("cannot dial API server: %w", err))
	}
	client.grpcConn = cc
	client.accountAttachments = &AccountAttachments{
		client: plumbing.NewAccountAttachmentsClient(client.grpcConn),
		parent: client,
	}
	client.accountGrants = &AccountGrants{
		client: plumbing.NewAccountGrantsClient(client.grpcConn),
		parent: client,
	}
	client.accounts = &Accounts{
		client: plumbing.NewAccountsClient(client.grpcConn),
		parent: client,
	}
	client.controlPanel = &ControlPanel{
		client: plumbing.NewControlPanelClient(client.grpcConn),
		parent: client,
	}
	client.nodes = &Nodes{
		client: plumbing.NewNodesClient(client.grpcConn),
		parent: client,
	}
	client.resources = &Resources{
		client: plumbing.NewResourcesClient(client.grpcConn),
		parent: client,
	}
	client.roleAttachments = &RoleAttachments{
		client: plumbing.NewRoleAttachmentsClient(client.grpcConn),
		parent: client,
	}
	client.roleGrants = &RoleGrants{
		client: plumbing.NewRoleGrantsClient(client.grpcConn),
		parent: client,
	}
	client.roles = &Roles{
		client: plumbing.NewRolesClient(client.grpcConn),
		parent: client,
	}
	client.secretStores = &SecretStores{
		client: plumbing.NewSecretStoresClient(client.grpcConn),
		parent: client,
	}
	return client, nil
}

type ClientOption func(c *Client)

func WithHost(host string) ClientOption {
	return func(c *Client) {
		c.apiHost = host
	}
}

func WithInsecure() ClientOption {
	return func(c *Client) {
		c.apiInsecureTransport = true
	}
}

func WithUserAgentExtra(userAgentExtra string) ClientOption {
	return func(c *Client) {
		c.userAgent += " " + userAgentExtra
	}
}

// AccountAttachments assign an account to a role or composite role.
func (c *Client) AccountAttachments() *AccountAttachments {
	return c.accountAttachments
}

// AccountGrants assign a resource directly to an account, giving the account the permission to connect to that resource.
func (c *Client) AccountGrants() *AccountGrants {
	return c.accountGrants
}

// Accounts are users that have access to strongDM. There are two types of accounts:
// 1. **Users:** humans who are authenticated through username and password or SSO.
// 2. **Service Accounts:** machines that are authenticated using a service token.
func (c *Client) Accounts() *Accounts {
	return c.accounts
}

// ControlPanel contains all administrative controls.
func (c *Client) ControlPanel() *ControlPanel {
	return c.controlPanel
}

// Nodes make up the strongDM network, and allow your users to connect securely to your resources. There are two types of nodes:
// - **Gateways** are the entry points into network. They listen for connection from the strongDM client, and provide access to databases and servers.
// - **Relays** are used to extend the strongDM network into segmented subnets. They provide access to databases and servers but do not listen for incoming connections.
func (c *Client) Nodes() *Nodes {
	return c.nodes
}

func (c *Client) Resources() *Resources {
	return c.resources
}

// RoleAttachments represent relationships between composite roles and the roles
// that make up those composite roles. When a composite role is attached to another
// role, the permissions granted to members of the composite role are augmented to
// include the permissions granted to members of the attached role.
//
// Deprecated: use multi-role instead.
func (c *Client) RoleAttachments() *RoleAttachments {
	return c.roleAttachments
}

// RoleGrants represent relationships between composite roles and the roles
// that make up those composite roles. When a composite role is attached to another
// role, the permissions granted to members of the composite role are augmented to
// include the permissions granted to members of the attached role.
//
// Deprecated: use access rules instead.
func (c *Client) RoleGrants() *RoleGrants {
	return c.roleGrants
}

// Roles are tools for controlling user access to resources. Each Role holds a
// list of resources which they grant access to. Composite roles are a special
// type of Role which have no resource associations of their own, but instead
// grant access to the combined resources associated with a set of child roles.
// Each user can be a member of one Role or composite role.
func (c *Client) Roles() *Roles {
	return c.roles
}

// SecretStores are servers where resource secrets (passwords, keys) are stored.
func (c *Client) SecretStores() *SecretStores {
	return c.secretStores
}

// Sign returns the signature for the given byte array
func (c *Client) Sign(methodName string, message []byte) string {
	// Current UTC date
	y, m, d := time.Now().UTC().Date()
	currentUTCDate := fmt.Sprintf("%04d-%02d-%02d", y, m, d)

	signingKey := hmacHelper(c.apiSecret, []byte(currentUTCDate))
	signingKey = hmacHelper(signingKey, []byte("sdm_api_v1"))

	hash := sha256.New()
	hash.Write([]byte(methodName))
	hash.Write([]byte{'\n'})
	hash.Write(message)
	hashedMessage := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(hmacHelper(signingKey, hashedMessage))
}

func hmacHelper(key, msg []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func (c *Client) wrapContext(ctx context.Context, req proto.Message, methodName string) context.Context {
	msg, _ := proto.Marshal(req)
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"x-sdm-authentication": c.apiToken,
		"x-sdm-signature":      c.Sign(methodName, msg),
		"x-sdm-api-version":    apiVersion,
		"x-sdm-user-agent":     c.userAgent,
	}))
}

func (c *Client) testOption(key string) interface{} {
	c.testOptionsMu.RLock()
	defer c.testOptionsMu.RUnlock()
	return c.testOptions[key]
}

// These defaults are taken from AWS. Customization of these values
// is a future step in the API.
const (
	defaultMaxRetries     = 3
	defaultBaseRetryDelay = 30 * time.Millisecond
	defaultMaxRetryDelay  = 300 * time.Second
)

func (c *Client) jitterSleep(iter int) {
	durMax := c.baseRetryDelay * time.Duration(2<<iter)
	if durMax > c.maxRetryDelay {
		durMax = c.maxRetryDelay
	}
	// This is a full jitter, ranging from no delay to the maximum
	// this jittering aims to prevent clients that start and conflict
	// at the same time from retrying at the same intervals
	dur := rand.Intn(int(durMax))
	// TODO: use resource.Retry instead of time.Sleep in Terraform provider
	// see https://github.com/bflad/tfproviderlint/tree/main/passes/R018
	time.Sleep(time.Duration(dur)) //lintignore:R018
}

func (c *Client) shouldRetry(iter int, err error) bool {
	if iter >= c.maxRetries-1 {
		return false
	}
	// Internal and unknown errors should be retried
	// Other error types are likely not brief temporary errors
	if s, ok := status.FromError(err); ok {
		return s.Code() == codes.Internal
	}
	return true
}

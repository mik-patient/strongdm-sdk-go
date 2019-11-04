/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	v1 "github.com/strongdm/strongdm-sdk-go/internal/v1"
)

var (
	host = "app.strongdm.com"
)

// Client is the strongDM API client.
type Client struct {
	grpcConn *grpc.ClientConn


	// Nodes are proxies in strongDM responsible to communicate with servers
	// (relays) and clients (gateways).
	Nodes *Nodes

}

// New creates a new strongDM API client.
func New(token string) (*Client, error) {
	headers := metadata.Pairs(
		"Authorization", "Bearer "+token,
	)
	cc, err := grpc.Dial(
		host,
		grpc.WithDefaultCallOptions(grpc.Header(&headers)),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot dial API server: %w", err)
	}
	client := &Client{
		grpcConn: cc,
	}
	
	client.Nodes = &Nodes{client: v1.NewNodesClient(client.grpcConn),}
	
	return client, nil
}

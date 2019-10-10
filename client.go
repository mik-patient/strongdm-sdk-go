/*
Package sdm implements an API client to strongDM restful API.
*/
package sdm

import (
	"net/http"
)


type V1CreateRelayRequest struct {

	Metadata V1RequestMetadata `json:"metadata"`

	Relay V1Relay `json:"relay"`
}

type V1CreateRelayResponse struct {

	Metadata V1ResponseMetadata `json:"metadata"`

	Relay V1Relay `json:"relay"`
}

type V1DeleteRelayResponse struct {

	Metadata V1ResponseMetadata `json:"metadata"`
}

type V1GetRelayResponse struct {

	Metadata V1ResponseMetadata `json:"metadata"`

	Relay V1Relay `json:"relay"`
}

type V1ListRelayResponse struct {

	Metadata V1ResponseMetadata `json:"metadata"`

	Relays V1Relay `json:"relays"`
}

type V1Relay struct {
	// Id is id is the unique ID for this relay.
	Id string `json:"id"`
	// Name is name is the human readable unique name for this relay.
	Name string `json:"name"`
}

type V1RequestMetadata struct {

	Page string `json:"page"`
}

type V1ResponseMetadata struct {

	Affected string `json:"affected"`

	Found string `json:"found"`

	NextPage string `json:"next_page"`
}

type V1UpdateRelayRequest struct {

	Metadata V1RequestMetadata `json:"metadata"`

	Relay V1Relay `json:"relay"`
}

type V1UpdateRelayResponse struct {

	Metadata V1ResponseMetadata `json:"metadata"`

	Relay V1Relay `json:"relay"`
}

type Relays struct{
	client *http.Client
}

func NewRelays() *Relays {
	return &Relays{
		client: http.DefaultClient,
	}
}

func (r *Relays) SetHTTPClient(client *http.Client) *Relays {
	r .client = client
	return r
}

// List returns all existing relays
func (r *Relays) List() (V1ListRelayResponse, error) {
	panic("not implemented")
}

// Create adds a new relay, and returns the new relay
func (r *Relays) Create() (V1CreateRelayResponse, error) {
	panic("not implemented")
}

// Delete removes an existing relay
func (r *Relays) Delete() (V1DeleteRelayResponse, error) {
	panic("not implemented")
}

// Get finds a sandwich by id
func (r *Relays) Get() (V1GetRelayResponse, error) {
	panic("not implemented")
}

// Update modifies an existing relay
func (r *Relays) Update() (V1UpdateRelayResponse, error) {
	panic("not implemented")
}

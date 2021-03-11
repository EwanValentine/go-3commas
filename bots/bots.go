package bots

import (
	"go-3commas/types"
	"net/http"
)

type requester interface {
	Request(endpoint, method string, request *types.Request) (*types.Response, error)
}

// Bots -
type Bots struct {
	requester
}

// NewBots -
func NewBots(requester requester) *Bots {
	return &Bots{requester}
}

// GetStrategyListRequest -
type GetStrategyListRequest struct {
	AccountID string   `json:"account_id"`
	Type      Type     `json:"type"`
	Strategy  Strategy `json:"strategy"`
}

// GetStrategyListResponse -
type GetStrategyListResponse struct {

}

// CreateRequest -
type CreateRequest Bot

// CreateResponse -
type CreateResponse struct{}

// Create a new bot
func (b *Bots) Create(request *CreateRequest) (*CreateResponse, error) {
	r, _ := types.NewRequest().Marshal(request)
	response, err := b.requester.Request("", http.MethodPost, r)
	if err != nil {
		return nil, err
	}

	var createResponse *CreateResponse
	response.Unmarshal(createResponse)

	return createResponse, nil
}

// ListRequest -
type ListRequest struct{
	// Default 50
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	AccountID string `json:"account_id"`
	// Enabled/Disabled
	Scope Scope `json:"scope"`
}

// ListResponse -
type ListResponse struct {
	Bots []Bot
}

// List bots
func (b *Bots) List() (*ListResponse, error) {
	request := &types.Request{}
	response, err := b.requester.Request("", http.MethodGet, request)
	if err != nil {
		return nil, err
	}

	var listResponse *ListResponse
	_ = response.Unmarshal(listResponse)

	return listResponse, nil
}

// UpdateRequest -
type UpdateRequest struct{}

// UpdateResponse -
type UpdateResponse struct{}

// Update -
func (b *Bots) Update() (*UpdateResponse, error) {
	request := &types.Request{}
	response, err := b.requester.Request("", http.MethodPatch, request)
	if err != nil {
		return nil, err
	}

	var updateResponse *UpdateResponse
	_ = response.Unmarshal(updateResponse)

	return updateResponse, nil
}

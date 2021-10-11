package bots

import (
	"net/http"

	"github.com/EwanValentine/go-3commas/types"
)

type requestAdapter interface {
	Request(endpoint, method string, request *types.Request, val interface{}) error
}

// Bots -
type Bots struct {
	requestAdapter
}

// NewBots -
func NewBots(requester requestAdapter) *Bots {
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
type CreateRequest struct {
	Body *Bot `json:"body"`
}

// CreateResponse -
type CreateResponse struct{}

// Create a new bot
func (b *Bots) Create(bot *Bot) (*CreateResponse, error) {
	var createBotResponse *CreateResponse

	request := &types.Request{
		Body: bot,
	}
	err := b.requestAdapter.Request("/bots", http.MethodPost, request, &createBotResponse)
	if err != nil {
		return nil, err
	}

	return createBotResponse, nil
}

// ListRequest -
type ListRequest struct {
	// Default 50
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	AccountID string `json:"account_id"`
	// Enabled/Disabled
	Scope Scope `json:"scope"`
}

// ListRequestV1 -
type ListRequestV1 struct {
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
}

// ListResponseV1 -
type ListResponseV1 struct {
	Bots []Bot `json:"bots"`
}

// ListResponse -
type ListResponse []Bot

// List bots
func (b *Bots) List() (*ListResponse, error) {

	var listBotsRequest ListRequestV1

	listBotsRequest.Limit = 50
	listBotsRequest.SortBy = "created_at"
	listBotsRequest.SortDirection = "desc"

	request := &types.Request{
		Body: listBotsRequest,
	}

	var listResponse *ListResponse
	err := b.requestAdapter.Request("/bots", http.MethodGet, request, &listResponse)
	if err != nil {
		return nil, err
	}

	return listResponse, nil
}

// UpdateRequest -
type UpdateRequest struct{}

// UpdateResponse -
type UpdateResponse struct{}

// Update -
func (b *Bots) Update() (*UpdateResponse, error) {
	request := &types.Request{}
	var updateResponse *UpdateResponse
	err := b.requestAdapter.Request("/bots", http.MethodPatch, request, &updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse, nil
}

// The code that interacts with the API for bots
package bots

import (
	"fmt"
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

// PauseRequest -
type PauseRequest struct {
	ID int `json:"id"`
}

// PauseResponse - @todo better types
type PauseResponse map[string]interface{}

// Pause -
func (b *Bots) Pause(id int) (*PauseResponse, error) {
	request := &types.Request{
		Body: PauseRequest{
			ID: id,
		},
	}
	var pauseResponse *PauseResponse
	err := b.requestAdapter.Request(fmt.Sprintf("/bots/%d/disable", id), http.MethodPost, request, &pauseResponse)
	if err != nil {
		return nil, err
	}

	return pauseResponse, nil
}

// UnpauseRequest -
type UnpauseRequest struct {
	ID int `json:"id"`
}

// UnpauseResponse -
type UnpauseResponse map[string]interface{}

// Unpause -
func (b *Bots) Unpause(id int) (*UnpauseResponse, error) {
	request := &types.Request{
		Body: UnpauseRequest{
			ID: id,
		},
	}
	var unpauseResponse *UnpauseResponse
	err := b.requestAdapter.Request(fmt.Sprintf("/bots/%d/enable", id), http.MethodPost, request, &unpauseResponse)
	if err != nil {
		return nil, err
	}

	return unpauseResponse, nil
}

type StatsRequest struct {
	ID int `json:"id"`
}

type StatsResponse map[string]interface{}

// Stats -
func (b *Bots) Stats(id int) (*StatsResponse, error) {
	request := &types.Request{
		Body: StatsRequest{ID: id},
	}
	var statsResponse *StatsResponse
	err := b.requestAdapter.Request(fmt.Sprintf("/bots/%d/deals_stats", id), http.MethodGet, request, &statsResponse)
	if err != nil {
		return nil, err
	}

	return statsResponse, nil
}

type ShowRequest struct {
	ID int `json:"id"`
}

type ShowResponse map[string]interface{}

// Show -
func (b *Bots) Show(id int) (*ShowResponse, error) {
	request := &types.Request{
		Body: ShowRequest{ID: id},
	}
	var showResponse *ShowResponse
	err := b.requestAdapter.Request(fmt.Sprintf("/bots/%d/show", id), http.MethodGet, request, &showResponse)
	if err != nil {
		return nil, err
	}

	return showResponse, nil
}

// ListRequest -
type ListRequest struct {
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
}

// ListResponse -
type ListResponse []Bot

// List bots
func (b *Bots) List() (*ListResponse, error) {

	var listBotsRequest ListRequest

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

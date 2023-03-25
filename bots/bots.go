package bots

import (
	"fmt"
	"net/http"

	"github.com/EwanValentine/go-3commas/types"
)

type RequestAdapter interface {
	Request(endpoint, method string, request *types.Request, val interface{}) error
}

type Bots struct {
	RequestAdapter
}

func NewBots(requester RequestAdapter) *Bots {
	return &Bots{requester}
}

type GetStrategyListRequest struct {
	AccountID string   `json:"account_id"`
	Type      Type     `json:"type"`
	Strategy  Strategy `json:"strategy"`
}

type GetStrategyListResponse struct {
}

type CreateRequest struct {
	Body *Bot `json:"body"`
}

type CreateResponse struct{}

func (b *Bots) Create(bot *Bot) (*CreateResponse, error) {
	createBotResponse := &CreateResponse{}

	request := &types.Request{
		Body: bot,
	}
	err := b.RequestAdapter.Request("/bots", http.MethodPost, request, createBotResponse)
	if err != nil {
		return nil, err
	}

	return createBotResponse, nil
}

type PauseRequest struct {
	ID int `json:"id"`
}

type PauseResponse map[string]interface{}

func (b *Bots) Pause(id int) (*PauseResponse, error) {
	pauseResponse := &PauseResponse{}
	request := &types.Request{
		Body: PauseRequest{
			ID: id,
		},
	}
	err := b.RequestAdapter.Request(fmt.Sprintf("/bots/%d/disable", id), http.MethodPost, request, pauseResponse)
	if err != nil {
		return nil, err
	}

	return pauseResponse, nil
}

type UnpauseRequest struct {
	ID int `json:"id"`
}

type UnpauseResponse map[string]interface{}

func (b *Bots) Unpause(id int) (*UnpauseResponse, error) {
	unpauseResponse := &UnpauseResponse{}
	request := &types.Request{
		Body: UnpauseRequest{
			ID: id,
		},
	}
	err := b.RequestAdapter.Request(fmt.Sprintf("/bots/%d/enable", id), http.MethodPost, request, unpauseResponse)
	if err != nil {
		return nil, err
	}

	return unpauseResponse, nil
}

type StatsRequest struct {
	ID int `json:"id"`
}

type StatsResponse map[string]interface{}

func (b *Bots) Stats(id int) (*StatsResponse, error) {
	statsResponse := &StatsResponse{}
	request := &types.Request{
		Body: StatsRequest{ID: id},
	}
	err := b.RequestAdapter.Request(fmt.Sprintf("/bots/%d/deals_stats", id), http.MethodGet, request, statsResponse)
	if err != nil {
		return nil, err
	}

	return statsResponse, nil
}

type ShowRequest struct {
	ID int `json:"id"`
}

type ShowResponse map[string]interface{}

func (b *Bots) Show(id int) (*ShowResponse, error) {
	showResponse := &ShowResponse{}
	request := &types.Request{
		Body: ShowRequest{ID: id},
	}
	err := b.RequestAdapter.Request(fmt.Sprintf("/bots/%d/show", id), http.MethodGet, request, showResponse)
	if err != nil {
		return nil, err
	}

	return showResponse, nil
}

type ListRequest struct {
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
}

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
	err := b.RequestAdapter.Request("/bots", http.MethodGet, request, &listResponse)
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
	err := b.RequestAdapter.Request("/bots", http.MethodPatch, request, &updateResponse)
	if err != nil {
		return nil, err
	}

	return updateResponse, nil
}

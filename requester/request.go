package requester

import (
	"bytes"
	"github.com/EwanValentine/go-3commas/types"
	"io/ioutil"
	"net/http"
)

// Requester -
type Requester struct {
	client *http.Client
}

// RequesterFactory -
func RequesterFactory() *Requester {
	client := &http.Client{}
	return NewRequester(client)
}

// NewRequester -
func NewRequester(client *http.Client) *Requester {
	return &Requester{client}
}

// Request -
func (r *Requester) Request(endpoint, method string, payload []byte) (*types.Response, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return &types.Response{
		Body: body,
	}, nil
}

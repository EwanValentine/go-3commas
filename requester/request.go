package requester

import "net/http"

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
func (r *Requester) Request(endpoint, method string) (*Response, error) {
	http.NewRequest()
}

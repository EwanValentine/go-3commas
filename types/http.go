package types

import "encoding/json"

// NewRequest -
func NewRequest() *Request {
	return &Request{
		Body: make([]byte, 0),
	}
}

// Request -
type Request struct {
	Body interface{}
}

// Marshal -
func (r *Request) Marshal(val interface{}) (*Request, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	r.Body = b

	return r, nil
}

// Response -
type Response struct {
	Body   []byte
	Status int
}

// Unmarshal -
func (r *Response) Unmarshal(val interface{}) error {
	return json.Unmarshal(r.Body, val)
}

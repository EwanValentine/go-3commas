package requester

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/EwanValentine/go-3commas/types"
)

const (
	// BaseURL -
	BaseURL = "https://api.3commas.io"
	V1      = "public/api/v1"
	V2      = "public/api/v2"
	WS      = "'wss://ws.3commas.io/websocket"
)

/*
const ENDPOINT = 'https://api.3commas.io';
const V1 = '/public/api/ver1';
const V2 = '/public/api/v2';
const WS = 'wss://ws.3commas.io/websocket';
*/

// Requester -
type Requester struct {
	client *http.Client
	key    string
}

// Factory -
func Factory() *Requester {
	client := &http.Client{}
	key := os.Getenv("API_KEY")
	return NewRequester(client, key)
}

// NewRequester -
func NewRequester(client *http.Client, key string) *Requester {
	return &Requester{client, key}
}

// Request -
func (r *Requester) Request(endpoint, method string, payload *types.Request) (*types.Response, error) {
	url := fmt.Sprintf("%s/%s/%s", BaseURL, V2, endpoint)
	log.Println("url: ", url)
	req, err := http.NewRequest(method, url, bytes.NewReader(payload.Body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("APIKEY", r.key)

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	log.Println(string(body))

	return &types.Response{
		Body: body,
	}, nil
}

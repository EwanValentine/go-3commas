package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/EwanValentine/go-3commas/signer"
	"github.com/EwanValentine/go-3commas/types"
)

const (
	BaseURL = "https://api.3commas.io"
	V1      = "/public/api/ver1"
	V2      = "/public/api/v2"
	WS      = "'wss://ws.3commas.io/websocket"
)

type Requester struct {
	client *http.Client
	key    string
	secret string
}

func FromEnv() *Requester {
	client := &http.Client{}
	key := os.Getenv("API_KEY")
	secret := os.Getenv("SECRET_KEY")

	return NewRequester(client, key, secret)
}

func NewRequester(client *http.Client, key, secret string) *Requester {
	return &Requester{client, key, secret}
}

type Payload map[string]interface{}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (r *Requester) Request(endpoint, method string, payload *types.Request, val interface{}) error {
	signerService := signer.New()

	requiredParams := Payload{
		"type":    "binance",
		"name":    "Binance",
		"api_key": r.key,
		"secret":  r.secret,
	}

	u := BaseURL + V1 + endpoint

	signature := ""

	values := url.Values{}
	if payload.Body != nil {
		v := reflect.ValueOf(payload.Body)
		typeOfS := v.Type()

		for i := 0; i < typeOfS.NumField(); i++ {
			field := typeOfS.Field(i)
			value := v.Field(i).Interface()
			values.Add(ToSnakeCase(field.Name), fmt.Sprintf("%v", value))
		}
	}

	for k, v := range requiredParams {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	queryParams := values.Encode()
	endpoint = fmt.Sprintf("%s%s", V1, endpoint)
	signature = signerService.Do(r.secret, endpoint, queryParams)
	u = fmt.Sprintf("%s?%s", u, queryParams)

	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return err
	}

	req.Header.Set("APIKEY", r.key)
	req.Header.Set("signature", signature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if method == http.MethodPost {
		req.PostForm = values
		req.Method = http.MethodPost
	}

	response, err := r.client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.Unmarshal(body, val)
}

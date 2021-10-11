package bots

import (
	"log"
	"net/http"
	"testing"

	"github.com/EwanValentine/go-3commas/conf"
	"github.com/EwanValentine/go-3commas/requester"
	"github.com/matryer/is"
)

func TestCanListBots(t *testing.T) {
	c := conf.Load()

	is := is.New(t)

	r := requester.NewRequester(http.DefaultClient, c.APIKey, c.SecretKey)
	b := NewBots(r)
	response, err := b.List()
	is.NoErr(err)

	log.Println(response)
}

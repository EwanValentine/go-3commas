package bots

import (
	"log"
	"testing"

	"github.com/EwanValentine/go-3commas/requester"
	"github.com/matryer/is"
)

func TestCanListBots(t *testing.T) {
	is := is.New(t)

	r := requester.Factory()
	b := NewBots(r)
	response, err := b.List()
	is.NoErr(err)

	log.Println(response)
}

package bots

import (
	"log"
	"net/http"
	"sync"
	"testing"

	"github.com/EwanValentine/go-3commas/conf"
	"github.com/EwanValentine/go-3commas/requester"
	"github.com/matryer/is"
)

var (
	bot  *Bot
	once *sync.Once
	c    *conf.Config
	b    *Bots
	r    *requester.Requester
)

func init() {
	once = &sync.Once{}
}

func setup(t *testing.T) *is.I {
	is := is.New(t)
	once.Do(func() {
		c = conf.Load()
		r = requester.NewRequester(http.DefaultClient, c.APIKey, c.SecretKey)
		b = NewBots(r)
	})
	return is
}

func TestCanListBots(t *testing.T) {
	is := setup(t)

	response, err := b.List()
	is.NoErr(err)
	is.True(len(*response) > 0)

	// Set a bot for future tests, this is stateful and weird,
	// need a better way to get a bot, or mock them
	res := *response
	bot = &res[0]
}

func TestCanGetStatus(t *testing.T) {
	is := setup(t)

	response, err := b.Stats(bot.ID)
	is.NoErr(err)

	log.Println(response)

	is.True(len(*response) > 0)
}

func TestCanPause(t *testing.T) {
	is := setup(t)

	response, err := b.Pause(bot.ID)
	is.NoErr(err)

	log.Println(response)
}

func TestCanUnpause(t *testing.T) {
	is := setup(t)

	response, err := b.Unpause(bot.ID)
	is.NoErr(err)

	log.Println(response)
}

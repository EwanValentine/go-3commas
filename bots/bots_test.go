package bots

import (
	"net/http"
	"testing"

	"github.com/EwanValentine/go-3commas/conf"
	"github.com/EwanValentine/go-3commas/requester"
	"github.com/matryer/is"
)

func newTestBot() *Bot {
	return &Bot{
		// Initialize the bot with the required fields
	}
}

func newTestBots() *Bots {
	c := conf.Load()
	r := requester.NewRequester(http.DefaultClient, c.APIKey, c.SecretKey)
	return NewBots(r)
}

func TestBots(t *testing.T) {
	is := is.New(t)
	b := newTestBots()

	t.Run("List", func(t *testing.T) {
		response, err := b.List()
		is.NoErr(err)
		is.True(len(*response) > 0)
	})

	bot := newTestBot()

	t.Run("Stats", func(t *testing.T) {
		response, err := b.Stats(bot.ID)
		is.NoErr(err)
		is.True(len(*response) > 0)
	})

	t.Run("Pause", func(t *testing.T) {
		_, err := b.Pause(bot.ID)
		is.NoErr(err)
	})

	t.Run("Unpause", func(t *testing.T) {
		_, err := b.Unpause(bot.ID)
		is.NoErr(err)
	})
}

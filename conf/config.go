package conf

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

// Load -
func load(c interface{}) error {
	ctx := context.Background()
	return envconfig.Process(ctx, c)
}

// Load -
func Load() *Config {
	var c Config
	if err := load(&c); err != nil {
		log.Panic(err)
	}
	return &c
}

// Config -
type Config struct {
	APIKey    string `env:"API_KEY,required"`
	SecretKey string `env:"SECRET_KEY,required"`
}

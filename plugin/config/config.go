package config

import (
	"context"

	"github.com/drone/drone-go/drone"
	pluginConfig "github.com/drone/drone-go/plugin/config"
)

type Config struct{}

func NewConfigPlugin() *Config {
	return &Config{}
}

func (c *Config) Find(ctx context.Context, req *pluginConfig.Request) (*drone.Config, error) {
	return &drone.Config{}, nil
}

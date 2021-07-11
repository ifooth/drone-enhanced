package config

import (
	"context"
	"encoding/json"

	"github.com/drone/drone-go/drone"
	pluginConfig "github.com/drone/drone-go/plugin/config"
	"github.com/sirupsen/logrus"
)

type Config struct{}

func NewConfigPlugin() *Config {
	return &Config{}
}

func (c *Config) Find(ctx context.Context, req *pluginConfig.Request) (*drone.Config, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("request body: %s", reqBody)

	if req.Build.Target == "drone-ci-enhanced" {
		return &drone.Config{Data: ""}, nil
	}
	return nil, nil
}

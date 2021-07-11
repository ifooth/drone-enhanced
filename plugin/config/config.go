package config

import (
	"context"

	"github.com/drone/drone-go/drone"
	pluginConfig "github.com/drone/drone-go/plugin/config"
	"github.com/sirupsen/logrus"
)

type Config struct{}

func NewConfigPlugin() *Config {
	return &Config{}
}

func (c *Config) Find(ctx context.Context, req *pluginConfig.Request) (*drone.Config, error) {
	logrus.Infof("req %v", req)
	if req.Repo.Branch == "drone-ci-enhanced" {
		logrus.Infof("trigger req %v", req)
		return &drone.Config{Data: ""}, nil
	}
	logrus.Infof("return null %v", req)
	return nil, nil
}

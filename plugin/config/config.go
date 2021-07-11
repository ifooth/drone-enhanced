package config

import (
	"context"
	"encoding/json"

	"github.com/drone/drone-go/drone"
	pluginConfig "github.com/drone/drone-go/plugin/config"
	"github.com/sirupsen/logrus"

	"github.com/ifooth/drone-ci-enhanced/providers"
)

type ConfigPlugin struct {
	provider providers.Provider
}

func NewConfigPlugin(provider providers.Provider) *ConfigPlugin {
	return &ConfigPlugin{provider: provider}
}

func (p *ConfigPlugin) Find(ctx context.Context, req *pluginConfig.Request) (*drone.Config, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("request body: %s", reqBody)

	content, err := p.provider.GetFileContent(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, ".drone.yml")

	config := &drone.Config{Data: content}
	return config, nil
}

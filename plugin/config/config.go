package config

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/drone/drone-go/drone"
	pluginConfig "github.com/drone/drone-go/plugin/config"
	pluginConverter "github.com/drone/drone-go/plugin/converter"
	"github.com/sirupsen/logrus"

	"github.com/ifooth/drone-ci-enhanced/plugin/converter"
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
	if err != nil {
		return nil, nil
	}

	fileListing, err := p.provider.GetFileListing(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, ".drone")
	if err != nil {
		logrus.Debugf(".drone not exist, just ignore")
	}

	converReq := &pluginConverter.Request{Repo: req.Repo, Build: req.Build, Config: drone.Config{}}
	yamlConverter := converter.NewYamlPlugin(p.provider)
	starlarkConverter := converter.NewStarlarkPlugin(p.provider)

	for _, file := range fileListing {
		if file.Type != "file" {
			continue
		}

		if yamlConverter.IsValidFilename(file.Name) {
			droneConfig, err := yamlConverter.ConvertContent(ctx, converReq, file)
			if err != nil {
				logrus.Warn("yaml convert content error, %s", err)
				continue
			}
			if content != "" {
				content = droneConfigAppend(content, droneConfig)
			}

		}

		if starlarkConverter.IsValidFilename(file.Name) {
			droneConfig, err := starlarkConverter.ConvertContent(ctx, converReq, file)
			if err != nil {
				logrus.Warn("yaml convert content error, %s", err)
				continue
			}
			if content != "" {
				content = droneConfigAppend(content, droneConfig)
			}

		}
	}

	config := &drone.Config{Data: content}
	logrus.Debugf("render content: %s", config.Data)

	return config, nil
}

func droneConfigAppend(droneConfig string, appends ...string) string {
	for _, a := range appends {
		a = strings.Trim(a, " \n")
		if a != "" {
			if !strings.HasPrefix(a, "---\n") {
				a = "---\n" + a
			}
			droneConfig += a
			if !strings.HasSuffix(droneConfig, "\n") {
				droneConfig += "\n"
			}
		}
	}
	return droneConfig
}

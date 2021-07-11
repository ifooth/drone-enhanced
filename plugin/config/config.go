package config

import (
	"context"
	"encoding/json"
	"strings"

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
	if err != nil {
		return nil, nil
	}

	fileListing, err := p.provider.GetFileListing(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, ".drone")
	if err != nil {
		logrus.Debugf(".drone not exist, just ignore")
	}

	for _, file := range fileListing {
		if f, _ := p.FindYaml(ctx, req, file); f != "" {
			content = droneConfigAppend(content, f)
		}

	}

	config := &drone.Config{Data: content}
	logrus.Debugf("render content: %s", config.Data)

	return config, nil
}

func (p *ConfigPlugin) FindYaml(ctx context.Context, req *pluginConfig.Request, fileEntry providers.FileListingEntry) (fileContent string, err error) {
	if fileEntry.Type != "file" {
		return "", nil
	}

	switch {
	case strings.HasSuffix(fileEntry.Name, ".yaml"):
	case strings.HasSuffix(fileEntry.Name, ".yml"):
	default:
		return "", nil
	}

	content, err := p.provider.GetFileContent(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, fileEntry.Path)

	return content, nil
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

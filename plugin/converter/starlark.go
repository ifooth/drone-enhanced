package converter

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	pluginConverter "github.com/drone/drone-go/plugin/converter"
	"github.com/ifooth/drone-ci-enhanced/plugin/converter/starlark"
	"github.com/ifooth/drone-ci-enhanced/providers"
)

type StarlarkPlugin struct {
	provider providers.Provider
}

func NewStarlarkPlugin(provider providers.Provider) *StarlarkPlugin {
	return &StarlarkPlugin{provider: provider}
}

func (p *StarlarkPlugin) Convert(ctx context.Context, req *pluginConverter.Request) (*drone.Config, error) {
	return nil, nil
}

func (p *StarlarkPlugin) ConvertContent(ctx context.Context, req *pluginConverter.Request, fileEntry providers.FileListingEntry) (string, error) {
	if !p.IsValidFilename(fileEntry.Name) {
		return "", nil
	}

	content, err := p.provider.GetFileContent(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, fileEntry.Path)
	if err != nil {
		return "", err
	}

	droneConfig, err := starlark.Parse(req, fileEntry.Name, content)
	if err != nil {
		return "", err
	}
	return droneConfig, nil
}

func (p *StarlarkPlugin) IsValidFilename(name string) bool {
	switch {
	case strings.HasSuffix(name, ".star"):
		return true
	case strings.HasSuffix(name, ".starlark"):
		return true
	case strings.HasSuffix(name, ".bzl"):
		return true
	default:
		return false
	}
}

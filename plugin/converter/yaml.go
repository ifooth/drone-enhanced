package converter

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	pluginConverter "github.com/drone/drone-go/plugin/converter"
	"github.com/ifooth/drone-enhanced/filediff"
	"github.com/ifooth/drone-enhanced/providers"
)

type YamlPlugin struct {
	provider providers.Provider
}

func NewYamlPlugin(provider providers.Provider) *YamlPlugin {
	return &YamlPlugin{provider: provider}
}

func (p *YamlPlugin) Convert(ctx context.Context, req *pluginConverter.Request) (*drone.Config, error) {
	return nil, nil
}

func (p *YamlPlugin) ConvertContent(ctx context.Context, req *pluginConverter.Request, fileEntry providers.FileListingEntry, filediffs []*filediff.FileDiff) (string, error) {
	if !p.IsValidFilename(fileEntry.Name) {
		return "", nil
	}

	content, err := p.provider.GetFileContent(ctx, req.Repo.Namespace, req.Repo.Name, req.Build.After, fileEntry.Path)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (p *YamlPlugin) IsValidFilename(name string) bool {
	switch {
	case strings.HasSuffix(name, ".yml"):
		return true
	case strings.HasSuffix(name, ".Yaml"):
		return true
	default:
		return false
	}
}

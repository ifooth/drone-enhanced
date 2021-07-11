package converter

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	pluginConverter "github.com/drone/drone-go/plugin/converter"
	"github.com/ifooth/drone-ci-enhanced/plugin/converter/starlark"
)

type StarlarkPlugin struct{}

func (p *StarlarkPlugin) Convert(ctx context.Context, req *pluginConverter.Request) (*drone.Config, error) {
	return nil, nil
}

func (p *StarlarkPlugin) ConvertContent(ctx context.Context, req *pluginConverter.Request, fileName string, fileContent string) (string, error) {

	// if the file extension is not jsonnet we can
	// skip this plugin by returning zero values.
	switch {
	case strings.HasSuffix(fileName, ".script"):
	case strings.HasSuffix(fileName, ".star"):
	case strings.HasSuffix(fileName, ".starlark"):
	default:
		return "", nil
	}

	content, err := starlark.Parse(req, fileName, fileContent)
	if err != nil {
		return "", err
	}
	return content, nil
}

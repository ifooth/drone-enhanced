package starlark

import (
	"github.com/drone/drone-go/drone"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"

	"github.com/ifooth/drone-enhanced/filediff"
)

func createArgs(repo *drone.Repo, build *drone.Build, input map[string]interface{}, filediffs []*filediff.FileDiff) []starlark.Value {
	return []starlark.Value{
		starlarkstruct.FromStringDict(
			starlark.String("context"),
			starlark.StringDict{
				"repo":      starlarkstruct.FromStringDict(starlark.String("repo"), fromRepo(repo)),
				"build":     starlarkstruct.FromStringDict(starlark.String("build"), fromBuild(build)),
				"input":     starlarkstruct.FromStringDict(starlark.String("input"), fromInput(input)),
				"filediffs": fromFileDiffList(filediffs),
			},
		),
	}
}

func fromInput(input map[string]interface{}) starlark.StringDict {
	out := map[string]starlark.Value{}
	for k, v := range input {
		if s, ok := v.(string); ok {
			out[k] = starlark.String(s)
		}
	}
	return out
}

func fromFileDiffList(filediffs []*filediff.FileDiff) *starlark.List {
	list := new(starlark.List)
	for _, v := range filediffs {
		list.Append(fromFileDiffDict(v))
	}
	return list
}

func fromFileDiffDict(v *filediff.FileDiff) *starlark.Dict {
	dict := new(starlark.Dict)
	dict.SetKey(starlark.String("name"), starlark.String(v.Name))
	dict.SetKey(starlark.String("path"), starlark.String(v.Path))
	dict.SetKey(starlark.String("type"), starlark.String(v.Type))
	dict.SetKey(starlark.String("extensions"), fromMap(v.Extensions))
	return dict
}

func fromBuild(v *drone.Build) starlark.StringDict {
	return starlark.StringDict{
		"event":         starlark.String(v.Event),
		"action":        starlark.String(v.Action),
		"cron":          starlark.String(v.Cron),
		"environment":   starlark.String(v.Deploy),
		"link":          starlark.String(v.Link),
		"branch":        starlark.String(v.Target),
		"source":        starlark.String(v.Source),
		"before":        starlark.String(v.Before),
		"after":         starlark.String(v.After),
		"target":        starlark.String(v.Target),
		"ref":           starlark.String(v.Ref),
		"commit":        starlark.String(v.After),
		"title":         starlark.String(v.Title),
		"message":       starlark.String(v.Message),
		"source_repo":   starlark.String(v.Fork),
		"author_login":  starlark.String(v.Author),
		"author_name":   starlark.String(v.AuthorName),
		"author_email":  starlark.String(v.AuthorEmail),
		"author_avatar": starlark.String(v.AuthorAvatar),
		"sender":        starlark.String(v.Sender),
		"debug":         starlark.Bool(v.Debug),
		"params":        fromMap(v.Params),
	}
}

func fromRepo(v *drone.Repo) starlark.StringDict {
	return starlark.StringDict{
		"uid":                  starlark.String(v.UID),
		"name":                 starlark.String(v.Name),
		"namespace":            starlark.String(v.Namespace),
		"slug":                 starlark.String(v.Slug),
		"git_http_url":         starlark.String(v.HTTPURL),
		"git_ssh_url":          starlark.String(v.SSHURL),
		"link":                 starlark.String(v.Link),
		"branch":               starlark.String(v.Branch),
		"config":               starlark.String(v.Config),
		"private":              starlark.Bool(v.Private),
		"visibility":           starlark.String(v.Visibility),
		"active":               starlark.Bool(v.Active),
		"trusted":              starlark.Bool(v.Trusted),
		"protected":            starlark.Bool(v.Protected),
		"ignore_forks":         starlark.Bool(v.IgnoreForks),
		"ignore_pull_requests": starlark.Bool(v.IgnorePulls),
	}
}

func fromMap(m map[string]string) *starlark.Dict {
	dict := new(starlark.Dict)
	for k, v := range m {
		dict.SetKey(
			starlark.String(k),
			starlark.String(v),
		)
	}
	return dict
}

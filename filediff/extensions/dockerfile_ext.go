package extensions

import (
	"regexp"
	"strings"
)

var argVersion = regexp.MustCompile(`ARG VERSION=(?P<version>[\w\-\.]+)`)

func ParseDockerfileVersion(content string) string {
	match := argVersion.FindStringSubmatch(content)
	if len(match) == 0 {
		return ""
	}

	result := make(map[string]string)
	for i, name := range argVersion.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	version := result["version"]

	if strings.HasPrefix(version, "v") {
		return version
	}

	return "v" + version
}

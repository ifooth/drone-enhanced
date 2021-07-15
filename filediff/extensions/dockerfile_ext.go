package extensions

import (
	"regexp"
	"strings"
)

var argVersion = regexp.MustCompile(`(?i)[ARG|ENV|LABEL]{1} VERSION=(?P<version>[\w\-\.]+)`)

func ParseDockerfileVersion(content string) string {
	lines := strings.Split(content, "\n")

	// 拿最后一个版本号
	version := ""

	for _, line := range lines {
		_version := ParseDockerfileVersionLine(line)
		if _version != "" {
			version = _version
		}
	}
	return version
}

func ParseDockerfileVersionLine(line string) string {
	match := argVersion.FindStringSubmatch(line)
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

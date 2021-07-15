package extensions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerfileExtensions(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"arg":       {input: "From ab\nARG VERSION=1.1 \n", want: "v1.1"},
		"multi_arg": {input: "From ab\nARG VERSION=1.1 \n ARG VERSION=1.0", want: "v1.0"},
		//"env":   {input: "From ab\nENV VERSION=1.1 \n", want: "v1.1"},
		//"label": {input: "From ab\nlabel VERSION=1.1 \n", want: "v1.1"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			version := ParseDockerfileVersion(tc.input)
			assert.NotEmpty(t, version)
			assert.Equal(t, tc.want, version)
		})
	}

}

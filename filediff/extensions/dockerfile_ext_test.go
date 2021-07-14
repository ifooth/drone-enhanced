package extensions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerfileExtensions(t *testing.T) {
	version := ParseDockerfileVersion("From ARG VERSION=1.1")

	assert.NotEmpty(t, version)
	assert.Equal(t, "v1.1", version)
}

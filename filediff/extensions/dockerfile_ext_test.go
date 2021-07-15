package extensions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerfileExtensions(t *testing.T) {
	version := ParseDockerfileVersion("From ab\nARG VERSION=1.1 \n")

	assert.NotEmpty(t, version)
	assert.Equal(t, "v1.1", version)
}

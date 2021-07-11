package providers

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiteaGetFileListing(t *testing.T) {
	ctx := context.Background()
	cred := &GiteaCredential{
		URL:   os.Getenv("TEST_URL"),
		Token: os.Getenv("TEST_TOKEN"),
		Debug: true,
	}
	client, err := NewGiteaClient(cred)
	assert.NoError(t, err)
	entrys, err := client.GetFileListing(ctx, os.Getenv("TEST_NAMESPACE"), os.Getenv("TEST_NAME"), os.Getenv("TEST_COMMITREF"), "")
	assert.NoError(t, err)
	assert.NotEmpty(t, entrys)
	t.Log(entrys)
}

func TestGiteaGetFileContent(t *testing.T) {
	ctx := context.Background()
	cred := &GiteaCredential{
		URL:   os.Getenv("TEST_URL"),
		Token: os.Getenv("TEST_TOKEN"),
		Debug: true,
	}
	client, err := NewGiteaClient(cred)
	assert.NoError(t, err)
	content, err := client.GetFileContent(ctx, os.Getenv("TEST_NAMESPACE"), os.Getenv("TEST_NAME"), os.Getenv("TEST_COMMITREF"), ".drone.yml")
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
	t.Log(content)
}

func TestGiteaChangedFilesInDiff(t *testing.T) {
	ctx := context.Background()
	cred := &GiteaCredential{
		URL:   os.Getenv("TEST_URL"),
		Token: os.Getenv("TEST_TOKEN"),
		Debug: true,
	}
	client, err := NewGiteaClient(cred)
	assert.NoError(t, err)
	content, err := client.ChangedFilesInDiff(ctx, os.Getenv("TEST_NAMESPACE"), os.Getenv("TEST_NAME"), "", os.Getenv("TEST_COMMITREF"))
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
	t.Logf("%v", content)
}

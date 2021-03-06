package providers

import (
	"context"
	"fmt"
	"path"

	"code.gitea.io/sdk/gitea"
	"github.com/ifooth/drone-enhanced/filediff"
)

type GiteaCredential struct {
	Server string
	Token  string
	Debug  bool
}

type GiteaClient struct {
	delegate *gitea.Client
}

func NewGiteaClient(cred *GiteaCredential) (*GiteaClient, error) {
	client, err := gitea.NewClient(cred.Server, gitea.SetToken(cred.Token))
	if err != nil {
		return nil, err
	}
	if cred.Debug {
		gitea.SetDebugMode()(client)
	}

	giteaClient := &GiteaClient{
		delegate: client,
	}
	return giteaClient, nil
}

func (c *GiteaClient) GetFileListing(ctx context.Context, namespace string, name string, commitRef string, path string) (fileListing []FileListingEntry, err error) {
	c.delegate.SetContext(ctx)

	contents, _, err := c.delegate.ListContents(namespace, name, commitRef, path)
	if err != nil {
		return nil, err
	}

	fileListing = make([]FileListingEntry, 0, len(contents))
	for _, content := range contents {
		entry := FileListingEntry{Type: content.Type, Name: content.Name, Path: content.Path}
		fileListing = append(fileListing, entry)
	}
	return fileListing, nil
}

func (c *GiteaClient) GetFileContent(ctx context.Context, namespace string, name string, commitRef string, path string) (fileContent string, err error) {
	c.delegate.SetContext(ctx)

	data, _, err := c.delegate.GetFile(namespace, name, commitRef, path)
	return fmt.Sprintf("%s", data), err
}

func (c *GiteaClient) ChangedFilesInDiff(ctx context.Context, namespace string, name string, base string, head string) ([]*filediff.FileDiff, error) {
	c.delegate.SetContext(ctx)

	commit, _, err := c.delegate.GetSingleCommit(namespace, name, head)
	if err != nil {
		return nil, err
	}

	diffs := make([]*filediff.FileDiff, 0, len(commit.Files))
	for _, v := range commit.Files {
		d := &filediff.FileDiff{
			Name:       path.Base(v.Filename),
			Path:       v.Filename,
			Type:       "file",
			Extensions: map[string]string{},
		}
		diffs = append(diffs, d)
	}
	return diffs, nil
}

package providers

import (
	"context"

	"github.com/ifooth/drone-ci-enhanced/filediff"
)

type FileListingEntry struct {
	Type string
	Name string
	Path string
}

type Provider interface {
	GetFileContent(ctx context.Context, namespace string, name string, commitRef string, path string) (fileContent string, err error)
	GetFileListing(ctx context.Context, namespace string, name string, commitRef string, path string) (fileListing []FileListingEntry, err error)
	ChangedFilesInDiff(ctx context.Context, namespace string, name string, base string, head string) ([]*filediff.FileDiff, error)
}

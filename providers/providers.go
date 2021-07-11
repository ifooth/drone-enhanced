package providers

import "context"

type FileListingEntry struct {
	Type string
	Name string
	Path string
}

type Provider interface {
	GetFileContent(ctx context.Context, namespace string, name string, commitRef string, path string) (fileContent string, err error)
	GetFileListing(ctx context.Context, namespace string, name string, commitRef string, path string) (fileListing []FileListingEntry, err error)
}

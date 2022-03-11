package content

import "github.com/kubeshop/testkube/pkg/api/v1/testkube"

// ContentFetcher is interface container for all possible fetchers
type ContentFetcher interface {
	StringFetcher
	URIFetcher
	GitDirFetcher
	GitFileFetcher

	Fetch(content *testkube.TestContent) (path string, err error)
}

// StringFetcher interface for fetching string based content to file
type StringFetcher interface {
	FetchString(str string) (path string, err error)
}

// URIFetcher interface for fetching URI based content to file
type URIFetcher interface {
	FetchURI(uri string) (path string, err error)
}

// GitDirFetcher interface for fetching GitDir based content to local directory
type GitDirFetcher interface {
	FetchGitDir(repo *testkube.Repository) (path string, err error)
}

// GitFileFetcher interface for fetching GitDir based content to local file
type GitFileFetcher interface {
	FetchGitFile(repo *testkube.Repository) (path string, err error)
}

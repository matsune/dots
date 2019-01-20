package dots

import (
	"io"
	"net/http"
	"net/url"
	"path"
)

type GithubResolver struct {
	repo   string
	host   string
	branch string
}

func NewGithubResolver(repo string) *GithubResolver {
	return &GithubResolver{
		repo:   repo,
		host:   "https://raw.githubusercontent.com/",
		branch: "master",
	}
}

func (r *GithubResolver) ReadFile(sub, file string) (io.ReadCloser, error) {
	url := r.url(sub, file)
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (r *GithubResolver) url(sub, file string) *url.URL {
	u, err := url.Parse(r.host)
	if err != nil {
		return nil
	}
	u.Path = path.Join(u.Path, r.repo, r.branch, sub, file)
	return u
}

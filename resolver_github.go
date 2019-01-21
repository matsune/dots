package dots

import (
	"fmt"
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
	url, err := r.url(sub, file)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url.String())
	if resp.StatusCode == 200 {
		return resp.Body, nil
	} else if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf(resp.Status)
}

func (r *GithubResolver) url(sub, file string) (*url.URL, error) {
	u, err := url.Parse(r.host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, r.repo, r.branch, sub, file)
	return u, nil
}

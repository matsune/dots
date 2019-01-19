package dots

import "io"

type GithubResolver struct {
	repo string
}

func NewGithubResolver(repo string) *GithubResolver {
	return &GithubResolver{
		repo: repo,
	}
}

// - TODO: implement resolver

func (r *GithubResolver) Targets() ([]Target, error) {
	return nil, nil
}

func (r *GithubResolver) ReadYml(sub string) ([]byte, error) {
	return nil, nil
}

func (r *GithubResolver) ReadFile(t Target) (io.ReadCloser, error) {
	return nil, nil
}

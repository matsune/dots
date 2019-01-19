package dots

import "io"

type Resolver interface {
	// Get all targets under repository
	Targets() ([]Target, error)
	// Read sub-directory's dots.yml
	ReadYml(sub string) ([]byte, error)
	// Read contents of target file
	ReadFile(t Target) (io.ReadCloser, error)
}

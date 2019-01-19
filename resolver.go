package dots

import "io"

type resolver interface {
	// Get all targets under repository
	Targets() ([]target, error)
	// Read sub-directory's dots.yml
	readYml(sub string) ([]byte, error)
	// Read contents of target file
	readFile(target) (io.ReadCloser, error)
}

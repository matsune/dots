package dots

import "io"

type Resolver interface {
	ReadFile(sub, file string) (io.ReadCloser, error)
}

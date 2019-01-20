package dots

import (
	"io"
	"os"
	"path/filepath"
)

type localResolver struct {
	repo string
}

func (r *localResolver) ReadFile(sub, file string) (io.ReadCloser, error) {
	filePath := r.filePath(sub, file)
	return os.Open(filePath)
}

func (r *localResolver) filePath(sub, file string) string {
	return filepath.Join(r.repo, sub, file)
}

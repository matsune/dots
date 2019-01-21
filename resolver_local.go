package dots

import (
	"io"
	"os"
	"path/filepath"
)

type LocalResolver struct {
	Repo string
}

func (r *LocalResolver) ReadFile(sub, file string) (io.ReadCloser, error) {
	filePath := r.filePath(sub, file)
	return os.Open(filePath)
}

func (r *LocalResolver) filePath(sub, file string) string {
	return filepath.Join(r.Repo, sub, file)
}

package dots

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type resolver interface {
	ReadYaml() ([]target, error)
	do(target) error
}

type localResolver struct {
	repo string
}

func (r *localResolver) ReadYaml() ([]target, error) {
	p := r.repo + "/dots.yml"
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("Could not read %s\n", p)
	}
	t, err := parseYaml(data)
	if err != nil {
		return nil, fmt.Errorf("Could not parse dots.yaml\n")
	}
	return t, nil
}

func (r *localResolver) do(t target) error {
	filePath := filepath.Join(r.repo, t.File)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	dstPath, err := homedir.Expand(t.Dst)
	if err != nil {
		return err
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}
	return nil
}

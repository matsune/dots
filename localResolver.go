package dots

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

type localResolver struct {
	repo string
}

func (r *localResolver) ReadTargets() ([]target, error) {
	return r.readDotsYml("")
}

// Get dots.yml path
func (r *localResolver) ymlPath(sub string) string {
	return filepath.Join(r.repo, sub, "dots.yml")
}

func (r *localResolver) readDotsYml(sub string) ([]target, error) {
	p := r.ymlPath(sub)
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("Could not read %s", p)
	}
	yml, err := parseYaml(data)
	if err != nil {
		return nil, err
	}
	for i := range yml.Targets {
		yml.Targets[i].Sub = sub
	}
	res := yml.Targets
	for _, s := range yml.Sub {
		// recursively read dots.yml of sub directories
		ts, err := r.readDotsYml(filepath.Join(sub, s))
		if err != nil {
			return nil, err
		}
		for _, t := range ts {
			res = append(res, t)
		}
	}
	return res, nil
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

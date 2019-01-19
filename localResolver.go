package dots

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type localResolver struct {
	repo string
}

func (r *localResolver) Targets() ([]target, error) {
	return r.targets("")
}

func (r *localResolver) filePath(sub, file string) string {
	return filepath.Join(r.repo, sub, file)
}

func (r *localResolver) ymlPath(sub string) string {
	return r.filePath(sub, "dots.yml")
}

func (r *localResolver) readYml(sub string) ([]byte, error) {
	p := r.ymlPath(sub)
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *localResolver) targets(sub string) ([]target, error) {
	data, err := r.readYml(sub)
	if err != nil {
		return nil, err
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
		ts, err := r.targets(filepath.Join(sub, s))
		if err != nil {
			return nil, err
		}
		for _, t := range ts {
			res = append(res, t)
		}
	}
	return res, nil
}

func (r *localResolver) readFile(t target) (io.ReadCloser, error) {
	filePath := r.filePath(t.Sub, t.File)
	return os.Open(filePath)
}

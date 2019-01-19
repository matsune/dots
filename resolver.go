package dots

import (
	"fmt"
	"io/ioutil"
)

type resolver interface {
	ReadYaml() ([]target, error)
}

type localResolver struct {
	repo string
}

func (l *localResolver) ReadYaml() ([]target, error) {
	p := l.repo + "/dots.yml"
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

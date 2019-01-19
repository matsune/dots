package dots

import (
	"golang.org/x/sync/errgroup"
	yaml "gopkg.in/yaml.v2"
)

func ParseYaml(str string) ([]target, error) {
	var y struct {
		Targets []target `yaml:"targets"`
	}
	err := yaml.Unmarshal([]byte(str), &y)
	if err != nil {
		return nil, err
	}

	eg := errgroup.Group{}
	for _, t := range y.Targets {
		eg.Go(func() error {
			return t.validate()
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	return y.Targets, nil
}

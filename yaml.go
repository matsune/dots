package dots

import (
	"golang.org/x/sync/errgroup"
	yaml "gopkg.in/yaml.v2"
)

type YamlFile struct {
	Targets []Target `yaml:"targets,omitempty"`
	Sub     []string `yaml:"sub,omitempty"`
}

func ParseYaml(str []byte) (YamlFile, error) {
	var y YamlFile
	err := yaml.Unmarshal(str, &y)
	if err != nil {
		return YamlFile{}, err
	}

	eg := errgroup.Group{}
	for _, t := range y.Targets {
		eg.Go(func() error {
			return t.validate()
		})
	}
	err = eg.Wait()
	if err != nil {
		return YamlFile{}, err
	}
	return y, nil
}

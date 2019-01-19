package dots

import (
	"golang.org/x/sync/errgroup"
	yaml "gopkg.in/yaml.v2"
)

type yamlFile struct {
	Targets []target `yaml:"targets,omitempty"`
	Sub     []string `yaml:"sub,omitempty"`
}

func parseYaml(str []byte) (yamlFile, error) {
	var y yamlFile
	err := yaml.Unmarshal(str, &y)
	if err != nil {
		return yamlFile{}, err
	}

	eg := errgroup.Group{}
	for _, t := range y.Targets {
		eg.Go(func() error {
			return t.validate()
		})
	}
	err = eg.Wait()
	if err != nil {
		return yamlFile{}, err
	}
	return y, nil
}

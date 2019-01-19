package dots

import "fmt"

type target struct {
	Name string `yaml:"name"`
	File string `yaml:"file"`
	Dst  string `yaml:"dst"`
}

func (t target) validate() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("target's name is empty")
	}
	if len(t.File) == 0 {
		return fmt.Errorf("target's file is empty")
	}
	if len(t.Dst) == 0 {
		return fmt.Errorf("target's dst is empty")
	}
	return nil
}

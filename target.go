package dots

import "fmt"

type target struct {
	Name string `yaml:"name"`
	Src  string `yaml:"src"`
	Dst  string `yaml:"dst"`
}

func (t target) validate() error {
	if len(t.Name) == 0 {
		return fmt.Errorf("target's name is empty")
	}
	if len(t.Src) == 0 {
		return fmt.Errorf("target's src is empty")
	}
	if len(t.Dst) == 0 {
		return fmt.Errorf("target's dst is empty")
	}
	return nil
}

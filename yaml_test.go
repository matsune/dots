package dots

import (
	"reflect"
	"testing"
)

func TestParseYamlSuccess(t *testing.T) {

	type test struct {
		str    string
		expect []target
	}

	successTests := []test{
		test{
			str: `
targets: 
  - dst: ~/.vimrc
    name: vimrc
    src: .vimrc
`,
			expect: []target{
				target{
					Name: "vimrc",
					Src:  ".vimrc",
					Dst:  "~/.vimrc",
				},
			},
		},
	}

	for _, c := range successTests {
		res, err := ParseYaml(c.str)
		if err != nil {
			t.Error(err)
		} else {
			if !reflect.DeepEqual(c.expect, res) {
				t.Errorf("expected %v, but got %v", c.expect, res)
			}
		}
	}
}

func TestParseYamlFail(t *testing.T) {

	failTests := map[string]string{
		"no name": `
targets: 
  - dst: ~/.vimrc
    src: .vimrc
`,
		"no dst": `
targets: 
  - name: vimrc
    src: .vimrc
`,
		"no src": `
targets: 
  - dst: ~/.vimrc
    name: vimrc
`,
	}

	for k, c := range failTests {
		res, err := ParseYaml(c)
		if err == nil {
			t.Errorf("[test %s] expected yaml should return error, but got %v", k, res)
		}
	}
}

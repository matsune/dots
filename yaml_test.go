package dots

import (
	"reflect"
	"testing"
)

func Test_parseYaml(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    yamlFile
		wantErr bool
	}{
		{
			name: "parse targets",
			str: `
targets: 
  - dst: ~/.vimrc
    name: vimrc
    file: .vimrc
`,
			want: yamlFile{
				Targets: []target{
					target{
						Name: "vimrc",
						File: ".vimrc",
						Dst:  "~/.vimrc",
					},
				},
			},
		},
		{
			name: "parse sub",
			str: `
targets:
sub:
  - a
  - ./b
`,
			want: yamlFile{
				Sub: []string{"a", "./b"},
			},
		},
		// - fail tests
		{
			name: "no name",
			str: `
targets:
  - dst: ~/.vimrc
    file: .vimrc
`,
			wantErr: true,
		},
		{
			name: "no dst",
			str: `
targets:
  - name: vimrc
    file: .vimrc
`,
			wantErr: true,
		},
		{
			name: "no src",
			str: `
targets:
  - name: vimrc
    dst: .vimrc
`,
			wantErr: true,
		},
		{
			name: "duplicated key",
			str: `
targets:
  - dst: ~/.vimrc
    dst: ~/.vimrc
`,
			wantErr: true,
		},
		{
			name: "invalid format",
			str: `
- targets
  - dst: ~/.vimrc
`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseYaml([]byte(tt.str))
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

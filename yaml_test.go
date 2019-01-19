package dots

import (
	"reflect"
	"testing"
)

func Test_ParseYaml(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    YamlFile
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
			want: YamlFile{
				Targets: []Target{
					Target{
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
			want: YamlFile{
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
			got, err := ParseYaml([]byte(tt.str))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

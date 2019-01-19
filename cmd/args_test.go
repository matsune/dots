package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    cmdArgs
		wantErr bool
	}{
		{
			name: "no target",
			args: []string{"dots", "matsune/dotfiles"},
			want: cmdArgs{
				cmd:     "dots",
				repo:    "matsune/dotfiles",
				targets: nil,
			},
			wantErr: false,
		}, {
			name: "1 target",
			args: []string{"dots", "matsune/dotfiles", "vim"},
			want: cmdArgs{
				cmd:     "dots",
				repo:    "matsune/dotfiles",
				targets: []string{"vim"},
			},
			wantErr: false,
		},
		{
			name: "2 target",
			args: []string{"dots", "matsune/dotfiles", "vim", "zsh", "tmux"},
			want: cmdArgs{
				cmd:     "dots",
				repo:    "matsune/dotfiles",
				targets: []string{"vim", "zsh", "tmux"},
			},
			wantErr: false,
		},
		// fail tests
		{
			name:    "no Repo",
			args:    []string{"dots"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

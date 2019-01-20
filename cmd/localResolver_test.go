package main

import (
	"reflect"
	"testing"

	"github.com/matsune/dots"
)

func Test_localResolver_ymlPath(t *testing.T) {
	tests := []struct {
		name string
		repo string
		sub  string
		want string
	}{
		{
			name: "root dots.yml",
			repo: "test",
			sub:  "",
			want: "test/dots.yml",
		},
		{
			name: "sub dots.yml",
			repo: "test",
			sub:  "aa",
			want: "test/aa/dots.yml",
		},
		{
			name: "sub dots.yml",
			repo: "test",
			sub:  "aa/bb",
			want: "test/aa/bb/dots.yml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &localResolver{
				repo: tt.repo,
			}
			if got := r.ymlPath(tt.sub); got != tt.want {
				t.Errorf("localResolver.ymlPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localResolver_readDotsYml(t *testing.T) {
	r := localResolver{
		repo: "../test_dotfiles",
	}
	got, err := r.Targets()
	if err != nil {
		t.Errorf("localResolver.Targets() error = %v", err)
		return
	}

	want := []dots.Target{
		{
			Name: "vimrc",
			File: "./.vimrc",
			Dst:  "~/go/src/github.com/matsune/dots/.vimrc",
			Sub:  "",
		},
		{
			Name: "zshrc",
			File: ".zshrc",
			Dst:  "~/go/src/github.com/matsune/dots/.zshrc",
			Sub:  "zsh",
		},
		{
			Name: "zprofile",
			File: ".zprofile",
			Dst:  "~/go/src/github.com/matsune/dots/.zprofile",
			Sub:  "zsh/zprofile",
		},
	}
	t.Run("test_dotfiles", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Targets() = %v, want %v", got, want)
		}
	})
}
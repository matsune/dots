package dots

import (
	"reflect"
	"testing"
)

func TestParseSuccess(t *testing.T) {

	type test struct {
		args   []string
		expect cmd
	}

	successTests := map[string]test{
		"no target": test{
			args: []string{"dots", "matsune/dotfiles"},
			expect: cmd{
				Self:    "dots",
				Repo:    "matsune/dotfiles",
				Targets: nil,
			},
		},
		"1 target": test{
			args: []string{"dots", "matsune/dotfiles", "vim"},
			expect: cmd{
				Self:    "dots",
				Repo:    "matsune/dotfiles",
				Targets: []string{"vim"},
			},
		},
		"2 target": test{
			args: []string{"dots", "matsune/dotfiles", "vim", "zsh", "tmux"},
			expect: cmd{
				Self:    "dots",
				Repo:    "matsune/dotfiles",
				Targets: []string{"vim", "zsh", "tmux"},
			},
		},
	}

	for name, c := range successTests {
		t.Logf("test [%s]", name)
		if res, err := Parse(c.args); err != nil {
			t.Fatal(err)
		} else {
			if res.Self != c.expect.Self {
				t.Errorf("expected Self is %s, but got %s", c.expect.Self, res.Self)
			}
			if res.Repo != c.expect.Repo {
				t.Errorf("expected Repo is %s, but got %s", c.expect.Repo, res.Repo)
			}
			if !reflect.DeepEqual(res.Targets, c.expect.Targets) {
				t.Errorf("expected Targets is %s, but got %s", c.expect.Targets, res.Targets)
			}
		}
	}
}

func TestParseError(t *testing.T) {

	type test struct {
		args []string
	}

	failTests := map[string]test{
		"no Repo": test{
			args: []string{"dots"},
		},
	}

	for name, c := range failTests {
		t.Logf("test [%s]", name)
		if _, err := Parse(c.args); err == nil {
			t.Fatal("Parse should return error")
		}
	}
}

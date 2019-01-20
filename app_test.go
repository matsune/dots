package dots

import (
	"io"
	"os"
	"testing"
)

type testResolver struct{}

func (testResolver) ReadFile(sub, file string) (io.ReadCloser, error) {
	return nil, nil
}

func TestSetResolver(t *testing.T) {

	tests := []struct {
		name string
		res  Resolver
	}{
		{
			name: "",
			res:  testResolver{},
		},
	}
	for _, tt := range tests {
		r = nil
		t.Run(tt.name, func(t *testing.T) {
			SetResolver(tt.res)
			if r != tt.res {
				t.Errorf("SetResolver r = %v, want %v", r, tt.res)
			}
		})
	}
}

func Test_doTarget(t *testing.T) {
	SetResolver(&localResolver{
		repo: "test_dotfiles",
	})

	tests := []struct {
		name    string
		t       Target
		after   func()
		wantErr bool
	}{
		{
			name: "success copy test_dotfiles/.vimrc to ./.vimrc",
			t: Target{
				Name: "vimrc",
				File: "./.vimrc",
				Dst:  "./.vimrc",
			},
			after: func() {
				os.Remove(".vimrc")
			},
		},
		// fail tests
		{
			name: "no such file",
			t: Target{
				Name: "a",
				File: "./.a",
				Dst:  "./.a",
			},
			wantErr: true,
		},
		{
			name: "can't expand dst path",
			t: Target{
				Name: "vimrc",
				File: "./.vimrc",
				Dst:  "~fffsa/dfa",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := doTarget(tt.t); (err != nil) != tt.wantErr {
				t.Errorf("doTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil && tt.wantErr {
				return
			}

			if _, err := os.Stat(tt.t.Dst); err != nil {
				t.Errorf("file %s does not exist", tt.t.Dst)
				return
			}
			tt.after()
		})
	}
}

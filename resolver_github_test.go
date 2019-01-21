package dots

import (
	"io"
	"reflect"
	"testing"
)

func TestNewGithubResolver(t *testing.T) {
	tests := []struct {
		name string
		repo string
		want *GithubResolver
	}{
		{
			name: "matsune/dots_sample",
			repo: "matsune/dots_sample",
			want: &GithubResolver{
				repo:   "matsune/dots_sample",
				host:   "https://raw.githubusercontent.com/",
				branch: "master",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGithubResolver(tt.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGithubResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubResolver_ReadFile(t *testing.T) {
	type fields struct {
		repo   string
		host   string
		branch string
	}
	type args struct {
		sub  string
		file string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &GithubResolver{
				repo:   tt.fields.repo,
				host:   tt.fields.host,
				branch: tt.fields.branch,
			}
			got, err := r.ReadFile(tt.args.sub, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubResolver.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GithubResolver.ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGithubResolver_url(t *testing.T) {
	type fields struct {
		repo   string
		host   string
		branch string
	}
	type args struct {
		sub  string
		file string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success githubresolver_url",
			fields: fields{
				repo:   "matsune/dots_sample",
				host:   "https://raw.githubusercontent.com",
				branch: "master",
			},
			args: args{
				sub:  "zsh",
				file: ".zshrc",
			},
			want: "https://raw.githubusercontent.com/matsune/dots_sample/master/zsh/.zshrc",
		},
		// fail tests
		{
			name: "fail with invalid host",
			fields: fields{
				repo:   "matsune/dots_sample",
				host:   "]]aa://",
				branch: "master",
			},
			args: args{
				sub:  "zsh",
				file: ".zshrc",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &GithubResolver{
				repo:   tt.fields.repo,
				host:   tt.fields.host,
				branch: tt.fields.branch,
			}
			got, err := r.url(tt.args.sub, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GithubResolver.url() got %v, error = %v, wantErr %v", got, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(got.String(), tt.want) {
					t.Errorf("GithubResolver.url().String() = %v, want %v", got.String(), tt.want)
				}
			}
		})
	}
}

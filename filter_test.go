package dots

import (
	"reflect"
	"testing"
)

var all = []Target{
	Target{
		Name: "a",
		Tags: []string{},
	},
	Target{
		Name: "b",
		Tags: []string{},
	},
	Target{
		Name: "c",
		Tags: []string{
			">>>",
		},
	},
	Target{
		Name: "d",
		Tags: []string{
			">>>",
			"???",
		},
	},
}

func Test_filterNames(t *testing.T) {
	tests := []struct {
		name  string
		names []string
		want  []Target
	}{
		{
			name:  "empty name",
			names: []string{},
			want: []Target{
				Target{
					Name: "a",
					Tags: []string{},
				},
				Target{
					Name: "b",
					Tags: []string{},
				},
				Target{
					Name: "c",
					Tags: []string{
						">>>",
					},
				},
				Target{
					Name: "d",
					Tags: []string{
						">>>",
						"???",
					},
				},
			},
		},
		{
			name:  "filter 1 name",
			names: []string{"a"},
			want: []Target{
				Target{
					Name: "a",
					Tags: []string{},
				},
			},
		},
		{
			name:  "filter 2 names",
			names: []string{"a", "b"},
			want: []Target{
				Target{
					Name: "a",
					Tags: []string{},
				},
				Target{
					Name: "b",
					Tags: []string{},
				},
			},
		},
		{
			name:  "filter duplicated names",
			names: []string{"a", "a", "a"},
			want: []Target{
				Target{
					Name: "a",
					Tags: []string{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterNames(all, tt.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterTags(t *testing.T) {
	tests := []struct {
		name string
		tags []string
		want []Target
	}{
		{
			name: "filter empty tag",
			tags: []string{},
			want: []Target{
				Target{
					Name: "a",
					Tags: []string{},
				},
				Target{
					Name: "b",
					Tags: []string{},
				},
			},
		},
		{
			name: "filter 1tag",
			tags: []string{">>>"},
			want: []Target{
				Target{
					Name: "c",
					Tags: []string{
						">>>",
					},
				},
				Target{
					Name: "d",
					Tags: []string{
						">>>",
						"???",
					},
				},
			},
		},
		{
			name: "filter 2tags",
			tags: []string{">>>", "???"},
			want: []Target{
				Target{
					Name: "c",
					Tags: []string{
						">>>",
					},
				},
				Target{
					Name: "d",
					Tags: []string{
						">>>",
						"???",
					},
				},
			},
		},
		{
			name: "filter unknown tag",
			tags: []string{"PPP"},
			want: []Target{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterTags(all, tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filter(t *testing.T) {
	type args struct {
		all     []Target
		targets []string
		tags    []string
	}
	tests := []struct {
		name string
		args args
		want []Target
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.all, tt.args.targets, tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

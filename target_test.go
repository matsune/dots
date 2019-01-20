package dots

import "testing"

func Test_Target_validate(t *testing.T) {
	tests := []struct {
		name    string
		t       Target
		wantErr bool
	}{
		{
			name: "no name",
			t: Target{
				Name: "dd",
				File: "aa",
				Dst:  "bb",
				Tags: []string{"a", "b"},
			},
			wantErr: false,
		},
		{
			name: "no tags",
			t: Target{
				Name: "dd",
				File: "aa",
				Dst:  "bb",
				Tags: []string{},
			},
			wantErr: false,
		},
		// fail tests
		{
			name: "no name",
			t: Target{
				File: "aa",
				Dst:  "bb",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Target.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

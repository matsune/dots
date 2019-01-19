package dots

import "testing"

func Test_target_validate(t *testing.T) {
	tests := []struct {
		name    string
		t       target
		wantErr bool
	}{
		{
			name: "no name",
			t: target{
				Name: "dd",
				File: "aa",
				Dst:  "bb",
			},
			wantErr: false,
		},
		// fail tests
		{
			name: "no name",
			t: target{
				File: "aa",
				Dst:  "bb",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.validate(); (err != nil) != tt.wantErr {
				t.Errorf("target.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

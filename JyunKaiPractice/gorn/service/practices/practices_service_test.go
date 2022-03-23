package practices

import (
	"testing"

	practices "bitbucket.org/henry111666888/hs_gorm/structure/practices"
)

func Test_practices_Created(t *testing.T) {
	type args struct {
		input practices.Table
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case 1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Created(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Created() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

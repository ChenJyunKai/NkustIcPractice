package accounts

import (
	"testing"

	model "bitbucket.org/henry111666888/hs_gorm/structure/accounts"
)

func TestCreated(t *testing.T) {
	type args struct {
		input model.Table
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				input: model.Table{
					Account:  "henry111666888@gmail.com",
					Name:     "Henry",
					Password: "123456",
				},
			},
			wantErr: false,
		}, {
			name: "case 2",
			args: args{
				input: model.Table{
					Account:  "peter@gmail.com",
					Name:     "Peter",
					Password: "654321",
				},
			},
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

func TestList(t *testing.T) {
	type args struct {
		input model.Table
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
			if err := List(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetByAccount(t *testing.T) {
	type args struct {
		input      model.Table
		account_id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				account_id: "e91bc037-2bf3-4075-bb61-059749ba3321",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetByAccount(tt.args.input, tt.args.account_id); (err != nil) != tt.wantErr {
				t.Errorf("GetByAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdated(t *testing.T) {
	type args struct {
		input      model.Table
		account_id string
		account    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				account_id: "e91bc037-2bf3-4075-bb61-059749ba3321",
				account:    "corn@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Updated(tt.args.input, tt.args.account_id, tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Updated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleted(t *testing.T) {
	type args struct {
		input      model.Table
		account_id string
		deleted    bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				account_id: "e91bc037-2bf3-4075-bb61-059749ba3321",
				deleted:    false,
			},
			wantErr: false,
		}, {
			name: "case 2",
			args: args{
				account_id: "61e5a4b6-9550-4025-9d52-d9435e780573",
				deleted:    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Deleted(tt.args.input, tt.args.account_id, tt.args.deleted); (err != nil) != tt.wantErr {
				t.Errorf("Deleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

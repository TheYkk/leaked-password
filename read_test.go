package leakedpassword

import "testing"

func TestIsLeaked(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Exist Password",
			args: args{
				"password",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "NotExist Password",
			args: args{
				"BAB8FC11-4A4C-46BD-9008-B65311148ADF",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsLeaked(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Leaked() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Leaked() got = %v, want %v", got, tt.want)
			}
		})
	}
}

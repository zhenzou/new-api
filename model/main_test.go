package model

import "testing"

func Test_extractTiDBHostname(t *testing.T) {
	type args struct {
		connectionString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TiDB",
			args: args{
				connectionString: "mysql://root:123456@@tcp(gateway01.us-east-1.prod.aws.tidbcloud.com:4000)/new-api?tls=tidb",
			},
			want:    "gateway01.us-east-1.prod.aws.tidbcloud.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractTiDBHostname(tt.args.connectionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractTiDBHostname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractTiDBHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

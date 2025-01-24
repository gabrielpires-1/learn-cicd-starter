package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		errMsg  string
	}{
		{
			"no auth header",
			args{http.Header{}},
			"",
			true,
			"no authorization header included",
		},
		{
			"malformed auth header",
			args{http.Header{"Authorization": []string{"Bearer"}}},
			"",
			true,
			"malformed authorization header",
		},
		{
			"valid auth header",
			args{http.Header{"Authorization": []string{"ApiKey 123456"}}},
			"123456",
			false,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

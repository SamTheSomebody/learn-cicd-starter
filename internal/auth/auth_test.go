package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	AuthHeaderNoValue := http.Header{"Authorization": []string{}}
	MalformedAuthHeaderNoSeperator := http.Header{"Authorization": []string{"ApiKey012345"}}
	MalformedAuthHeaderInvalidKey := http.Header{"Authorization": []string{"ApiKay 012345"}}
	ValidAuthHeader := http.Header{"Authorization": []string{"ApiKey 012345"}}

	tests := map[string]struct {
		args    http.Header
		want    string
		wantErr bool
	}{
		"No Auth Header":                     {args: make(http.Header), want: "", wantErr: true},
		"No Auth Header Value":               {args: AuthHeaderNoValue, want: "", wantErr: true},
		"Malformed Auth Header No Seperator": {args: MalformedAuthHeaderNoSeperator, want: "", wantErr: true},
		"Malformed Auth Header Invalid Key":  {args: MalformedAuthHeaderInvalidKey, want: "", wantErr: true},
		"Valid Auth Header":                  {args: ValidAuthHeader, want: "012345", wantErr: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args)
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

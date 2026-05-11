package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		headers http.Header
		want    string
		err     error
	}

	tests := map[string]test{
		"valid_api_key": {
			headers: http.Header{"Authorization": []string{"ApiKey 12345abcdeZ"}},
			want:    "12345abcde",
			err:     nil,
		},
		"invalid_authorization_header": {
			headers: http.Header{"Authorization": []string{"Bearer token123"}},
			want:    "",
			err:     errors.New("malformed authorization header"),
		},
		"no_authorization_header": {
			headers: http.Header{},
			want:    "",
			err:     ErrNoAuthHeaderIncluded,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if err != nil {
				if tc.err == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error %v, got %v", tc.err, err)
				}
			} else {
				if tc.err != nil {
					t.Fatalf("expected error %v, got none", tc.err)
				}
				if got != tc.want {
					t.Errorf("got %s, want %s", got, tc.want)
				}
			}
		})
	}
}

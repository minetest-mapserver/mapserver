package web

import (
	"mapserver/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newAuthTestApi() *Api {
	return NewApi(&app.App{Config: &app.Config{WebApi: &app.WebApiConfig{SecretKey: "s3cret"}}})
}

func TestCheckAuth(t *testing.T) {
	tests := []struct {
		name   string
		header *string
		ok     bool
	}{
		{"valid key", ptr("s3cret"), true},
		{"wrong key", ptr("other"), false},
		{"empty key", ptr(""), false},
		{"missing header", nil, false},
		{"key with whitespace", ptr(" s3cret"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/x", nil)
			if tc.header != nil {
				req.Header.Set("Authorization", *tc.header)
			}
			rec := httptest.NewRecorder()

			assert.Equal(t, tc.ok, newAuthTestApi().check_auth(rec, req))
			if tc.ok {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Empty(t, rec.Body.String())
			} else {
				assert.Equal(t, http.StatusForbidden, rec.Code)
				assert.Equal(t, "invalid key!", rec.Body.String())
			}
		})
	}
}

func ptr(s string) *string { return &s }

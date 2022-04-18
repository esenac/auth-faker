package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esenac/auth-faker/jwt"
)

type mockTokenGetter struct {
	ret string
	err error
}

func (m mockTokenGetter) GetSignedToken(sub, iss, aud, scope string, customClaims jwt.CustomClaims) (string, error) {
	return m.ret, m.err
}

func TestCreateTokenHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	body := map[string]interface{}{
		"custom_claims": nil,
		"subject":       "test-subject",
		"audience":      "test-audience",
		"scope":         "test-scope",
		"issuer":        "test-issuer",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fail()
	}
	req := httptest.NewRequest(http.MethodPost, "http://localhost", bytes.NewReader(jsonBody))

	srv := mockTokenGetter{ret: "signedToken", err: nil}
	CreateTokenHandler(srv)(recorder, req)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}

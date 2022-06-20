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

func (m mockTokenGetter) GetSignedToken(_, _, _, _ string, _ jwt.CustomClaims) (string, error) {
	return m.ret, m.err
}

func TestCreateTokenHandler(t *testing.T) {
	tests := []struct {
		testCase           string
		input              map[string]interface{}
		expectedStatusCode int
		expectedBody       string
	}{
		{
			testCase: "Given a valid payload the token is correctly generated",
			input: map[string]interface{}{
				"custom_claims": nil,
				"subject":       "test-subject",
				"audience":      "test-audience",
				"scope":         "test-scope",
				"issuer":        "test-issuer",
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       "{\"token\":\"signedToken\"}",
		},
		{
			testCase: "Given an invalid payload we get a 400 status code",
			input: map[string]interface{}{
				"custom_claims": nil,
				"issuer":        435,
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "json: cannot unmarshal number into Go struct field TokenRequest.issuer of type string\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			body := tc.input
			jsonBody, err := json.Marshal(body)
			if err != nil {
				t.Fail()
			}
			req := httptest.NewRequest(http.MethodPost, "http://localhost", bytes.NewReader(jsonBody))

			srv := mockTokenGetter{ret: "signedToken", err: nil}
			CreateTokenHandler(srv)(recorder, req)
			if recorder.Result().StatusCode != tc.expectedStatusCode {
				t.Logf("Expected status code: %d - received status code: %d", tc.expectedStatusCode, recorder.Result().StatusCode)
				t.Fail()
			}
			if recorder.Body.String() != tc.expectedBody {
				t.Logf("Expected body: %s - received body: %s", tc.expectedBody, recorder.Body.String())
				t.Fail()
			}
		})
	}
}

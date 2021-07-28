package http

import (
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	CustomClaims map[string]interface{} `json:"custom_claims"`
	Subject      string                 `json:"subject"`
	Audience     string                 `json:"audience"`
	Scope        string                 `json:"scope"`
}

func GetTokenRequest(r *http.Request) (TokenRequest, error) {
	var result TokenRequest
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return TokenRequest{}, err
	}

	return result, nil
}

package main

import (
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	CustomClaims map[string]interface{} `json:"custom_claims"`
	//Username          string `json:"username"`
	Subject  string `json:"subject"`
	Audience string `json:"audience"`
	//AccountId         int    `json:"account_id"`
	//UserId            int    `json:"user_id"`
	//ExternalAccountId string `json:"external_account_id"`
	//ExternalUserId    string `json:"external_user_id"`
	//TenantId          string `json:"tenant_id"`
	//Email             string `json:"email"`
}

func GetTokenRequest(r *http.Request) (TokenRequest, error) {
	var result TokenRequest
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return TokenRequest{}, err
	}

	return result, nil
}

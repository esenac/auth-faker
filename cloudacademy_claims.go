package main

import "github.com/dgrijalva/jwt-go"

type CloudAcademyClaims struct {
	Username             string `json:"username"`
	AccountId            int    `json:"account_id"`
	UserId               int    `json:"user_id"`
	ExternalAccountId    string `json:"external_account_id"`
	ExternalUserId       string `json:"external_user_id"`
	TenantId             string `json:"tenant_id"`
	Email                string `json:"email"`
	OriginalIssuedAtTime int64  `json:"orig_iat"`
	jwt.StandardClaims
}

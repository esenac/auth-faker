package main

import "github.com/dgrijalva/jwt-go"

type CloudAcademyClaims struct {
	Username             string `json:"https://clda.co/claims/username"`
	AccountId            int    `json:"https://clda.co/claims/account_id"`
	UserId               int    `json:"https://clda.co/claims/user_id"`
	ExternalAccountId    string `json:"https://clda.co/claims/external_account_id"`
	ExternalUserId       string `json:"https://clda.co/claims/external_user_id"`
	TenantId             string `json:"https://clda.co/claims/tenant_id"`
	Email                string `json:"https://clda.co/claims/email"`
	OriginalIssuedAtTime int64  `json:"orig_iat"`
	Scope string `json:"scope"`
	jwt.StandardClaims
}

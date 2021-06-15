package main

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var key *rsa.PrivateKey

func init() {
	key = LoadRSAPrivateKeyFromDisk("./resources/jwt-keys/private")
}

func GetTokenHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenRequest, err := GetTokenRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		claims := CloudAcademyClaims{
			tokenRequest.Username,
			tokenRequest.AccountId,
			tokenRequest.UserId,
			tokenRequest.ExternalAccountId,
			tokenRequest.ExternalUserId,
			tokenRequest.TenantId,
			tokenRequest.Email,
			time.Now().Unix(),
			jwt.StandardClaims{
				Subject:   tokenRequest.Subject,
				Issuer:    "auth-faker",
				Audience:  tokenRequest.Audience,
				ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		s, e := token.SignedString(key)

		if e != nil {
			panic(e.Error())
		}

		w.Header().Add("Content-Type", "application/json")
		jsonBody, err := json.Marshal(map[string]string{"token": s})
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		_, _ = w.Write(jsonBody)
	}
}

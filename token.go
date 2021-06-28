package main

import (
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key *rsa.PrivateKey
var issuer string

func init() {
	key = LoadRSAPrivateKeyFromDisk("./resources/certs/certificate.key.pem")
	issuer = os.Getenv("TOKEN_ISSUER")
	if issuer == "" {
		issuer = "auth-faker"
	}
}

func CreateTokenHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenRequest, err := GetTokenRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		claims := Claims{}
		claims.StandardClaims =
			jwt.StandardClaims{
				Subject:   tokenRequest.Subject,
				Issuer:    issuer,
				Audience:  tokenRequest.Audience,
				ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			}
		for k, v := range tokenRequest.CustomClaims {
			claims.Add(k, v)
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

package http

import (
	"encoding/json"
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/esenac/auth-faker/jwt"
)

const contentTypeHeader = "Content-Type"

func GetHandler(data interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(contentTypeHeader, "application/json")
		jsonBody, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		w.Write(jsonBody)
	}
}

func CreateTokenHandler(key interface{}, issuer string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenRequest, err := GetTokenRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		claims := jwt.New(tokenRequest.Subject, issuer, tokenRequest.Audience, tokenRequest.CustomClaims)

		token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
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

package http

import (
	"encoding/json"
	"net/http"

	"github.com/esenac/auth-faker/jwt"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

type TokenGetter interface {
	GetSignedToken(sub, iss, aud, scope string, customClaims jwt.CustomClaims) (string, error)
}

func GetHandler(data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(contentTypeHeader, jsonContentType)
		jsonBody, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		_, _ = w.Write(jsonBody)
	}
}

func CreateTokenHandler(tokenGetter TokenGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenRequest, err := decodeRequest[TokenRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		signedString, err := tokenGetter.GetSignedToken(
			tokenRequest.Subject,
			tokenRequest.Issuer,
			tokenRequest.Audience,
			tokenRequest.Scope,
			tokenRequest.CustomClaims)
		if err != nil {
			panic(err.Error())
		}

		w.Header().Add(contentTypeHeader, jsonContentType)
		jsonBody, err := json.Marshal(map[string]string{"token": signedString})
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		_, _ = w.Write(jsonBody)
	}
}

func decodeRequest[T any](r *http.Request) (T, error) {
	var result T
	err := json.NewDecoder(r.Body).Decode(&result)
	return result, err
}

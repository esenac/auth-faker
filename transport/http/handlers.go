package http

import (
	"encoding/json"
	"net/http"

	"github.com/esenac/auth-faker/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

// TokenGetter defines a set of methods to retrieve a token.
type TokenGetter interface {
	GetSignedToken(sub, iss, aud, scope string, customClaims jwt.CustomClaims) (string, error)
}

// KeySetGetter defines a set of methods to retrieve a jwk.Set.
type KeySetGetter interface {
	GetKeySet() jwk.Set
}

// GetJWKSHandler handles the JWKS retrieval.
func GetJWKSHandler(keySetGetter KeySetGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		set := keySetGetter.GetKeySet()

		w.Header().Add(contentTypeHeader, jsonContentType)
		jsonBody, err := json.Marshal(set)
		if err != nil {
			http.Error(w, "Error converting results to json",
				http.StatusInternalServerError)
		}
		_, _ = w.Write(jsonBody)
	}
}

// CreateTokenHandler handles the creation of a token
func CreateTokenHandler(tokenGetter TokenGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenRequest, err := decodeRequest[TokenRequest](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
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

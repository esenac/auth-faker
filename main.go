package main

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	key := LoadRSAPrivateKeyFromDisk("/keys/private")

	claims := jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	publicKey := LoadRSAPublicKeyFromDisk("/keys/public")
	k, e := jwk.New(publicKey)

	if e != nil {
		panic(e.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/get-token", GetHandler(map[string]string{"token": s}))
	mux.HandleFunc("/.well-known/jwks.json", GetHandler(k))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

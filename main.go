package main

import (
	"log"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	publicKey := LoadRSAPublicKeyFromDisk("./resources/jwt-keys/public")
	k, e := jwk.New(publicKey)

	if e != nil {
		panic(e.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/token", GetTokenHandler())
	mux.HandleFunc("/.well-known/jwks.json", GetHandler(k))

	log.Fatal(http.ListenAndServeTLS(":8080", "./resources/certs/auth-faker+3.pem", "./resources/certs/auth-faker+3-key.pem", mux))
}

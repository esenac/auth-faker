package main

import (
	"github.com/lestrrat-go/jwx/jwk"
	"log"
	"net/http"
	"strings"
)


type JWKSResult struct {
	Keys [1]jwk.Key `json:"keys"`
}

func main() {
	location := "./resources/certs/certificate.pem"
	publicKey := LoadRSAPublicKeyFromDisk(location)
	x5cBytes, e := ReadKeyAsX5C(location)
	x5c := strings.Replace(strings.Replace(string(x5cBytes), "-----BEGIN CERTIFICATE-----\n", "", -1),
		"\n-----END CERTIFICATE-----", "", -1)

	if e != nil {
		panic(e.Error())
	}

	k, e := jwk.New(publicKey)
	if e != nil {
		panic(e.Error())
	}
	err := k.Set("x5c", x5c)
	if err != nil {
		panic(e.Error())
	}

	var keys [1]jwk.Key
	keys[0] = k
	result := JWKSResult{
		Keys: keys,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/token", GetTokenHandler())
	mux.HandleFunc("/.well-known/jwks.json", GetHandler(result))

	log.Fatal(http.ListenAndServe(":80", mux))
}

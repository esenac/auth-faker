package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"
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

	keyset := jwk.NewSet()
	keyset.Add(k)
	mux := http.NewServeMux()


	mux.HandleFunc("/token", GetTokenHandler())
	mux.HandleFunc("/.well-known/jwks.json", GetHandler(keyset))

	log.Fatal(http.ListenAndServe(":80", mux))
}

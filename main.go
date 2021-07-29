package main

import (
	"log"
	"strings"

	"github.com/esenac/auth-faker/transport/http"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	location := "./resources/certs/certificate.pem"
	publicKey := LoadRSAPublicKeyFromDisk(location)
	x5cBytes, e := ReadKeyAsX5C(location)
	if e != nil {
		panic(e.Error())
	}
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

	server := http.New()
	server.AddRoute("/token", http.CreateTokenHandler(key))
	server.AddRoute("/.well-known/jwks.json", http.GetHandler(keyset))
	log.Fatal(server.Start(80))
}

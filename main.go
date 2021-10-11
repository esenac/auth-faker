package main

import (
	"log"
	"os"
	"strings"

	"github.com/esenac/auth-faker/jwk"
	"github.com/esenac/auth-faker/transport/http"
)

func main() {
	location := "./resources/certs/certificate.pem"
	publicKey, err := LoadRSAPublicKeyFromDisk(location)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	certificateKey, err := LoadRSAPrivateKeyFromDisk("./resources/certs/certificate.key.pem")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	x5cBytes, err := ReadFile(location)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	x5c := strings.Replace(strings.Replace(string(x5cBytes), "-----BEGIN CERTIFICATE-----\n", "", -1),
		"\n-----END CERTIFICATE-----", "", -1)

	key, err := jwk.NewKey(publicKey, &x5c)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	keyset := jwk.GetKeyset(key)
	server := http.New()
	server.AddRoute("/token", http.CreateTokenHandler(certificateKey))
	server.AddRoute("/.well-known/jwks.json", http.GetHandler(keyset))
	log.Fatal(server.Start(80))
}

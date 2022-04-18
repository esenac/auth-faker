package main

import (
	"log"

	"github.com/esenac/auth-faker/jwk"
	"github.com/esenac/auth-faker/jwt"
	"github.com/esenac/auth-faker/keys"
	"github.com/esenac/auth-faker/transport/http"
)

func main() {
	keysManager := keys.NewManager("./resources/certs/certificate.pem", "./resources/certs/certificate.key.pem")
	publicKey, err := keysManager.LoadPublicKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	certificateKey, err := keysManager.LoadPrivateKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	x5c, err := keysManager.LoadX5C()
	if err != nil {
		log.Fatal(err.Error())
	}
	key, err := jwk.NewKey(publicKey, x5c)
	if err != nil {
		log.Fatal(err.Error())
	}
	keyset := jwk.GetKeyset(key)

	jwtService := jwt.NewService(certificateKey)
	server := http.New()
	server.AddRoute("/token", http.CreateTokenHandler(jwtService))
	server.AddRoute("/.well-known/jwks.json", http.GetHandler(keyset))
	log.Fatal(server.Start(80))
}

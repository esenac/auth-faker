package main

import (
	"log"

	"github.com/esenac/auth-faker/jwk"
	"github.com/esenac/auth-faker/jwt"
	"github.com/esenac/auth-faker/keys"
	"github.com/esenac/auth-faker/transport/http"
)

func main() {
	keysManager := keys.NewLoader("./resources/certs/certificate.pem", "./resources/certs/certificate.key.pem")
	jwkService, err := jwk.NewService(keysManager)
	if err != nil {
		log.Fatal(err.Error())
	}

	jwtService, err := jwt.NewService(keysManager)
	if err != nil {
		log.Fatal(err.Error())
	}
	server := http.New()
	server.AddRoute("/token", http.CreateTokenHandler(jwtService))
	server.AddRoute("/.well-known/jwks.json", http.GetJWKSHandler(jwkService))
	log.Fatal(server.Start(80))
}

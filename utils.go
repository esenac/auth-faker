package main

import (
	"crypto/rsa"
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var key *rsa.PrivateKey
var issuer string

func init() {
	key = LoadRSAPrivateKeyFromDisk("./resources/certs/certificate.key.pem")
	issuer = os.Getenv("TOKEN_ISSUER")
	if issuer == "" {
		issuer = "auth-faker"
	}
}

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := ReadFile(location)
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func ReadFile(location string) ([]byte, error) {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	return keyData, e
}

func ReadKeyAsX5C(location string) ([]byte, error) {
	return ReadFile(location)
}

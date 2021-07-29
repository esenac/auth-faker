package main

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

func LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}

func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	keyData, err := ReadFile(location)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func ReadFile(location string) ([]byte, error) {
	return ioutil.ReadFile(location)
}

package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

type Service struct {
	key interface{}
}

func NewService(signingKey interface{}) Service {
	return Service{key: signingKey}
}

func (s Service) GetSignedToken(sub, iss, aud, scope string, customClaims CustomClaims) (string, error) {
	claims := newClaims(
		sub, iss, aud, scope, customClaims)

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	return token.SignedString(s.key)
}

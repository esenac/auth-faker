package jwt

import (
	"crypto/rsa"
	"github.com/lestrrat-go/jwx/jwk"

	jwtgo "github.com/dgrijalva/jwt-go"
)

// PrivateKeyLoader defines a set of methods to load a private key.
type PrivateKeyLoader interface {
	LoadPrivateKey() (*rsa.PrivateKey, error)
}

// Service manages the creation and retrieval of a JWT token.
type Service struct {
	key   *rsa.PrivateKey
	keyId string
}

// NewService creates a Service using the provided loader to retrieve the private key.
func NewService(keyId string, loader PrivateKeyLoader) (*Service, error) {
	key, err := loader.LoadPrivateKey()
	if err != nil {
		return nil, err
	}
	return &Service{key: key, keyId: keyId}, nil
}

// GetSignedToken retrieves a signed JWT token with provided sub, iss, aud, scope and customClaims.
func (s Service) GetSignedToken(sub, iss, aud, scope string, customClaims CustomClaims) (string, error) {
	claims := newClaims(
		sub, iss, aud, scope, customClaims)

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)
	token.Header[jwk.KeyIDKey] = s.keyId
	return token.SignedString(s.key)
}

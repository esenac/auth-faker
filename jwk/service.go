package jwk

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
)

// PublicKeyLoader defines a set of methods to load a public key.
type PublicKeyLoader interface {
	LoadPublicKey() (*rsa.PublicKey, error)
	LoadX5C() (*string, error)
}

// Service manages the creation and retrieval of a JWK Key and Set.
type Service struct {
	key jwk.Key
}

// NewService creates a Service using the provided keyLoader to retrieve the public key.
func NewService(keyLoader PublicKeyLoader) (*Service, error) {
	key, err := createKey(keyLoader)
	if err != nil {
		return nil, err
	}
	return &Service{key: key}, nil
}

// GetKeySet returns a jwk.Set containing the generated public Key.
func (s Service) GetKeySet() jwk.Set {
	return newKeyset(s.key)
}

func createKey(keyLoader PublicKeyLoader) (jwk.Key, error) {
	publicKey, err := keyLoader.LoadPublicKey()
	if err != nil {
		return nil, err
	}

	x5c, err := keyLoader.LoadX5C()
	if err != nil {
		return nil, err
	}
	return NewKey(publicKey, x5c)
}

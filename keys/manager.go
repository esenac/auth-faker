package keys

import (
	"crypto/rsa"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Loader loads public and private keys from provided locations.
type Loader struct {
	publicKeyLocation  string
	privateKeyLocation string
}

// NewLoader creates a Loader to load the public and private keys
// from the given publicKeyLocation and privateKeyLocation, respectively.
func NewLoader(publicKeyLocation, privateKeyLocation string) *Loader {
	return &Loader{
		publicKeyLocation:  publicKeyLocation,
		privateKeyLocation: privateKeyLocation,
	}
}

// LoadPublicKey retrieves the public key from the disk and returns a rsa.PublicKey
func (l *Loader) LoadPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(l.publicKeyLocation)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

// LoadPrivateKey retrieves the public key from the disk and returns a rsa.PrivateKey
func (l *Loader) LoadPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(l.privateKeyLocation)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}

// LoadX5C generates the x5c parameter starting from the public key.
func (l *Loader) LoadX5C() (*string, error) {
	x5cBytes, err := ioutil.ReadFile(l.publicKeyLocation)
	if err != nil {
		return nil, err
	}
	x5c := strings.Replace(strings.Replace(string(x5cBytes), "-----BEGIN CERTIFICATE-----\n", "", -1),
		"\n-----END CERTIFICATE-----", "", -1)
	return &x5c, nil
}

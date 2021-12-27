package keys

import (
	"crypto/rsa"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Manager struct {
	publicKeyLocation  string
	privateKeyLocation string
}

func NewManager(publicKeyLocation, privateKeyLocation string) *Manager {
	return &Manager{
		publicKeyLocation:  publicKeyLocation,
		privateKeyLocation: privateKeyLocation,
	}
}

func (m *Manager) LoadPublicKey() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(m.publicKeyLocation)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(keyData)
}

func (m *Manager) LoadPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(m.privateKeyLocation)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(keyData)
}

func (m *Manager) LoadX5C() (*string, error) {
	x5cBytes, err := ioutil.ReadFile(m.publicKeyLocation)
	if err != nil {
		return nil, err
	}
	x5c := strings.Replace(strings.Replace(string(x5cBytes), "-----BEGIN CERTIFICATE-----\n", "", -1),
		"\n-----END CERTIFICATE-----", "", -1)
	return &x5c, nil
}

package jwk

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/lestrrat-go/jwx/jwk"
)

type mockPublicKeyLoader struct {
	pubKey           *rsa.PublicKey
	errLoadPublicKey error
	x5c              *string
	errLoadX5C       error
}

func (m mockPublicKeyLoader) LoadPublicKey() (r *rsa.PublicKey, err error) {
	return m.pubKey, m.errLoadPublicKey
}

func (m mockPublicKeyLoader) LoadX5C() (*string, error) {
	return m.x5c, m.errLoadX5C
}

func TestCreateKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	tests := []struct {
		testCase    string
		keyLoader   PublicKeyLoader
		expectedKey jwk.Key
		expectedErr error
	}{
		{
			testCase:    "The Public Key is successfully loaded",
			keyLoader:   mockPublicKeyLoader{pubKey: &privateKey.PublicKey},
			expectedKey: nil,
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			key, err := createKey(tc.keyLoader)
			if err != tc.expectedErr {
				t.Logf("Expected error: %v - received error: %v", tc.expectedErr, err)
				t.Fail()
			}
			if key != tc.expectedKey {
				t.Logf("Expected key: %v - received key: %v", tc.expectedKey, key)
				t.Fail()
			}
		})
	}
}

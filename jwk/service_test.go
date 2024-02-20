package jwk

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
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
	keyLoaderErr := errors.New("key loader err")
	x5cErr := errors.New("x5c loader err")
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	tests := []struct {
		testCase    string
		keyLoader   PublicKeyLoader
		expectKey   bool
		expectedErr error
	}{
		{
			testCase:    "The Public Key is successfully loaded",
			keyLoader:   mockPublicKeyLoader{pubKey: &privateKey.PublicKey},
			expectKey:   true,
			expectedErr: nil,
		},
		{
			testCase:    "Key loader returns error",
			keyLoader:   mockPublicKeyLoader{errLoadPublicKey: keyLoaderErr},
			expectKey:   false,
			expectedErr: keyLoaderErr,
		},
		{
			testCase:    "x5c loader returns error",
			keyLoader:   mockPublicKeyLoader{pubKey: &privateKey.PublicKey, errLoadX5C: x5cErr},
			expectKey:   false,
			expectedErr: x5cErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			key, err := createKey("some-key-id", tc.keyLoader)
			if err != tc.expectedErr {
				t.Logf("Expected error: %v - received error: %v", tc.expectedErr, err)
				t.Fail()
			}
			if _, ok := key.(jwk.Key); tc.expectKey && !ok {
				t.Logf("Expected key type: jwk.Key - received key: %v", key)
				t.Fail()
			}
		})
	}
}

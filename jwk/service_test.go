package jwk

import (
	"testing"

	"github.com/lestrrat-go/jwx/jwk"
)

func TestCreateKey(t *testing.T) {
	tests := []struct {
		testCase    string
		keyLoader   PublicKeyLoader
		expectedKey jwk.Key
		expectedErr error
	}{
		{
			testCase:    "TBD",
			keyLoader:   nil,
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

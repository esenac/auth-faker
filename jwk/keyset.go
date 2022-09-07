package jwk

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
)

func newKeyset(keys ...jwk.Key) jwk.Set {
	keySet := jwk.NewSet()
	for _, k := range keys {
		keySet.Add(k)
	}
	return keySet
}

// NewKey creates a JWK key from the given public key and x5c.
func NewKey(publicKey *rsa.PublicKey, x5c *string) (jwk.Key, error) {
	k, err := jwk.New(publicKey)
	if err != nil {
		return nil, err
	}
	if x5c != nil {
		err = k.Set("x5c", *x5c)
	}
	return k, err
}

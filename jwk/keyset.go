package jwk

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/jwk"
)

func GetKeyset(keys ...jwk.Key) jwk.Set {
	keyset := jwk.NewSet()
	for _, k := range keys {
		keyset.Add(k)
	}
	return keyset
}

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

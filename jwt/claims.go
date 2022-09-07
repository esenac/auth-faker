package jwt

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims represents the JWT claims.
type Claims struct {
	jwt.StandardClaims
	jwt.MapClaims
}

// CustomClaims represents a set of custom JWT claims different from the standard ones.
type CustomClaims jwt.MapClaims

func newClaims(subject string, issuer string, audience string, scope string, customClaims CustomClaims) Claims {
	claims := Claims{
		jwt.StandardClaims{
			Subject:   subject,
			Issuer:    issuer,
			Audience:  audience,
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
		jwt.MapClaims{
			"scope": scope,
		},
	}

	for k, v := range customClaims {
		claims.Add(k, v)
	}
	return claims
}

// Valid validates time based claims "exp, iat, nbf". It's based on jwt.StandardClaims Valid() method.
func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}

// Add adds a claim to the JWT.
func (c *Claims) Add(k string, v interface{}) {
	if c.MapClaims == nil {
		c.MapClaims = jwt.MapClaims{}
	}
	c.MapClaims[k] = v
}

// MarshalJSON returns the JSON encoding of c.
func (c Claims) MarshalJSON() ([]byte, error) {
	finalMap := make(map[string]interface{})
	structType := reflect.TypeOf(c.StandardClaims)
	rt := reflect.ValueOf(&c.StandardClaims).Elem()
	for i := 0; i < rt.NumField(); i++ {
		f := structType.Field(i)

		v := strings.Split(f.Tag.Get("json"), ",")[0]
		if v != "" && rt.Field(i).Interface() != reflect.Zero(f.Type).Interface() {
			finalMap[v] = rt.Field(i).Interface()
		}
	}

	for k, v := range c.MapClaims {
		finalMap[k] = v
	}

	return json.Marshal(finalMap)
}

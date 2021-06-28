package main

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	jwt.MapClaims
}

func (c Claims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *Claims) Add(k string, v interface{}) {
	if c.MapClaims == nil {
		c.MapClaims = jwt.MapClaims{}
	}
	c.MapClaims[k] = v
}

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

package lambdas

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwk struct {
	Alg string   `json:"alg"`
	Kty string   `json:"kty"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	Kid string   `json:"kid"`
	X5t string   `json:"x5t"`
	X5c []string `json:"x5c"`
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	jwkString := os.Getenv("JWKS")
	var jwks []Jwk
	err := json.Unmarshal([]byte(jwkString), &jwks)
	if err != nil {
		return nil, err
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing 'kid' in token header")
		}

		for _, jwk := range jwks {
			if jwk.Kid == kid {
				n := jwk.N
				e := jwk.E

				// Decode base64url-encoded values
				nBytes, err := base64.RawURLEncoding.DecodeString(n)
				if err != nil {
					return nil, err
				}
				eBytes, err := base64.RawURLEncoding.DecodeString(e)
				if err != nil {
					return nil, err
				}

				// Construct the RSA public key
				publicKey := &rsa.PublicKey{
					N: new(big.Int).SetBytes(nBytes),
					E: int(new(big.Int).SetBytes(eBytes).Int64()),
				}

				return publicKey, nil
			}
		}
		return nil, fmt.Errorf("no matching 'kid' found in JWK keys")
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["exp"] == nil || !checkExpiration(claims["exp"]) {
			return nil, fmt.Errorf("token is expired or exp claim is missing")
		}
	} else {
		return nil, err
	}

	return token, nil
}

func checkExpiration(timestamp interface{}) bool {
	var expTime int64

	switch v := timestamp.(type) {
	case float64:
		expTime = int64(v)
	case json.Number:
		expTime, _ = v.Int64()
	}

	return time.Unix(expTime, 0).After(time.Now())
}

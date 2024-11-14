package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	// ssh-keygen -t ecdsa -b 521 -m PEM -f example
	privKey, err := os.ReadFile("./example")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(privKey)

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	privateKey := key
	publicKey := &key.PublicKey

	accessToken := newToken(privateKey)

	fmt.Println("access token is: ", accessToken)

	readToken(accessToken, publicKey)

	jwkPublicKey := jwk.NewECDSAPublicKey()
	if err := jwkPublicKey.FromRaw(publicKey); err != nil {
		panic(err)
	}

	jwkBytes, err := json.Marshal(jwkPublicKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("jwk public key: %s\n", jwkBytes)
}

func newToken(privateKey *ecdsa.PrivateKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES512,
		jwt.MapClaims{
			"jti": uuid.NewString(),
			"iss": "my-app",
			"exp": time.Now().Add(3 * time.Minute).Unix(),
			"uid": "123",
		},
	)

	accessToken, err := token.SignedString(privateKey)
	if err != nil {
		panic(fmt.Errorf("jwt: %w", err))
	}

	return accessToken
}

func readToken(token string, publicKey *ecdsa.PublicKey) {
	mapClaims := new(jwt.MapClaims)

	_, err := jwt.ParseWithClaims(token, mapClaims, func(t *jwt.Token) (interface{}, error) {
		// if t.Method.Alg() != publicKey. {
		// 	return nil, ErrInvalidSigningAlgo
		// }

		return publicKey, nil
	})
	if err != nil {
		panic(err)
	}

	var userID string

	userIDRaw, ok := (*mapClaims)["uid"]
	if ok {
		userID = userIDRaw.(string)
	}

	fmt.Println(userID)
}

// access token is:  eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE0MTI0NDgsImlzcyI6Im15LWFwcCIsImp0aSI6IjZjYmZiZTVlLWY2OGUtNDIxMC1iMDUwLTI5ODA1Y2IzZTNmMCIsInVpZCI6IjEyMyJ9.AB1tDxzBStglHHZHr-oP-gWjxop3e9AamuRmsZ3fCfMj-b5EZA700cg8LGxwLwF50YDK-CVcAclJFxu3l_zGuCnMAKzK23lUZXEtDmjY15KKZz55yfH_fuVnF18PBEp0jfdeF68xbz_EcbOh2O7BoA7iOWjbD8wEvHF9Iea2PXSyZPdc
// 123
// jwk public key: {"crv":"P-521","kty":"EC","x":"AVWkeNAk21b71foCAZoI8DSReMqyMp0eFnUjyd-O9eZzBgyG_BnS975kdIg8LZuixTIaVIAyLKINvPAnnASSnigK","y":"AHdW6YDlxH4XVot5romqN7rGcyfFDfz0fgr-uordjFYoFd8EwlBjzst0Yt_uIsSEIknSgyucfDD6LlQGyBV8rpEV

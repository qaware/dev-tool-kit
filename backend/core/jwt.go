package core

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/cristalhq/jwt/v3"
	"github.com/square/go-jose/v3"
	"golang.org/x/crypto/ed25519"
	"strings"
)

func DecodeJwt(input string, key string) (bool, string, error) {
	if input == "" {
		return false, "", errors.New("Empty JWT")
	}

	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\n", "")

	token, err := jwt.ParseString(input)
	if err != nil {
		DebugError(err)
		return false, "", errors.New("Error parsing JWT")
	}

	payload, err := FormatJson(string(token.RawClaims()))
	if err != nil {
		DebugError(err)
		payload = string(token.RawClaims())
	}

	verifier, err := selectVerifierByAlgorithm(token.Header().Algorithm, []byte(key))
	if err != nil {
		DebugError(err)
		return false, payload, &Information{"Signature verification failed: " + err.Error()}
	}

	valid := false
	err = verifier.Verify(token.Payload(), token.Signature())
	if err != nil {
		DebugError(err)
	} else {
		valid = true
	}

	return valid, payload, nil
}

func selectVerifierByAlgorithm(algorithm jwt.Algorithm, key []byte) (jwt.Verifier, error) {
	if algorithm == jwt.HS256 || algorithm == jwt.HS384 || algorithm == jwt.HS512 {
		if len(key) == 0 {
			return nil, errors.New("Empty key")
		}
		return jwt.NewVerifierHS(algorithm, key)
	}

	var publicKey interface{}

	jwk, err := extractKeyFromJwks(key)
	if err == nil {
		publicKey = jwk
	} else {
		block, _ := pem.Decode(key)
		if block == nil || block.Type != "PUBLIC KEY" {
			return nil, errors.New("Invalid public key format")
		}

		publicKey, err = x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			DebugError(err)
			return nil, errors.New("Invalid public key data")
		}
	}

	switch algorithm {
	case jwt.RS256, jwt.RS384, jwt.RS512:
		return jwt.NewVerifierRS(algorithm, publicKey.(*rsa.PublicKey))
	case jwt.PS256, jwt.PS384, jwt.PS512:
		return jwt.NewVerifierPS(algorithm, publicKey.(*rsa.PublicKey))
	case jwt.ES256, jwt.ES384, jwt.ES512:
		return jwt.NewVerifierES(algorithm, publicKey.(*ecdsa.PublicKey))
	case jwt.EdDSA:
		verifier, err := jwt.NewVerifierEdDSA(publicKey.(ed25519.PublicKey))
		if err != nil {
			return nil, errors.New("Invalid public key length")
		}
		return verifier, nil
	}

	return nil, errors.New("Unsupported signature algorithm")
}

func extractKeyFromJwks(keyJson []byte) (interface{}, error) {
	var jwk jose.JSONWebKey
	err := jwk.UnmarshalJSON(keyJson)
	if err != nil {
		return nil, err
	}

	return jwk.Key, nil
}

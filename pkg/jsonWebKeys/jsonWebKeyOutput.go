package jsonWebKeys

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type JsonWebKeyOutput struct {
	JsonWebKey             jwk.Key `json:"jwk,omitempty"`
	JsonWebKeyPublic       jwk.Key `json:"jwk_public_key,omitempty"`
	JsonWebKeyString       string  `json:"jwk_string,omitempty"`
	JsonWebKeyPublicString string  `json:"jwk_public_string,omitempty"`
	Base64JsonWebKey       string  `json:"base64_jwk,omitempty"`
	AlgPublicKey           string  `json:"alg_public_key,omitempty"`
	AlgPrivateKey          string  `json:"alg_private_key,omitempty"`
	Algorithm              string  `json:"algorithm,omitempty"`
}

func NewJsonWebKeyOutput(
	jsonWebKey jwk.Key,
	algPrivateKey string,
	algPublicKey string) (*JsonWebKeyOutput, error) {
	j := &JsonWebKeyOutput{
		JsonWebKey:    jsonWebKey,
		AlgPrivateKey: algPrivateKey,
		AlgPublicKey:  algPublicKey,
	}

	jsonbuf, err := json.Marshal(jsonWebKey)
	if err != nil {
		return nil, err
	}

	j.JsonWebKeyString, err = getJsonOutputFromKey(j.JsonWebKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := j.JsonWebKey.PublicKey()
	if err != nil {
		return nil, err
	}
	j.JsonWebKeyPublic = publicKey

	j.JsonWebKeyPublicString, err = getJsonOutputFromKey(publicKey)
	if err != nil {
		return nil, err
	}

	j.Base64JsonWebKey = base64.StdEncoding.EncodeToString(jsonbuf)

	return j, nil
}

func getJsonOutputFromKey(key jwk.Key) (string, error) {
	jsonbuf, err := json.Marshal(key)
	if err != nil {
		return "", err
	}

	jsonWebKey := &bytes.Buffer{}
	if err := json.Indent(jsonWebKey, jsonbuf, "", "  "); err != nil {
		return "", err
	}

	return jsonWebKey.String(), nil
}

package jwk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

// JSONWebKeyOutput contains the output of a JSON Web Key creation operation.
type JSONWebKeyOutput struct {
	JSONWebKey             jwk.Key `json:"jwk,omitempty"`
	JSONWebKeyPublic       jwk.Key `json:"jwk_public_key,omitempty"`
	JSONWebKeyString       string  `json:"jwk_string,omitempty"`
	JSONWebKeyPublicString string  `json:"jwk_public_string,omitempty"`
	Base64JSONWebKey       string  `json:"base64_jwk,omitempty"`
	AlgPublicKey           string  `json:"alg_public_key,omitempty"`
	AlgPrivateKey          string  `json:"alg_private_key,omitempty"`
	Algorithm              string  `json:"algorithm,omitempty"`
}

// NewJSONWebKeyOutput creates a new JSONWebKeyOutput from the provided key and PEM-encoded keys.
func NewJSONWebKeyOutput(
	jsonWebKey jwk.Key,
	algPrivateKey string,
	algPublicKey string) (*JSONWebKeyOutput, error) {
	j := &JSONWebKeyOutput{
		JSONWebKey:    jsonWebKey,
		AlgPrivateKey: algPrivateKey,
		AlgPublicKey:  algPublicKey,
	}

	jsonBuf, err := json.Marshal(jsonWebKey)
	if err != nil {
		return nil, err
	}

	j.JSONWebKeyString, err = getJSONOutputFromKey(j.JSONWebKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := j.JSONWebKey.PublicKey()
	if err != nil {
		return nil, err
	}
	j.JSONWebKeyPublic = publicKey

	j.JSONWebKeyPublicString, err = getJSONOutputFromKey(publicKey)
	if err != nil {
		return nil, err
	}

	j.Base64JSONWebKey = base64.StdEncoding.EncodeToString(jsonBuf)

	return j, nil
}

func getJSONOutputFromKey(key jwk.Key) (string, error) {
	jsonBuf, err := json.Marshal(key)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := json.Indent(buf, jsonBuf, "", "  "); err != nil {
		return "", err
	}

	return buf.String(), nil
}

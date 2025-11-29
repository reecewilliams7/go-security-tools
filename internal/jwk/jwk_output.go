package jwk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

// JWKOutput contains the output of a JSON Web Key creation operation.
type JWKOutput struct {
	JWK             jwk.Key `json:"jwk,omitempty"`
	JWKPublic       jwk.Key `json:"jwk_public_key,omitempty"`
	JWKString       string  `json:"jwk_string,omitempty"`
	JWKPublicString string  `json:"jwk_public_string,omitempty"`
	Base64JWK       string  `json:"base64_jwk,omitempty"`
	PEMPublicKey    string  `json:"pem_public_key,omitempty"`
	PEMPrivateKey   string  `json:"pem_private_key,omitempty"`
}

// NewJWKOutput creates a new JWKOutput from the provided key and PEM-encoded keys.
func NewJWKOutput(
	key jwk.Key) (*JWKOutput, error) {
	j := &JWKOutput{
		JWK: key,
	}

	jsonBuf, err := json.Marshal(key)
	if err != nil {
		return nil, err
	}

	j.JWKString, err = getJSONOutputFromKey(j.JWK)
	if err != nil {
		return nil, err
	}

	pemPrivateKey, err := jwk.EncodePEM(key)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwk.PublicKeyOf(key)
	if err != nil {
		return nil, err
	}

	pemPublicKey, err := jwk.EncodePEM(publicKey)
	if err != nil {
		return nil, err
	}

	j.JWKPublicString, err = getJSONOutputFromKey(publicKey)
	if err != nil {
		return nil, err
	}

	j.Base64JWK = base64.StdEncoding.EncodeToString(jsonBuf)
	j.PEMPrivateKey = string(pemPrivateKey)
	j.PEMPublicKey = string(pemPublicKey)

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

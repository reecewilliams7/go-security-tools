package jsonWebKeys

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type jsonWebKeyOutput struct {
	JsonWebKey             jwk.Key
	JsonWebKeyString       string
	JsonWebKeyPublicString string
	Base64JsonWebKey       string
	RsaPublicKey           string
	RsaPrivateKey          string
}

func NewJsonWebKeyOutput(
	jsonWebKey jwk.Key,
	rsaPrivateKey string,
	rsaPublicKey string) (*jsonWebKeyOutput, error) {
	j := &jsonWebKeyOutput{
		JsonWebKey:    jsonWebKey,
		RsaPrivateKey: rsaPrivateKey,
		RsaPublicKey:  rsaPublicKey,
	}

	jsonbuf, err := json.Marshal(jsonWebKey)
	if err != nil {
		return nil, err
	}

	jwkPrivateJson := &bytes.Buffer{}
	if err := json.Indent(jwkPrivateJson, jsonbuf, "", "  "); err != nil {
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

	jwkPrivateJson := &bytes.Buffer{}
	if err := json.Indent(jwkPrivateJson, jsonbuf, "", "  "); err != nil {
		return "", err
	}

	return jwkPrivateJson.String(), nil
}

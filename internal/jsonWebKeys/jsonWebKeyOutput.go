package jsonWebKeys

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type jsonWebKeyOutput struct {
	JsonWebKey       *jwk.Key
	JsonWebKeyString string
	Base64JsonWebKey string
	RsaPublicKey     string
	RsaPrivateKey    string
}

func NewJsonWebKeyOutput(jsonWebKey *jwk.Key, rsaPrivateKey string, rsaPublicKey string) (*jsonWebKeyOutput, error) {
	j := &jsonWebKeyOutput{JsonWebKey: jsonWebKey, RsaPrivateKey: rsaPrivateKey, RsaPublicKey: rsaPublicKey}

	jsonbuf, err := json.Marshal(jsonWebKey)
	if err != nil {
		return nil, err
	}

	jwkJson := &bytes.Buffer{}
	if err := json.Indent(jwkJson, jsonbuf, "", "  "); err != nil {
		return nil, err
	}

	j.JsonWebKeyString = jwkJson.String()

	j.Base64JsonWebKey = base64.StdEncoding.EncodeToString(jsonbuf)

	return j, nil
}

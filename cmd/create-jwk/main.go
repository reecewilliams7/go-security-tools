package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	//jose "gopkg.in/square/go-jose.v2"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func main() {
	// rsaKey, _ := rsa.GenerateKey(rand.Reader, 2048)           // XXX Check err
	// serialNumber, _ := rand.Int(rand.Reader, big.NewInt(100)) // XXX Check err

	// template := x509.Certificate{
	// 	SerialNumber: serialNumber,
	// 	Subject: pkix.Name{
	// 		Organization: []string{"Example Co"},
	// 	},
	// 	NotBefore:             time.Now(),
	// 	NotAfter:              time.Now().Add(2 * time.Hour),
	// 	KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	// 	ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	// 	BasicConstraintsValid: true,
	// }

	// derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &rsaKey.PublicKey, rsaKey) // XXX Check err
	// certificate, _ := x509.ParseCertificate(derBytes)

	// jwk := jose.JSONWebKey{
	// 	Certificates: []*x509.Certificate{certificate},
	// 	Key:          &rsaKey.PublicKey,
	// 	KeyID:        "someKeyId",
	// 	Use:          "sig",
	// }

	// jwkJsonBytes, _ := jwk.MarshalJSON()
	// jwkJson := string(jwkJsonBytes)

	// log.Println(jwkJson)

	var rawKey interface{}
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf(`failed to generate rsa private key: %w`, err)
	}
	rawKey = rsaKey

	key, err := jwk.FromRaw(rawKey)

	if err != nil {
		log.Fatalf(`failed to create new JWK from raw key: %w`, err)
	}
}

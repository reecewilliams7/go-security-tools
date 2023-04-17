package main

import (
	"log"

	"github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

func main() {
	jc := jsonWebKeys.NewJsonWebKeyCreator()
	o, err := jc.Create()
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Println(o.JsonWebKeyString)
	log.Println("---------------------------")
	log.Println(o.Base64JsonWebKey)
	log.Println("---------------------------")
	log.Println(o.RsaPrivateKey)
	log.Println("---------------------------")
	log.Println(o.RsaPublicKey)
}

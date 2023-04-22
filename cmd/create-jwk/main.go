package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/reecewilliams7/go-security-tools/internal/jsonWebKeys"
)

func main() {
	log.SetPrefix("")

	var outputBase64Ptr = flag.Bool("output-base64", false, "")
	var outputRsaKeysPtr = flag.Bool("output-rsa-keys", false, "")

	flag.Parse()

	outputBase64 := *outputBase64Ptr
	outputRsaKeys := *outputRsaKeysPtr

	jc := jsonWebKeys.NewJsonWebKeyCreator()
	o, err := jc.Create()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(o.JsonWebKeyString)
	if outputBase64 {
		fmt.Println("")
		fmt.Println("---------------------------")
		fmt.Println("")
		fmt.Println(o.Base64JsonWebKey)
		fmt.Println("")
	}
	if outputRsaKeys {
		fmt.Println("---------------------------")
		fmt.Println("")
		fmt.Println(o.RsaPrivateKey)
		fmt.Println("")
		fmt.Println("---------------------------")
		fmt.Println("")
		fmt.Println(o.RsaPublicKey)
		fmt.Println("")
	}
}

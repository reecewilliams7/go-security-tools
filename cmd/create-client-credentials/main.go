package main

import (
	"fmt"

	"github.com/reecewilliams7/go-security-tools/internal/clientCredentials"
)

func main() {
	cc := clientCredentials.NewClientCredentialsCreator()
	o := cc.Create()

	fmt.Printf("ClientId: %s \n", o.ClientId)
	fmt.Printf("ClientSecret %s \n", o.ClientSecret)
}

package jwk

import "fmt"

const (
	asteriskLine = "**********************************************************"
	dashLine     = "---------------------------"
	emptyLine    = ""
)

type FmtJWKOutputWriter struct {
	outputBase64  bool
	outputPemKeys bool
}

func NewFmtJWKOutputWriter(outputBase64 bool, outputPemKeys bool) *FmtJWKOutputWriter {
	return &FmtJWKOutputWriter{
		outputBase64:  outputBase64,
		outputPemKeys: outputPemKeys,
	}
}

func (w *FmtJWKOutputWriter) Write(o *JWKOutput, i int) error {
	fmt.Println(asteriskLine)
	fmt.Printf("Writing JWK %d\n", i)
	fmt.Println(asteriskLine)
	fmt.Println("JWK Private Key:")
	fmt.Println(o.JWKString)
	fmt.Println(emptyLine)
	fmt.Println(dashLine)
	fmt.Println(emptyLine)
	fmt.Println("JWK Public Key:")
	fmt.Println(o.JWKPublicString)

	if w.outputBase64 {
		fmt.Println(emptyLine)
		fmt.Println(dashLine)
		fmt.Println(emptyLine)
		fmt.Println("Base64 Encoded JWK Private Key:")
		fmt.Println(o.Base64JWK)
		fmt.Println(emptyLine)
	}
	if w.outputPemKeys {
		fmt.Println(dashLine)
		fmt.Println(emptyLine)
		fmt.Printf("PEM Encoded %s Private Key:\n", o.JWK.KeyType())
		fmt.Println(o.PEMPrivateKey)
		fmt.Println(emptyLine)
		fmt.Println(dashLine)
		fmt.Println(emptyLine)
		fmt.Printf("PEM Encoded %s Public Key:\n", o.JWK.KeyType())
		fmt.Println(o.PEMPublicKey)
		fmt.Println(emptyLine)
	}
	fmt.Println(asteriskLine)

	return nil
}

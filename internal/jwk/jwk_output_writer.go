package jwk

type JWKOutputWriter interface {
	Write(output *JWKOutput, i int) error
}

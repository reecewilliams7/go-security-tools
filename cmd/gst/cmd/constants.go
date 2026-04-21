package cmd

const (
	LogLevelFlag = "log-level"

	// JWK
	OutputBase64Flag  = "output-base64"
	OutputPemKeysFlag = "output-pem-keys"
	OutputJWKSFlag    = "output-jwks"
	OutputPathFlag    = "output-path"
	OutputFileFlag    = "output-file"
	CountFlag         = "count"
	KeyTypeFlag       = "kty"

	// JWK Algorithm Types
	JwkAlgorithmRsa2048   = "RSA-2048"
	JwkAlgorithmRsa4096   = "RSA-4096"
	JwkAlgorithmEcdsaP256 = "ECDSA-P256"
	JwkAlgorithmEcdsaP384 = "ECDSA-P384"
	JwkAlgorithmEcdsaP521 = "ECDSA-P521"
	JwkAlgorithmOkpEd25519 = "OKP-Ed25519"
	JwkAlgorithmOkpX25519  = "OKP-X25519"
	JwkAlgorithmHs256      = "HS256"
	JwkAlgorithmHs384      = "HS384"
	JwkAlgorithmHs512      = "HS512"

	OutputFileDefault = "create-jwk"

	// Client Credentials
	ClientIdTypeFlag      = "client-id-type"
	ClientSecretTypeFlag  = "client-secret-type"
	SecretLengthFlag      = "secret-length"
	SecretEncodingFlag    = "secret-encoding"

	// Client ID Types
	ClientIdTypeUUIDv7  = "uuidv7"
	ClientIdTypeShort   = "short-uuid"
	ClientIdTypeNanoid  = "nanoid"

	// Client Secret Types
	ClientSecretTypeCryptoRand = "crypto-rand"

	// Secret Encodings
	SecretEncodingBase64    = "base64"
	SecretEncodingBase64URL = "base64url"
	SecretEncodingHex       = "hex"

	// PKCE
	PKCEMethodFlag = "method"
	PKCEMethodS256 = "S256"

	// JWT
	JWTTokenFlag = "token"
)

package cmd

const (
	LogLevelFlag = "log-level"

	// JWK
	OutputBase64Flag  = "output-base64"
	OutputPemKeysFlag = "output-pem-keys"
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

	OutputFileDefault = "create-jwk"

	// Client Credentials
	ClientIdTypeFlag     = "client-id-type"
	ClientSecretTypeFlag = "client-secret-type"

	// Client ID Types
	ClientIdTypeUUIDv7 = "uuidv7"
	ClientIdTypeShort  = "short-uuid"

	// Client Secret Types
	ClientSecretTypeCryptoRand = "crypto-rand"
)

package cmd

const (
	LogLevelFlagName = "log-level"

	OutputBase64FlagName   = "output-base64"
	OutputRsaKeysFlagName  = "output-rsa-keys"
	OutputPathFlagName     = "output-path"
	OutputFileNameFlagName = "output-file-name"
	CountFlagName          = "count"
	OutputFileDefaultName  = "create-jwk"

	// Client Credentials
	ClientIdTypeFlag     = "client-id-type"
	ClientSecretTypeFlag = "client-secret-type"

	// Client ID Types
	ClientIdTypeUUIDv7 = "uuidv7"
	ClientIdTypeShort  = "short-uuid"

	// Client Secret Types
	ClientSecretTypeCryptoRand = "crypto-rand"
)

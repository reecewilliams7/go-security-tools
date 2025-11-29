package cmd

import (
	"os"
	"testing"
)

func TestBuildClientCredentialsCreator_UUIDv7(t *testing.T) {
	ccc, err := buildClientCredentialsCreator(ClientIdTypeUUIDv7, ClientSecretTypeCryptoRand)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ccc == nil {
		t.Fatal("expected non-nil ClientCredentialsCreator")
	}

	// Verify it works
	cc, err := ccc.CreateClientCredentials()
	if err != nil {
		t.Fatalf("unexpected error creating credentials: %v", err)
	}
	if cc.ClientID() == "" {
		t.Error("expected non-empty ClientID")
	}
	if cc.ClientSecret() == "" {
		t.Error("expected non-empty ClientSecret")
	}
}

func TestBuildClientCredentialsCreator_ShortUUID(t *testing.T) {
	ccc, err := buildClientCredentialsCreator(ClientIdTypeShort, ClientSecretTypeCryptoRand)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ccc == nil {
		t.Fatal("expected non-nil ClientCredentialsCreator")
	}
}

func TestBuildClientCredentialsCreator_UnknownIDType(t *testing.T) {
	_, err := buildClientCredentialsCreator("unknown", ClientSecretTypeCryptoRand)
	if err == nil {
		t.Fatal("expected error for unknown ID type")
	}
}

func TestBuildClientCredentialsCreator_UnknownSecretType(t *testing.T) {
	_, err := buildClientCredentialsCreator(ClientIdTypeUUIDv7, "unknown")
	if err == nil {
		t.Fatal("expected error for unknown secret type")
	}
}

func TestBuildJWKCreator_RSA2048(t *testing.T) {
	creator, err := buildJWKCreator(JwkAlgorithmRsa2048)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if creator == nil {
		t.Fatal("expected non-nil JWKCreator")
	}
}

func TestBuildJWKCreator_RSA4096(t *testing.T) {
	creator, err := buildJWKCreator(JwkAlgorithmRsa4096)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if creator == nil {
		t.Fatal("expected non-nil JWKCreator")
	}
}

func TestBuildJWKCreator_ECDSA(t *testing.T) {
	tests := []string{
		JwkAlgorithmEcdsaP256,
		JwkAlgorithmEcdsaP384,
		JwkAlgorithmEcdsaP521,
	}

	for _, alg := range tests {
		t.Run(alg, func(t *testing.T) {
			creator, err := buildJWKCreator(alg)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if creator == nil {
				t.Fatal("expected non-nil JWKCreator")
			}
		})
	}
}

func TestBuildJWKCreator_Unknown(t *testing.T) {
	_, err := buildJWKCreator("unknown")
	if err == nil {
		t.Fatal("expected error for unknown algorithm")
	}
}

func TestBuildJWKWriter_Console(t *testing.T) {
	writer, err := buildJWKWriter("", "test", false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if writer == nil {
		t.Fatal("expected non-nil JWKOutputWriter")
	}
}

func TestBuildJWKWriter_File(t *testing.T) {
	tempDir := t.TempDir()

	writer, err := buildJWKWriter(tempDir, "test", false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if writer == nil {
		t.Fatal("expected non-nil JWKOutputWriter")
	}
}

func TestBuildJWKWriter_NonExistentPath(t *testing.T) {
	_, err := buildJWKWriter("/non/existent/path", "test", false, false)
	if err == nil {
		t.Fatal("expected error for non-existent path")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected os.IsNotExist error, got %v", err)
	}
}

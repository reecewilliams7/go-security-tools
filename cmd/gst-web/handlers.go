package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/reecewilliams7/go-security-tools/clientcredentials"
	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
	"github.com/reecewilliams7/go-security-tools/jwk"
)

const (
	algorithmRSA2048   = "RSA-2048"
	algorithmRSA4096   = "RSA-4096"
	algorithmECDSAP256 = "ECDSA-P256"
	algorithmECDSAP384 = "ECDSA-P384"
	algorithmECDSAP521 = "ECDSA-P521"

	idTypeUUIDv7    = "uuidv7"
	idTypeShortUUID = "short-uuid"

	secretTypeCryptoRand = "crypto-rand"

	maxJWKCount = 10
	maxCCCount  = 20
)

var (
	jwkKeyTypes       = []string{algorithmRSA2048, algorithmRSA4096, algorithmECDSAP256, algorithmECDSAP384, algorithmECDSAP521}
	ccClientIDTypes   = []string{idTypeUUIDv7, idTypeShortUUID}
	ccSecretTypes     = []string{secretTypeCryptoRand}
	validJWKKeyTypes  = map[string]bool{algorithmRSA2048: true, algorithmRSA4096: true, algorithmECDSAP256: true, algorithmECDSAP384: true, algorithmECDSAP521: true}
	validCCIDTypes    = map[string]bool{idTypeUUIDv7: true, idTypeShortUUID: true}
	validSecretTypes  = map[string]bool{secretTypeCryptoRand: true}
)

// JWKPageData holds template data for the JWK creation page.
type JWKPageData struct {
	Title         string
	KeyTypes      []string
	KeyType       string
	Count         int
	OutputBase64  bool
	OutputPemKeys bool
	Results       []*internaljwk.JWKOutput
	Error         string
	Submitted     bool
}

// CCPageData holds template data for the client credentials creation page.
type CCPageData struct {
	Title            string
	ClientIDTypes    []string
	SecretTypes      []string
	ClientIDType     string
	SecretType       string
	Count            int
	Results          []*clientcredentials.ClientCredentials
	Error            string
	Submitted        bool
}

func (s *server) render(w http.ResponseWriter, name string, data any) {
	tmpl, ok := s.templates[name]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) handleIndex(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "index", map[string]string{"Title": "Home"})
}

func (s *server) handleJWKGet(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "jwk", JWKPageData{
		Title:    "Create JWK",
		KeyTypes: jwkKeyTypes,
		KeyType:  algorithmRSA2048,
		Count:    1,
	})
}

func (s *server) handleJWKPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	keyType := r.FormValue("keyType")
	outputBase64 := r.FormValue("outputBase64") == "on"
	outputPemKeys := r.FormValue("outputPemKeys") == "on"

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil || count < 1 || count > maxJWKCount {
		count = 1
	}

	data := JWKPageData{
		Title:         "Create JWK",
		KeyTypes:      jwkKeyTypes,
		KeyType:       keyType,
		Count:         count,
		OutputBase64:  outputBase64,
		OutputPemKeys: outputPemKeys,
		Submitted:     true,
	}

	if !validJWKKeyTypes[keyType] {
		data.Error = fmt.Sprintf("invalid key type: %q", keyType)
		s.render(w, "jwk", data)
		return
	}

	creator, err := buildJWKCreator(keyType)
	if err != nil {
		data.Error = err.Error()
		s.render(w, "jwk", data)
		return
	}

	for i := range count {
		o, err := creator.Create()
		if err != nil {
			data.Error = fmt.Sprintf("failed to create JWK %d: %v", i+1, err)
			s.render(w, "jwk", data)
			return
		}
		data.Results = append(data.Results, o)
	}

	s.render(w, "jwk", data)
}

func (s *server) handleCCGet(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "cc", CCPageData{
		Title:         "Create Client Credentials",
		ClientIDTypes: ccClientIDTypes,
		SecretTypes:   ccSecretTypes,
		ClientIDType:  idTypeUUIDv7,
		SecretType:    secretTypeCryptoRand,
		Count:         1,
	})
}

func (s *server) handleCCPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	clientIDType := r.FormValue("clientIdType")
	secretType := r.FormValue("secretType")

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil || count < 1 || count > maxCCCount {
		count = 1
	}

	data := CCPageData{
		Title:         "Create Client Credentials",
		ClientIDTypes: ccClientIDTypes,
		SecretTypes:   ccSecretTypes,
		ClientIDType:  clientIDType,
		SecretType:    secretType,
		Count:         count,
		Submitted:     true,
	}

	if !validCCIDTypes[clientIDType] {
		data.Error = fmt.Sprintf("invalid client ID type: %q", clientIDType)
		s.render(w, "cc", data)
		return
	}

	if !validSecretTypes[secretType] {
		data.Error = fmt.Sprintf("invalid secret type: %q", secretType)
		s.render(w, "cc", data)
		return
	}

	creator, err := buildCCCreator(clientIDType, secretType)
	if err != nil {
		data.Error = err.Error()
		s.render(w, "cc", data)
		return
	}

	for i := range count {
		cc, err := creator.CreateClientCredentials()
		if err != nil {
			data.Error = fmt.Sprintf("failed to create credentials %d: %v", i+1, err)
			s.render(w, "cc", data)
			return
		}
		data.Results = append(data.Results, cc)
	}

	s.render(w, "cc", data)
}

func buildJWKCreator(algorithm string) (jwk.JWKCreator, error) {
	switch algorithm {
	case algorithmRSA2048:
		return jwk.NewRSAJSONWebKeyCreator(2048), nil
	case algorithmRSA4096:
		return jwk.NewRSAJSONWebKeyCreator(4096), nil
	case algorithmECDSAP256:
		return jwk.NewECDSAJWKCreator("P256"), nil
	case algorithmECDSAP384:
		return jwk.NewECDSAJWKCreator("P384"), nil
	case algorithmECDSAP521:
		return jwk.NewECDSAJWKCreator("P521"), nil
	default:
		return nil, fmt.Errorf("unknown JWK algorithm: %s", algorithm)
	}
}

func buildCCCreator(clientIDType, secretType string) (*clientcredentials.ClientCredentialsCreator, error) {
	var idCreator clientcredentials.ClientIDCreator
	switch clientIDType {
	case idTypeUUIDv7:
		idCreator = clientcredentials.NewUUIDv7ClientIDCreator()
	case idTypeShortUUID:
		idCreator = clientcredentials.NewShortUUIDClientIDCreator()
	default:
		return nil, fmt.Errorf("unknown client ID type: %s", clientIDType)
	}

	var secretCreator clientcredentials.ClientSecretCreator
	switch secretType {
	case secretTypeCryptoRand:
		secretCreator = clientcredentials.NewCryptoRandClientSecretCreator()
	default:
		return nil, fmt.Errorf("unknown secret type: %s", secretType)
	}

	return clientcredentials.NewClientCredentialsCreator(idCreator, secretCreator), nil
}

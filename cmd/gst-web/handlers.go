package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/reecewilliams7/go-security-tools/clientcredentials"
	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
	internaljwt "github.com/reecewilliams7/go-security-tools/internal/jwt"
	"github.com/reecewilliams7/go-security-tools/jwk"
	"github.com/reecewilliams7/go-security-tools/pkce"
)

const (
	algorithmRSA2048    = "RSA-2048"
	algorithmRSA4096    = "RSA-4096"
	algorithmECDSAP256  = "ECDSA-P256"
	algorithmECDSAP384  = "ECDSA-P384"
	algorithmECDSAP521  = "ECDSA-P521"
	algorithmOKPEd25519 = "OKP-Ed25519"
	algorithmOKPX25519  = "OKP-X25519"
	algorithmHS256      = "HS256"
	algorithmHS384      = "HS384"
	algorithmHS512      = "HS512"

	idTypeUUIDv7    = "uuidv7"
	idTypeShortUUID = "short-uuid"
	idTypeNanoid    = "nanoid"

	secretTypeCryptoRand = "crypto-rand"

	maxJWKCount = 10
	maxCCCount  = 20
	maxPKCECount = 20
)

var (
	jwkKeyTypes = []string{
		algorithmRSA2048, algorithmRSA4096,
		algorithmECDSAP256, algorithmECDSAP384, algorithmECDSAP521,
		algorithmOKPEd25519, algorithmOKPX25519,
		algorithmHS256, algorithmHS384, algorithmHS512,
	}
	ccClientIDTypes  = []string{idTypeUUIDv7, idTypeShortUUID, idTypeNanoid}
	ccSecretTypes    = []string{secretTypeCryptoRand}
	secretEncodings  = []string{"base64", "base64url", "hex"}
	validJWKKeyTypes = map[string]bool{
		algorithmRSA2048: true, algorithmRSA4096: true,
		algorithmECDSAP256: true, algorithmECDSAP384: true, algorithmECDSAP521: true,
		algorithmOKPEd25519: true, algorithmOKPX25519: true,
		algorithmHS256: true, algorithmHS384: true, algorithmHS512: true,
	}
	validCCIDTypes       = map[string]bool{idTypeUUIDv7: true, idTypeShortUUID: true, idTypeNanoid: true}
	validSecretTypes     = map[string]bool{secretTypeCryptoRand: true}
	validSecretEncodings = map[string]bool{"base64": true, "base64url": true, "hex": true}
)

// JWKPageData holds template data for the JWK creation page.
type JWKPageData struct {
	Title         string
	KeyTypes      []string
	KeyType       string
	Count         int
	OutputBase64  bool
	OutputPemKeys bool
	OutputJWKS    bool
	JWKSString    string
	Results       []*internaljwk.JWKOutput
	Error         string
	Submitted     bool
}

// CCPageData holds template data for the client credentials creation page.
type CCPageData struct {
	Title           string
	ClientIDTypes   []string
	SecretTypes     []string
	SecretEncodings []string
	ClientIDType    string
	SecretType      string
	SecretLength    int
	SecretEncoding  string
	Count           int
	Results         []*clientcredentials.ClientCredentials
	Error           string
	Submitted       bool
}

// PKCEPageData holds template data for the PKCE creation page.
type PKCEPageData struct {
	Title     string
	Count     int
	Results   []*pkce.PKCEPair
	Error     string
	Submitted bool
}

// JWTPageData holds template data for the JWT decode page.
type JWTPageData struct {
	Title       string
	Token       string
	HeaderJSON  string
	PayloadJSON string
	IsExpired   bool
	ExpiresAt   string
	NoExpiry    bool
	Error       string
	Submitted   bool
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
	outputJWKS := r.FormValue("outputJWKS") == "on"

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
		OutputJWKS:    outputJWKS,
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

	if outputJWKS {
		var rawKeys []json.RawMessage
		for _, o := range data.Results {
			if o.JWKPublicString != "" {
				rawKeys = append(rawKeys, json.RawMessage(o.JWKPublicString))
			} else if o.JWKString != "" {
				rawKeys = append(rawKeys, json.RawMessage(o.JWKString))
			}
		}
		jwksBytes, _ := json.MarshalIndent(map[string]any{"keys": rawKeys}, "", "  ")
		data.JWKSString = string(jwksBytes)
	}

	s.render(w, "jwk", data)
}

func (s *server) handleCCGet(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "cc", CCPageData{
		Title:           "Create Client Credentials",
		ClientIDTypes:   ccClientIDTypes,
		SecretTypes:     ccSecretTypes,
		SecretEncodings: secretEncodings,
		ClientIDType:    idTypeUUIDv7,
		SecretType:      secretTypeCryptoRand,
		SecretLength:    32,
		SecretEncoding:  "base64",
		Count:           1,
	})
}

func (s *server) handleCCPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	clientIDType := r.FormValue("clientIdType")
	secretType := r.FormValue("secretType")
	secretEncoding := r.FormValue("secretEncoding")
	if !validSecretEncodings[secretEncoding] {
		secretEncoding = "base64"
	}

	secretLength, err := strconv.Atoi(r.FormValue("secretLength"))
	if err != nil || secretLength < 16 || secretLength > 64 {
		secretLength = 32
	}

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil || count < 1 || count > maxCCCount {
		count = 1
	}

	data := CCPageData{
		Title:           "Create Client Credentials",
		ClientIDTypes:   ccClientIDTypes,
		SecretTypes:     ccSecretTypes,
		SecretEncodings: secretEncodings,
		ClientIDType:    clientIDType,
		SecretType:      secretType,
		SecretLength:    secretLength,
		SecretEncoding:  secretEncoding,
		Count:           count,
		Submitted:       true,
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

	creator, err := buildCCCreator(clientIDType, secretType, secretLength, secretEncoding)
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
	case algorithmOKPEd25519:
		return jwk.NewOKPJWKCreator("Ed25519"), nil
	case algorithmOKPX25519:
		return jwk.NewOKPJWKCreator("X25519"), nil
	case algorithmHS256:
		return jwk.NewHMACJWKCreator("HS256"), nil
	case algorithmHS384:
		return jwk.NewHMACJWKCreator("HS384"), nil
	case algorithmHS512:
		return jwk.NewHMACJWKCreator("HS512"), nil
	default:
		return nil, fmt.Errorf("unknown JWK algorithm: %s", algorithm)
	}
}

func buildCCCreator(clientIDType, secretType string, secretLength int, secretEncoding string) (*clientcredentials.ClientCredentialsCreator, error) {
	var idCreator clientcredentials.ClientIDCreator
	switch clientIDType {
	case idTypeUUIDv7:
		idCreator = clientcredentials.NewUUIDv7ClientIDCreator()
	case idTypeShortUUID:
		idCreator = clientcredentials.NewShortUUIDClientIDCreator()
	case idTypeNanoid:
		idCreator = clientcredentials.NewNanoidClientIDCreator()
	default:
		return nil, fmt.Errorf("unknown client ID type: %s", clientIDType)
	}

	var secretCreator clientcredentials.ClientSecretCreator
	switch secretType {
	case secretTypeCryptoRand:
		secretCreator = clientcredentials.NewCryptoRandClientSecretCreatorWithConfig(secretLength, secretEncoding)
	default:
		return nil, fmt.Errorf("unknown secret type: %s", secretType)
	}

	return clientcredentials.NewClientCredentialsCreator(idCreator, secretCreator), nil
}

func (s *server) handlePKCEGet(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "pkce", PKCEPageData{
		Title: "Create PKCE",
		Count: 1,
	})
}

func (s *server) handlePKCEPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil || count < 1 || count > maxPKCECount {
		count = 1
	}

	data := PKCEPageData{Title: "Create PKCE", Count: count, Submitted: true}
	creator := pkce.NewS256Creator()
	for i := range count {
		pair, err := creator.Create()
		if err != nil {
			data.Error = fmt.Sprintf("failed to create PKCE pair %d: %v", i+1, err)
			s.render(w, "pkce", data)
			return
		}
		data.Results = append(data.Results, pair)
	}
	s.render(w, "pkce", data)
}

func (s *server) handleJWTGet(w http.ResponseWriter, _ *http.Request) {
	s.render(w, "jwt", JWTPageData{Title: "Decode JWT"})
}

func (s *server) handleJWTPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	token := strings.TrimSpace(r.FormValue("token"))
	data := JWTPageData{Title: "Decode JWT", Token: token, Submitted: true}

	if token == "" {
		data.Error = "JWT token is required"
		s.render(w, "jwt", data)
		return
	}

	decoded, err := internaljwt.Decode(token)
	if err != nil {
		data.Error = fmt.Sprintf("failed to decode JWT: %v", err)
		s.render(w, "jwt", data)
		return
	}

	headerJSON, _ := json.MarshalIndent(decoded.Header, "", "  ")
	payloadJSON, _ := json.MarshalIndent(decoded.Payload, "", "  ")
	data.HeaderJSON = string(headerJSON)
	data.PayloadJSON = string(payloadJSON)
	data.IsExpired = decoded.IsExpired
	if decoded.ExpiresAt == nil {
		data.NoExpiry = true
	} else {
		data.ExpiresAt = decoded.ExpiresAt.UTC().Format("2006-01-02 15:04:05 UTC")
	}
	s.render(w, "jwt", data)
}

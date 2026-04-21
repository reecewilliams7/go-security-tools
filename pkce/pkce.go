package pkce

// PKCEPair holds a PKCE code_verifier and its corresponding code_challenge.
type PKCEPair struct {
	CodeVerifier  string
	CodeChallenge string
	Method        string
}

// Creator generates PKCE code verifier/challenge pairs.
type Creator interface {
	Create() (*PKCEPair, error)
}

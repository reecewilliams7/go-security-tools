package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// DecodeOutput holds the decoded parts of a JWT.
type DecodeOutput struct {
	Header    map[string]any
	Payload   map[string]any
	IsExpired bool
	ExpiresAt *time.Time
}

// Decode splits and base64url-decodes a JWT's header and payload without
// verifying the signature. It returns an error for malformed tokens.
func Decode(token string) (*DecodeOutput, error) {
	token = strings.TrimSpace(token)
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT: expected 3 dot-separated parts")
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid JWT header encoding: %w", err)
	}

	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid JWT payload encoding: %w", err)
	}

	var header map[string]any
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, fmt.Errorf("invalid JWT header JSON: %w", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, fmt.Errorf("invalid JWT payload JSON: %w", err)
	}

	out := &DecodeOutput{
		Header:  header,
		Payload: payload,
	}

	if expVal, ok := payload["exp"]; ok {
		if expNum, ok := expVal.(float64); ok {
			t := time.Unix(int64(expNum), 0).UTC()
			out.ExpiresAt = &t
			out.IsExpired = time.Now().After(t)
		}
	}

	return out, nil
}

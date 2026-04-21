package tui

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	internaljwt "github.com/reecewilliams7/go-security-tools/internal/jwt"
)

type jwtField int

const (
	jwtFieldToken  jwtField = iota // JWT token input
	jwtFieldSubmit                 // decode button
	jwtNumFields                   // sentinel
)

type jwtFormModel struct {
	focused    jwtField
	tokenInput textinput.Model
	decoding   bool
	width      int
	height     int
}

func newJWTFormModel() jwtFormModel {
	token := textinput.New()
	token.Placeholder = "Paste JWT token here…"
	token.CharLimit = 8192
	token.Width = 60
	token.Focus()

	return jwtFormModel{
		focused:    jwtFieldToken,
		tokenInput: token,
	}
}

func (m jwtFormModel) init() tea.Cmd { return nil }

func (m jwtFormModel) update(msg tea.Msg) (jwtFormModel, tea.Cmd) {
	if m.decoding {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, navigate(screenMenu)
		case "tab", "down":
			m = m.moveFocus(1)
			return m, nil
		case "shift+tab", "up":
			m = m.moveFocus(-1)
			return m, nil
		case "enter":
			if m.focused == jwtFieldSubmit {
				m.decoding = true
				return m, m.submitCmd()
			}
			m = m.moveFocus(1)
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.focused == jwtFieldToken {
		m.tokenInput, cmd = m.tokenInput.Update(msg)
	}
	return m, cmd
}

func (m jwtFormModel) moveFocus(delta int) jwtFormModel {
	next := int(m.focused) + delta
	if next < 0 {
		next = int(jwtNumFields) - 1
	} else if next >= int(jwtNumFields) {
		next = 0
	}
	m.focused = jwtField(next)

	if m.focused == jwtFieldToken {
		m.tokenInput.Focus()
	} else {
		m.tokenInput.Blur()
	}
	return m
}

func (m jwtFormModel) submitCmd() tea.Cmd {
	token := strings.TrimSpace(m.tokenInput.Value())

	return func() tea.Msg {
		if token == "" {
			return navigateMsg{to: screenResult, result: "Error: no JWT token provided"}
		}

		decoded, err := internaljwt.Decode(token)
		if err != nil {
			return navigateMsg{to: screenResult, result: fmt.Sprintf("Error decoding JWT: %v", err)}
		}

		headerJSON, _ := json.MarshalIndent(decoded.Header, "", "  ")
		payloadJSON, _ := json.MarshalIndent(decoded.Payload, "", "  ")

		sep := strings.Repeat("─", 60)
		var sb strings.Builder
		sb.WriteString(sep + "\n")
		sb.WriteString("Header\n")
		sb.WriteString(sep + "\n")
		sb.WriteString(string(headerJSON))
		sb.WriteString("\n\n")
		sb.WriteString(sep + "\n")
		sb.WriteString("Payload\n")
		sb.WriteString(sep + "\n")
		sb.WriteString(string(payloadJSON))
		sb.WriteString("\n\n")
		sb.WriteString(sep + "\n")
		if decoded.IsExpired {
			sb.WriteString("Status: EXPIRED")
		} else if decoded.ExpiresAt == nil {
			sb.WriteString("Status: No expiry claim (exp)")
		} else {
			sb.WriteString(fmt.Sprintf("Status: Valid — expires at %s", decoded.ExpiresAt.UTC().Format("2006-01-02 15:04:05 UTC")))
		}
		sb.WriteString("\n")
		return navigateMsg{to: screenResult, result: sb.String()}
	}
}

func (m jwtFormModel) view() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Decode JWT"))
	sb.WriteString("\n\n")

	if m.decoding {
		sb.WriteString(selectedStyle.Render("Decoding token…"))
		sb.WriteString("\n")
		return sb.String()
	}

	focused := m.focused == jwtFieldToken
	var lbl string
	if focused {
		lbl = selectedStyle.Render(fmt.Sprintf("▶  %-20s", "JWT Token"))
	} else {
		lbl = normalStyle.Render(fmt.Sprintf("   %-20s", "JWT Token"))
	}
	sb.WriteString(lbl + "  " + m.tokenInput.View() + "\n")
	sb.WriteString("\n")

	if m.focused == jwtFieldSubmit {
		sb.WriteString(selectedStyle.Render("▶  [ Decode ]"))
	} else {
		sb.WriteString(normalStyle.Render("   [ Decode ]"))
	}
	sb.WriteString("\n")

	sb.WriteString(helpStyle.Render("\ntab/↓ next  shift+tab/↑ prev  enter confirm  esc back"))
	return sb.String()
}

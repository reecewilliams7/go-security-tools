package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	internaljwk "github.com/reecewilliams7/go-security-tools/internal/jwk"
	"github.com/reecewilliams7/go-security-tools/jwk"
)

type jwkField int

const (
	jwkFieldKeyType    jwkField = iota // key-type selector
	jwkFieldCount                      // count text input
	jwkFieldBase64                     // base64 toggle
	jwkFieldPEMKeys                    // PEM keys toggle
	jwkFieldOutputPath                 // output-path text input
	jwkFieldOutputFile                 // output-file text input
	jwkFieldSubmit                     // generate button
	jwkNumFields                       // sentinel
)

var keyTypeOptions = []string{
	"RSA-2048",
	"RSA-4096",
	"ECDSA-P256",
	"ECDSA-P384",
	"ECDSA-P521",
}

type jwkFormModel struct {
	focused    jwkField
	keyTypeIdx int
	countInput textinput.Model
	base64     bool
	pemKeys    bool
	outputPath textinput.Model
	outputFile textinput.Model
	generating bool
	width      int
	height     int
}

func newJWKFormModel() jwkFormModel {
	count := textinput.New()
	count.Placeholder = "1"
	count.CharLimit = 2
	count.Width = 5

	path := textinput.New()
	path.Placeholder = "(leave empty for console output)"
	path.CharLimit = 256
	path.Width = 45

	file := textinput.New()
	file.Placeholder = "create-jwk"
	file.CharLimit = 64
	file.Width = 30
	file.SetValue("create-jwk")

	m := jwkFormModel{
		focused:    jwkFieldKeyType,
		countInput: count,
		outputPath: path,
		outputFile: file,
	}
	return m
}

func (m jwkFormModel) init() tea.Cmd { return nil }

func (m jwkFormModel) update(msg tea.Msg) (jwkFormModel, tea.Cmd) {
	if m.generating {
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
			if m.focused == jwkFieldSubmit {
				m.generating = true
				return m, m.submitCmd()
			}
			m = m.moveFocus(1)
			return m, nil
		case "left":
			if m.focused == jwkFieldKeyType && m.keyTypeIdx > 0 {
				m.keyTypeIdx--
			}
			return m, nil
		case "right":
			if m.focused == jwkFieldKeyType && m.keyTypeIdx < len(keyTypeOptions)-1 {
				m.keyTypeIdx++
			}
			return m, nil
		case " ":
			switch m.focused {
			case jwkFieldBase64:
				m.base64 = !m.base64
			case jwkFieldPEMKeys:
				m.pemKeys = !m.pemKeys
			}
			return m, nil
		}
	}

	// Forward remaining messages to the active text input.
	var cmd tea.Cmd
	switch m.focused {
	case jwkFieldCount:
		m.countInput, cmd = m.countInput.Update(msg)
	case jwkFieldOutputPath:
		m.outputPath, cmd = m.outputPath.Update(msg)
	case jwkFieldOutputFile:
		m.outputFile, cmd = m.outputFile.Update(msg)
	}
	return m, cmd
}

func (m jwkFormModel) moveFocus(delta int) jwkFormModel {
	next := int(m.focused) + delta
	if next < 0 {
		next = int(jwkNumFields) - 1
	} else if next >= int(jwkNumFields) {
		next = 0
	}
	m.focused = jwkField(next)

	// Sync text-input focus state.
	m.countInput.Blur()
	m.outputPath.Blur()
	m.outputFile.Blur()
	switch m.focused {
	case jwkFieldCount:
		m.countInput.Focus()
	case jwkFieldOutputPath:
		m.outputPath.Focus()
	case jwkFieldOutputFile:
		m.outputFile.Focus()
	}
	return m
}

func (m jwkFormModel) submitCmd() tea.Cmd {
	// Capture values now; the closure runs in a goroutine.
	keyType := keyTypeOptions[m.keyTypeIdx]
	base64 := m.base64
	pemKeys := m.pemKeys
	outputPath := strings.TrimSpace(m.outputPath.Value())
	outputFile := strings.TrimSpace(m.outputFile.Value())
	if outputFile == "" {
		outputFile = "create-jwk"
	}
	countStr := m.countInput.Value()
	if countStr == "" {
		countStr = "1"
	}
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 1
	}
	if count > 10 {
		count = 10
	}

	return func() tea.Msg {
		var creator jwk.JWKCreator
		switch keyType {
		case "RSA-2048":
			creator = jwk.NewRSAJSONWebKeyCreator(2048)
		case "RSA-4096":
			creator = jwk.NewRSAJSONWebKeyCreator(4096)
		case "ECDSA-P256":
			creator = jwk.NewECDSAJWKCreator("P256")
		case "ECDSA-P384":
			creator = jwk.NewECDSAJWKCreator("P384")
		case "ECDSA-P521":
			creator = jwk.NewECDSAJWKCreator("P521")
		default:
			creator = jwk.NewRSAJSONWebKeyCreator(2048)
		}

		writeToFile := outputPath != ""
		var writer internaljwk.JWKOutputWriter
		if writeToFile {
			writer = internaljwk.NewFileJwkOutputWriter(outputPath, outputFile, base64, pemKeys)
		}

		var sb strings.Builder
		for i := 1; i <= count; i++ {
			o, err := creator.Create()
			if err != nil {
				return navigateMsg{to: screenResult, result: fmt.Sprintf("Error creating JWK: %v", err)}
			}

			if writeToFile {
				if err := writer.Write(o, i); err != nil {
					return navigateMsg{to: screenResult, result: fmt.Sprintf("Error writing JWK to file: %v", err)}
				}
				sb.WriteString(fmt.Sprintf("JWK %d — private key written to %s/%s-%d.jwk\n", i, outputPath, outputFile, i))
				sb.WriteString(fmt.Sprintf("JWK %d — public  key written to %s/%s-pub-%d.jwk\n", i, outputPath, outputFile, i))
				if base64 {
					sb.WriteString(fmt.Sprintf("JWK %d — base64      written to %s/%s-base64-%d.jwk\n", i, outputPath, outputFile, i))
				}
				if pemKeys {
					sb.WriteString(fmt.Sprintf("JWK %d — PEM public  written to %s/%s-%d.pub\n", i, outputPath, outputFile, i))
					sb.WriteString(fmt.Sprintf("JWK %d — PEM private written to %s/%s-%d.key\n", i, outputPath, outputFile, i))
				}
			} else {
				sep := strings.Repeat("─", 60)
				sb.WriteString(sep + "\n")
				sb.WriteString(fmt.Sprintf("JWK %d  (%s)\n", i, keyType))
				sb.WriteString(sep + "\n\n")
				sb.WriteString("Private Key:\n")
				sb.WriteString(o.JWKString)
				sb.WriteString("\n\nPublic Key:\n")
				sb.WriteString(o.JWKPublicString)
				sb.WriteString("\n")
				if base64 {
					sb.WriteString("\nBase64 Encoded Private Key:\n")
					sb.WriteString(o.Base64JWK)
					sb.WriteString("\n")
				}
				if pemKeys {
					sb.WriteString(fmt.Sprintf("\nPEM %s Private Key:\n", o.JWK.KeyType()))
					sb.WriteString(o.PEMPrivateKey)
					sb.WriteString(fmt.Sprintf("\nPEM %s Public Key:\n", o.JWK.KeyType()))
					sb.WriteString(o.PEMPublicKey)
				}
			}
		}
		return navigateMsg{to: screenResult, result: sb.String()}
	}
}

func (m jwkFormModel) view() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Create JWK"))
	sb.WriteString("\n\n")

	if m.generating {
		sb.WriteString(selectedStyle.Render("Generating keys, please wait…"))
		sb.WriteString("\n")
		return sb.String()
	}

	sb.WriteString(m.renderField(jwkFieldKeyType, "Key Type",
		m.renderSelector(keyTypeOptions, m.keyTypeIdx, m.focused == jwkFieldKeyType)))
	sb.WriteString(m.renderField(jwkFieldCount, "Count", m.countInput.View()))
	sb.WriteString(m.renderField(jwkFieldBase64, "Output Base64",
		m.renderToggle(m.base64, m.focused == jwkFieldBase64)))
	sb.WriteString(m.renderField(jwkFieldPEMKeys, "Output PEM Keys",
		m.renderToggle(m.pemKeys, m.focused == jwkFieldPEMKeys)))
	sb.WriteString(m.renderField(jwkFieldOutputPath, "Output Path", m.outputPath.View()))
	sb.WriteString(m.renderField(jwkFieldOutputFile, "Output File", m.outputFile.View()))
	sb.WriteString("\n")

	if m.focused == jwkFieldSubmit {
		sb.WriteString(selectedStyle.Render("▶  [ Generate ]"))
	} else {
		sb.WriteString(normalStyle.Render("   [ Generate ]"))
	}
	sb.WriteString("\n")

	sb.WriteString(helpStyle.Render("\ntab/↓ next  shift+tab/↑ prev  ←/→ change option  space toggle  enter confirm  esc back"))
	return sb.String()
}

func (m jwkFormModel) renderField(f jwkField, label, value string) string {
	focused := m.focused == f
	var lbl string
	if focused {
		lbl = focusedLabelStyle.Render(fmt.Sprintf("▶  %-18s", label+":"))
	} else {
		lbl = unfocusedLabelStyle.Render(fmt.Sprintf("   %-18s", label+":"))
	}
	return lbl + " " + value + "\n"
}

func (m jwkFormModel) renderSelector(options []string, idx int, focused bool) string {
	left, right := "  ", "  "
	if idx > 0 {
		left = "◀ "
	}
	if idx < len(options)-1 {
		right = " ▶"
	}
	text := left + options[idx] + right
	if focused {
		return selectedStyle.Render(text)
	}
	return normalStyle.Render(text)
}

func (m jwkFormModel) renderToggle(val, focused bool) string {
	label := "[ ]  No"
	if val {
		label = "[✓]  Yes"
	}
	if focused {
		return selectedStyle.Render(label)
	}
	return normalStyle.Render(label)
}

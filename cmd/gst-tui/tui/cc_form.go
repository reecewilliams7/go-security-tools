package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/reecewilliams7/go-security-tools/clientcredentials"
)

type ccField int

const (
	ccFieldCount      ccField = iota // count text input
	ccFieldIDType                    // client-id-type selector
	ccFieldSecretType                // client-secret-type selector
	ccFieldSubmit                    // generate button
	ccNumFields                      // sentinel
)

var (
	idTypeOptions = []string{"uuidv7", "short-uuid"}
	secretOptions = []string{"crypto-rand"}
)

type ccFormModel struct {
	focused        ccField
	countInput     textinput.Model
	idTypeIdx      int
	secretTypeIdx  int
	generating     bool
	width          int
	height         int
}

func newCCFormModel() ccFormModel {
	count := textinput.New()
	count.Placeholder = "1"
	count.CharLimit = 3
	count.Width = 5
	count.Focus()

	return ccFormModel{
		focused:    ccFieldCount,
		countInput: count,
	}
}

func (m ccFormModel) init() tea.Cmd { return nil }

func (m ccFormModel) update(msg tea.Msg) (ccFormModel, tea.Cmd) {
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
			if m.focused == ccFieldSubmit {
				m.generating = true
				return m, m.submitCmd()
			}
			m = m.moveFocus(1)
			return m, nil
		case "left":
			switch m.focused {
			case ccFieldIDType:
				if m.idTypeIdx > 0 {
					m.idTypeIdx--
				}
			case ccFieldSecretType:
				if m.secretTypeIdx > 0 {
					m.secretTypeIdx--
				}
			}
			return m, nil
		case "right":
			switch m.focused {
			case ccFieldIDType:
				if m.idTypeIdx < len(idTypeOptions)-1 {
					m.idTypeIdx++
				}
			case ccFieldSecretType:
				if m.secretTypeIdx < len(secretOptions)-1 {
					m.secretTypeIdx++
				}
			}
			return m, nil
		}
	}

	// Forward remaining messages to count input when focused.
	var cmd tea.Cmd
	if m.focused == ccFieldCount {
		m.countInput, cmd = m.countInput.Update(msg)
	}
	return m, cmd
}

func (m ccFormModel) moveFocus(delta int) ccFormModel {
	next := int(m.focused) + delta
	if next < 0 {
		next = int(ccNumFields) - 1
	} else if next >= int(ccNumFields) {
		next = 0
	}
	m.focused = ccField(next)

	if m.focused == ccFieldCount {
		m.countInput.Focus()
	} else {
		m.countInput.Blur()
	}
	return m
}

func (m ccFormModel) submitCmd() tea.Cmd {
	idType := idTypeOptions[m.idTypeIdx]
	secretType := secretOptions[m.secretTypeIdx]
	countStr := m.countInput.Value()
	if countStr == "" {
		countStr = "1"
	}
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 1
	}
	if count > 100 {
		count = 100
	}

	return func() tea.Msg {
		var idCreator clientcredentials.ClientIDCreator
		switch idType {
		case "uuidv7":
			idCreator = clientcredentials.NewUUIDv7ClientIDCreator()
		case "short-uuid":
			idCreator = clientcredentials.NewShortUUIDClientIDCreator()
		default:
			return navigateMsg{to: screenResult, result: fmt.Sprintf("Unknown client ID type: %s", idType)}
		}

		var secretCreator clientcredentials.ClientSecretCreator
		switch secretType {
		case "crypto-rand":
			secretCreator = clientcredentials.NewCryptoRandClientSecretCreator()
		default:
			return navigateMsg{to: screenResult, result: fmt.Sprintf("Unknown client secret type: %s", secretType)}
		}

		ccc := clientcredentials.NewClientCredentialsCreator(idCreator, secretCreator)

		sep := strings.Repeat("─", 60)
		var sb strings.Builder
		for i := 1; i <= count; i++ {
			cc, err := ccc.CreateClientCredentials()
			if err != nil {
				return navigateMsg{to: screenResult, result: fmt.Sprintf("Error creating credentials: %v", err)}
			}
			sb.WriteString(sep + "\n")
			sb.WriteString(fmt.Sprintf("Credentials %d\n", i))
			sb.WriteString(sep + "\n")
			sb.WriteString(fmt.Sprintf("Client ID:     %s\n", cc.ClientID()))
			sb.WriteString(fmt.Sprintf("Client Secret: %s\n\n", cc.ClientSecret()))
		}
		return navigateMsg{to: screenResult, result: sb.String()}
	}
}

func (m ccFormModel) view() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Create Client Credentials"))
	sb.WriteString("\n\n")

	if m.generating {
		sb.WriteString(selectedStyle.Render("Generating credentials, please wait…"))
		sb.WriteString("\n")
		return sb.String()
	}

	sb.WriteString(m.renderField(ccFieldCount, "Count", m.countInput.View()))
	sb.WriteString(m.renderField(ccFieldIDType, "Client ID Type",
		m.renderSelector(idTypeOptions, m.idTypeIdx, m.focused == ccFieldIDType)))
	sb.WriteString(m.renderField(ccFieldSecretType, "Client Secret Type",
		m.renderSelector(secretOptions, m.secretTypeIdx, m.focused == ccFieldSecretType)))
	sb.WriteString("\n")

	if m.focused == ccFieldSubmit {
		sb.WriteString(selectedStyle.Render("▶  [ Generate ]"))
	} else {
		sb.WriteString(normalStyle.Render("   [ Generate ]"))
	}
	sb.WriteString("\n")

	sb.WriteString(helpStyle.Render("\ntab/↓ next  shift+tab/↑ prev  ←/→ change option  enter confirm  esc back"))
	return sb.String()
}

func (m ccFormModel) renderField(f ccField, label, value string) string {
	focused := m.focused == f
	var lbl string
	if focused {
		lbl = focusedLabelStyle.Render(fmt.Sprintf("▶  %-20s", label+":"))
	} else {
		lbl = unfocusedLabelStyle.Render(fmt.Sprintf("   %-20s", label+":"))
	}
	return lbl + " " + value + "\n"
}

func (m ccFormModel) renderSelector(options []string, idx int, focused bool) string {
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

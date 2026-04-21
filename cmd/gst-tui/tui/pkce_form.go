package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/reecewilliams7/go-security-tools/pkce"
)

type pkceField int

const (
	pkceFieldCount  pkceField = iota // count text input
	pkceFieldSubmit                  // generate button
	pkceNumFields                    // sentinel
)

type pkceFormModel struct {
	focused    pkceField
	countInput textinput.Model
	generating bool
	width      int
	height     int
}

func newPKCEFormModel() pkceFormModel {
	count := textinput.New()
	count.Placeholder = "1"
	count.CharLimit = 2
	count.Width = 5
	count.Focus()

	return pkceFormModel{
		focused:    pkceFieldCount,
		countInput: count,
	}
}

func (m pkceFormModel) init() tea.Cmd { return nil }

func (m pkceFormModel) update(msg tea.Msg) (pkceFormModel, tea.Cmd) {
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
			if m.focused == pkceFieldSubmit {
				m.generating = true
				return m, m.submitCmd()
			}
			m = m.moveFocus(1)
			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.focused == pkceFieldCount {
		m.countInput, cmd = m.countInput.Update(msg)
	}
	return m, cmd
}

func (m pkceFormModel) moveFocus(delta int) pkceFormModel {
	next := int(m.focused) + delta
	if next < 0 {
		next = int(pkceNumFields) - 1
	} else if next >= int(pkceNumFields) {
		next = 0
	}
	m.focused = pkceField(next)

	if m.focused == pkceFieldCount {
		m.countInput.Focus()
	} else {
		m.countInput.Blur()
	}
	return m
}

func (m pkceFormModel) submitCmd() tea.Cmd {
	countStr := m.countInput.Value()
	if countStr == "" {
		countStr = "1"
	}
	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		count = 1
	}
	if count > 20 {
		count = 20
	}

	return func() tea.Msg {
		creator := pkce.NewS256Creator()
		sep := strings.Repeat("─", 60)
		var sb strings.Builder
		for i := 1; i <= count; i++ {
			pair, err := creator.Create()
			if err != nil {
				return navigateMsg{to: screenResult, result: fmt.Sprintf("Error creating PKCE pair: %v", err)}
			}
			sb.WriteString(sep + "\n")
			sb.WriteString(fmt.Sprintf("PKCE Pair %d\n", i))
			sb.WriteString(sep + "\n")
			sb.WriteString(fmt.Sprintf("Code Verifier:  %s\n", pair.CodeVerifier))
			sb.WriteString(fmt.Sprintf("Code Challenge: %s\n", pair.CodeChallenge))
			sb.WriteString(fmt.Sprintf("Method:         %s\n\n", pair.Method))
		}
		return navigateMsg{to: screenResult, result: sb.String()}
	}
}

func (m pkceFormModel) view() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Create PKCE"))
	sb.WriteString("\n\n")

	if m.generating {
		sb.WriteString(selectedStyle.Render("Generating PKCE pairs, please wait…"))
		sb.WriteString("\n")
		return sb.String()
	}

	sb.WriteString(m.renderField(pkceFieldCount, "Count", m.countInput.View()))
	sb.WriteString("\n")

	if m.focused == pkceFieldSubmit {
		sb.WriteString(selectedStyle.Render("▶  [ Generate ]"))
	} else {
		sb.WriteString(normalStyle.Render("   [ Generate ]"))
	}
	sb.WriteString("\n")

	sb.WriteString(helpStyle.Render("\ntab/↓ next  shift+tab/↑ prev  enter confirm  esc back"))
	return sb.String()
}

func (m pkceFormModel) renderField(f pkceField, label, value string) string {
	focused := m.focused == f
	var lbl string
	if focused {
		lbl = selectedStyle.Render(fmt.Sprintf("▶  %-20s", label))
	} else {
		lbl = normalStyle.Render(fmt.Sprintf("   %-20s", label))
	}
	return lbl + "  " + value + "\n"
}

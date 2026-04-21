package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	cursor int
	width  int
	height int
}

var menuItems = []struct {
	label       string
	description string
}{
	{"Create JWK", "Generate JSON Web Keys (RSA-2048, RSA-4096, ECDSA-P256/384/521)"},
	{"Create Client Credentials", "Generate OAuth 2.0 client ID and secret pair"},
	{"Quit", "Exit the application"},
}

func newMenuModel() menuModel { return menuModel{} }

func (m menuModel) update(msg tea.Msg) (menuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(menuItems)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.cursor {
			case 0:
				return m, navigate(screenJWKForm)
			case 1:
				return m, navigate(screenCCForm)
			case 2:
				return m, tea.Quit
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) view() string {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("gst-tui"))
	sb.WriteString("\n")
	sb.WriteString(subtitleStyle.Render("Go Security Tools  —  OAuth 2.0 · OpenID Connect · JWK management"))
	sb.WriteString("\n\n")

	for i, item := range menuItems {
		if i == m.cursor {
			sb.WriteString(selectedStyle.Render(fmt.Sprintf("▶  %s", item.label)))
			sb.WriteString("\n")
			sb.WriteString(subtitleStyle.Render(fmt.Sprintf("   %s", item.description)))
		} else {
			sb.WriteString(normalStyle.Render(fmt.Sprintf("   %s", item.label)))
		}
		sb.WriteString("\n")
	}

	sb.WriteString(helpStyle.Render("\n↑/↓  navigate   enter  select   q  quit"))
	return sb.String()
}

package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type resultModel struct {
	viewport viewport.Model
	width    int
	height   int
}

func newResultModel(content string, width, height int) resultModel {
	if width == 0 {
		width = 80
	}
	if height == 0 {
		height = 24
	}
	// Reserve 4 lines for title and footer.
	vp := viewport.New(width, height-4)
	vp.SetContent(content)
	return resultModel{viewport: vp, width: width, height: height}
}

func (m resultModel) init() tea.Cmd { return nil }

func (m resultModel) update(msg tea.Msg) (resultModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, navigate(screenMenu)
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 4
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m resultModel) view() string {
	var sb strings.Builder
	sb.WriteString(titleStyle.Render("Results"))
	sb.WriteString("\n")
	sb.WriteString(m.viewport.View())
	sb.WriteString("\n")
	sb.WriteString(helpStyle.Render("↑/↓  scroll   esc/q  back to menu"))
	return sb.String()
}

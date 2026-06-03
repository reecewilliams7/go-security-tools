package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type resultModel struct {
	viewport viewport.Model
	content  string // raw unwrapped content for re-wrapping on resize
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
	// Reserve 5 lines: title (2 with MarginBottom), blank, newline after viewport, footer.
	vp := viewport.New(width, height-5)
	vp.SetContent(wrapContent(content, width))
	return resultModel{viewport: vp, content: content, width: width, height: height}
}

// wrapContent hard-wraps each line of plain-text content at width characters.
func wrapContent(content string, width int) string {
	if width <= 0 {
		return content
	}
	lines := strings.Split(content, "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		for len(line) > width {
			out = append(out, line[:width])
			line = line[width:]
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
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
		m.viewport.Height = msg.Height - 5
		m.viewport.SetContent(wrapContent(m.content, msg.Width))
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

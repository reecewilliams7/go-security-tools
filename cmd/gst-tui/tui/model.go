package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenMenu screen = iota
	screenJWKForm
	screenCCForm
	screenResult
)

// navigateMsg triggers a screen transition.
type navigateMsg struct {
	to     screen
	result string
}

func navigate(to screen) tea.Cmd {
	return func() tea.Msg { return navigateMsg{to: to} }
}

// Model is the root TUI model that owns all screen sub-models.
type Model struct {
	screen screen
	menu   menuModel
	jwk    jwkFormModel
	cc     ccFormModel
	result resultModel
	width  int
	height int
}

// NewModel returns the initial root model showing the main menu.
func NewModel() Model {
	return Model{
		screen: screenMenu,
		menu:   newMenuModel(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.menu.width, m.menu.height = msg.Width, msg.Height
		m.jwk.width, m.jwk.height = msg.Width, msg.Height
		m.cc.width, m.cc.height = msg.Width, msg.Height
		if m.screen == screenResult {
			m.result.viewport.Width = msg.Width
			m.result.viewport.Height = msg.Height - 4
		}
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case navigateMsg:
		return m.handleNavigate(msg)
	}

	switch m.screen {
	case screenMenu:
		var cmd tea.Cmd
		m.menu, cmd = m.menu.update(msg)
		return m, cmd
	case screenJWKForm:
		var cmd tea.Cmd
		m.jwk, cmd = m.jwk.update(msg)
		return m, cmd
	case screenCCForm:
		var cmd tea.Cmd
		m.cc, cmd = m.cc.update(msg)
		return m, cmd
	case screenResult:
		var cmd tea.Cmd
		m.result, cmd = m.result.update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	switch m.screen {
	case screenMenu:
		return m.menu.view()
	case screenJWKForm:
		return m.jwk.view()
	case screenCCForm:
		return m.cc.view()
	case screenResult:
		return m.result.view()
	}
	return ""
}

func (m Model) handleNavigate(nav navigateMsg) (Model, tea.Cmd) {
	m.screen = nav.to
	switch nav.to {
	case screenJWKForm:
		m.jwk = newJWKFormModel()
		m.jwk.width, m.jwk.height = m.width, m.height
		return m, m.jwk.init()
	case screenCCForm:
		m.cc = newCCFormModel()
		m.cc.width, m.cc.height = m.width, m.height
		return m, m.cc.init()
	case screenMenu:
		m.menu = newMenuModel()
		m.menu.width, m.menu.height = m.width, m.height
		return m, nil
	case screenResult:
		m.result = newResultModel(nav.result, m.width, m.height)
		return m, m.result.init()
	}
	return m, nil
}

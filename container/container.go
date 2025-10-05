package container

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	content string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *Model) SetContent(content string) {
	m.content = content
}

func (m Model) View() string {
	return m.content
}

func New() Model {
	return Model{content: ""}
}

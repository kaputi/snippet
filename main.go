package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kaputi/snippets/content"
	"github.com/kaputi/snippets/lang"
	"github.com/kaputi/snippets/logger"
	"github.com/kaputi/snippets/snippet"
	"github.com/kaputi/snippets/theme"
	"github.com/kaputi/snippets/tree"
)

type focusState uint

const (
	langPanel focusState = iota
	treePanel
	snipeetPanel
	contentPanel
)

type column struct {
	content string
}

func (c column) Init() tea.Cmd {
	return nil
}

func (c column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c column) SetContent(content string) {
	c.content = content
}

func (c column) View() string {
	return c.content
}

func newColumn() column {
	return column{content: ""}
}

type maiContainer struct {
	focusPanel   focusState
	leftColumn   column
	rightColumn  column
	langModel    lang.Model
	treeModel    tree.Model
	snippetModel snippet.Model
	contentModel content.Model
}

func newModel() maiContainer {
	return maiContainer{
		focusPanel:   0,
		leftColumn:   newColumn(),
		rightColumn:  newColumn(),
		langModel:    lang.NewModel(),
		treeModel:    tree.NewModel(),
		snippetModel: snippet.NewModel(),
		contentModel: content.NewModel(),
	}
}

func (m maiContainer) Init() tea.Cmd {
	return tea.Batch(m.langModel.Init(), m.treeModel.Init(), m.snippetModel.Init(), m.contentModel.Init())
}

func (m maiContainer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "j", "down":
			m.focusPanel++
			if m.focusPanel > snipeetPanel {
				m.focusPanel = langPanel
			}
		case "shift+tab", "k", "up":
			m.focusPanel++
			if m.focusPanel > snipeetPanel {
				m.focusPanel = langPanel
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m maiContainer) View() string {

	// leftColumnStyle := theme.LeftColumn()
	// rightColumnSyle := theme.RightColumn()
	navPanelStyle := theme.NavPanel()
	focusedNavPanelStyle := theme.FocusedNavPanel()

	contentPanelStyle := theme.ContentPanel()

	var langString, treeString, snippetString string

	switch m.focusPanel {
	case langPanel:
		langString = focusedNavPanelStyle.Render(m.langModel.View())
		treeString = navPanelStyle.Render(m.treeModel.View())
		snippetString = navPanelStyle.Render(m.snippetModel.View())
	case treePanel:
		langString = navPanelStyle.Render(m.langModel.View())
		treeString = focusedNavPanelStyle.Render(m.treeModel.View())
		snippetString = navPanelStyle.Render(m.snippetModel.View())
	case snipeetPanel:
		langString = navPanelStyle.Render(m.langModel.View())
		treeString = navPanelStyle.Render(m.treeModel.View())
		snippetString = focusedNavPanelStyle.Render(m.snippetModel.View())
	}

	leftContent := lipgloss.JoinVertical(lipgloss.Top, langString, treeString, snippetString)
	m.leftColumn.SetContent(leftContent)

	rightContent := contentPanelStyle.Render(m.contentModel.View())
	m.rightColumn.SetContent("aaaaaaaaaaaa")

	s := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, rightContent)

	return s
}

func main() {
	err := logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	logger.Log("Application started")

	p := tea.NewProgram(newModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	err = logger.Close()
	if err != nil {
		log.Fatal(err)
	}
}

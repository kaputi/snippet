package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaputi/snippets/logger"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select your choices (press q to quit):\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		checked := "[ ]" // not selected
		if _, ok := m.selected[i]; ok {
			checked = "[x]" // selected
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)
	}
	s += "\nPress enter or space to toggle selection.\n"
	return s
}

func main() {
	err := logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	logger.Log("Application started")

	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	err = logger.Close()
	if err != nil {
		log.Fatal(err)
	}
}

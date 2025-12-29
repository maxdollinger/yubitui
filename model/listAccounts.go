package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListAccountsModel struct {
	key      ListAccountsI
	accounts []string
	cursor   int
}

func NewListAccountsModel(key ListAccountsI) *ListAccountsModel {
	return &ListAccountsModel{key: key}
}

func (m *ListAccountsModel) Init() tea.Cmd {
	return ListAccountsCmd(m.key)
}

var selectedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))

func (m *ListAccountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AccountsListedMsg:
		m.accounts = msg
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q", tea.KeyEsc.String():
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.accounts)-1 {
				m.cursor++
			}

		case "enter", " ":
			return m, SwitchToGenerateCodeModelCmd(m.accounts[m.cursor])
		}
	}

	return m, nil
}

func (m *ListAccountsModel) View() string {
	var s strings.Builder
	s.WriteString("Generate code for:\n\n")

	for i, choice := range m.accounts {

		line := choice
		if m.cursor == i {
			line = selectedStyle.Render(line)
		}

		fmt.Fprintf(&s, "%s\n", line)
	}

	s.WriteString("\nPress q to quit, n to add account.\n")

	return s.String()
}

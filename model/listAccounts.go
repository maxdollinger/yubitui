package model

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ListAccountsKeyI interface {
	ListAccountsI
	DeleteAccountI
}

type ListAccountsModel struct {
	key      ListAccountsKeyI
	accounts []string
	cursor   int
}

func NewListAccountsModel(key ListAccountsKeyI) *ListAccountsModel {
	return &ListAccountsModel{key: key}
}

func (m *ListAccountsModel) Init() tea.Cmd {
	return ListAccountsCmd(m.key)
}

func (m *ListAccountsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AccountsListedMsg:
		m.accounts = msg
	case AccountDeletedMsg:
		m.accounts = m.accounts[0:0]
		return m, ListAccountsCmd(m.key)
	case tea.KeyMsg:

		switch msg.String() {
		case "q", tea.KeyEsc.String():
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.accounts)-1 {
				m.cursor++
			}

		case "n":
			return m, SwitchToAddAccountModelCmd()

		case "d":
			return m, SwitchToDeleteAccountModel(m.accounts[m.cursor])

		case "enter":
			return m, SwitchToGenerateCodeModelCmd(m.accounts[m.cursor])
		}
	}

	return m, nil
}

func (m *ListAccountsModel) View() string {
	var s strings.Builder
	s.WriteString("Accounts:\n\n")

	for i, choice := range m.accounts {

		line := choice
		if m.cursor == i {
			line = listSelectedItem.Render(line)
		}

		fmt.Fprintf(&s, "%s\n", line)
	}

	s.WriteString("\nPress q to quit, enter to generate code, n to add account, d to delete account.\n")

	return s.String()
}

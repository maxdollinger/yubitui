package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
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
	spinner  spinner.Model
}

func NewListAccountsModel(key ListAccountsKeyI) *ListAccountsModel {
	sModel := spinner.New()

	return &ListAccountsModel{
		key:     key,
		spinner: sModel,
	}
}

func (m *ListAccountsModel) Init() tea.Cmd {
	return tea.Batch(ListAccountsCmd(m.key), m.spinner.Tick)
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

		case "r":
			return m, SwitchToRenameAccountModel(m.accounts[m.cursor])

		case "d":
			return m, SwitchToDeleteAccountModel(m.accounts[m.cursor])

		case "enter":
			return m, SwitchToGenerateCodeModelCmd(m.accounts[m.cursor])
		}
	}

	if len(m.accounts) == 0 {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *ListAccountsModel) View() string {
	if len(m.accounts) == 0 {
		return fmt.Sprintf("\n%s loading accounts...\n", m.spinner.View())
	}

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

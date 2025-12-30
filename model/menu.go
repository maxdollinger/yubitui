package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	key      ListAccountsI
	accounts []string
	cursor   int
	spinner  spinner.Model
	showHelp bool
}

func NewMainMenuModel(key ListAccountsI) *MenuModel {
	sModel := spinner.New()

	return &MenuModel{
		key:     key,
		spinner: sModel,
	}
}

func (m *MenuModel) Init() tea.Cmd {
	return tea.Batch(ListAccountsCmd(m.key), m.spinner.Tick)
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "h":
			m.showHelp = !m.showHelp
			return m, nil

		case "n":
			return m, NewAddModelCmd()

		case "r":
			return m, NewRenameModelCmd(m.accounts[m.cursor])

		case "d":
			return m, NewDeleteModelCmd(m.accounts[m.cursor])

		case "enter":
			return m, NewCodeModelCmd(m.accounts[m.cursor])
		}
	}

	if len(m.accounts) == 0 {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *MenuModel) View() string {
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

	s.WriteString("\n")
	if m.showHelp {
		s.WriteString(menuHelpText())
	} else {
		s.WriteString("Press q to quit, h to toggle help.\n")
	}

	return s.String()
}

func menuHelpText() string {
	var sb strings.Builder

	sb.WriteString("j,k    - move down/up\n")
	sb.WriteString("enter  - generate code for selected account\n")
	sb.WriteString("n      - add new accound\n")
	sb.WriteString("d      - delete selected account\n")
	sb.WriteString("r      - rename selected account\n")
	sb.WriteString("esc,q  - exit the program\n")

	return sb.String()
}

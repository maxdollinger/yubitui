package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type DeleteAccountModel struct {
	key     DeleteAccountI
	account string
}

func NewDeleteAccountModel(key DeleteAccountI, account string) *DeleteAccountModel {
	return &DeleteAccountModel{
		key:     key,
		account: account,
	}
}

func (m *DeleteAccountModel) Init() tea.Cmd {
	return nil
}

func (m *DeleteAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AccountDeletedMsg:
		return m, SwitchToListAccountsModelCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "y":
			return m, DeleteAccountCmd(m.key, m.account)
		case "n":
			return m, SwitchToListAccountsModelCmd()
		}
	}

	return m, nil
}

func (m *DeleteAccountModel) View() string {
	return fmt.Sprintf("\nAre you sure to delete account \"%s\" [y/n]\n", m.account)
}

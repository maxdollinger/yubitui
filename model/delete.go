package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type DeleteModel struct {
	key     DeleteAccountI
	account string
}

func NewDeleteModel(key DeleteAccountI, account string) *DeleteModel {
	return &DeleteModel{
		key:     key,
		account: account,
	}
}

func (m *DeleteModel) Init() tea.Cmd {
	return nil
}

func (m *DeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case AccountDeletedMsg:
		return m, NewMainMenuModelCmd()
	case tea.KeyMsg:
		switch msg.String() {
		case "y":
			return m, DeleteAccountCmd(m.key, m.account)
		case "n":
			return m, NewMainMenuModelCmd()
		}
	}

	return m, nil
}

func (m *DeleteModel) View() string {
	return fmt.Sprintf("\nAre you sure to delete account \"%s\" [y/n]\n", m.account)
}

package model

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type GenerateCodeModel struct {
	key         GenerateCodeI
	clipboard   CopyToClipboardI
	account     string
	copySuccess bool
	code        string
}

func NewGenerateCodeModel(key GenerateCodeI, clipboard CopyToClipboardI, account string) *GenerateCodeModel {
	return &GenerateCodeModel{
		key:         key,
		clipboard:   clipboard,
		account:     account,
		copySuccess: false,
	}
}

func (m *GenerateCodeModel) Init() tea.Cmd {
	return GenerateCodeCmd(m.key, m.account)
}

func (m *GenerateCodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ErrMsg:
		log.Fatal(msg)

	case CodeGeneratedMsg:
		m.code = string(msg)
		if m.clipboard != nil {
			if err := m.clipboard.Copy(m.code); err == nil {
				m.copySuccess = true
			}
		}
		return m, tea.Quit
	}

	return m, nil
}

func (m *GenerateCodeModel) View() string {
	copied := ""
	if m.copySuccess {
		copied = " -> copied to clipboard"
	}
	return fmt.Sprintf("Code for %s:\n%s%s\n", m.account, m.code, copied)
}

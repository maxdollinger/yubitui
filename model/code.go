package model

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type CodeModel struct {
	key         GenerateCodeI
	clipboard   CopyToClipboardI
	account     string
	copySuccess bool
	code        string
	waitCnt     int
	waitTime    int
}

func NewCodeModel(key GenerateCodeI, clipboard CopyToClipboardI, account string) *CodeModel {
	return &CodeModel{
		key:         key,
		clipboard:   clipboard,
		account:     account,
		copySuccess: false,
		waitTime:    5,
	}
}

func (m *CodeModel) Init() tea.Cmd {
	return GenerateCodeCmd(m.key, m.account)
}

func (m *CodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ErrMsg:
		log.Fatal(msg)

	case IntervalMsg:
		m.waitCnt += int(msg.duration.Seconds())
		if m.waitCnt == 5 {
			return m, NewMainMenuModelCmd()
		} else {
			return m, IntervalCmd(time.Second)
		}

	case CodeGeneratedMsg:
		m.code = string(msg)
		if m.clipboard != nil {
			if err := m.clipboard.Copy(m.code); err == nil {
				m.copySuccess = true
			}
		}

		return m, IntervalCmd(time.Second)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, NewMainMenuModelCmd()
		}
	}

	return m, nil
}

func (m *CodeModel) View() string {
	copied := ""
	if m.copySuccess {
		copied = "  ... code copied"
	}

	left := m.waitTime - m.waitCnt
	waitMsg := fmt.Sprintf("back to accounts in %ds", left)
	return fmt.Sprintf("Code for %s:\n%s\n%s\n\n%s\n", m.account, m.code, copied, waitMsg)
}

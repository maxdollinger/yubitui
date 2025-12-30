package model

import (
	"fmt"
	"strings"
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
				return m, IntervalCmd(time.Second)
			}
		}

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, NewMainMenuModelCmd()
		}
	}

	return m, nil
}

func (m *CodeModel) View() string {
	if len(m.code) == 0 {
		return "\n...generating code\n"
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "TOTP code for '%s':\n\n", m.account)
	sb.WriteString(codeStyle.Render(m.code))
	if m.copySuccess {
		sb.WriteString("  ... code copied\n\n")
		left := m.waitTime - m.waitCnt
		fmt.Fprintf(&sb, "back to accounts in %ds\n", left)
	} else {
		sb.WriteString("\n\npress q to go back or ctrl+c to quit")
	}

	return sb.String()
}

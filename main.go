package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdollinger/yubitui/clipboard"
	"github.com/mdollinger/yubitui/yubikey"
)

type (
	codeMsg    string
	errMsg     error
	accountMsg []string
)

type model struct {
	accounts  []string
	cursor    int
	code      string
	key       *yubikey.Yubikey
	clipboard *clipboard.Clipboard
	copied    bool
}

func initialModel() *model {
	key, err := yubikey.InitYubikey()
	if err != nil {
		log.Fatal(err)
	}

	clip, err := clipboard.InitClipboard()
	if err != nil {
		log.Println(err.Error())
	}
	return &model{
		key:       key,
		clipboard: clip,
	}
}

func (m model) Init() tea.Cmd {
	return ListAccountsCmd(m.key)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		log.Fatal(msg)

	case codeMsg:
		m.code = string(msg)
		if m.clipboard != nil {
			if err := m.clipboard.Copy(m.code); err == nil {
				m.copied = true
			}
		}
		return m, tea.Quit

	case accountMsg:
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
			account := m.accounts[m.cursor]
			return m, GenerateCodeCmd(m.key, account)
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.code != "" {
		account := m.accounts[m.cursor]
		copied := ""
		if m.copied {
			copied = " -> copied to clipboard"
		}
		return fmt.Sprintf("Code for %s:\n%s%s\n", account, m.code, copied)
	}

	return AccountView(&m)
}

func main() {
	model := initialModel()
	defer model.key.Close()

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %s", err)
	}
}

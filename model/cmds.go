package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ErrMsg error

func ErrCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return ErrMsg(err)
	}
}

type NewCodeModelMsg struct {
	Account string
}

func NewCodeModelCmd(account string) tea.Cmd {
	return func() tea.Msg {
		return NewCodeModelMsg{Account: account}
	}
}

type NewMainMenuModelMsg struct{}

func NewMainMenuModelCmd() tea.Cmd {
	return func() tea.Msg {
		return NewMainMenuModelMsg{}
	}
}

type NewAddModelMsg struct{}

func NewAddModelCmd() tea.Cmd {
	return func() tea.Msg {
		return NewAddModelMsg{}
	}
}

type NewDeleteModelMsg struct {
	account string
}

func NewDeleteModelCmd(account string) tea.Cmd {
	return func() tea.Msg {
		return NewDeleteModelMsg{
			account: account,
		}
	}
}

type NewRenameModelMsg struct {
	account string
}

func NewRenameModelCmd(account string) tea.Cmd {
	return func() tea.Msg {
		return NewRenameModelMsg{
			account: account,
		}
	}
}

type CodeGeneratedMsg string

func GenerateCodeCmd(key GenerateCodeI, account string) func() tea.Msg {
	return func() tea.Msg {
		code, err := key.GenerateCode(account)
		if err != nil {
			return ErrMsg(err)
		}

		return CodeGeneratedMsg(code)
	}
}

type AccountsListedMsg []string

func ListAccountsCmd(key ListAccountsI) tea.Cmd {
	return func() tea.Msg {
		accounts, err := key.ListAccounts()
		if err != nil {
			return ErrMsg(err)
		}

		return AccountsListedMsg(accounts)
	}
}

type AccountDeletedMsg bool

func DeleteAccountCmd(key DeleteAccountI, account string) tea.Cmd {
	return func() tea.Msg {
		err := key.DeleteAccount(account)
		if err != nil {
			return ErrMsg(err)
		}

		return AccountDeletedMsg(true)
	}
}

type AccountRenamedMsg struct {
	old string
	new string
}

func RenameAccountCmd(key RenameAccountI, account string, name string) tea.Cmd {
	return func() tea.Msg {
		err := key.RenameAccount(account, name)
		if err != nil {
			return ErrMsg(err)
		}

		return AccountRenamedMsg{
			old: account,
			new: name,
		}
	}
}

type IntervalMsg struct {
	duration time.Duration
	tickTime time.Time
}

func IntervalCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return IntervalMsg{
			duration: d,
			tickTime: t,
		}
	})
}

func KeyCmd(msg tea.KeyMsg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ErrMsg error

type GenerateCodeModelMsg struct {
	Account string
}

func SwitchToGenerateCodeModelCmd(account string) tea.Cmd {
	return func() tea.Msg {
		return GenerateCodeModelMsg{Account: account}
	}
}

type ListAccountsModelMsg struct{}

func SwitchToListAccountsModelCmd() tea.Cmd {
	return func() tea.Msg {
		return ListAccountsModelMsg{}
	}
}

type AddAccountModelMsg struct{}

func SwitchToAddAccountModelCmd() tea.Cmd {
	return func() tea.Msg {
		return AddAccountModelMsg{}
	}
}

type DeleteAccountModelMsg struct {
	account string
}

func SwitchToDeleteAccountModel(account string) tea.Cmd {
	return func() tea.Msg {
		return DeleteAccountModelMsg{
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

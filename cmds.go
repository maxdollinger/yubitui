package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdollinger/yubitui/yubikey"
)

func ListAccountsCmd(key *yubikey.Yubikey) func() tea.Msg {
	return func() tea.Msg {
		accounts, err := key.ListAccounts()
		if err != nil {
			return errMsg(err)
		}

		return accountMsg(accounts)
	}
}

func GenerateCodeCmd(key *yubikey.Yubikey, account string) func() tea.Msg {
	return func() tea.Msg {
		code, err := key.GenerateCode(account)
		if err != nil {
			return errMsg(err)
		}

		return codeMsg(code)
	}
}

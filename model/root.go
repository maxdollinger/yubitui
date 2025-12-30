package model

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdollinger/yubitui/clipboard"
	"github.com/mdollinger/yubitui/yubikey"
)

type RootModel struct {
	key         YubiKeyI
	clipboard   *clipboard.Clipboard
	activeModel tea.Model
}

func NewRootModel() *RootModel {
	key, err := yubikey.InitYubikey()
	if err != nil {
		log.Fatal(err)
	}

	clip, err := clipboard.InitClipboard()
	if err != nil {
		log.Println(err.Error())
	}

	return &RootModel{
		key:         key,
		clipboard:   clip,
		activeModel: NewMainMenuModel(key),
	}
}

func (m RootModel) Init() tea.Cmd {
	return m.activeModel.Init()
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case NewMainMenuModelMsg:
		listAccountsModel := NewMainMenuModel(m.key)
		m.activeModel = listAccountsModel
		return m, m.activeModel.Init()

	case NewCodeModelMsg:
		generateCodeModel := NewCodeModel(m.key, m.clipboard, msg.Account)
		m.activeModel = generateCodeModel
		return m, m.activeModel.Init()

	case NewAddModelMsg:
		addAccountModel := NewAddModel(m.key, m.clipboard)
		m.activeModel = addAccountModel
		return m, m.activeModel.Init()

	case NewDeleteModelMsg:
		deleteAccountModel := NewDeleteModel(m.key, msg.account)
		m.activeModel = deleteAccountModel
		return m, m.activeModel.Init()

	case NewRenameModelMsg:
		renameAccountModel := NewRenameModel(m.key, msg.account)
		m.activeModel = renameAccountModel
		return m, m.activeModel.Init()

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.activeModel, cmd = m.activeModel.Update(msg)
	return m, cmd
}

func (m *RootModel) View() string {
	return m.activeModel.View()
}

func (m *RootModel) Cleanup() {
	m.key.Close()
}

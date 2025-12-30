package model

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RenameModel struct {
	focusIndex int
	nameInput  textinput.Model
	key        RenameAccountI
	account    string
	inputMode  bool
	showHelp   bool
	keyStack   []string
}

func NewRenameModel(key RenameAccountI, account string) *RenameModel {
	m := RenameModel{
		key:     key,
		account: account,
	}

	t := textinput.New()
	t.Cursor.Style = cursorStyle
	t.Cursor.SetMode(cursor.CursorHide)
	t.CharLimit = 64
	t.Placeholder = "name"
	t.SetValue(m.account)
	t.Focus()
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle

	m.nameInput = t

	return &m
}

func (m *RenameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *RenameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputMode {
		return m.InputModeUpdates(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.keyStack = append(m.keyStack, msg.String())
		switch msg.String() {
		case "i":
			if input, ok := m.getSelectedInput(); ok {
				cmd := input.Cursor.SetMode(cursor.CursorBlink)
				m.inputMode = true
				return m, cmd
			}

		case "q":
			return m, NewMainMenuModelCmd()

		case "h":
			m.showHelp = !m.showHelp
			return m, nil

		case "d":
			lastKeyIdx := len(m.keyStack) - 1
			if lastKeyIdx > 0 && m.keyStack[lastKeyIdx-1] != "d" {
				return m, nil
			}

			if input, ok := m.getSelectedInput(); ok {
				input.Reset()
			}

		// Set focus to next input
		case "tab", "enter", "up", "down", "j", "k":
			s := msg.String()

			if s == "enter" && m.focusIndex == 1 {
				err := m.key.RenameAccount(m.account, m.nameInput.Value())
				if err != nil {
					fmt.Println(err)
				}

				return m, NewMainMenuModelCmd()
			}
			// Cycle indexes
			m.moveFocus(s)

			var cmd tea.Cmd
			if m.focusIndex == 0 {
				m.nameInput.PromptStyle = focusedStyle
				m.nameInput.TextStyle = focusedStyle
				cmd = m.nameInput.Focus()
			} else {
				m.nameInput.Blur()
				m.nameInput.PromptStyle = noStyle
				m.nameInput.TextStyle = noStyle
			}

			return m, cmd
		}
	}

	return m, nil
}

func (m *RenameModel) moveFocus(s string) {
	if s == "up" || s == "k" {
		m.focusIndex--
	} else {
		m.focusIndex++
	}

	if m.focusIndex > 1 {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = 0
	}
}

func (m *RenameModel) getSelectedInput() (*textinput.Model, bool) {
	if m.focusIndex == 0 {
		return &m.nameInput, true
	}

	return nil, false
}

func (m *RenameModel) InputModeUpdates(msg tea.Msg) (tea.Model, tea.Cmd) {
	input, ok := m.getSelectedInput()
	if !ok {
		m.inputMode = false
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter":
			m.inputMode = false
			input.SetCursor(len(input.Value()))
			cmd := input.Cursor.SetMode(cursor.CursorHide)
			return m, tea.Batch(func() tea.Msg {
				return msg
			}, cmd)
		case "ctrl+h":
			current := input.Position()
			input.SetCursor(current - 1)
			return m, nil
		case "ctrl+l":
			current := input.Position()
			input.SetCursor(current + 1)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.nameInput, cmd = input.Update(msg)

	defaultWidth := len(input.Placeholder)
	inputWidth := len(input.Value())
	if inputWidth > 0 {
		input.Width = inputWidth
	} else {
		input.Width = defaultWidth
	}

	return m, cmd
}

func (m *RenameModel) View() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Rename account \"%s\"\n\n", m.account)
	fmt.Fprintf(&b, "%s\n", m.nameInput.View())

	button := bluredButton("rename account")
	if m.focusIndex == 1 {
		button = focusedButton("rename account")
	}
	fmt.Fprintf(&b, "\n%s\n\n", button)
	if m.showHelp {
		fmt.Fprint(&b, addHelptText())
	} else {
		fmt.Fprintf(&b, "press q to go back, h to show help")
	}

	return b.String()
}

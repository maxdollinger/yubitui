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

type AddAccountModel struct {
	focusIndex int
	inputs     []textinput.Model
	key        AddAccountI
	inputMode  bool
	showHelp   bool
	keyStack   []string
}

func NewAddAcountModel(key AddAccountI, clip PasteI) *AddAccountModel {
	m := AddAccountModel{
		inputs: make([]textinput.Model, 2),
		key:    key,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.Cursor.SetMode(cursor.CursorHide)
		t.CharLimit = 64

		switch i {
		case 0:
			t.Placeholder = "account"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "secret"
			t.SetValue(clip.Paste())
		}

		t.Width = len(t.Placeholder)
		m.inputs[i] = t
	}

	return &m
}

func (m *AddAccountModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *AddAccountModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, SwitchToListAccountsModelCmd()

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

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				err := m.key.AddAccount(m.inputs[0].Value(), m.inputs[1].Value(), 6)
				if err != nil {
					fmt.Println(err)
				}

				return m, SwitchToGenerateCodeModelCmd(m.inputs[0].Value())
			}
			// Cycle indexes
			if s == "up" || s == "shift+tab" || s == "k" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	return m, nil
}

func (m *AddAccountModel) getSelectedInput() (*textinput.Model, bool) {
	if m.focusIndex < len(m.inputs) && m.focusIndex >= 0 {
		return &m.inputs[m.focusIndex], true
	}

	return nil, false
}

func (m *AddAccountModel) InputModeUpdates(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, cmd
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
	m.inputs[m.focusIndex], cmd = input.Update(msg)

	defaultWidth := len(input.Placeholder)
	inputWidth := len(input.Value())
	if inputWidth > 0 {
		input.Width = inputWidth
	} else {
		input.Width = defaultWidth
	}

	return m, cmd
}

func (m *AddAccountModel) View() string {
	var b strings.Builder

	b.WriteString("Add new TOTP account\n\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := bluredButton("add account")
	if m.focusIndex == len(m.inputs) {
		button = focusedButton("add account")
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", button)
	if m.showHelp {
		fmt.Fprint(&b, HelpText())
	} else {
		fmt.Fprintf(&b, "press q to go back, h to show help")
	}

	return b.String()
}

func HelpText() string {
	var b strings.Builder

	fmt.Fprintln(&b, "i   - enter input mode")
	fmt.Fprintln(&b, "esc - leave input mode")
	fmt.Fprintln(&b, "dd  - delete input field")
	fmt.Fprintln(&b, "q   - go back to accounts list")

	return b.String()
}

package model

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdollinger/yubitui/utils"
)

type AddModel struct {
	focusIndex int
	inputs     []textinput.Model
	key        AddAccountI
	clipboard  PasteI
	inputMode  bool
	showHelp   bool
	keyStack   []string
}

func NewAddModel(key AddAccountI, clip PasteI) *AddModel {
	m := AddModel{
		inputs:    make([]textinput.Model, 2),
		key:       key,
		clipboard: clip,
	}

	for i := range m.inputs {
		t := textinput.New()
		t.Cursor.Style = cursorStyle
		t.Cursor.SetMode(cursor.CursorHide)
		t.CharLimit = 64
		t.Width = 64

		switch i {
		case 0:
			t.Placeholder = "account"
			t.Focus()
			setInputFocusedStyle(&t)
		case 1:
			t.Placeholder = "secret"
		}

		m.inputs[i] = t
	}

	return &m
}

func (m *AddModel) Init() tea.Cmd {
	return nil
}

func (m *AddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputMode {
		return m.InputModeUpdates(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		m.keyStack = append(m.keyStack, s)
		switch s {
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

		case "p":
			if input, ok := m.getSelectedInput(); ok {
				str := m.clipboard.Paste()
				str = strings.ReplaceAll(str, "\n", " ")
				str = strings.TrimSpace(str)
				if len(str) > 64 {
					str = str[:64]
				}

				input.SetValue(str)
			}

			return m, nil

		// Set focus to next input
		case "tab", "enter", "up", "down", "j", "k":
			// Did the user press enter while the submit button was focused?
			if s == "enter" && m.focusIndex == len(m.inputs) {
				name := m.inputs[0].Value()
				secret := m.inputs[1].Value()

				return m, AddAccountCmd(m.key, name, secret)
			}
			// Cycle indexes
			if s == "up" || s == "k" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			m.focusIndex = utils.Clamp(m.focusIndex, 0, len(m.inputs))

			var cmd tea.Cmd
			for i := 0; i <= len(m.inputs)-1; i++ {
				input := &m.inputs[i]
				if i == m.focusIndex {
					// Set focused state
					cmd = input.Focus()
					setInputFocusedStyle(input)
				} else {
					// Remove focused state
					input.Blur()
					setInputNoStyle(input)
				}
			}

			return m, cmd
		}
	}

	return m, nil
}

func (m *AddModel) getSelectedInput() (*textinput.Model, bool) {
	if m.focusIndex < len(m.inputs) && m.focusIndex >= 0 {
		return &m.inputs[m.focusIndex], true
	}

	return nil, false
}

func (m *AddModel) InputModeUpdates(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, tea.Batch(cmd, KeyCmd(msg))
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
	return m, cmd
}

func (m *AddModel) View() string {
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
		fmt.Fprint(&b, addHelptText())
	} else {
		fmt.Fprintf(&b, "press q to go back, h to toggle help")
	}

	return b.String()
}

func addHelptText() string {
	var b strings.Builder

	fmt.Fprintln(&b, "i     - enter input mode")
	fmt.Fprintln(&b, "p     - to paste from clipboard")
	fmt.Fprintln(&b, "esc   - leave input mode")
	fmt.Fprintln(&b, "enter - leave input mode and go to next field")
	fmt.Fprintln(&b, "ctr+h - move cursor left while in input mode")
	fmt.Fprintln(&b, "ctr+l - move cursor right while in input mode")
	fmt.Fprintln(&b, "dd    - delete input field")
	fmt.Fprintln(&b, "q     - go back to accounts list")

	return b.String()
}

package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const (
	SELECTED_COLOR = "#7D56F4"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color(SELECTED_COLOR))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()
	codeStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00cc22"))
	errStyle     = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cc0000")).
			Border(lipgloss.RoundedBorder(), true, true, true, true).
			MarginLeft(4).
			Padding(2)

	listSelectedItem = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color(SELECTED_COLOR))
)

func focusedButton(text string) string {
	return focusedStyle.Render(fmt.Sprintf("[ %s ]", text))
}

func bluredButton(text string) string {
	return fmt.Sprintf("[ %s ]", blurredStyle.Render(text))
}

func setInputFocusedStyle(input *textinput.Model) {
	input.PromptStyle = focusedStyle
	input.TextStyle = focusedStyle
	input.Cursor.TextStyle = focusedStyle
}

func setInputNoStyle(input *textinput.Model) {
	input.PromptStyle = noStyle
	input.TextStyle = noStyle
	input.Cursor.TextStyle = noStyle
}

package model

import (
	"fmt"

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

	listSelectedItem = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color(SELECTED_COLOR))
)

func focusedButton(text string) string {
	return focusedStyle.Render(fmt.Sprintf("[ %s ]", text))
}

func bluredButton(text string) string {
	return fmt.Sprintf("[ %s ]", blurredStyle.Render(text))
}

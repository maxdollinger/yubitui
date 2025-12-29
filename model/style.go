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

	focusedButton    = focusedStyle.Render("[ add account ]")
	blurredButton    = fmt.Sprintf("[ %s ]", blurredStyle.Render("add account"))
	listSelectedItem = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color(SELECTED_COLOR))
)

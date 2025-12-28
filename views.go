package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var selectedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))

func AccountView(m *model) string {
	var s strings.Builder
	s.WriteString("Generate code for:\n\n")

	for i, choice := range m.accounts {

		line := choice
		if m.cursor == i {
			line = selectedStyle.Render(line)
		}

		fmt.Fprintf(&s, "%s\n", line)
	}

	s.WriteString("\nPress q to quit.\n")

	return s.String()
}

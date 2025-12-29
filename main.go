package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mdollinger/yubitui/model"
)

func main() {
	model := model.NewRootModel()
	defer model.Cleanup()

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %s", err)
	}
}

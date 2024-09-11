package main

import (
	"gomusik/ui"

	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.MainProgram{}
	if _, err := tea.NewProgram(&m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

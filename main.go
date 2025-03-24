package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  if _, err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Could not run the program")
    os.Exit(1)
  }
}

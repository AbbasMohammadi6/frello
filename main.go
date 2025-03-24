package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type task struct {
	title       string
}

func (t task) FilterValue() string {
	return t.title
}

func (t task) Title() string {
	return t.title
}

func (t task) Description() string {
  // TODO: there is an empty space under the title, because of this empty string, somehow remove it...
  // ...maybe create a new delegate and use the render method of it
	return ""
}

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.list.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func main() {
	items := []list.Item{
		task{"title 1"},
		task{"title 2"},
		task{"title 3"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Todos"

	m := model{list: l}
	tea.NewProgram(m, tea.WithAltScreen()).Run()
}

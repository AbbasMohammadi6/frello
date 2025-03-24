package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type task struct {
	title string
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

type status int

const (
	todo status = iota
	doing
	done
)

type model struct {
	lists   []list.Model
	fucosed status
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	doingStyle := lipgloss.NewStyle()
	doneStyle := lipgloss.NewStyle()

	return lipgloss.JoinHorizontal(
		0,
		m.lists[todo].View(), doingStyle.Render(m.lists[doing].View()), doneStyle.Render(m.lists[done].View()),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.lists[todo], cmd = m.lists[todo].Update(msg)

	return m, cmd
}

func main() {
	items := []list.Item{
		task{"title 1"},
		task{"title 2"},
		task{"title 3"},
	}

	todo := list.New(items, list.NewDefaultDelegate(), 20, 30)
	todo.Title = "Todos"
  todo.SetShowHelp(false)

	doing := list.New(items, list.NewDefaultDelegate(), 20, 30)
	doing.Title = "Doing"
  doing.SetShowHelp(false)

	done := list.New(items, list.NewDefaultDelegate(), 20, 30)
	done.Title = "Done"
  done.SetShowHelp(false)

	m := model{
		lists: []list.Model{todo, doing, done},
	}

	tea.NewProgram(m, tea.WithAltScreen()).Run()
}

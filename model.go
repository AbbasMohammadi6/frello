package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	lists        []list.Model
	fucosed      status
	loaded       bool
	windowWidth  int
	windowHeight int
}

var (
	divisor      = 3
	PaddingLeft  = 2
	PaddingRight = 2
	columnStyle  = lipgloss.NewStyle().PaddingRight(PaddingRight).PaddingLeft(PaddingLeft)
)

func (m *model) initialize() {
	items := []list.Item{
		task{"the quick brown fox jumped over the lazy dog"},
		task{"title 2"},
		task{"title 3"},
	}

	columnWidth := m.windowWidth/divisor - (PaddingLeft + PaddingRight)
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), columnWidth, m.windowHeight)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[todo].SetItems(items)
	m.lists[todo].Title = "Todos"

	m.lists[doing].SetItems(items)
	m.lists[doing].Title = "Doing"

	m.lists[done].SetItems(items)
	m.lists[done].Title = "Done"

	m.loaded = true
}

func initialModel() model {
	return model{
		loaded: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if !m.loaded {
		return "loading..."
	}

	columnWidth := m.windowWidth/divisor - (PaddingLeft + PaddingRight)
	columnStyle.Width(columnWidth)

	return lipgloss.JoinHorizontal(
		0,
		columnStyle.Render(m.lists[todo].View()),
		columnStyle.Render(m.lists[doing].View()),
		columnStyle.Render(m.lists[done].View()),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		{
			m.windowWidth = msg.Width
			m.windowHeight = msg.Height
			m.initialize() // Does this cause race condition???
		}
	}

	var cmd tea.Cmd
	m.lists[todo], cmd = m.lists[todo].Update(msg)

	return m, cmd
}

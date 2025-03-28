package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	lists        []list.Model
	tasks        [][]list.Item
	focused      status
	loaded       bool
	windowWidth  int
	windowHeight int
	showModal    bool
	textarea     textarea.Model
}

var (
	divisor      = 3 // TODO: get this from the number of columns
	PaddingLeft  = 2
	PaddingRight = 2
	BorderWidth  = 1
	ColumnStyle  = lipgloss.NewStyle().PaddingRight(PaddingRight).PaddingLeft(PaddingLeft)
	FocusedStyle = lipgloss.NewStyle().
			PaddingRight(PaddingRight-BorderWidth).
			PaddingLeft(PaddingLeft-BorderWidth).
			Border(lipgloss.RoundedBorder(), false, true, false, true). // TODO: figure a way to show border top and bottom without removing the titles
			BorderForeground(lipgloss.Color("62"))
)

func (m *model) initialize() {
	items := []list.Item{
		task{"the quick brown fox jumped over the lazy dog"},
		task{"title 2"},
		task{"title 3"},
	}

	m.tasks[todo] = items
	m.tasks[doing] = items
	m.tasks[done] = items

	columnWidth := m.windowWidth/divisor - (PaddingLeft + PaddingRight)
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), columnWidth, m.windowHeight)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// m.lists[todo].SetItems(items)
	m.lists[todo].Title = "Todos"

	// m.lists[doing].SetItems(items)
	m.lists[doing].Title = "Doing"

	// m.lists[done].SetItems(items)
	m.lists[done].Title = "Done"

	m.loaded = true

	m.textarea = textarea.New()
}

func initialModel() model {
	return model{
		loaded: false,
		tasks:  make([][]list.Item, 3),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if !m.loaded {
		return "loading..."
	}

	if m.showModal {
		return m.textarea.View()
	} else {

		m.lists[todo].SetItems(m.tasks[todo])
		m.lists[doing].SetItems(m.tasks[doing])
		m.lists[done].SetItems(m.tasks[done])

		columnWidth := m.windowWidth/divisor - (PaddingLeft + PaddingRight)
		ColumnStyle.Width(columnWidth)
		FocusedStyle.Width(columnWidth)

		todoView := m.lists[todo].View()
		doingView := m.lists[doing].View()
		doneView := m.lists[done].View()

		switch m.focused {
		case doing:
			return lipgloss.JoinHorizontal(
				0,
				ColumnStyle.Render(todoView),
				FocusedStyle.Render(doingView),
				ColumnStyle.Render(doneView),
			)

		case done:
			return lipgloss.JoinHorizontal(
				0,
				ColumnStyle.Render(todoView),
				ColumnStyle.Render(doingView),
				FocusedStyle.Render(doneView),
			)

		default:
			return lipgloss.JoinHorizontal(
				0,
				FocusedStyle.Render(todoView),
				ColumnStyle.Render(doingView),
				ColumnStyle.Render(doneView),
			)
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			if m.showModal {
				m.showModal = false
				return m, nil // nesseccary to prevent bubbletea exiting the program
			} else {
				return m, tea.Quit
			}

		case "ctrl+s":
			if m.showModal {
				value := m.textarea.Value()
				if len(value) > 0 {
          m.tasks[todo] = append(m.tasks[todo], task{title: value})
					m.showModal = false
          m.textarea.Reset()
				}
			}

		case "h", "left":
			if m.focused != todo && !m.showModal {
				m.focused--
			}

		case "l", "right":
			if m.focused != done && !m.showModal {
				m.focused++
			}

		case "ctrl+n":
			if !m.showModal {
				m.showModal = true
				m.textarea.Focus()
			}
		}

	case tea.WindowSizeMsg:
		{
			m.windowWidth = msg.Width
			m.windowHeight = msg.Height
			m.initialize() // Does this cause race condition???
		}
	}

	var cmd tea.Cmd
	if m.showModal {
		m.textarea, cmd = m.textarea.Update(msg)
	} else {
		m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	}

	return m, cmd
}

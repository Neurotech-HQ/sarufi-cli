package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sarufi-io/sarufi-golang-sdk"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	sarufi  sarufi.Application
	options list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() (s string) {
	return lipgloss.NewStyle().PaddingTop(1).Render(m.options.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)
	return m, cmd
}

func main() {
	m := model{}
	options := []list.Item{
		item{title: "Read from .env file", desc: "Read sarufi api key from .env file"},
		item{title: "Read from OS environment", desc: "Read credentials from OS environment"},
	}
	m.options = list.New(options, list.NewDefaultDelegate(), 50, 10)
	m.options.Title = "Welcome to Sarufi CLI"
	m.options.SetShowHelp(false)
	m.options.SetShowStatusBar(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("What did you do???")
		os.Exit(1)
	}

}

package app

import (
	// "fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "j", "down":
				if m.cursor > len(m.services)-1 {
					m.cursor = 4
				} else {
					m.cursor++
				}
			case "k", "up":
				if m.cursor <= 0 {
					m.cursor = 0
				} else {
					m.cursor--
				}
			case "enter", " ":
				m.chosen = true
				m.loaded = false
				m.choice = m.cursor - 1
				// updateChoices(m)
			}

		return m, nil
	}
	return m, nil
}

// func updateChoices(m model) tea.Cmd {
// 	// This function is called when the user has chosen a service
// 	if m.chosen {
// 		m.choice = m.cursor
// 	}

// 	return nil
// }
package tui

import (
    "github.com/PFNASS/pfnass-tui/cmd"

    tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyMsg:
        switch msg.String() {
        case "up":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        case "enter":
            switch m.cursor {
            case 0:
                // Call Cobra command directly
                cmd.Execute()
            case 1:
                return m, tea.Quit
            }
        }
    }

    return m, nil
}
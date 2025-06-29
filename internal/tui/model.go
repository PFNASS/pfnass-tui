package tui

import (
	tea  "github.com/charmbracelet/bubbletea"
)

type menuOption int

const (
    StartService menuOption = iota
    StopService
    Quit
)

type model struct {
    cursor int
    choices []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModel() model {
    return model{
        choices: []string{"Start Service", "Stop Service", "Quit"},
    }
}
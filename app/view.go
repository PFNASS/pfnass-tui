package app

import (
	"fmt"
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/PFNASS/pfnass-tui/app/home"
)


func (m model) View() string {
	// The he
	// s := "PFNASS Services\n\n"
	switch(m.choice) {
		case 0:
			return HomeView(m)
		case 1:
			return ProxyView(m)
		case 2:
			return IdPView(m)
		case 3:
			return MediaServerView(m)
		case 4:
			return K3SView(m)
		default:
			return fmt.Sprintf("Unknown service choice: %d", m.choice)
	}
}

package app
 
import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	return model{
		services:  []string{"Proxy", "IdP", "Media Server", "K3S-PFNASS"},
		// selected:  make(map[int]struct{}), // Initialize the map to track selected services
		choice: 0,
		chosen:  false,
		cursor:  0,
		loaded: false,
		quitting: false,
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func Start() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
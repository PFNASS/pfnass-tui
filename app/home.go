package app

import (
	"fmt"
)

func HomeView(m model) string {
	// The header
	var s string
	s += "PFNASS Services\n\n"
	s += "Choose a service to view:\n"
	for i, service := range m.services {
		// Is the cursor pointing at this service?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this service selected?
		// checked := " " // not selected
		// if _, ok := m.selected[i]; ok {
		// 	checked = "x" // selected!
		// }

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, service)
	}
	// The footer
	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
package app

import (
	// "fmt"
)

func MediaServerView(m model) string {
	// The header
	var s string
	s += "PFNASS Media Server Service\n\n"
	s += "This is the Media service for PFNASS.\n\n"

	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
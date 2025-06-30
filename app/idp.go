package app

import (
	// "fmt"
)

func IdPView(m model) string {
	// The header
	var s string
	s += "PFNASS IdP Service\n\n"
	s += "This is the IdP service for PFNASS.\n\n"

	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
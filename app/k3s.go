package app

import (
	// "fmt"
)

func K3SView(m model) string {
	// The header
	var s string
	s += "PFNASS K3S Service\n\n"
	s += "This is the k3s service for PFNASS.\n\n"

	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
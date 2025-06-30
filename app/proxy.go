package app

import (
	// "fmt"
)

func ProxyView(m model) string {
	// The header
	var s string
	s += "PFNASS Proxy Service\n\n"
	s += "This is the proxy service for PFNASS.\n\n"

	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
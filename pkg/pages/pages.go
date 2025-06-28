package pages

import (
	// "github.com/charmbracelet/bubbles"
// 	"github.com/charmbracelet/lipgloss"
)

type PageName string

const (
	MediaServer          PageName = "mediaServer"
	Docker         		 PageName = "docker"
	k3s					 PageName = "k3s"
	HomeAssistant		 PageName = "homeAssistant"
)
package tui

import (
    "fmt"
    "strings"
)

func (m model) View() string {
    var b strings.Builder
    b.WriteString("pfnass-tui - Homelab Controller\n\n")
    for i, choice := range m.choices {
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor
        }
        b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
    }
    b.WriteString("\nUse ↑/↓ to navigate, enter to select.")
    return b.String()
}

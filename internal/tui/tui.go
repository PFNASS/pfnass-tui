package tui

import (
    "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type tui struct {
	viewport viewport.Model
}

const content = `
# PFNASS Terminal User Interface Controller

## PFNASS Services

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
|    | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you'd like |
`

func newTUI() (*tui, error) {
	const width = 78

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	// We need to adjust the width of the glamour render from our main width
	// to account for a few things:
	//
	//  * The viewport border width
	//  * The viewport padding
	//  * The viewport margins
	//  * The gutter glamour applies to the left side of the content
	//
	const glamourGutter = 2
	glamourRenderWidth := width - vp.Style.GetHorizontalFrameSize() - glamourGutter

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(glamourRenderWidth),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &tui{
		viewport: vp,
	}, nil
}

func (e tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return e, tea.Quit
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}

func (e tui) Init() tea.Cmd {
	return nil
}

func (e tui) View() string {
	return e.viewport.View() + e.helpView()
}

func (e tui) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}

func Start() {
    model, err := newTUI()
    if err != nil {
        panic(err)
    }
    if _, err := tea.NewProgram(model).Run(); err != nil {
        panic(err)
    }
    // if err := p.Start(); err != nil {
    //     panic(err)
    // }
}
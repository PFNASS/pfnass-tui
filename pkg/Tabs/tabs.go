package tabs

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tabModel struct {
	Tabs       []string
	TabContent []string
	activeTab  int
}

func (tab tabModel) Init() tea.Cmd {
	return nil
}

func (tab tabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return tab, tea.Quit
		case "right", "l", "n", "tab":
			tab.activeTab = min(tab.activeTab+1, len(tab.Tabs)-1)
			return tab, nil
		case "left", "h", "p", "shift+tab":
			tab.activeTab = max(tab.activeTab-1, 0)
			return tab, nil
		}
	}

	return tab, nil
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#91ffff", Dark: "#08a4a7"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func (tabModel tabModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range tabModel.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(tabModel.Tabs)-1, i == tabModel.activeTab
		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(tabModel.TabContent[tabModel.activeTab]))
	return docStyle.Render(doc.String())
}

func main() {

	tabs := []string{"Media Server", "Docker", "k3s", "Home Assistant"}
	tabContent := []string{"Start", "Containers", "k3s containers", "Status"}
	tabModel := tabModel{Tabs: tabs, TabContent: tabContent}
	r := tea.NewProgram(tabModel)
	if _, err := r.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
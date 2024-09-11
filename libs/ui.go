package libs

import (
	"github.com/charmbracelet/lipgloss"
)

// Const values
var (
	GapHeight = 4
	GapWidth  = 2
)

// Color
var (
	HighlightColor         = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	HighlightChooseDirFile = lipgloss.NewStyle().
   		Background(lipgloss.Color("#7d56f3")). // Orange background for example
    	Foreground(lipgloss.Color("white")). // Black text for example
    	Bold(true)
)

var (
	InactiveTabBorder      = tabBorderWithBottom("┴", "─", "┴")
	ActiveTabBorder        = tabBorderWithBottom("╯", " ", "╰")
	DocStyle               = lipgloss.NewStyle()
	InactiveTabStyle       = lipgloss.NewStyle().Border(InactiveTabBorder).BorderForeground(HighlightColor).Padding(0, 1)
	ActiveTabStyle         = InactiveTabStyle.Border(ActiveTabBorder, true)
	WindowStyle            = lipgloss.NewStyle().BorderForeground(HighlightColor).Align(lipgloss.Center).Border(lipgloss.RoundedBorder()).UnsetBorderTop()
	BorderOnlyBottom = lipgloss.Border{
        TopLeft:     " ", // Remove top left border
        Top:         " ", // Remove top border
        TopRight:    " ", // Remove top right border
		BottomLeft:  "─", // Keep bottom border
        Bottom:      "─", // Keep bottom border
        BottomRight: "╮", // Keep bottom right border
        Left:        " ", // Keep left border
        Right:       " ", // Keep right border
    }
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

package tabs

import (
	"gomusik/libs"

	"github.com/charmbracelet/lipgloss"
)

func AddEmptyBorderBottomTab(renderedTabs *[]string, termWidth int, tabWidth int) {
	emptyString := ""
	style := lipgloss.NewStyle().Border(libs.BorderOnlyBottom).BorderForeground(libs.HighlightColor)
	if termWidth > tabWidth {
		// Fix the border when terminal width near tab length
		if termWidth == tabWidth + 1 {
			newBorder := libs.BorderOnlyBottom
			newBorder.BottomLeft = "╮"
			style = style.Border(newBorder)
		}
		for i := 0; i < termWidth - libs.GapWidth - tabWidth; i++ {
			emptyString += " "
		}
		*renderedTabs = append(*renderedTabs, style.Render(emptyString))
	} 
}

package tabs

import (
	"gomusik/libs"

	"github.com/charmbracelet/lipgloss"
)

func AddMainTab(renderedTabs *[]string, tabs *[]string, activeTab, termWidth, tabWidth int) {
	for i, t := range *tabs {
		style := lipgloss.Style{}
		isFirst, isLast, isActive := i == 0, i == len(*tabs)-1, i == activeTab

		// Set border if is active or inactive
		if isActive {
			style = libs.ActiveTabStyle
		} else {
			style = libs.InactiveTabStyle
		}
		border, _, _, _, _ := style.GetBorder()

		// Change first border bottom left 
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		}

		// Change last boder bottom right
		if isLast && isActive && termWidth == tabWidth {
			border.BottomRight = "│"
		} else if isLast && !isActive && termWidth == tabWidth {
			border.BottomRight = "┤"
		}

		style = style.Border(border)
		*renderedTabs = append(*renderedTabs, style.Render(t))
	}
}

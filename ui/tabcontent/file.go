package tabcontent

import (
	"gomusik/libs"

	"github.com/charmbracelet/lipgloss"
)

func GetFile(
	list       []string,
	listCount  int,
	listChoose int,
	filePass   int,
	currentDir string,
	termWidth  int,
	termHeight int,
) string {
	// Trim down list based on height terminal and gaps
	//pass := filePass
	for i := 0; i <= listCount - (termHeight - libs.GapHeight); i++ {
		if filePass > 0 {
			list = list[1:]
			filePass -= 1
			continue
		}

		if listChoose + libs.GapHeight < termHeight {
			list = list[:len(list)-1]
		}
	} 

	// Add item from list to stringContent
	stringContent := ""
	for i, item := range list {
		if i == listChoose {
			stringContent += libs.HighlightChooseDirFile.Render(item) + "\n"
		} else {
			stringContent += item + "\n"
		}
	}
	stringContent = stringContent[:len(stringContent)-1] // Remove new line

	return libs.WindowStyle.
		Width(termWidth - libs.GapWidth).
		Height(termHeight - libs.GapHeight).
		Align(lipgloss.Left).
		Render(stringContent)
}

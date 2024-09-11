package ui

import (
	"gomusik/libs"
	"gomusik/ui/tabcontent"
	"gomusik/ui/tabs"

	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainProgram struct {
	// UI
	Refresh    bool
	Width      int
	Height     int

	// Tabs
	Tabs       []string
	TabContent []string
	ActiveTab  int

	// Files
	CurrentDir string
	FileList   []string
	FileCount  int
	FileChoose int
	FilePass   int
}

func (m *MainProgram) switchFileUp() {
	if m.FileChoose <= 0 {
		m.FileChoose = 0
		m.FilePass -= 1
		if m.FilePass < 0 {
			m.FilePass = 0
		}
		return
	}

	m.FileChoose -= 1
}

func (m *MainProgram) switchFileDown() {
	// Create maxChoose to limit FileChoose
	maxChoose := 0
	if m.FileCount >= m.Height - libs.GapHeight {
		maxChoose = m.Height - libs.GapHeight - 1
	} else {
		maxChoose = m.FileCount
	}

	// Update FileChoose and FilePass
	if m.FileChoose >= maxChoose {
		m.FilePass += 1
		if m.FilePass + m.FileChoose >= m.FileCount {
			m.FilePass = m.FileCount - maxChoose
			m.FileChoose = maxChoose
		}
		return
	}
	m.FileChoose += 1
}

func (m *MainProgram) updateDir() {
	// Get list and count
	list, count, _ := libs.ListDirectory(m.CurrentDir)
	m.FileList = list
	m.FileCount = count
	m.FileChoose = 0
	m.FilePass = 0
}

func (m *MainProgram) Init() tea.Cmd {
	// Get current dir, list and count
	dir, _ := os.Getwd()
	list, count, _ := libs.ListDirectory(dir)

	// Initialize value
	m.Tabs = []string{"Player", "Playlist", "File Manager", "Settings"}
	m.CurrentDir = dir
	m.FileList = list
	m.FileCount = count
	m.Refresh = true

	return nil
}

func (m *MainProgram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Refresh = true
		if m.FileCount >= m.Height - libs.GapHeight {
			maxChoose := m.Height - libs.GapHeight - 1
			m.FileChoose = maxChoose
			m.FilePass = m.FileCount - maxChoose
		} else {
			m.FilePass = 0
			m.FileChoose = m.FileCount
		}
		return m, nil

	case tea.KeyMsg:
		m.Refresh = true

		switch keypress := msg.String(); keypress {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			m.ActiveTab = libs.SwitchToMax(m.ActiveTab, len(m.Tabs)-1)
			return m, nil

		case "shift+tab":
			m.ActiveTab = libs.SwitchToMin(m.ActiveTab, 0)
			return m, nil

		case "up":
			if  m.Tabs[m.ActiveTab] == "File Manager" {
				m.switchFileUp()
			}
			return m, nil

		case "down":
			if  m.Tabs[m.ActiveTab] == "File Manager" {
				m.switchFileDown()
			}
			return m, nil

		case "enter":
			if m.Tabs[m.ActiveTab] == "File Manager" {
				if m.FileList[m.FileChoose] == ".." {
					// Remove last directory from m.CurrentDir
					lastIndex := strings.LastIndex(m.CurrentDir, "/")
					if lastIndex != -1 {
						m.CurrentDir = m.CurrentDir[:lastIndex+1]
					}
					m.CurrentDir = m.CurrentDir[:len(m.CurrentDir)-1]

					m.updateDir()
				} else if strings.Contains(m.FileList[m.FileChoose + m.FilePass], "/") {
					// Add directory to m.CurrentDir
					newDir := ""
					currentFile := m.FileList[m.FileChoose + m.FilePass]
					newDir = currentFile[:len(currentFile)-1]
					m.CurrentDir += ("/" + newDir)

					m.updateDir()
				}
			}
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}

func (m *MainProgram) View() string {
	// Get tab length
	tabWidth := 0
	for _, item := range m.Tabs {
		tabWidth += len(item) + 4
	}

	// Get render tabs
	renderedTabs := []string{}
	tabs.AddMainTab(&renderedTabs, &m.Tabs, m.ActiveTab, m.Width, tabWidth)
	tabs.AddEmptyBorderBottomTab(&renderedTabs, m.Width, tabWidth)

	// Add all renderedTabs to row
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	// Add tab content based on main tab
	content := ""
	if m.Refresh {
		if m.Tabs[m.ActiveTab] == "File Manager" {
			content = tabcontent.GetFile(
				m.FileList,
				m.FileCount,
				m.FileChoose,
				m.FilePass,
				m.CurrentDir,
				m.Width,
				m.Height,
			)
		} else { 
			content = libs.WindowStyle.
				Width(m.Width - libs.GapWidth).
				Height(m.Height - libs.GapHeight).
				Align(lipgloss.Left).
				Render("Nothing here! :D")
		}
		m.Refresh = false
	}

	// Generate UI
	doc := strings.Builder{}
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(content)

	return libs.DocStyle.Render(doc.String())
}

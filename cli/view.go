package cli

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle   = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	faintStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	enumerateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m Model) View() string {
	s := appNameStyle.Render(
		`----^^^^^^^^^^^^^         ------^^^^^^^^^^^^
		-----^^^^^^^^^^^^^NOTES APP------^^^^^^^^^^^^
		_____________________________________________
		`) + "\n\n"
	if m.state == listView {
		for i, n := range m.notes {
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}
			shortBody := strings.ReplaceAll(n.Body, "\n", "")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += enumerateStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}
		s += faintStyle.Render("n-  new notes,\n q- quit")
	}
	return s
}

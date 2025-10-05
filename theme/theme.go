package theme

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/lipgloss"
	"github.com/kaputi/snippets/logger"
	"golang.org/x/term"
)

var colors = map[string]string{
	"foreground": "#FFFFFF",
	"background": "#000000",
	"primary":    "#FFA500",
	"secondary":  "#00FFFF",
	"accent":     "#FF00FF",
	"selected":   "#008000",
}

var (
	termWidth  int
	termHeight int

	navPanelWidthPercent float32 = 0.3
	navPanelWidth        int
	navPanelHeight       int

	contentPanelWidth  int
	contentPanelHeight int

	panelStyle = lipgloss.NewStyle().
			Margin(0, 0).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder())

	leftColumnStyle  lipgloss.Style
	rightColumnStyle lipgloss.Style

	navPanelStyle     lipgloss.Style
	contentPanelStyle lipgloss.Style

	focusedNavPanelStyle     lipgloss.Style
	focusedContentPanelStyle lipgloss.Style
)

func init() {
	updateStyle()
	go watchWindowResize()
}

func updateStyle() {
	width, height := getTermSize()
	accentColor, _ := Color("accent")

	termWidth = width
	termHeight = height

	leftColumnWidth := int(float32(termWidth) * navPanelWidthPercent)

	leftColumnStyle = lipgloss.NewStyle().
		Width(leftColumnWidth).
		Height(termHeight)

	rightColumnStyle = lipgloss.NewStyle().
		Width(termWidth - leftColumnWidth).
		Height(termHeight)

	navPanelWidth = int(leftColumnWidth)
	navPanelHeight = int((float32(termHeight) - 2) * 0.3) // -2 for the border

	contentPanelWidth = termWidth - navPanelWidth
	contentPanelHeight = navPanelHeight * 3

	navPanelStyle = panelStyle.Width(navPanelWidth).Height(navPanelHeight)
	contentPanelStyle = panelStyle.Width(contentPanelWidth).Height(contentPanelHeight)

	focusedNavPanelStyle = navPanelStyle.
		BorderForeground(lipgloss.Color(accentColor))

	focusedContentPanelStyle = contentPanelStyle.
		BorderForeground(lipgloss.Color(accentColor))

}

func getTermSize() (int, int) {
	width, height, err := term.GetSize(0) // Get terminal size (0 is stdin)
	if err != nil {
		logger.Log("Error getting terminal size")
	}
	return width, height
}

func watchWindowResize() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)

	for range sig {
		updateStyle()
	}
}

func Color(name string) (string, error) {
	if color, exists := colors[name]; exists {
		return color, nil
	}

	return "", fmt.Errorf("color %s not found", name)
}

func NavPanel() lipgloss.Style {
	return navPanelStyle
}

func ContentPanel() lipgloss.Style {
	return contentPanelStyle
}

func FocusedNavPanel() lipgloss.Style {
	return focusedNavPanelStyle
}

func FocusedContentPanel() lipgloss.Style {
	return focusedContentPanelStyle
}

func LeftColumn() lipgloss.Style {
	return leftColumnStyle
}

func RightColumn() lipgloss.Style {
	return rightColumnStyle
}

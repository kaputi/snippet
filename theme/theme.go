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
	navPanelWidthPercent float32 = 0.3

	TermWidth  int
	TermHeight int

	NavPanelWidth int

	LangPanelHeight    int
	TreePanelHeight    int
	SnippetPanelHeight int

	ContentPanelWidth  int
	ContentPanelHeight int

	PanelStyle        lipgloss.Style
	LangPanelStyle    lipgloss.Style
	TreePanelStyle    lipgloss.Style
	SnippetPanelStyle lipgloss.Style
	ContentPanelStyle lipgloss.Style
)

func init() {
	updateStyle()
	go watchWindowResize()
}

func updateStyle() {
	width, height := getTermSize()

	TermWidth = width
	TermHeight = height

	leftColumnWidth := int(float32(TermWidth) * navPanelWidthPercent)
	navPanelHeight := int((float32(TermHeight - 2)) * 0.3) // -2 for the border
	navPanelWidth := leftColumnWidth

	LangPanelHeight = navPanelHeight
	TreePanelHeight = navPanelHeight
	SnippetPanelHeight = TermHeight - 2 - LangPanelHeight - TreePanelHeight

	ContentPanelWidth = TermWidth - navPanelWidth
	ContentPanelHeight = navPanelHeight * 3

	PanelStyle = lipgloss.NewStyle().
		Margin(0, 0).
		Padding(0, 1).
		Border(lipgloss.RoundedBorder())

	LangPanelStyle = PanelStyle.Width(navPanelWidth).Height(LangPanelHeight)
	TreePanelStyle = PanelStyle.Width(navPanelWidth).Height(TreePanelHeight)
	SnippetPanelStyle = PanelStyle.Width(navPanelWidth).Height(SnippetPanelHeight)
	ContentPanelStyle = PanelStyle.Width(ContentPanelWidth).Height(ContentPanelHeight)
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

func FocusPanel(inputStyle lipgloss.Style) lipgloss.Style {
	accentColor, _ := Color("accent")
	return inputStyle.BorderForeground(lipgloss.Color(accentColor))
}

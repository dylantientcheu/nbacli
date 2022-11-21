package styles

import (
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var EmphasisText = lipgloss.NewStyle().Bold(true)
var PrimaryText = lipgloss.NewStyle().Bold(true)
var SecondaryText = lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))

var LiveText = lipgloss.NewStyle().Background(lipgloss.AdaptiveColor{Light: "#ef2929", Dark: "#ef2929"}).Foreground(lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"}).Bold(true)
var FinalText = lipgloss.NewStyle().Background(lipgloss.Color("#9356DF")).Foreground(lipgloss.Color("#ffffff")).Bold(true)
var DescText = lipgloss.NewStyle().Foreground(lipgloss.Color("#818181"))
var ScoreText = lipgloss.NewStyle().Background(lipgloss.AdaptiveColor{Light: "214", Dark: "#181818"}).Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "214"})

const (
	magentaDark = "200"
	yellowDark  = "214"
	blueDark    = "33"
	pinkDark    = "219"

	logoBgDark           = "#0f1429"
	headerBgDark         = "#2d3454"
	unselectedItemFgDark = "251"
	paginatorBgDark      = logoBgDark
	selectedPageFgDark   = unselectedItemFgDark
	unselectedPageFgDark = "239"

	magentaLight = magentaDark
	yellowLight  = "208"
	blueLight    = blueDark
	pinkLight    = pinkDark

	logoBgLight           = "252"
	headerBgLight         = "254"
	unselectedItemFgLight = "235"
	paginatorBgLight      = logoBgLight
	selectedPageFgLight   = unselectedItemFgLight
	unselectedPageFgLight = "247"
)

func GetSpinner() spinner.Spinner {
	normal := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: selectedPageFgLight, Dark: selectedPageFgDark}).
		Background(lipgloss.AdaptiveColor{Light: unselectedPageFgLight, Dark: unselectedPageFgDark}).
		Faint(false)

	color := normal.Copy()

	magenta := lipgloss.AdaptiveColor{Light: magentaLight, Dark: magentaDark}
	yellow := lipgloss.AdaptiveColor{Light: yellowLight, Dark: yellowDark}
	blue := lipgloss.AdaptiveColor{Light: blueLight, Dark: blueDark}

	return spinner.Spinner{
		Frames: []string{
			normal.Render("fetching" + color.Foreground(magenta).Render(".") + color.Foreground(yellow).Render(".") + color.Foreground(blue).Render(".")),
			normal.Render("fetching   "),
			normal.Render("fetching" + color.Foreground(magenta).Render(".") + normal.Render("  ")),
			normal.Render("fetching" + color.Foreground(magenta).Render(".") + color.Foreground(yellow).Render(".") + normal.Render(" ")),
		},
		FPS: 1500 * time.Millisecond,
	}
}

func LiveStyle() string {
	return LiveText.Render(" LIVE ")
}

func FinalStyle() string {
	return FinalText.Render(" FINAL ")
}

func ScoreStyle(homeScore int, awayScore int) string {
	return ScoreText.Render(" " + strconv.Itoa(homeScore) + " - " + strconv.Itoa(awayScore) + " ")
}

func DescStyle(desc string) string {
	return DescText.Render(desc)
}

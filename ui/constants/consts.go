package constants

import (
	"strconv"

	"github.com/evertras/bubble-table/table"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* CONSTANTS */

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(Secondary)

var (
	// P the current tea program
	P *tea.Program

	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg

	CustomTableBorder = table.Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "┬",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "┴",
		InnerJunction:  "┼",

		InnerDivider: "│",
	}

	LiveText = lipgloss.NewStyle().Background(lipgloss.AdaptiveColor{Light: "#ef2929", Dark: "#ef2929"}).Foreground(lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"}).Bold(true)

	FinalText = lipgloss.NewStyle().Background(lipgloss.Color("#9356DF")).Foreground(lipgloss.Color("#ffffff")).Bold(true)
	DescText  = lipgloss.NewStyle().Foreground(lipgloss.Color("#818181"))

	ScoreText = lipgloss.NewStyle().Background(lipgloss.AdaptiveColor{Light: "214", Dark: "#181818"}).Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "214"})

	Accent       = lipgloss.AdaptiveColor{Light: "#5b1b7b", Dark: "#5b1b7b"}
	AccentDarker = lipgloss.AdaptiveColor{Light: "#5b1b7b", Dark: "#5b1b7b"}
	Secondary    = lipgloss.AdaptiveColor{Light: "#ed2265", Dark: "#ed2265"}
	Tertiary     = lipgloss.AdaptiveColor{Light: "#f69053", Dark: "#f69053"}

	activeTabBorder = lipgloss.Border{
		Bottom:      "─",
		Top:         "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Bottom: "─",
		// Top:         "─",
		// Left:        "│",
		// Right:       "│",
		// TopLeft:     "╭",
		// TopRight:    "╮",
		BottomLeft:  "─",
		BottomRight: "─",
	}

	TabStyle = lipgloss.NewStyle().
			Border(tabBorder, true).
			BorderForeground(Accent).
			Background(Accent).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	ActiveTabStyle = lipgloss.NewStyle().
			Border(activeTabBorder, true).
			BorderForeground(Secondary).
			Background(Secondary).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1)

	BleedSpaceWidth = 4
)

/* STYLING */

// DocStyle styling for viewports
var DocStyle = lipgloss.NewStyle().Margin(1, 2)

// TitleStyle styling for titles
var TitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Background(Accent).Padding(0, 2)

// HelpStyle styling for help context menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

// ErrStyle provides styling for error messages
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

// AlertStyle provides styling for alert messages
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render

type keymap struct {
	Enter     key.Binding
	Yesterday key.Binding
	Tomorrow  key.Binding
	Back      key.Binding
	Quit      key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↲/enter", "select"),
	),
	Yesterday: key.NewBinding(
		key.WithKeys("i", "left"),
		key.WithHelp("←/i", "previous day"),
	),
	Tomorrow: key.NewBinding(
		key.WithKeys("o", "right"),
		key.WithHelp("→/o", "next day"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

package constants

import (
	"github.com/dylantientcheu/nbacli/nba"
	"github.com/evertras/bubble-table/table"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* CONSTANTS */

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(Accent)

var (
	// P the current tea program
	P *tea.Program

	Gm *nba.BoxScoreRepository
	Sb *nba.ScoreboardRepository
	St *nba.StandingsRepository

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

	Accent = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	TabStyle = lipgloss.NewStyle().
			Border(tabBorder, true).
			BorderForeground(Accent).
			Padding(0, 1)

	ActiveTabStyle = lipgloss.NewStyle().
			Border(activeTabBorder, true).
			BorderForeground(Accent).
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

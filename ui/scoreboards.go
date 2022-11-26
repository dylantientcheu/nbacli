package ui

import (
	"fmt"
	"nba-cli/nba"
	"nba-cli/ui/constants"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

// SelectMsg the message to change the view to the selected entry
type SelectMsg struct {
	ActiveScorebardID uint
}

const (
	nav mode = iota
	edit
)

type Model struct {
	mode        mode
	list        list.Model
	currentDate time.Time
	quitting    bool
	gameview    bool
}

func InitScoreboard(date time.Time) tea.Model {
	// TODO: add spinners
	items := newScoreboardList(constants.Sb, date)
	m := Model{mode: nav, currentDate: date, list: list.NewModel(items, list.NewDefaultDelegate(), 8, 8)}
	if constants.WindowSize.Height != 0 {
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-1)
	}
	m.list.Title = "NBA Games - " + m.currentDate.Format("Monday, 2 Jan 06")
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Tomorrow,
			constants.Keymap.Yesterday,
			constants.Keymap.Back,
		}
	}
	return m
}

func newScoreboardList(scbrd *nba.ScoreboardRepository, date time.Time) []list.Item {
	games := scbrd.GetGames(date)
	return gamesToItems(games)
}

// Init run any intial IO on program start
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Yesterday):
			// TODO: add spinners
			var previousDay nba.ScoreboardRepository
			m.currentDate = m.currentDate.AddDate(0, 0, -1)
			games := previousDay.GetGames(m.currentDate)
			items := gamesToItems(games)
			m.list.Title = "NBA Games - " + m.currentDate.Format("Monday, 2 Jan 06")
			m.list.SetItems(items)
		case key.Matches(msg, constants.Keymap.Tomorrow):
			// TODO: add spinners
			var nextDay nba.ScoreboardRepository
			m.currentDate = m.currentDate.AddDate(0, 0, 1)
			games := nextDay.GetGames(m.currentDate)
			items := gamesToItems(games)
			m.list.Title = "NBA Games - " + m.currentDate.Format("Monday, 2 Jan 06")
			m.list.SetItems(items)
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			m.gameview = true
			activeGame := m.list.SelectedItem().(nba.BoxScoreSummary)
			fmt.Printf("%#v", activeGame.GameId)
			gameView := InitGameView(activeGame.GameId)
			return gameView.Update(constants.WindowSize)
		default:
			m.list, cmd = m.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	return constants.DocStyle.Render(m.list.View() + "\n")
}

func gamesToItems(games []nba.BoxScoreSummary) []list.Item {
	items := make([]list.Item, len(games))
	for i, proj := range games {
		items[i] = list.Item(proj)
	}
	return items
}

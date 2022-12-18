package ui

import (
	"time"

	"github.com/dylantientcheu/nbacli/nba"
	"github.com/dylantientcheu/nbacli/ui/constants"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type mode int

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
	items := newScoreboardList(nba.Sb, date)
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
			var previousDay nba.ScoreboardRepository
			m.currentDate = m.currentDate.AddDate(0, 0, -1)
			games := previousDay.GetGames(m.currentDate)
			items := gamesToItems(games)
			m.list.Title = "NBA Games - " + m.currentDate.Format("Monday, 2 Jan 06")
			m.list.SetItems(items)
		case key.Matches(msg, constants.Keymap.Tomorrow):
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
			gameView := InitGameView(activeGame.GameId, activeGame, m)
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

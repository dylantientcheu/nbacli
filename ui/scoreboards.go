package ui

import (
	"fmt"
	"nba-cli/nba"
	"nba-cli/ui/constants"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
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
	create
)

type Model struct {
	mode     mode
	list     list.Model
	quitting bool
}

func InitScoreboard(date time.Time) tea.Model {
	items := newScoreboardList(constants.Sb, date)
	m := Model{mode: nav, list: list.NewModel(items, list.NewDefaultDelegate(), 8, 8)}
	if constants.WindowSize.Height != 0 {
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-1)
	}
	m.list.Title = "NBA Games"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Create,
			constants.Keymap.Rename,
			constants.Keymap.Delete,
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
		case key.Matches(msg, constants.Keymap.Create):
			m.mode = create
			cmd = textinput.Blink
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			activeGame := m.list.SelectedItem().(nba.Game)
			fmt.Printf("%#v", activeGame.GameId)
			return m, tea.Quit // enter the game id
			// entry := InitEntry(constants.Er, activeProject.ID, constants.P)
			// return entry.Update(constants.WindowSize)
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

// TODO: use generics
// gamesToItems convert []model.Project to []list.Item
func gamesToItems(games []nba.Game) []list.Item {
	items := make([]list.Item, len(games))
	for i, proj := range games {
		items[i] = list.Item(proj)
	}
	return items
}

func (m Model) getActiveGameID() string {
	items := m.list.Items()
	activeItem := items[m.list.Index()]
	return activeItem.(nba.Game).GameId
}

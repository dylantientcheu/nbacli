package ui

import (
	"nba-cli/nba"
	"nba-cli/styles"
	"nba-cli/ui/constants"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type EntryModel struct {
	table        table.Model
	activeGameID string
}

func (m EntryModel) Init() tea.Cmd { return nil }

func (m EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m EntryModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func InitGameView(activeGameID string) *EntryModel {
	columns := []table.Column{
		{Title: "POS", Width: 2},
		{Title: "NAME", Width: 16},
		{Title: "MIN", Width: 6},
		{Title: "FG", Width: 6},
		{Title: "3PT", Width: 3},
		{Title: "FT", Width: 3},
		{Title: "REB", Width: 3},
		{Title: "AST", Width: 3},
		{Title: "STL", Width: 3},
		{Title: "BLK", Width: 3},
		{Title: "TO", Width: 3},
		{Title: "+/-", Width: 4},
		{Title: "PTS", Width: 3},
	}

	rows := newStatsBoard(constants.Gm, activeGameID)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(12),
	)

	// TODO: Add more styles
	// - Game Score
	// - Team Name
	// - Team Color (optional)
	// - Logo (optional)
	// - Separate Benchers from Starters
	// - Add a header for each section
	// - Separate teams by tables
	// - Handle non active games

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := EntryModel{t, activeGameID}
	return &m
}

func newStatsBoard(game *nba.BoxScoreRepository, gameID string) []table.Row {
	// testId := "0022200248"
	gameStats := game.GetSingleGameStats(gameID)
	return statsToRows(gameStats)
}

func statsToRows(gameStats []nba.GameStat) []table.Row {
	var rows []table.Row
	areBenchers := false

	for _, stat := range gameStats {
		// format plus minus
		plusMinus := "0"
		if stat.PlusMinus > 0 {
			plusMinus = "+" + strconv.FormatInt(stat.PlusMinus, 10)
		}
		rows = append(rows, table.Row{
			stat.StartPosition, // POS - C
			stat.PlayerName,    // NAME - LeBron James
			stat.Min,           // MIN - 36:00
			strconv.FormatInt(stat.Fgm, 10) + "-" + strconv.FormatInt(stat.Fga, 10), // FG - 10-20
			strconv.FormatInt(stat.Fg3M, 10),                                        // 3PT - 2-5
			strconv.FormatInt(stat.Ftm, 10),                                         // FT - 10-10
			strconv.FormatInt(stat.Reb, 10),                                         // REB - 10
			strconv.FormatInt(stat.AST, 10),                                         // AST - 10
			strconv.FormatInt(stat.Stl, 10),                                         // STL - 10
			strconv.FormatInt(stat.Blk, 10),                                         // BLK - 10
			strconv.FormatInt(stat.To, 10),                                          // TO - 10
			plusMinus,                                                               // +/- - 10
			strconv.FormatInt(stat.Pts, 10),                                         // PTS - 10
		})

		if (stat.StartPosition == "") && !areBenchers {
			rows = append(rows, table.Row{"", styles.ScoreText.Render("BENCH"), "", "", "", "", "", "", "", "", "", "", ""})
			areBenchers = true
		}
	}
	return rows
}

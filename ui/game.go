package ui

import (
	"strconv"

	"github.com/dylantientcheu/nbacli/nba"
	"github.com/dylantientcheu/nbacli/ui/constants"
	"github.com/dylantientcheu/nbacli/ui/gameboard/scoretext"

	"github.com/evertras/bubble-table/table"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(constants.Accent)

type GameModel struct {
	table                 table.Model
	activeGameID          string
	activeGame            nba.BoxScoreSummary
	previousModel         Model
	help                  help.Model
	width, height, margin int
}

var gameKM = GameKM{
	Down:     key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "highlight next row")),
	Up:       key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "highlight previous row")),
	Previous: key.NewBinding(key.WithKeys("esc", "q"), key.WithHelp("q/esc", "back to games list")),
}

func (m *GameModel) recalculateTable() {
	m.table = m.table.WithTargetWidth(m.width)
}

func (m GameModel) Init() tea.Cmd { return nil }

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			// return to previous page
			return m.previousModel, tea.Batch()
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			// TODO: to player view
			return m, tea.Batch()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.recalculateTable()
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m GameModel) View() string {
	table := m.table.View() + "\n"

	helpContainer := lipgloss.NewStyle().
		SetString(m.help.View(gameKM)).
		Width(m.width).
		Align(lipgloss.Center).
		PaddingTop(1).
		String()

	return scoretext.RenderScoreText(m.activeGame.ArenaName, m.activeGame.GameDate, m.activeGame.HomeTeamScore, m.activeGame.VisitorTeamScore, m.activeGame.HomeTeamName, m.activeGame.VisitorTeamName) + table + helpContainer
}

func InitGameView(activeGameID string, activeGame nba.BoxScoreSummary, previousModel Model) *GameModel {
	columns := []table.Column{
		table.NewFlexColumn("POS", "POS", 2),
		table.NewFlexColumn("NAME", "NAME", 10),
		table.NewFlexColumn("MIN", "MIN", 6),
		table.NewFlexColumn("FG", "FG", 6),
		table.NewFlexColumn("3PT", "3PT", 3),
		table.NewFlexColumn("FT", "FT", 3),
		table.NewFlexColumn("REB", "REB", 3),
		table.NewFlexColumn("AST", "AST", 3),
		table.NewFlexColumn("STL", "STL", 3),
		table.NewFlexColumn("BLK", "BLK", 3),
		table.NewFlexColumn("TO", "TO", 3),
		table.NewFlexColumn("+/-", "+/-", 4),
		table.NewFlexColumn("PTS", "PTS", 3),
	}

	rows := newStatsBoard(nba.Gm, activeGameID)

	t := table.New(columns).WithRows(rows).
		Focused(true).
		Border(constants.CustomTableBorder).WithBaseStyle(baseStyle).WithPageSize(constants.WindowSize.Height / 3)

	m := GameModel{t, activeGameID, activeGame, previousModel, help.New(), constants.WindowSize.Height, constants.WindowSize.Width, 3}
	return &m
}

func newStatsBoard(game *nba.BoxScoreRepository, gameID string) []table.Row {
	gameStats := game.GetSingleGameStats(gameID)
	return statsToRows(gameStats)
}

func statsToRows(gameStats []nba.GameStat) []table.Row {
	var rows []table.Row
	areBenchers := false

	rows = append(rows, table.NewRow(renderTeamRow("AWAY TEAM")).
		WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).
			Background(constants.Secondary)))

	for idx, stat := range gameStats {
		// format plus minus
		plusMinus := "0"
		if stat.PlusMinus > 0 {
			plusMinus = "+" + strconv.FormatInt(stat.PlusMinus, 10)
		} else {
			plusMinus = strconv.FormatInt(stat.PlusMinus, 10)
		}

		if (stat.StartPosition == "") && !areBenchers {
			rows = append(rows, table.NewRow(
				renderBenchRow(),
			).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Background(lipgloss.AdaptiveColor{Light: "214", Dark: "#181818"})))
			areBenchers = true
		}

		rows = append(rows, table.NewRow(
			table.RowData{
				"POS":  stat.StartPosition,
				"NAME": stat.PlayerName,
				"MIN":  stat.Min,
				"FG":   strconv.FormatInt(stat.Fgm, 10) + "-" + strconv.FormatInt(stat.Fga, 10),
				"3PT":  strconv.FormatInt(stat.Fg3M, 10),
				"FT":   strconv.FormatInt(stat.Ftm, 10),
				"REB":  strconv.FormatInt(stat.Reb, 10),
				"AST":  strconv.FormatInt(stat.AST, 10),
				"STL":  strconv.FormatInt(stat.Stl, 10),
				"BLK":  strconv.FormatInt(stat.Blk, 10),
				"TO":   strconv.FormatInt(stat.To, 10),
				"+/-":  plusMinus,
				"PTS":  strconv.FormatInt(stat.Pts, 10),
			},
		))
		if stat.StartPosition != "" {
			areBenchers = false
		}

		if idx < len(gameStats)-1 && gameStats[idx].TeamID != gameStats[idx+1].TeamID {
			rows = append(rows, table.NewRow(renderTeamRow("HOME TEAM")).WithStyle(lipgloss.NewStyle().
				AlignHorizontal(lipgloss.Center).
				Background(constants.Secondary)))
		}
	}
	return rows
}

func renderBenchRow() table.RowData {
	return table.RowData{
		"POS":  "",
		"NAME": table.NewStyledCell("B E N C H", lipgloss.NewStyle().Foreground(constants.Tertiary).Padding(0)),
		"MIN":  "",
		"FG":   "",
		"3PT":  "",
		"FT":   "",
		"REB":  "",
		"AST":  "",
		"STL":  "",
		"BLK":  "",
		"TO":   "",
		"+/-":  "",
		"PTS":  "",
	}
}

func renderTeamRow(team string) table.RowData {
	return table.RowData{
		"POS":  "",
		"NAME": table.NewStyledCell(team, lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))),
		"MIN":  "",
		"FG":   "",
		"3PT":  "",
		"FT":   "",
		"REB":  "",
		"AST":  "",
		"STL":  "",
		"BLK":  "",
		"TO":   "",
		"+/-":  "",
		"PTS":  "",
	}
}

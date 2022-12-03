package ui

import (
	"nba-cli/nba"
	"nba-cli/ui/constants"
	"nba-cli/ui/gameboard/scoretext"
	"strconv"

	"github.com/evertras/bubble-table/table"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#874BFD"))

var (
	customBorder = table.Border{
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
)

type keyMap struct {
	Down     key.Binding
	Up       key.Binding
	Previous key.Binding
}

// FullHelp implements help.KeyMap
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

// ShortHelp implements help.KeyMap
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Down, k.Up, k.Previous}
}

type GameModel struct {
	table                 table.Model
	activeGameID          string
	activeGame            nba.BoxScoreSummary
	previousModel         Model
	help                  help.Model
	width, height, margin int
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
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.GetFocused()),
			)
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

	keyMap := keyMap{
		Down:     key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "highlight next row")),
		Up:       key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "highlight previous row")),
		Previous: key.NewBinding(key.WithKeys("esc", "q"), key.WithHelp("q/esc", "back to games list")),
	}
	helpContainer := lipgloss.NewStyle().
		SetString(m.help.View(keyMap)).
		Width(m.width).
		Align(lipgloss.Center).
		PaddingTop(1).
		String()

	// helpText :=
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

	rows := newStatsBoard(constants.Gm, activeGameID)

	t := table.New(columns).WithRows(rows).
		Focused(true).
		Border(customBorder).WithBaseStyle(baseStyle).WithPageSize(constants.WindowSize.Height / 3)

	// TODO: Add more styles
	// // - Game Score
	// // - Team Name
	// - Team Color (optional)
	// - Logo (optional)
	// // - Separate Benchers from Starters
	// // - Add a header for each section
	// // - Separate teams by tables (paginate)
	// // - Help text
	// - Handle non active games

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

	rows = append(rows, table.NewRow(
		table.RowData{
			"POS":  "",
			"NAME": table.NewStyledCell("AWAY TEAM", lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "214", Dark: "0"})),
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
		},
	).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Background(lipgloss.AdaptiveColor{Light: "0", Dark: "214"})))

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
				table.RowData{
					"POS":  "",
					"NAME": table.NewStyledCell("B E N C H", lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "214"}).Padding(0)),
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
				},
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
			rows = append(rows, table.NewRow(
				table.RowData{
					"POS":  "",
					"NAME": table.NewStyledCell("HOME TEAM", lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "214", Dark: "0"})),
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
				},
			).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Background(lipgloss.AdaptiveColor{Light: "0", Dark: "214"})))
		}
	}
	return rows
}

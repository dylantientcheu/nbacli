package ui

import (
	"github.com/dylantientcheu/nbacli/nba"
	"github.com/dylantientcheu/nbacli/ui/constants"

	"github.com/evertras/bubble-table/table"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StandingsModel struct {
	easternConfTable      table.Model
	westernConfTable      table.Model
	help                  help.Model
	width, height, margin int
}

func (m *StandingsModel) recalculateTable() {
	m.easternConfTable = m.easternConfTable.WithTargetWidth(m.width)
	m.westernConfTable = m.westernConfTable.WithTargetWidth(m.width)
}

func (m StandingsModel) Init() tea.Cmd { return nil }

func (m StandingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			// TODO: to team view
			return m, tea.Batch()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.recalculateTable()
	}

	m.easternConfTable, _ = m.easternConfTable.Update(msg)
	m.westernConfTable, cmd = m.westernConfTable.Update(msg)

	return m, cmd
}

func (m StandingsModel) View() string {
	easternTable := m.easternConfTable.View() + "\n"
	westernTable := m.westernConfTable.View() + "\n"

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

	easternConfText := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Padding(1, 3).Background(lipgloss.AdaptiveColor{Light: "214", Dark: "#181818"}).Render("EASTERN CONFERENCE")
	westernConfText := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Padding(1, 3).Background(lipgloss.AdaptiveColor{Light: "214", Dark: "#181818"}).Render("WESTERN CONFERENCE")

	return easternConfText + "\n" + easternTable + "\n" + westernConfText + "\n" + westernTable + "\n" + helpContainer
}

func InitStandingsView() *StandingsModel {
	columns := []table.Column{
		table.NewFlexColumn("POS", "POS", 2),
		table.NewFlexColumn("TEAM", "TEAM", 10),
		table.NewFlexColumn("PCT", "PCT", 5),
		table.NewFlexColumn("HOME", "HOME", 5),
		table.NewFlexColumn("AWAY", "AWAY", 5),
		table.NewFlexColumn("CONF", "CONF", 3),
		table.NewFlexColumn("PPG", "PPG", 3),
		table.NewFlexColumn("OPPPPG", "OPPPPG", 3),
		table.NewFlexColumn("DIFF", "DIFF", 3),
		table.NewFlexColumn("STRK", "STRK", 3),
		table.NewFlexColumn("L10", "L10", 3),
	}

	easternRows, westernRows := newStandingsBoard(constants.St)

	tEast := table.New(columns).WithRows(easternRows).Focused(true).Border(constants.CustomTableBorder).WithBaseStyle(constants.BaseStyle).WithPageSize(constants.WindowSize.Height / 3)
	tWest := table.New(columns).WithRows(westernRows).Border(constants.CustomTableBorder).WithBaseStyle(constants.BaseStyle).WithPageSize(constants.WindowSize.Height / 3)

	m := StandingsModel{tEast, tWest, help.New(), constants.WindowSize.Height, constants.WindowSize.Width, 3}
	return &m
}

func newStandingsBoard(standings *nba.StandingsRepository) ([]table.Row, []table.Row) {
	easternConference, westernConference := standings.GetSeasonStandings()
	return standingsToRows(easternConference, westernConference)
}

func standingsToRows(easternConferenceStandings []nba.Standing, westernConferenceStandings []nba.Standing) ([]table.Row, []table.Row) {
	var (
		eastRows []table.Row
		westRows []table.Row
	)

	for _, stat := range easternConferenceStandings {
		eastRows = append(eastRows, table.NewRow(
			table.RowData{
				"POS":    stat.PlayoffRank,
				"TEAM":   stat.TeamName,
				"PCT":    stat.WinPCT,
				"HOME":   stat.Home,
				"AWAY":   stat.Road,
				"CONF":   stat.ConferenceRecord,
				"PPG":    stat.PointsPG,
				"OPPPPG": stat.OppPointsPG,
				"DIFF":   stat.DiffPointsPG,
				"STRK":   stat.StrCurrentStreak,
				"L10":    stat.L10,
			},
		))
	}

	for _, stat := range westernConferenceStandings {
		westRows = append(westRows, table.NewRow(
			table.RowData{
				"POS":    stat.PlayoffRank,
				"TEAM":   stat.TeamName,
				"PCT":    stat.WinPCT,
				"HOME":   stat.Home,
				"AWAY":   stat.Road,
				"CONF":   stat.ConferenceRecord,
				"PPG":    stat.PointsPG,
				"OPPPPG": stat.OppPointsPG,
				"DIFF":   stat.DiffPointsPG,
				"STRK":   stat.StrCurrentStreak,
				"L10":    stat.L10,
			},
		))
	}
	return eastRows, westRows
}

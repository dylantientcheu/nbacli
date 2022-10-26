package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blurdylan/go-nba/entity"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Game struct {
	entity.Game
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (i Game) Title() string {
	vTeam, _ := GetTeamByIdOrTricode(i.VTeam.TeamID, i.VTeam.TriCode)
	hTeam, _ := GetTeamByIdOrTricode(i.HTeam.TeamID, i.HTeam.TriCode)
	return hTeam.TeamShortName + " vs " + vTeam.TeamShortName
}
func (i Game) Description() string {
	timeUntil := time.Until(i.StartTimeUTC).Round(time.Minute)
	venue := fmt.Sprintf("%s - %s, %s", i.Arena.Name, i.Arena.City, i.Arena.StateAbbr)
	return fmt.Sprintf("Tip-off in %s | %s", timeUntil.String(), venue)
}
func (i Game) FilterValue() string { return i.GameID }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func FetchUpcomingGames() {
	// an upcoming game is a game which starts in less than 23hours,
	// we need to fetch all today and yesterday games which may fit in this criteria

	currentDate := GetCurrentDate()
	GAME_URL := fmt.Sprintf("https://data.nba.net/prod/v1/%s/scoreboard.json", currentDate)
	resp, err := http.Get(GAME_URL)

	if err != nil {
		log.Fatal("an error occurred, please try again")
	}

	defer resp.Body.Close()

	var scoreboard entity.Scoreboard
	if err := json.NewDecoder(resp.Body).Decode(&scoreboard); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	games := make([]Game, 0, len(scoreboard.Games))

	for _, game := range scoreboard.Games {
		games = append(games, Game{game})
	}

	items := []list.Item{}

	// put all games into items
	for _, game := range games {
		items = append(items, game)
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = fmt.Sprintf("There are %d upcoming games today\n", scoreboard.NumGames)

	if (scoreboard.NumGames) == 0 {
		fmt.Printf("There are no upcoming games today")
		// todo: incase we have no games print something sweeeet!!
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Team struct {
	IsNBAFranchise bool   `json:"isNBAFranchise"`
	IsAllStar      bool   `json:"isAllStar"`
	City           string `json:"city"`
	AltCityName    string `json:"altCityName"`
	FullName       string `json:"fullName"`
	Tricode        string `json:"tricode"`
	TeamID         string `json:"teamId"`
	Nickname       string `json:"nickname"`
	URLName        string `json:"urlName"`
	TeamShortName  string `json:"teamShortName"`
	ConfName       string `json:"confName"`
	DivName        string `json:"divName"`
}

// GetDateInFormat returns the current date in the format YYYYMMDD
func GetCurrentDate() string {

	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day() // - 1 // todo: get timezone before showing today's games.

	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

func GetTeamByIdOrTricode(id string, tricode string) (Team, error) {
	jsonFile, err := os.Open("./static/teams.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Teams array
	var teams []Team
	json.Unmarshal(byteValue, &teams)

	// find the team with the id or tricode
	for i := 0; i < len(teams); i++ {
		if teams[i].TeamID == id || teams[i].Tricode == tricode {
			return teams[i], nil
		}
	}

	// return an empty team if not found
	return Team{}, fmt.Errorf("Team not found")
}

// format a date passed as DD/MM/YYYY to YYYYMMDD
func FormatDate(date string) string {
	return fmt.Sprintf("%s%s%s", date[6:], date[3:5], date[0:2])
}

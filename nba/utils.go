package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Team struct {
	IsNBAFranchise bool   `json:"isNBAFranchise"`
	IsAllStar      bool   `json:"isAllStar"`
	City           string `json:"city"`
	AltCityName    string `json:"altCityName"`
	FullName       string `json:"fullName"`
	Tricode        string `json:"tricode"`
	TeamID         int64  `json:"teamId"`
	Nickname       string `json:"nickname"`
	URLName        string `json:"urlName"`
	TeamShortName  string `json:"teamShortName"`
	ConfName       string `json:"confName"`
	DivName        string `json:"divName"`
}

func GetTeamByIdOrTricode(id int64, tricode string) (Team, error) {
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

func GetDateTimeFromESTInUTC(estTime string, gameDate string) time.Time {
	cleanDate := gameDate[:len(gameDate)-9]

	cleanTime := strings.Replace(estTime, " ", "", -1)
	cleanTime = strings.TrimSpace(strings.ToUpper(cleanTime[:len(cleanTime)-2]))

	timeMeridian := cleanTime[len(cleanTime)-2:]

	fullTime := cleanTime[:len(cleanTime)-2] + ":00"

	t, _ := time.Parse("03:04:05PM", fullTime+timeMeridian)
	fullTime = t.Format("15:04:05")

	EST, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// parse the date and time in the lyt format
	const lyt = "2006-01-02 15:04:05 MST"

	timeAndLoc := fmt.Sprintf("%s %s", cleanDate, fullTime+" EST")

	dt, err := time.ParseInLocation(lyt, timeAndLoc, EST)
	if err != nil {
		panic(err)
	}

	return dt.UTC()

}

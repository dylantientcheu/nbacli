package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dylantientcheu/nbacli/nba"
	"github.com/dylantientcheu/nbacli/ui/constants"

	tea "github.com/charmbracelet/bubbletea"
)

// StartTea the entry point for the UI. Initializes the model.
func StartTea(date time.Time) {
	scbrd := nba.ScoreboardRepository{}
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	constants.Sb = &scbrd

	m := InitScoreboard(date)
	UpdateTeaView(m)
}

func StartStanding() {
	m := InitStandingsView()
	UpdateTeaView(m)
}

func UpdateTeaView(m tea.Model) {
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

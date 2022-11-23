package ui

import (
	"fmt"
	"log"
	"nba-cli/nba"
	"nba-cli/ui/constants"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// StartTea the entry point for the UI. Initializes the model.
func StartTea(sb nba.BoxScoreSummaryRepository, date time.Time) {
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
	constants.Sb = &sb
	// constants.Gm = &gm

	m := InitScoreboard(date)
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if err := constants.P.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

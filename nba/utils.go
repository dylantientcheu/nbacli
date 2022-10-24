package nba

import (
	"fmt"
	"time"
)

// GetDateInFormat returns the current date in the format YYYYMMDD
func GetCurrentDate() string {

	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day() // - 1 // todo: get timezone before showing today's games.

	return fmt.Sprintf("%d%02d%02d", year, month, day)
}

// format a date passed as DD/MM/YYYY to YYYYMMDD
func FormatDate(date string) string {
	return fmt.Sprintf("%s%s%s", date[6:], date[3:5], date[0:2])
}

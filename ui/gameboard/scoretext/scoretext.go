package scoretext

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	subtle         = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#212121"}
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 1).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	scoreTextStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginTop(1)
)

// gotten from https://fsymbols.com/generators/tarty/

var scoreTextFont = map[int]string{
	420: `░░░░░░
░░░░░░
█████╗
╚════╝
░░░░░░
░░░░░░`,
	0: `░█████╗░
██╔══██╗
██║░░██║
██║░░██║
╚█████╔╝
░╚════╝░`,
	1: `░░███╗░░
░████║░░
██╔██║░░
╚═╝██║░░
███████╗
╚══════╝`,
	2: `██████╗░
╚════██╗
░░███╔═╝
██╔══╝░░
███████╗
╚══════╝`,
	3: `██████╗░
╚════██╗
░█████╔╝
░╚═══██╗
██████╔╝
╚═════╝░`,
	4: `░░██╗██╗
░██╔╝██║
██╔╝░██║
███████║
╚════██║
░░░░░╚═╝`,
	5: `███████╗
██╔════╝
██████╗░
╚════██╗
██████╔╝
╚═════╝░`,
	6: `░█████╗░
██╔═══╝░
██████╗░
██╔══██╗
╚█████╔╝
░╚════╝░`,
	7: `███████╗
╚════██║
░░░░██╔╝
░░░██╔╝░
░░██╔╝░░
░░╚═╝░░░`,
	8: `░█████╗░
██╔══██╗
╚█████╔╝
██╔══██╗
╚█████╔╝
░╚════╝░`,
	9: `░█████╗░
██╔══██╗
╚██████║
░╚═══██║
░█████╔╝
░╚════╝░`,
}

func RenderScoreText(score string) string {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	{
		// game board
		scoreTeamHome := lipgloss.JoinHorizontal(lipgloss.Top, getBigScoreText(98), getBigScoreText(420))
		scoreAwayTeam := getBigScoreText(96)

		scoreText := lipgloss.JoinHorizontal(lipgloss.Center, scoreTeamHome, scoreAwayTeam)

		stadium := lipgloss.NewStyle().Width(90).Align(lipgloss.Center).Render("Toyota Center | Houston, TX\n22 Oct 2021")

		score := lipgloss.NewStyle().Width(80).MarginTop(2).Align(lipgloss.Center).Render("Houston Rockets          Golden State Warriors")
		gameInfo := lipgloss.JoinVertical(lipgloss.Center, stadium, score)
		ui := lipgloss.JoinVertical(lipgloss.Center, gameInfo, scoreText)

		gameBoard := lipgloss.Place(width, 18,
			lipgloss.Center, lipgloss.Top,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("篮球"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(gameBoard + "\n\n")
	}

	return doc.String()
}

func getBigScoreText(number int) string {
	if number == 420 {
		return scoreTextStyle.Render(scoreTextFont[420])
	}

	scoreSlice := splitInt(number)
	result := ""

	for _, v := range scoreSlice {
		result = lipgloss.JoinHorizontal(lipgloss.Top, result, scoreTextStyle.Render(scoreTextFont[v]))
	}

	return result
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}

	return result
}

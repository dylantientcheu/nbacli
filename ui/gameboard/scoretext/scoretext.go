package scoretext

import (
	"fmt"
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

	teamNameStyle = lipgloss.NewStyle().Margin(0, 1)
)

// gotten from https://fsymbols.com/generators/tarty/

var scoreTextFont = map[int]string{
	420: `      
      
█████╗
╚════╝
      
      `,
	0: ` █████╗ 
██╔══██╗
██║  ██║
██║  ██║
╚█████╔╝
 ╚════╝ `,
	1: `  ███╗  
 ████║  
██╔██║  
╚═╝██║  
███████╗
╚══════╝`,
	2: `██████╗ 
╚════██╗
  ███╔═╝
██╔══╝  
███████╗
╚══════╝`,
	3: `██████╗ 
╚════██╗
 █████╔╝
 ╚═══██╗
██████╔╝
╚═════╝ `,
	4: `  ██╗██╗
 ██╔╝██║
██╔╝ ██║
███████║
╚════██║
     ╚═╝`,
	5: `███████╗
██╔════╝
██████╗ 
╚════██╗
██████╔╝
╚═════╝ `,
	6: ` █████╗ 
██╔═══╝ 
██████╗ 
██╔══██╗
╚█████╔╝
 ╚════╝ `,
	7: `███████╗
╚════██║
    ██╔╝
   ██╔╝ 
  ██╔╝  
  ╚═╝   `,
	8: ` █████╗ 
██╔══██╗
╚█████╔╝
██╔══██╗
╚█████╔╝
 ╚════╝ `,
	9: ` █████╗ 
██╔══██╗
╚██████║
 ╚═══██║
 █████╔╝
 ╚════╝ `,
}

func RenderScoreText(ArenaName string, gameDate string, scoreHome int, scoreAway int, HomeTeamName string, VisitorTeamName string) string {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	{
		// game board
		scoreTeamHome := lipgloss.JoinHorizontal(lipgloss.Center, teamNameStyle.Render(HomeTeamName), lipgloss.JoinHorizontal(lipgloss.Top, getBigScoreText(scoreHome), getBigScoreText(420)))
		scoreAwayTeam := lipgloss.JoinHorizontal(lipgloss.Center, getBigScoreText(scoreAway), teamNameStyle.Render(VisitorTeamName))

		scoreText := lipgloss.JoinHorizontal(lipgloss.Center, scoreTeamHome, scoreAwayTeam)

		stadium := lipgloss.NewStyle().Width(90).Align(lipgloss.Center).Render(fmt.Sprintf("%s\n%s", ArenaName, gameDate[:len(gameDate)-9]))

		ui := lipgloss.JoinVertical(lipgloss.Center, stadium, scoreText)

		gameBoard := lipgloss.Place(width, 17,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("░"),
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

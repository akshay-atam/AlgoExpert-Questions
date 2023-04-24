package main

import "fmt"

const HOME_TEAM_WON = 1

// TournamentWinner function
func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}

	for i, competition := range competitions {
		result := results[i]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}
	return currentBestTeam
}

// updateScores function
func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}

func main() {
	competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	results := []int{0, 0, 1}
	fmt.Println("Winner: ", TournamentWinner(competitions, results))
}

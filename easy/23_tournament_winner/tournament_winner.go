package tournament_winner

const homeTeamWon = 1

// TournamentWinner
// Better solution
func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}

	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == homeTeamWon {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}

	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}

// My solution
func tournamentWinner(competitions [][]string, results []int) string {

	algoPoint := make(map[string]int)
	for idx, competition := range competitions {
		homeTeam := competition[0]
		awayTeam := competition[1]

		result := results[idx]

		if result == homeTeamWon {
			algoPoint[homeTeam] += 3
		} else {
			algoPoint[awayTeam] += 3
		}
	}

	var winner string
	var point int
	for team, p := range algoPoint {
		if p > point {
			winner = team
		}
		point += p
	}

	return winner
}

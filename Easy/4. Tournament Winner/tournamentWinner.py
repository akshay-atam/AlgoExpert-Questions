HOME_TEAM_WON = 1

# O(n) time | O(k) space
def tournamentWinner(competitions, results):
    currentBestTeam = ""
    scores = {currentBestTeam: 0}

    for i, competition in enumerate(competitions):
        result = results[i]
        homeTeam, awayTeam = competition
        
        winningTeam = homeTeam if result == HOME_TEAM_WON else awayTeam
        
        updateScores(winningTeam, 3, scores)
        
        if scores[winningTeam] > scores[currentBestTeam]:
            currentBestTeam = winningTeam
            
    return currentBestTeam

def updateScores(team, points, scores):
    if team not in scores:
        scores[team] = 0
        
    scores[team] += points
	
if __name__ == "__main__":
    competitions = [
    ["HTML", "C#"],
    ["C#", "Python"],
    ["Python", "HTML"]
    ]
    results = [0, 0, 1]
	
    print(f"Winner: {tournamentWinner(competitions, results)}")

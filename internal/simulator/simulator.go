package simulator

import (
	"flip7-simulator/internal/game"
	"fmt"
)

// SimulationResult holds the results for an algorithm
type SimulationResult struct {
	AlgorithmName string
	GamesWon      int
	TotalScore    int
	AverageScore  float64
	Flip7Count    int
	BustCount     int
}

// Simulator runs multiple games with different algorithms
type Simulator struct {
	algorithms []game.Algorithm
	numGames   int
}

// NewSimulator creates a new simulator
func NewSimulator(algorithms []game.Algorithm, numGames int) *Simulator {
	return &Simulator{
		algorithms: algorithms,
		numGames:   numGames,
	}
}

// Run executes the simulation
func (s *Simulator) Run() {
	results := make([]SimulationResult, len(s.algorithms))

	// Initialize results
	for i, algo := range s.algorithms {
		results[i] = SimulationResult{
			AlgorithmName: algo.GetName(),
		}
	}

	// Run games
	for gameNum := 0; gameNum < s.numGames; gameNum++ {
		winner, scores, flip7s, busts := s.playGame()

		// Update results
		for i := range results {
			results[i].TotalScore += scores[i]
			results[i].Flip7Count += flip7s[i]
			results[i].BustCount += busts[i]

			if winner == i {
				results[i].GamesWon++
			}
		}

		// Progress indicator
		if (gameNum+1)%100 == 0 {
			fmt.Printf("Completed %d/%d games\n", gameNum+1, s.numGames)
		}
	}

	// Calculate averages
	for i := range results {
		if s.numGames > 0 {
			results[i].AverageScore = float64(results[i].TotalScore) / float64(s.numGames)
		}
	}

	s.displayResults(results)
}

// playGame runs a single game and returns winner index, scores, flip7 counts, and bust counts
func (s *Simulator) playGame() (int, []int, []int, []int) {
	g := game.NewGame(len(s.algorithms))
	scores := make([]int, len(s.algorithms))
	flip7s := make([]int, len(s.algorithms))
	busts := make([]int, len(s.algorithms))

	// Play until someone reaches 200 points
	for {
		g.CreateDeck()
		g.StartNewRound()

		// Deal initial cards
		g.DealInitialCard()

		// Play the round
		for !g.IsRoundOver() {
			for playerID := 0; playerID < len(s.algorithms); playerID++ {
				player := g.Players[playerID]

				if player.IsBust || player.HasStood {
					continue
				}

				// Get algorithm decision
				gameState := g.GetGameState()
				cardsRemaining := g.GetCardsRemaining()
				decision := s.algorithms[playerID].MakeDecision(player, gameState, cardsRemaining)

				if decision.Action == "hit" {
					g.PlayerHit(playerID)
				} else {
					g.PlayerStand(playerID)
				}

				// Check if round is over due to Flip 7
				if g.IsRoundOver() {
					break
				}
			}
		}

		// Calculate round scores
		roundWinner := -1
		highestScore := -1

		for playerID := 0; playerID < len(s.algorithms); playerID++ {
			roundScore := g.CalculateScore(playerID)
			g.Players[playerID].GameScore += roundScore
			scores[playerID] = g.Players[playerID].GameScore

			if g.HasFlip7(playerID) {
				flip7s[playerID]++
			}

			if g.Players[playerID].IsBust {
				busts[playerID]++
			}

			// Track highest score for potential game end
			if scores[playerID] > highestScore {
				highestScore = scores[playerID]
				roundWinner = playerID
			}
		}

		// Check if game is over (someone reached 200)
		if highestScore >= 200 {
			return roundWinner, scores, flip7s, busts
		}
	}
}

// displayResults shows the simulation results
func (s *Simulator) displayResults(results []SimulationResult) {
	fmt.Printf("\n=== Flip 7 Simulation Results ===\n")
	fmt.Printf("Total Games: %d\n\n", s.numGames)

	// Sort by win rate
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].GamesWon > results[i].GamesWon {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	fmt.Printf("%-20s %8s %8s %12s %8s %8s\n", "Algorithm", "Wins", "Win%", "Avg Score", "Flip 7s", "Busts")
	fmt.Printf("%-20s %8s %8s %12s %8s %8s\n", "=========", "====", "====", "=========", "=======", "=====")

	for _, result := range results {
		winRate := float64(result.GamesWon) / float64(s.numGames) * 100
		fmt.Printf("%-20s %8d %7.1f%% %11.1f %8d %8d\n",
			result.AlgorithmName,
			result.GamesWon,
			winRate,
			result.AverageScore,
			result.Flip7Count,
			result.BustCount)
	}

	fmt.Printf("\n")
}

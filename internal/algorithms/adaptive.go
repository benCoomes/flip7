package algorithms

import (
	"flip7-simulator/internal/game"
)

// AdaptiveAlgorithm adapts strategy based on opponents' scores
type AdaptiveAlgorithm struct {
	name string
}

func NewAdaptiveAlgorithm() *AdaptiveAlgorithm {
	return &AdaptiveAlgorithm{
		name: "Adaptive",
	}
}

func (a *AdaptiveAlgorithm) MakeDecision(playerState game.PlayerState, gameState game.GameState, cardsRemaining map[int]int) game.Decision {
	// Calculate our current position
	ourScore := playerState.GameScore

	// Find highest opponent score
	maxOpponentScore := 0
	for _, player := range gameState.Players {
		if player.ID != playerState.ID && player.GameScore > maxOpponentScore {
			maxOpponentScore = player.GameScore
		}
	}

	// Count unique values and current round score
	uniqueValues := make(map[int]bool)
	currentRoundScore := 0
	hasX2 := false

	for _, card := range playerState.Cards {
		if card.CardType == "number" {
			uniqueValues[card.Value] = true
			currentRoundScore += card.Value
		} else if card.CardType == "modifier" {
			if card.IsX2 {
				hasX2 = true
			} else {
				currentRoundScore += card.Modifier
			}
		}
	}

	if hasX2 {
		currentRoundScore *= 2
	}

	// If we're behind, be more aggressive
	scoreDifference := maxOpponentScore - ourScore

	if scoreDifference > 50 {
		// Far behind - go for Flip 7 or high scores
		if len(uniqueValues) >= 3 {
			return game.Decision{Action: "hit"}
		}
		if currentRoundScore < 40 {
			return game.Decision{Action: "hit"}
		}
	} else if scoreDifference > 20 {
		// Slightly behind - moderate risk
		if len(uniqueValues) >= 4 {
			return game.Decision{Action: "hit"}
		}
		if currentRoundScore < 35 {
			return game.Decision{Action: "hit"}
		}
	} else {
		// Ahead or close - be conservative
		if len(uniqueValues) >= 6 {
			return game.Decision{Action: "hit"}
		}
		if currentRoundScore >= 25 {
			return game.Decision{Action: "stand"}
		}
	}

	// Calculate bust risk
	safeCards := 0
	totalCards := 0

	for value, count := range cardsRemaining {
		totalCards += count
		if !uniqueValues[value] {
			safeCards += count
		}
	}

	if totalCards == 0 {
		return game.Decision{Action: "stand"}
	}

	bustRisk := float64(totalCards-safeCards) / float64(totalCards)

	if bustRisk > 0.6 {
		return game.Decision{Action: "stand"}
	}

	return game.Decision{Action: "hit"}
}

func (a *AdaptiveAlgorithm) GetName() string {
	return a.name
}

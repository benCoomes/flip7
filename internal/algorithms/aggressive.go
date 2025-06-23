package algorithms

import (
	"flip7-simulator/internal/game"
)

// AggressiveAlgorithm goes for Flip 7 more aggressively
type AggressiveAlgorithm struct {
	name string
}

func NewAggressiveAlgorithm() *AggressiveAlgorithm {
	return &AggressiveAlgorithm{
		name: "Aggressive",
	}
}

func (a *AggressiveAlgorithm) MakeDecision(playerState game.PlayerState, gameState game.GameState, cardsRemaining map[int]int) game.Decision {
	// Count unique number values we have
	uniqueValues := make(map[int]bool)
	currentScore := 0
	hasX2 := false

	for _, card := range playerState.Cards {
		if card.CardType == "number" {
			uniqueValues[card.Value] = true
			currentScore += card.Value
		} else if card.CardType == "modifier" {
			if card.IsX2 {
				hasX2 = true
			} else {
				currentScore += card.Modifier
			}
		}
	}

	if hasX2 {
		currentScore *= 2
	}

	// Always go for Flip 7 if we have 4+ unique cards
	if len(uniqueValues) >= 4 {
		return game.Decision{Action: "hit"}
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

	// More aggressive thresholds
	if currentScore >= 45 && bustRisk > 0.4 {
		return game.Decision{Action: "stand"}
	}

	if currentScore >= 60 && bustRisk > 0.2 {
		return game.Decision{Action: "stand"}
	}

	return game.Decision{Action: "hit"}
}

func (a *AggressiveAlgorithm) GetName() string {
	return a.name
}

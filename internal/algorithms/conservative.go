package algorithms

import (
	"flip7-simulator/internal/game"
)

// ConservativeAlgorithm uses risk assessment based on cards seen
type ConservativeAlgorithm struct {
	name string
}

func NewConservativeAlgorithm() *ConservativeAlgorithm {
	return &ConservativeAlgorithm{
		name: "Conservative",
	}
}

func (a *ConservativeAlgorithm) MakeDecision(playerState game.PlayerState, gameState game.GameState, cardsRemaining map[int]int) game.Decision {
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

	// If we have 6 unique values, go for Flip 7
	if len(uniqueValues) == 6 {
		return game.Decision{Action: "hit"}
	}

	// Calculate risk of busting
	safeCards := 0
	totalCards := 0

	for value, count := range cardsRemaining {
		totalCards += count
		if !uniqueValues[value] {
			safeCards += count
		}
	}

	bustRisk := float64(totalCards-safeCards) / float64(totalCards)

	// Conservative thresholds
	if currentScore >= 35 && bustRisk > 0.3 {
		return game.Decision{Action: "stand"}
	}

	if currentScore >= 25 && bustRisk > 0.5 {
		return game.Decision{Action: "stand"}
	}

	return game.Decision{Action: "hit"}
}

func (a *ConservativeAlgorithm) GetName() string {
	return a.name
}

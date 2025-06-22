package algorithms

import (
	"flip7-simulator/internal/game"
	"fmt"
)

// AlwaysHitAlgorithm always chooses to hit until bust or Flip 7
type AlwaysHitAlgorithm struct {
	name string
}

func NewAlwaysHitAlgorithm() *AlwaysHitAlgorithm {
	return &AlwaysHitAlgorithm{
		name: "Always Hit",
	}
}

func (a *AlwaysHitAlgorithm) MakeDecision(playerState game.PlayerState, gameState game.GameState, cardsRemaining map[int]int) game.Decision {
	return game.Decision{Action: "hit"}
}

func (a *AlwaysHitAlgorithm) GetName() string {
	return a.name
}

// StopAtScoreAlgorithm stops when reaching a target score
type StopAtScoreAlgorithm struct {
	name        string
	targetScore int
}

func NewStopAtScoreAlgorithm(targetScore int) *StopAtScoreAlgorithm {
	return &StopAtScoreAlgorithm{
		name:        fmt.Sprintf("Stop at %d", targetScore),
		targetScore: targetScore,
	}
}

func (a *StopAtScoreAlgorithm) MakeDecision(playerState game.PlayerState, gameState game.GameState, cardsRemaining map[int]int) game.Decision {
	// Calculate current score
	currentScore := 0
	hasX2 := false

	for _, card := range playerState.Cards {
		if card.CardType == "number" {
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

	if currentScore >= a.targetScore {
		return game.Decision{Action: "stand"}
	}

	return game.Decision{Action: "hit"}
}

func (a *StopAtScoreAlgorithm) GetName() string {
	return a.name
}

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

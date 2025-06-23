package algorithms

import (
	"flip7-simulator/internal/game"
	"fmt"
)

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

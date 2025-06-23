package algorithms

import (
	"flip7-simulator/internal/game"
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

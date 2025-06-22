package game

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	game := NewGame(2)
	game.CreateDeck()

	// Check that deck has the expected number of cards
	// 0(1) + 1(1) + 2(2) + 3(3) + ... + 12(12) = 1+1+2+3+4+5+6+7+8+9+10+11+12 = 79
	// Plus 6 modifier cards (+1, +2, +3 x2 each) + 1 x2 card = 7 modifier cards
	expectedCards := 79 + 7 // number cards + modifiers
	if len(game.Deck) != expectedCards {
		t.Errorf("Expected %d cards in deck, got %d", expectedCards, len(game.Deck))
	}
}

func TestPlayerHit(t *testing.T) {
	game := NewGame(2)
	game.CreateDeck()

	// Give player a card first
	game.DealInitialCard()
	initialCards := len(game.Players[0].Cards)

	// Hit should add a card
	success := game.PlayerHit(0)
	if !success {
		t.Error("PlayerHit should succeed")
	}

	if len(game.Players[0].Cards) != initialCards+1 {
		t.Error("PlayerHit should add one card")
	}
}

func TestBustCondition(t *testing.T) {
	game := NewGame(1)

	// Manually create a scenario where player will bust
	game.Players[0].Cards = []Card{
		{Value: 5, CardType: "number"},
	}

	// Add the same value card to deck
	game.Deck = []Card{
		{Value: 5, CardType: "number"},
	}

	// Hit should cause bust
	game.PlayerHit(0)

	if !game.Players[0].IsBust {
		t.Error("Player should be bust after receiving duplicate card")
	}

	if len(game.Players[0].Cards) != 0 {
		t.Error("Bust player should have no cards")
	}
}

func TestFlip7Detection(t *testing.T) {
	game := NewGame(1)

	// Create a Flip 7 scenario
	game.Players[0].Cards = []Card{
		{Value: 0, CardType: "number"},
		{Value: 1, CardType: "number"},
		{Value: 2, CardType: "number"},
		{Value: 3, CardType: "number"},
		{Value: 4, CardType: "number"},
		{Value: 5, CardType: "number"},
		{Value: 6, CardType: "number"},
		{Value: 0, CardType: "modifier", Modifier: 2}, // modifier doesn't count
	}

	if !game.HasFlip7(0) {
		t.Error("Should detect Flip 7 with 7 unique number values")
	}
}

func TestScoreCalculation(t *testing.T) {
	game := NewGame(1)

	// Test basic scoring
	game.Players[0].Cards = []Card{
		{Value: 5, CardType: "number"},
		{Value: 10, CardType: "number"},
		{Value: 0, CardType: "modifier", Modifier: 3},
	}

	expectedScore := 5 + 10 + 3 // = 18
	score := game.CalculateScore(0)

	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}

func TestScoreWithX2Multiplier(t *testing.T) {
	game := NewGame(1)

	// Test x2 multiplier
	game.Players[0].Cards = []Card{
		{Value: 5, CardType: "number"},
		{Value: 10, CardType: "number"},
		{Value: 0, CardType: "modifier", IsX2: true},
	}

	expectedScore := (5 + 10) * 2 // = 30
	score := game.CalculateScore(0)

	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}

func TestFlip7Bonus(t *testing.T) {
	game := NewGame(1)

	// Test Flip 7 bonus (should not be doubled)
	game.Players[0].Cards = []Card{
		{Value: 0, CardType: "number"},
		{Value: 1, CardType: "number"},
		{Value: 2, CardType: "number"},
		{Value: 3, CardType: "number"},
		{Value: 4, CardType: "number"},
		{Value: 5, CardType: "number"},
		{Value: 6, CardType: "number"},
		{Value: 0, CardType: "modifier", IsX2: true},
	}

	baseScore := (0 + 1 + 2 + 3 + 4 + 5 + 6) * 2 // = 42
	flip7Bonus := 15
	expectedScore := baseScore + flip7Bonus // = 57

	score := game.CalculateScore(0)

	if score != expectedScore {
		t.Errorf("Expected score %d, got %d", expectedScore, score)
	}
}

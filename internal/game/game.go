package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Card represents a card in the Flip 7 deck
type Card struct {
	Value    int    // 0-12 for number cards
	CardType string // "number", "modifier", "action"
	Modifier int    // For +1, +2, +3 cards
	IsX2     bool   // For x2 multiplier card
}

// PlayerState represents the current state of a player
type PlayerState struct {
	ID        int
	Cards     []Card
	Score     int
	IsBust    bool
	HasStood  bool
	GameScore int // Total score across all games
}

// GameState represents the current state of the game
type GameState struct {
	Players      []PlayerState
	Deck         []Card
	DiscardPile  []Card
	CurrentRound int
	IsGameOver   bool
	Winner       int
}

// Decision represents a player's choice
type Decision struct {
	Action string // "hit" or "stand"
}

// Algorithm interface for different playing strategies
type Algorithm interface {
	MakeDecision(playerState PlayerState, gameState GameState, cardsRemaining map[int]int) Decision
	GetName() string
}

// Game manages the Flip 7 game logic
type Game struct {
	Players     []PlayerState
	Deck        []Card
	DiscardPile []Card
	rng         *rand.Rand
}

// NewGame creates a new Flip 7 game
func NewGame(numPlayers int) *Game {
	game := &Game{
		Players: make([]PlayerState, numPlayers),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// Initialize players
	for i := 0; i < numPlayers; i++ {
		game.Players[i] = PlayerState{
			ID:    i,
			Cards: make([]Card, 0),
		}
	}

	return game
}

// CreateDeck creates the Flip 7 deck according to the rules
func (g *Game) CreateDeck() {
	g.Deck = make([]Card, 0)

	// Number cards: 12 twelve-value cards down to 2 two-value cards, 1 one-value, 1 zero-value
	for value := 0; value <= 12; value++ {
		count := 1
		if value >= 2 {
			count = value
		}

		for i := 0; i < count; i++ {
			g.Deck = append(g.Deck, Card{
				Value:    value,
				CardType: "number",
			})
		}
	}

	// Modifier cards (+1, +2, +3, x2)
	// Adding some modifier cards (exact counts not specified in rules)
	modifiers := []int{1, 2, 3}
	for _, mod := range modifiers {
		for i := 0; i < 2; i++ { // 2 of each modifier
			g.Deck = append(g.Deck, Card{
				Value:    0,
				CardType: "modifier",
				Modifier: mod,
			})
		}
	}

	// Add x2 multiplier card
	g.Deck = append(g.Deck, Card{
		Value:    0,
		CardType: "modifier",
		IsX2:     true,
	})

	// Shuffle the deck
	g.ShuffleDeck()
}

// ShuffleDeck shuffles the current deck
func (g *Game) ShuffleDeck() {
	for i := len(g.Deck) - 1; i > 0; i-- {
		j := g.rng.Intn(i + 1)
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	}
}

// DealInitialCard deals one card to each player to start a round
func (g *Game) DealInitialCard() {
	for i := range g.Players {
		if !g.Players[i].HasStood && !g.Players[i].IsBust {
			card := g.DrawCard()
			if card != nil {
				g.Players[i].Cards = append(g.Players[i].Cards, *card)
			}
		}
	}
}

// DrawCard draws a card from the deck, reshuffling if necessary
func (g *Game) DrawCard() *Card {
	if len(g.Deck) == 0 {
		g.ReshuffleDeck()
	}

	if len(g.Deck) == 0 {
		return nil // No cards available
	}

	card := g.Deck[0]
	g.Deck = g.Deck[1:]
	return &card
}

// ReshuffleDeck reshuffles the discard pile back into the deck
func (g *Game) ReshuffleDeck() {
	g.Deck = append(g.Deck, g.DiscardPile...)
	g.DiscardPile = make([]Card, 0)
	g.ShuffleDeck()
}

// PlayerHit gives a player another card
func (g *Game) PlayerHit(playerID int) bool {
	if playerID < 0 || playerID >= len(g.Players) {
		return false
	}

	player := &g.Players[playerID]
	if player.IsBust || player.HasStood {
		return false
	}

	card := g.DrawCard()
	if card == nil {
		return false
	}

	// Check for bust condition (duplicate number value)
	if card.CardType == "number" {
		for _, existingCard := range player.Cards {
			if existingCard.CardType == "number" && existingCard.Value == card.Value {
				player.IsBust = true
				g.DiscardPile = append(g.DiscardPile, player.Cards...)
				player.Cards = make([]Card, 0)
				return true
			}
		}
	}

	player.Cards = append(player.Cards, *card)

	// Check for Flip 7
	if g.HasFlip7(playerID) {
		return true
	}

	return true
}

// PlayerStand makes a player stand with their current cards
func (g *Game) PlayerStand(playerID int) {
	if playerID >= 0 && playerID < len(g.Players) {
		g.Players[playerID].HasStood = true
	}
}

// HasFlip7 checks if a player has achieved Flip 7
func (g *Game) HasFlip7(playerID int) bool {
	if playerID < 0 || playerID >= len(g.Players) {
		return false
	}

	player := g.Players[playerID]
	uniqueValues := make(map[int]bool)

	for _, card := range player.Cards {
		if card.CardType == "number" {
			uniqueValues[card.Value] = true
		}
	}

	return len(uniqueValues) == 7
}

// CalculateScore calculates a player's score for the round
func (g *Game) CalculateScore(playerID int) int {
	if playerID < 0 || playerID >= len(g.Players) {
		return 0
	}

	player := g.Players[playerID]
	if player.IsBust {
		return 0
	}

	score := 0
	hasX2 := false

	// Calculate base score
	for _, card := range player.Cards {
		if card.CardType == "number" {
			score += card.Value
		} else if card.CardType == "modifier" {
			if card.IsX2 {
				hasX2 = true
			} else {
				score += card.Modifier
			}
		}
	}

	// Apply x2 multiplier (but not to Flip 7 bonus)
	if hasX2 {
		score *= 2
	}

	// Add Flip 7 bonus (cannot be doubled)
	if g.HasFlip7(playerID) {
		score += 15
	}

	return score
}

// IsRoundOver checks if the current round is over
func (g *Game) IsRoundOver() bool {
	activePlayers := 0

	for _, player := range g.Players {
		if !player.IsBust && !player.HasStood {
			activePlayers++
		}
		// Check for Flip 7
		if g.HasFlip7(player.ID) {
			return true
		}
	}

	return activePlayers == 0
}

// StartNewRound resets players for a new round
func (g *Game) StartNewRound() {
	for i := range g.Players {
		g.Players[i].Cards = make([]Card, 0)
		g.Players[i].IsBust = false
		g.Players[i].HasStood = false
	}
}

// GetCardsRemaining returns a map of remaining cards in the deck by value
func (g *Game) GetCardsRemaining() map[int]int {
	remaining := make(map[int]int)

	for _, card := range g.Deck {
		if card.CardType == "number" {
			remaining[card.Value]++
		}
	}

	return remaining
}

// GetGameState returns the current game state
func (g *Game) GetGameState() GameState {
	return GameState{
		Players:     g.Players,
		Deck:        g.Deck,
		DiscardPile: g.DiscardPile,
	}
}

// PrintGameState prints the current state for debugging
func (g *Game) PrintGameState() {
	fmt.Printf("=== Game State ===\n")
	fmt.Printf("Deck size: %d\n", len(g.Deck))

	for _, player := range g.Players {
		fmt.Printf("Player %d: ", player.ID)
		if player.IsBust {
			fmt.Printf("BUST")
		} else if player.HasStood {
			fmt.Printf("STOOD (Score: %d)", g.CalculateScore(player.ID))
		} else {
			fmt.Printf("Cards: ")
			for _, card := range player.Cards {
				if card.CardType == "number" {
					fmt.Printf("%d ", card.Value)
				} else if card.CardType == "modifier" {
					if card.IsX2 {
						fmt.Printf("x2 ")
					} else {
						fmt.Printf("+%d ", card.Modifier)
					}
				}
			}
			fmt.Printf("(Score: %d)", g.CalculateScore(player.ID))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("================\n")
}

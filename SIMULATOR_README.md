# Flip 7 Simulator

A Go-based simulator for the card game Flip 7, where different algorithms compete against each other.

## Building

```bash
go build -o flip7-simulator ./cmd/
```

## Usage

Run the simulator with default settings (1000 games):
```bash
./flip7-simulator
```

Run with a specific number of games:
```bash
./flip7-simulator -games 500
```

Show help:
```bash
./flip7-simulator -help
```

## Game Rules

Flip 7 is a card game where players try to collect cards without getting duplicates:

- **Deck**: Contains 12 twelve-value cards, 11 eleven-value cards, down to 2 two-value cards, plus single 1-value and 0-value cards, along with modifier cards (+1, +2, +3, x2)
- **Goal**: Reach 200 points to win the game
- **Rounds**: Each round, players can "hit" (take another card) or "stand" (keep current cards)
- **Busting**: Getting a duplicate number value causes you to bust and score 0 for the round
- **Flip 7**: Having 7 unique number values gives a 15-point bonus and ends the round
- **Scoring**: Sum of all card values plus modifiers, with potential x2 multiplier

## Algorithms

The simulator includes several different playing strategies:

- **Always Hit**: Never stops until bust or Flip 7
- **Stop at X**: Stops when reaching X points in a round (X=25, 30, 40)
- **Conservative**: Uses risk assessment based on remaining cards
- **Aggressive**: Aggressively pursues Flip 7 opportunities
- **Adaptive**: Adjusts strategy based on opponents' scores

## Output

The simulator shows:
- Win count and percentage for each algorithm
- Average score across all games
- Number of Flip 7s achieved
- Number of busts

## Example Output

```
=== Flip 7 Simulation Results ===
Total Games: 1000

Algorithm                Wins     Win%    Avg Score  Flip 7s    Busts
=========                ====     ====    =========  =======    =====
Stop at 30                472    47.2%       178.3        9     4257
Conservative              304    30.4%       158.0      241     5684
Stop at 40                187    18.7%       134.6      154     6491
Always Hit                 37     3.7%        56.9      757     8662
```

# Flip 7 Simulator

Develop different algorithms to compete against each other in the flip 7 game.

This version has no 'draw 3s', 'second chances' or 'freezes'.

## Quick Start

Build and run the simulator:
```bash
make quick
```

Or manually:
```bash
go build -o flip7-simulator ./cmd/
./flip7-simulator -games 100
```

See `SIMULATOR_README.md` for detailed documentation.

## Available Commands

- `make build` - Build the simulator
- `make test` - Run unit tests  
- `make run` - Run 1000 game simulation
- `make quick` - Run 100 game simulation
- `make help` - Show available commands


## Flip 7 Rules

Modified from https://boardgamegeek.com/thread/3386548/never-tell-me-the-oddsflip-7-a-detailed-review to exclude action and modifier cards from the rules.

### Deck
The deck is made up of twelve 12-value cards, eleven 11-value cards...all the way down to two 2-value cards. There is a single 1-value and a single 0-value card. Then there are a range of Action and Modifier Cards.

### Summary
Players must decide if they will risk taking another card 'Hit' or if they will 'Sit', sticking with what they have. The relative scores and what they may have already observed to have come out of the deck may influence a player's decision.

### Rules
The game plays out as such :

* Deal a Card to All Players – The dealer, will begin any round by dealing a single card to each player face-up.
* Stick or Twist? – When the action comes back around to a player with one or more cards they will have to make a simple decision. They can choose to pass\stay, in which case they are out of the round and will score the points that are in front of them.
  * If they want to take the risk and 'Hit' (go again), the dealer will flip over another card for them. In this way a round is played out, with each player deciding to pass (thus scoring) or busting and getting nothing for the round.
* Busting! – If a player receives a second card of the same value as one they already have in play...they bust. It's as simple as that.
* The Magical 'Flip 7' – The game is called Flip 7 because this is a magical moment if someone can achieve it. To 'Flip 7', a player must stay in a round until they have 7 face-up cards with unique values (numbers). Action Cards and the Modifier Cards ('+' cards and the x2 card) do not count towards a player's total of 7.
* This means a player needs to avoid receiving a duplicate value\number card. In all there are 13 different values in the deck, so needing to have 7 unique ones is a lot. Of course having any of the 0-3 values can be quite helpful and having a Second Chance up your sleeve can come in handy as well.
* Managing a 'Flip 7' allows a player to score and they gain 15 additional points for achieving the feat. This bonus cannot be doubled by the x2 Modifier Card either.
* Ending a Round and Scoring – A round can come to an an in one of 3 ways. All players might bust, all players might pass and of course there can be any combination of players busting or passing to score within a single round.
* The third way to end a round is if a player managed to 'Flip 7'. This ends a round immediately and all players still in-play will score.
* All players that did not bust can add up their points

### Scoring
Example – This example is one of the most complex in the game. If you can understand the following, you can score any combination in the game.

A player holds a 5, 9, 11, 12, 0, 2 and 8 to form a 'Flip 7'. Their score is -
12 + 11 + 9 + 5 + 2 + 0 + 8 = 47 and finally +15 for the 'Flip 7' = 62

Another player passed and is holding 11, 10, 4 and a +2. Their score is more indicative of most scores in a round and results in a score of -
11 + 10 + 4 + 2 = 27

### The Goal and Winning the Game
If any player manages to reach 200 or more at the end of a round, the game is at an end. Multiple players can achieve the feat and in this case, the player with the highest score takes the win. No tie-breaker rules are given, but it would be easy to house-rule one or more ideas. I think 'the player that comes from furthest back takes the win' isn't a bad idea, or perhaps you track 'Flip 7' results and a tied player with more of them under their belt gets the win.

It's a good idea to keep the players informed of the relative scores at the end of each round so they can determine the level of risk they might need to get back in the game or to push for a good lead or for the target of 200 points (if the competition are some way behind).

### Exhausting the Deck and Large Player Counts
It is quite likely that the deck will become exhausted in a game of Flip 7. When this happens simply reshuffle all the cards in the discard pile. It is worth noting that players do not score until the end of a round, so all cards in front of players when the deck runs out will not be added to those cards that are reshuffled, should the round continue.

The game's player count is listed as 3+ and conceivably any number of players could play at the one time. The rulebook suggests that when playing with more than 18 players it is best to use two copies of the game and mash all the cards together. For mine, I would probably limit the game to no more than 10 players with a single deck. Plus it would become rather unwieldly with any more anyway.

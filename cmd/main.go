package main

import (
	"flag"
	"flip7-simulator/internal/algorithms"
	"flip7-simulator/internal/game"
	"flip7-simulator/internal/simulator"
	"fmt"
)

func main() {
	// Command line flags
	numGames := flag.Int("games", 1000, "Number of games to simulate")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		fmt.Println("=== Flip 7 Simulator ===")
		fmt.Println("Simulates multiple games of Flip 7 with different algorithms competing.")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		fmt.Println("\nAlgorithms included:")
		fmt.Println("  - Always Hit: Always takes another card until bust or Flip 7")
		fmt.Println("  - Stop at X: Stops when reaching X points in a round")
		fmt.Println("  - Conservative: Uses risk assessment based on cards seen")
		fmt.Println("  - Aggressive: Aggressively goes for Flip 7")
		fmt.Println("  - Adaptive: Adapts strategy based on opponents' scores")
		return
	}

	fmt.Println("=== Flip 7 Simulator ===")
	fmt.Printf("Running %d games...\n\n", *numGames)

	// Create different algorithms
	algoList := []game.Algorithm{
		algorithms.NewAlwaysHitAlgorithm(),
		algorithms.NewStopAtScoreAlgorithm(25),
		algorithms.NewStopAtScoreAlgorithm(30),
		algorithms.NewStopAtScoreAlgorithm(40),
		algorithms.NewConservativeAlgorithm(),
		algorithms.NewAggressiveAlgorithm(),
		algorithms.NewAdaptiveAlgorithm(),
	}

	// Run simulation
	sim := simulator.NewSimulator(algoList, *numGames)
	sim.Run()
}

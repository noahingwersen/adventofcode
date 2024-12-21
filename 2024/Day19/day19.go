package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const filepath = "input.txt"

func parseInput() ([][]byte, [][]byte) {
	file, _ := os.ReadFile(filepath)
	parts := bytes.Split(file, []byte{'\n', '\n'})
	towels := bytes.Split(parts[0], []byte{',', ' '})
	patterns := bytes.Split(parts[1], []byte{'\n'})

	return towels, patterns
}

func patternIsPossible(pattern []byte, towels [][]byte) bool {
	for _, towel := range towels {
		if bytes.HasPrefix(pattern, towel) {
			remainder := bytes.TrimPrefix(pattern, towel)
			if len(remainder) == 0 {
				return true
			}

			possible := patternIsPossible(remainder, towels)
			if possible {
				return true
			}
		}
	}
	return false
}

func possibleCombinations(pattern []byte, towels [][]byte, cache *map[string]int) int {
	key := string(pattern)
	combos, seen := (*cache)[key]
	if seen {
		return combos
	}

	total := 0
	for _, towel := range towels {
		if bytes.HasPrefix(pattern, towel) {
			remainder := bytes.TrimPrefix(pattern, towel)
			if len(remainder) == 0 {
				total++
			}

			total += possibleCombinations(remainder, towels, cache)
		}
	}

	(*cache)[key] = total
	return total
}

func part1() int {
	towels, patterns := parseInput()
	possible := 0

	for _, pattern := range patterns {
		if patternIsPossible(pattern, towels) {
			possible++
		}
	}
	return possible
}

func part2() int {
	towels, patterns := parseInput()
	combos := 0

	cache := make(map[string]int)
	for _, pattern := range patterns {
		total := possibleCombinations(pattern, towels, &cache)
		combos += total
	}
	return combos
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const filepath = "input.txt"

func parseInput() []string {
	input, _ := os.ReadFile(filepath)
	return strings.Split(string(input), " ")
}

func digits(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func splitInt(i int) [2]int {
	half := digits(i) / 2
	divisor := 1
	for i := 0; i < half; i++ {
		divisor *= 10
	}

	return [2]int{i / divisor, i % divisor}
}

func blink(stone int, rounds int, seen map[[2]int]int) int {
	answer, check := seen[[2]int{stone, rounds}]
	if check {
		return answer
	}

	if rounds == 0 {
		return 1
	}

	newStones := 0
	if stone == 0 {
		newStones += blink(1, rounds-1, seen)
	} else if digits(stone)%2 == 0 {
		parts := splitInt(stone)
		newStones += blink(parts[0], rounds-1, seen)
		newStones += blink(parts[1], rounds-1, seen)
	} else {
		newStones += blink(stone*2024, rounds-1, seen)
	}

	seen[[2]int{stone, rounds}] = newStones
	return newStones
}

func part1() int {
	stones := parseInput()
	rounds := 25
	total := 0
	seen := make(map[[2]int]int)
	for _, stone := range stones {
		stoneValue, _ := strconv.Atoi(stone)
		total += blink(stoneValue, rounds, seen)
	}

	return total
}

func part2() int {
	stones := parseInput()
	rounds := 75
	total := 0
	seen := make(map[[2]int]int)
	for _, stone := range stones {
		stoneValue, _ := strconv.Atoi(stone)
		total += blink(stoneValue, rounds, seen)
	}

	return total
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

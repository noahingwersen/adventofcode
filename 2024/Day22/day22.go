package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const filepath = "input.txt"

func parseInput() []int {
	file, _ := os.ReadFile(filepath)
	var numbers []int
	for _, line := range strings.Split(string(file), "\n") {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}
	return numbers
}

func findSecretNumber(number int) int {
	number = ((number * 64) ^ number) % 16777216
	number = ((number / 32) ^ number) % 16777216
	number = ((number * 2048) ^ number) % 16777216

	return number
}

func sequencePrice(number int, iterations int) map[[4]int]int {
	prices := make(map[[4]int]int)
	var previousNumber int
	var changes []int
	for range iterations {
		previousNumber = number
		number = findSecretNumber(number)
		changes = append(changes, (number%10)-(previousNumber%10))
		length := len(changes)
		if length < 4 {
			continue
		}

		if length > 4 {
			changes = changes[length-4 : length]
		}
		key := [4]int{changes[0], changes[1], changes[2], changes[3]}
		_, seen := prices[key]
		if !seen {
			prices[key] = number % 10
		}
	}

	return prices
}

func part1() int {
	secrets := parseInput()
	total := 0
	for _, secret := range secrets {
		for range 2000 {
			secret = findSecretNumber(secret)
		}
		total += secret
	}

	return total
}

func part2() int {
	secrets := parseInput()

	totalPriceMap := make(map[[4]int]int)
	maxBananas := 0
	for _, secret := range secrets {
		priceMap := sequencePrice(secret, 2000)
		for key, value := range priceMap {
			totalPriceMap[key] += value
			if totalPriceMap[key] > maxBananas {
				maxBananas = totalPriceMap[key]
			}
		}
	}
	return maxBananas
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

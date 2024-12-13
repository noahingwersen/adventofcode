package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDistance(current string, end string, traveled int, compare string, distances map[string]map[string]int, seen map[string]bool) int {
	seen[current] = true
	if current == end {
		if len(seen) == len(distances) {
			return traveled
		} else {
			return -1
		}
	}

	newTravel := -1
	for city, distance := range distances[current] {
		if seen[city] {
			continue
		}

		seenCopy := make(map[string]bool, len(seen))
		for key, value := range seen {
			seenCopy[key] = value
		}

		distanceToEnd := findDistance(city, end, distance, compare, distances, seenCopy)

		if distanceToEnd == -1 {
			continue
		}
		if newTravel == -1 {
			newTravel = distanceToEnd
		} else {
			if compare == "min" {
				newTravel = min(newTravel, distanceToEnd)
			} else if compare == "max" {
				newTravel = max(newTravel, distanceToEnd)
			}
		}
	}

	return traveled + newTravel
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	distances := make(map[string]map[string]int)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " = ")
		cities := strings.Split(parts[0], " to ")
		distance, _ := strconv.Atoi(parts[1])
		for _, city := range cities {
			if distances[city] == nil {
				distances[city] = make(map[string]int)
			}
		}

		distances[cities[0]][cities[1]] = distance
		distances[cities[1]][cities[0]] = distance
	}

	minTravel := 0
	maxTravel := 0
	for city1 := range distances {
		for city2 := range distances {
			if city1 == city2 {
				continue
			}

			minDistance := findDistance(city1, city2, 0, "min", distances, make(map[string]bool))
			maxDistance := findDistance(city1, city2, 0, "max", distances, make(map[string]bool))

			maxTravel = max(maxTravel, maxDistance)
			if minTravel == 0 {
				minTravel = minDistance
			} else {
				minTravel = min(minTravel, minDistance)
			}

		}
	}

	fmt.Println(minTravel)
	fmt.Println(maxTravel)
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func bounds(position1 string, position2 string) ([2]int, [2]int) {
	parts1 := strings.Split(position1, ",")
	parts2 := strings.Split(position2, ",")
	start1, _ := strconv.Atoi(parts1[0])
	start2, _ := strconv.Atoi(parts2[0])
	stop1, _ := strconv.Atoi(parts1[1])
	stop2, _ := strconv.Atoi(parts2[1])

	return [2]int{start1, start2}, [2]int{stop1, stop2}
}

func turnedOn(lights [1000][1000]bool) int {
	count := 0
	for _, row := range lights {
		for _, on := range row {
			if on {
				count++
			}
		}
	}
	return count
}

func brightness(lights [1000][1000]int) int {
	total := 0
	for _, row := range lights {
		for _, light := range row {
			total += light
		}
	}
	return total
}

func main() {
	var p1 [1000][1000]bool
	var p2 [1000][1000]int

	input, _ := os.ReadFile("input.txt")
	r := regexp.MustCompile(`((?:turn (?:on|off)|toggle)) (\d+,\d+) through (\d+,\d+)`)
	matches := r.FindAllSubmatch(input, -1)

	for _, match := range matches {
		command := string(match[1])
		start, stop := bounds(string(match[2]), string(match[3]))

		for x := start[0]; x <= start[1]; x++ {
			for y := stop[0]; y <= stop[1]; y++ {
				switch command {
				case "turn on":
					p1[x][y] = true
					p2[x][y]++
				case "turn off":
					p1[x][y] = false
					p2[x][y] = max(0, p2[x][y]-1)
				case "toggle":
					p1[x][y] = !p1[x][y]
					p2[x][y] += 2
				}
			}
		}
	}

	fmt.Println(turnedOn(p1))
	fmt.Println(brightness(p2))
}

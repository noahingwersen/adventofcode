package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() ([]string, map[string][][2]int) {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)
	var grid []string
	antennas := make(map[string][][2]int)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			char := string(c)
			if char != "." {
				antennas[char] = append(antennas[char], [2]int{x, y})
			}
		}

		grid = append(grid, line)
		y++
	}
	file.Close()
	return grid, antennas
}

func isValid(position [2]int, grid []string) bool {
	return position[0] >= 0 && position[0] < len(grid[0]) && position[1] >= 0 && position[1] < len(grid)
}

func createAntinodes(antennas map[string][][2]int, grid []string, part1 bool) int {
	antinodes := make(map[[2]int]bool)
	for _, locations := range antennas {
		for _, antenna1 := range locations {
			for _, antenna2 := range locations {
				if antenna1[0] == antenna2[0] && antenna1[1] == antenna2[1] {
					continue
				}

				distance := [2]int{antenna2[0] - antenna1[0], antenna2[1] - antenna1[1]}

				// There's probably a way to have a single block handle both cases, but I'm lazy
				if part1 {
					antinode1 := [2]int{antenna1[0] - distance[0], antenna1[1] - distance[1]}
					antinode2 := [2]int{antenna2[0] + distance[0], antenna2[1] + distance[1]}
					for _, antinode := range [][2]int{antinode1, antinode2} {
						if isValid(antinode, grid) {
							antinodes[antinode] = true
						}
					}
				} else {
					// Start at one antenna, and move towards the other antenna. This will make sure antinodes are added between
					// the antenna pair as well as on top of the antenna themselvess
					for antenna, direction := range map[[2]int]int{antenna1: 1, antenna2: -1} {
						added := 0
						position := [2]int{antenna[0] + (distance[0] * direction), antenna[1] + (distance[1] * direction)}
						for isValid(position, grid) {
							antinodes[position] = true
							position = [2]int{position[0] + (distance[0] * direction), position[1] + (distance[1] * direction)}
							added++
						}
					}

				}

			}
		}
	}
	return len(antinodes)
}

func part1() int {
	grid, antennas := parseInput()
	return createAntinodes(antennas, grid, true)
}

func part2() int {
	grid, antennas := parseInput()
	return createAntinodes(antennas, grid, false)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

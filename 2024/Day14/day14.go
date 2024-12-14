package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

const filepath = "input.txt"

func parseInput() [][4]int {
	input, _ := os.ReadFile(filepath)
	r := regexp.MustCompile(`-?\d+`)

	matches := r.FindAll(input, -1)
	var robots [][4]int
	for i := 0; i < len(matches); i += 4 {
		var robot [4]int
		for m := range 4 {
			value, _ := strconv.Atoi(string(matches[i+m]))
			robot[m] = value
		}
		robots = append(robots, robot)
	}
	return robots
}

func addToQuadrant(robot [2]int, quadrants *[4]int, bounds *[2]int) {
	middleX := (*bounds)[0] / 2
	middleY := (*bounds)[1] / 2

	if robot[0] < middleX {
		if robot[1] < middleY {
			(*quadrants)[0]++
		} else if robot[1] > middleY {
			(*quadrants)[1]++
		}
	} else if robot[0] > middleX {
		if robot[1] < middleY {
			(*quadrants)[2]++
		} else if robot[1] > middleY {
			(*quadrants)[3]++
		}
	}
}

func danger(quadrants *[4]int) int {
	return (*quadrants)[0] * (*quadrants)[1] * (*quadrants)[2] * (*quadrants)[3]
}

func propagate(robot *[4]int, seconds int, bounds *[2]int) {
	x := ((*robot)[0] + (*robot)[2]*seconds) % (*bounds)[0]
	y := ((*robot)[1] + (*robot)[3]*seconds) % (*bounds)[1]
	for i, value := range [2][2]int{{x, (*bounds)[0]}, {y, (*bounds)[1]}} {
		if value[0] < 0 {
			value[0] += value[1]
		}
		(*robot)[i] = value[0]
	}
}

func print(robots [][4]int, iteration int) {
	filename := fmt.Sprintf("robots%v.txt", iteration)
	file, _ := os.Create(filename)
	defer file.Close()

	var grid [103][101]int
	for _, r := range robots {
		grid[r[1]][r[0]]++
	}

	for _, row := range grid {
		line := make([]byte, 102)
		line[101] = 10
		for x, value := range row {
			if value > 0 {
				line[x] = 35
			} else {
				line[x] = 46
			}
		}
		file.Write(line)
	}
}

func part1() int {
	robots := parseInput()
	bounds := [2]int{101, 103}

	var quadrants [4]int
	for _, robot := range robots {
		propagate(&robot, 100, &bounds)
		addToQuadrant([2]int{robot[0], robot[1]}, &quadrants, &bounds)
	}

	return danger(&quadrants)
}

func part2() int {
	robots := parseInput()
	bounds := [2]int{101, 103}

	scores := make(map[int][][4]int)
	minScore := math.MaxInt
	minTime := 0
	var i int
	// Robots repeat after about 10k iterations
	for i = 0; i < 12000; i++ {
		quadrants := [4]int{0, 0, 0, 0}
		var grid [103][101]int
		newRobots := make([][4]int, len(robots))
		for r := range robots {
			robot := robots[r]
			propagate(&robot, 1, &bounds)
			addToQuadrant([2]int{robot[0], robot[1]}, &quadrants, &bounds)
			grid[robot[1]][robot[0]]++
			newRobots[r] = robot
		}

		// With my input, the lowest danger score shows the christmas tree. This is not guaranteed for all inputs
		dangerScore := danger(&quadrants)
		scores[dangerScore] = newRobots
		if dangerScore < minScore {
			minScore = dangerScore
			minTime = i + 1
		}
		copy(robots, newRobots)
	}

	print(scores[minScore], minTime)
	return minTime
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

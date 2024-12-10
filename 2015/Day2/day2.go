package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	p1 := 0
	p2 := 0
	for scanner.Scan() {
		minArea := math.MaxInt
		minPerimeter := math.MaxInt
		sides := strings.Split(scanner.Text(), "x")
		volume := 1
		for i := range len(sides) {
			side1, _ := strconv.Atoi(sides[i])
			next := i + 1
			if i == len(sides)-1 {
				next = 0
			}
			side2, _ := strconv.Atoi(sides[next])
			area := side1 * side2
			perimeter := 2*side1 + 2*side2
			volume *= side1
			p1 += 2 * area
			if area < minArea {
				minArea = area
			}
			if perimeter < minPerimeter {
				minPerimeter = perimeter
			}
		}
		p1 += minArea
		p2 += volume + minPerimeter
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

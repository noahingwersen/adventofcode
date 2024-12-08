package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const filepath = "input.txt"

func parseInput() ([]int, []int) {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)

	var slice1 []int
	var slice2 []int
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		id1, err := strconv.Atoi(values[0])
		check(err)
		slice1 = append(slice1, id1)
		id2, err := strconv.Atoi(values[1])
		check(err)
		slice2 = append(slice2, id2)

	}
	file.Close()

	return slice1, slice2
}

func part1() int {
	idList1, idList2 := parseInput()
	slices.Sort(idList1)
	slices.Sort(idList2)

	total := 0
	for i := range len(idList1) {
		distance := math.Abs(float64(idList1[i]) - float64(idList2[i]))
		total += int(distance)
	}

	return total

}

func part2() int {
	idList1, idList2 := parseInput()
	counter := make(map[int]int)
	for _, value := range idList2 {
		counter[value] += 1
	}

	similarity := 0
	for _, value := range idList1 {
		similarity += value * counter[value]
	}

	return similarity
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

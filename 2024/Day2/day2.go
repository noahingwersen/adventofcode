package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const filepath = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() [][]int {
	file, ferr := os.Open(filepath)
	check(ferr)

	scanner := bufio.NewScanner(file)

	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		report := make([]int, len(values))
		for i, value := range values {
			level, err := strconv.Atoi(value)
			check(err)
			report[i] = level
		}

		reports = append(reports, report)
	}
	file.Close()

	return reports
}

func validateReport(report []int) bool {
	// Undeclared ints initialize to 0 which isn't helpful, so I'm using a separate variable "initialized"
	var direction int
	intialized := false

	valid := true
	for i := range len(report) - 1 {
		thisValue := report[i]
		nextValue := report[i+1]

		change := nextValue - thisValue
		absChange := int(math.Abs(float64(change)))
		if !intialized {
			// We need to compute an initial direction, but check the other potential issues
			intialized = true
			if change == 0 || absChange > 3 {
				valid = false
				break
			}
			direction = change / absChange

			continue
		}

		// direction * change will produce a negative number if the new difference changes direction from
		// ascending to descending (or vice versa)
		if direction*change <= 0 || absChange > 3 {
			valid = false
			break
		}

		direction = change / absChange

	}

	return valid
}

// From https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
// The top voted answer modifies the original slice, this option doesn't
func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func part1() int {
	reports := parseInput()
	validReports := 0
	for _, report := range reports {
		if validateReport(report) {
			validReports++
		}
	}
	return validReports
}

func part2() int {
	reports := parseInput()
	validReports := 0
	for _, report := range reports {
		indexToRemove := len(report) - 1
		valid := validateReport(report)
		// If a report isn't valid, start removing single values from the report and try again until it is
		for !valid && indexToRemove >= 0 {
			newReport := removeIndex(report, indexToRemove)
			valid = validateReport(newReport)
			indexToRemove--
		}
		if valid {
			validReports++
		}
	}
	return validReports
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

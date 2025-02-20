package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

func validPassword(password []byte) bool {
	increasingStraight := false
	pairs := make(map[byte]bool)

	for index, char := range password {
		// Passwords may not contain the letters i, o, or l
		for _, invalid := range []byte{'i', 'o', 'l'} {
			if char == invalid {
				return false
			}
		}

		// Passwords must include one increasing straight of at least three letters
		if !increasingStraight && index <= len(password)-3 {
			if password[index+1] == char+1 && password[index+2] == char+2 {
				increasingStraight = true
			}
		}

		// Passwords must contain at least two different, non-overlapping pairs of letters
		if len(pairs) < 2 && index <= len(password)-2 {
			if char == password[index+1] {
				pairs[char] = true
			}
		}

		if increasingStraight && len(pairs) >= 2 {
			return true
		}
	}

	return false
}

func intToPassword(number int) []byte {
	if number == 0 {
		return []byte{97}
	}

	var values []byte
	for number > 0 {
		values = append(values, byte((number%26)+97))
		number /= 26
	}

	slices.Reverse(values)
	return values
}

func passwordToInt(password []byte) int {
	number := 0
	for index, char := range password {
		number += (int(char) - 97) * int(math.Pow(26, float64(len(password)-index-1)))
	}

	return number
}

func findNextPassword(password *[]byte) {
	number := passwordToInt(*password) + 1
	*password = intToPassword(number)
	for !validPassword(*password) {
		number++
		*password = intToPassword(number)
	}
}

func main() {
	input, _ := os.ReadFile("input.txt")
	findNextPassword(&input)
	fmt.Println(string(input))
	findNextPassword(&input)
	fmt.Println(string(input))
}

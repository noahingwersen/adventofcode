package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	p1 := 0
	p2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		unquoted, _ := strconv.Unquote(line)
		quoted := strconv.Quote(line)
		p1 += len(line) - len(unquoted)
		p2 += len(quoted) - len(line)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

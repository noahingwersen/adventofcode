package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	counter := 0
	p1, p2 := false, false
	for !p1 || !p2 {
		counter++
		password := append(input, strconv.Itoa(counter)...)
		hash := md5.Sum(password)
		if !p1 && hex.EncodeToString(hash[:])[:5] == "00000" {
			fmt.Println(counter)
			p1 = true
		}
		if hex.EncodeToString(hash[:])[:6] == "000000" {
			fmt.Println(counter)
			p2 = true
		}
	}
}

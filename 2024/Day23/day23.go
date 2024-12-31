package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

const filepath = "input.txt"

func parseInput() [][]byte {
	file, _ := os.ReadFile(filepath)
	return bytes.Split(file, []byte("\n"))
}

func countNetworks(networks map[string][]string, prefix byte) int {
	seen := make(map[[3]string]bool)
	for computer, connections := range networks {
		for _, connection1 := range connections {
			if connection1 == computer {
				continue
			}

			for _, connection2 := range networks[connection1] {
				if connection2 == computer || connection2 == connection1 {
					continue
				}

				if slices.Contains(networks[connection2], computer) {
					if computer[0] == prefix || connection1[0] == prefix || connection2[0] == prefix {
						network := []string{computer, connection1, connection2}
						slices.Sort(network)
						seen[[3]string(network)] = true
					}
				}
			}
		}
	}
	return len(seen)
}

func part1() int {
	computers := parseInput()
	networks := make(map[string][]string)
	for _, line := range computers {
		parts := strings.Split(string(line), "-")
		networks[parts[0]] = append(networks[parts[0]], parts[1])
		networks[parts[1]] = append(networks[parts[1]], parts[0])
	}

	return countNetworks(networks, 't')
}

func dfs(computer string, network *[]string, connections map[string][]string) {
	for _, connection := range connections[computer] {
		valid := true
		for _, c := range *network {
			if !slices.Contains(connections[c], connection) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		*network = append(*network, connection)
		dfs(connection, network, connections)
	}
}

func part2() string {
	computers := parseInput()
	networks := make(map[string][]string)
	for _, line := range computers {
		parts := strings.Split(string(line), "-")
		networks[parts[0]] = append(networks[parts[0]], parts[1])
		networks[parts[1]] = append(networks[parts[1]], parts[0])
	}

	longestNetwork := make([]string, 0)
	for computer := range networks {
		var network []string
		dfs(computer, &network, networks)
		if len(network) > len(longestNetwork) {
			longestNetwork = network
		}
	}

	slices.Sort(longestNetwork)
	return strings.Join(longestNetwork, ",")
}

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %v in %v\n", part1(), time.Since(start))
	half := time.Now()
	fmt.Printf("Part 2: %v in %v\n", part2(), time.Since(half))
}

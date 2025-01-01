package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type problemData [][2]string

func readData(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var data problemData

	for scan.Scan() {
		p := strings.Split(scan.Text(), "-")
		data = append(data, [2]string(p))
	}

	return data
}

func intersect(a, b []string) []string {
	var output []string

	m := make(map[string]bool, len(b))

	for _, t := range a {
		m[t] = true
	}

	for _, t := range b {
		_, ok := m[t]
		if ok {
			output = append(output, t)
		}
	}

	return output
}

func runPart1(data problemData) int {
	var total int
	nodeConnections := make(map[string][]string)

	for _, pair := range data {
		triples := intersect(nodeConnections[pair[0]], nodeConnections[pair[1]])
		for _, t := range triples {
			if pair[0][0] == 't' ||
				pair[1][0] == 't' ||
				t[0] == 't' {
				total += 1
			}
		}

		nodeConnections[pair[0]] = append(nodeConnections[pair[0]], pair[1])
		nodeConnections[pair[1]] = append(nodeConnections[pair[1]], pair[0])
	}

	return total
}

func runPart2(data problemData) string {
	nodeConnections := make(map[string][]string)
	var networks [][]string
	var biggestNetLength int
	var biggestNet string

	for _, pair := range data {
		triples := intersect(nodeConnections[pair[0]], nodeConnections[pair[1]])
		for _, t := range triples {
			networks = append(networks, []string{pair[0], pair[1], t})
		}

		nodeConnections[pair[0]] = append(nodeConnections[pair[0]], pair[1])
		nodeConnections[pair[1]] = append(nodeConnections[pair[1]], pair[0])

		// This isn't super-efficient and has a 50 second run-time, so could probably be improved
		// with some more efficient data structures.
	netLoop:
		for _, net := range networks {
			for _, node := range net {
				if slices.Index(nodeConnections[pair[0]], node) == -1 {
					continue netLoop
				}
			}

			newNet := make([]string, len(net))
			copy(newNet, net)
			newNet = append(newNet, pair[0])
			networks = append(networks, newNet)

			if biggestNetLength < len(newNet) {
				biggestNetLength = len(newNet)
				slices.Sort(newNet)
				biggestNet = strings.Join(newNet, ",")
			}
		}
	}

	return biggestNet
}

func main() {
	var inputFile = flag.String("input", "inputs/day23.txt", "Problem input file")

	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %s\n", part1, part2)
}

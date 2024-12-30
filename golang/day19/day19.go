package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problemData struct {
	source []string
	target []string
}

func readData(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var data problemData

	scan.Scan()
	data.source = strings.Split(scan.Text(), ", ")

	for scan.Scan() {
		if scan.Text() != "" {
			data.target = append(data.target, scan.Text())
		}
	}

	return data
}

func tryTowel(data problemData, design string) bool {
	if design == "" {
		return true
	}
	for _, towel := range data.source {
		if len(design) >= len(towel) {
			if design[0:len(towel)] == towel {
				if tryTowel(data, design[len(towel):]) {
					return true
				}
			}
		}
	}
	return false
}

func runPart1(data problemData) int {
	total := 0
	for _, design := range data.target {
		if tryTowel(data, design) {
			total++
		}
	}

	return total
}

func tryTowelPart2(data problemData, design string, cache *map[string]int) int {
	// Like part 1 but we return the number of matches intstead of true/false,
	// and run a cache of designs we've already tried.
	_, ok := (*cache)[design]
	if ok {
		return (*cache)[design]
	}

	total := 0
	for _, towel := range data.source {
		if len(design) >= len(towel) {
			if design[0:len(towel)] == towel {
				total += tryTowelPart2(data, design[len(towel):], cache)
			}
		}
	}

	(*cache)[design] = total
	return total
}

func runPart2(data problemData) int {
	total := 0
	cache := map[string]int{
		"": 1,
	}

	for _, design := range data.target {
		total += tryTowelPart2(data, design, &cache)
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day19.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d %d\n", part1, part2)
}

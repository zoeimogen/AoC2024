package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(inputFile string) string {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return (string(data))
}

func runPart1(data string) int {
	var a, b int
	total := 0

	for i := range data {
		_, err := fmt.Sscanf(data[i:], "mul(%d,%d)", &a, &b)
		if err == nil {
			total += a * b
		}
	}

	return total
}

func runPart2(data string) int {
	var a, b int
	total := 0
	enabled := true

	for i := range data {
		if strings.HasPrefix(data[i:], "do()") {
			enabled = true
		} else if strings.HasPrefix(data[i:], "don't()") {
			enabled = false
		} else if enabled {
			_, err := fmt.Sscanf(data[i:], "mul(%d,%d)", &a, &b)
			if err == nil {
				total += a * b
			}
		}
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day03.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

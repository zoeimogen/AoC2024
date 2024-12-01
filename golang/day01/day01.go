package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type problemData struct {
	left, right []int
}

func diffInt(a, b int) int {
	var diff int = a - b
	if diff < 0 {
		return 0 - diff
	}
	return diff
}

func readData(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var left, right int
	var data problemData

	for {
		_, err := fmt.Fscanln(file, &left, &right)

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		data.left = append(data.left, left)
		data.right = append(data.right, right)
	}

	return data
}

func runPart1(data problemData) int {
	left := make([]int, len(data.left))
	right := make([]int, len(data.right))

	copy(left, data.left)
	slices.Sort(left)

	copy(right, data.right)
	slices.Sort(right)

	difference := 0

	for i := range left {
		difference += diffInt(left[i], right[i])
	}

	return difference
}

func runPart2(data problemData) int {
	score := 0

	for _, left := range data.left {
		// Inefficient but not a speed-critical puzzle today
		count := 0
		for _, right := range data.right {
			if left == right {
				count++
			}
		}
		score += left * count
	}

	return score
}

func main() {
	var inputFile = flag.String("input", "inputs/day01.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

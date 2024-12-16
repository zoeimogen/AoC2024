package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type problemData []int

var cache = make(map[int]map[int]int)

func readData(inputFile string) problemData {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var stones []int
	text := strings.Split(string(data), " ")

	for _, t := range text {
		stone, _ := strconv.Atoi(t)
		stones = append(stones, stone)
	}

	return stones
}

func iterateStone(stone int, count int) int {
	if count == 0 {
		return 1
	}
	if stone == 0 {
		return (iterateStone(1, count-1))
	}

	_, ok := cache[count]
	if !ok {
		cache[count] = make(map[int]int)
	}
	r, ok := cache[count][stone]
	if ok {
		return (r)
	}

	if len(strconv.Itoa(stone))%2 == 0 {
		s := strconv.Itoa(stone)
		l := len(s)
		left, _ := strconv.Atoi(s[:l/2])
		right, _ := strconv.Atoi(s[l/2:])
		result := iterateStone(left, count-1) + iterateStone(right, count-1)
		cache[count][stone] = result
		return (result)
	}

	result := iterateStone(stone*2024, count-1)
	cache[count][stone] = result
	return (result)
}

func runProblem(data problemData, count int) int {
	total := 0
	for _, s := range data {
		total += iterateStone(s, count)
	}
	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day10.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runProblem(data, 25)
	runProblem(data, 40)
	runProblem(data, 50)
	runProblem(data, 60)
	part2 := runProblem(data, 75)
	fmt.Printf("%d, %d\n", part1, part2)
}

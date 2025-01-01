package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type problemData []int

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
		i, _ := strconv.Atoi(scan.Text())
		data = append(data, i)
	}

	return data
}

func cycle(secret int) int {
	// This could be optimised but the solution is already sub-second.
	secret = ((secret * 64) ^ secret) % 16777216
	secret = ((secret / 32) ^ secret) % 16777216
	secret = ((secret * 2048) ^ secret) % 16777216
	return secret
}

func runPart1(data problemData) int {
	var total int

	for _, secret := range data {
		s := secret
		for i := 0; i < 2000; i++ {
			s = cycle(s)
		}
		total += s
	}

	return total
}

func runPart2(data problemData) int {
	var totalBananas int

	// Keep track of how many bananas each sequence has generated so far.
	bananas := make([]int, 19*20*20*20)

	for _, secret := range data {
		// Keep track of sequences we've seen for this monkey before, so we can ignore repeats.
		seenSequences := make([]bool, 19*20*20*20)
		s := secret
		var d []int
		for i := 0; i < 2000; i++ {
			olds := s
			s = cycle(s)
			d = append(d, (s%10)-(olds%10))
			if len(d) > 3 {
				bananaKey := ((d[i-3] + 9) * 20 * 20 * 20) +
					((d[i-2] + 9) * 20 * 20) +
					((d[i-1] + 9) * 20) +
					(d[i] + 9)

				if !seenSequences[bananaKey] {
					seenSequences[bananaKey] = true
					bananas[bananaKey] += s % 10
					if bananas[bananaKey] > totalBananas {
						totalBananas = bananas[bananaKey]
					}
				}
			}
		}
	}

	return totalBananas
}

func main() {
	var inputFile = flag.String("input", "inputs/day22.txt", "Problem input file")

	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

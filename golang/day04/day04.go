package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type problemData []string

func searchFrom(data problemData, row int, col int, search_string string, offset_row int, offset_col int) int {
	if search_string == "" {
		return 1
	}

	x := row + offset_row
	y := col + offset_col

	if x < 0 || x >= len(data) {
		return 0
	}

	if y < 0 || y >= len(data[row]) {
		return 0
	}

	if data[x][y] == search_string[0] {
		return searchFrom(data, x, y, search_string[1:], offset_row, offset_col)
	}

	return 0
}

func readData(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var data problemData

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		data = append(data, scan.Text())
	}

	return data
}

func runPart1(data problemData) int {
	total := 0
	searchString := "XMAS"

	for i, row := range data {
		for j, c := range row {
			if byte(c) == searchString[0] {
				total += searchFrom(data, i, j, searchString[1:], -1, -1)
				total += searchFrom(data, i, j, searchString[1:], -1, 0)
				total += searchFrom(data, i, j, searchString[1:], -1, 1)
				total += searchFrom(data, i, j, searchString[1:], 0, 1)
				total += searchFrom(data, i, j, searchString[1:], 0, -1)
				total += searchFrom(data, i, j, searchString[1:], 1, -1)
				total += searchFrom(data, i, j, searchString[1:], 1, 0)
				total += searchFrom(data, i, j, searchString[1:], 1, 1)
			}
		}
	}

	return total
}

func runPart2(data problemData) int {
	total := 0

	for i, row := range data {
		for j, c := range row {
			if i > 0 && i < len(data)-1 && j > 0 && j < len(data)-1 && c == 'A' {
				if data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S' &&
					((data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S') ||
						(data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S')) {
					total += 1
				} else if data[i+1][j+1] == 'M' && data[i-1][j-1] == 'S' &&
					((data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S') ||
						(data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S')) {
					total += 1
				}
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

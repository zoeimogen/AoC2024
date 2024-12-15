package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type problemData [][]int

type rowcol struct {
	row int
	col int
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

	for scan.Scan() {
		var line []int
		for _, c := range scan.Text() {
			i, err := strconv.Atoi(string(c))
			if err == nil {
				line = append(line, i)
			}
		}
		data = append(data, line)
	}

	return data
}

func searchFromPart1(data problemData, row int, col int) []rowcol {
	currentHeight := data[row][col]

	if currentHeight == 9 {
		return []rowcol{rowcol{row, col}}
	}

	var found []rowcol

	if row > 0 && data[row-1][col] == currentHeight+1 {
		found = append(found, searchFromPart1(data, row-1, col)...)
	}
	if col > 0 && data[row][col-1] == currentHeight+1 {
		found = append(found, searchFromPart1(data, row, col-1)...)
	}
	if row < (len(data)-1) && data[row+1][col] == currentHeight+1 {
		found = append(found, searchFromPart1(data, row+1, col)...)
	}
	if col < (len(data[row])-1) && data[row][col+1] == currentHeight+1 {
		found = append(found, searchFromPart1(data, row, col+1)...)
	}

	return found
}

func runPart1(data problemData) int {
	total := 0

	for row := range data {
		for col := range data {
			if data[row][col] == 0 {
				found := searchFromPart1(data, row, col)

				unique := map[rowcol]bool{}
				for _, rc := range found {
					unique[rc] = true
				}

				total += len(unique)
			}
		}
	}

	return total
}

func searchFromPart2(data problemData, row int, col int) int {
	currentHeight := data[row][col]

	if currentHeight == 9 {
		return 1
	}

	total := 0

	if row > 0 && data[row-1][col] == currentHeight+1 {
		total += searchFromPart2(data, row-1, col)
	}
	if col > 0 && data[row][col-1] == currentHeight+1 {
		total += searchFromPart2(data, row, col-1)
	}
	if row < (len(data)-1) && data[row+1][col] == currentHeight+1 {
		total += searchFromPart2(data, row+1, col)
	}
	if col < (len(data[row])-1) && data[row][col+1] == currentHeight+1 {
		total += searchFromPart2(data, row, col+1)
	}

	return total
}

func runPart2(data problemData) int {
	total := 0

	for row := range data {
		for col := range data {
			if data[row][col] == 0 {
				total += searchFromPart2(data, row, col)
			}
		}
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day10.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

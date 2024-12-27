package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type problemData struct {
	warehouse [][]rune
	moves     []rune
	x, y      int
}

func readDataPart1(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var data problemData

	x := 0

	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		var line []rune
		for y, c := range scan.Text() {
			if c == '@' {
				c = '.'
				data.x = x
				data.y = y
			}

			line = append(line, c)
		}
		data.warehouse = append(data.warehouse, line)
		x++
	}

	for scan.Scan() {
		data.moves = append(data.moves, []rune(scan.Text())...)
	}

	return data
}

func readDataPart2(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var data problemData

	x := 0

	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		var line []rune
		for y, c := range scan.Text() {
			if c == '@' {
				line = append(line, '.', '.')
				data.x = x
				data.y = y * 2
			} else if c == '#' {
				line = append(line, '#', '#')
			} else if c == 'O' {
				line = append(line, '[', ']')
			} else if c == '.' {
				line = append(line, '.', '.')
			} else {
				panic("Unknown object in warehouse")
			}
		}
		data.warehouse = append(data.warehouse, line)
		x++
	}

	for scan.Scan() {
		data.moves = append(data.moves, []rune(scan.Text())...)
	}

	return data
}

func moveTo(data problemData, x int, y int, dx int, dy int, dryRun bool) bool {
	if data.warehouse[x][y] == '.' {
		return true
	}

	if data.warehouse[x][y] == '#' {
		return false
	}

	if data.warehouse[x][y] == '[' && dy == 0 {
		if moveTo(data, x+dx, y+dy, dx, dy, true) && moveTo(data, x+dx, y+dy+1, dx, dy, true) {
			if !dryRun {
				moveTo(data, x+dx, y+dy, dx, dy, false)
				moveTo(data, x+dx, y+dy+1, dx, dy, false)
				data.warehouse[x+dx][y+dy] = data.warehouse[x][y]
				data.warehouse[x][y] = '.'
				data.warehouse[x+dx][y+dy+1] = data.warehouse[x][y+1]
				data.warehouse[x][y+1] = '.'
			}
			return true
		}
		return false
	}

	if data.warehouse[x][y] == ']' && dy == 0 {
		if moveTo(data, x+dx, y+dy, dx, dy, true) && moveTo(data, x+dx, y+dy-1, dx, dy, true) {
			if !dryRun {
				moveTo(data, x+dx, y+dy, dx, dy, false)
				moveTo(data, x+dx, y+dy-1, dx, dy, false)
				data.warehouse[x+dx][y+dy] = data.warehouse[x][y]
				data.warehouse[x][y] = '.'
				data.warehouse[x+dx][y+dy-1] = data.warehouse[x][y-1]
				data.warehouse[x][y-1] = '.'
			}
			return true
		}
		return false
	}

	if moveTo(data, x+dx, y+dy, dx, dy, dryRun) {
		if !dryRun {
			data.warehouse[x+dx][y+dy] = data.warehouse[x][y]
			data.warehouse[x][y] = '.'
		}
		return true
	}

	return false
}

func runProblem(data problemData) int {
	for _, move := range data.moves {
		if move == '^' {
			if moveTo(data, data.x-1, data.y, -1, 0, false) {
				data.x--
			}
		} else if move == '>' {
			if moveTo(data, data.x, data.y+1, 0, 1, false) {
				data.y++
			}
		} else if move == 'v' {
			if moveTo(data, data.x+1, data.y, 1, 0, false) {
				data.x++
			}
		} else if move == '<' {
			if moveTo(data, data.x, data.y-1, 0, -1, false) {
				data.y--
			}
		}
	}

	total := 0
	for x := range data.warehouse {
		for y := range data.warehouse[x] {
			if data.warehouse[x][y] == 'O' || data.warehouse[x][y] == '[' {
				total += 100*x + y
			}
		}
	}
	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day15.txt", "Problem input file")
	flag.Parse()
	dataPart1 := readDataPart1(*inputFile)
	part1 := runProblem(dataPart1)
	dataPart2 := readDataPart2(*inputFile)
	part2 := runProblem(dataPart2)
	fmt.Printf("%d, %d\n", part1, part2)
}

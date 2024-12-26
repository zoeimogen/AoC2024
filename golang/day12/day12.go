package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"
)

type problemData [][]rune

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
		var line []rune
		for _, c := range scan.Text() {
			line = append(line, c)
		}
		data = append(data, line)
	}

	return data
}

func getPlot(data problemData, row int, col int) rune {
	if row < 0 || row >= len(data) || col < 0 || col >= len(data[row]) {
		return '.'
	}
	return unicode.ToUpper(data[row][col])
}

func traceEdges(data problemData, plot rune, row int, col int, endRow int, endCol int, direction int, initial bool) int {
	if !initial && direction == 0 && row == endRow && col == endCol {
		return 0
	}

	var nextRow, nextCol, otherRow, otherCol int

	if direction == 0 {
		nextRow = row - 1
		nextCol = col
		otherRow = row - 1
		otherCol = col - 1
		// We only mark vertical edges as visited on the right hand side.
		data[row][col] = unicode.ToLower(data[row][col])
	} else if direction == 1 {
		nextRow = row
		nextCol = col + 1
		otherRow = row - 1
		otherCol = col + 1
	} else if direction == 2 {
		nextRow = row + 1
		nextCol = col
		otherRow = row + 1
		otherCol = col + 1
	} else if direction == 3 {
		nextRow = row
		nextCol = col - 1
		otherRow = row + 1
		otherCol = col - 1
	} else {
		panic("Bad direction")
	}

	nextPlot := getPlot(data, nextRow, nextCol)
	otherPlot := getPlot(data, otherRow, otherCol)

	if nextPlot != plot {
		// Plot that would be next doesn't match, turn clockwise
		return traceEdges(data, plot, row, col, endRow, endCol, dir(direction+1), false) + 1
	} else if otherPlot == plot {
		// Concave corner, turn anticlockwise
		return traceEdges(data, plot, otherRow, otherCol, endRow, endCol, dir(direction-1), false) + 1
	} else {
		// Continue along the edge
		return traceEdges(data, plot, nextRow, nextCol, endRow, endCol, direction, false)
	}
}

func searchFrom(data problemData, dataEdge problemData, plot rune, row int, col int) (int, int, int) {
	// Scan for a plot from a given row/col, marking visited squares by converting to lower case as we go.
	if data[row][col] == unicode.ToLower(plot) {
		// Been here already
		return 0, 0, 0
	}

	if data[row][col] != plot {
		// Gone outside of our plot
		return 0, 1, 0
	}

	data[row][col] = unicode.ToLower(data[row][col])
	area := 1
	perimeter := 0
	edge := 0

	if getPlot(dataEdge, row, col-1) != plot && unicode.IsUpper(dataEdge[row][col]) {
		// An unscanned edge has been found!
		edge += traceEdges(dataEdge, plot, row, col, row, col, 0, true)
	}

	if row > 0 {
		a, p, e := searchFrom(data, dataEdge, plot, row-1, col)
		area += a
		perimeter += p
		edge += e
	} else {
		perimeter++
	}

	if col > 0 {
		a, p, e := searchFrom(data, dataEdge, plot, row, col-1)
		area += a
		perimeter += p
		edge += e
	} else {
		perimeter++
	}

	if row < len(data)-1 {
		a, p, e := searchFrom(data, dataEdge, plot, row+1, col)
		area += a
		perimeter += p
		edge += e
	} else {
		perimeter++
	}

	if col < len(data[row])-1 {
		a, p, e := searchFrom(data, dataEdge, plot, row, col+1)
		area += a
		perimeter += p
		edge += e
	} else {
		perimeter++
	}

	return area, perimeter, edge
}

func runPart1(data problemData) int {
	total := 0

	datacopyArea := make([][]rune, len(data))
	datacopyEdge := make([][]rune, len(data))
	for r := range data {
		datacopyArea[r] = make([]rune, len(data[r]))
		datacopyEdge[r] = make([]rune, len(data[r]))
		copy(datacopyArea[r], data[r])
		copy(datacopyEdge[r], data[r])
	}

	for row := range datacopyArea {
		for col := range datacopyArea {
			if unicode.IsUpper(datacopyArea[row][col]) {
				a, p, _ := searchFrom(datacopyArea, datacopyEdge, datacopyArea[row][col], row, col)
				total += a * p
			}
		}
	}

	return total
}

func dir(d int) int {
	if d < 0 {
		return d + 4
	}
	if d > 3 {
		return d - 4
	}
	return d
}

func runPart2(data problemData) int {
	total := 0

	datacopyArea := make([][]rune, len(data))
	datacopyEdge := make([][]rune, len(data))
	for r := range data {
		datacopyArea[r] = make([]rune, len(data[r]))
		datacopyEdge[r] = make([]rune, len(data[r]))
		copy(datacopyArea[r], data[r])
		copy(datacopyEdge[r], data[r])
	}

	for row := range datacopyArea {
		for col := range datacopyArea {
			if unicode.IsUpper(datacopyArea[row][col]) {
				a, _, e := searchFrom(datacopyArea, datacopyEdge, datacopyArea[row][col], row, col)
				total += a * e
			}
		}
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day12.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

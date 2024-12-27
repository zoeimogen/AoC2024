package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type robot struct {
	px, py, vx, vy int
}

type problemData []robot

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
		var r robot
		_, err := fmt.Sscanf(scan.Text(), "p=%d,%d v=%d,%d", &r.px, &r.py, &r.vx, &r.vy)
		if err == nil {
			data = append(data, r)
		}
	}

	return data
}

func runPart1(data problemData, testGrid bool) int {
	var gridX, gridY int

	quadrants := [4]int{0, 0, 0, 0}

	if testGrid {
		gridX = 11
		gridY = 7
	} else {
		gridX = 101
		gridY = 103
	}

	for _, r := range data {
		new_x := (r.px + r.vx*100) % gridX
		if new_x < 0 {
			new_x += gridX
		}
		new_y := (r.py + r.vy*100) % gridY
		if new_y < 0 {
			new_y += gridY
		}

		if new_x < (gridX / 2) {
			if new_y < (gridY / 2) {
				quadrants[0]++
			} else if new_y > (gridY / 2) {
				quadrants[1]++
			}
		} else if new_x > (gridX / 2) {
			if new_y < (gridY / 2) {
				quadrants[2]++
			} else if new_y > (gridY / 2) {
				quadrants[3]++
			}
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func runPart2(data problemData, showImage bool) int {
	t := 0

	for {
		overlaps := false

		grid := make([][]bool, 103)
		for y := range grid {
			grid[y] = make([]bool, 101)
		}

		for i, r := range data {
			new_x := (r.px + r.vx) % 101
			if new_x < 0 {
				new_x += 101
			}
			new_y := (r.py + r.vy) % 103
			if new_y < 0 {
				new_y += 103
			}
			data[i].px = new_x
			data[i].py = new_y
			if grid[new_y][new_x] {
				overlaps = true
			} else {
				grid[new_y][new_x] = true
			}
		}

		t++

		if !overlaps {
			// Too high a number to view images even at 10 frames a second, so had to guess at
			// "victory" conditions. Stray robots mean we can't check for empty first row/column,
			// but at least for my dataset there are no overlapping robots in the image.
			if showImage {
				fmt.Printf("Time %d\n", t)
				for y := range grid {
					for x := range grid[y] {
						if grid[y][x] {
							print("*")
						} else {
							print(".")
						}
					}
					fmt.Println("")
				}
			}
			return t
		}
	}

}

func main() {
	var inputFile = flag.String("input", "inputs/day14.txt", "Problem input file")
	var testGrid = flag.Bool("test", false, "Use smaller test grid")
	var showImage = flag.Bool("image", false, "Show the image for part 2")
	flag.Parse()

	data := readData(*inputFile)
	part1 := runPart1(data, *testGrid)
	part2 := runPart2(data, *showImage)
	fmt.Printf("%d, %d\n", part1, part2)
}

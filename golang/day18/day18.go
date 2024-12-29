package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type xy struct {
	x int
	y int
}

type problemData []xy

type position struct {
	x, y, steps int
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
		var x, y int
		_, err := fmt.Sscanf(scan.Text(), "%d,%d", &x, &y)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		data = append(data, xy{x, y})
	}

	return data
}

func runMaze(data problemData, maxXY int, inputLen int) int {
	var memory, visited [][]bool
	for y := 0; y <= maxXY; y++ {
		memory = append(memory, make([]bool, maxXY+1))
		visited = append(visited, make([]bool, maxXY+1))
	}

	for _, p := range data[0:inputLen] {
		memory[p.y][p.x] = true
	}

	queue := []position{position{0, 0, 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x == maxXY && p.y == maxXY {
			return p.steps
		}

		if p.y > 0 && !memory[p.y-1][p.x] && !visited[p.y-1][p.x] {
			visited[p.y-1][p.x] = true
			queue = append(queue, position{p.x, p.y - 1, p.steps + 1})
		}
		if p.y < maxXY && !memory[p.y+1][p.x] && !visited[p.y+1][p.x] {
			visited[p.y+1][p.x] = true
			queue = append(queue, position{p.x, p.y + 1, p.steps + 1})
		}
		if p.x > 0 && !memory[p.y][p.x-1] && !visited[p.y][p.x-1] {
			visited[p.y][p.x-1] = true
			queue = append(queue, position{p.x - 1, p.y, p.steps + 1})
		}
		if p.x < maxXY && !memory[p.y][p.x+1] && !visited[p.y][p.x+1] {
			visited[p.y][p.x+1] = true
			queue = append(queue, position{p.x + 1, p.y, p.steps + 1})
		}
	}
	return -1
}

func runPart1(data problemData, test bool) int {
	if test {
		return runMaze(data, 6, 12)
	}
	return runMaze(data, 70, 1024)
}

func runPart2(data problemData, test bool) string {
	for i := 1; i < len(data); i++ {
		var result int
		if test {
			result = runMaze(data, 6, i)
		} else {
			result = runMaze(data, 70, i)
		}
		if result == -1 {
			return fmt.Sprintf("%d,%d", data[i-1].x, data[i-1].y)
		}
	}
	return "No solution found"
}

func main() {
	var inputFile = flag.String("input", "inputs/day18.txt", "Problem input file")
	var test = flag.Bool("test", false, "Use smaller test map")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data, *test)
	part2 := runPart2(data, *test)
	fmt.Printf("%d %s\n", part1, part2)
}

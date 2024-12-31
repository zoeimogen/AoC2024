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

type problemData struct {
	maze       [][]bool
	start, end xy
}

type position struct {
	xy    xy
	steps int
	route []xy
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

	y := 0
	for scan.Scan() {
		var line []bool
		for x, c := range scan.Text() {
			if c == 'S' {
				line = append(line, false)
				data.start = xy{x, y}
			} else if c == 'E' {
				line = append(line, false)
				data.end = xy{x, y}
			} else if c == '.' {
				line = append(line, false)
			} else {
				line = append(line, true)
			}
		}
		data.maze = append(data.maze, line)
		y++
	}

	return data
}

func offset(data problemData, x, y int) int {
	// We use this so that the slice used to keep track of steps required to reach each square/
	// can be 1-dimensional.
	return y*len(data.maze) + x
}

func abs(i int) int {
	if i < 0 {
		return 0 - i
	}
	return i
}

func runSimpleMaze(data problemData) ([]xy, []int) {
	// There is, per the problem statement, only one route through the maze. Get it, and return
	// both the route taken and the steps required to reach each square.
	visited := make([]int, len(data.maze[0])*len(data.maze))
	// Step counts all start from 1 as 0 is used to mark unvisited.
	visited[offset(data, data.start.x, data.start.y)] = 1
	queue := []position{position{data.start, 2, []xy{data.start}}}

	// Given there is only one route, a queue is overkill but a relic of earlier attempts to solve
	// the problem.
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.xy.x == data.end.x && p.xy.y == data.end.y {
			return p.route, visited
		}

		for _, t := range []xy{
			xy{p.xy.x - 1, p.xy.y},
			xy{p.xy.x + 1, p.xy.y},
			xy{p.xy.x, p.xy.y - 1},
			xy{p.xy.x, p.xy.y + 1}} {
			if !data.maze[t.y][t.x] && visited[offset(data, t.x, t.y)] == 0 {
				visited[offset(data, t.x, t.y)] = p.steps
				route := make([]xy, len(p.route))
				copy(route, p.route)
				route = append(route, xy{t.x, t.y})
				queue = append(queue, position{xy{t.x, t.y}, p.steps + 1, route})
			}

		}
	}
	panic("No route found through maze")
}

func runCheatMaze(data problemData, cheats int, timeSaved int, route []xy, scores []int) int {
	// Now for the main event. Go back over the route we found, and for each spot on the route,
	// scan out the number of squares (as Manhattan distance) that we have cheats. These are
	// the squares we can skip to. If any of those skip further ahead than the required minimum
	// time saved, increment the total.
	total := 0
	for _, p := range route {
		for y := 0 - cheats; y <= cheats; y++ {
			for x := 0 - cheats + abs(y); x <= cheats-abs(y); x++ {
				if (x != 0 || y != 0) && p.y+y >= 0 && p.x+x >= 0 && p.y+y < len(data.maze) && p.x+x < len(data.maze[p.y+y]) {
					a := scores[offset(data, p.x, p.y)]
					b := scores[offset(data, p.x+x, p.y+y)]
					s := b - a - abs(x) - abs(y)
					if s >= timeSaved {
						total++
					}
				}
			}
		}
	}
	return total
}

func runPart1(data problemData, test bool) int {
	route, scores := runSimpleMaze(data)
	if test {
		return runCheatMaze(data, 2, 2, route, scores)
	} else {
		return runCheatMaze(data, 2, 100, route, scores)
	}
}

func runPart2(data problemData, test bool) int {
	route, scores := runSimpleMaze(data)
	if test {
		return runCheatMaze(data, 20, 50, route, scores)
	} else {
		return runCheatMaze(data, 20, 100, route, scores)
	}
}

func main() {
	var inputFile = flag.String("input", "inputs/day20.txt", "Problem input file")
	var test = flag.Bool("test", false, "Lower time saving thresholds to 2ps/50ps for test data")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data, *test)
	part2 := runPart2(data, *test)
	fmt.Printf("%d %d\n", part1, part2)
}

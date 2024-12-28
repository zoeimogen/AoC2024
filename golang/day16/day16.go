package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type problemData [][]rune

type xy struct {
	x int
	y int
}

type position struct {
	x, y, direction, score int
	// The copying around of the tile list slows things down considerably, but
	// the part 2 solution is still achievable in under a second.
	tileList []xy
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
		var line []rune
		for _, c := range scan.Text() {
			if c == 'S' || c == 'E' {
				c = '.'
			}

			line = append(line, c)
		}
		data = append(data, line)
	}

	return data
}

func getDelta(direction int) (int, int) {
	if direction == 0 {
		return -1, 0
	}
	if direction == 1 {
		return 0, 1
	}
	if direction == 2 {
		return 1, 0
	}
	return 0, -1
}

func runProblem(data problemData) (int, int) {
	var queue []position
	scores := make([][]int, len(data))
	for d := range data {
		scores[d] = make([]int, len(data[d]))
	}

	// The queue isn't ordered by score, we just run the queue to exhaution at the end. This
	// doesn't take too long, because the moment we have a first solution we can start
	// discarding any higher scoring queue entries.
	queue = append(queue, position{len(data) - 2, 1, 1, 0, []xy{xy{len(data) - 2, 1}}})
	queue = append(queue, position{len(data) - 2, 1, 2, 1000, []xy{xy{len(data) - 2, 1}}})
	queue = append(queue, position{len(data) - 2, 1, 0, 1000, []xy{xy{len(data) - 2, 1}}})
	winningScore := 0
	var winningTileList []xy

	for {
		if len(queue) == 0 {
			// Queue is empty, we reached the end. (Or didn't and will just return zero)
			// But first, count the number of visited tiles for the part 2 solution.
			count := 0
			unique := map[xy]bool{}
			for _, xy := range winningTileList {
				if !unique[xy] {
					count++
					unique[xy] = true
				}
			}
			return winningScore, count
		}
		p := queue[0]
		queue = queue[1:]

		if p.x == 1 && p.y == len(data[0])-2 {
			// Reached the end. Check if we need to record the score.
			if winningScore == 0 || p.score < winningScore {
				// Yes, and this is the first or best result so overwrite the visited tiles list.
				winningScore = p.score
				winningTileList = make([]xy, len(p.tileList))
				copy(winningTileList, p.tileList)
				unique := map[xy]bool{}
				for _, xy := range winningTileList {
					if !unique[xy] {
						unique[xy] = true
					}
				}
			} else if p.score == winningScore {
				// Joint win, append to the visited tiles list.
				winningTileList = append(winningTileList, p.tileList...)
				unique := map[xy]bool{}
				for _, xy := range winningTileList {
					if !unique[xy] {
						unique[xy] = true
					}
				}
			}
		} else if winningScore == 0 || p.score <= winningScore {
			dx, dy := getDelta(p.direction)
			// Skip the move into the new tile if it's a wall, or if it already has a much better score.
			// (Best score for a tile ignores direction, hence an allowance of 1000)
			if data[p.x+dx][p.y+dy] == '.' && (scores[p.x+dx][p.y+dy] == 0 || scores[p.x+dx][p.y+dy]+1000 > p.score) {
				newTileList := make([]xy, len(p.tileList))
				copy(newTileList, p.tileList)
				newTileList = append(newTileList, xy{p.x + dx, p.y + dy})
				queue = append(queue, position{p.x + dx, p.y + dy, p.direction, p.score + 1, newTileList})
				queue = append(queue, position{p.x + dx, p.y + dy, (p.direction + 1) % 4, p.score + 1001, newTileList})
				queue = append(queue, position{p.x + dx, p.y + dy, (p.direction + 3) % 4, p.score + 1001, newTileList})
				scores[p.x+dx][p.y+dy] = p.score + 1
			}
		}
	}
}

func main() {
	var inputFile = flag.String("input", "inputs/day15.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1, part2 := runProblem(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

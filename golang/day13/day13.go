package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

type machine struct {
	ax, ay int
	bx, by int
	px, py int
}

type problemData []machine

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
		var m machine
		_, err := fmt.Sscanf(scan.Text(), "Button A: X+%d, Y+%d", &m.ax, &m.ay)
		if err != nil {
			panic(err)
		}

		scan.Scan()
		_, err = fmt.Sscanf(scan.Text(), "Button B: X+%d, Y+%d", &m.bx, &m.by)
		if err != nil {
			panic(err)
		}

		scan.Scan()
		_, err = fmt.Sscanf(scan.Text(), "Prize: X=%d, Y=%d", &m.px, &m.py)
		if err != nil {
			panic(err)
		}
		data = append(data, m)

		scan.Scan()
	}

	return data
}

func runProblem(data problemData, part2 bool) int {
	total := 0

	for _, m := range data {
		ax := float64(m.ax)
		ay := float64(m.ay)
		bx := float64(m.bx)
		by := float64(m.by)
		px := float64(m.px)
		py := float64(m.py)
		if part2 {
			px += 10000000000000.0
			py += 10000000000000.0
		}
		solution := (px*by - py*bx) / ((ax * by) - (ay * bx))
		if solution >= 0 && math.Trunc(solution) == solution {
			solutionB := (px - solution*ax) / bx
			if solutionB >= 0 && math.Trunc(solutionB) == solutionB {
				total += int(solution*3 + solutionB)
			}
		}
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day13.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runProblem(data, false)
	part2 := runProblem(data, true)
	fmt.Printf("%d, %d\n", part1, part2)
}

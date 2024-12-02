package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type problemData [][]int

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
		var report []int
		text := strings.Split(scan.Text(), " ")

		for _, t := range text {
			v, _ := strconv.Atoi(t)
			report = append(report, v)
		}

		data = append(data, report)
	}

	return data
}

func checkReport(report []int) bool {
	// Check a report for safety
	if report[0] < report[1] {
		// Values increasing
		for i := range report[:len(report)-1] {
			if (report[i+1] > report[i]+3) || (report[i+1] <= report[i]) {
				return false
			}
		}
	} else {
		// Values decreasing
		for i := range report[:len(report)-1] {
			if (report[i+1] < report[i]-3) || (report[i+1] >= report[i]) {
				return false
			}
		}
	}

	return true
}

func runPart1(data problemData) int {
	safe := 0

	for _, report := range data {
		if checkReport(report) {
			safe += 1
		}
	}

	return safe
}

func runPart2(data problemData) int {
	safe := 0

	for _, report := range data {
		if checkReport(report) {
			safe += 1
		} else {
			for i := range report {
				newReport := make([]int, 0, len(report)-1)
				newReport = append(newReport, report[:i]...)
				newReport = append(newReport, report[i+1:]...)
				if checkReport(newReport) {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}

func main() {
	var inputFile = flag.String("input", "inputs/day01.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

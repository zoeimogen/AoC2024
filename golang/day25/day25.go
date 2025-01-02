package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type problemData struct {
	locks [][5]int
	keys  [][5]int
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
		if scan.Text()[0] == '.' {
			// This is a key
			keyData := [5]int{-1, -1, -1, -1, -1}
			level := 5
			for scan.Scan() {
				if scan.Text() == "" {
					break
				}
				for i, c := range scan.Text() {
					if c == '#' && keyData[i] == -1 {
						keyData[i] = level
					}
				}
				level--
			}
			data.keys = append(data.keys, keyData)
		} else {
			// This is a lock
			lockData := [5]int{-1, -1, -1, -1, -1}
			level := 0
			for scan.Scan() {
				if scan.Text() == "" {
					break
				}
				for i, c := range scan.Text() {
					if c == '.' && lockData[i] == -1 {
						lockData[i] = level
					}
				}
				level++
			}
			data.locks = append(data.locks, lockData)
		}
	}

	return data
}

func runPart1(data problemData) int {
	var total int

	for _, key := range data.keys {
	lock:
		for _, lock := range data.locks {
			for pin := range key {
				if key[pin]+lock[pin] > 5 {
					continue lock
				}
			}
			total++
		}
	}
	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day25.txt", "Problem input file")

	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	fmt.Printf("%d\n", part1)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type problemData struct {
	a, b, c, pc   int
	program       []int
	programString string
	output        []int
}

func readData(inputFile string) problemData {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var data problemData

	data.pc = 0
	_, err = fmt.Fscanf(file, "Register A: %d\n", &data.a)
	if err != nil {
		panic("Can't read Register A")
	}
	_, err = fmt.Fscanf(file, "Register B: %d\n", &data.b)
	if err != nil {
		panic("Can't read Register B")
	}
	_, err = fmt.Fscanf(file, "Register C: %d\n\n", &data.c)
	if err != nil {
		panic("Can't read Register C")
	}

	var program string
	_, err = fmt.Fscanf(file, "Program: %s", &program)
	if err != nil {
		panic("Can't read Program")
	}

	data.programString = program

	for _, p := range strings.Split(program, ",") {
		i, _ := strconv.Atoi(p)
		data.program = append(data.program, i)
	}

	return data
}

func combo(data problemData) int {
	if data.program[data.pc+1] >= 0 && data.program[data.pc+1] <= 3 {
		return data.program[data.pc+1]
	} else if data.program[data.pc+1] == 4 {
		return data.a
	} else if data.program[data.pc+1] == 5 {
		return data.b
	} else if data.program[data.pc+1] == 6 {
		return data.c
	}
	panic("Bad combo operand")
}

func adv(data *problemData) {
	data.a = data.a >> combo(*data)
	data.pc += 2
}

func bxl(data *problemData) {
	data.b = data.b ^ data.program[data.pc+1]
	data.pc += 2
}

func bst(data *problemData) {
	data.b = combo(*data) % 8
	data.pc += 2
}

func jnz(data *problemData) {
	if data.a == 0 {
		data.pc += 2
		return
	}
	data.pc = data.program[data.pc+1]
}

func bxc(data *problemData) {
	data.b = data.b ^ data.c
	data.pc += 2
}

func out(data *problemData) {
	data.output = append(data.output, combo(*data)%8)
	data.pc += 2
}

func bdv(data *problemData) {
	data.b = data.a >> combo(*data)
	data.pc += 2
}

func cdv(data *problemData) {
	data.c = data.a >> combo(*data)
	data.pc += 2
}

func runProgram(data problemData) []int {
	for data.pc < len(data.program) {
		if data.program[data.pc] == 0 {
			adv(&data)
		} else if data.program[data.pc] == 1 {
			bxl(&data)
		} else if data.program[data.pc] == 2 {
			bst(&data)
		} else if data.program[data.pc] == 3 {
			jnz(&data)
		} else if data.program[data.pc] == 4 {
			bxc(&data)
		} else if data.program[data.pc] == 5 {
			out(&data)
		} else if data.program[data.pc] == 6 {
			bdv(&data)
		} else if data.program[data.pc] == 7 {
			cdv(&data)
		} else {
			panic("Bad instruction")
		}
	}

	return data.output
}

func runPart1(data problemData) string {
	result := runProgram(data)

	var output []string
	for _, i := range result {
		output = append(output, strconv.Itoa(i))
	}

	return strings.Join(output, ",")
}

func solveForPosition(data problemData, position int, init int) int {
	// Part 2 is a bit of a pain and requires examining the program output,
	// although the 3-bit nature of the program and the fact that register
	// A can only ever be shifted right are clues. The output value length
	// also depends on the input length - increasing by one for every three
	// bits in the input. This means the solution is a 48 bit number (given
	// a 16 instruction program) which is far too big to brute force even at
	// millions of runs a second.

	// The solution is that the last output digit can only be influenced by
	// the first three bits of register A as that is all that is left by that
	// point in the execution. So we find a value of the first three bits that
	// gives the desired last digit then repeat for the next three bits and
	// last but one output position and so on. Occasionally we don't find a
	// suitable output at later positions so need to back out to the previous
	// position and try again with a different value.

	i := init
	increment := 1 << (position * 3)

	for j := 0; j < 8; j++ {
		data.a = i
		result := runProgram(data)

		if len(result) == len(data.program) {
			if result[position] == data.program[position] {
				if position == 0 {
					// This is the last position, solution found
					return i
				}
				nextPosition := solveForPosition(data, position-1, i)
				if nextPosition != -1 {
					return nextPosition
				}
			}
		}
		i += increment
	}

	// No match, back out to the previous position.
	return -1
}

func runPart2(data problemData) int {
	i := 0
	position := len(data.program) - 1

	for {
		i = solveForPosition(data, position, i)
		if i == -1 {
			panic("No solution found")
		}
		if position == 0 {
			return i
		}
		position--
	}
}

func main() {
	var inputFile = flag.String("input", "inputs/day17.txt", "Problem input file")
	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%s %d\n", part1, part2)
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"
)

type pin struct {
	isSet         bool
	value         bool
	attachedGates []*gate
	upstreamGates []*gate
}

type gate struct {
	inputs    [2]string
	operation string
	output    string
	isSet     bool
	value     bool
}

type problemData struct {
	pins  map[string]pin
	gates []*gate
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
	data.pins = make(map[string]pin)

	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		pinData := strings.Split(scan.Text(), ": ")
		pinValue := false
		if pinData[1] == "1" {
			pinValue = true
		}
		data.pins[pinData[0]] = pin{
			true,
			pinValue,
			[]*gate{},
			[]*gate{}}
	}

	for scan.Scan() {
		gateData := strings.Split(scan.Text(), " ")

		newGate := gate{
			[2]string{gateData[0], gateData[2]},
			gateData[1],
			gateData[4],
			false,
			false,
		}

		data.gates = append(data.gates, &newGate)

		// Upstream pin/gate creation andupdate
		if g, ok := data.pins[gateData[4]]; !ok {
			data.pins[gateData[4]] = pin{
				false,
				false,
				[]*gate{},
				[]*gate{&newGate}}
		} else {
			g.upstreamGates = append(g.upstreamGates, &newGate)
			data.pins[gateData[4]] = g
		}

		// Downstream pin/gate creation and update
		for _, p := range []string{gateData[0], gateData[2]} {
			if g, ok := data.pins[p]; !ok {
				data.pins[p] = pin{
					false,
					false,
					[]*gate{&newGate},
					[]*gate{}}
			} else {
				g.attachedGates = append(g.attachedGates, &newGate)
				data.pins[p] = g
			}
		}

	}

	return data
}

func doGate(data *problemData, g *gate) bool {
	if !data.pins[g.inputs[0]].isSet || !data.pins[g.inputs[1]].isSet {
		return false
	}

	if data.pins[g.output].isSet {
		return false
	}

	a := data.pins[g.inputs[0]].value
	b := data.pins[g.inputs[1]].value

	var result bool
	if g.operation == "AND" {
		result = a && b
	} else if g.operation == "XOR" {
		result = a != b
	} else if g.operation == "OR" {
		result = a || b
	}

	g.isSet = true
	g.value = result

	p := data.pins[g.output]
	p.isSet = true
	p.value = result
	data.pins[g.output] = p

	for _, nextGate := range data.pins[g.output].attachedGates {
		doGate(data, nextGate)
	}

	return true

}

func runPart1(data problemData) int {
	zPins := []string{}
	for p := range maps.Keys(data.pins) {
		if p[0] == 'z' {
			zPins = append(zPins, p)
		}
	}

	sort.Strings(zPins)
	slices.Reverse(zPins)

	for _, g := range data.gates {
		doGate(&data, g)
	}

	var result int
	for _, p := range zPins {
		result <<= 1
		if data.pins[p].value {
			result++
		}
	}
	return result
}

func findPin(data *problemData, p string, operator string) (string, string) {
	for _, g := range data.pins[p].attachedGates {
		if g.operation == operator {
			var otherPin string
			if g.inputs[0] == p {
				otherPin = g.inputs[1]
			} else {
				otherPin = g.inputs[0]
			}
			return g.output, otherPin
		}
	}

	return "", ""
}

func runPart2(data problemData) string {
	// Logic gates are an unobfuscated add-with-carry, so we can scan through pin by pin looking
	// for obvious mismatches. Some assumptions are made to keep the solution simples.
	zPins := []string{}
	for p := range maps.Keys(data.pins) {
		if p[0] == 'z' {
			zPins = append(zPins, p)
		}
	}

	sort.Strings(zPins)
	var carry string
	var output []string

	for bit, zPin := range zPins {
		if bit == 0 {
			// Assume z00 is OK
			carry, _ = findPin(&data, "y00", "AND")
			continue
		}
		yPin := fmt.Sprintf("y%02d", bit)

		pendingSwap := ""

		//////////////////////////////////////////////////
		// Check the two XOR gates for the sum are correct
		sumPin, _ := findPin(&data, yPin, "XOR")
		if sumPin == "" {
			// Assume the final z pin output is OK
			break
		}

		outPin, sumPinB := findPin(&data, carry, "XOR")
		if outPin == "" {
			panic("Can't find sum gate, previous carry was wrong")
		} else if sumPin != sumPinB {
			// First XOR wrong: Got sumPin, should be sumPinB
			output = append(output, sumPin)
			if pendingSwap == "" {
				pendingSwap = sumPin
			} else {
				pendingSwap = ""
			}
		}

		if outPin != zPin {
			// Output XOR wrong: should always be zPin
			output = append(output, outPin)
			if pendingSwap == "" {
				pendingSwap = outPin
			} else {
				pendingSwap = ""
			}
		}

		////////////////
		// Carry section
		carryA, _ := findPin(&data, carry, "AND")
		carryB, _ := findPin(&data, yPin, "AND")
		if carryB == "" {
			panic("No gate AND with yPin")
		}

		if carryA == "" {
			panic("No gate AND with carry")
		}

		var checkCarryB string
		carry, checkCarryB = findPin(&data, carryA, "OR")

		if carry == "" {
			carry, _ = findPin(&data, carryB, "OR")
			if carry == "" {
				panic("Can't find third carry gate")
			}

			// Carry A wrong
			output = append(output, carryA)
			if pendingSwap == "" {
				pendingSwap = carryA
			} else {
				pendingSwap = ""
			}
		} else if carryB != checkCarryB {
			// Carry B wrong
			output = append(output, carryB)
			if pendingSwap == "" {
				pendingSwap = carryB
			} else {
				pendingSwap = ""
			}
		}

		if pendingSwap != "" {
			// Detecting incorrect carry is harder, but the problem seems to only swap one pair
			// within a z-pin output block so just assume we got it wrong if we only found one
			// swap so far.
			output = append(output, carry)
			carry = pendingSwap
		}
	}

	sort.Strings(output)
	return strings.Join(output, ",")
}

func main() {
	var inputFile = flag.String("input", "inputs/day24.txt", "Problem input file")

	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %s\n", part1, part2)
}

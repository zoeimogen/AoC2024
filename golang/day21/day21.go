package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type padLayout [][]rune

var padLayoutA = padLayout{
	[]rune{'7', '8', '9'},
	[]rune{'4', '5', '6'},
	[]rune{'1', '2', '3'},
	[]rune{' ', '0', 'A'},
}

var padLayoutB = padLayout{
	[]rune{' ', '^', 'A'},
	[]rune{'<', 'v', '>'},
}

type Pad struct {
	padLayout padLayout
	padMoves  map[rune]map[rune]string
	cache     map[string]int
}

type problemData []string

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

		data = append(data, scan.Text())
	}

	return data
}

func (p Pad) getMove(sx, sy, dx, dy int) string {
	// One of the two keys to solving this is getting the right set of keypad moves. We always
	// want to chain identical moves (eg always v<< or <<v, not <v<) as more efficient for the
	// next robot in the stack which can then just keep hammering A.

	// But we also want to operate the keypad from left-to-right if we can, because that also
	// minimises moves for robots further up the stack which then can get all the more expensive
	// (because further from A) leftward moves done at the start, and then just do the cheaper
	// up/down/right moves next.

	// Similarly, we prefer to do down before right when possible as more expensive (two moves
	// from A versus one for up/right)

	if dy == sy && dx == sx {
		return "A"
	}
	if dy > sy {
		if dx < sx {
			if p.padLayout[sy][dx] != ' ' {
				return "<<<"[:sx-dx] + "vvv"[:dy-sy] + "A"
			} else {
				return "vvv"[:dy-sy] + "<<<"[:sx-dx] + "A"
			}
		} else if dx >= sx && p.padLayout[dy][sx] != ' ' {
			return "vvv"[:dy-sy] + ">>>"[:dx-sx] + "A"
		}
	} else if dy <= sy {
		if dx < sx {
			if p.padLayout[sy][dx] != ' ' {
				return "<<<"[:sx-dx] + "^^^"[:sy-dy] + "A"
			} else {
				return "^^^"[:sy-dy] + "<<<"[:sx-dx] + "A"
			}
		} else if p.padLayout[dy][sx] != ' ' {
			return "^^^"[:sy-dy] + ">>>"[:dx-sx] + "A"
		}
	}

	if dx < sx {
		if dy > sy && p.padLayout[dy][sx] != ' ' {
			return "<<<"[:sx-dx] + "vvv"[:dy-sy] + "A"
		} else if dy < sy {
			return "<<<"[:sx-dx] + "^^^"[:sy-dy] + "A"
		}
	} else if dx > sx {
		if dy > sy {
			return ">>>"[:dx-sx] + "vvv"[:dy-sy] + "A"
		} else if dy < sy {
			return ">>>"[:dx-sx] + "^^^"[:sy-dy] + "A"
		}
	}
	panic("No option")
}

func (p *Pad) initPad(pad padLayout) {
	p.padMoves = make(map[rune]map[rune]string)
	p.padLayout = pad
	p.cache = make(map[string]int)

	for sy := range pad {
		for sx, skey := range pad[sy] {

			p.padMoves[skey] = make(map[rune]string)

			if skey != ' ' {
				for dy := range pad {
					for dx, dkey := range pad[dy] {
						if dkey != ' ' {
							p.padMoves[skey][dkey] = p.getMove(sx, sy, dx, dy)
						}
					}
				}
			}
		}
	}
}

func (p *Pad) usePad(input string) string {
	var output string

	currentKey := 'A'
	for _, k := range input {
		output += p.padMoves[currentKey][k]
		currentKey = k
	}

	return output
}

func (p *Pad) usePadDepthFirst(input string, depth int) int {
	// This is the other key to the solution. The part 2 results are ~100GB long, so take far
	// too long to generate in full. Instead we just return the length of moves required to
	// generate the given sequence, which requires a depth first search. Even that is slow,
	// so we cache results.
	var output int

	cacheKey := strconv.Itoa(depth) + input
	_, ok := p.cache[cacheKey]
	if ok {
		return p.cache[cacheKey]
	}

	currentKey := 'A'
	for _, k := range input {
		if depth == 1 {
			output += len(p.padMoves[currentKey][k])
		} else {
			output += p.usePadDepthFirst(p.padMoves[currentKey][k], depth-1)
		}
		currentKey = k
	}
	p.cache[cacheKey] = output
	return output
}

func runPart1(data problemData) int {
	// We could now solve part 1 using the part 2 code but this was the initial solution.
	var total int
	var padA, padB Pad
	padA.initPad(padLayoutA)
	padB.initPad(padLayoutB)

	for _, input := range data {
		resultA := padA.usePad(input)
		resultB := padB.usePad(resultA)
		resultC := padB.usePad(resultB)
		code, _ := strconv.Atoi(input[0:3])
		total += code * len(resultC)
	}

	return total
}

func runPart2(data problemData) int {
	var total int
	var padA, padB Pad
	padA.initPad(padLayoutA)
	padB.initPad(padLayoutB)

	for _, input := range data {
		result := padA.usePad(input)
		r := padB.usePadDepthFirst(result, 25)
		code, _ := strconv.Atoi(input[0:3])
		total += code * r
	}

	return total
}

func main() {
	var inputFile = flag.String("input", "inputs/day21.txt", "Problem input file")

	flag.Parse()
	data := readData(*inputFile)
	part1 := runPart1(data)
	part2 := runPart2(data)
	fmt.Printf("%d, %d\n", part1, part2)
}

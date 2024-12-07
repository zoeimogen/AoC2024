#!/usr/bin/env python3
'''Advent of Code 2024 Day 6 solution'''
import argparse
import copy

def read_data(filename):
    '''Read input data'''
    data = { "map": [],
             "start": None }

    with open(filename, 'r', encoding="utf8") as file:
        for line in file:
            if "^" in line:
                data["start"] = (len(data["map"]), line.index("^"))
                line = line.replace("^", "X")
            data["map"].append(line.rstrip())

    return data

def runpart1(data):
    '''Run part one'''
    direction = 0
    location = data["start"]
    this_map = copy.deepcopy(data["map"])
    while True:
        if direction == 0:
            new_location = (location[0] - 1, location[1])
        elif direction == 1:
            new_location = (location[0], location[1] + 1)
        elif direction == 2:
            new_location = (location[0] + 1, location[1])
        else:
            new_location = (location[0], location[1] - 1)

        if (new_location[0] < 0 or new_location[0] >= len(this_map[0]) or
            new_location[1] < 0 or new_location[1] >= len(this_map)):
            break

        if this_map[new_location[0]][new_location[1]] != "#":
            location = new_location
            this_map[location[0]] = (this_map[location[0]][:location[1]] +
                                     "X" + 
                                     this_map[location[0]][location[1]+1:])
        else:
            direction += 1
            if direction == 4:
                direction = 0

    return(sum((len([c for c in line if c == "X"])
                 for line in this_map)))

def check_loop(this_map, start):
    '''Check for loops on a map'''
    # We run through the map storing where we've been in the high nibble (0) and the directions
    # visited in the low nibble of each location. I suspect the use of strings is less efficient
    # than proper arrays of ints here, but "fast enough" applies.
    location = start
    direction = 0

    while True:
        if direction == 0:
            new_location = (location[0] - 1, location[1])
        elif direction == 1:
            new_location = (location[0], location[1] + 1)
        elif direction == 2:
            new_location = (location[0] + 1, location[1])
        else:
            new_location = (location[0], location[1] - 1)

        if (new_location[0] < 0 or new_location[0] >= len(this_map[0]) or
            new_location[1] < 0 or new_location[1] >= len(this_map)):
            # Exited the map; no loop
            return False

        if this_map[new_location[0]][new_location[1]] != "#":
            direction_bit = 1 << direction
            if ord(this_map[new_location[0]][new_location[1]]) & ord('0') == ord('0'):
                # Been here before
                if ord(this_map[new_location[0]][new_location[1]]) & direction_bit:
                    # Been here before in this direction: loop
                    return True
                new_marker = ord(this_map[new_location[0]][new_location[1]]) | direction_bit
            else:
                new_marker = ord('0') | direction_bit

            location = new_location
            this_map[location[0]] = (this_map[location[0]][:location[1]] +
                                     chr(new_marker) +
                                     this_map[location[0]][location[1]+1:])
        else:
            direction += 1
            if direction == 4:
                direction = 0

def runpart2(data):
    '''Run part two'''
    # We could speed this up by only checking locations that are visited in part one (As otherwise
    # putting an obstacle there would have no effect)
    total = 0
    for row, _ in enumerate(data["map"]):
        for col, _ in enumerate(data["map"][row]):
            this_map = copy.deepcopy(data["map"])
            if this_map[row][col] == ".":
                this_map[row] = this_map[row][0:col] + "#" + this_map[row][col+1:]
            if check_loop(this_map, data["start"]):
                total += 1

    return total

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day06.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)

    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

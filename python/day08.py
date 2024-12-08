#!/usr/bin/env python3
'''Advent of Code 2024 Day 8 solution'''
import argparse
from collections import defaultdict

def read_data(filename):
    '''Read input data'''
    data = { "nodes": defaultdict(list),
             "dimensions": None }

    with open(filename, 'r', encoding="utf8") as file:
        for row, line in enumerate(file):
            for col, c in enumerate(line.rstrip()):
                if c != ".":
                    data["nodes"][c].append((row, col))

    data["dimensions"] = (row+1, len(line)) # pylint: disable=undefined-loop-variable
    return data


def iterator(data):
    '''Get all tx/pairs across all frequencies'''
    for frequency in data["nodes"]:
        for tx in data["nodes"][frequency]:
            for pair in data["nodes"][frequency]:
                if tx[0] != pair[0] or tx[1] != pair[1]:
                    yield (tx, pair)


def runpart1(data):
    '''Run part one'''
    antinodes = set()

    for tx, pair in iterator(data):
        antinode = (tx[0]*2 - pair[0], tx[1]*2 - pair[1])
        if (antinode[0] >= 0 and antinode[0] < data["dimensions"][0] and
            antinode[1] >= 0 and antinode[1] < data["dimensions"][1]):
            antinodes.add(antinode)

    return len(antinodes)


def runpart2(data):
    '''Run part two'''
    antinodes = set()

    for tx, pair in iterator(data):
        offset = (tx[0] - pair[0], tx[1] - pair[1])
        antinode = tx
        while True:
            if (antinode[0] < 0 or antinode[0] >= data["dimensions"][0] or
                antinode[1] < 0 or antinode[1] >= data["dimensions"][1]):
                break
            antinodes.add(antinode)
            antinode = (antinode[0] + offset[0], antinode[1] + offset[1])

    return len(antinodes)


def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day08.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)
    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)


if __name__ == '__main__':
    print(main())

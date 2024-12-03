#!/usr/bin/env python3
'''Advent of Code 2024 Day 3 solution'''
import argparse
import re

# Surprisingly efficient to use regexps with ^, the runtime is about 20ms on the full problem.
MUL = re.compile(r"^mul\(([0-9]+),([0-9]+)\)")
DO = re.compile(r"^do\(\)")
DONT = re.compile(r"^don't\(\)")

def runpart1(data):
    '''Run part one'''
    total = 0
    for i in range(len(data)):
        r = MUL.match(data[i-1:])
        if r:
            total += int(r.group(1)) * int(r.group(2))

    return total

def runpart2(data):
    '''Run part two'''
    enabled = True

    total = 0
    for i in range(len(data)):
        r = MUL.match(data[i-1:])
        if r and enabled:
            total += int(r.group(1)) * int(r.group(2))
        elif DO.match(data[i-1:]):
            enabled = True
        elif DONT.match(data[i-1:]):
            enabled = False

    return total

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day03.txt")
    args = parser.parse_args()

    # Read inputa file
    with open(args.filename, 'r', encoding="utf8") as file:
        data = file.read()

    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

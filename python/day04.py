#!/usr/bin/env python3
# pylint: disable=too-many-boolean-expressions
'''Advent of Code 2024 Day 4 solution'''
import argparse

def search_from(data, row, col, search_string, offset):
    '''Do the part one search in a certain direction (offset)'''
    if search_string == "":
        return 1

    x = row + offset[0]
    y = col + offset[1]

    if x < 0 or x >= len(data):
        return 0

    if y < 0 or y >= len(data[row]):
        return 0

    if data[x][y] == search_string[0]:
        return search_from(data, x, y, search_string[1:], offset)

    return 0

def runpart1(data):
    '''Run part one'''
    total = 0
    search_string = "XMAS"

    for i, row in enumerate(data):
        for j, c in enumerate(row):
            if c == search_string[0]:
                for offset in ((-1, -1), (-1, 0), (-1,  1),
                               (0,  -1),          (0,   1),
                               (1,  -1), ( 1, 0), ( 1,  1)):
                    total += search_from(data, i, j, search_string[1:], offset)

    return total

def runpart2(data):
    '''Run part two'''
    total = 0

    for i, row in enumerate(data):
        for j, c in enumerate(row):
            if i > 0 and j > 0 and c == 'A':
                try:
                    if (data[i-1][j-1] == 'M' and data[i+1][j+1] == 'S' and
                       ((data[i+1][j-1] == 'M' and data[i-1][j+1] == 'S') or
                        (data[i-1][j+1] == 'M' and data[i+1][j-1] == 'S'))):
                        total += 1
                    elif (data[i+1][j+1] == 'M' and data[i-1][j-1] == 'S' and
                       ((data[i+1][j-1] == 'M' and data[i-1][j+1] == 'S') or
                        (data[i-1][j+1] == 'M' and data[i+1][j-1] == 'S'))):
                        total += 1
                except IndexError:
                    pass

    return total

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day04.txt")
    args = parser.parse_args()

    # Read input file
    with open(args.filename, 'r', encoding="utf8") as file:
        data = file.readlines()

    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

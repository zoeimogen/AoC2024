#!/usr/bin/env python3
'''Advent of Code 2024 Day 1 solution'''
import argparse
import copy

def read_data(filename):
    '''Read input data'''
    left_list = []
    right_list = []

    with open(filename, 'r', encoding="utf8") as file:
        for line in file:
            line_data = line.split()
            left_list.append(int(line_data[0]))
            right_list.append(int(line_data[1]))

    return(left_list, right_list)

def runpart1(data):
    '''Run part one'''
    left = data[0]
    left.sort()

    right = data[1]
    right.sort()

    difference = 0

    for index, l in enumerate(left):
        difference += abs(l - right[index])

    return difference

def runpart2(data):
    '''Run part two'''
    score = 0

    for i in data[0]:
        # Inefficient but not a speed-critical puzzle today
        score += i * len([j for j in data[1] if i == j])

    return score

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day01.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)

    # Run the solution
    part1 = runpart1(copy.deepcopy(data))
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

#!/usr/bin/env python3
'''Advent of Code 2024 Day 7 solution'''
import argparse

def read_data(filename):
    '''Read input data'''
    data = []

    with open(filename, 'r', encoding="utf8") as file:
        for line in file:
            target = int(line.split(":",1)[0])
            values = [int(x) for x in line.split()[1:]]
            data.append((target, values))

    return data

def try_operators(target, calculated, values, part_two=False):
    '''Iterate through all possible combinations of operator looking for a match'''
    if len(values) == 0:
        return target == calculated

    if try_operators(target, calculated + values[0], values[1:], part_two):
        return True
    if try_operators(target, calculated * values[0], values[1:], part_two):
        return True
    if part_two and try_operators(target, int(str(calculated) + str(values[0])),
                                  values[1:], part_two):
        return True
    return False

def runpart1(data):
    '''Run part one'''
    total = 0

    for calibration in data:
        target = calibration[0]
        values = calibration[1]
        if try_operators(target, values[0], values[1:]):
            total += target

    return total

def runpart2(data):
    '''Run part two'''
    total = 0

    for calibration in data:
        target = calibration[0]
        values = calibration[1]
        if try_operators(target, values[0], values[1:], part_two=True):
            total += target

    return total

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day07.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)

    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

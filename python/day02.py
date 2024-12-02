#!/usr/bin/env python3
'''Advent of Code 2024 Day 2 solution'''
import argparse

def read_data(filename):
    '''Read input data'''
    data = []

    with open(filename, 'r', encoding="utf8") as file:
        for report in file:
            levels = report.split()
            data.append([int(i) for i in levels])

    return data


def check_report(report):
    '''Check a report for safety'''
    if report[0] < report[1]:
        # Values increasing
        for i, _ in enumerate(report[:-1]):
            if (report[i+1] > report[i] + 3) or (report[i+1] <= report[i]):
                return False
    else:
        # Values decreasing
        for i, _ in enumerate(report[:-1]):
            if (report[i+1] < report[i] - 3) or (report[i+1] >= report[i]):
                return False

    return True


def runpart1(data):
    '''Run part one'''
    return len([True for report in data if check_report(report)])


def runpart2(data):
    '''Run part two'''
    return len([True
                for report in data
                if check_report(report) or any(True
                                               for i, _ in enumerate(report)
                                               if check_report(report[:i] + report[i+1:]))])


def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day02.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)

    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

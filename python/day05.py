#!/usr/bin/env python3
'''Advent of Code 2024 Day 5 solution'''
import argparse
from math import floor
from collections import defaultdict
from functools import cmp_to_key

def read_data(filename):
    '''Read input data'''
    data = { "pairs": defaultdict(list),
             "pages": []
             }

    with open(filename, 'r', encoding="utf8") as file:
        for line in file:
            if line == "\n":
                break
            pair_text = line.rstrip().split("|")
            data["pairs"][int(pair_text[0])].append(int(pair_text[1]))

        for line in file:
            pages = line.split(",")
            data["pages"].append([int(i) for i in pages])

    return data


def is_correct(pages, pairs):
    '''Old school comparison function for pages (Quicker to write than a proper keys solution)'''
    disallowed = []
    for page in reversed(pages):
        if page in disallowed:
            return False
        if page in pairs:
            disallowed += pairs[page]

    return True

def runpart1(data):
    '''Run part one'''
    total = 0

    for pages in data["pages"]:
        if is_correct(pages, data["pairs"]):
            middle = pages[int(floor(len(pages)/2.0))]
            total += middle

    return total

def runpart2(data):
    '''Run part two'''
    total = 0

    def compare_pages(a, b):
        if a in data["pairs"] and b in data["pairs"][a]:
            return -1
        if b in data["pairs"] and a in data["pairs"][b]:
            return 1
        return 0

    for pages in data["pages"]:
        if not is_correct(pages, data["pairs"]):
            new_order = sorted(pages, key=cmp_to_key(compare_pages))
            middle = new_order[int(floor(len(new_order)/2.0))]
            total += middle

    return total

def main():
    '''Main'''
    parser = argparse.ArgumentParser()
    parser.add_argument("filename", nargs='?', default="inputs/day05.txt")
    args = parser.parse_args()

    # Read input file
    data = read_data(args.filename)
    # Run the solution
    part1 = runpart1(data)
    part2 = runpart2(data)
    return (part1, part2)

if __name__ == '__main__':
    print(main())

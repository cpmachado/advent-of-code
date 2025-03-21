"""
Program that implements a solution for day 1 of the Advent Of Code 2023.
See LICENSE for details.
"""

import argparse
from typing import List


TRANSLATABLE_DIGITS = [
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]


def compute_calibration(digits: List[int]) -> int:
    """
    Computes the calibration value, given digits of a given line
    """
    if digits and len(digits) > 0:
        first, last = digits[0], digits[-1]
        return first * 10 + last
    return 0


def parse_part1(line: str) -> List[int]:
    """
    Parsing function for the line in part 1
    """
    return [int(c) for c in line if c.isdigit()]


def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(compute_calibration(parse_part1(line)) for line in file_handle)


def parse_part2(line: str) -> List[int]:
    """
    Parsing function for the line in part 1
    """
    return [
        x
        for x in [
            int(c)
            if c.isdigit()
            else next(
                (
                    v
                    for v, k in enumerate(TRANSLATABLE_DIGITS)
                    if line[i:].startswith(k)
                ),
                None,
            )
            for i, c in enumerate(line)
        ]
        if x is not None
    ]


def part2(filename: str) -> int:
    """
    Solution of part 2
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(compute_calibration(parse_part2(line)) for line in file_handle)


def main():
    """
    Main function of day 01 of Advent of code 2023
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("file", action="store", type=str)
    parser.add_argument("-p", "--part", action="store", type=int, default=1)
    args = parser.parse_args()
    if args.part == 1:
        print(part1(args.file))
    elif args.part == 2:
        print(part2(args.file))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

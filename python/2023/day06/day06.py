"""
Program that implements a solution for day 6 of the Advent Of Code 2023.
See LICENSE for details.
"""

import argparse
from functools import reduce
from math import sqrt, floor
from operator import mul


def parse_values(line: str) -> list[int]:
    """
    Parse line of values for part1
    """
    return [int(x) for x in line.split(":")[1].strip().split(" ") if x]


def parse_big_value(line: str) -> int:
    """
    Parse big value from line for part2
    """
    return int(line.split(":")[1].strip().replace(" ", ""))


def dmr(t: int, r: int, i: int) -> int:
    """
    Computes distance minus record
    """
    return i * (t - i) - r


def dmr_count(t: int, r: int) -> int:
    """
    Counts how many integers does dmr > 0
    """
    # d(i) = i * (t - i)
    # dmr(i) = d(i) - r -- (d minus record)
    # we want dmr > 0
    # the roots of dmr are i = (t pm sqrt(t^2 - 4 * r)) / 2
    sqrt_of_discriminant = sqrt(t**2 - 4 * r)
    lower_limit = floor((t - sqrt_of_discriminant) / 2)
    upper_limit = floor((t + sqrt_of_discriminant) / 2)
    while dmr(t, r, lower_limit) <= 0:
        lower_limit += 1
    while dmr(t, r, upper_limit) <= 0:
        upper_limit -= 1
    return upper_limit - lower_limit + 1


def part1(filename: str) -> int:
    """
    Part 1 of the challenge
    """
    with open(filename, encoding="utf8") as file_handle:
        times = parse_values(next(file_handle))
        distances = parse_values(next(file_handle))
        return reduce(mul, (dmr_count(*t) for t in zip(times, distances)), 1)


def part2(filename: str) -> int:
    """
    Part 2 of the challenge
    """
    with open(filename, encoding="utf8") as file_handle:
        time = parse_big_value(next(file_handle))
        distance = parse_big_value(next(file_handle))
        return dmr_count(time, distance)


def main():
    """
    Main function of day 5 of Advent of code 2023
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

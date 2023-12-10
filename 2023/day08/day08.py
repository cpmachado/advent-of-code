"""
    Program that implements a solution for day 8 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
import re
from itertools import cycle


def parse_node(line: str) -> tuple[str, tuple[str, str]]:
    name, nodes = line.split(" = ")
    left, right = re.sub("[() \n]", "", nodes).split(",")

    return (name, (left, right))


def part1(filename: str) -> int:
    """
    Part 1 of the challenge
    """
    with open(filename, encoding="utf8") as file_handle:
        path = next(file_handle).strip()
        nodes = {}
        curr = None
        next(file_handle)
        for line in file_handle:
            k, v = parse_node(line)
            nodes[k] = v
        curr = "AAA"
        cnt = 0
        for cmd in cycle(path):
            if curr == "ZZZ":
                break
            cnt += 1
            if cmd == "L":
                curr, _ = nodes[curr]
            else:
                _, curr = nodes[curr]
    return cnt


def main():
    """
    Main function of day 8 of Advent of code 2023
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("file", action="store", type=str)
    parser.add_argument("-p", "--part", action="store", type=int, default=1)
    args = parser.parse_args()
    if args.part == 1:
        print(part1(args.file))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

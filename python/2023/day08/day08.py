"""
    Program that implements a solution for day 8 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
import re
from functools import partial
from itertools import takewhile, cycle
from math import lcm


def parse_node(line: str) -> tuple[str, tuple[str, str]]:
    name, nodes = line.split(" = ")
    left, right = re.sub("[() \n]", "", nodes).split(",")

    return (name, (left, right))


def count_steps(name: str, cmds: list[str], nodes: dict[str, tuple[str]], pred) -> int:
    cnt = 0
    for cmd in cycle(cmds):
        if pred(name):
            break
        if cmd == "L":
            name, _ = nodes[name]
        else:
            _, name = nodes[name]
        cnt += 1
    return cnt


def part1(filename: str) -> int:
    """
    Part 1 of the challenge
    """
    with open(filename, encoding="utf8") as file_handle:
        cmds = next(file_handle).strip()
        nodes = {}
        next(file_handle)
        for line in file_handle:
            k, v = parse_node(line)
            nodes[k] = v
        if "AAA" not in nodes.keys():
            raise RuntimeError("AAA not present")
        return count_steps("AAA", cmds, nodes, (lambda curr: curr == "ZZZ"))


def part2(filename: str) -> int:
    """
    Part 2 of the challenge
    """
    with open(filename, encoding="utf8") as file_handle:
        cmds = next(file_handle).strip()
        nodes = {}
        next(file_handle)
        for line in file_handle:
            k, v = parse_node(line)
            nodes[k] = v
        cnt = []
        return lcm(
            *list(
                count_steps(curr, cmds, nodes, (lambda curr: curr.endswith("Z")))
                for curr in (key for key in nodes.keys() if key.endswith("A"))
            )
        )


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
    elif args.part == 2:
        print(part2(args.file))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

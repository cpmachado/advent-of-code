"""
Program that implements a solution for day 3 of the Advent Of Code 2023.
See LICENSE for details.
"""

import argparse
from dataclasses import dataclass
from functools import reduce
from itertools import takewhile
from operator import mul

GEAR_SYMBOL = "*"


@dataclass
class Coordinate:
    """
    Class to represent coordinates in N^n
    """

    coord: list[int]

    def __add__(self, other):
        """
        Overloader for '+' operator
        """
        return Coordinate([x + y for x, y in zip(self.coord, other.coord)])

    def is_adjacent(self, other) -> bool:
        """
        Check if two coordinates are adjacent
        """
        return all(abs(x - y) <= 1 for x, y in zip(self.coord, other.coord))

    def is_span_adjacent(self, other, cols: int) -> bool:
        """
        Check if a coordinate + (0, 1...n) is adjacent to other coordinate
        """
        return any((self + Coordinate([i, 0])).is_adjacent(other) for i in range(cols))


@dataclass
class EngineSchematic:
    """
    EngineSchematic
    """

    parts: list[tuple[Coordinate, str]]
    symbols: list[tuple[Coordinate, str]]

    def add_line(self, j: int, line: str):
        """
        Parses a line of input
        """
        relevant_chars = [
            (
                Coordinate([i, j]),
                c if not c.isdigit() else "".join(takewhile(str.isdigit, line[i:])),
            )
            for i, c in enumerate(line)
            if c != "."
            and (
                i < 1
                or line[i - 1] == "."
                or c.isdigit()
                and not line[i - 1].isdigit()
                or not c.isdigit()
                and line[i - 1].isdigit()
            )
        ]
        self.parts += [
            (coord, content) for coord, content in relevant_chars if content.isdigit()
        ]
        self.symbols += [
            (coord, content)
            for coord, content in relevant_chars
            if not content.isdigit()
        ]

    def compute_part_sum(self) -> int:
        """
        Computes part sum
        """
        return sum(
            int(val)
            for coord, val in self.parts
            if any(
                coord.is_span_adjacent(coord_symb, len(val))
                for coord_symb, _ in self.symbols
            )
        )

    def compute_gear_ratio_sum(self) -> int:
        """
        Computes gear ratio sum
        """
        return sum(
            reduce(mul, x, 1)
            for x in [
                [
                    int(val)
                    for coord, val in self.parts
                    if coord.is_span_adjacent(gear_coord, len(val))
                ]
                for gear_coord in (
                    coord for coord, content in self.symbols if content == GEAR_SYMBOL
                )
            ]
            if len(x) > 1
        )


def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    schematic = EngineSchematic([], [])
    with open(filename, encoding="utf8") as file_handle:
        for j, line in enumerate(file_handle):
            schematic.add_line(j, line.strip())
    return schematic.compute_part_sum()


def part2(filename: str) -> int:
    """
    Solution of part 2
    """
    schematic = EngineSchematic([], [])
    with open(filename, encoding="utf8") as file_handle:
        for j, line in enumerate(file_handle):
            schematic.add_line(j, line.strip())
    return schematic.compute_gear_ratio_sum()


def main():
    """
    Main function of day 03 of Advent of code 2023
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

"""
    Program that implements a solution for day 3 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
from dataclasses import dataclass
from enum import Enum
from itertools import takewhile


@dataclass
class Coordinate:
    coord: list[int]

    def __add__(self, other):
        return Coordinate([x + y for x, y in zip(self.coord, other.coord)])

    def is_adjacent(self, other) -> bool:
        return all(abs(x - y) <= 1 for x, y in zip(self.coord, other.coord))


@dataclass
class EngineSchematic:
    parts: list[tuple[Coordinate, str]]
    symbols: list[tuple[Coordinate, str]]

    def add_line(self, j: int, line: str):
        relevant_chars = [
            (
                Coordinate([i, j]),
                line[i]
                if not line[i].isdigit()
                else "".join(takewhile(str.isdigit, line[i:])),
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
            coord for coord, content in relevant_chars if not content.isdigit()
        ]

    def compute_sum(self):
        return sum(
            int(val)
            for coord, val in self.parts
            if any(
                any(
                    (coord + Coordinate([i, 0])).is_adjacent(coord_symb)
                    for i in range(len(val))
                )
                for coord_symb in self.symbols
            )
        )


def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    schematic = EngineSchematic([], [])
    with open(filename, encoding="utf8") as file_handle:
        for j, line in enumerate(file_handle):
            schematic.add_line(j, line.strip())
    return schematic.compute_sum()


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
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

"""
    Program that implements a solution for day 4 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
from dataclasses import dataclass
from functools import reduce
from operator import mul


@dataclass
class ScratchCard:
    card_id: int
    winning: list[int]
    numbers: list[int]

    @staticmethod
    def from_line(line: str):
        card_section, numbers_section = line.strip().split(":")
        card_id = int(card_section.strip().split(" ")[-1])
        winning, numbers = [
            [int(x) for x in arr.strip().split(" ") if x]
            for arr in numbers_section.strip().split("|")
        ]
        return ScratchCard(card_id, winning, numbers)

    def score(self):
        n = sum(1 for x in self.winning if x in self.numbers)
        return 2 ** (n - 1) if n > 0 else 0


def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(ScratchCard.from_line(line).score() for line in file_handle)


def main():
    """
    Main function of day 04 of Advent of code 2023
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

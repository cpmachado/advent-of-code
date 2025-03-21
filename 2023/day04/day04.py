"""
    Program that implements a solution for day 4 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
from dataclasses import dataclass
from functools import reduce


@dataclass
class ScratchCard:
    """
    Describes a ScratchCard
    """

    card_id: int
    winning: list[int]
    numbers: list[int]

    @staticmethod
    def from_line(line: str):
        """
        Parses a Line into a ScratchCard
        """
        card_section, numbers_section = line.strip().split(":")
        card_id = int(card_section.strip().split(" ")[-1])
        winning, numbers = [
            [int(x) for x in arr.strip().split(" ") if x]
            for arr in numbers_section.strip().split("|")
        ]
        return ScratchCard(card_id, winning, numbers)

    def number_of_matches(self):
        """
        Number of winning numbers matches
        """
        return sum(1 for x in self.winning if x in self.numbers)

    def score(self) -> int:
        """
        Score of ScratchCard
        """
        number_of_matches = self.number_of_matches()
        return int(2 ** (number_of_matches - 1)) if number_of_matches > 0 else 0


def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(ScratchCard.from_line(line).score() for line in file_handle)


def scratch_reducer(
    acc: tuple[int, tuple[int]], scratch: ScratchCard
) -> tuple[int, tuple[int]]:
    """
    Reducer for Part 2
    """
    val, carry = acc
    multiplier, *tail = carry
    number_of_matches = scratch.number_of_matches()
    if len(tail) == 0:
        tail = [0 for _ in range(len(scratch.winning))]
    new_tail = [
        x if i >= number_of_matches else x + 1 + multiplier
        for i, x in enumerate([*tail, 0])
    ]
    return (val + 1 + multiplier, new_tail)


def part2(filename: str) -> int:
    """
    Solution of part 2
    """
    with open(filename, encoding="utf8") as file_handle:
        val, _ = reduce(
            scratch_reducer,
            (ScratchCard.from_line(line) for line in file_handle),
            (0, [0]),
        )
        return val


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
    elif args.part == 2:
        print(part2(args.file))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

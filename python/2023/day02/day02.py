"""
Program that implements a solution for day 2 of the Advent Of Code 2023.
See LICENSE for details.
"""

import argparse
from dataclasses import dataclass
from functools import reduce
from operator import mul


@dataclass
class CubeBag:
    """
    Describes a bag of cubes
    """

    cubes: dict[str, int]

    @staticmethod
    def from_line(line: str):
        """
        Builds a Cube Bag from a record from the file
        """
        cubes = {}
        for record in line.split(","):
            val, k = record.strip().split(" ")
            cubes[k] = int(val)
        return CubeBag(cubes)

    def has_enough(self, other):
        """
        Checks if a given bag could have been used from a draw of another
        """
        return all(
            k in self.cubes and self.cubes[k] >= v for k, v in other.cubes.items()
        )

    def minimal_common_bag(self, other):
        """
        Computes the minimal bag, that contains enough cubes in both
        """
        return CubeBag(
            dict(
                (k, max(self.cubes.get(k, 0), other.cubes.get(k, 0)))
                for k in set((*self.cubes.keys(), *other.cubes.keys()))
            )
        )


@dataclass
class CubeGame:
    """
    Describes a game of cubes
    """

    game_id: int
    bags: list[CubeBag]

    @staticmethod
    def from_line(line: str):
        """
        Builds a Cube Game from a line of the file
        """
        game_data, cube_data = line.strip().split(":")
        _, game_id = game_data.split(" ")
        bags = [CubeBag.from_line(record) for record in cube_data.split(";")]
        return CubeGame(int(game_id), bags)

    def is_valid(self, bag: CubeBag) -> bool:
        """
        Checks if it is a valid game, based on a reference bag
        """
        return all(bag.has_enough(bagi) for bagi in self.bags)

    def power_set(self):
        """
        Computes the power set of the game
        """
        first, *rest = self.bags
        return reduce(
            mul, reduce(CubeBag.minimal_common_bag, rest, first).cubes.values(), 1
        )


def part1(filename: str, bag: CubeBag) -> int:
    """
    Solution of part 1
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(
            game.game_id
            for line in file_handle
            for game in [CubeGame.from_line(line)]
            if game.is_valid(bag)
        )


def part2(filename: str) -> int:
    """
    Solution of part 2
    """
    with open(filename, encoding="utf8") as file_handle:
        return sum(CubeGame.from_line(line).power_set() for line in file_handle)


def main():
    """
    Main function of day 02 of Advent of code 2023
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("file", action="store", type=str)
    parser.add_argument("-p", "--part", action="store", type=int, default=1)
    parser.add_argument("-r", "--red", action="store", type=int, default=0)
    parser.add_argument("-g", "--green", action="store", type=int, default=0)
    parser.add_argument("-b", "--blue", action="store", type=int, default=0)
    args = parser.parse_args()
    if args.part == 1:
        bag = CubeBag({"red": args.red, "green": args.green, "blue": args.blue})
        print(part1(args.file, bag))
    elif args.part == 2:
        print(part2(args.file))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

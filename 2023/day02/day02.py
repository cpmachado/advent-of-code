"""
    Program that implements a solution for day 2 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
from dataclasses import dataclass


@dataclass
class CubeBag:
    cubes: dict[str, int]

    @staticmethod
    def from_line(line: str):
        cubes = dict()
        for record in line.split(","):
            v, k = record.strip().split(" ")
            cubes[k] = int(v)
        return CubeBag(cubes)

    def has_enough(self, other):
        return all(
            k in self.cubes and self.cubes[k] >= v for k, v in other.cubes.items()
        )


@dataclass
class CubeGame:
    game_id: int
    bags: list[CubeBag]

    @staticmethod
    def from_line(line: str):
        game_data, cube_data = line.strip().split(":")
        _, game_id = game_data.split(" ")
        bags = [CubeBag.from_line(record) for record in cube_data.split(";")]
        return CubeGame(int(game_id), bags)

    def is_valid(self, bag: CubeBag) -> bool:
        return all(bag.has_enough(bagi) for bagi in self.bags)


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
    bag = CubeBag({"red": args.red, "green": args.green, "blue": args.blue})
    if args.part == 1:
        print(part1(args.file, bag))
    else:
        raise NotImplementedError("Not implemented")


if __name__ == "__main__":
    main()

"""
    Program that implements a solution for day 3 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse

def part1(filename: str) -> int:
    """
    Solution of part 1
    """
    with open(filename, encoding="utf8") as file_handle:
        for j, line in enumerate(file_handle):
            pass
    return 0


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

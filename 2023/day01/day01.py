import argparse
from typing import List


def parse_digits(digits: List[str]) -> int:
    if len(digits) == 1:
        a = digits[0]
        return int(a + a)
    elif len(digits) > 1:
        a, *_, b = digits
        return int(a + b)
    return 0


def part1(filename: str) -> int:
    x = []
    with open(filename) as f:
        x = [parse_digits([d for d in line if d.isdigit()]) for line in f]
    return sum(x)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("file", action="store", type=str)
    parser.add_argument("-p", "--part", action="store", type=int, default=1)
    args = parser.parse_args()
    if args.part == 1:
        print(part1(args.file))


if __name__ == "__main__":
    main()

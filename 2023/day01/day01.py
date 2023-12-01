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


def day01(filename: str) -> int:
    x = []
    with open(filename) as f:
        x = [parse_digits([d for d in line if d.isdigit()]) for line in f]
    return sum(x)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("file", action="store", type=str)
    args = parser.parse_args()
    print(day01(args.file))


if __name__ == "__main__":
    main()

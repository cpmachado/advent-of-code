"""
    Program that implements a solution for day 5 of the Advent Of Code 2023.
    See LICENSE for details.
"""
import argparse
from math import inf
from dataclasses import dataclass
from itertools import takewhile


@dataclass
class IRange:
    """An integer Range"""

    a: int
    b: int

    @staticmethod
    def checkedInit(a: int, b: int):
        if a > b:
            return None
        else:
            return IRange(a, b)

    def belongs(self, x: int) -> bool:
        return self.left_limit() <= x <= self.right_limit()

    def intersects(self, other) -> bool:
        return other is not None and (
            self.belongs(other.left_limit()) or other.belongs(self.left_limit())
        )

    def contains(self, other) -> bool:
        return (
            other is None
            or self.belongs(other.left_limit())
            and self.belongs(other.right_limit())
        )

    def intersection(self, other):
        if not self.intersects(other):
            return None
        else:
            a = max(self.left_limit(), other.left_limit())
            b = min(self.right_limit(), other.right_limit())
            return IRange(a, b)

    def complement(self, other) -> list:
        if (intersection := self.intersection(other)) is None:
            return [self.copy()]
        elif intersection.contains(self):
            return None
        else:
            left = IRange.checkedInit(self.left_limit(), intersection.left_limit() - 1)
            right = IRange.checkedInit(
                intersection.right_limit() + 1, self.right_limit()
            )
            if left is None:
                return [right]
            elif right is None:
                return [left]
            else:
                return [left, right]

    def left_limit(self):
        return self.a

    def right_limit(self):
        return self.b

    def shift(self, d: int):
        return IRange(self.left_limit() + d, self.right_limit() + d)

    def copy(self):
        return IRange(self.left_limit(), self.right_limit())


@dataclass
class IRangeMapper:
    rng: IRange
    d: int

    def map(self, rng: IRange) -> (IRange, IRange):
        if not self.applies(rng):
            return (None, rng.copy())
        elif self.rng.contains(rng):
            return (rng.shift(self.d), None)
        else:
            intersection = self.rng.intersection(rng)
            complement = rng.complement(self.rng)
            return (intersection.shift(self.d), complement)

    def applies(self, rng: IRange) -> bool:
        return self.rng.intersects(rng)

    def get_left_limit(self):
        return self.rng.left_limit()


@dataclass
class PropertyMapper:
    name: str
    mappers: list[IRangeMapper]

    def map_ranges(self, rngs: list[IRange]) -> list[IRange]:
        stack = rngs.copy()
        res = []
        while len(stack) > 0:
            rng = stack.pop()
            mapped, complement = self.get_applied_mapper(rng).map(rng)
            res.append(mapped)
            if complement is not None:
                for c in complement:
                    stack.append(c)
        return res

    def get_applied_mapper(self, rng: IRange):
        if (
            mapper := next(
                (mapper for mapper in self.mappers if mapper.applies(rng)), None
            )
        ) is not None:
            return mapper
        return IRangeMapper(IRange(-inf, inf), 0)


def part1(filename: str) -> int:
    with open(filename, encoding="utf-8") as file_handle:
        seeds = [
            IRange(x, x)
            for x in [int(x) for x in next(file_handle, "").strip().split(" ")[1:] if x]
        ]
        maps = []
        for line in file_handle:
            if "map" in line:
                name = line.split(" ")[0]
                f = lambda dest, src, length: IRangeMapper(
                    IRange(src, src + length - 1), dest - src
                )
                maps.append(
                    PropertyMapper(
                        name,
                        sorted(
                            (
                                f(*map(int, line.strip().split(" ")))
                                for line in takewhile(
                                    (lambda line: line.strip()), file_handle
                                )
                            ),
                            key=IRangeMapper.get_left_limit,
                        ),
                    )
                )
        curr = seeds
        for mapper in maps:
            curr = mapper.map_ranges(curr)
        return min((rng.left_limit() for rng in curr))


def part2(filename: str) -> int:
    with open(filename, encoding="utf-8") as file_handle:
        process_seeds = lambda raw_seeds: [
            IRange(a, b)
            for a, b in (
                (raw_seeds[i], raw_seeds[i] + raw_seeds[i + 1] - 1)
                for i in range(0, len(raw_seeds), 2)
            )
        ]
        seeds = process_seeds(
            [int(x) for x in next(file_handle, "").strip().split(" ")[1:] if x]
        )
        maps = []
        for line in file_handle:
            if "map" in line:
                name = line.split(" ")[0]
                f = lambda dest, src, length: IRangeMapper(
                    IRange(src, src + length - 1), dest - src
                )
                maps.append(
                    PropertyMapper(
                        name,
                        sorted(
                            (
                                f(*map(int, line.strip().split(" ")))
                                for line in takewhile(
                                    (lambda line: line.strip()), file_handle
                                )
                            ),
                            key=IRangeMapper.get_left_limit,
                        ),
                    )
                )
        curr = seeds
        for mapper in maps:
            curr = mapper.map_ranges(curr)
        return min((rng.left_limit() for rng in curr))


def main():
    """
    Main function of day 05 of Advent of code 2023
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

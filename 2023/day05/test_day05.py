"""
    Tests for day 05 of Advent of Code 2023
    See LICENSE for details
"""
import os
import unittest
from day05 import IRange

TEST_FILE = os.path.join("data", "test.txt")


class Day05Tests(unittest.TestCase):
    """
    Tests for day 5
    """

    def test_part1(self):
        """
        Tests for day 5 part 1
        """
        actual = part1(TEST_FILE)
        expected = 35
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")

    def test_part2(self):
        """
        Tests for day 5 part 2
        """
        actual = part2(TEST_FILE)
        expected = 46
        self.assertEqual(
            actual, expected, f"Expected {expected} scratch cards, not {actual}"
        )


if __name__ == "__main__":
    unittest.main()

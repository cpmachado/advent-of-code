"""
Tests for day 04 of Advent of Code 2023
See LICENSE for details
"""

import os
import unittest
from day04 import part1, part2

TEST_DIR = os.path.dirname(__file__)
TEST_FILE = os.path.join(TEST_DIR, "example.txt")


class Day04Tests(unittest.TestCase):
    """
    Tests for day 4
    """

    def test_part1(self):
        """
        Tests for day 4 part 1
        """
        actual = part1(TEST_FILE)
        expected = 13
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")

    def test_part2(self):
        """
        Tests for day 4 part 2
        """
        actual = part2(TEST_FILE)
        expected = 30
        self.assertEqual(
            actual, expected, f"Expected {expected} scratch cards, not {actual}"
        )


if __name__ == "__main__":
    unittest.main()

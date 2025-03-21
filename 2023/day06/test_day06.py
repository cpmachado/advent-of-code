"""
    Tests for day 06 of Advent of Code 2023
    See LICENSE for details
"""
import os
import unittest
from day06 import part1, part2

TEST_FILE = os.path.join("data", "test.txt")


class Day06Tests(unittest.TestCase):
    """
    Tests for day 6
    """

    def test_part1(self):
        """
        Tests for day 6 part 1
        """
        actual = part1(TEST_FILE)
        expected = 288
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")

    def test_part2(self):
        """
        Tests for day 6 part 2
        """
        actual = part2(TEST_FILE)
        expected = 71503
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")


if __name__ == "__main__":
    unittest.main()

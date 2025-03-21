"""
Tests for day 08 of Advent of Code 2023
See LICENSE for details
"""

import os
import unittest
from day08 import part1, part2

TEST_FILE_1 = os.path.join("data", "test_1.txt")
TEST_FILE_2 = os.path.join("data", "test_2.txt")
TEST_FILE_3 = os.path.join("data", "test_3.txt")


class Day08Tests(unittest.TestCase):
    """
    Tests for day 8
    """

    def test_part1_1(self):
        """
        Tests for day 8 part 1 - test 1
        """
        actual = part1(TEST_FILE_1)
        expected = 2
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")

    def test_part1_2(self):
        """
        Tests for day 8 part 1 - test 2
        """
        actual = part1(TEST_FILE_2)
        expected = 6
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")

    def test_part2(self):
        """
        Tests for day 8 part 2
        """
        actual = part2(TEST_FILE_3)
        expected = 6
        self.assertEqual(actual, expected, f"Expected {expected} points, not {actual}")


if __name__ == "__main__":
    unittest.main()

"""
    Tests for day 01 of Advent of Code 2023
    See LICENSE for details
"""
import os
import unittest
from day01 import part1, part2

TEST_FILE_PART1 = os.path.join("data", "test_part1.txt")
TEST_FILE_PART2 = os.path.join("data", "test_part2.txt")


class Day01Tests(unittest.TestCase):
    """
    Tests for day 1
    """

    def test_part1(self):
        """
        Tests for day 1 part 1
        """
        actual = part1(TEST_FILE_PART1)
        expected = 142
        self.assertEqual(
            actual, expected, f"Expected calibration values {expected}, not {actual}"
        )

    def test_part2(self):
        """
        Tests for day 1 part 2
        """
        actual = part2(TEST_FILE_PART2)
        expected = 281
        self.assertEqual(
            actual, expected, f"Expected calibration values {expected}, not {actual}"
        )


if __name__ == "__main__":
    unittest.main()

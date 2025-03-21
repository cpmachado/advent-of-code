"""
    Tests for day 03 of Advent of Code 2023
    See LICENSE for details
"""
import os
import unittest
from day03 import part1, part2

TEST_FILE = os.path.join("data", "test.txt")


class Day03Tests(unittest.TestCase):
    """
    Tests for day 3
    """

    def test_part1(self):
        """
        Tests for day 3 part 1
        """
        actual = part1(TEST_FILE)
        expected = 4361
        self.assertEqual(
            actual,
            expected,
            f"Expected engine part sum to be {expected}, not {actual}",
        )

    def test_part2(self):
        """
        Tests for day 3 part 2
        """
        actual = part2(TEST_FILE)
        expected = 467835
        self.assertEqual(
            actual,
            expected,
            f"Expected engine gear ratio sum to be {expected}, not {actual}",
        )


if __name__ == "__main__":
    unittest.main()

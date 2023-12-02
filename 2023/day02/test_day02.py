"""
    Tests for day 02 of Advent of Code 2023
    See LICENSE for details
"""
import os
import unittest
from day02 import part1, CubeBag

TEST_FILE_PART1 = os.path.join("data", "test.txt")


class Day02Tests(unittest.TestCase):
    """
    Tests for day 2
    """

    def test_part1(self):
        """
        Tests for day 2 part 1
        """
        bag = CubeBag({"red": 12, "green": 13, "blue": 14})
        actual = part1(TEST_FILE_PART1, bag)
        expected = 8
        self.assertEqual(
            actual,
            expected,
            f"Expected valid game id sum to be {expected}, not {actual}",
        )


if __name__ == "__main__":
    unittest.main()

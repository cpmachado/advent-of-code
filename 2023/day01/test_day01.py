import os
import unittest
from day01 import part1

TEST_FILE = os.path.join("data", "test_part1.txt")


class Day01Tests(unittest.TestCase):
    def test_part1(self):
        actual = part1(TEST_FILE)
        expected = 142
        self.assertEqual(
            actual, expected, f"Expected calibration values {expected}, not {actual}"
        )

if __name__ == '__main__':
    unittest.main()

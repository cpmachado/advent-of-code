import os
import unittest
from day01 import day01

TEST_FILE = os.path.join("data", "test.txt")


class Day01Tests(unittest.TestCase):
    def test_day01(self):
        actual = day01(TEST_FILE)
        expected = 142
        self.assertEqual(
            actual, expected, f"Expected calibration values {expected}, not {actual}"
        )

if __name__ == '__main__':
    unittest.main()

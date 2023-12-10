from unittest import TestCase

from day04.first.main import first
from input.file import input

example_input = input("input.txt")
daily_input = input("../input.txt")


class TestMain(TestCase):
    def test_example(self):
        out = first(example_input)

        self.assertEqual(13, out)

    def test_specific_row_example(self):
        out = first(input("input_row.txt"))

        self.assertEqual(72, out)

    def test_with_input(self):
        out = first(daily_input)

        self.assertEqual(27845, out)

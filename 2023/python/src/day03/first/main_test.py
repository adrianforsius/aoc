from unittest import TestCase

from day03.first.main import first
from input.file import input

example_input = input("input.txt")
daily_input = input("../input.txt")


class TestMain(TestCase):
    def test_example(self):
        out = first(example_input)

        self.assertEqual(4361, out)

    def test_numer_last_in_row_example(self):
        out = first(example_input)

        self.assertEqual(4361, out)

    def test_with_input(self):
        out = first(daily_input)

        self.assertEqual(517021, out)

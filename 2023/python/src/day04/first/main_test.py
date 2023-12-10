from unittest import TestCase

from day01.first.main import first
from input.file import input

example_input = input("input.txt")
daily_input = input("../input.txt")


class TestMain(TestCase):
    def test_example(self):
        out = first(example_input)

        self.assertEqual([], out)

    def test_with_input(self):
        out = first(daily_input)

        expected = []
        print(sum(expected))
        self.assertEqual(sum(expected), sum(out))

from unittest import TestCase

from day02.first.main import first
from input.file import input

example_input = input("input.txt").rstrip("\n")
daily_input = input("../input.txt").rstrip("\n")


class TestMain(TestCase):
    def test_example(self):
        out = first(example_input)

        self.assertEqual(8, out)

    def test_with_input(self):
        out = first(daily_input)

        print(out)
        self.assertEqual(3099, out)

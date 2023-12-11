from unittest import TestCase

from day05.second.main import second
from input.file import input

example_input = input("input.txt")
daily_input = input("../input.txt")


class TestMain(TestCase):
    def test_example(self):
        out = second(example_input)

        self.assertEqual(0, out)

    def test_with_input(self):
        out = second(daily_input)

        self.assertEqual(0, out)

from unittest import TestCase

from day01.second.main import second
from input.file import input

example_input = input("input.txt")
daily_input = input("../input.txt")


class TestMain(TestCase):
    def test_example(self):
        out = second(example_input)

        self.assertEqual([29, 83, 13, 24, 42, 14, 76], out)

    def test_with_input(self):
        data = input(daily_input)
        out = second(data)

        self.assertEqual(53340, sum(out))

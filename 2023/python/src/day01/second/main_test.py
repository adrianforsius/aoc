import os
from unittest import TestCase

from input.file import test_input, input
from day01.second.main import second

path = os.path.abspath(os.path.dirname(__file__))

class TestMain(TestCase):

    def test_example(self):
        data = test_input(path)
        out = second(data)

        self.assertEqual([29, 83, 13, 24, 42, 14, 76], out)

    def test_with_input(self):
        data = input(path)
        out = second(data)

        self.assertEqual(53340, sum(out))

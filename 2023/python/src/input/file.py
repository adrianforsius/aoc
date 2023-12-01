import os

def input(path):
    with open(os.path.join(path, "input.txt")) as f:
        return f.read()

def test_input(path):
    with open(os.path.join(path, "input_test.txt")) as f:
        return f.read()

import os

def input():
    path = os.path.abspath(os.path.dirname(__file__))
    with open(os.path.join(path, "input.txt")) as f:
        return f.read()

data = input()



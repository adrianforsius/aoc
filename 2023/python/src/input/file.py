import os


def input(path):
    with open(path) as f:
        return f.read().rstrip("\n")

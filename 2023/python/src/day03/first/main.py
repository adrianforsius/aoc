from typing import Any
from day03.grid import Grid


def first(data: Any) -> Any:
    grid = Grid(data)
    numbers = []
    for i, row in enumerate(grid.grid):
        num = {"valid": False, "chars": []}
        for e, char in enumerate(row):
            try:
                int(char)
                coord = (e, i)
                num["chars"].append(char)
                if check_coord(grid.neighbours(coord)):
                    num["valid"] = True
            except ValueError:
                if num["valid"]:
                    number = int("".join([char for char in num["chars"]]))
                    numbers.append(number)
                num = {"valid": False, "chars": []}

        if num["valid"]:
            number = int("".join([char for char in num["chars"]]))
            numbers.append(number)

    print(numbers)
    return sum([int(num) for num in numbers])


def check_coord(neighboors):
    for _, val in neighboors:
        if not val:
            continue

        if val == ".":
            continue

        # adjecent numbers
        if val.isdigit():
            continue

        return True
    return False

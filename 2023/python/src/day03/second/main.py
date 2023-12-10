from typing import Any
from day03.grid import Grid


def second(data: Any) -> Any:
    grid = Grid(data)
    gears = {}
    for i, row in enumerate(grid.grid):
        num = {"valid": False, "gear": False, "chars": []}
        for e, char in enumerate(row):
            try:
                int(char)
                coord = (e, i)
                num["chars"].append(char)
                if check_coord(grid.neighbours(coord)):
                    num["valid"] = True

                gear_coord = check_gear(grid.neighbours(coord))
                if gear_coord:
                    num["gear"] = gear_coord
            except ValueError:
                if num["valid"] and num["gear"]:
                    number = int("".join([char for char in num["chars"]]))
                    gears.setdefault(num["gear"], []).append(number)
                num = {"valid": False, "gear": False, "chars": []}

        if num["valid"] and num["gear"]:
            number = int("".join([char for char in num["chars"]]))
            gears.setdefault(num["gear"], []).append(number)

    return sum([gear[0] * gear[1] for gear in gears.values() if len(gear) == 2])


def check_gear(neighboors):
    for coord, val in neighboors:
        if val == "*":
            return coord
    return False


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

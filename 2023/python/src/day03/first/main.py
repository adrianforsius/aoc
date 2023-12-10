from typing import Any

top_left = (-1, 1)
top = (0, 1)
top_right = (1, 1)
right = (1, 0)
bottom_right = (1, -1)
bottom = (0, -1)
bottom_left = (-1, -1)
left = (-1, 0)

coords = [
    top_left,
    top,
    top_right,
    right,
    bottom_right,
    bottom,
    bottom_left,
    left,
]


class Grid:
    def __init__(self, data):
        self.grid = [list(row) for row in data.split("\n")]
        self.width = len(self.grid[0])
        self.length = len(self.grid)

    def get(self, coord):
        x, y = coord
        if y >= self.length or y < 0 or x >= self.width or x < 0:
            return None
        return self.grid[y][x]

    def neighbours(self, coord):
        for c in coords:
            yield self.get((coord[0] + c[0], coord[1] + c[1]))


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
    for val in neighboors:
        if not val:
            continue

        if val == ".":
            continue

        # adjecent numbers
        if val.isdigit():
            continue

        return True
    return False

from typing import Any
from functools import reduce


def second(data: Any) -> Any:
    rows = data.split("\n")
    game_count = 0
    for game_number, row in enumerate(rows, start=1):
        colors = {
            "blue": 0,
            "green": 0,
            "red": 0,
        }
        sets = row.split(": ")[1].split("; ")
        for set in sets:
            cubes = set.split(", ")
            for cube in cubes:
                parts = cube.split(" ")
                num = int(parts[0])
                color = parts[1]
                if colors[color] < num:
                    colors[color] = num

        game_count += reduce(lambda x, y: x * y, colors.values())

    return game_count

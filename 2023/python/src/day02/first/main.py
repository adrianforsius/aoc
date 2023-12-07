from typing import Any

first_colors = {
    "blue": 14,
    "green": 13,
    "red": 12,
}


def first(data: Any) -> Any:
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

        if enough_colors(colors, first_colors):
            game_count += game_number

    return game_count


def enough_colors(colors, compare) -> bool:
    for color, value in colors.items():
        if value > compare[color]:
            return False
    return True

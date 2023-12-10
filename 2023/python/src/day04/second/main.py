from typing import Any
from itertools import zip_longest


def second(data: Any) -> Any:
    total = []
    winners = []
    rows = data.split("\n")

    carries = [0]
    for row in rows:
        instances = 1
        try:
            instances += carries.pop(0)
        except IndexError:
            pass
        game = row.split(": ")[1]
        parts = game.split(" | ")
        winning = parts[0].split()
        winning = list(set(winning))
        numbers = parts[1].split()
        numbers = list(set(numbers))

        count = len(set(winning) & set(numbers))
        carries = [
            x + y for x, y in zip_longest(carries, [instances] * count, fillvalue=0)
        ]

        total.append(instances)

    print(total)
    print(winners)
    return sum(total)

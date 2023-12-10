from typing import Any


def first(data: Any) -> Any:
    total = []
    winners = []
    rows = data.split("\n")
    for row in rows:
        game = row.split(": ")[1]
        parts = game.split(" | ")
        winning = parts[0].split()
        winning = list(set(winning))
        numbers = parts[1].split()
        numbers = list(set(numbers))
        count = len(set(winning) & set(numbers))

        score = 0
        if count > 0:
            score = 2 ** (count - 1)

        total.append(score)

    print(total)
    print(winners)
    return sum(total)

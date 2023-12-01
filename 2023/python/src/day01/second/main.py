from typing import Any

number_groups = [
    {
        "one": 1,
        "two": 2,
        "six": 6,
    },
    {
        "four": 4,
        "nine": 9,
        "five": 5,
    },
    {
        "three": 3,
        "seven": 7,
        "eight": 8,
    }
]

numbers_to_int = {
    "one": 1,
    "two": 2,
    "six": 6,

    "four": 4,
    "nine": 9,
    "five": 5,

    "three": 3,
    "seven": 7,
    "eight": 8,
}


def get_number(data):
    group = number_groups[0]
    lastIndex = min(len(data), 3)
    try:
        return group[data[:lastIndex]]
    except KeyError:
        pass

    group = number_groups[1]
    lastIndex = min(len(data), 4)
    try:
        return group[data[:lastIndex]]
    except KeyError:
        pass

    group = number_groups[2]
    lastIndex = min(len(data), 5)
    try:
        return group[data[:lastIndex]]
    except KeyError:
        pass
    return None


def second(data: Any) -> Any:
    out = []
    for line in data.splitlines():
        first = 0
        last = 0
        for i in range(0, len(line)):
            c = line[i]
            if c.isdigit():
                if first == 0:
                    first = int(c)
                last = int(c)

            number = get_number(line[i:])
            if number:
                if first == 0:
                    first = number
                last = number
        out.append(int(f"{first}{last}"))

    return out

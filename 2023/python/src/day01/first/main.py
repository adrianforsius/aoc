from typing import Any

def first(data: Any) -> Any:
    out = []
    for line in data.splitlines():
        first = 0
        last = 0
        for c in line:
            if c.isdigit():
                if first == 0:
                    first = int(c)
                last = int(c)
        out.append(int(f"{first}{last}"))

    return out

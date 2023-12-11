import sys
from typing import Any


def first(data: Any) -> Any:
    parts = data.split("\n\n")
    seeds = parts[0].split(": ")[1].strip().split(" ")
    seed_to_soil = parts[1].split("\n")[1:]
    soil_to_fer = parts[2].split("\n")[1:]
    fert_to_wat = parts[3].split("\n")[1:]
    wat_to_lig = parts[4].split("\n")[1:]
    lig_to_temp = parts[5].split("\n")[1:]
    temp_to_hum = parts[6].split("\n")[1:]
    hum_to_loc = parts[7].split("\n")[1:]

    min_seed = sys.maxsize
    for seed in seeds:
        soil = to_map(seed, seed_to_soil)
        fer = to_map(soil, soil_to_fer)
        wat = to_map(fer, fert_to_wat)
        lig = to_map(wat, wat_to_lig)
        temp = to_map(lig, lig_to_temp)
        hum = to_map(temp, temp_to_hum)
        loc = to_map(hum, hum_to_loc)
        if loc < min_seed:
            min_seed = loc

    return min_seed


def to_map(value, items):
    val = int(value)
    for item in items:
        parts = item.split(" ")
        dst = int(parts[0])
        src = int(parts[1])
        rng = int(parts[2])
        if val <= src + rng and val >= src:
            diff = dst - src + val
            if diff < 0:
                return diff * -1
            return diff

    print(val)
    return val

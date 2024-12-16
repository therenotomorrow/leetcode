import functools


@functools.cache
def _transform_stone(stone: int) -> tuple[int, ...]:
    if stone == 0:
        return (1,)

    as_str = str(stone)
    year = 2024

    if (lstr := len(as_str)) and lstr % 2 == 0:
        half = lstr // 2
        return int(as_str[:half]), int(as_str[half:])

    return (stone * year,)


@functools.cache
def _blink(stones: tuple[int, ...], blink: int) -> int:
    if blink == 0:
        return len(stones)

    if len(stones) == 1:
        stones = _transform_stone(stones[0])

        return _blink(stones, blink - 1)

    return sum(_blink((stone,), blink) for stone in stones)


def first_star(stones: list[int]) -> int:
    blinks = 25

    return _blink(tuple(stones), blinks)


def second_star(stones: list[int]) -> int:
    blinks = 75

    return _blink(tuple(stones), blinks)

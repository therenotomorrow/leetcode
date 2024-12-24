import collections
import dataclasses

_asterisk = '*'


@dataclasses.dataclass(slots=True, frozen=True)
class Point:
    row: int
    col: int


def _find_adjustment(
    engine: list[str],
    row: int,
    left: int,
    right: int,
) -> bool:
    left -= 1 if left > 0 else 0
    right += 1 if right < len(engine[0]) - 1 else 0

    if row == 0:
        top = '.' * (right - left)
    else:
        top = engine[row - 1][left:right]

    if row == len(engine) - 1:
        bot = '.' * (right - left)
    else:
        bot = engine[row + 1][left:right]

    return all(
        (
            all((char == '.' for char in top)),
            all((char == '.' for char in bot)),
            engine[row][left].isdigit() or engine[row][left] == '.',
            engine[row][right - 1].isdigit() or engine[row][right - 1] == '.',
        ),
    )


def _collect_number(text: str) -> str:
    idx = 0
    num = ''

    while idx < len(text) and text[idx].isdigit():
        num += text[idx]
        idx += 1

    return num


def _find_gear_top(
    engine: list[str],
    row: int,
    left: int,
    right: int,
) -> Point | None:
    if row == 0:
        return None

    position: Point | None = None

    for idx in range(left, right):
        char = engine[row - 1][idx]

        if char not in {'.', _asterisk}:
            return None

        if char == _asterisk:
            return Point(row=row - 1, col=idx)

    return position


def _find_gear_bot(
    engine: list[str],
    row: int,
    left: int,
    right: int,
) -> Point | None:
    if row == len(engine) - 1:
        return None

    position: Point | None = None

    for idx in range(left, right):
        char = engine[row + 1][idx]

        if char not in {'.', _asterisk}:
            return None

        if char == _asterisk:
            return Point(row=row + 1, col=idx)

    return position


def _find_gear_left(engine: list[str], row: int, left: int) -> Point | None:
    position: Point | None = None

    if engine[row][left] == _asterisk:
        return Point(row=row, col=left)

    return position


def _find_gear_right(engine: list[str], row: int, right: int) -> Point | None:
    position: Point | None = None

    right -= 1

    if engine[row][right] == _asterisk:
        return Point(row=row, col=right)

    return position


def _find_gears(
    engine: list[str],
    row: int,
    left: int,
    right: int,
) -> Point | None:
    left -= 1 if left > 0 else 0
    right += 1 if right < len(engine[0]) - 1 else 0

    return (
        _find_gear_top(engine, row, left, right)
        or _find_gear_bot(engine, row, left, right)
        or _find_gear_left(engine, row, left)
        or _find_gear_right(engine, row, right)
    )


def _calculate_ratio(numbers: list[int]) -> int:
    if len(numbers) < 2:
        return 0

    mul = 1
    for numa in numbers:
        mul *= numa

    return mul


def first_star(engine: list[str]) -> int:
    nums = []

    for row, line in enumerate(engine):
        col = 0

        while col < len(line):
            num = _collect_number(line[col:])

            if num:
                if not _find_adjustment(engine, row, col, col + len(num)):
                    nums.append(int(num))

                col += len(num) - 1

            col += 1

    return sum(nums)


def second_star(engine: list[str]) -> int:
    nums = collections.defaultdict(list)

    for row, line in enumerate(engine):
        col = 0

        while col < len(line):
            if num := _collect_number(line[col:]):
                gear = _find_gears(engine, row, col, col + len(num))

                if gear is not None:
                    nums[gear].append(int(num))

                col += len(num) - 1

            col += 1

    return sum(_calculate_ratio(numbers) for numbers in nums.values())

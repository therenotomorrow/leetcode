import dataclasses

_hashtag = '#'
_directions = {
    '^': (-1, 0),
    'v': (1, 0),
    '<': (0, -1),
    '>': (0, 1),
}


@dataclasses.dataclass(slots=True)
class Point:
    row: int
    col: int
    direct: str

    @property
    def as_tuple(self) -> tuple[int, int]:
        return self.row, self.col


def _move_forward(point: Point) -> Point:
    direction = _directions[point.direct]

    point.row += direction[0]
    point.col += direction[1]

    return point


def _move_backward(point: Point) -> Point:
    direction = _directions[point.direct]

    point.row -= direction[0]
    point.col -= direction[1]

    return point


def _change_direction(point: Point) -> Point:
    match point.direct:
        case '^':
            point.direct = '>'
        case '>':
            point.direct = 'v'
        case 'v':
            point.direct = '<'
        case '<':
            point.direct = '^'

    return point


def _is_guard_inside(grid: list[list[str]], point: Point) -> bool:
    return 0 <= point.row < len(grid) and 0 <= point.col < len(grid[0])


def _is_obstacle(grid: list[list[str]], point: Point) -> bool:
    return grid[point.row][point.col] == _hashtag


def _travel(grid: list[list[str]], start: Point) -> set[tuple[int, int]]:
    point = Point(row=start.row, col=start.col, direct=start.direct)
    visited = set()

    while True:
        point = _move_forward(point)

        if not _is_guard_inside(grid, point):
            break

        if _is_obstacle(grid, point):
            point = _move_backward(point)
            point = _change_direction(point)

            continue

        visited.add(point.as_tuple)

    return visited


def _detect_cycle(grid: list[list[str]], start: Point) -> bool:
    point = Point(row=start.row, col=start.col, direct=start.direct)
    tries = len(grid) * len(grid[0])  # empirically

    while tries > 0:
        tries -= 1

        point = _move_forward(point)

        if not _is_guard_inside(grid, point):
            break

        if not _is_obstacle(grid, point):
            continue

        point = _move_backward(point)
        point = _change_direction(point)

    return tries == 0


def first_star(grid: list[list[str]], start: Point) -> int:
    return len(_travel(grid, start))


def second_star(grid: list[list[str]], start: Point) -> int:
    ans = 0

    for row, col in _travel(grid, start):
        if start.as_tuple == (row, col):
            continue

        grid[row][col] = _hashtag
        ans += _detect_cycle(grid, start)
        grid[row][col] = '.'

    return ans

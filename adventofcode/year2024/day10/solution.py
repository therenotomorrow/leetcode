import dataclasses

_directions = {
    (-1, 0),
    (1, 0),
    (0, -1),
    (0, 1),
}


@dataclasses.dataclass(slots=True)
class Point:
    row: int
    col: int
    height: int

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented

        return self.row == other.row and self.col == other.col

    def __hash__(self) -> int:
        return hash((self.row, self.col))


def _is_inside_map(topomap: list[list[Point]], row: int, col: int) -> bool:
    return 0 <= row < len(topomap) and 0 <= col < len(topomap[0])


def _can_move(from_point: Point, to_point: Point) -> bool:
    return to_point.height - from_point.height == 1


def _find_next_moves(topomap: list[list[Point]], point: Point) -> list[Point]:
    points = []

    for direct in _directions:
        next_row = point.row + direct[0]
        next_col = point.col + direct[1]

        if not _is_inside_map(topomap, next_row, next_col):
            continue

        step = topomap[next_row][next_col]

        if _can_move(point, step):
            points.append(step)

    return points


def _travel_map(topomap: list[list[Point]], point: Point) -> set[Point]:
    if point.height == 9:
        return {point}

    ans = set()

    for next_point in _find_next_moves(topomap, point):
        ans.update(_travel_map(topomap, next_point))

    return ans


def _count_map(topomap: list[list[Point]], point: Point) -> int:
    if point.height == 9:
        return 1

    return sum(_count_map(topomap, next_point) for next_point in _find_next_moves(topomap, point))


def first_star(topomap: list[list[Point]]) -> int:
    ans = 0

    for line in topomap:
        for point in line:
            if point.height != 0:
                continue

            ans += len(_travel_map(topomap, point))

    return ans


def second_star(topomap: list[list[Point]]) -> int:
    ans = 0

    for line in topomap:
        for point in line:
            if point.height != 0:
                continue

            ans += _count_map(topomap, point)

    return ans

import dataclasses
import typing


@dataclasses.dataclass(slots=True)
class Point:
    row: int
    col: int

    def __add__(self, other: 'Point') -> 'Point':
        return Point(row=self.row + other.row, col=self.col + other.col)

    def __sub__(self, other: 'Point') -> 'Point':
        return Point(row=self.row - other.row, col=self.col - other.col)

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented

        return self.row == other.row and self.col == other.col

    def __hash__(self) -> int:
        return hash((self.row, self.col))


Grid = list[list[str]]
Freq = dict[str, list[Point]]


def _is_point_inside(grid: Grid, point: Point) -> bool:
    return 0 <= point.row < len(grid) and 0 <= point.col < len(grid[0])


def _create_diagonal(
    grid: Grid,
    antenna1: Point,
    antenna2: Point,
) -> list[Point]:
    diagonal = []

    diff = antenna1 - antenna2
    opposite = antenna1 + diff

    while _is_point_inside(grid, opposite):
        diagonal.append(opposite)
        opposite += diff

    return diagonal


def _yield_antennas(
    freq: Freq,
) -> typing.Generator[tuple[Point, Point], None, None]:
    for antennas in freq.values():
        yield from (
            (antenna1, antenna2)
            for antenna1 in antennas
            for antenna2 in antennas
            if antenna1 != antenna2
        )


def first_star(grid: Grid, freq: Freq) -> int:
    anti_nodes = set()

    for antenna1, antenna2 in _yield_antennas(freq):
        diff = antenna1 - antenna2
        opposite = antenna1 + diff

        if not _is_point_inside(grid, opposite):
            continue

        anti_nodes.add(opposite)

    return len(anti_nodes)


def second_star(grid: Grid, freq: Freq) -> int:
    anti_nodes: set[Point] = set()

    for antenna1, antenna2 in _yield_antennas(freq):
        anti_nodes.update((antenna1, antenna2))

        diagonal = _create_diagonal(grid, antenna1, antenna2)
        anti_nodes.update(diagonal)

    return len(anti_nodes)

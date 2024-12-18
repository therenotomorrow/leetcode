import dataclasses
import typing

_directions = {
    (-1, 0),
    (1, 0),
    (0, -1),
    (0, 1),
}

_checks = {
    ((-1, 0), (0, -1)),
    ((0, 1), (-1, 0)),
    ((1, 0), (0, 1)),
    ((0, -1), (1, 0)),
}


@dataclasses.dataclass(slots=True)
class Point:
    row: int
    col: int
    plant: str

    def __add__(self, other: 'Point') -> 'Point':
        return Point(
            row=self.row + other.row,
            col=self.col + other.col,
            plant='',
        )

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented

        return self.row == other.row and self.col == other.col

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def is_inside(self, garden: 'Garden') -> bool:
        return 0 <= self.row < len(garden) and 0 <= self.col < len(garden[0])


Garden = list[list[Point]]
Region = set[Point]


def _move(point: Point, direct: tuple[int, int]) -> Point:
    return point + Point(row=direct[0], col=direct[1], plant='')


def _travel(point: Point, garden: Garden) -> typing.Iterator[Point]:
    for direct in _directions:
        neighbor = _move(point, direct)

        if not neighbor.is_inside(garden):
            continue

        neighbor.plant = garden[neighbor.row][neighbor.col].plant

        yield neighbor


def _collect_region(visited: Region, garden: Garden, start: Point) -> Region:
    if start in visited:
        return set()

    visited.add(start)
    region = {start}

    for neighbor in _travel(start, garden):
        if neighbor.plant != start.plant:
            continue

        region.update(_collect_region(visited, garden, neighbor))

    return region


def _collect_regions(garden: Garden) -> typing.Iterator[Region]:
    visited: set[Point] = set()

    for row in garden:
        for point in row:
            if point in visited:
                continue

            region = _collect_region(visited, garden, point)
            if region:
                yield region


def first_star(garden: Garden) -> int:
    ans = 0

    for region in _collect_regions(garden):
        perimeter = 0

        for point in region:
            for direct in _directions:
                neighbor = _move(point, direct)

                if neighbor not in region:
                    perimeter += 1

        ans += len(region) * perimeter

    return ans


def second_star(garden: Garden) -> int:  # noqa:WPS210,WPS231
    ans = 0

    for region in _collect_regions(garden):
        num_sides = 0

        for point in region:
            for check in _checks:
                point_a = _move(point, check[0])
                point_b = _move(point, check[1])
                point_ab = _move(point_a, check[1])

                a_ok = point_a in region
                b_ok = point_b in region
                ab_ok = point_ab in region

                both_outside = not a_ok and not b_ok
                forms_corner = a_ok and b_ok and not ab_ok

                if both_outside or forms_corner:
                    num_sides += 1

        ans += len(region) * num_sides

    return ans

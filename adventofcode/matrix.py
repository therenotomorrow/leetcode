import dataclasses


@dataclasses.dataclass
class Point:
    row: int
    col: int
    cell: str = ''

    @property
    def tuple(self) -> tuple[int, int]:
        return self.row, self.col

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented

        return self.tuple == other.tuple

    def __hash__(self) -> int:
        return hash(self.tuple)


class Matrix:
    def __init__(self, data: list[list[Point]]) -> None:
        self.rows = len(data)
        self.cols = len(data[0])
        self.data = data

    def inside(self, point: Point) -> bool:
        return 0 <= point.row < self.rows and 0 <= point.col < self.cols


def printer(matrix: Matrix) -> str:
    lines = []
    print()
    for row in matrix.data:
        lines.append(''.join(point.cell for point in row))

    print('\n'.join(lines))

    return ''

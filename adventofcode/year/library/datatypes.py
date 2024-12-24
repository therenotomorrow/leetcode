import dataclasses
import typing

Ints: typing.TypeAlias = list[int]
PointTuple: typing.TypeAlias = tuple[int, int]


@dataclasses.dataclass(slots=True, kw_only=True)
class Point:
    row: int
    col: int
    mark: str
    grid: 'Grid'

    def __eq__(self, other: object) -> bool:
        if not isinstance(other, Point):
            return NotImplemented

        return (self.row, self.col) == (other.row, other.col)

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def move(self, direct: 'Direction') -> tuple['Point', bool]:
        next_row = self.row + direct.row
        next_col = self.col + direct.col

        if 0 <= next_row < self.grid.rows and 0 <= next_col < self.grid.cols:
            return self.grid[next_row][next_col], True

        return self, False

    def back(self, direct: 'Direction') -> tuple['Point', bool]:
        return self.move(direct.opposite)

    @property
    def clone(self) -> 'Point':
        return Point(row=self.row, col=self.col, mark=self.mark, grid=self.grid)


_up = '^'
_down = 'v'
_left = '<'
_right = '>'


class Direction:  # noqa: WPS214
    def __init__(self, *, row: int = 0, col: int = 0, name: str = ''):
        self.name = name
        self.row = row
        self.col = col

        if self.name == _up:
            self.row = -1
            self.col = 0
        elif self.name == _down:
            self.row = 1
            self.col = 0
        elif self.name == _left:
            self.row = 0
            self.col = -1
        elif self.name == _right:
            self.row = 0
            self.col = 1

    def increase(self, mul: int) -> 'Direction':
        mul_row = mul * self.row
        mul_col = mul * self.col

        direct = Direction(row=mul_row, col=mul_col)
        direct.name = self.name

        return direct

    @property
    def opposite(self) -> 'Direction':
        if self.name == _up:
            return self.down()
        if self.name == _down:
            return self.up()
        if self.name == _left:
            return self.right()
        if self.name == _right:
            return self.left()

        raise NotImplementedError()

    @property
    def turn_right(self) -> 'Direction':
        if self.name == _up:
            return self.right()
        if self.name == _right:
            return self.down()
        if self.name == _down:
            return self.left()
        if self.name == _left:
            return self.up()

        raise NotImplementedError()

    @classmethod
    def up(cls) -> 'Direction':
        return cls(name=_up)

    @classmethod
    def down(cls) -> 'Direction':
        return cls(name=_down)

    @classmethod
    def left(cls) -> 'Direction':
        return cls(name=_left)

    @classmethod
    def right(cls) -> 'Direction':
        return cls(name=_right)

    @classmethod
    def up_left(cls) -> 'Direction':
        return cls(row=-1, col=-1)

    @classmethod
    def down_left(cls) -> 'Direction':
        return cls(row=1, col=-1)

    @classmethod
    def up_right(cls) -> 'Direction':
        return cls(row=-1, col=1)

    @classmethod
    def down_right(cls) -> 'Direction':
        return cls(row=1, col=1)


GridRaw: typing.TypeAlias = list[list[str]]


class Grid:
    def __init__(self, grid: GridRaw):
        self.points = []

        for row, line in enumerate(grid):
            points = []

            for col, cell in enumerate(line):
                points.append(Point(row=row, col=col, mark=cell, grid=self))

            self.points.append(points)

        self.rows = len(self.points)
        self.cols = len(self.points[0])
        self.size = self.rows * self.cols

    def __iter__(self) -> typing.Iterator[list[Point]]:
        return iter(self.points)

    def __getitem__(self, idx: int) -> list[Point]:
        return self.points[idx]

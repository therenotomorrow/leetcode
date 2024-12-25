import dataclasses
import typing

Ints: typing.TypeAlias = list[int]
Strs: typing.TypeAlias = list[str]


@dataclasses.dataclass(slots=True, kw_only=True)
class Point:
    row: int
    col: int
    mark: str
    grid: 'Grid'

    def __eq__(self, other: object) -> bool:
        if isinstance(other, str):
            return self.mark == other

        if not isinstance(other, Point):
            return NotImplemented

        return (self.row, self.col) == (other.row, other.col)

    def __hash__(self) -> int:
        return hash((self.row, self.col))

    def __str__(self) -> str:
        return self.mark

    @property
    def inside(self) -> bool:
        return 0 <= self.row < self.grid.rows and 0 <= self.col < self.grid.cols

    def move(self, direct: 'Direction') -> tuple['Point', bool]:
        next_point = Point(row=self.row + direct.row, col=self.col + direct.col, mark='', grid=self.grid)

        if next_point.inside:
            return self.grid[next_point], True

        return self, False


_up = '^'
_down = 'v'
_left = '<'
_right = '>'


@dataclasses.dataclass(slots=True, kw_only=True, frozen=True)
class Direction:  # noqa:WPS214
    row: int = 0
    col: int = 0
    name: str = ''

    def increase(self, mul: int) -> 'Direction':
        mul_row = mul * self.row
        mul_col = mul * self.col

        return Direction(row=mul_row, col=mul_col, name=self.name)

    @classmethod
    def new(cls, name: str) -> 'Direction':
        if name == _up:
            return cls.up()
        if name == _down:
            return cls.down()
        if name == _left:
            return cls.left()
        if name == _right:
            return cls.right()

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
        return cls(row=-1, name=_up)

    @classmethod
    def down(cls) -> 'Direction':
        return cls(row=1, name=_down)

    @classmethod
    def left(cls) -> 'Direction':
        return cls(col=-1, name=_left)

    @classmethod
    def right(cls) -> 'Direction':
        return cls(col=1, name=_right)

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


class Grid:
    def __init__(self, grid: list[Strs]):
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

    def __getitem__(self, point: Point) -> Point:
        return self.points[point.row][point.col]

    def point(self, row: int, col: int, mark: str = '') -> Point:
        return Point(row=row, col=col, mark=mark, grid=self)

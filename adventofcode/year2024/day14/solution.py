import collections
import dataclasses
import pprint


@dataclasses.dataclass(slots=True)
class Position:
    row: int
    col: int


@dataclasses.dataclass(slots=True)
class Velocity:
    row: int
    col: int


@dataclasses.dataclass(slots=True)
class Robot:
    pos: Position
    vel: Velocity


def _move(tiles: list[list[int]], robot: Robot, seconds: int) -> list[list[int]]:
    rows = len(tiles)
    cols = len(tiles[0])

    new_row = robot.pos.row + seconds * robot.vel.row
    new_col = robot.pos.col + seconds * robot.vel.col

    new_row = new_row % rows
    new_col = new_col % cols

    tiles[new_row][new_col] += 1

    return tiles


def _calc_safety_factor(tiles: list[list[int]]) -> int:
    rows = len(tiles) // 2
    cols = len(tiles[0]) // 2

    robots = [0, 0, 0, 0]

    for row in range(rows):
        for col in range(cols):
            robots[0] += tiles[row][col]

    for row in range(rows + 1, len(tiles)):
        for col in range(cols):
            robots[1] += tiles[row][col]

    for row in range(rows):
        for col in range(cols + 1, len(tiles[0])):
            robots[2] += tiles[row][col]

    for row in range(rows + 1, len(tiles)):
        for col in range(cols + 1, len(tiles[0])):
            robots[3] += tiles[row][col]

    mul = 1
    for num in robots:
        mul *= num

    return mul


def first_star(robots: list[Robot], rows: int, cols: int) -> int:
    seconds = 100
    tiles = [[0] * cols for _ in range(rows)]

    for robot in robots:
        tiles = _move(tiles, robot, seconds)

    return _calc_safety_factor(tiles)


def second_star(robots: list[Robot], rows: int, cols: int) -> int:
    seconds = 10000
    ans = 0
    while seconds > 0:
        tiles = [[0] * cols for _ in range(rows)]

        for robot in robots:
            tiles = _move(tiles, robot, seconds)

        for row in tiles:
            for col, val in enumerate(row):
                if val == 0:
                    row[col] = '.'
                else:
                    row[col] = '*'

            gg = collections.Counter(row).get('*', 0)
            ans = max(gg, ans)

        if ans == 33:
            for row in tiles:
                print(''.join(row))
            return seconds

        seconds -= 1

    print(ans)
    return 0

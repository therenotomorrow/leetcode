import collections

from adventofcode import matrix


class Point(matrix.Point):
    @property
    def is_box(self) -> bool:
        return self.cell == 'O'

    @property
    def is_wall(self) -> bool:
        return self.cell == '#'

    @property
    def is_free(self) -> bool:
        return self.cell == '.'


class Map(matrix.Matrix):
    def move(self, point: Point, direct: str) -> Point:
        next_row = point.row
        next_col = point.col

        match direct:
            case '^':
                next_row -= 1
            case 'v':
                next_row += 1
            case '<':
                next_col -= 1
            case '>':
                next_col += 1

        return self.data[next_row][next_col]

    def move_boxes(self, point: Point, direct: str) -> bool:
        prev_free = point
        next_free = point

        while next_free.is_box:
            next_free = self.move(next_free, direct)

            if next_free.is_wall:
                return False

        if next_free.is_wall:
            return False

        prev_free.cell, next_free.cell = next_free.cell, prev_free.cell

        return True

    def travel(self, point: Point, directions: list[str]) -> None:
        for direct in directions:
            next_point = self.move(point, direct)

            if next_point.is_wall:
                continue

            if next_point.is_free:
                point.cell, next_point.cell = next_point.cell, point.cell
                point = next_point
                continue

            if next_point.is_box:
                if self.move_boxes(next_point, direct):
                    point.cell, next_point.cell = next_point.cell, point.cell
                    point = next_point

    def calc_gps(self) -> int:
        gps = 0

        for row, line in enumerate(self.data):
            for col, point in enumerate(line):
                if point.cell == 'O':
                    gps += 100 * row + col

        return gps


def first_star(topomap: Map, start: Point, moves: list[str]) -> int:
    print('--')
    topomap.travel(start, moves)
    return topomap.calc_gps()


def second_star(grid: list[list[str]], directions: list[str]) -> int:
    m, n = len(grid), len(grid[0])
    for i in range(m):
        for j in reversed(range(n)):
            if grid[i][j] == '#':
                grid[i].insert(j, '#')
            if grid[i][j] == '.':
                grid[i].insert(j, '.')
            if grid[i][j] == '@':
                robot = (i, j * 2)
                grid[i][j : j + 1] = ['.', '.']
            if grid[i][j] == 'O':
                grid[i][j : j + 1] = ['[', ']']

    for d in directions[:]:
        i, j = robot

        if d == '<':
            k = j - 1
            while grid[i][k] == ']':
                k -= 2
            if grid[i][k] == '.':
                for l in range(k, j):
                    grid[i][l] = grid[i][l + 1]
                robot = (i, j - 1)

        elif d == '>':
            k = j + 1
            while grid[i][k] == '[':
                k += 2
            if grid[i][k] == '.':
                for l in reversed(range(j + 1, k + 1)):
                    grid[i][l] = grid[i][l - 1]
                robot = (i, j + 1)

        elif d == '^':
            queue = {(i - 1, j)}
            rows = collections.defaultdict(set)
            while queue:
                x, y = queue.pop()
                match grid[x][y]:
                    case '#':
                        break
                    case ']':
                        rows[x] |= {y - 1, y}
                        queue |= {(x - 1, y), (x - 1, y - 1)}
                    case '[':
                        rows[x] |= {y, y + 1}
                        queue |= {(x - 1, y), (x - 1, y + 1)}
                    case '.':
                        rows[x].add(y)
            else:
                for x in sorted(rows):
                    for y in rows[x]:
                        grid[x][y] = grid[x + 1][y] if y in rows[x + 1] else '.'
                robot = (i - 1, j)

        elif d == 'v':
            queue = {(i + 1, j)}
            rows = collections.defaultdict(set)
            while queue:
                x, y = queue.pop()
                match grid[x][y]:
                    case '#':
                        break
                    case ']':
                        rows[x] |= {y - 1, y}
                        queue |= {(x + 1, y), (x + 1, y - 1)}
                    case '[':
                        rows[x] |= {y, y + 1}
                        queue |= {(x + 1, y), (x + 1, y + 1)}
                    case '.':
                        rows[x].add(y)
            else:
                for x in sorted(rows, reverse=True):
                    for y in rows[x]:
                        grid[x][y] = grid[x - 1][y] if y in rows[x - 1] else '.'
                robot = (i + 1, j)

    total = 0
    for i in range(m):
        for j in range(n * 2):
            if grid[i][j] == '[':
                total += 100 * i + j

    return total

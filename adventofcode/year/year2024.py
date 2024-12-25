import collections
import dataclasses
import functools
import re
import typing

from adventofcode.year.library import datatypes, days, funcs


@dataclasses.dataclass
class Day01(days.Day):
    def parse(self) -> tuple[datatypes.Ints, datatypes.Ints]:
        left_list = []
        right_list = []

        for line in funcs.split(self.indata):
            left, right = map(int, line.split())

            left_list.append(left)
            right_list.append(right)

        return left_list, right_list

    def first_star(self) -> int:
        left_list, right_list = self.parse()

        left_list.sort()
        right_list.sort()

        return sum(abs(right - left) for right, left in zip(right_list, left_list))

    def second_star(self) -> int:
        left_list, right_list = self.parse()

        count = collections.Counter(right_list)

        return sum(num * count.get(num, 0) for num in left_list)


@dataclasses.dataclass
class Day02(days.Day):
    def parse(self) -> list[datatypes.Ints]:
        reports = []

        for line in funcs.split(self.indata):
            reports.append(list(map(int, line.split())))

        return reports

    def first_star(self) -> int:
        return sum(self._is_safe(report) for report in self.parse())

    def second_star(self) -> int:
        safe = 0

        for report in self.parse():
            for exclude, _ in enumerate(report):
                left, right = exclude, exclude + 1

                if self._is_safe(report[:left] + report[right:]):
                    safe += 1

                    break

        return safe

    @classmethod
    def _is_safe(cls, report: datatypes.Ints) -> bool:
        is_asc = (report[1] - report[0]) > 0
        diffs = []

        for idx, curr in enumerate(report[1:], start=1):
            diff = curr - report[idx - 1]

            if not is_asc:
                diff = -diff

            diffs.append(1 <= diff <= 3)

        return all(diffs)


@dataclasses.dataclass
class Day03(days.Day):
    mul_pattern = re.compile(r'mul\((\d+),(\d+)\)')
    do_pattern = re.compile(r'(do\(\))|(don\'t\(\))')

    def parse(self) -> str:
        return self.indata.strip()

    def first_star(self) -> int:
        return sum(int(num1) * int(num2) for num1, num2 in self.mul_pattern.findall(self.parse()))

    def second_star(self) -> int:
        mul = self._find_mul(self.parse())
        doe = self._find_do(self.parse())

        ans = 0
        allow = True

        for idx, _ in enumerate(self.parse()):
            if doe.get(idx) is not None:
                allow = doe[idx]
                continue

            if mul.get(idx) is not None and allow:
                ans += mul[idx]

        return ans

    def _find_mul(self, memory: str) -> dict[int, int]:
        statements = {}

        for group in self.mul_pattern.finditer(memory):
            num1, num2 = map(int, group.groups())
            statements[group.start(0)] = num1 * num2

        return statements

    def _find_do(self, memory: str) -> dict[int, bool]:
        statements = {}

        for group in self.do_pattern.finditer(memory):
            statements[group.start(0)] = bool(group.group(1))

        return statements


@dataclasses.dataclass
class Day04(days.Day):
    xmas_pattern = 'XMAS'
    mas_pattern = {'MAS', 'SAM'}
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
        datatypes.Direction.up_left(),
        datatypes.Direction.up_right(),
        datatypes.Direction.down_left(),
        datatypes.Direction.down_right(),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> int:
        puzzle = self.parse()
        count = 0

        for line in puzzle:
            for point in line:
                if point != 'X':
                    continue

                count += self._find_xmas_pattern(point)

        return count

    def second_star(self) -> int:
        puzzle = self.parse()

        return sum(self._find_mas_pattern(point) for line in puzzle for point in line)

    def _find_xmas_pattern(self, start: datatypes.Point) -> int:
        count = 0

        for direct in self.directions:
            for idx, _ in enumerate(self.xmas_pattern):
                point = start.next(direct.increase(idx))

                if point.outside or point != self.xmas_pattern[idx]:
                    break
            else:
                count += 1

        return count

    def _find_mas_pattern_main_diag(self, point: datatypes.Point) -> bool:
        up_left = point.next(datatypes.Direction.up_left())
        down_right = point.next(datatypes.Direction.down_right())

        line = ''.join(map(str, (up_left, point, down_right)))

        return up_left.inside and down_right.inside and line in self.mas_pattern

    def _find_mas_pattern_sub_diag(self, point: datatypes.Point) -> bool:
        up_right = point.next(datatypes.Direction.up_right())
        down_left = point.next(datatypes.Direction.down_left())

        line = ''.join(map(str, (up_right, point, down_left)))

        return up_right.inside and down_left.inside and line in self.mas_pattern

    def _find_mas_pattern(self, start: datatypes.Point) -> bool:
        return self._find_mas_pattern_main_diag(start) and self._find_mas_pattern_sub_diag(start)


Update: typing.TypeAlias = datatypes.Ints
Updates: typing.TypeAlias = list[Update]
Ruleset: typing.TypeAlias = dict[int, set[int]]


@dataclasses.dataclass
class Day05(days.Day):
    def parse(self) -> tuple[Updates, Ruleset]:
        parts = funcs.parts(self.indata)

        ruleset = self._parse_ruleset(parts[0])
        updates = self._parse_updates(parts[1])

        return updates, ruleset

    def first_star(self) -> int:
        updates, ruleset = self.parse()
        total = 0

        for update in updates:
            if self._is_correct_update(update, ruleset):
                total += update[len(update) // 2]

        return total

    def second_star(self) -> int:
        updates, ruleset = self.parse()
        total = 0

        for update in updates:
            if self._is_correct_update(update, ruleset):
                continue

            update = self._sort_update(update, ruleset)
            total += update[len(update) // 2]

        return total

    @classmethod
    def _parse_ruleset(cls, indata: str) -> Ruleset:
        ruleset = collections.defaultdict(set)

        for line in funcs.split(indata):
            before, after = tuple(map(int, line.split('|')))
            ruleset[before].add(after)

        return ruleset

    @classmethod
    def _parse_updates(cls, indata: str) -> Updates:
        updates = []

        for line in funcs.split(indata):
            updates.append([int(num) for num in line.split(',')])

        return updates

    @classmethod
    def _is_correct_update(cls, update: Update, rules: Ruleset) -> bool:
        used = set()

        for num in update:
            for related_page in rules.get(num, set()):
                if related_page in used:
                    return False

            used.add(num)

        return True

    @classmethod
    def _sort_update(cls, update: Update, rules: Ruleset) -> Update:
        return sorted(
            update,
            key=lambda rule: sum(rule in rules.get(num, set()) for num in update),
            reverse=True,
        )


@dataclasses.dataclass
class Day06(days.Day):
    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> int:
        grid = self.parse()
        start = self._find_start(grid)
        direction = datatypes.Direction.new(start.mark)

        return len(self._travel(start, direction))

    def second_star(self) -> int:
        grid = self.parse()
        start = self._find_start(grid)
        direction = datatypes.Direction.new(start.mark)

        ans = 0

        for point in self._travel(start, direction):
            if start == point:
                continue

            grid[point] = '#'
            ans += self._detect_cycle(start, direction, grid.size)
            grid[point] = '.'

        return ans

    @classmethod
    def _detect_cycle(cls, point: datatypes.Point, direct: datatypes.Direction, tries: int) -> bool:
        while tries > 0:
            tries -= 1

            next_point = point.next(direct)

            if next_point.outside:
                break

            if next_point == '#':
                direct = direct.turn_right
            else:
                point = next_point

        return tries == 0

    @classmethod
    def _find_start(cls, grid: datatypes.Grid) -> datatypes.Point:
        start = datatypes.Point(row=0, col=0, mark='', grid=grid)

        for row in grid:
            for point in row:
                if point.mark in {'<', '>', '^', 'v'}:
                    start = point

                    break

        return start

    @classmethod
    def _travel(cls, point: datatypes.Point, direct: datatypes.Direction) -> datatypes.UniqPoints:
        visited = set()

        while True:
            next_point = point.next(direct)

            if next_point.outside:
                break

            if next_point == '#':
                direct = direct.turn_right
                continue

            point = next_point
            visited.add(point)

        return visited


Combination = tuple[int, datatypes.Ints]


@dataclasses.dataclass
class Day07(days.Day):
    def parse(self) -> list[Combination]:
        calibration = []

        for line in funcs.split(self.indata):
            part1, part2 = line.split(': ')
            numbers = list(map(int, part2.split()))
            calibration.append((int(part1), numbers))

        return calibration

    def first_star(self) -> int:
        calibration = self.parse()
        total = 0

        for ans, numbers in calibration:
            if self._is_equal_first(ans, numbers[0], numbers):
                total += ans

        return total

    def second_star(self) -> int:
        calibration = self.parse()
        total = 0

        for ans, numbers in calibration:
            if self._is_equal_second(ans, numbers[0], numbers):
                total += ans

        return total

    @classmethod
    def _is_equal_first(cls, ans: int, curr: int, numbers: datatypes.Ints) -> bool:
        if len(numbers) == 1:
            return ans == curr

        other = numbers[1]

        is_sum = cls._is_equal_first(ans, curr + other, numbers[1:])
        is_mul = cls._is_equal_first(ans, curr * other, numbers[1:])

        return is_sum or is_mul

    @classmethod
    def _is_equal_second(cls, ans: int, curr: int, numbers: datatypes.Ints) -> bool:
        if len(numbers) == 1:
            return ans == curr

        other = numbers[1]
        another = int(str(curr) + str(other))

        is_con = cls._is_equal_second(ans, another, numbers[1:])
        is_sum = cls._is_equal_second(ans, curr + other, numbers[1:])
        is_mul = cls._is_equal_second(ans, curr * other, numbers[1:])

        return is_con or is_sum or is_mul


Frequencies = dict[str, list[datatypes.Point]]


@dataclasses.dataclass
class Day08(days.Day):
    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> int:
        grid = self.parse()
        anti_nodes = set()

        for antenna1, antenna2 in self._yield_antennas(grid):
            point = grid.point(row=2 * antenna1.row - antenna2.row, col=2 * antenna1.col - antenna2.col)

            if not point.inside:
                continue

            anti_nodes.add(point)

        return len(anti_nodes)

    def second_star(self) -> int:
        grid = self.parse()
        anti_nodes: datatypes.UniqPoints = set()

        for antenna1, antenna2 in self._yield_antennas(grid):
            anti_nodes.update((antenna1, antenna2))
            diagonal = self._create_diagonal(grid, antenna1, antenna2)
            anti_nodes.update(diagonal)

        return len(anti_nodes)

    @classmethod
    def _collect_frequencies(cls, grid: datatypes.Grid) -> Frequencies:
        frequencies = collections.defaultdict(list)

        for row in grid:
            for point in row:
                if point == '.':
                    continue

                frequencies[point.mark].append(point)

        return frequencies

    @classmethod
    def _yield_antennas(cls, grid: datatypes.Grid) -> typing.Iterable[tuple[datatypes.Point, datatypes.Point]]:
        frequencies = cls._collect_frequencies(grid)

        for antennas in frequencies.values():
            yield from ((antenna1, antenna2) for antenna1 in antennas for antenna2 in antennas if antenna1 != antenna2)

    @classmethod
    def _create_diagonal(
        cls,
        grid: datatypes.Grid,
        antenna1: datatypes.Point,
        antenna2: datatypes.Point,
    ) -> list[datatypes.Point]:
        diagonal = []

        diff_row = antenna1.row - antenna2.row
        diff_col = antenna1.col - antenna2.col

        point = grid.point(row=antenna1.row + diff_row, col=antenna1.col + diff_col)

        while point.inside:
            diagonal.append(point)

            point = grid.point(row=point.row + diff_row, col=point.col + diff_col)

        return diagonal


IdxSize: typing.TypeAlias = tuple[int, int]


@dataclasses.dataclass
class Day09(days.Day):
    def parse(self) -> datatypes.Strs:
        fragments = []
        ident = 0

        for idx, char in enumerate(map(int, self.indata.strip())):
            if idx % 2 == 0:
                fill = str(ident)
                ident += 1
            else:
                fill = '.'

            fragments.extend([fill for _ in range(char)])

        return fragments

    def first_star(self) -> int:
        fragments = self.parse()
        left = 0
        right = len(fragments) - 1

        while left < right:
            if fragments[left] != '.':
                left += 1
                continue

            if fragments[right] == '.':
                right -= 1
                continue

            fragments[left], fragments[right] = fragments[right], fragments[left]  # noqa:WPS414
            left += 1
            right -= 1

        return self._calc_checksum(fragments)

    def second_star(self) -> int:  # noqa:WPS231
        fragments = self.parse()

        left = 0
        right = len(fragments) - 1

        while right >= 0:
            if fragments[left] != '.':
                left += 1
                continue

            if fragments[right] == '.':
                right -= 1
                continue

            right, want_size = self._find_desired_size(fragments, right)

            if left >= right:
                left = 0
                continue

            left, allowed_size = self._find_allowed_size(fragments, left)

            if want_size > allowed_size:
                right += want_size
                continue

            fragments = self._swap_fragments(
                fragments,
                left,
                allowed_size,
                right,
                want_size,
            )
            left = 0

        return self._calc_checksum(fragments)

    @classmethod
    def _calc_checksum(cls, fragments: datatypes.Strs) -> int:
        ans = 0

        for ident, num in enumerate(fragments):
            if num == '.':
                continue

            ans += ident * int(num)

        return ans

    @classmethod
    def _find_desired_size(cls, fragments: datatypes.Strs, idx: int) -> IdxSize:
        want = 0
        fragment = fragments[idx]

        while fragments[idx] == fragment:
            idx -= 1
            want += 1

        return idx, want

    @classmethod
    def _find_allowed_size(cls, fragments: datatypes.Strs, idx: int) -> IdxSize:
        allowed = 0

        while fragments[idx] == '.':
            idx += 1
            allowed += 1

        return idx, allowed

    @classmethod
    def _swap_fragments(  # noqa:WPS211
        cls,
        fragments: datatypes.Strs,
        start: int,
        allowed: int,
        end: int,
        want: int,
    ) -> datatypes.Strs:
        right1 = end + 1
        right2 = right1 + want

        left1 = start - allowed
        left2 = left1 + want

        fragments[left1:left2], fragments[right1:right2] = (  # noqa:WPS362,WPS414
            fragments[right1:right2],
            fragments[left1:left2],
        )

        return fragments


@dataclasses.dataclass
class Day10(days.Day):
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> int:
        grid = self.parse()
        ans = 0

        for line in grid:
            for point in line:
                if point != '0':
                    continue

                ans += len(self._travel_map(point))

        return ans

    def second_star(self) -> int:
        grid = self.parse()
        ans = 0

        for line in grid:
            for point in line:
                if point != '0':
                    continue

                ans += self._count_map(point)

        return ans

    @classmethod
    def _can_move(cls, from_point: datatypes.Point, to_point: datatypes.Point) -> bool:
        return int(to_point.mark) - int(from_point.mark) == 1

    def _find_next_moves(self, point: datatypes.Point) -> list[datatypes.Point]:
        points = []

        for direct in self.directions:
            next_point = point.next(direct)

            if next_point.outside:
                continue

            if self._can_move(point, next_point):
                points.append(next_point)

        return points

    def _travel_map(self, point: datatypes.Point) -> datatypes.UniqPoints:
        if point == '9':
            return {point}

        ans = set()

        for next_point in self._find_next_moves(point):
            ans.update(self._travel_map(next_point))

        return ans

    def _count_map(self, point: datatypes.Point) -> int:
        if point == '9':
            return 1

        return sum(self._count_map(next_point) for next_point in self._find_next_moves(point))


@dataclasses.dataclass
class Day11(days.Day):
    def parse(self) -> datatypes.Ints:
        stones: datatypes.Ints = []

        for line in funcs.split(self.indata):
            stones.extend(map(int, line.split()))

        return stones

    def first_star(self) -> int:
        stones = self.parse()
        blinks = 25

        return self._blink(tuple(stones), blinks)

    def second_star(self) -> int:
        stones = self.parse()
        blinks = 75

        return self._blink(tuple(stones), blinks)

    @classmethod
    @functools.cache
    def _transform_stone(cls, stone: int) -> datatypes.IntsTuple:
        if stone == 0:
            return (1,)

        as_str = str(stone)
        lstr = len(as_str)
        year = 2024

        if lstr % 2 == 0:
            half = lstr // 2
            return int(as_str[:half]), int(as_str[half:])

        return (stone * year,)

    @classmethod
    @functools.cache
    def _blink(cls, stones: datatypes.IntsTuple, blink: int) -> int:
        if blink == 0:
            return len(stones)

        if len(stones) == 1:
            stones = cls._transform_stone(stones[0])

            return cls._blink(stones, blink - 1)

        return sum(cls._blink((stone,), blink) for stone in stones)


@dataclasses.dataclass
class Day12(days.Day):
    directions = {
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
    }
    _checks = {
        (datatypes.Direction.up(), datatypes.Direction.left()),
        (datatypes.Direction.right(), datatypes.Direction.up()),
        (datatypes.Direction.down(), datatypes.Direction.right()),
        (datatypes.Direction.left(), datatypes.Direction.down()),
    }

    def parse(self) -> datatypes.Grid:
        raw = []

        for line in funcs.split(self.indata):
            raw.append(list(line))

        return datatypes.Grid(raw)

    def first_star(self) -> int:
        ans = 0

        for region in self._collect_regions(self.parse()):
            perimeter = 0

            for point in region:
                for direct in self.directions:
                    if point.next(direct) not in region:
                        perimeter += 1

            ans += len(region) * perimeter

        return ans

    def second_star(self) -> int:  # noqa:WPS210,WPS231
        grid = self.parse()
        ans = 0

        for region in self._collect_regions(grid):
            num_sides = 0

            for point in region:
                for direct in self._checks:
                    point_a = point.next(direct[0])
                    point_b = point.next(direct[1])
                    point_ab = point_a.next(direct[1])

                    a_ok = point_a in region
                    b_ok = point_b in region
                    ab_ok = point_ab in region

                    both_outside = not a_ok and not b_ok
                    forms_corner = a_ok and b_ok and not ab_ok

                    if both_outside or forms_corner:
                        num_sides += 1

            ans += len(region) * num_sides

        return ans

    def _travel(self, point: datatypes.Point) -> typing.Iterator[datatypes.Point]:
        for direct in self.directions:
            next_point = point.next(direct)

            if next_point.outside:
                continue

            yield next_point

    def _collect_region(self, visited: datatypes.UniqPoints, start: datatypes.Point) -> datatypes.UniqPoints:
        if start in visited:
            return set()

        visited.add(start)
        region = {start}

        for next_point in self._travel(start):
            if next_point.mark != start.mark:
                continue

            region.update(self._collect_region(visited, next_point))

        return region

    def _collect_regions(self, grid: datatypes.Grid) -> typing.Iterator[datatypes.UniqPoints]:
        visited: datatypes.UniqPoints = set()

        for row in grid:
            for point in row:
                if point in visited:
                    continue

                region = self._collect_region(visited, point)
                if region:
                    yield region


@dataclasses.dataclass(slots=True)
class Button:
    cost_x: int
    cost_y: int


@dataclasses.dataclass(slots=True)
class Prize:
    value_x: int
    value_y: int


NumPair: typing.TypeAlias = tuple[int, int]
Game = tuple[Button, Button, Prize]


@dataclasses.dataclass
class Day13(days.Day):

    def parse(self) -> list[Game]:  # noqa:WPS210
        button_a = Button(cost_x=0, cost_y=0)
        button_b = Button(cost_x=0, cost_y=0)
        prize = Prize(value_x=0, value_y=0)

        games = []
        times = 1

        for line in funcs.split(self.indata):
            match times:
                case 1:
                    button_a = Button(*self._parse_nums(line.split('+')))
                case 2:
                    button_b = Button(*self._parse_nums(line.split('+')))
                case 3:
                    prize = Prize(*self._parse_nums(line.split('=')))

            times += 1
            if times > 3:
                games.append((button_a, button_b, prize))
                times = 1

        return games

    def first_star(self) -> int:
        games = self.parse()
        shift = 0

        return sum(self._resolve(*game, shift=shift) for game in games)

    def second_star(self) -> int:
        games = self.parse()
        shift = 10000000000000

        return sum(self._resolve(*game, shift=shift) for game in games)

    @classmethod
    def _parse_nums(cls, parts: datatypes.Strs) -> NumPair:
        left = int(parts[-2].split(',')[0])
        right = int(parts[-1])

        return left, right

    @classmethod
    def _calc(cls, left: Button, right: Button, prize: Prize) -> tuple[int, bool]:
        divider = right.cost_x * left.cost_y - left.cost_x * right.cost_y
        dividend = prize.value_x * left.cost_y - prize.value_y * left.cost_x

        if divider == 0:
            return 0, False

        ans = dividend / divider

        return int(ans), ans == int(ans)

    @classmethod
    def _resolve(cls, button_a: Button, button_b: Button, prize: Prize, *, shift: int) -> int:
        prize.value_x += shift
        prize.value_y += shift

        times_a, ok_a = cls._calc(button_b, button_a, prize)
        times_b, ok_b = cls._calc(button_a, button_b, prize)

        return 3 * times_a + times_b if ok_a and ok_b else 0


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


Tiles: typing.TypeAlias = list[datatypes.Ints]


@dataclasses.dataclass
class Day14(days.Day):
    rows: int
    cols: int

    def parse(self) -> list[Robot]:  # noqa:WPS210
        robots = []

        for line in funcs.split(self.indata):
            parts: datatypes.Strs = line.split()

            col, row = map(int, parts[0].lstrip('p=').split(','))
            v_col, v_row = map(int, parts[1].lstrip('v=').split(','))

            robots.append(Robot(pos=Position(row=row, col=col), vel=Velocity(row=v_row, col=v_col)))

        return robots

    def first_star(self) -> int:
        robots = self.parse()
        seconds = 100

        return self._run(robots, seconds)

    def second_star(self) -> int:
        robots = self.parse()
        seconds = 10000
        tree_fact = 69552000  # empirically

        while seconds > 0:
            fact = self._run(robots, seconds)
            if fact == tree_fact:
                return seconds

            seconds -= 1

        return seconds

    def _run(self, robots: list[Robot], seconds: int) -> int:
        tiles = [[0 for _ in range(self.cols)] for _ in range(self.rows)]

        for robot in robots:
            tiles = self._move(tiles, robot, seconds)

        return self._calc_safety_factor(tiles)

    def _move(self, tiles: Tiles, robot: Robot, seconds: int) -> Tiles:
        new_row = robot.pos.row + seconds * robot.vel.row
        new_col = robot.pos.col + seconds * robot.vel.col

        new_row = new_row % self.rows
        new_col = new_col % self.cols

        tiles[new_row][new_col] += 1

        return tiles

    def _calc_first_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2) for row in range(self.rows // 2))

    def _calc_second_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2) for row in range(self.rows // 2 + 1, len(tiles)))

    def _calc_third_squad(self, tiles: Tiles) -> int:
        return sum(tiles[row][col] for col in range(self.cols // 2 + 1, len(tiles[0])) for row in range(self.rows // 2))

    def _calc_fourth_squad(self, tiles: Tiles) -> int:
        return sum(
            tiles[row][col]
            for col in range(self.cols // 2 + 1, len(tiles[0]))
            for row in range(self.rows // 2 + 1, len(tiles))
        )

    def _calc_safety_factor(self, tiles: Tiles) -> int:
        robots = {
            self._calc_first_squad(tiles),
            self._calc_second_squad(tiles),
            self._calc_third_squad(tiles),
            self._calc_fourth_squad(tiles),
        }

        mul = 1
        for num in robots:
            mul *= num

        return mul

import collections
import dataclasses
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
    directions = [
        datatypes.Direction.up(),
        datatypes.Direction.down(),
        datatypes.Direction.left(),
        datatypes.Direction.right(),
        datatypes.Direction.up_left(),
        datatypes.Direction.up_right(),
        datatypes.Direction.down_left(),
        datatypes.Direction.down_right(),
    ]

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
                point, is_ok = start.move(direct.increase(idx))

                if not is_ok or point != self.xmas_pattern[idx]:
                    break
            else:
                count += 1

        return count

    def _find_mas_pattern_main_diag(self, point: datatypes.Point) -> bool:
        up_left, ok_left = point.move(datatypes.Direction.up_left())
        down_right, ok_right = point.move(datatypes.Direction.down_right())

        line = ''.join(map(str, (up_left, point, down_right)))

        return ok_left and ok_right and line in self.mas_pattern

    def _find_mas_pattern_sub_diag(self, point: datatypes.Point) -> bool:
        up_right, ok_right = point.move(datatypes.Direction.up_right())
        down_left, ok_left = point.move(datatypes.Direction.down_left())

        line = ''.join(map(str, (up_right, point, down_left)))

        return ok_right and ok_left and line in self.mas_pattern

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

            point.mark = '#'
            ans += self._detect_cycle(start, direction, grid.size)
            point.mark = '.'

        return ans

    @classmethod
    def _detect_cycle(cls, point: datatypes.Point, direct: datatypes.Direction, tries: int) -> bool:
        while tries > 0:
            tries -= 1

            next_point, is_ok = point.move(direct)

            if not is_ok:
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
    def _travel(cls, point: datatypes.Point, direct: datatypes.Direction) -> set[datatypes.Point]:
        visited = set()

        while True:
            next_point, is_ok = point.move(direct)

            if not is_ok:
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
        anti_nodes: set[datatypes.Point] = set()

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
    ) -> list[str]:
        right1 = end + 1
        right2 = right1 + want

        left1 = start - allowed
        left2 = left1 + want

        fragments[left1:left2], fragments[right1:right2] = (  # noqa:WPS362,WPS414
            fragments[right1:right2],
            fragments[left1:left2],
        )

        return fragments

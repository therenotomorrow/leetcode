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

        return sum(abs(right - left) for right, left in zip(right_list, left_list))  # noqa: WPS221

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
        return sum(int(num1) * int(num2) for num1, num2 in self.mul_pattern.findall(self.parse()))  # noqa: WPS221

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
                if point.mark != 'X':
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

                if not is_ok or point.mark != self.xmas_pattern[idx]:
                    break
            else:
                count += 1

        return count

    def _find_mas_pattern_main_diag(self, point: datatypes.Point) -> bool:
        up_left, ok_left = point.move(datatypes.Direction.up_left())
        down_right, ok_right = point.move(datatypes.Direction.down_right())

        line = ''.join((up_left.mark, point.mark, down_right.mark))

        return ok_left and ok_right and line in self.mas_pattern

    def _find_mas_pattern_sub_diag(self, point: datatypes.Point) -> bool:
        up_right, ok_right = point.move(datatypes.Direction.up_right())
        down_left, ok_left = point.move(datatypes.Direction.down_left())

        line = ''.join((up_right.mark, point.mark, down_left.mark))

        return ok_right and ok_left and line in self.mas_pattern

    def _find_mas_pattern(self, start: datatypes.Point) -> bool:
        return self._find_mas_pattern_main_diag(start) and self._find_mas_pattern_sub_diag(start)


Update: typing.TypeAlias = list[int]
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
            key=lambda rule: sum(rule in rules.get(num, set()) for num in update),  # noqa: WPS221
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
        direction = datatypes.Direction(name=start.mark)

        return len(self._travel(start, direction))

    def second_star(self) -> int:
        grid = self.parse()
        start = self._find_start(grid)
        direction = datatypes.Direction(name=start.mark)

        ans = 0

        for point in self._travel(start, direction):
            if start == point:
                continue

            point.mark = '#'
            ans += self._detect_cycle(start, direction, grid.size)
            point.mark = '.'

        return ans

    @classmethod
    def _detect_cycle(cls, start: datatypes.Point, direct: datatypes.Direction, tries: int) -> bool:
        point = start.clone

        while tries > 0:
            tries -= 1

            point, is_ok = point.move(direct)

            if not is_ok:
                break

            if point.mark == '#':
                point, _ = point.back(direct)
                direct = direct.turn_right

        return tries == 0

    @classmethod
    def _find_start(cls, grid: datatypes.Grid) -> datatypes.Point:
        start = datatypes.Point(row=0, col=0, mark='', grid=grid)

        for line in grid:
            for point in line:
                if point.mark in {'<', '>', '^', 'v'}:
                    start = point

                    break

        return start

    @classmethod
    def _travel(cls, start: datatypes.Point, direct: datatypes.Direction) -> set[datatypes.Point]:
        point = start.clone
        visited = set()

        while True:
            point, is_ok = point.move(direct)

            if not is_ok:
                break

            if point.mark == '#':
                point, _ = point.back(direct)
                direct = direct.turn_right

                continue

            visited.add(point)

        return visited

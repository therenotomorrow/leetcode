import dataclasses
import typing

from adventofcode.library import datatypes, funcs


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
class Day(datatypes.Day):
    def parse(self) -> list[Game]:
        button_a = Button(cost_x=0, cost_y=0)
        button_b = Button(cost_x=0, cost_y=0)
        prize = Prize(value_x=0, value_y=0)

        games = []
        times = 1

        for line in funcs.split(self.indata):
            match times:
                case 1:
                    button_a = Button(*_parse_nums(line.split('+')))
                case 2:
                    button_b = Button(*_parse_nums(line.split('+')))
                case 3:
                    prize = Prize(*_parse_nums(line.split('=')))

            times += 1
            if times > 3:
                games.append((button_a, button_b, prize))
                times = 1

        return games

    def first_star(self) -> int:
        games = self.parse()
        shift = 0

        return sum(_resolve(*game, shift=shift) for game in games)

    def second_star(self) -> int:
        games = self.parse()
        shift = 10000000000000

        return sum(_resolve(*game, shift=shift) for game in games)


def _resolve(button_a: Button, button_b: Button, prize: Prize, *, shift: int) -> int:
    prize.value_x += shift
    prize.value_y += shift

    times_a, ok_a = _calc(button_b, button_a, prize)
    times_b, ok_b = _calc(button_a, button_b, prize)

    return 3 * times_a + times_b if ok_a and ok_b else 0


def _parse_nums(parts: datatypes.Strs) -> NumPair:
    left = int(parts[-2].split(',')[0])
    right = int(parts[-1])

    return left, right


def _calc(left: Button, right: Button, prize: Prize) -> tuple[int, bool]:
    divider = right.cost_x * left.cost_y - left.cost_x * right.cost_y
    dividend = prize.value_x * left.cost_y - prize.value_y * left.cost_x

    if divider == 0:
        return 0, False

    ans = dividend / divider

    return int(ans), ans == int(ans)

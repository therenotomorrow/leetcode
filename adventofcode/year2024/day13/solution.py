import dataclasses


@dataclasses.dataclass(slots=True)
class Button:
    cost_x: int
    cost_y: int


@dataclasses.dataclass(slots=True)
class Prize:
    value_x: int
    value_y: int


Game = tuple[Button, Button, Prize]


def _calc(left: Button, right: Button, prize: Prize) -> tuple[int, bool]:
    divider = right.cost_x * left.cost_y
    divider -= left.cost_x * right.cost_y

    if divider == 0:
        return 0, False

    dividend = prize.value_x * left.cost_y
    dividend -= prize.value_y * left.cost_x

    ans = dividend / divider

    return int(ans), ans == int(ans)


def _resolve(
    button_a: Button,
    button_b: Button,
    prize: Prize,
    *,
    shift: int,
) -> int:
    prize.value_x += shift
    prize.value_y += shift

    times_a, ok_a = _calc(button_b, button_a, prize)
    times_b, ok_b = _calc(button_a, button_b, prize)

    return 3 * times_a + times_b if ok_a and ok_b else 0


def first_star(games: list[Game]) -> int:
    shift = 0

    return sum(_resolve(*game, shift=shift) for game in games)


def second_star(games: list[Game]) -> int:
    shift = 10000000000000

    return sum(_resolve(*game, shift=shift) for game in games)

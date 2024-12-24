import collections

Rounds = list[dict[str, int]]


def _check_valid_rounds(rounds: Rounds) -> bool:
    restrict = {'red': 12, 'green': 13, 'blue': 14}
    valid = True

    for turn in rounds:
        valid_turn = all(scr <= restrict[clr] for clr, scr in turn.items())
        valid = valid and valid_turn

    return valid


def _find_cubes_power(rounds: Rounds) -> int:
    total: dict[str, int] = collections.defaultdict(int)

    for turn in rounds:
        for color, score in turn.items():
            total[color] = max(total[color], score)

    mul = 1

    for cubes in total.values():
        mul *= cubes

    return mul


def first_star(games: dict[int, Rounds]) -> int:
    score = 0

    for game_id, rounds in games.items():
        if _check_valid_rounds(rounds):
            score += game_id

    return score


def second_star(games: dict[int, Rounds]) -> int:
    return sum(_find_cubes_power(rounds) for rounds in games.values())

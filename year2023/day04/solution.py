import functools

Cards = list[tuple[set[int], set[int]]]


def first_star(cards: Cards) -> int:
    points = 0

    for win, own in cards:
        score = 1

        for card in win:
            if card in own:
                score *= 2

        points += score // 2

    return points


def second_star(cards: Cards) -> int:
    @functools.cache
    def play(card: int) -> int:  # noqa:WPS430
        if card >= len(cards):
            return 0

        win, own = cards[card]
        matches = len(win & own)

        card += 1

        return 1 + sum(play(card) for card in range(card, card + matches))

    return sum(play(card) for card in range(len(cards)))

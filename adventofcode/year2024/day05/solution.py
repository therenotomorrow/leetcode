import collections

Rules = list[tuple[int, int]]
Updates = list[list[int]]
Ruleset = dict[int, set[int]]


def _is_correct_update(update: list[int], rules: Ruleset) -> bool:
    used = set()

    for num in update:
        for related_page in rules.get(num, set()):
            if related_page in used:
                return False

        used.add(num)

    return True


def _sort_update(update: list[int], rules: Ruleset) -> list[int]:
    return sorted(
        update,
        key=lambda rule: sum(rule in rules.get(num, set()) for num in update),
        reverse=True,
    )


def first_star(rules: Rules, updates: Updates) -> int:
    ruleset = collections.defaultdict(set)

    for before, after in rules:
        ruleset[before].add(after)

    total = 0

    for update in updates:
        if _is_correct_update(update, ruleset):
            total += update[len(update) // 2]

    return total


def second_star(rules: Rules, updates: Updates) -> int:
    ruleset = collections.defaultdict(set)

    for before, after in rules:
        ruleset[before].add(after)

    total = 0

    for update in updates:
        if _is_correct_update(update, ruleset):
            continue

        update = _sort_update(update, ruleset)
        total += update[len(update) // 2]

    return total

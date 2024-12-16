def _make_fragments(disk: str) -> list[str]:
    fragments = []
    ident = 0

    for idx, char in enumerate(map(int, disk)):
        if idx % 2 == 0:
            fill = str(ident)
            ident += 1
        else:
            fill = '.'

        fragments.extend([fill for _ in range(char)])

    return fragments


def _calc_checksum(fragments: list[str]) -> int:
    ans = 0

    for ident, num in enumerate(fragments):
        if num == '.':
            continue

        ans += ident * int(num)

    return ans


def _swap_fragments(
    fragments: list[str],
    start: int,
    allowed: int,
    end: int,
    want: int,
) -> list[str]:
    right1 = end + 1
    right2 = right1 + want

    left1 = start - allowed
    left2 = left1 + want
    fragments[left1:left2], fragments[right1:right2] = (  # noqa:WPS362
        fragments[right1:right2],
        fragments[left1:left2],
    )

    return fragments


def _find_desired_size(fragments: list[str], idx: int) -> tuple[int, int]:
    want = 0
    fragment = fragments[idx]

    while fragments[idx] == fragment:
        idx -= 1
        want += 1

    return idx, want


def _find_allowed_size(fragments: list[str], idx: int) -> tuple[int, int]:
    allowed = 0

    while fragments[idx] == '.':
        idx += 1
        allowed += 1

    return idx, allowed


def first_star(disk: str) -> int:
    fragments = _make_fragments(disk)

    left = 0
    right = len(fragments) - 1

    while left < right:
        if fragments[left] != '.':
            left += 1
            continue

        if fragments[right] == '.':
            right -= 1
            continue

        fragments[left], fragments[right] = fragments[right], fragments[left]
        left += 1
        right -= 1

    return _calc_checksum(fragments)


def second_star(disk: str) -> int:  # noqa:WPS231
    fragments = _make_fragments(disk)

    left = 0
    right = len(fragments) - 1

    while right >= 0:
        if fragments[left] != '.':
            left += 1
            continue

        if fragments[right] == '.':
            right -= 1
            continue

        right, want_size = _find_desired_size(fragments, right)

        if left >= right:
            left = 0
            continue

        left, allowed_size = _find_allowed_size(fragments, left)

        if want_size > allowed_size:
            right += want_size
            continue

        fragments = _swap_fragments(
            fragments,
            left,
            allowed_size,
            right,
            want_size,
        )
        left = 0

    return _calc_checksum(fragments)

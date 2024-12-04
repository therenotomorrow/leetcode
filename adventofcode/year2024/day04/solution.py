_m_letter = 'M'
_a_letter = 'A'
_s_letter = 'S'


def _find_xmas_pattern(
    puzzle: list[list[str]],
    row: int,
    col: int,
    direction: tuple[int, int],
) -> bool:
    word = 'XMAS'

    for idx, _ in enumerate(word):
        new_row = row + direction[0] * idx
        new_col = col + direction[1] * idx

        if not (0 <= new_row < len(puzzle) and 0 <= new_col < len(puzzle[0])):
            return False

        if puzzle[new_row][new_col] != word[idx]:
            return False

    return True


def _find_mas_pattern_main_diag(
    puzzle: list[list[str]],
    row: int,
    col: int,
) -> bool:
    line = ''.join(
        (
            puzzle[row - 1][col - 1],
            puzzle[row][col],
            puzzle[row + 1][col + 1],
        ),
    )

    return line in {'MAS', 'SAM'}


def _find_mas_pattern_sub_diag(
    puzzle: list[list[str]],
    row: int,
    col: int,
) -> bool:
    line = ''.join(
        (
            puzzle[row - 1][col + 1],
            puzzle[row][col],
            puzzle[row + 1][col - 1],
        ),
    )

    return line in {'MAS', 'SAM'}


def _find_mas_pattern(puzzle: list[list[str]], row: int, col: int) -> bool:
    if row - 1 < 0 or row + 1 >= len(puzzle):
        return False

    if col - 1 < 0 or col + 1 >= len(puzzle[0]):
        return False

    return _find_mas_pattern_main_diag(
        puzzle,
        row,
        col,
    ) and _find_mas_pattern_sub_diag(
        puzzle,
        row,
        col,
    )


def first_star(puzzle: list[list[str]]) -> int:  # noqa:WPS231
    count = 0
    directions = [
        (0, 1),
        (0, -1),
        (1, 0),
        (-1, 0),
        (1, 1),
        (-1, -1),
        (1, -1),
        (-1, 1),
    ]

    for row, line in enumerate(puzzle):
        for col, _ in enumerate(line):
            if puzzle[row][col] != 'X':
                continue

            for direction in directions:
                if _find_xmas_pattern(puzzle, row, col, direction):
                    count += 1

    return count


def second_star(puzzle: list[list[str]]) -> int:
    return sum(
        _find_mas_pattern(puzzle, row, col)
        for row, line in enumerate(puzzle)
        for col, _ in enumerate(line)
    )

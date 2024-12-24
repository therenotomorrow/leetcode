def _found_digit_by_value(text: str, idx: int) -> int:
    return int(digit) if (digit := text[idx]).isdigit() else 0


def _found_digit_by_name(text: str, idx: int, *, asc: bool) -> int:
    digits = {
        'one': 1,
        'two': 2,
        'three': 3,
        'four': 4,
        'five': 5,
        'six': 6,
        'seven': 7,
        'eight': 8,
        'nine': 9,
    }

    for digit_name, digit_value in digits.items():
        if asc and text.startswith(digit_name, idx):
            return digit_value

        if not asc and text.endswith(digit_name, 0, idx + 1):
            return digit_value

    return 0


def _find_left(text: str, *, skipper: bool) -> int:
    num = 0

    for left, _ in enumerate(text):
        num += _found_digit_by_value(text, left)

        if not skipper:
            num += _found_digit_by_name(text, left, asc=True)

        if num:
            break

    return num


def _find_right(text: str, *, skipper: bool) -> int:
    num = 0

    for right in range(len(text) - 1, -1, -1):
        num += _found_digit_by_value(text, right)

        if not skipper:
            num += _found_digit_by_name(text, right, asc=False)

        if num:
            break

    return num


def first_star(document: list[str]) -> int:
    return sum(10 * _find_left(line, skipper=True) + _find_right(line, skipper=True) for line in document)


def second_star(document: list[str]) -> int:
    return sum(10 * _find_left(line, skipper=False) + _find_right(line, skipper=False) for line in document)

Combination = tuple[int, list[int]]


def _is_equal_math(ans: int, curr: int, numbers: list[int]) -> bool:
    if len(numbers) == 1:
        return ans == curr

    other = numbers[1]

    is_sum = _is_equal_math(ans, curr + other, numbers[1:])
    is_mul = _is_equal_math(ans, curr * other, numbers[1:])

    return is_sum or is_mul


def _is_equal_concat(ans: int, curr: int, numbers: list[int]) -> bool:
    if len(numbers) == 1:
        return ans == curr

    other = numbers[1]

    is_con = _is_equal_concat(ans, int(str(curr) + str(other)), numbers[1:])
    is_sum = _is_equal_concat(ans, curr + other, numbers[1:])
    is_mul = _is_equal_concat(ans, curr * other, numbers[1:])

    return is_con or is_sum or is_mul


def first_star(calibration: list[Combination]) -> int:
    total = 0

    for ans, numbers in calibration:
        if _is_equal_math(ans, numbers[0], numbers):
            total += ans

    return total


def second_star(calibration: list[Combination]) -> int:
    total = 0

    for ans, numbers in calibration:
        if _is_equal_concat(ans, numbers[0], numbers):
            total += ans

    return total

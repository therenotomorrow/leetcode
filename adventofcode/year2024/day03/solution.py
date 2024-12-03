import re


def _find_mul_statements(memory: str) -> dict[int, int]:
    pattern = re.compile(r'mul\((\d+),(\d+)\)')
    statements = {}

    for group in pattern.finditer(memory):
        num1, num2 = map(int, group.groups())
        statements[group.start(0)] = num1 * num2

    return statements


def _find_do_statements(memory: str) -> dict[int, bool]:
    pattern = re.compile(r'(do\(\))|(don\'t\(\))')
    statements = {}

    for group in pattern.finditer(memory):
        statements[group.start(0)] = bool(group.group(1))

    return statements


def first_star(memory: str) -> int:
    pattern = re.compile(r'mul\((\d+),(\d+)\)')

    return sum(int(num1) * int(num2) for num1, num2 in pattern.findall(memory))


def second_star(memory: str) -> int:
    mul = _find_mul_statements(memory)
    doe = _find_do_statements(memory)

    ans = 0
    allow = True

    for idx, _ in enumerate(memory):
        if doe.get(idx) is not None:
            allow = doe[idx]
            continue

        if mul.get(idx) is not None and allow:
            ans += mul[idx]

    return ans

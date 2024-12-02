def check_order(report: list[int]) -> bool:
    idx = 1
    asc = (report[1] - report[0]) > 0
    diffs: list[bool] = []

    while idx < len(report):
        diff = report[idx] - report[idx - 1]
        idx += 1

        diffs.append(1 <= diff <= 3 if asc else -3 <= diff <= -1)

    return all(diffs)


def first_star(reports: list[list[int]]) -> int:
    return sum(check_order(report) for report in reports)


def second_star(reports: list[list[int]]) -> int:
    safe = 0

    for report in reports:
        for exclude, _ in enumerate(report):
            left, right = exclude, exclude + 1

            if check_order(report[:left] + report[right:]):
                safe += 1

                break

    return safe

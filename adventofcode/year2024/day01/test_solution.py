import pathlib

import pytest

from .solution import first_star, second_star

_dir = pathlib.Path(__file__).parent


def parse(filename: str) -> tuple[list[int], list[int]]:
    list1 = []
    list2 = []

    with open(filename) as input_file:
        for line in input_file:
            left, right = map(int, line.strip().split())

            list1.append(left)
            list2.append(right)

    return list1, list2


class TestDay01:
    @pytest.mark.parametrize(
        'filename, want',
        [
            ('example.txt', 11),
            ('input.txt', 2375403),
        ],
        ids=('example', 'result'),
    )
    def test_first_star(self, filename: str, want: int) -> None:
        list1, list2 = parse(f'{_dir}/{filename}')

        got = first_star(list1, list2)

        assert got == want, f'first_star() = {got}, want = {want}'

    @pytest.mark.parametrize(
        'filename, want',
        [
            ('example.txt', 31),
            ('input.txt', 23082277),
        ],
        ids=('example', 'result'),
    )
    def test_second_star(self, filename: str, want: int) -> None:
        list1, list2 = parse(f'{_dir}/{filename}')

        got = second_star(list1, list2)

        assert got == want, f'second_star() = {got}, want = {want}'

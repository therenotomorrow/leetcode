import io

import pytest

from .solution import first_star, second_star

_example = """
125 17
"""
_input = """
8069 87014 98 809367 525 0 9494914 5
"""


def parse(indata: str) -> list[int]:
    infile = io.StringIO(indata)

    stones: list[int] = []

    for line in filter(bool, map(str.strip, infile)):
        stones.extend(map(int, line.split()))

    return stones


class TestDay11:
    @pytest.mark.parametrize(
        'indata, want',
        [(_example, 55312), (_input, 183484)],
        ids=('example', 'result'),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        stones = parse(indata)

        got = first_star(stones)

        assert got == want, f'first_star() = {got}, want = {want}'

    @pytest.mark.parametrize(
        'indata, want',
        [(_example, 65601038650482), (_input, 218817038947400)],
        ids=('example', 'result'),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        stones = parse(indata)

        got = second_star(stones)

        assert got == want, f'second_star() = {got}, want = {want}'

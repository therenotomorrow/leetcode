import pytest

from adventofcode.library import testit
from adventofcode.year2024 import day21

example = """
029A
980A
179A
456A
379A
"""
answer = """
671A
826A
670A
085A
283A
"""


class TestDay:
    @pytest.mark.parametrize(testit.names, [(example, 126384), (answer, 182844)], ids=testit.ids())
    def test_first_star(self, indata: str, want: int) -> None:
        testit.first_star(day21.Day(indata), want)

    @pytest.mark.parametrize(testit.names, [(example, 154115708116294), (answer, 226179529377982)], ids=testit.ids())
    def test_second_star(self, indata: str, want: int) -> None:
        testit.second_star(day21.Day(indata), want)

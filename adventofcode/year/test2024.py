import pytest

from adventofcode.year import data2024, year2024

_test_params = 'indata, want'


def _ids() -> tuple[str, str]:
    return 'example', 'result'


def _assert(got: int, want: int, *, star: int) -> None:
    msg = 'first_star' if star == 1 else 'second_star'

    assert got == want, f'{msg}() = {got}, want = {want}'


class TestDay01:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day01_example, 11), (data2024.day01_input, 2375403)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day01(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day01_example, 31), (data2024.day01_input, 23082277)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day01(indata).second_star()

        _assert(got, want, star=2)


class TestDay02:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day02_example, 2), (data2024.day02_input, 379)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day02(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day02_example, 4), (data2024.day02_input, 430)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day02(indata).second_star()

        _assert(got, want, star=2)


class TestDay03:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day03_example, 161), (data2024.day03_input, 168539636)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day03(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day03_example, 48), (data2024.day03_input, 97529391)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day03(indata).second_star()

        _assert(got, want, star=2)


class TestDay04:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day04_example, 18), (data2024.day04_input, 2618)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day04(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day04_example, 12), (data2024.day04_input, 2011)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day04(indata).second_star()

        _assert(got, want, star=2)


class TestDay05:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day05_example, 143), (data2024.day05_input, 7307)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day05(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day05_example, 123), (data2024.day05_input, 4713)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day05(indata).second_star()

        _assert(got, want, star=2)


class TestDay06:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day06_example, 41), (data2024.day06_input, 4883)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day06(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day06_example, 6), (data2024.day06_input, 1655)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day06(indata).second_star()

        _assert(got, want, star=2)


class TestDay07:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day07_example, 3749), (data2024.day07_input, 3312271365652)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day07(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day07_example, 11387), (data2024.day07_input, 509463489296712)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day07(indata).second_star()

        _assert(got, want, star=2)


class TestDay08:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day08_example, 14), (data2024.day08_input, 313)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day08(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day08_example, 34), (data2024.day08_input, 1064)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day08(indata).second_star()

        _assert(got, want, star=2)


class TestDay09:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day09_example, 1928), (data2024.day09_input, 6401092019345)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day09(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day09_example, 2858), (data2024.day09_input, 6431472344710)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day09(indata).second_star()

        _assert(got, want, star=2)


class TestDay10:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day10_example, 36), (data2024.day10_input, 820)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day10(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day10_example, 81), (data2024.day10_input, 1786)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day10(indata).second_star()

        _assert(got, want, star=2)


class TestDay11:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day11_example, 55312), (data2024.day11_input, 183484)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day11(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day11_example, 65601038650482), (data2024.day11_input, 218817038947400)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day11(indata).second_star()

        _assert(got, want, star=2)


class TestDay12:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day12_example, 1930), (data2024.day12_input, 1396298)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day12(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day12_example, 1206), (data2024.day12_input, 853588)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day12(indata).second_star()

        _assert(got, want, star=2)


class TestDay13:
    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day13_example, 480), (data2024.day13_input, 36838)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int) -> None:
        got = year2024.Day13(indata).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        _test_params,
        [(data2024.day13_example, 875318608908), (data2024.day13_input, 83029436920891)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int) -> None:
        got = year2024.Day13(indata).second_star()

        _assert(got, want, star=2)


class TestDay14:
    @pytest.mark.parametrize(
        f'{_test_params}, rows, cols',
        [(data2024.day14_example, 12, 7, 11), (data2024.day14_input, 214400550, 103, 101)],
        ids=_ids(),
    )
    def test_first_star(self, indata: str, want: int, rows: int, cols: int) -> None:
        got = year2024.Day14(indata, rows, cols).first_star()

        _assert(got, want, star=1)

    @pytest.mark.parametrize(
        f'{_test_params}, rows, cols',
        [(data2024.day14_example, 0, 7, 11), (data2024.day14_input, 8149, 103, 101)],
        ids=_ids(),
    )
    def test_second_star(self, indata: str, want: int, rows: int, cols: int) -> None:
        got = year2024.Day14(indata, rows, cols).second_star()

        _assert(got, want, star=2)

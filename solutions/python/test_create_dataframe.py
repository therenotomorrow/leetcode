import pandas
import pytest

from .create_dataframe import create_dataframe


class TestCreateDataframe:
    @pytest.mark.parametrize(
        'args, want',
        [
            (
                [[1, 15], [2, 11], [3, 11], [4, 20]],
                pandas.DataFrame.from_records(
                    data=[(1, 15), (2, 11), (3, 11), (4, 20)],
                    columns=('student_id', 'age'),
                ),
            ),
        ],
        ids=('smoke 1',),
    )
    def tests(self, args: list[list[int]], want: pandas.DataFrame) -> None:
        got = create_dataframe(args)
        diff = got.compare(want, keep_equal=True, align_axis=0)

        assert diff.empty, f'create_dataframe() = {got}, want = {want}'

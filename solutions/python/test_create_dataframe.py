import pandas
import pytest

from .create_dataframe import create_dataframe


class TestCreateDataframe(object):
    @pytest.mark.parametrize(
        'args, want',
        [
            (
                [[1, 15], [2, 11], [3, 11], [4, 20]],
                pandas.DataFrame.from_dict(
                    {'student_id': [1, 2, 3, 4], 'age': [15, 11, 11, 20]},
                ),
            ),
        ],
        ids=('smoke 1',),
    )
    def tests(self, args: list[list[int]], want: pandas.DataFrame) -> None:
        got = create_dataframe(args)
        diff = got.compare(want, keep_equal=True, align_axis=0)

        assert diff.empty, f'create_dataframe() = {got}, want = {want}'

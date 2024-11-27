import pandas
import pytest

from .modify_salary_column import modify_salary_column


class TestModifySalaryColumn:
    @pytest.mark.parametrize(
        'frame, want',
        [
            (
                pandas.DataFrame.from_records(
                    data=[
                        ('Jack', 19666),
                        ('Piper', 74754),
                        ('Mia', 62509),
                        ('Ulysses', 54866),
                    ],
                    columns=('name', 'salary'),
                ),
                pandas.DataFrame.from_records(
                    data=[
                        ('Jack', 39332),
                        ('Piper', 149508),
                        ('Mia', 125018),
                        ('Ulysses', 109732),
                    ],
                    columns=('name', 'salary'),
                ),
            ),
        ],
        ids=('smoke 1',),
    )
    def tests(self, frame: pandas.DataFrame, want: pandas.DataFrame) -> None:
        got = modify_salary_column(frame)
        diff = got.compare(want, keep_equal=True, align_axis=0)

        assert diff.empty, f'modify_salary_column() = {got}, want = {want}'

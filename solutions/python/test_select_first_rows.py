import pandas
import pytest

from .select_first_rows import select_first_rows


class TestSelectFirstRows:
    @pytest.mark.parametrize(
        'frame, want',
        [
            (
                pandas.DataFrame.from_records(
                    data=[
                        (3, 'Bob', 'Operations', 48675),
                        (90, 'Alice', 'Sales', 11096),
                        (9, 'Tatiana', 'Engineering', 33805),
                        (60, 'Annabelle', 'InformationTechnology', 37678),
                        (49, 'Jonathan', 'HumanResources', 23793),
                        (43, 'Khaled', 'Administration', 40454),
                    ],
                    columns=('employee_id', 'name', 'department', 'salary'),
                ),
                pandas.DataFrame.from_records(
                    data=[
                        (3, 'Bob', 'Operations', 48675),
                        (90, 'Alice', 'Sales', 11096),
                        (9, 'Tatiana', 'Engineering', 33805),
                    ],
                    columns=('employee_id', 'name', 'department', 'salary'),
                ),
            ),
        ],
        ids=('smoke 1',),
    )
    def tests(self, frame: pandas.DataFrame, want: pandas.DataFrame) -> None:
        got = select_first_rows(frame)
        diff = got.compare(want, keep_equal=True, align_axis=0)

        assert diff.empty, f'select_first_rows() = {got}, want = {want}'

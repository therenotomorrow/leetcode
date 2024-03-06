import pandas


def create_dataframe(student_data: list[list[int]]) -> pandas.DataFrame:
    return pandas.DataFrame(student_data, columns=('student_id', 'age'))

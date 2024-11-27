import pandas


def select_first_rows(employees: pandas.DataFrame) -> pandas.DataFrame:
    return employees.head(3)

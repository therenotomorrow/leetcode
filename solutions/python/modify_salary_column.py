import pandas


def modify_salary_column(employees: pandas.DataFrame) -> pandas.DataFrame:
    return employees.mul({'name': 1, 'salary': 2})

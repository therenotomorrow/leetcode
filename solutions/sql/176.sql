select (select distinct e.salary as "SecondHighestSalary"
        from employee as e
        order by e.salary desc
        offset 1 limit 1);

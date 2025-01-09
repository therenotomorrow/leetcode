select e.employee_id
from employees as e
where e.salary < 30000
  and e.manager_id not in (select e.employee_id from employees as e)
order by e.employee_id;

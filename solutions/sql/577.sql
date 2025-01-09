select e.name, b.bonus
from employee as e
         left join public.bonus as b on e.empid = b.empid
where b.bonus is null
   or b.bonus < 1000;

select c.class
from courses as c
group by c.class
having count(c.student) >= 5;

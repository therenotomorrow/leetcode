select p.email
from person as p
group by p.email
having count(p.email) > 1;
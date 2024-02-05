select p.firstname, p.lastname, a.city, a.state
from person as p
         left join address as a on p.personid = a.personid
order by p.firstname;

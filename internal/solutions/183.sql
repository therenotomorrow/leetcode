select c.name as Customers
from customers as c
         left join orders as o on c.id = o.customerid
where o.id is null;

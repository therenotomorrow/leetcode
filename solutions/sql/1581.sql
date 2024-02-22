select v.customer_id, count(v.customer_id) as count_no_trans
from visits as v
         left outer join transactions as t on v.visit_id = t.visit_id
where transaction_id is null
group by v.customer_id
order by count_no_trans desc, v.customer_id;

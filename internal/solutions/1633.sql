select r.contest_id,
       round(numeric_div(count(r.contest_id), (select count(u.*) from users as u)) * 100, 2) as percentage
from users as u
         join register as r on u.user_id = r.user_id
group by r.contest_id
order by percentage desc, r.contest_id;
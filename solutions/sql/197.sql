select w1.id
from weather as w1,
     weather as w2
where w1.temperature > w2.temperature
  and Date(w1.recorddate) - interval '1 day' = w2.recorddate;
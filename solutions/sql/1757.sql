select p.product_id
from products as p
where low_fats = 'Y'
  and recyclable = 'Y';

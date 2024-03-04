Create type low_fats as ENUM ('Y', 'N');
Create type recyclable as ENUM ('Y', 'N');

Create table If Not Exists Products
(
    product_id int,
    low_fats   low_fats,
    recyclable recyclable
);

Truncate table Products;

insert into Products (product_id, low_fats, recyclable)
values ('0', 'Y', 'N');
insert into Products (product_id, low_fats, recyclable)
values ('1', 'Y', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('2', 'N', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('3', 'Y', 'Y');
insert into Products (product_id, low_fats, recyclable)
values ('4', 'N', 'N');

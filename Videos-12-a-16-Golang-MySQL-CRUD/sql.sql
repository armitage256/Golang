create database test;

use test;

create table product(
id int auto_increment primary key,
description varchar(255),
quantity int,
price decimal(10,2),
amount decimal(10,2),
created_at timestamp default current_timestamp()   
);

insert into product (description, quantity, price, amount) values
('Orange', 50, 2.50, (50 * 2.50)), 
('Banana', 10, 1.99, (10 * 1.99)),
('Morango', 200, 5.45, (200 * 5.45));
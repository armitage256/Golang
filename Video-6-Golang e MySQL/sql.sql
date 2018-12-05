create database test;

use test;

create table if not exists users(
id int auto_increment primary key,
name varchar(50) not null,
email varchar(50) not null,
password varchar(100) not null
);
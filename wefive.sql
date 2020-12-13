CREATE DATABASE `wefive_government` CHARACTER SET 'utf8';

create table user
(
user_id int not null primary key auto_increment,
name varchar(20) not null,
password varchar(50) not null,
card_id varchar(20) not null unique,
phone varchar(20) unique
);

create table admin
(
admin_id int not null primary key auto_increment,
name varchar(20) not null,
password varchar(50) not null,
phone varchar(20) unique
);


create table governor
(
gover_id int not null primary key auto_increment,
dept_id int not null references department(dept_id),
password varchar(50) not null,
phone varchar(20) not null unique
);

create table business
(
bus_id int not null primary key auto_increment,
dept_id int not null references department(dept_id),
bus_name varchar(50),
description varchar(500),
requirement varchar(500),
cost real
);

create table department
(
dept_id int not null primary key auto_increment,
dept_name varchar(20) not null unique,
location varchar(50),
work_time varchar(100),
description varchar(200)
);
create table user_business
(
user_id int not null references user(user_id),
bus_id int not null references business(bus_id),
primary key (user_id, bus_id)
);

create table material
(
material_id int not null primary key auto_increment,
bus_id int not null references business(bus_id),
material_name varchar(20) not null,
description varchar(500),
photo_url varchar(100)
);

create table search
(
search_id int not null primary key auto_increment,
info varchar(50),
user_id int not null references user(user_id),
created_time timestamp DEFAULT CURRENT_TIMESTAMP
);

create table comment
(
comment_id int not null primary key auto_increment,
user_id int not null references user(user_id),
dept_id int not null references department(dept_id),
bus_id int not null references business(bus_id),
content varchar(300),
reply varchar(300)
);

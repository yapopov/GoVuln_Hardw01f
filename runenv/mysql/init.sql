create database vulnapp;
create table vulnapp.user (id int not null auto_increment primary key, name varchar(255) not null,age int not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into vulnapp.user (name,age) values ("Amuro Ray",15),("Char Aznable",20),("Kamille Bidan",17),("Judau Ashta",14),("Banagher Links",16);


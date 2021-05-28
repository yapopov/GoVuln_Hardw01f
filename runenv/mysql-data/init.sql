SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;
create database vulnapp;
grant all privileges on vulnapp.* to root@"vulnapp-goapp.runenv_default" identified by 'rootwolf' with grant option;
create table vulnapp.user (id int not null auto_increment primary key, name varchar(255) not null,mail varchar(255),age int not null,passwd varchar(255) not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into vulnapp.user (name,mail,age,passwd) values ("Amuro Ray","RX-78-2@EFSF.com",15,"Amuro,Ikima-su!"),("Char Aznable","MS-06-S@Zeon.com",20,"AkaiSuisei"),("Banagher Links","RX-0@londo.bell",16,"CongratulationsNowYouHaveMasteredSQLinjection");
create table vulnapp.sessions (uid int,sessionid varchar(128));
create table vulnapp.userdetails (uid int not null primary key, userimage varchar(64), address varchar(64), animal varchar(32), word varchar(64));
insert vulnapp.userdetails(uid,userimage,address,animal,word) values (1,"amuro.png","SIDE-7","GANDOM","こいつ...、動くぞ!"),(2,"char.png","SIDE-3","ZAKU","連邦のMSは化物か!?"),(3,"unicorn.png","INDUSTRIAL-7","UNICORN GANDOM","人の未来は...人が作るものだ!!");
create table vulnapp.posts (postid int not null primary key auto_increment, uid int not null, post varchar(256) not null, created_at timestamp not null default current_timestamp);
create table vulnapp.admins (adminid int primary key not null auto_increment, mail varchar(32), passwd varchar(64));
insert into vulnapp.admins(mail,passwd) values ("admin@admin.com","Qwerty1234"),("amuro@ray.com","RX-78-2");
create table vulnapp.adminsessions(adminsid int auto_increment not null primary key, adminsessionid varchar(64));

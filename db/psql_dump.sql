create database titanic;
create role container with password 'test' login;
grant all privileges on database titanic to container;
alter database titanic owner to container;
\connect titanic container;
create schema person;
create table people(uuid varchar(255) NOT NULL,info json NOT NULL);

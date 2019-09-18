# golangDataBase
Simple program to perform CRUD Operations using Golang.

Make sure to create a table and database in MYSQL: 


create database go_db;

use go_db;

create table customer(id int, name varchar(50));


Configure you're port number, username, MYSQL password in the code under 1.function to get a database connection.




This is a simple program that consists of the following Operations in Golang :


1.function to get a database connection.

2.function to insert a row in customer table.

3.function to update a customer record by customer id.

4.function to delete a customer by customer id.



After successfully creating the program open shell and execute the code by:
(Open powershell window in the source file directory)


Go build main.go

Go run main.go



Now check you're table from MYSQL using :   select * from customer;

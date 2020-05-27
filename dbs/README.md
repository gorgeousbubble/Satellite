# Database package interfaces
The Database package function interfaces description.

## Introduction
Database package is a functional package. It can connect to various of database by similar interface. Related packages are mainly support by open source code in Github. Now support SQL and Non-SQL database like SQLServer, MySQL, SQLite3, Redis, etc.

## Feature of package
The package is mainly used for connect to database which has been realized by 'dbs' package. You can refer to corresponding go file.

#### Interface of database
  * The dbs package offer interfaces let program easy to connect to database
  * Support various databases which has been operated by 'dbs' package, like SQLServer, MySQL, SQLite3, Redis, etc.
  * Similar interfaces and methods
  * Useful and effective

## Usage of interfaces
  * When you first connect to database, you can call method 'Connect(...)' to connect to database.
  * When program finished or need close connection with database, you can call method 'Close(...)' to disconnect with database.
  * You can call 'Insert', 'Update', 'Delete' functions to operate the database.
  * If you want to get data from database, please use 'Query' or 'QueryRow' to fetch the data from database.
  * It just a demo of database operation, you may extend the interfaces when you want to use function futher more.

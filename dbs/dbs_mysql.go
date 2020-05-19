package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// MYSQL Connect function
// this function is mainly use to connect to MYSQL database
// return err indicate the success or failure function execute
func (db *TMySQL) Connect() (err error) {
	// splice data source name
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", db.User, db.Password, db.Protocol, db.Host, db.Port, db.DataBase)
	// connect to mysql database
	db.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connect to MySQL database:", err)
		return err
	}
	// set database configuration
	db.DB.SetConnMaxLifetime(3 * time.Minute)
	db.DB.SetMaxOpenConns(100)
	db.DB.SetMaxIdleConns(16)
	return err
}

// MYSQL Close function
// this function is mainly use to close the connection of MYSQL database
// return err indicate the success or failure function execute
func (db *TMySQL) Close() (err error) {
	err = db.DB.Close()
	if err != nil {
		log.Println("Error close MySQL database:", err)
		return err
	}
	return err
}

// MYSQL QueryRow function
// this function is mainly used to search user by user id segment
// id is a number which define in structure 'TMySQL'
// return err indicate the success or failure function execute
func (db *TMySQL) QueryRow(id int64) (user TUser, err error) {
	query := fmt.Sprintf("select * from %s where id=?", db.DataBase)
	// query one row from database
	err = db.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		log.Println("Error query row from database:", err)
		return user, err
	}
	return user, err
}

// MYSQL Query function
// this function is mainly used to search user
// return err indicate the success or failure function execute
func (db *TMySQL) Query() (users []TUser, err error) {
	query := fmt.Sprintf("select * from %s", db.DataBase)
	// query all rows from database
	rows, err := db.DB.Query(query)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		log.Println("Error query from database:", err)
		return users, err
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Println("Error get columns:", err)
		return users, err
	}
	i := 0
	users = make([]TUser, len(cols))
	for rows.Next() {
		err = rows.Scan(&users[i].ID, &users[i].Name, &users[i].Password)
		if err != nil {
			log.Println("Error scan:", err)
			return users, err
		}
		i++
	}
	return users, err
}

// MYSQL Insert function
// this function is mainly used to add user to database
// input user instance 'TUser'
// return err indicate the success or failure function execute
func (db *TMySQL) Insert(user TUser) (id int64, ra int64, err error) {
	query := fmt.Sprintf("insert into %s(name,password) value(?,?)", db.DataBase)
	// insert one row into database
	r, err := db.DB.Exec(query, user.Name, user.Password)
	if err != nil {
		log.Println("Error insert into database:", err)
		return id, ra, err
	}
	// get last insert id
	id, err = r.LastInsertId()
	if err != nil {
		log.Println("Error get last insert id:", err)
		return id, ra, err
	}
	// get rows affected
	ra, err = r.RowsAffected()
	if err != nil {
		log.Println("Error get rows affected:", err)
		return id, ra, err
	}
	return id, ra, err
}

// MYSQL Update function
// this function is mainly used to update a user in database
// input user instance 'TUser'
// return err indicate the success or failure function execute
func (db *TMySQL) Update(user TUser) (ra int64, err error) {
	query := fmt.Sprintf("update %s set password=? where id=?", db.DataBase)
	// update one row in database
	r, err := db.DB.Exec(query, user.Password, user.ID)
	if err != nil {
		log.Println("Error update one row in database:", err)
		return ra, err
	}
	// get rows affected
	ra, err = r.RowsAffected()
	if err != nil {
		log.Println("Error get rows affected:", err)
		return ra, err
	}
	return ra, err
}

// MYSQL Delete function
// this function is mainly used to delete a user in database
// input user instance 'TUser'
// return err indicate the success or failure function execute
func (db *TMySQL) Delete(user TUser) (id int64, ra int64, err error) {
	query := fmt.Sprintf("delete from %s where id=?", db.DataBase)
	// delete one row in database
	r, err := db.DB.Exec(query, user.ID)
	if err != nil {
		log.Println("Error delete one row in database:", err)
		return id, ra, err
	}
	// get last insert id
	id, err = r.LastInsertId()
	if err != nil {
		log.Println("Error get last insert id:", err)
		return id, ra, err
	}
	// get rows affected
	ra, err = r.RowsAffected()
	if err != nil {
		log.Println("Error get rows affected:", err)
		return id, ra, err
	}
	return id, ra, err
}

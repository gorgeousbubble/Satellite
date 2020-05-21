package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// SQLite3 Connect function
// this function is mainly use to connect to SQLite3 database
// return err indicate the success or failure function execute
func (db *TSQLite3) Connect() (err error) {
	// splice data source name
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=memory", db.DataBase)
	// connect to sqlite3 database
	db.DB, err = sql.Open("sqlite3", dsn)
	if err != nil {
		log.Println("Error connect to SQLite3 database:", err)
		return err
	}
	// set database configuration
	db.DB.SetConnMaxLifetime(3 * time.Minute)
	db.DB.SetMaxOpenConns(100)
	db.DB.SetMaxIdleConns(16)
	return err
}

// SQLite3 Close function
// this function is mainly use to close the connection of SQLite3 database
// return err indicate the success or failure function execute
func (db *TSQLite3) Close() (err error) {
	err = db.DB.Close()
	if err != nil {
		log.Println("Error close SQLite3 database:", err)
		return err
	}
	return err
}

func (db *TSQLite3) QueryRow(id int64) (user TUser, err error) {
	query := fmt.Sprintf("select * from %s where id=?", db.DataBase)
	// query one row from database
	err = db.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		log.Println("Error query row from database:", err)
		return user, err
	}
	return user, err
}

func (db *TSQLite3) Query() (users []TUser, err error) {
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

func (db *TSQLite3) Insert(user TUser) (id int64, ra int64, err error) {
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

func (db *TSQLite3) Update(user TUser) (ra int64, err error) {
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

func (db *TSQLite3) Delete(user TUser) (id int64, ra int64, err error) {
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

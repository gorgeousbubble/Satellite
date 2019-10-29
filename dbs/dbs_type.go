package dbs

import "database/sql"

type TMySQL struct {
	DB       *sql.DB
	User     string
	Password string
	Host     string
	Port     int
	Protocol string
	DataBase string
}

type TUser struct {
	ID       int64          `db:"id"`
	Name     sql.NullString `db:"name"`
	Password sql.NullString `db:"password"`
}

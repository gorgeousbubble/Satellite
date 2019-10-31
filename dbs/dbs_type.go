package dbs

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type TMySQL struct {
	DB       *sql.DB
	User     string
	Password string
	Host     string
	Port     int
	Protocol string
	DataBase string
}

type TMSSQL struct {
	DB       *sql.DB
	User     string
	Password string
	Host     string
	Port     int
	DataBase string
}

type TSQLite3 struct {
	DB       *sql.DB
	DataBase string
}

type TRedis struct {
	DB       *redis.Client
	Host     string
	Port     string
	Password string
}

type TUser struct {
	ID       int64          `db:"id"`
	Name     sql.NullString `db:"name"`
	Password sql.NullString `db:"password"`
}

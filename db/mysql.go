package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Mysql *sql.DB

func InitMysql(dsn string) {
	handler, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Mysql = handler
}

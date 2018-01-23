package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/CX1ng/proxypool/common"
)

var Mysql *sql.DB

func InitMysql(cfg *common.MysqlConfig) {
	handler, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		panic(err)
	}
	handler.SetMaxIdleConns(cfg.MaxIdle)
	handler.SetMaxOpenConns(cfg.MaxOpen)

	Mysql = handler
}

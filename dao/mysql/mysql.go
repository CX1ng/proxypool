package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/CX1ng/proxypool/common"
)

const (
	createDatabaseSql = `CREATE DATABASE IF NOT EXISTS %s`
	dropDatabaseSql   = `DROP DATABASE %s`
	createIPListSql   = `CREATE TABLE IF NOT EXISTS ip_list(
		id bigint PRIMARY KEY AUTO_INCREMENT,
		ip varchar(16) NOT NULL COMMENT "抓取的代理地址",
    	port int NOT NULL COMMENT "代理地址端口",
    	type varchar(8) NOT NULL COMMENT "类型(http/https)",
   	 	origin varchar(16) NOT NULL COMMENT "来源站",
    	raw_time varchar(32) NOT NULL COMMENT "源站爬取时间",
   	 	region varchar(64)  COMMENT "地区",
    	capture_time datetime NOT NULL COMMENT "爬取时间",
    	last_verify_time datetime NOT NULL COMMENT "最后验证时间",
		create_time datetime NOT NULL COMMENT "创建时间",
		UNIQUE KEY(ip, port) 
	)DEFAULT CHARSET=utf8`
	userDatabaseSql = `use %s`
)

var sqlExecHandler *sql.DB

func GetDBHandler() *sql.DB {
	if sqlExecHandler == nil {
		initMysqlStorage(common.GetConfigHandler().Mysql)
	}
	return sqlExecHandler
}

func initMysqlStorage(cfg *common.MysqlConfig) {
	initMysql(cfg)
	initDatabase(cfg)
}

func initMysql(cfg *common.MysqlConfig) {
	handler, err := sql.Open("mysql", cfg.Dsn)
	if err != nil {
		panic(err)
	}
	handler.SetMaxIdleConns(cfg.MaxIdle)
	handler.SetMaxOpenConns(cfg.MaxOpen)

	sqlExecHandler = handler
}

func initDatabase(cfg *common.MysqlConfig) {
	if sqlExecHandler == nil {
		panic(common.ErrMysqlHandlerNotInit)
	}
	createDbSql := fmt.Sprintf(createDatabaseSql, cfg.DbName)
	if _, err := sqlExecHandler.Exec(createDbSql); err != nil {
		panic(err)
	}
	useDbSql := fmt.Sprintf(userDatabaseSql, cfg.DbName)
	if _, err := sqlExecHandler.Exec(useDbSql); err != nil {
		panic(err)
	}
	if _, err := sqlExecHandler.Exec(createIPListSql); err != nil {
		panic(err)
	}
}

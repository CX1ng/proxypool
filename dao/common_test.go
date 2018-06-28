package dao

import (
	"database/sql"
)

var testDB *sql.DB

const (
	testCreateDatabaseSql = "create database test_proxy_pool"
	testDropDatabaseSql   = "drop database if exists test_proxy_pool"
	testChangeDatabaseSql = "use test_proxy_pool"
)

func initTestDB() {
	var err error
	testDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	if _, err := testDB.Exec(testDropDatabaseSql); err != nil {
		panic(err)
	}
	if _, err := testDB.Exec(testCreateDatabaseSql); err != nil {
		panic(err)
	}
	if _, err := testDB.Exec(testChangeDatabaseSql); err != nil {
		panic(err)
	}
	if _, err := testDB.Exec(createIPListSql); err != nil {
		panic(err)
	}
}

func init() {
	initTestDB()
}

package main

import (
	"flag"
	"net/http"
	"strings"

	. "github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
	. "github.com/CX1ng/proxypool/dao/mysql"
	"github.com/CX1ng/proxypool/processord"
	"github.com/CX1ng/proxypool/server"
)

var configPath = flag.String("config", "./config/config.dev.toml", "config path")

func InitStorage() (Import, error) {
	switch strings.ToLower(GetConfigHandler().Storage) {
	case "mysql":
		InitMysqlStorage(GetConfigHandler().Mysql)
		return DBConnector{DB: GetDBHandler()}, nil
	default:
		return nil, ErrStorageNotSupport
	}
}

func main() {
	flag.Parse()

	//init config
	InitConfig(*configPath)

	//init storage
	db, err := InitStorage()
	if err != nil {
		panic(err)
	}

	storage := processord.NewStorage(db)
	for _, detail := range GetConfigHandler().ProxyWebs {
		processor, err := processord.NewProcessor(detail, storage.Queue)
		if err != nil {
			panic(err)
		}
		go processor.Run()
	}

	go storage.GetIPInfoFromChannel()

	router := server.NewProxyPoolRouter()
	if err := http.ListenAndServe(GetConfigHandler().Listen, router); err != nil {
		panic(err)
	}
}

package main

import (
	"flag"
	"net/http"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/dao"
	"github.com/CX1ng/proxypool/processord"
	"github.com/CX1ng/proxypool/server"
)

var configPath = flag.String("config", "./config/config.dev.toml", "config path")

func InitStorage() (processord.Import, error) {
	switch common.Config.Storage {
	case "Mysql":
		dao.InitMysql(common.Config.Mysql)
		dao.InitDatabase()
		return dao.DBConnector{DB: dao.Mysql}, nil
	default:
		return nil, common.StorageNotSupport
	}
}

func main() {
	flag.Parse()

	//init config
	common.InitConfig(*configPath)

	//init storage
	db, err := InitStorage()
	if err != nil {
		panic(err)
	}

	storage := processord.NewStorage(db)
	kddPro, err := processord.NewProcessor("kuaidaili", 1, 100, storage.Queue)
	if err != nil {
		panic(err)
	}
	go kddPro.Run()
	xcPro, err := processord.NewProcessor("xici", 1, 100, storage.Queue)
	if err != nil {
		panic(err)
	}
	go xcPro.Run()

	go storage.GetIPInfoFromChannel()

	router := server.NewProxyPoolRouter()
	if err := http.ListenAndServe(common.Config.Listen, router); err != nil {
		panic(err)
	}
}

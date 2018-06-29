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

func main() {
	flag.Parse()

	//init config
	common.InitConfig(*configPath)

	//init mysql
	dao.InitMysql(common.Config.Mysql)

	//init database
	dao.InitDatabase()

	// 考虑将采集相关操作交由Restful API控制
	//init processord

	//init Storage
	storage := processord.NewStorage()
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

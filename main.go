package main

import (
	"flag"
	"net/http"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/db"
	"github.com/CX1ng/proxypool/server"
)

var configPath = flag.String("config", "./config/config.dev.toml", "config path")

func main() {
	flag.Parse()

	//init config
	common.InitConfig(*configPath)

	//init mysql
	db.InitMysql(common.Config.Mysql)

	//init processord

	//init Storage
	// storage := processord.NewStorage()

	// kddPro := processord.NewProcessor(parser.NewKuaiDaiLi(1, 100, storage.Queue))
	// go kddPro.StartExec()
	// xcPro := processord.NewProcessor(parser.NewXiCi(1, 100, storage.Queue))
	// go xcPro.StartExec()

	// storage.GetIPInfoFromChannel()

	router := server.NewProxyPoolRouter()
	if err := http.ListenAndServe(common.Config.Listen, router); err != nil {
		panic(err)
	}
}

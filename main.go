package main

import (
	"flag"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/db"
	"github.com/CX1ng/proxypool/processord"
	"github.com/CX1ng/proxypool/processord/parser"
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
	storage := processord.NewStorage()

	pro := processord.NewProcessor(parser.NewKuaiDaiLi(1, 5, storage.Queue))
	go pro.StartExec()

	storage.GetIPInfoFromChannel()
}

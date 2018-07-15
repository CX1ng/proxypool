package main

import (
	"flag"
	"net/http"
	"strings"

	. "github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/dao"
	"github.com/CX1ng/proxypool/dao/mysql"
	"github.com/CX1ng/proxypool/dao/redis"
	"github.com/CX1ng/proxypool/processord"
	"github.com/CX1ng/proxypool/server"
	"runtime"
)

var configPath = flag.String("config", "./config/config.dev.toml", "config path")

func main() {
	flag.Parse()

	//init config
	InitConfig(*configPath)
	runtime.GOMAXPROCS(GetConfigHandler().MaxProcs)

	//init storage
	initializer, ok := dao.StorageInitializer[strings.ToLower(GetConfigHandler().Storage)]
	if !ok {
		panic(ErrStorageNotSupport)
	}

	storage := processord.NewStorage(initializer())
	for _, detail := range GetConfigHandler().ProxyWebs {
		processor, err := processord.NewProcessor(detail, storage.Queue)
		if err != nil {
			panic(err)
		}
		go processor.Run()
	}

	go storage.VerifyAndInsertIPSWithLoop()

	router := server.NewProxyPoolRouter()
	if err := http.ListenAndServe(GetConfigHandler().Listen, router); err != nil {
		panic(err)
	}
}

// TODO:
// 解决Mysql/Redis引用问题，不然无法触发两个包的init函数
func Noop() {
	redis.Noop()
	mysql.Noop()
}

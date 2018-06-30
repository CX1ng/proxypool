package common

import (
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	Listen  string
	Storage string
	Time    *TimeConfig
	Mysql   *MysqlConfig
}

type TimeConfig struct {
	TimeInterval int //请求的时间间隔
}

type MysqlConfig struct {
	Dsn     string // database dsn
	DbName  string
	MaxOpen int
	MaxIdle int
}

type ProxyPool struct {
	Storage string
}

var configHandler *config

func InitConfig(path string) {
	cfg := new(config)
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		panic(err)
	}
	if _, ok := StorageMaps[strings.ToLower(cfg.Storage)]; !ok {
		panic(ErrStorageNotSupport)
	}
	configHandler = cfg
}

func GetConfigHandler() *config {
	if configHandler == nil {
		panic(ErrConfigHandlerNotInit)
	}
	return configHandler
}

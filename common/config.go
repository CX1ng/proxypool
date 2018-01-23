package common

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Time  *TimeConfig
	Mysql *MysqlConfig
}

type TimeConfig struct {
	TimeInterval int //请求的时间间隔
}

type MysqlConfig struct {
	Dsn     string // database dsn
	MaxOpen int
	MaxIdle int
}

var Config *config

func InitConfig(path string) {
	cfg := new(config)
	if _, err := toml.DecodeFile(path, cfg); err != nil {
		panic(err)
	}
	Config = cfg
}

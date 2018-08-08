package common

import (
	"time"

	"github.com/BurntSushi/toml"
)

type config struct {
	Listen    string
	Storage   string
	Mysql     *MysqlConfig
	Redis     *RedisConfig
	ProxyWebs []WebDetail `toml:"proxy_web"`
	MaxProcs  int
}

type MysqlConfig struct {
	Dsn     string // database dsn
	DbName  string
	MaxOpen int
	MaxIdle int
}

type RedisConfig struct {
	Dsn      string
	Timeout  Duration
	Protocol string
	Db       int
}

type WebDetail struct {
	Name             string
	TaskType         string
	BeginPageNum     int
	EndPageNum       int
	TimeInterval     Duration
	LoopTimeInterval Duration
}

func (w *WebDetail) Validation() error {
	if w.BeginPageNum < 0 {
		return ErrBeginPageNumLessThanOne
	}
	if w.EndPageNum < 0 {
		return ErrEndPageNumLessThanZero
	}
	return nil
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
	configHandler = cfg
}

func GetConfigHandler() *config {
	if configHandler == nil {
		panic(ErrConfigHandlerNotInit)
	}
	return configHandler
}

// Duration 配置中使用的时长
type Duration struct {
	time.Duration
}

// UnmarshalText 将字符串形式的时长信息转换为Duration类型
func (d *Duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// D 从Duration struct中取出time.Duration类型的值
func (d *Duration) D() time.Duration {
	return d.Duration
}

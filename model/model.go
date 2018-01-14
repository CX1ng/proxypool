package model

import (
	"time"
)

type ProxyIP struct {
	ID             int       //编号
	IP             string    //代理IP
	Port           int       //代理端口
	Type           string    //类型(http/https)
	Origin         string    //来源站
	RawTime        string    //来源站爬取时间
	Region         string    //代理IP地区
	CaptureTime    time.Time //从来源站爬取的时间
	LastVerifyTime time.Time //最后检查时间
	CreateTime     time.Time //创建时间
}

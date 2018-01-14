package main

import (
	"fmt"
	"time"

	"github.com/CX1ng/proxypool/db"
	"github.com/CX1ng/proxypool/model"
	"github.com/CX1ng/proxypool/processord/parser"
)

func test() {
	channel := make(chan *model.ProxyIP, 100)
	kdd := parser.NewKuaiDaiLi(1, 5, channel)
	go func() {
		for {
			select {
			case info := <-channel:
				fmt.Printf("get ip:%s\n", info.IP)
				err := model.InsertIP(db.Mysql, info)
				if err != nil {
					fmt.Printf("err : %s\n", err)
				}
			case <-time.After(60 * time.Second):
				fmt.Printf("channel quit\n")
				break
			}
		}
	}()
	kdd.RangePage()
}

func main() {
	db.InitMysql("root:@tcp(127.0.0.1:3306)/proxy_pool?charset=utf8mb4,utf8&parseTime=True&loc=Local")
	test()
}

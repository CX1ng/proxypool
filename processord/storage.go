package processord

import (
	"fmt"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/db"
	"github.com/CX1ng/proxypool/model"
)

type Storage struct {
	Done  <-chan bool
	Queue chan *model.ProxyIP
}

func NewStorage() *Storage {
	return &Storage{
		Done:  make(chan bool),
		Queue: make(chan *model.ProxyIP, common.StorageChannelCapacity),
	}
}

func (s *Storage) GetIPInfoFromChannel() {
	for {
		select {
		case info := <-s.Queue:
			fmt.Printf("get ip:%s\n", info.IP)
			err := model.InsertIP(db.Mysql, info)
			if err != nil {
				fmt.Printf("err : %s\n", err)
			}
		case <-s.Done:
			break
		case <-time.After(common.StorageChannelTimeout * time.Second):
			fmt.Printf("channel quit\n")
			break
		}
	}
}

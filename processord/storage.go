package processord

import (
	"fmt"
	"strconv"
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
	exit := false
	routinesCount := make(chan bool, 20)
	for {
		select {
		case info := <-s.Queue:
			routinesCount <- true
			go verifyAndInsert(info)
			<-routinesCount
		case <-s.Done:
			break
		case <-time.After(common.StorageChannelTimeout * time.Second):
			fmt.Printf("channel quit\n")
			exit = true
		}
		if exit == true {
			break
		}
	}
}

func verifyAndInsert(info *model.ProxyIP) {
	if VerifyProxy(info.IP, strconv.Itoa(info.Port)) == false {
		return
	}
	fmt.Printf("Insert %s into mysql\n", info.IP)
	err := model.InsertIP(db.Mysql, info)
	if err != nil {
		fmt.Printf("err : %s\n", err)
	}
}

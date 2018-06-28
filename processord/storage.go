package processord

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
	"github.com/CX1ng/proxypool/models"
)

type Storage struct {
	done  <-chan bool
	Queue chan []models.ProxyIP
	db    DBConnector
}

func NewStorage() *Storage {
	return &Storage{
		done:  make(chan bool),
		Queue: make(chan []models.ProxyIP, common.StorageChannelCapacity),
		db:    DBConnector{DB: Mysql},
	}
}

func (s *Storage) GetIPInfoFromChannel() {
	exit := false
	routinesCount := make(chan bool, 20)
	for {
		select {
		case info := <-s.Queue:
			routinesCount <- true
			go verifyAndInsert(s.db, info)
			<-routinesCount
		case <-s.done:
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

func verifyAndInsert(db DBConnector, infos []models.ProxyIP) {
	ipList := make([]models.ProxyIP, 0, len(infos))
	for _, info := range infos {
		if ok := VerifyProxy(info["ip"].(string), strconv.Itoa(info["port"].(int))); !ok {
			continue
		}
		info["last_verify_time"] = time.Now()
		info["create_time"] = time.Now()
		ipList = append(ipList, info)
	}

	if len(ipList) == 0 {
		return
	}
	for _, info := range ipList {
		fmt.Printf("Proxy_ip:%+v\n", info)
	}
	err := db.BulkInsertProxyIPs(ipList)
	if err != nil {
		fmt.Printf("err : %s\n", err)
	}
}

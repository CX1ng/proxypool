package processord

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
)

type Storage struct {
	done  <-chan bool
	Queue chan []models.ProxyIP
	imp   Import
}

func NewStorage(imp Import) *Storage {
	return &Storage{
		done:  make(chan bool),
		Queue: make(chan []models.ProxyIP, common.StorageChannelCapacity),
		imp:   imp,
	}
}

func (s *Storage) GetIPInfoFromChannel() {
	exit := false
	routinesCount := make(chan bool, 20)
	for {
		select {
		case info := <-s.Queue:
			routinesCount <- true
			go verifyAndInsert(s.imp, info)
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

func verifyAndInsert(imp Import, infos []models.ProxyIP) {
	ipList := make([]models.ProxyIP, 0, len(infos))
	for _, info := range infos {
		if ok := VerifyProxy(info["ip"].(string), strconv.Itoa(info["port"].(int))); !ok {
			continue
		}
		info["last_verify_time"] = time.Now().Format("2006-01-02 15:04:05")
		info["create_time"] = time.Now().Format("2006-01-02 15:04:05")
		ipList = append(ipList, info)
	}

	if len(ipList) == 0 {
		return
	}
	for _, info := range ipList {
		fmt.Printf("Proxy_ip:%+v\n", info)
	}
	err := imp.BulkInsertProxyIPs(ipList)
	if err != nil {
		fmt.Printf("err : %s\n", err)
	}
}

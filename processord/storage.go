package processord

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
)

type Storage struct {
	Queue chan []models.ProxyIP
	imp   Import
}

func NewStorage(imp Import) *Storage {
	return &Storage{
		Queue: make(chan []models.ProxyIP, common.StorageChannelCapacity),
		imp:   imp,
	}
}

func (s *Storage) GetIPInfoFromChannel() {
	// routinesCount控制并发度
	routinesCount := make(chan bool, common.StorageConcurrencyRoutineCount)
	for {
		select {
		case info := <-s.Queue:
			routinesCount <- true
			go s.verifyAndInsert(info)
			<-routinesCount
		case <-time.After(3 * time.Second):
			fmt.Printf("Not Receive Data. Waiting...\n")
		}
	}
}

func (s *Storage) verifyAndInsert(infos []models.ProxyIP) {
	for _, info := range infos {
		if ok := VerifyProxy(info["ip"].(string), strconv.Itoa(info["port"].(int))); !ok {
			continue
		}
		info["last_verify_time"] = time.Now().Format("2006-01-02 15:04:05")
		info["create_time"] = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("Proxy_ip:%+v\n", info)
	}
	if len(infos) == 0 {
		return
	}
	if err := s.imp.BulkInsertProxyIPs(infos); err != nil {
		fmt.Printf("err : %s\n", err)
	}
}

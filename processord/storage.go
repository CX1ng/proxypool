package processord

import (
	"fmt"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/dao"
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
	infos = BulkVerifyProxyIPs(infos)
	if len(infos) == 0 {
		return
	}
	if err := s.imp.BulkInsertProxyIPs(infos); err != nil {
		fmt.Printf("err : %s\n", err)
	}
}

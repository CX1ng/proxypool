package processord

import (
	"fmt"
	"time"

	"github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/dao"
	"github.com/CX1ng/proxypool/models"
)

const (
	arrayLen = 1
)

type Storage struct {
	Queue      chan []models.ProxyIP
	waitInsert chan models.ProxyIP
	imp        Importer
}

func NewStorage(imp Importer) *Storage {
	return &Storage{
		Queue:      make(chan []models.ProxyIP, common.StorageChannelCapacity),
		imp:        imp,
		waitInsert: make(chan models.ProxyIP, common.StorageChannelCapacity),
	}
}

func (s *Storage) VerifyAndInsertIPSWithLoop() {
	// routinesCount控制并发度
	routinesCount := make(chan bool, common.StorageConcurrencyRoutineCount)
	infos := make([]models.ProxyIP, 0, arrayLen)
	for {
		select {
		case info := <-s.Queue: // 验证队列
			routinesCount <- true
			go BulkVerifyProxyIPs(s.waitInsert, info)
			<-routinesCount
		case info := <-s.waitInsert: // 存储队列
			infos = append(infos, info)
			if len(infos) > arrayLen {
				go s.imp.BulkInsertProxyIPs(infos)
				infos = make([]models.ProxyIP, 0, arrayLen)
			}
		case <-time.After(3 * time.Second): // 超时等待
			fmt.Printf("Not Receive Data. Waiting...\n")
		}
	}
}

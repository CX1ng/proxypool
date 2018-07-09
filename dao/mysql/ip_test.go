package mysql

import (
	"testing"
	"time"

	. "github.com/CX1ng/proxypool/models"
	"github.com/stretchr/testify/assert"
)

func TestBulkInsertProxyIPsAndGetLimitProxyIP(t *testing.T) {
	as := assert.New(t)
	db := DBConnector{DB: testDB}

	now := time.Now().Format("2006-01-02 15:04:05")
	ip1 := NewProxyIP()
	ip1["ip"] = "1.1.1.1"
	ip1["port"] = "80"
	ip1["type"] = "http"
	ip1["origin"] = "kuaidaili"
	ip1["raw_time"] = now
	ip1["region"] = "beijing"
	ip1["capture_time"] = now
	ip1["last_verify_time"] = now
	ip1["create_time"] = now

	ip2 := NewProxyIP()
	ip2["ip"] = "2.2.2.2"
	ip2["port"] = "80"
	ip2["type"] = "http"
	ip2["origin"] = "kuaidaili"
	ip2["raw_time"] = now
	ip2["region"] = "shanghai"
	ip2["capture_time"] = now
	ip2["last_verify_time"] = now
	ip2["create_time"] = now
	infos := []ProxyIP{ip1, ip2}
	err := db.BulkInsertProxyIPs(infos)
	as.Nil(err)
	ips, err := db.GetLimitProxyIP(0)
	as.Nil(err)
	as.Len(ips, 2)
	as.Contains(ips, joinIPPort("1.1.1.1", 80))
	as.Contains(ips, joinIPPort("2.2.2.2", 80))
}

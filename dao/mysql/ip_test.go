package mysql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/CX1ng/proxypool/models"
)

func TestBulkInsertProxyIPsAndGetLimitProxyIP(t *testing.T) {
	as := assert.New(t)
	db := DBConnector{DB: testDB}

	now := time.Now().Format("2006-01-02 15:04:05")
	ip1 := NewProxyIP()
	ip1.Set("ip", "1.1.1.1")
	ip1.Set("port", "80")
	ip1.Set("type", "http")
	ip1.Set("origin", "kuaidaili")
	ip1.Set("raw_time", now)
	ip1.Set("region", "beijing")
	ip1.Set("capture_time", now)
	ip1.Set("last_verify_time", now)
	ip1.Set("create_time", now)

	ip2 := NewProxyIP()
	ip2.Set("ip", "2.2.2.2")
	ip2.Set("port", "80")
	ip2.Set("type", "http")
	ip2.Set("origin", "kuaidaili")
	ip2.Set("raw_time", now)
	ip2.Set("region", "shanghai")
	ip2.Set("capture_time", now)
	ip2.Set("last_verify_time", now)
	ip2.Set("create_time", now)
	infos := []ProxyIP{ip1, ip2}
	err := db.BulkInsertProxyIPs(infos)
	as.Nil(err)
	ips, err := db.GetLimitProxyIP(0)
	as.Nil(err)
	as.Len(ips, 2)
	as.Contains(ips, joinIPPort("1.1.1.1", "80"))
	as.Contains(ips, joinIPPort("2.2.2.2", "80"))
}

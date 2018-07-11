package redis

import (
	"testing"

	. "github.com/CX1ng/proxypool/models"
	"github.com/stretchr/testify/assert"
)

func TestBulkInsertProxyIPs(t *testing.T) {
	as := assert.New(t)

	obj := NewProxyIP()
	obj.Set("ip", "122.114.31.177")
	obj.Set("last_verify_time", "2018-07-14 22:05:14")
	obj.Set("type", "HTTP")
	obj.Set("origin", "西刺")
	obj.Set("capture_time", "20187-14 22:05:14")
	obj.Set("region", "河南郑州")
	obj.Set("port", 808)
	obj.Set("raw_time", "18-0721:51")
	ips := []ProxyIP{obj}

	err := testRedisConnector.BulkInsertProxyIPs(ips)
	as.Nil(err)
}

package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/CX1ng/proxypool/common"
	. "github.com/CX1ng/proxypool/models"
)

func TestBulkInsertProxyIPsAndGetLimitProxyIP(t *testing.T) {
	as := assert.New(t)

	tesIP1 := NewProxyIP()
	tesIP1.Set("ip", "122.114.31.177")
	tesIP1.Set("last_verify_time", "2018-07-14 22:05:14")
	tesIP1.Set("type", "HTTP")
	tesIP1.Set("origin", "西刺")
	tesIP1.Set("capture_time", "20187-14 22:05:14")
	tesIP1.Set("region", "河南郑州")
	tesIP1.Set("port", "808")
	tesIP1.Set("raw_time", "18-0721:51")
	testIP2 := NewProxyIP()
	testIP2.Set("ip", "1.1.1.1")
	testIP2.Set("last_verify_time", "2018-07-14 22:00:00")
	testIP2.Set("type", "HTTP")
	testIP2.Set("origin", "快代理")
	testIP2.Set("capture_time", "2018-07-14 22:00:01")
	testIP2.Set("region", "北京")
	testIP2.Set("port", "80")
	testIP2.Set("raw_time", "2018-07-13 22:00:01")
	ips := []ProxyIP{tesIP1, testIP2}

	err := testRedisConnector.BulkInsertProxyIPs(ips)
	as.Nil(err)
	respIPs, err := testRedisConnector.GetLimitProxyIP(1)
	as.Nil(err)
	as.Len(respIPs, 1)
	respIPs, err = testRedisConnector.GetLimitProxyIP(2)
	as.Nil(err)
	as.Len(respIPs, 2)
	as.Contains(respIPs, "122.114.31.177:808")
	as.Contains(respIPs, "1.1.1.1:80")
	respIPs, err = testRedisConnector.GetLimitProxyIP(3)
	as.Nil(err)
	as.Len(respIPs, 2)
	as.Contains(respIPs, "122.114.31.177:808")
	as.Contains(respIPs, "1.1.1.1:80")
	respIPs, err = testRedisConnector.GetLimitProxyIP(-1)
	as.NotNil(err)
	as.Equal(ErrModelLimitInvalid, err)
}

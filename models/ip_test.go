package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProxyIPGetFIelds(t *testing.T) {
	as := assert.New(t)

	ip := NewProxyIP()
	fields := ip.GetFields()
	as.Equal(len(fields), len(proxyIPFields))
	for _, field := range fields {
		as.Contains(proxyIPFields, field)
	}
}

func TestProxyIPFieldsNum(t *testing.T) {
	as := assert.New(t)

	ip := NewProxyIP()
	num := ip.FieldsNum()
	as.Equal(num, len(proxyIPFields))
}

func TestProxyIPGetMethods(t *testing.T) {
	as := assert.New(t)

	info := NewProxyIP()
	info.Set("ip", "1.1.1.1")
	ip, err := info.IP()
	as.Nil(err)
	as.Equal(ip, "1.1.1.1")
	info.Set("ip", "2.2.2.2")
	ip, err = info.IP()
	as.Nil(err)
	as.Equal(ip, "2.2.2.2")

	info.Set("port", int64(80))
	port, err := info.Port()
	as.Nil(err)
	as.Equal(port, int64(80))

	info.Set("type", "https")
	typeValue, err := info.Type()
	as.Nil(err)
	as.Equal(typeValue, "https")

	info.Set("origin", "xici")
	origin, err := info.Origion()
	as.Nil(err)
	as.Equal(origin, "xici")

	info.Set("region", "beijing")
	region, err := info.Region()
	as.Nil(err)
	as.Equal(region, "beijing")

	now := time.Now().Format("2006-01-02 15:04:05")
	info.Set("create_time", now)
	createTime, err := info.CreateTime()
	as.Nil(err)
	as.Equal(createTime, now)

	info.Set("raw_time", now)
	rawTime, err := info.RawTime()
	as.Nil(err)
	as.Equal(rawTime, now)

	info.Set("capture_time", now)
	captureTime, err := info.CaptureTime()
	as.Nil(err)
	as.Equal(captureTime, now)

	info.Set("last_verify_time", now)
	lastVerifyTime, err := info.LastVerifyTime()
	as.Nil(err)
	as.Equal(lastVerifyTime, now)
}

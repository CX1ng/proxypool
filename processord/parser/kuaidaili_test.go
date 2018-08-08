package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/CX1ng/proxypool/utils"
	"github.com/bouk/monkey"
)

func TestKuaidailiageParser(t *testing.T) {
	as := assert.New(t)
	mockReq := monkey.Patch(utils.HttpRequestWithUserAgent, MockReqWithKdd)
	defer mockReq.Unpatch()

	html, err := utils.HttpRequestWithUserAgent("https://www.kuaidaili.com/free/inha/1")
	as.Nil(err)

	setter := &KuaidailiSetter{}
	kdd := setter.SettingParser()
	as.Nil(err)
	info, err := kdd.PageParser(html)
	as.Nil(err)
	as.NotNil(info)
	as.NotEqual(len(info), 0)
	ip, ok := info[0]["ip"]
	as.NotNil(ip)
	as.True(ok)
	port, ok := info[0]["port"]
	as.NotNil(port)
	as.True(ok)
	origin, ok := info[0]["origin"]
	as.NotNil(origin)
	as.True(ok)
	region, ok := info[0]["region"]
	as.NotNil(region)
	as.True(ok)
	rawTime, ok := info[0]["raw_time"]
	as.NotNil(rawTime)
	as.True(ok)
	captureTime, ok := info[0]["capture_time"]
	as.NotNil(captureTime)
	as.True(ok)
	typeValue, ok := info[0]["type"]
	as.NotNil(typeValue)
	as.True(ok)
}

func TestGetMaxPageNumWithKdd(t *testing.T) {
	as := assert.New(t)
	mockReq := monkey.Patch(utils.HttpRequestWithUserAgent, MockReqWithKdd)
	defer mockReq.Unpatch()

	setter := KuaidailiSetter{}
	kdd := setter.SettingParser()
	maxPage, err := kdd.GetMaxPageNum(10)
	as.Nil(err)
	time.Sleep(1 * time.Second)
	maxPage, err = kdd.GetMaxPageNum(999999)
	as.Nil(err)
	as.NotEqual(999999, maxPage)
}

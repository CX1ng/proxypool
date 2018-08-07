package parser

import (
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/CX1ng/proxypool/utils"
	"github.com/bouk/monkey"
)

func TestXiciageParser(t *testing.T) {
	as := assert.New(t)
	mockReq := monkey.Patch(utils.HttpRequestWithUserAgent, MockReqWithXici)
	defer mockReq.Unpatch()

	html, err := utils.HttpRequestWithUserAgent("http://www.xicidaili.com/nn/1")
	as.Nil(err)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	setter := &XiciSetter{}
	xici := setter.SettingParser()
	as.Nil(err)
	info := xici.PageParser(doc)
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

func TestGetMaxPageNumWithXici(t *testing.T) {
	as := assert.New(t)
	mockReq := monkey.Patch(utils.HttpRequestWithUserAgent, MockReqWithXici)
	defer mockReq.Unpatch()

	setter := XiciSetter{}
	xc := setter.SettingParser()
	maxPage, err := xc.GetMaxPageNum(10)
	as.Nil(err)
	as.Equal(10, maxPage)
	time.Sleep(1 * time.Second)
	maxPage, err = xc.GetMaxPageNum(999999)
	as.Nil(err)
	as.NotEqual(999999, maxPage)
}

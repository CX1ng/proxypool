package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/CX1ng/proxypool/utils"
)

func TestIp66PageParser(t *testing.T) {
	as := assert.New(t)

	html, err := utils.HttpRequestWithUserAgent("http://www.66ip.cn/1.html")
	as.Nil(err)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	setter := &Ip66Setter{}
	ip66 := setter.SettingParser()
	as.Nil(err)
	info := ip66.PageParser(doc)
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

func TestGetMaxPageNumWith66Ip(t *testing.T) {
	as := assert.New(t)
	setter := Ip66Setter{}
	ip66 := setter.SettingParser()
	pageNum, err := ip66.GetMaxPageNum(10)
	as.Nil(err)
	as.Equal(10, pageNum)
}

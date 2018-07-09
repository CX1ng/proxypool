package parser

import (
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/CX1ng/proxypool/common"
)

func TestXiciageParser(t *testing.T) {
	as := assert.New(t)

	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.xicidaili.com/nn/1", nil)
	as.Nil(err)
	request.Header.Add("User-Agent", common.UserAgent)
	resp, err := client.Do(request)
	as.Nil(err)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)

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
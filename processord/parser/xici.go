package parser

import (
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/models"
)

func init() {
	WebParsers.SetParserSetter("xici", &XiciSetter{})
}

type XiciSetter struct{}

func (x *XiciSetter) SettingParser() Parser {
	return &Xici{
		name: "xici",
		url:  "http://www.xicidaili.com/nn/",
	}
}

type Xici struct {
	name string
	url  string
}

func (x *Xici) GetParserName() string {
	return x.name
}

func (x *Xici) GetUrl() string {
	return x.url
}

func (x *Xici) PageParser(doc *goquery.Document) []ProxyIP {
	var ips []ProxyIP
	ipTable := doc.Find("table#ip_list").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(1).Text()
		port := node.Find("td").Eq(2).Text()
		proxyType := node.Find("td").Eq(5).Text()
		region := node.Find("td").Eq(3).Find("a").Text()
		rawTime := node.Find("td").Eq(9).Text()
		ipInfo := NewProxyIP()
		ipInfo["ip"] = ip
		ipInfo["port"] = port
		ipInfo["type"] = proxyType
		ipInfo["origin"] = "西刺"
		ipInfo["region"] = region
		ipInfo["raw_time"] = rawTime
		ipInfo["capture_time"] = time.Now().Format("2006-01-02 15:04:05")
		ips = append(ips, ipInfo)
	})
	return ips
}

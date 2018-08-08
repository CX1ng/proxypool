package parser

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/utils"
)

func init() {
	WebParsers.SetParserSetter("kuaidaili", &KuaidailiSetter{})
}

type KuaidailiSetter struct{}

func (k *KuaidailiSetter) SettingParser() Parser {
	return &Kuaidaili{
		name: "kuaidaili",
		url:  "https://www.kuaidaili.com/free/inha/",
	}
}

type Kuaidaili struct {
	name string
	url  string
}

func (k *Kuaidaili) GetParserName() string {
	return k.name
}

func (k *Kuaidaili) GetUrl() string {
	return k.url
}

func (k *Kuaidaili) PageParser(html string) ([]ProxyIP, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	var ips []ProxyIP
	ipTable := doc.Find("table.table.table-bordered.table-striped").Find("tbody").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(0).Text()
		port := node.Find("td").Eq(1).Text()
		proxyType := node.Find("td").Eq(3).Text()
		region := node.Find("td").Eq(4).Text()
		rawTime := node.Find("td").Eq(6).Text()
		ipInfo := NewProxyIP()
		ipInfo["ip"] = ip
		ipInfo["port"] = port
		ipInfo["type"] = proxyType
		ipInfo["origin"] = "快代理"
		ipInfo["region"] = region
		ipInfo["raw_time"] = rawTime
		ipInfo["capture_time"] = time.Now().Format("2006-01-02 15:04:05")
		ips = append(ips, ipInfo)
	})
	return ips, nil
}

func (k *Kuaidaili) GetMaxPageNum(maxNum int) (int, error) {
	html, err := HttpRequestWithUserAgent(k.url + "1")
	if err != nil {
		return -1, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return -1, err
	}
	maxPageNum, err := strconv.Atoi(doc.Find("#listnav").Find("ul").Find("li:nth-of-type(9)").Text())
	if err != nil {
		return -1, err
	}
	if maxNum == 0 || maxPageNum < maxNum {
		return maxPageNum, nil
	}
	return maxNum, nil
}

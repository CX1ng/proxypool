package parser

import (
	"strings"
	"time"
	"strconv"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/utils"
)

func init() {
	WebParsers.SetParserSetter("66ip", &Ip66Setter{})
}

type Ip66Setter struct{}

// FIXME:
// [ ] Support: url: http://www.66ip.cn/%d.html
func (i *Ip66Setter) SettingParser() Parser {
	return &Ip66{
		name: "66ip",
		url:  "http://www.66ip.cn/",
	}
}

type Ip66 struct {
	name string
	url  string
}

func (i *Ip66) GetParserName() string {
	return i.name
}

func (i *Ip66) GetUrl() string {
	return i.url
}

func (i *Ip66) PageParser(doc *goquery.Document) []ProxyIP {
	var ips []ProxyIP
	ipTable := doc.Find("div.containerbox").Find("table").Find("tbody").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		port := node.Find("td").Eq(1).Text()
		ip := node.Find("td").Eq(0).Text()
		region := GBK2UTF(node.Find("td").Eq(2).Text())
		rawTime := GBK2UTF(node.Find("td").Eq(4).Text())
		ipInfo := NewProxyIP()
		ipInfo["ip"] = ip
		ipInfo["port"] = port
		ipInfo["type"] = "http"
		ipInfo["origin"] = "66ip"
		ipInfo["region"] = region
		ipInfo["raw_time"] = rawTime
		ipInfo["capture_time"] = time.Now().Format("2006-01-02 15:04:05")
		ips = append(ips, ipInfo)
	})
	return ips
}

func (i *Ip66) GetMaxPageNum(maxNum int) (int, error) {
	html, err := HttpRequestWithUserAgent(i.url + "1.html")
	if err != nil {
		return -1, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return -1, err
	}
	maxPageNum, err := strconv.Atoi(doc.Find("div#PageList").Find("a:nth-last-child(2)").Text())
	if err != nil {
		return -1, err
	}
	if maxPageNum < maxNum {
		return maxPageNum, nil
	}
	return maxNum, nil
}

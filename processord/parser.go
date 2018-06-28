package processord

import (
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/models"
)

type Parser func(doc *goquery.Document) []ProxyIP

var WebParser = make(map[string]Parser)

func init() {
	WebParser["kuaidaili"] = KuaiDaiLiPageParser
	WebParser["xici"] = XiciPageParser
}

// map 将代理名与解析器绑定到一起
func KuaiDaiLiPageParser(doc *goquery.Document) []ProxyIP {
	var ips []ProxyIP
	ipTable := doc.Find("table.table.table-bordered.table-striped").Find("tbody").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(0).Text()
		port, _ := strconv.Atoi(node.Find("td").Eq(1).Text())
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
		ipInfo["capture_time"] = time.Now()
		ips = append(ips, ipInfo)
	})
	return ips
}

func XiciPageParser(doc *goquery.Document) []ProxyIP {
	var ips []ProxyIP
	ipTable := doc.Find("table#ip_list").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(1).Text()
		port, _ := strconv.Atoi(node.Find("td").Eq(2).Text())
		proxyType := node.Find("td").Eq(5).Text()
		region := node.Find("td").Eq(3).Find("a").Text()
		rawTime := node.Find("td").Eq(9).Text()
		ipInfo := NewProxyIP()
		ipInfo["ip"] = ip
		ipInfo["port"] = port
		ipInfo["type"] = proxyType
		ipInfo["origin"] = "快代理"
		ipInfo["region"] = region
		ipInfo["raw_time"] = rawTime
		ipInfo["capture_time"] = time.Now()
		ips = append(ips, ipInfo)
	})
	return ips
}

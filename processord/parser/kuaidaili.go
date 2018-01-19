package parser

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/CX1ng/proxypool/model"

	"github.com/PuerkitoBio/goquery"
)

const (
	KuaiDaiLiUrl = "https://www.kuaidaili.com/free/inha/"
)

type WebSiteKuaiDaiLi struct {
	BaseUrl      string
	MaxPageNum   int
	BeginPageNum int
	channel      chan<- *model.ProxyIP
}

var once sync.Once
var kddStruct *WebSiteKuaiDaiLi

// NewKuaiDaiLi 实例使用单例模式进行创建，创建多个同一代理网站的实例并没什么用
func NewKuaiDaiLi(beginPageNum int, maxPageNum int, channel chan<- *model.ProxyIP) *WebSiteKuaiDaiLi {
	if beginPageNum <= 0 {
		beginPageNum = 1
	}
	once.Do(func() {
		kddStruct = &WebSiteKuaiDaiLi{
			BaseUrl:      KuaiDaiLiUrl,
			MaxPageNum:   maxPageNum,
			BeginPageNum: beginPageNum,
			channel:      channel,
		}
	})
	return kddStruct
}

// GetProxyIPList 迭代代理网站页面
func (w *WebSiteKuaiDaiLi) Exec() {
	for i := w.BeginPageNum; i < w.MaxPageNum; i++ {
		w.ParsePage(i)
		time.Sleep(10 * time.Second)
	}
}

// GetKuaiDaiLiIPList 解析“快代理”单页面，提取高匿IP
func (w *WebSiteKuaiDaiLi) ParsePage(pageNum int) {
	//后续请求连接使用net/http，配置header头
	doc, err := goquery.NewDocument(fmt.Sprintf("%s/%d/", KuaiDaiLiUrl, pageNum))
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}
	ipTable := doc.Find("table.table.table-bordered.table-striped").Find("tbody").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(0).Text()
		port, _ := strconv.Atoi(node.Find("td").Eq(1).Text())
		proxyType := node.Find("td").Eq(3).Text()
		region := node.Find("td").Eq(4).Text()
		rawTime := node.Find("td").Eq(6).Text()
		info := &model.ProxyIP{
			IP:      ip,
			Port:    port,
			Type:    proxyType,
			Origin:  "快代理",
			Region:  region,
			RawTime: rawTime,
		}
		w.channel <- info
		fmt.Printf("puts ip: %s to channel\n", ip)
	})
}

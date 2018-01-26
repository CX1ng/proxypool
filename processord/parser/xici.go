package parser

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/model"

	"github.com/PuerkitoBio/goquery"
)

type WebSiteXiCi struct {
	BaseUrl      string
	MaxPageNum   int
	BeginPageNum int
	channel      chan<- *model.ProxyIP
}

var xcOnce sync.Once
var xcStruct *WebSiteXiCi

func NewXiCi(beginPageNum int, maxPageNum int, channel chan<- *model.ProxyIP) *WebSiteXiCi {
	if beginPageNum <= 0 {
		beginPageNum = 1
	}
	xcOnce.Do(func() {
		xcStruct = &WebSiteXiCi{
			BaseUrl:      common.XiCi,
			MaxPageNum:   maxPageNum,
			BeginPageNum: beginPageNum,
			channel:      channel,
		}
	})
	return xcStruct
}

// Exec 迭代代理网站页面
func (w *WebSiteXiCi) Exec() {
	for i := w.BeginPageNum; i < w.MaxPageNum; i++ {
		w.ParsePage(i)
		time.Sleep(time.Duration(common.Config.Time.TimeInterval) * time.Second)
	}
}

// ParsePage 解析”西刺"单页面，提取高匿IP
func (w *WebSiteXiCi) ParsePage(pageNum int) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%d/", common.XiCi, pageNum), nil)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	request.Header.Add("User-Agent", common.UserAgent)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}
	ipTable := doc.Find("table#ip_list").Find("tr")
	ipTable.Each(func(index int, node *goquery.Selection) {
		ip := node.Find("td").Eq(1).Text()
		port, _ := strconv.Atoi(node.Find("td").Eq(2).Text())
		proxyType := node.Find("td").Eq(5).Text()
		region := node.Find("td").Eq(3).Find("a").Text()
		rawTime := node.Find("td").Eq(9).Text()
		info := &model.ProxyIP{
			IP:      ip,
			Port:    port,
			Type:    proxyType,
			Origin:  "西刺",
			Region:  region,
			RawTime: rawTime,
		}
		w.channel <- info
		fmt.Printf("puts ip: %s to channel\n", ip)
	})
}

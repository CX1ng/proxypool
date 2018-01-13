package parser

import (
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const (
	KuaiDaiLiUrl = "https://www.kuaidaili.com/free/inha/"
	MaxPageNum = 5
)

func GetKuaiDaiLiProxyIPList(){
	for i :=0;i < MaxPageNum;i++ {
		ParseKuaiDaiLiPage(i)
		time.Sleep(10 * time.Second)
	}
}

// GetKuaiDaiLiIPList 解析“快代理”单页面，提取高匿IP
func ParseKuaiDaiLiPage(pageNum int){
	//请求连接使用net/http，配置header头
	doc,err := goquery.NewDocument(fmt.Sprintf("%s/%d/",KuaiDaiLiUrl,pageNum))
	if err != nil {
		fmt.Printf("Error: %+v\n",err)
		return
	}
	ipTable := doc.Find("table.table.table-bordered.table-striped").Find("tbody").Find("tr")
	ipTable.Each(func(index int,node *goquery.Selection){
		ip := node.Find("td").Eq(0).Text()
		fmt.Printf("ip:%v\n",ip)
	})
}
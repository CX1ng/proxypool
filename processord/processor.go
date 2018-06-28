package processord

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
)

type Processor struct {
	Name         string
	beginPageNum int
	maxPageNum   int
	url          string
	parser       Parser
	queue        chan []models.ProxyIP
}

func NewProcessor(name string, beginPageNum, maxPageNum int, queue chan []models.ProxyIP) (*Processor, error) {
	url, ok := WebUrl[strings.ToLower(name)]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Proxy Web %s Not Support", name))
	}
	parser, ok := WebParser[strings.ToLower(name)]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Proxy Web %s Not Support", name))
	}

	// 获取解析器及baseUrl
	return &Processor{
		Name:         name,
		beginPageNum: beginPageNum,
		maxPageNum:   maxPageNum,
		url:          url,
		parser:       parser,
		queue:        queue,
	}, nil
}

func (p *Processor) Run() {
	for i := p.beginPageNum; i < p.maxPageNum; i++ {
		p.ParserPage(i)
		time.Sleep(time.Duration(Config.Time.TimeInterval) * time.Second)
	}
}

func (p *Processor) ParserPage(pageNum int) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%d/", p.url, pageNum), nil)
	if err != nil {
		return err
	}
	request.Header.Add("User-Agent", UserAgent)
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//后续请求连接使用net/http，配置header头
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	infoList := p.parser(doc)

	fmt.Printf("Processor:%+v\n", len(infoList))
	p.queue <- infoList

	return nil
}

package processord

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/processord/parser"
)

type Processor struct {
	url    string
	parser Parser
	queue  chan []models.ProxyIP
	config *webDetail
}

type webDetail struct {
	name         string
	beginPageNum int
	endPageNum   int
	timeInterval int
}

func NewProcessor(detail WebDetail, queue chan []models.ProxyIP) (*Processor, error) {
	parserSetter, ok := WebParsers.GetParserSetter(detail.Name)
	if !ok {
		return nil, ErrParserNotSupport
	}
	parser := parserSetter.SettingParser()

	// 获取解析器及baseUrl
	return &Processor{
		url:    parser.GetUrl(),
		parser: parser,
		queue:  queue,
		config: &webDetail{
			name:         detail.Name,
			beginPageNum: detail.BeginPageNum,
			endPageNum:   detail.EndPageNum,
			timeInterval: detail.TimeInterval,
		},
	}, nil
}

func (p *Processor) Run() {
	// TODO:最大页数不能超过网站最大页数
	for i := p.config.beginPageNum; i < p.config.endPageNum; i++ {
		p.ParserPage(i)
		time.Sleep(time.Duration(p.config.timeInterval) * time.Second)
	}
}

func (p *Processor) ParserPage(pageNum int) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%d/", p.url, pageNum), nil)
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
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	infoList := p.parser.PageParser(doc)

	fmt.Printf("Processor:%+v\n", len(infoList))
	p.queue <- infoList

	return nil
}

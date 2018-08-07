package processord

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/processord/parser"
	. "github.com/CX1ng/proxypool/utils"
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
	maxPage, err := p.parser.GetMaxPageNum(p.config.endPageNum)
	if err != nil {
		return
	}
	for i := p.config.beginPageNum; i <= maxPage; i++ {
		p.ParserPage(i)
		time.Sleep(time.Duration(p.config.timeInterval) * time.Second)
	}
}

func (p *Processor) ParserPage(pageNum int) error {
	html, err := HttpRequestWithUserAgent(fmt.Sprintf("%s%d/", p.url, pageNum))
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return err
	}
	infoList := p.parser.PageParser(doc)

	fmt.Printf("Processor:%+v\n", len(infoList))
	p.queue <- infoList

	return nil
}

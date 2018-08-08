package processord

import (
	"fmt"
	"strings"
	"time"

	. "github.com/CX1ng/proxypool/common"
	"github.com/CX1ng/proxypool/models"
	. "github.com/CX1ng/proxypool/processord/parser"
	. "github.com/CX1ng/proxypool/utils"
)

type Processor struct {
	url    string
	parser Parser
	queue  chan []models.ProxyIP
	detail WebDetail
}

func NewProcessor(detail WebDetail, queue chan []models.ProxyIP) (*Processor, error) {
	parserSetter, ok := WebParsers.GetParserSetter(detail.Name)
	if !ok {
		return nil, ErrParserNotSupport
	}
	parser := parserSetter.SettingParser()

	if err := detail.Validation(); err != nil {
		return nil, err
	}
	// 获取解析器及baseUrl
	return &Processor{
		url:    parser.GetUrl(),
		parser: parser,
		queue:  queue,
		detail: detail,
	}, nil
}

func (p *Processor) RunWithTaskType() {
	switch strings.ToLower(p.detail.TaskType) {
	case "once":
		p.OnceRun()
	case "loop":
		p.LoopRun()
	default:
		// 记录错误
	}
}

func (p *Processor) OnceRun() {
	maxPage, err := p.parser.GetMaxPageNum(p.detail.EndPageNum)
	if err != nil {
		return
	}
	for i := p.detail.BeginPageNum; i <= maxPage; i++ {
		p.ParserPage(i)
		time.Sleep(p.detail.TimeInterval.D())
	}
}

func (p *Processor) LoopRun() {
	for {
		p.OnceRun()
		time.Sleep(p.detail.LoopTimeInterval.D()) // 两次循环之间的时间间隔
	}
}

func (p *Processor) ParserPage(pageNum int) error {
	html, err := HttpRequestWithUserAgent(fmt.Sprintf("%s%d/", p.url, pageNum))
	if err != nil {
		return err
	}
	infoList, err := p.parser.PageParser(html)
	if err != nil {
		return err
	}

	fmt.Printf("Processor:%+v\n", len(infoList))
	p.queue <- infoList

	return nil
}

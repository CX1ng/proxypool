package parser

import (
	"github.com/PuerkitoBio/goquery"

	. "github.com/CX1ng/proxypool/models"
)

var WebParsers = NewParserSetterSet()

type ParserSetter interface {
	SettingParser() Parser
}

type Parser interface {
	PageParser(doc *goquery.Document) []ProxyIP
	GetUrl() string
	GetMaxPageNum(maxNum int) (int, error)
}

package parser

import (
	. "github.com/CX1ng/proxypool/models"
)

var WebParsers = NewParserSetterSet()

type ParserSetter interface {
	SettingParser() Parser
}

type Parser interface {
	PageParser(html string) ([]ProxyIP, error)
	GetUrl() string
	GetMaxPageNum(maxNum int) (int, error)
}

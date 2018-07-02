package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserSetterSet(t *testing.T) {
	as := assert.New(t)

	xici := &XiciSetter{}
	kdd := &KuaidailiSetter{}
	set := NewParserSetterSet()
	result := set.SetParserSetter("xici", xici)
	as.True(result)
	result = set.SetParserSetter("kuaidaili", kdd)
	as.True(result)
	result = set.SetParserSetter("kuaidaili", kdd)
	as.False(result)

	names := set.GetNames()
	as.Len(names, 2)
	as.Contains(names, "xici")
	as.Contains(names, "kuaidaili")

	parsers := set.GetParserSetters()
	as.Len(parsers, 2)
	as.Contains(parsers, kdd)
	as.Contains(parsers, xici)

	parser, result := set.GetParserSetter("xici")
	as.Equal(parser, xici)
	as.True(result)
	parser, result = set.GetParserSetter("kuaidaili")
	as.Equal(parser, kdd)
	as.True(result)
	parser, result = set.GetParserSetter("xx")
	as.False(result)
	as.Nil(parser)

	set.DelParserSetter("xici")
	parser, result = set.GetParserSetter("xici")
	as.False(result)
	as.Nil(parser)
	set.DelParserSetter("kuaidaili")
	parser, result = set.GetParserSetter("kuaidaili")
	as.False(result)
	as.Nil(parser)

}

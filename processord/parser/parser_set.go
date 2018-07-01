package parser

import (
	"strings"
)

type ParserSetterSet map[string]ParserSetter

func NewParserSetterSet() ParserSetterSet {
	return make(map[string]ParserSetter)
}

func (p ParserSetterSet) GetNames() []string {
	keys := make([]string, 0, len(p))
	for key, _ := range p {
		keys = append(keys, key)
	}
	return keys
}

func (p ParserSetterSet) GetParserSetters() []ParserSetter {
	setters := make([]ParserSetter, 0, len(p))
	for _, setter := range p {
		setters = append(setters, setter)
	}
	return setters
}

func (p ParserSetterSet) SetParserSetter(name string, setter ParserSetter) bool {
	name = strings.ToLower(name)
	_, ok := p[name]
	if ok {
		return false
	}
	p[name] = setter
	return true
}

func (p ParserSetterSet) DelParserSetter(name string) {
	name = strings.ToLower(name)
	delete(p, name)
}

func (p ParserSetterSet) GetParserSetter(name string) (ParserSetter, bool) {
	name = strings.ToLower(name)
	setter, ok := p[name]
	if !ok {
		return nil, false
	}
	return setter, true
}

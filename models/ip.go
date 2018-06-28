package models

import (
	. "github.com/CX1ng/proxypool/common"
)

type ProxyIP map[string]interface{}

func NewProxyIP() ProxyIP {
	proxyIPMap := make(map[string]interface{})
	for _, name := range ProxyIPFields {
		proxyIPMap[name] = nil
	}

	return proxyIPMap
}

func (p ProxyIP) GetFields() []string {
	keys := make([]string, 0, len(p))
	for key, _ := range p {
		keys = append(keys, key)
	}
	return keys
}

func (p ProxyIP) FieldsNum() int {
	return len(p)
}
